package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/opesun/goquery"
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

const (
	BufferSize = 4096
	IvSize     = 16
)

func sleepAndSay(sayWhat string, after time.Duration) {
	go func() {
		s := time.Now()
		time.Sleep(after)
		e := time.Now()

		fmt.Printf("[%s] - %s\n", e.Sub(s), sayWhat)

	}()
}

func grab() <-chan string {
	c := make(chan string)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				x, err := goquery.ParseUrl("http://vpustotu.ru/moderation/")
				if err == nil {
					if s := strings.TrimSpace(x.Find(".fi_text").Text()); s != "" {
						c <- s
					}
				}

				time.Sleep(2000 * time.Millisecond)
			}
		}()
	}
	fmt.Println("Spawned 10 go Grabber threads...")

	return c
}

func hashFile(file string) (string, error) {
	var md5hash string

	if exists(file) == false {
		return md5hash, errors.New("file does not exist")
	}

	f, err := os.Open(file)
	if err != nil {
		return md5hash, errors.New("cannot open file")
	}
	defer f.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, f); err != nil {
		return md5hash, err
	}

	md5hash = hex.EncodeToString(hash.Sum(nil))

	return md5hash, nil
}

func spawnWorkers(in chan string) {
	for index := 0; index < 5; index++ {
		go func() {
			for {
				select {
				case filePath := <-in:
					{
						fmt.Printf("Start hasing of [%s]\n", filePath)
						/*hash, err := hashFile(filePath)
						if err != nil {
							fmt.Printf("Error during hasing file:%s\n", filePath)
						} else {
							fmt.Printf("file=%s -> hash(%s)\n", filePath, hash)
						} */
					}
				}
			}
		}()
	}
}

func main() {

	bytes, err := ioutil.ReadFile("./123.torrent")
	if err != nil {
		os.Exit(1)
	}

	d, err := decode(string(bytes))
	if err != nil {
		os.Exit(1)
	}

	l := d["evg"]
	_ = l

	/*
		root := "c:\\tmp"

		//fileList := []string{}
		toProcessChan := make(chan string)
		spawnWorkers(toProcessChan)
		err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
			//fileList = append(fileList, path)
			if f.IsDir() {
				//fmt.Printf("DIR:%s\n", path)
			} else {
				//fmt.Printf("file=%s, size=%d\n", path, f.Size())
				if f.Size() > 1024*1024 {
					//fmt.Printf("File size bigger 1Mb pusing to pipe for processing==> %s", path)
					toProcessChan <- path
				}
			}

			return nil
		})

		if err != nil {
			panic(err)
		}

		fmt.Scanln()

			qChan := grab()
			ticker := time.NewTicker(500 * time.Millisecond)
			defer ticker.Stop()

			for i := 0; i < 10; i++ {
				quote := <-qChan

				fmt.Printf("----------------------------------\n%s\n---------------------------------\n", quote)
			}

			for {

				select {
				case <-ticker.C:
					{
						fmt.Println("tick")
					}
				case q := <-qChan:
					{
						fmt.Println(q)
					}
				}
			}

			fmt.Scanln()


				inFile, err := os.Open("c:\\Users\\ievgen_iukhymovych\\Downloads\\Net.Level_3.09.Winter_2017.zip")
				if err != nil {
					panic(err)
				}

				outFile, err := os.Create("c:\\temp\\Encypted.aes")
				if err != nil {
					panic(err)
				}

				iv := make([]byte, IvSize)
				_, err = rand.Read(iv)
				if err != nil {
					panic(err)
				}

				keyAes := []byte("1234567890098765")
				aes, err := aes.NewCipher(keyAes)
				if err != nil {
					panic(err)
				}

				ctr := cipher.NewCTR(aes, iv)

				buf := make([]byte, BufferSize)
				for {
					n, err := inFile.Read(buf)
					if err != nil && err != io.EOF {
						panic(err)
					}

					outBuf := make([]byte, n)
					ctr.XORKeyStream(outBuf, buf[:n])

					outFile.Write(outBuf)

					if err == io.EOF {
						break
					}
				}


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
