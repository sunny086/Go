package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//publicKeyStr, privateKeyStr, err := GenRsaKey(1024)
	//fmt.Println(publicKeyStr)
	//fmt.Println(privateKeyStr)
	//fmt.Println(err)
	//
	//var msg = "大猩猩玩泥巴"
	//fmt.Println("公钥加密前：", msg)

	//var data []byte
	//var err error
	//data, err = RsaEncrypt([]byte("大猩猩玩泥巴"))
	//if err != nil {
	//	panic(err)
	//}
	//
	//origData, err := RsaDecrypt(data)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(origData))

	err := fmt.Errorf("111")
	fmt.Println(err)
}

// GenRsaKey RSA公钥私钥产生
func GenRsaKey(bits int) (publicKeyStr, privateKeyStr string, err error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	bufferPrivate := new(bytes.Buffer)
	err = pem.Encode(bufferPrivate, block)
	if err != nil {
		return
	}
	privateKeyStr = bufferPrivate.String()
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	bufferPublic := new(bytes.Buffer)
	err = pem.Encode(bufferPublic, block)
	if err != nil {
		return
	}
	publicKeyStr = bufferPublic.String()
	return

}

var decrypted string
var privateKey, publicKey []byte

func init() {
	var err error
	flag.StringVar(&decrypted, "d", "", "加密过的数据")
	flag.Parse()
	publicKey, err = ioutil.ReadFile("public.pem")
	if err != nil {
		os.Exit(-1)
	}
	privateKey, err = ioutil.ReadFile("private.pem")
	if err != nil {
		os.Exit(-1)
	}
}

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)

}
