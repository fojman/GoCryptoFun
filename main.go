package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"os"
)

/*
	File encryptor:
	- load RSA from file
	- decrypt AES crypt key
	- using AES encrypt/decrypt file

	it's a multi tool it can:
	- generate PKI pari
	- crypt/decrypt file(or input string)
*/

// default mode is Crypt
const (
	Crypt = iota
	Generate
	Invalid
)

func createAes(buf []byte, aesKey string) ([]byte, int) {

	// Byte array of the string
	plaintext := buf

	// Key
	key := []byte(aesKey)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Empty array of 16 + plaintext length
	// Include the IV at the beginning
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// Slice of first 16 bytes
	iv := ciphertext[:aes.BlockSize]

	// Write 16 rand bytes to fill iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from plaintext to ciphertext
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return nil, string(ciphertext)
}

func aesEncryptBuf(inBuf []byte, len int) ([]byte, int) {

}

func main() {

	//mode := flag.String("mode", "generate", "Utility working modes: crypt|generate, default: generate")
	//dir := flag.String("dir", "./", "Utility ouput direcatory, i.e. for storing keys, etc")

	srcFile := flag.String("sfile", "", "source file location")
	//dstFile := flag.String("dfile", "", "dest file location")

	// parse command line
	flag.Parse()
	if false == exists(*srcFile) {
		fmt.Println("ERR: source file does not exit")
		flag.PrintDefaults()
		os.Exit(1)
	}

	path, err := processFile(*srcFile, "0123456789012345")
	if err != nil {
		panic(err)
	}

	fmt.Printf("file encrypted to: %s", path)

	/*
		if *mode != "gen" {
			flag.PrintDefaults()
			os.Exit(1)
		}

		ensureExists(*dir)

		fmt.Println("Generating RSA PKI keys...")

		priv, pub := generateKeys()

		fmt.Printf("Storing keys at: %s\n", *dir)

		privStr := rsaPrivKeyToPemString(priv, "string_in")
		pubStr, _ := rsaPubKeyToPemString(pub)

		privPath := filepath.Join(*dir, "private.key")
		pubPath := filepath.Join(*dir, "public.key")
		writeStringToFile(privStr, privPath)
		writeStringToFile(pubStr, pubPath) */

	fmt.Println("Keys generated..")

	os.Exit(0)
}
