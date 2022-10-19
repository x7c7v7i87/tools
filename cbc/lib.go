package cbc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

//参考文档
//http://www.topgoer.com/%E5%85%B6%E4%BB%96/%E5%8A%A0%E5%AF%86%E8%A7%A3%E5%AF%86/%E5%8A%A0%E5%AF%86%E8%A7%A3%E5%AF%86.html
//高级加密标准（Adevanced Encryption Standard ,AES）
//16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法

type Cbc struct {
	Key  string
	data string
}

func Set(key string) *Cbc {
	return &Cbc{
		Key: key,
	}
}

// PKCS7 填充模式
func (c *Cbc) pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 填充的反向操作，删除填充字符串
func (c *Cbc) pKCS7UnPadding1(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}

// 实现加密
func (c *Cbc) aesEcrypt(origData []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = c.pKCS7Padding(origData, blockSize)
	//采用AES加密方法中CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	//执行加密
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 实现解密
func (c *Cbc) aesDeCrypt(cypted []byte, key []byte) (string, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cypted))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(origData, cypted)
	//去除填充字符串
	origData, err = c.pKCS7UnPadding1(origData)
	if err != nil {
		return "", err
	}
	return string(origData), err
}

// 加密base64
func (c *Cbc) Encode(data string) string {
	c.data = data
	pwd := []byte(c.data)
	result, err := c.aesEcrypt(pwd, []byte(c.Key))
	if err != nil {
		return ""
	}
	return hex.EncodeToString(result)
}

// 解密
func (c *Cbc) Decode(data string) string {
	c.data = data
	temp, _ := hex.DecodeString(c.data)
	//执行AES解密
	res, _ := c.aesDeCrypt(temp, []byte(c.Key))
	return res
}
