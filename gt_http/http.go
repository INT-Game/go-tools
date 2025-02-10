package gt_http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// HttpGet 发送GET请求
// @param url string 请求地址
// @param header map[string]string 请求头
func HttpGet(url string, header map[string]string) ([]byte, error) {
	var err error
	var respBody []byte

	httpClient := new(http.Client)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return respBody, err
	}

	request.Header.Set("Content-Type", "application/json")

	for k, v := range header {
		request.Header.Set(k, v)
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return respBody, err
	}

	defer func() {
		if response != nil && response.Body != nil {
			err = response.Body.Close()
			if err != nil {
				return
			}
		}
	}()

	if response.StatusCode != http.StatusOK {
		return respBody, errors.New("请求错误")
	}

	respBody, err = io.ReadAll(response.Body)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

// HttpPost 发送POST请求
// @param url string 请求地址
// @param header map[string]string 请求头
// @param reqData any 请求数据
func HttpPost(url string, header map[string]string, reqData any) ([]byte, error) {
	var err error
	var respBody []byte

	httpClient := new(http.Client)

	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return respBody, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return respBody, err
	}

	request.Header.Set("Content-Type", "application/json")

	for k, v := range header {
		request.Header.Set(k, v)
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return respBody, err
	}

	defer func() {
		if response != nil && response.Body != nil {
			err = response.Body.Close()
			if err != nil {
				return
			}
		}
	}()

	if response.StatusCode != http.StatusOK {
		return respBody, errors.New("请求错误")
	}

	respBody, err = io.ReadAll(response.Body)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}
