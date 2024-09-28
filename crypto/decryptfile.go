package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
	"path/filepath"
	"github.com/pterm/pterm"
	"fmt"
	"io/ioutil"
)

func DecryptFile(dir string, key string) error {

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

	iv := make([]byte, aes.BlockSize)

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		pterm.Error.Println("Error reading IV:", err)
		return err
	}

	// Create a new Cipher Block from the key
	blockCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		pterm.Error.Println("Error creating block:", err)
		return err
	}

	mode := cipher.NewCBCDecrypter(blockCipher, iv)
// Read the encrypted data from the file
	encryptedBytes, err := ioutil.ReadAll(file)
	if err != nil {
    	pterm.Error.Println("Error reading file:", err)
    	return err
	}
	// Create a new byte array the size of the original file
	decryptedBytes := make([]byte, len(encryptedBytes))

	// Decrypt the file
	mode.CryptBlocks(decryptedBytes, encryptedBytes)

	// Write the decrypted file to disk

	pterm.FgCyan.Println("Enter new file name: ")
	var newFileName string
	fmt.Scanln(&newFileName)
	err = os.WriteFile(filepath.Join(dir, newFileName), decryptedBytes, 0644)
	if err != nil {
		pterm.Error.Println("Error writing decrypted file:", err)
		return err
	}

	pterm.Success.Println("File decrypted successfully.")

	return nil
}