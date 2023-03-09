package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/google/uuid"
)

// HashCookie возвращает набор байтов из случайной захэшированной строки
func HashCookie() ([]byte, error) {
	var id = uuid.New()

	key, err := GenerateRandom(2 * aes.BlockSize) // ключ шифрования
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return []byte{}, err

	}

	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return []byte{}, err

	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return []byte{}, err

	}

	// создаём вектор инициализации
	nonce, err := GenerateRandom(aesgcm.NonceSize())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return []byte{}, err

	}

	return aesgcm.Seal(nil, nonce, []byte(id.String()), nil), nil
}

// GenerateRandom генерирует случайный набор байтов
func GenerateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
