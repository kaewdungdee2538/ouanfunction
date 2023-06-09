package md5

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

const CodeDecryptedFailed = 33006
const MessageDecryptedFailedTh = "ถอดรหัสล้มเหลว"
const MessageDecryptedFailedEn = "Decryped failed"

const saltlen = 8
const keylen = 32
const iterations = 10002

func EncryptMD5(plaintext string, password string) (string, error) {
	header := make([]byte, saltlen+aes.BlockSize)

	salt := header[:saltlen]
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}

	iv := header[saltlen : aes.BlockSize+saltlen]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	key := pbkdf2.Key([]byte(password), salt, iterations, keylen, md5.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(header)+len(plaintext))
	copy(ciphertext, header)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize+saltlen:], []byte(plaintext))
	return base64Encode(ciphertext), nil
}

func DecryptMD5(encrypted string, password string) (string, error) {
	a, err := base64Decode([]byte(encrypted))
	if err != nil {
		return "", err
	}
	ciphertext := a
	salt := ciphertext[:saltlen]
	iv := ciphertext[saltlen : aes.BlockSize+saltlen]
	key := pbkdf2.Key([]byte(password), salt, iterations, keylen, md5.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New(MessageDecryptedFailedEn)
	}

	decrypted := ciphertext[saltlen+aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decrypted, decrypted)

	return string(decrypted), nil
}

func base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
