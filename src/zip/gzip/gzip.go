package gzip

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func RequestGzip(enable bool) (int, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if enable { // 请求 header 添加 gzip
		req.Header.Add("Content-Encoding", "gzip")
		req.Header.Add("Accept-Encoding", "gzip")
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do Error:", err)
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll Error", err)
		return 0, err
	}
	return len(body), nil
}

// client 解析 gzip 返回
func ClientUncompress() {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	req.Header.Add("Content-Encoding", "gzip")
	req.Header.Add("Accept-Encoding", "gzip")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var buf [1024 * 1024]byte
	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	_, err = reader.Read(buf[:])
	if err != nil {
		log.Fatal(err)
	}
	reader.Close()
}

func ClientNormal() {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var buf [1024 * 1024]byte
	_, err = resp.Body.Read(buf[:])
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
}
