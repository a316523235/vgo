package gpt35

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

const ModelGpt35Turbo = "gpt-3.5-turbo"

const MaxTokensGpt35Turbo = 4096

const (
	RoleUser      RoleType = "user"
	RoleAssistant RoleType = "assistant"
	RoleSystem    RoleType = "system"
)

const DefaultUrl = "https://api.openai.com/v1/chat/completions"
//const DefaultUrl = "https://www.google.com/"		//test

type Client struct {
	transport *http.Client
	apiKey    string
	url       string
	proxyUrl  string
}

func NewClient(apiKey string) *Client {
	return &Client{
		transport: http.DefaultClient,
		apiKey:    apiKey,
		url:       DefaultUrl,
	}
}

func NewProxyClient(apiKey string) *Client {
	return &Client{
		transport: http.DefaultClient,
		apiKey:    apiKey,
		url:       DefaultUrl,
		proxyUrl: "http://127.0.0.1:7890",
	}
}

func (c *Client) GetChat(r *Request) (*Response, error) {
	//jsonData, err := json.Marshal(r)
	//if err != nil {
	//	return nil, err
	//}

	jsonData := []byte(`{"model": "gpt-3.5-turbo", "messages": [{"role": "user", "content": "Hello!"}]}`)
	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	proxyUrlParsed, _ := url.Parse(c.proxyUrl)
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrlParsed)}}
	httpResp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()
	//defer func() {
	//	_ = httpResp.Body.Close()
	//}()

	var resp Response
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
