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
	"golang.org/x/sys/windows/registry"
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
		pterm.Error.Println("Error decrypting file:", err)
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
func CreateSettingsToRegedit(settingsName string, settingValue []byte) {
    k, err := registry.OpenKey(registry.CURRENT_USER, `Software\iod`, registry.ALL_ACCESS)
    if err != nil {
        k, _, err = registry.CreateKey(registry.CURRENT_USER, `Software\iod`, registry.ALL_ACCESS)
        if err != nil {
            return
        }
    }

    defer k.Close()

    err = k.SetBinaryValue(settingsName, settingValue)
    if err != nil {
        return
    }
}
func ReadRegistryValue(key registry.Key, subKey string, valueName string) ([]byte, error) {
    k, err := registry.OpenKey(key, subKey, registry.QUERY_VALUE)
    if err != nil {
        return nil, err
    }
    defer k.Close()

    val, _, err := k.GetBinaryValue(valueName)
    if err != nil {
        return nil, err
    }

    return val, nil
}
func CheckAviailableKey() bool {
	_, err := ReadRegistryValue(registry.CURRENT_USER, `Software\iod`, "key")
	if err != nil {
		return false
	}
	return true
}
func DeleteUserOriginalFile(path string) {
	if err := os.Remove(path); err != nil {
		pterm.Error.Println("Error deleting file:", err)
		return
	}
}