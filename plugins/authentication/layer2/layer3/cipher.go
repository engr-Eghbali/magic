package magic

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"log"
)

const keystr = "fuckoffyoupity,mortalcreaturs!!!"

func Cipher(str string) (ciphertxt []byte) {
	text := []byte(str)

	//key from file**	content, err := ioutil.ReadFile("/home/reza/Desktop/go/server/pluginRZA/authentication/layer2/layer3/checker/key.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	key := []byte(keystr)
	ciphertext, err := encrypt(text, key)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	return ciphertext
}
func Decipher(shit []byte) (txt string) {
	//	content, err := ioutil.ReadFile("/home/reza/Desktop/go/server/pluginRZA/authentication/layer2/layer3/checker/key.txt")
	//	if err != nil {
	//		log.Fatal(err)
	//}
	key := []byte(keystr)

	plaintext, err := decrypt(shit, key)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	return string(plaintext[:len(plaintext)])
}

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
