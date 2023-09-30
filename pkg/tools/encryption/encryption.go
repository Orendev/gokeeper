package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base32"

	"github.com/pkg/errors"
)

var ErrOpen = errors.Errorf("cipher: message authentication failed")

const Size = 32

type Enc struct {
	Key [Size]byte
}

func New(passphrase string) *Enc {
	return &Enc{Key: sha256.Sum256([]byte(passphrase))}
}

func (e *Enc) Encrypt(msg string) (string, error) {
	src := []byte(msg)

	block, err := aes.NewCipher(e.Key[:])
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce, err := e.getNonce(gcm.NonceSize())
	if err != nil {
		return "", err
	}

	return base32.StdEncoding.EncodeToString(gcm.Seal(nil, nonce, src, nil)), nil
}

func (e *Enc) Decrypt(msg string) (string, error) {
	src, err := base32.StdEncoding.DecodeString(msg)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(e.Key[:])
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(src) < nonceSize {
		return "", ErrOpen
	}

	nonce, err := e.getNonce(gcm.NonceSize())
	if err != nil {
		return "", err
	}

	ret, err := gcm.Open(nil, nonce, src, nil)
	if err != nil {
		return "", err
	}

	return string(ret), nil

}

func (e *Enc) getNonce(size int) ([]byte, error) {
	if len(e.Key) < size {
		return nil, ErrOpen
	}
	return e.Key[len(e.Key)-size:], nil
}
