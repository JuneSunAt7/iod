package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

func DecryptFile(key []byte, in *os.File, out *os.File) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	nonceSize := gcm.NonceSize()
	nonce, err := io.ReadAll(io.LimitReader(rand.Reader, int64(nonceSize)))
	if err != nil {
		return err
	}
	ciphertext, err := io.ReadAll(in)
	if err != nil {
		return err
	}
	if len(ciphertext) < nonceSize {
		return err
	}
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}
	_, err = out.Write(plaintext)
	return err
}