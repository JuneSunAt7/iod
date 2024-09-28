package functions

import (
	"os"
	"github.com/pterm/pterm"
	"path/filepath"
	"bufio"
	"fmt"
)

// open file and show content
func OpenFile(dir string) {
	pterm.FgCyan.Println("Enter file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	if _, err := os.Stat(filepath.Join(dir, fileName)); os.IsNotExist(err) {
		pterm.Error.Println("File does not exist.")
		return
	}

	file, err := os.Open(filepath.Join(dir, fileName))
	if err != nil {
		pterm.Error.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pterm.FgGreen.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		pterm.Error.Println("Error reading file:", err)
	}
}