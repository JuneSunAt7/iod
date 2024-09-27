package functions
import (
	"os"
	"sort"
	"github.com/pterm/pterm"
)

func ListSortedFilesAndDirs(dir string) { // sort by name
	files, err := os.ReadDir(dir)
	if err != nil {
		pterm.Error.Println("Error listing files:", err)
		return
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		pterm.FgMagenta.Println(file.Name())
	}
}