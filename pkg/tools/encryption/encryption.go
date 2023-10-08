package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"

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

func (e *Enc) EncryptByte(src []byte) ([]byte, error) {
	block, err := aes.NewCipher(e.Key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce, err := e.getNonce(gcm.NonceSize())
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nil, nonce, src, nil), nil
}

func (e *Enc) DecryptByte(src []byte) ([]byte, error) {

	block, err := aes.NewCipher(e.Key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(src) < nonceSize {
		return nil, ErrOpen
	}

	nonce, err := e.getNonce(gcm.NonceSize())
	if err != nil {
		return nil, err
	}

	ret, err := gcm.Open(nil, nonce, src, nil)
	if err != nil {
		return nil, err
	}

	return ret, nil

}

func (e *Enc) getNonce(size int) ([]byte, error) {
	if len(e.Key) < size {
		return nil, ErrOpen
	}
	return e.Key[len(e.Key)-size:], nil
}
