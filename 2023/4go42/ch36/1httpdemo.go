package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
	}

	cookie := &http.Cookie{Name: "userId", Value: strconv.Itoa(123456)}

	request.AddCookie(cookie)

	request.Header.Set("Accept", "text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8")
	request.Header.Set("Accept-Charset", "GBK, utf-8;q=0.7, *;q=0.3")
	request.Header.Set("Accept-Encoding", "gzip, deflate, sdch")
	request.Header.Set("Accept-Language", "zh-CN, zh;q=0.8")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Connection", "keep-alive")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	if response.StatusCode == 200 {
		body, err := gzip.NewReader(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		defer body.Close()
		r, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(r))
	}
}
