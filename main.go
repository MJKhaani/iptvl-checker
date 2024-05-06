package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"syscall"
)

func isPortOpen(port string) bool {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func killProcessByName(processName string) error {
	cmd := exec.Command("taskkill", "/F", "/IM", processName)
	return cmd.Run()
}

func main() {
	port := "8888" // Change this to the port you want to check
	processName := "ipvtl.dll"

	if !isPortOpen(port) {
		fmt.Printf("Port %s is not open. Killing process %s...\n", port, processName)
		if err := killProcessByName(processName); err != nil {
			fmt.Printf("Error killing process %s: %v\n", processName, err)
			//os.Exit(1)
		}
		fmt.Printf("Starting ipvtl_16ch_trial.exe...\n")
		cmd := exec.Command("./ipvtl_16ch_trial.exe")
		cmd.Dir = "."                                            // Set the directory where ipvtl_16ch_trial.exe is located
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // Start the process hidden
		if err := cmd.Start(); err != nil {
			fmt.Printf("Error starting ipvtl_16ch_trial.exe: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("ipvtl_16ch_trial.exe started successfully.")
	} else {
		fmt.Printf("Port %s is already open. No action taken.\n", port)
	}
}
