package requests

import (
	"L1273/model"
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/imroc/req/v3"
	"os"
)

// Transaction 请求 api
func Transaction(apiKey, address string) *model.Response {
	// api 地址
	url := fmt.Sprintf(`https://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=10&sort=desc&apikey=%s`, address, apiKey)

	//实例化 client
	client := req.NewClient()

	//使用 get 请求
	response := &model.Response{}
	_, err := client.NewRequest().SetResult(response).Get(url)
	if err != nil {
		return nil
	}

	return response
}

// IsValue 判断 md5 是否存在
func IsValue(str string, md5 []string) bool {
	for _, v := range md5 {
		if v == str {
			return true
		}
	}
	return false
}

// GetFileContents 读取文件返回一个字符串的切片
func GetFileContents() []string {
	file, err := os.Open("log.txt")
	if err != nil {
		return nil
	}
	defer file.Close()

	var logList []string
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		logList = append(logList, fileScanner.Text())
	}
	return logList
}

// ToMd5 字符串编码成 md5
func ToMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// WriteMd5 写入 md 值
func WriteMd5(md5 string) {
	filename := "log.txt"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(md5 + "\n")
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}
