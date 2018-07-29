package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// default rsa key len
const (
	RsaDefaultKeyLenght = 2048
)

func generateKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	pk, _ := rsa.GenerateKey(rand.Reader, RsaDefaultKeyLenght)

	return pk, &pk.PublicKey
}

// save rsa pk
func rsaPrivKeyToPemString(pk *rsa.PrivateKey, where string) string {
	pkBytes := x509.MarshalPKCS1PrivateKey(pk)
	pkBytesPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: pkBytes,
		},
	)

	return string(pkBytesPem)
}

func rsaPubKeyToPemString(pubKey *rsa.PublicKey) (string, error) {
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return "", err
	}
	pubKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubKeyBytes,
		},
	)
	return string(pubKeyPem), nil
}

func parseRsaPrivKeyFromPemString(privPem string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPem))
	if block == nil {
		return nil, errors.New("Failed to parse PEM block containing RSA Priv key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func parseRsaPubKeyFromPemString(pubPem string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPem))
	if block == nil {
		return nil, errors.New("cannot parse pub key from PEM")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break
	}

	return nil, errors.New("Key type is not RSA.PUB")
}
