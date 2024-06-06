package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// 終止 explorer.exe 進程
	cmd := exec.Command("taskkill", "/f", "/im", "explorer.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error terminating explorer.exe: %v\n", err)
		return
	}

	// 重啟 explorer.exe 進程
	cmd = exec.Command("start", "explorer.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error starting explorer.exe: %v\n", err)
		return
	}

	fmt.Println("Explorer.exe has been restarted successfully.")
}
