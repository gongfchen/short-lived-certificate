package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main_decrypt(domain_name string, num_of_cert int) {
	keypath := "../../CA-middle-daemon-storage/Daily Keys/" + domain_name + "/daily_key.txt"
	// count := 0
	// for count < num_of_cert {
	// Read entire file content
	daily_key, err := ioutil.ReadFile(keypath)
	check(err)

	// err = os.Remove(keypath)
	// check(err)

	// Convert []byte to string
	key := string(daily_key)

	// Read entire file content
	cipher, err := ioutil.ReadFile("../../CA-middle-daemon-storage/Encrypted Certificates/" + domain_name + "/enc_cert" + strconv.Itoa(num_of_cert) + ".txt")
	check(err)

	// Convert []byte to string
	encrypted := string(cipher)

	decrypted := decrypt(encrypted, key)

	certpath := "../../middle-daemon-website-daemon-storage/Daily Certificates/" + domain_name
	if _, err := os.Stat(certpath); os.IsNotExist(err) {
		os.Mkdir(certpath, 0700)
	}

	f, err := os.Create(certpath + "/certday" + strconv.Itoa(num_of_cert) + ".pem")
	check(err)
	l, err := f.WriteString(decrypted)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully --- Decrypted certificate written in shared storage")
	err = f.Close()
	check(err)
	// count += 1
	// }

}

func decrypt(encryptedString string, keyString string) (decryptedString string) {

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	check(err)

	return fmt.Sprintf("%s", plaintext)
}
