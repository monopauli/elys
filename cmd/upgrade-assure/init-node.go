package main

import (
	"log"
	"os/exec"
)

func initNode(cmdPath, moniker, chainId, homePath string) {
	// Command and arguments
	args := []string{"init", moniker, "--chain-id", chainId, "--home", homePath}

	// Execute the command
	if err := exec.Command(cmdPath, args...).Run(); err != nil {
		log.Fatalf(ColorRed+"Command execution failed: %v", err)
	}

	// If execution reaches here, the command was successful
	log.Printf(ColorYellow+"init node with moniker %s, chain id %s and home path: %s successfully", moniker, chainId, homePath)
}
