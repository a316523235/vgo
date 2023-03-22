package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/a316523235/wingo/conf"
	gpt35 "github.com/a316523235/wingo/third/gtp35"
	"github.com/a316523235/wingo/util"
	"net/http"
	"net/url"
	"time"
)

type ReqBody struct {
	Model string
}

func StartMyGpt() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("异常, 原因：")
			fmt.Println(err)
		}
	}()
	c := gpt35.NewProxyClient("sk-" +  conf.AIKey)
	req := &gpt35.Request{
		Model: gpt35.ModelGpt35Turbo,
		Messages: []*gpt35.Message{
			{
				Role:    gpt35.RoleUser,
				Content: "Hello",
			},
		},
	}

	resp, err := c.GetChat(req)
	if err != nil {
		panic(err)
	}

	time.Sleep(300 * time.Second)

	println(resp.Choices[0].Message.Content)
	println(resp.Usage.PromptTokens)
	println(resp.Usage.CompletionTokens)
	println(resp.Usage.TotalTokens)
}

func StartMyGpt2()  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("异常, 原因：")
			fmt.Println(err)
		}
	}()

	// 定义请求的URL
	api := "https://api.openai.com/v1/chat/completions"

	// 创建代理请求
	proxyUrl := "http://127.0.0.1:7890"
	proxyUrlParsed, _ := url.Parse(proxyUrl)
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrlParsed)}}

	//请求
	reqBody := gpt35.Request{
		Model:gpt35.ModelGpt35Turbo,
		Messages: []*gpt35.Message{},
	}

	for {
		if !Switch.IsTaskOpen() {
			break
		}
		str := utils.ReadLine("-----input question------ \n")

		if str == "q" {
			fmt.Println("over")
			break
		}

		reqBody.Messages = append(reqBody.Messages, &gpt35.Message{Role:gpt35.RoleUser, Content:str})
		requestBody, err :=json.Marshal(reqBody)
		if err != nil {
			fmt.Println("request body Marshal error, err: " + err.Error())
			fmt.Println("over")
			break
		}

		// 创建HTTP POST请求
		req, err := http.NewRequest("POST", api, bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Println("new request error, err: " + err.Error())
			fmt.Println("over")
			break
		}

		// 设置请求头
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer sk-" + conf.AIKey)

		// 发送HTTP请求
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("client do error, err: " + err.Error())
			fmt.Println("over")
			break
		}

		// 关闭响应体
		defer resp.Body.Close()

		// 读取响应体
		responseBody := new(bytes.Buffer)
		_, err = responseBody.ReadFrom(resp.Body)
		if err != nil {
			fmt.Println("read response body error, err: " + err.Error())
			fmt.Println("over")
			break
		}

		// 输出响应结果
		bodyStr := responseBody.String()
		var res gpt35.Response
		err = json.Unmarshal([]byte(bodyStr), &res)
		if err != nil {
			fmt.Println("response unmarshal error, err: " + err.Error())
			fmt.Println("over")
			break
		}
		if res.Error != nil {
			fmt.Println("api back error tip: " + res.Error.Message)
			fmt.Println("over")
			break
		}
		if res.Choices != nil {
			fmt.Println("api back Choices is nil")
			fmt.Println("over")
			break
		}

		if len(reqBody.Messages) > 5 {
			reqBody.Messages = reqBody.Messages[len(reqBody.Messages) - 5:]
		}
		for i, _ := range res.Choices {
			if res.Choices[i].Message != nil {
				reqBody.Messages = append(reqBody.Messages, res.Choices[i].Message)
			}
		}
		formattedData, err := json.MarshalIndent(res, "", "  ")
		fmt.Println(string(formattedData))

		time.Sleep(1 * time.Second)
	}

	time.Sleep(1 * time.Second)
}

func StartMyGpt3()  {
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("异常, 原因：")
	//		fmt.Println(err)
	//	}
	//}()
	//
	//c := openai.NewClient("sk-")
	//ctx := context.Background()
	//
	//req := openai.CompletionRequest{
	//	Model:     openai.GPT3Ada,
	//	MaxTokens: 5,
	//	Prompt:    "Lorem ipsum",
	//	Stream:    true,
	//}
	//stream, err := c.CreateCompletionStream(ctx, req)
	//if err != nil {
	//	fmt.Printf("CompletionStream error: %v\n", err)
	//	return
	//}
	//defer stream.Close()
	//
	//for {
	//	response, err := stream.Recv()
	//	if errors.Is(err, io.EOF) {
	//		fmt.Println("Stream finished")
	//		return
	//	}
	//
	//	if err != nil {
	//		fmt.Printf("Stream error: %v\n", err)
	//		return
	//	}
	//
	//
	//	fmt.Printf("Stream response: %v\n", response)
	//}
	//
	//time.Sleep(60 * time.Second)
}