package service

import (
	"encoding/json"
	"fmt"
	gpt35 "github.com/a316523235/wingo/third/gtp35"
	"testing"
)

func TestOcr(t *testing.T) {
	Ocr()
}

func TestOcr2(t *testing.T)  {
	bodyStr := `{"id":"chatcmpl-6w786nSV9i9CMkQPTQlP5INnRicxJ","object":"chat.completion","created":1679308450,
	 	"model":"gpt-3.5-turbo-0301","usage":{"prompt_tokens":30,"completion_tokens":190,"total_tokens":220},"choices":[{"message":{"role":"assistant","content":"\n\n7890端口是Shadowsocks代理协议默认的端口号，但并不仅仅用于代理。这个端口也可以用于其他类型的服务，例如微软的Hyper-V虚拟化管理。 \n\n如果你想使用Shadowsocks代理服务，请确保你已经正确地安装好了代理服务器和客户端，并且配置好了相应的端口和加密算法。如果端口不能使用，可能是因为代理服务器没有正确地配置或者被防火墙或ISP封锁了。 \n\n如果你对网络配置不是很熟悉，建议找到一些相关的教程或者咨询专业人士来帮助你解决问题。"},"finish_reason":"stop","index":0}]}`

	var res gpt35.Response
	err := json.Unmarshal([]byte(bodyStr), &res)
	if err != nil {
		fmt.Println("response unmarshal error, err: " + err.Error())
		fmt.Println("over")
	}
	formattedData, err := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(formattedData))
}