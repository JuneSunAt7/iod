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
    "encoding/hex"
    "crypto/sha256"
)
// EncryptFile encrypts a file using the provided key
func EncryptFile(key []byte, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	return append(nonce, ciphertext...), nil
}

// DecryptFile decrypts a file using the provided key
func DecryptFile(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce := ciphertext[:nonceSize]
	ciphertext = ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func EncryptFileTUI(dir string, key []byte) error {
    pterm.FgCyan.Println("Enter file name: ")
    var fileName string
    fmt.Scanln(&fileName)

    if _, err := os.Stat(filepath.Join(dir, fileName)); os.IsNotExist(err) {
        pterm.Error.Println("File does not exist.")
        return err
    }

    pterm.FgCyan.Println("Enter new file name: ")
    var newFileName string
    fmt.Scanln(&newFileName)

    newFile, err := os.Create(filepath.Join(dir, newFileName))
    if err != nil {
        pterm.Error.Println("Error creating file:", err)
        return err
    }
    defer newFile.Close()
    plaintext, err := os.ReadFile(filepath.Join(dir, fileName))
	if err != nil {
		return err
	}
    ciphertext, err := EncryptFile(key, plaintext)
	if err != nil {
		return err
	}
    DeleteUserOriginalFile(filepath.Join(dir, fileName))
    pterm.Success.Println("File encrypted successfully.")
	return os.WriteFile(filepath.Join(dir, newFileName), ciphertext, 0644)
}


func DecryptFileTUI(dir string, key []byte) error {
    pterm.FgCyan.Println("Enter file name: ")
    var fileName string
    fmt.Scanln(&fileName)

    pterm.FgCyan.Println("Enter new decrypted file name: ")
    var plaintextPath string
    fmt.Scanln(&plaintextPath)

    if _, err := os.Stat(filepath.Join(dir, fileName)); os.IsNotExist(err) {
        pterm.Error.Println("File does not exist.")
        return err
    }

    ciphertext, err := os.ReadFile(filepath.Join(dir, fileName))
	if err != nil {
		return err
	}

	plaintext, err := DecryptFile(key, ciphertext)
	if err != nil {
		return err
	}
    DeleteUserOriginalFile(filepath.Join(dir, fileName))
    pterm.Success.Println("File decrypted successfully.")
	return os.WriteFile(plaintextPath, plaintext, 0644)

}
func GenerateKey() []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		pterm.Error.Println("Error generating key:", err)
	}
	return key
}
func HashKey(key []byte) string {
	hash := sha256.Sum256(key)
	return hex.EncodeToString(hash[:])
}
func SaveKeyToRegedit(key string) {
	// Save the key to the registry
	if err := os.WriteFile("key.txt", []byte(key), 0644); err != nil {
		pterm.Error.Println("Error saving key to registry:", err)
		return
	}
}
func ReadKeyFromRegedit() []byte {
	// Read the key from the registry
	keyBytes, err := os.ReadFile("key.txt")
	if err != nil {
		pterm.Error.Println("Error reading key from registry:", err)
		return nil
	}
	return keyBytes
}

func DeleteUserOriginalFile(path string) {
	if err := os.Remove(path); err != nil {
		pterm.Error.Println("Error deleting file:", err)
		return
	}
}