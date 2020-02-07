package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	deleteCmd := exec.Command("kind", "delete", "cluster")
	deleteCmd.Stdout = os.Stdout
	deleteCmd.Stderr = os.Stderr
	defer deleteCmd.Run()

	createCmd := exec.Command("kind", "create", "cluster")
	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr

	err := createCmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Some test code in here

}
