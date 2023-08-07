package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserId    string `json:"user_id"`
}

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage： simpleDict WORD
		example: simpleDict hello
					`)
		os.Exit(1)
	}
	word := os.Args[1]
	query(word)
}

func query(word string) {
	//创建client
	client := &http.Client{}
	// var data = strings.NewReader(`{"trans_type":"zh2en","source":"翻译","user_id":"64cbb03d59ccfc0018a3b058"}`)
	request := DictRequest{
		TransType: "en2zh",
		Source:    "good",
		UserId:    "64cbb03d59ccfc0018a3b058",
	}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
		return
	}
	var data = bytes.NewReader(buf)

	//发起请求
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}

	//设置请求头
	req.Header.Set("authority", "api.interpreter.caiyunai.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("app-name", "xy")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("device-id", "")
	req.Header.Set("origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("sec-ch-ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")

	//获取响应
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	/**
	 *  golang编码习惯
	 *  defer resp.Body.Close()
	 *  response的body为流,为了避免资源泄露，defer手动关闭流
	 *  函数结束之后从下往上触发
	 *  此处只有一个defer，函数结束之后
	 */
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("there is a error")
		}
	}(resp.Body)

	//流读到内存里面成为byte数组
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}

	//fmt.Printf("%s\n", bodyText)
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n\n", dictResponse)
	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}

	for _, item := range dictResponse.Dictionary.Synonym {
		fmt.Print(item + " ")
	}

	for _, item := range dictResponse.Dictionary.Antonym {
		fmt.Print(item)
	}

	for _, item := range dictResponse.Dictionary.WqxExample {
		fmt.Println(item)
	}
}
