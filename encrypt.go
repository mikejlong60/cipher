package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	fmt.Println("Encryption program v0.01")
	//text := []byte("you are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great juliayou are great julia")
	key := []byte("passphrasewhichneedstobe32bytes!")

	text, err := ioutil.ReadFile("otis.txt") //"HIPAA.jpeg")
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
	nonce := make([]byte, gcm.NonceSize())
	//populates our nonce with a cryptographically secure random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	//Here you are encrypting the text using the Seal function.
	//Seal encrypts and authenticates plaintext, authenticates the
	//additional data and appends the result to dst, returning the updated
	//slice.  The nonce must be NonceSize() bytes long and unique for all time,
	//for a given key.
	//	fmt.Println(gcm.Seal(nonce, nonce, text, nil))

	//The WriteFile method returns an error if unsuccessful
	err = ioutil.WriteFile("encryptedhipaa", gcm.Seal(nonce, nonce, text, nil), 0777)

	if err != nil {
		fmt.Println(err)
	}
}
