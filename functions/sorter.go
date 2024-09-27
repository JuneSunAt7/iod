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

    for i := 0; i < len(files); i++ {
        pterm.FgMagenta.Println(files[i].Name())
    }
}
// sort files & dirs by change date

func ListSortedFilesAndDirsByChangeDate(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		pterm.Error.Println("Error listing files:", err)
		return
	}

	sort.Slice(files, func(i, j int) bool {
		infoI, err := files[i].Info()
		if err != nil {
			pterm.Error.Println("Error getting file info:", err)
		}
		infoJ, err := files[j].Info()
		if err != nil {
			pterm.Error.Println("Error getting file info:", err)
		}
		return infoI.ModTime().After(infoJ.ModTime())
	})

	for _, file := range files {
		pterm.FgMagenta.Println(file.Name())
	}
}