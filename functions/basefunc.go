package functions
import(
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/pterm/pterm"
)
func ListFiles(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		pterm.Error.Println("Error listing files:", err)
		return
	}

	for _, file := range files {
		pterm.FgMagenta.Println(file.Name())
	}
}


func ChangeDirectory(currentDir string) string {
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

func CreateFile(dir string) {
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

func DeleteFile(dir string) {
	pterm.FgCyan.Println("Enter file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	if err := os.Remove(filepath.Join(dir, fileName)); err != nil {
		pterm.Error.Println("Error deleting file:", err)
		return
	}
}