package functions

// find file in computer
import (
	"path/filepath"
	"github.com/pterm/pterm"
	"fmt"
)

func GlobalFinderFile() {
	pterm.FgCyan.Println("Enter the pattern(example: *.txt): ")
	var fileName string
	fmt.Scanln(&fileName)

	files, err := filepath.Glob(fileName)
	if err != nil {
		pterm.Error.Println("Error finding file:", err)
		return
	}
	if len(files) == 0 {
		pterm.Warning.Println("No files found")
		return
	}
	for _, file := range files {
		pterm.FgMagenta.Println(file)
	}
}

