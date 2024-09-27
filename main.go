package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"atomicgo.dev/keyboard/keys"
	"github.com/pterm/pterm"
)

// main runs an interactive shell allowing the user to list the files and directories in
// their current directory, change directories, create new files, delete files, and quit.
func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		pterm.Error.Println("Error getting current directory:", err)
		return
	}
	var options []string

	options = append(options, "1. List files and dirs")
	options = append(options, "2. Change directory")
	options = append(options, "3. Create new file")
	options = append(options, "4. Delete file")
	options = append(options, "5. Quit")

	pterm.FgCyan.Printf("Current directory: %s\n", currentDir)
	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.Filter = false
	printer.TextStyle.Add(*pterm.NewStyle(pterm.FgGreen))
	printer.KeyConfirm = keys.Enter

	for {
		selectedOptions, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

		switch selectedOptions {
		case "1. List files and dirs":
			listFiles(currentDir)
		case "2. Change directory":
			currentDir = changeDirectory(currentDir)
		case "3. Create new file":
			createFile(currentDir)
		case "4. Delete file":
			deleteFile(currentDir)
		case "5. Quit":
			return
		}
	}
}

func listFiles(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		pterm.Error.Println("Error listing files:", err)
		return
	}

	for _, file := range files {
		pterm.FgMagenta.Println(file.Name())
	}
}


func changeDirectory(currentDir string) string {
	pterm.FgCyan.Println("Enter new directory: ")
	var newDir string
	fmt.Scanln(&newDir)

	if newDir == ".." {
		return filepath.Dir(currentDir)
	}

	if !strings.HasPrefix(newDir, "/") {
		newDir = filepath.Join(currentDir, newDir)
	}

	if _, err := os.Stat(newDir); os.IsNotExist(err) {
		pterm.Error.Println("Directory does not exist.")
		return currentDir
	}

	return newDir
}

func createFile(dir string) {
	pterm.FgCyan.Println("Enter file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	if _, err := os.Stat(filepath.Join(dir, fileName)); !os.IsNotExist(err) {
		pterm.Warning.Println("File already exists.")
		return
	}

	file, err := os.Create(filepath.Join(dir, fileName))
	if err != nil {
		pterm.Error.Println("Error creating file:", err)
		return
	}

	defer file.Close()
}

func deleteFile(dir string) {
	pterm.FgCyan.Println("Enter file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	if err := os.Remove(filepath.Join(dir, fileName)); err != nil {
		pterm.Error.Println("Error deleting file:", err)
		return
	}
}