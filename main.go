package main

import (
	"crypto/aes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
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

func encrypteFile(file string, aesKey string) {

	key := []byte(aesKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// IV

}

func main() {

	mode := flag.String("mode", "generate", "Utility working modes: crypt|generate, default: generate")
	dir := flag.String("dir", "./", "Utility ouput direcatory, i.e. for storing keys, etc")
	// parse command line
	flag.Parse()

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
	writeStringToFile(pubStr, pubPath)

	fmt.Println("Keys generated..")
}
