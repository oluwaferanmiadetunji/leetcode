package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <solution>")
		return
	}

	solution := os.Args[1]
	solutionDir := filepath.Join("solutions", solution)
	if _, err := os.Stat(solutionDir); os.IsNotExist(err) {
		fmt.Printf("Error: Directory %s does not exist\n", solutionDir)
		return
	}

	executable := "temp_executable"

	// Build the Go source file into an executable
	buildCmd := exec.Command("go", "build", "-o", executable, filepath.Join(solutionDir, "main.go"))
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		fmt.Printf("Error: Failed to compile %s\n", filepath.Join(solutionDir, "main.go"))
		fmt.Println(err)
		return
	}

	// Ensure the binary is executable
	if err := os.Chmod(executable, 0755); err != nil {
		fmt.Println("Error: Failed to set executable permissions")
		fmt.Println(err)
		return
	}

	// Check if the executable file exists
	if _, err := os.Stat(executable); os.IsNotExist(err) {
		fmt.Println("Error: Executable file does not exist")
		return
	}

	// Run the compiled executable
	runCmd := exec.Command("./" + executable)
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	if err := runCmd.Run(); err != nil {
		fmt.Printf("Error: Failed to run %s\n", executable)
		fmt.Println(err)
		return
	}

	// Clean up the executable file
	if err := os.Remove(executable); err != nil {
		fmt.Printf("Error: Failed to delete temporary executable\n")
		fmt.Println(err)
		return
	}
}
