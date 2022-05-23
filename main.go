package main

import (
	"L1273/bot"
	"L1273/config"
	"L1273/requests"
	"fmt"
	"time"
)

func main() {
	// 主循环
	for true {
		fmt.Println("主循环")
		c := config.Viper()
		// 多少秒请求一次
		time.Sleep(c.Num * time.Second)
		for _, v := range c.Address {
			time.Sleep(2 * time.Second)
			print(v + "\n")
			info := requests.Transaction(c.Api, v)
			if len(info.Result) <= 0 {
				continue
			}
			results := info.Result[0]
			// 返回 log 中的 md5 值
			logList := requests.GetFileContents()
			fmt.Println(logList)
			// 字符串编码为 md5
			md5Str := requests.ToMd5(results.Hash + results.TimeStamp)
			fmt.Println(md5Str)
			//// 判断是否存在 md5
			isMd5 := requests.IsValue(md5Str, logList)
			fmt.Println(isMd5)
			if isMd5 == true {
				fmt.Println("存在不通知")
				continue
			} else {
				fmt.Println("通知")
				markdown := bot.Markdown(results.Hash, results.TransactionIndex, results.From, results.To, results.Value, results.Input)
				bot.Bot(c.BotApi, c.BotKey, markdown)
				requests.WriteMd5(md5Str)
			}
		}
	}
}
