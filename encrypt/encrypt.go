package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

const (
	KEY_LENGTH = 32
)

type Encrypt struct {
	Key []byte
}

func New(key []byte) (*Encrypt, error) {
	if len(key) != KEY_LENGTH {
		return nil, errors.New("Invalid key length")
	}

	return &Encrypt{Key: []byte(key)}, nil
}

func (e *Encrypt) UpdateKey(key string) error {
	if len(key) != KEY_LENGTH {
		return errors.New("Invalid key length")
	}

	e.Key = []byte(key)
	return nil
}

func (e *Encrypt) Encrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher(e.Key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nonce, nonce, src, nil)

	return cipherText, nil
}

func (e *Encrypt) Decrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher(e.Key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce := src[:nonceSize]
	cipher := src[nonceSize:]

	decrypted, err := gcm.Open(nil, []byte(nonce), []byte(cipher), nil)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
