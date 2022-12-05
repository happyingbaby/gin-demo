package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]

	}

	return string(result)
}

func MD5(str string) string {
	b := []byte(str)
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

//定义一个创建文件目录的方法
func Mkdir(basePath string) string {
	//	1.获取当前时间,并且格式化时间
	folderName := time.Now().Format("20060102")
	folderPath := filepath.Join(basePath, folderName)
	//使用mkdirall会创建多层级目录
	os.MkdirAll(folderPath, os.ModePerm)
	return folderPath
}
