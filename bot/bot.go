package bot

import (
	"L1273/transition"
	"fmt"
	"github.com/imroc/req/v3"
)

// Body api 的 json
type Body struct {
	Msgtype  string   `json:"msgtype"`
	Markdown *Message `json:"markdown"`
}

type Message struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// Markdown 把信息转换成 markdown 格式
func Markdown(hash, index, from, to, value, input string) string {
	transfer, _ := transition.GetValue(input)
	hash_ := fmt.Sprintf(`[检测地址](https://etherscan.io/tx/%s) %s %s`, hash, transfer, "  \n  ")
	value_ := value + "  \n  "
	from_ := fmt.Sprintf(`从[%s](https://etherscan.io/address/%s)%s`, from, from, "  \n  ")
	to_ := fmt.Sprintf(`到[%s](https://etherscan.io/address/%s)%s`, to, to, "  \n  ")
	index_ := index + "  \n  "
	markdown := fmt.Sprintf(`%s %s %s %s %s`, hash_, value_, from_, to_, index_)
	return markdown
}

// Bot 机器人发送消息
func Bot(api, key, markdown string) {
	// 实例化
	client := req.NewClient()
	message := &Message{Title: "测试", Text: markdown}
	url := api + key
	body := &Body{Msgtype: "markdown", Markdown: message}
	_, err := client.NewRequest().SetBody(body).Post(url)
	if err != nil {
		return
	}
}
