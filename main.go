package main

import (
	"os"
	"atomicgo.dev/keyboard/keys"
	"github.com/pterm/pterm"
	"iod/functions"
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
	options = append(options, "5. List sorted files and dirs")
	options = append(options, "6. Quit")

	pterm.FgCyan.Printf("Current directory: %s\n", currentDir)
	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.Filter = false
	printer.TextStyle.Add(*pterm.NewStyle(pterm.FgGreen))
	printer.KeyConfirm = keys.Enter

	for {
		selectedOptions, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

		switch selectedOptions {
		case "1. List files and dirs":
			functions.ListFiles(currentDir)
		case "2. Change directory":
			currentDir = functions.ChangeDirectory(currentDir)
		case "3. Create new file":
			functions.CreateFile(currentDir)
		case "4. Delete file":
			functions.DeleteFile(currentDir)
		case "5. List sorted files and dirs":
			var sortBy []string

			sortBy = append(sortBy, "1. Sort by name")
			sortBy = append(sortBy, "2. Sort by change date")

			printer := pterm.DefaultInteractiveMultiselect.WithOptions(sortBy)
			printer.Filter = false
			printer.TextStyle.Add(*pterm.NewStyle(pterm.FgGreen))
			printer.KeyConfirm = keys.Enter

			selectedOptions, _ := pterm.DefaultInteractiveSelect.WithOptions(sortBy).Show()

			switch selectedOptions {
			case "1. Sort by name":
				functions.ListSortedFilesAndDirs(currentDir)
			case "2. Sort by change date":
				functions.ListSortedFilesAndDirsByChangeDate(currentDir)
			}
		case "6. Quit":
			return
		}
	}
}


