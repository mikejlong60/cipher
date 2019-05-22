package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Decryption program v0.01")
	key := []byte("passphrasewhichneedstobe32bytes!")

	ciphertext, err := ioutil.ReadFile("HIPAA.jpeg")
	if err != nil {
		fmt.Println(err)
	}

	//generates a new AES cipher using our 32 byte key
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	//GCM or Calois/Counter Mode is a mode of operation
	//for symmetric key cryptographic block ciphers.
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	//Creates a new byte array the size of the nonce which must be passed to Seal
	nonceSize := gcm.NonceSize()

	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(plaintext))

}
