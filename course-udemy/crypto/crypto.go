package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func doCrypt() {
	block, str, iv, encrypted := encrypt("key4567890keykey", "cipher7890cipher", "Encrypt this!")
	decrypted := decrypt(block, str, iv, encrypted)
	fmt.Println(decrypted)
}

func encrypt(key string, cipherTextRaw string, textToEncrypt string) (cipher.Block, []byte, []byte, []uint8) {
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		panic(err)
	}

	str := []byte(textToEncrypt)

	/*
		16 bytes = AES-128
		24 bytes = AES-192
		32 bytes = AES-256
	*/

	cipherText := []byte(cipherTextRaw)
	iv := cipherText[:aes.BlockSize]
	encrypter := cipher.NewCFBEncrypter(block, iv)

	encrypted := make([]byte, len(str))
	encrypter.XORKeyStream(encrypted, str)

	fmt.Printf("%s encrypted to %v\n", str, encrypted)
	return block, str, iv, encrypted
}

func decrypt(block cipher.Block, str []byte, iv []byte, encrypted []uint8) string {
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypted := make([]byte, len(str))
	decrypter.XORKeyStream(decrypted, encrypted)
	fmt.Printf("%v decrypt to %s\n", encrypted, decrypted)
	return string(decrypted)
}
