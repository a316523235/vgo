package gpt35

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/a316523235/wingo/conf"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

func TestProxy(t *testing.T)  {
	// 配置代理地址
	proxyUrl, err := url.Parse("http://127.0.0.1:7890")
	if err != nil {
		panic(err)
	}
	// 通过ProxyURL函数设置代理地址
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
		// 指定代理类型
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		// 防止keepalive过多创建文件
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	// 带自定义Transport的http.Client，该Client对象可以自定义各种参数
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(60 * time.Second),
	}

	//jsonData := json.Marshal([]int{1,2})
	requestBody := []byte(`{"model": "gpt-3.5-turbo", "messages": [{"role": "user", "content": "bootstrap中的tooltip如何调整最大宽度"}]}`)
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
		return
	}
	bt := bytes.NewBuffer(jsonData)

	req, err := http.NewRequest("Post", "https://api.openai.com/v1/chat/completions", bt)
	if err != nil {
		t.Fatal(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer sk-" + conf.AIKey)

	//client := &http.Client{}
	resp, err := client.Do(req)

	// 请求Url
	//resp, err := client.Get("http://www.google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	// 打印结果
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func TestProxy2(t *testing.T)  {
	gClient := NewProxyClient("sk-" + conf.AIKey)

	req := &Request{
		Model: ModelGpt35Turbo,
		Messages: []*Message{
			{
				Role:    RoleUser,
				Content: "Hello",
			},
		},
	}
	resp, err := gClient.GetChat(req)
	if err != nil {
		panic(err)
	}

	time.Sleep(300 * time.Second)

	println(resp.Choices[0].Message.Content)
	println(resp.Usage.PromptTokens)
	println(resp.Usage.CompletionTokens)
	println(resp.Usage.TotalTokens)
}
