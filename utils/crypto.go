package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	b64 "encoding/base64"
	"fmt"
	"io"
	"log"
)

var PubKey *rsa.PublicKey

func init() {
	raw, err := b64.StdEncoding.DecodeString("MIIBCgKCAQEApvxieEC28whOn2Qi++uxVDM6m2spQHWoAZBz2WiqauAs9tpgeLuQgKBxpK4W+qm0oVksXXMKE5aDYdmvVz4X++S4KkElTJutg7VJzXVSlIOlCIRR6UyPmPQYmdGWx+gf+LBd6HbhAKbSzJBDQ7j4MW4dqjbuKBxqbJQcg13vnvefgIog9XD3k5eA+dkLk9DuMSZDd24ACcBtiwEj2ATy6ccEY4VXGk+AUWA6VowlcUTTNC7Yp+19QgpVpWIpJ8gcIptdBJ9TDbP0/ST48niqQQDWLeXsa0Wg+ZR+/uTV36aMqYQuf4xHrm3hhroK2diOasJEDrD1ml7Afty717YNMQIDAQAB")
	if err != nil {
		log.Fatal("something went horribly wrong")
	}
	PubKey, err = x509.ParsePKCS1PublicKey(raw)
}

// Encrypts text with the passphrase
func Encrypt(text string, passphrase string) string {
	salt := make([]byte, 8)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic(err.Error())
	}

	key, iv := __DeriveKeyAndIv(passphrase, string(salt))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	AntiCrack()

	pad := __PKCS5Padding([]byte(text), block.BlockSize())
	ecb := cipher.NewCBCEncrypter(block, []byte(iv))
	encrypted := make([]byte, len(pad))
	ecb.CryptBlocks(encrypted, pad)

	return b64.StdEncoding.EncodeToString([]byte("Salted__" + string(salt) + string(encrypted)))
}

func EncryptWithSalt(text string, passphrase string, salt string) string {
	key, iv := __DeriveKeyAndIv(passphrase, salt)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	pad := __PKCS5Padding([]byte(text), block.BlockSize())
	ecb := cipher.NewCBCEncrypter(block, []byte(iv))
	encrypted := make([]byte, len(pad))
	ecb.CryptBlocks(encrypted, pad)

	return b64.StdEncoding.EncodeToString([]byte(string(encrypted)))
}

// Decrypts encrypted text with the passphrase
func Decrypt(encrypted string, passphrase string) string {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			_, ok = r.(error)
			if !ok {
				_ = fmt.Errorf("pkg: %v", r)
			}
		}
	}()
	ct, _ := b64.StdEncoding.DecodeString(encrypted)
	if len(ct) < 16 || string(ct[:8]) != "Salted__" {
		return ""
	}

	salt := ct[8:16]
	ct = ct[16:]
	key, iv := __DeriveKeyAndIv(passphrase, string(salt))
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	cbc := cipher.NewCBCDecrypter(block, []byte(iv))
	dst := make([]byte, len(ct))
	cbc.CryptBlocks(dst, ct)

	return string(__PKCS5Trimming(dst))
}

func DecryptWithSalt(encrypted string, passphrase string, salt string) string {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			_, ok = r.(error)
			if !ok {
				_ = fmt.Errorf("pkg: %v", r)
			}
		}
	}()
	ct, _ := b64.StdEncoding.DecodeString(encrypted)
	key, iv := __DeriveKeyAndIv(passphrase, salt)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	cbc := cipher.NewCBCDecrypter(block, []byte(iv))
	dst := make([]byte, len(ct))
	cbc.CryptBlocks(dst, ct)

	return string(__PKCS5Trimming(dst))
}

func __PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func __PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func __DeriveKeyAndIv(passphrase string, salt string) (string, string) {
	salted := ""
	dI := ""

	for len(salted) < 48 {
		md := md5.New()
		md.Write([]byte(dI + passphrase + salt))
		dM := md.Sum(nil)
		dI = string(dM[:16])
		salted = salted + dI
	}

	key := salted[0:32]
	iv := salted[32:48]

	return key, iv
}

func RSAEncrypt(data string) string {
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		PubKey,
		[]byte(data),
		nil)
	if err != nil {
		panic(err)
	}

	return b64.StdEncoding.EncodeToString(encryptedBytes)
}