package gpt35

import (
	"bytes"
	"fmt"
	"github.com/a316523235/wingo/conf"
	"net/http"
	"net/url"
	"testing"
)

/**
原始curl
curl --proxy http://127.0.0.1:7890 https://api.openai.com/v1/chat/completions \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer sk-........' \
  -d '{
  "model": "gpt-3.5-turbo",
  "messages": [{"role": "user", "content": "Hello!"}]
}'
以下gpt回答的使用Go编写的等效代码：
 */
func TestByGpt(t *testing.T) {
	// 定义请求的URL
	api := "https://api.openai.com/v1/chat/completions"

	// 定义请求体
	requestBody := []byte(`{"model": "gpt-3.5-turbo", "messages": [{"role": "user", 
"content": "7890端口一般是用来做代理吗，怎么不行了"}]}`)

	// 创建代理请求
	proxyUrl := "http://127.0.0.1:7890"
	proxyUrlParsed, _ := url.Parse(proxyUrl)
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrlParsed)}}
	//client := &http.Client{Transport: &http.Transport{
	//	Proxy: http.ProxyURL(proxyUrlParsed),
	//	TLSClientConfig: &tls.Config{
	//		InsecureSkipVerify: true,
	//		MinVersion: tls.VersionTLS12,
	//		MaxVersion: tls.VersionTLS13,
	//	},
	//}}

	// 创建HTTP POST请求
	req, err := http.NewRequest("POST", api, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer sk-" + conf.AIKey)

	// 发送HTTP请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 关闭响应体
	defer resp.Body.Close()

	// 读取响应体
	responseBody := new(bytes.Buffer)
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 输出响应结果
	fmt.Println(responseBody.String())
}
