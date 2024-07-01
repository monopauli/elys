package app

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	m "github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	consumertypes "github.com/cosmos/interchain-security/v4/x/ccv/consumer/types"
)

func SetupHandlers(app *ElysApp) {
	setUpgradeHandler(app)
	loadUpgradeStore(app)
}

func setUpgradeHandler(app *ElysApp) {
	upgradeName := "v999.999.999" // Make sure this matches your upgrade plan's name
	app.UpgradeKeeper.SetUpgradeHandler(
		upgradeName,
		func(ctx sdk.Context, plan upgradetypes.Plan, vm m.VersionMap) (m.VersionMap, error) {
			app.Logger().Info("Running upgrade handler for " + upgradeName)

			// Initialize new stores or modules here
			// Add any other initialization logic needed for the upgrade

			return app.mm.RunMigrations(ctx, app.configurator, vm)
		},
	)
}

func loadUpgradeStore(app *ElysApp) {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("Failed to read upgrade info from disk: %v", err))
	}

	fmt.Printf("Upgrade info: %+v\n", upgradeInfo)

	if shouldLoadUpgradeStore(app, upgradeInfo) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{consumertypes.ModuleName},
			// Deleted: []string{},
		}
		fmt.Printf("Setting store loader with height %d and store upgrades: %+v\n", upgradeInfo.Height, storeUpgrades)

		// Use upgrade store loader for the initial loading of all stores when app starts,
		// it checks if version == upgradeHeight and applies store upgrades before loading the stores,
		// so that new stores start with the correct version (the current height of chain),
		// instead the default which is the latest version that store last committed i.e 0 for new stores.
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	} else {
		fmt.Println("No need to load upgrade store.")
	}
}

func shouldLoadUpgradeStore(app *ElysApp, upgradeInfo upgradetypes.Plan) bool {
	currentHeight := app.LastBlockHeight()
	fmt.Printf("Current block height: %d, Upgrade height: %d\n", currentHeight, upgradeInfo.Height)
	return upgradeInfo.Name == "v999.999.999" && currentHeight == upgradeInfo.Height && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height)
}
