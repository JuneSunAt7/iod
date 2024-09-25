package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	for {
		fmt.Printf("Current directory: %s\n", currentDir)
		fmt.Println("1. List files and dirs")
		fmt.Println("2. Change directory")
		fmt.Println("3. Create new file")
		fmt.Println("4. Delete file")
		fmt.Println("5. Quit")

		var choice string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			listFiles(currentDir)
		case "2":
			currentDir = changeDirectory(currentDir)
		case "3":
			createFile(currentDir)
		case "4":
			deleteFile(currentDir)
		case "5":
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func listFiles(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error listing files:", err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func changeDirectory(currentDir string) string {
	fmt.Print("Enter new directory: ")
	var newDir string
	fmt.Scanln(&newDir)

	if newDir == ".." {
		return filepath.Dir(currentDir)
	}

	if !strings.HasPrefix(newDir, "/") {
		newDir = filepath.Join(currentDir, newDir)
	}

	if _, err := os.Stat(newDir); os.IsNotExist(err) {
		fmt.Println("Directory does not exist.")
		return currentDir
	}

	return newDir
}

func createFile(dir string) {
	fmt.Print("Enter file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	if _, err := os.Stat(filepath.Join(dir, fileName)); !os.IsNotExist(err) {
		fmt.Println("File already exists.")
		return
	}

	file, err := os.Create(filepath.Join(dir, fileName))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()
}

func deleteFile(dir string) {
	fmt.Print("Enter file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	if err := os.Remove(filepath.Join(dir, fileName)); err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
}