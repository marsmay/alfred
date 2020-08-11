package util

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"time"
)

const HttpTimeout = 5 * time.Second

func HttpRequest(url string, header map[string]string, body []byte) ([]byte, error) {
	method := "GET"

	if len(body) > 0 {
		method = "POST"
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{Timeout: HttpTimeout}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error http code %d", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func GbkToUtf8(from string) (to string, err error) {
	reader := transform.NewReader(bytes.NewReader([]byte(from)), simplifiedchinese.GBK.NewDecoder())
	bs, err := ioutil.ReadAll(reader)

	if err == nil {
		to = string(bs)
	}

	return
}
