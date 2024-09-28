package crypto
// crypt file in AES 

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
	"path/filepath"
	"github.com/pterm/pterm"
	"fmt"
)

func EncryptFile(dir string, key string) error {
	pterm.FgCyan.Println("Enter file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	if _, err := os.Stat(filepath.Join(dir, fileName)); os.IsNotExist(err) {
		pterm.Error.Println("File does not exist.")
		return err
	}

	file, err := os.Open(filepath.Join(dir, fileName))
	if err != nil {
		pterm.Error.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	pterm.FgCyan.Println("Enter new file name: ")
	var newFileName string
	fmt.Scanln(&newFileName)

	newFile, err := os.Create(filepath.Join(dir, newFileName))
	if err != nil {
		pterm.Error.Println("Error creating file:", err)
		return err
	}

	defer newFile.Close()

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		pterm.Error.Println("Error creating cipher:", err)
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		pterm.Error.Println("Error creating GCM:", err)
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		pterm.Error.Println("Error reading nonce:", err)
		return err
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(fileName), nil)
	if _, err := newFile.Write(ciphertext); err != nil {
		pterm.Error.Println("Error writing to file:", err)
		return err
	}

	pterm.Success.Println("File encrypted and saved as", newFileName)

	return nil
}
   