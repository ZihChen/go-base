package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goformat/app/global"
	"goformat/app/global/helper"
	"goformat/library/errorcode"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
)

type Curl struct {
	send INotifier
}

type ICurl interface {
	SendGet(apiURL string, header map[string]string, param map[string]interface{}) (body []byte, apiErr errorcode.Error)
	SendPost(apiURL string, header map[string]string, param map[string]interface{}, rawData bool) (body []byte, apiErr errorcode.Error)
	SendPut(apiURL string, header map[string]string, param map[string]interface{}, rawData bool) (body []byte, apiErr errorcode.Error)
	SendDelete(apiURL string, header map[string]string, param map[string]interface{}, rawData bool) (body []byte, apiErr errorcode.Error)
	PostFile(apiURL string, header map[string]string, param map[string]interface{}, filename string) (body []byte, apiErr errorcode.Error)
}

// NewCurl Implement NewCurl
func NewCurl() ICurl {
	return &Curl{
		send: NewNotify(),
	}
}

// sendGet CURL GET
func (c *Curl) SendGet(apiURL string, header map[string]string, param map[string]interface{}) (body []byte, apiErr errorcode.Error) {
	client := &http.Client{}
	// 建立一個請求
	reqest, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "CURL_CREATE_FAIL", err.Error(), param)
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}
	// 組Header
	for hk, hv := range header {
		reqest.Header.Add(hk, hv)
	}
	//組參數
	q := reqest.URL.Query()
	for pk, pv := range param {
		paramV := reflect.ValueOf(pv)
		if paramV.Kind() == reflect.Slice {
			for i := 0; i < paramV.Len(); i++ {
				value := paramV.Index(i)
				q.Add(pk, fmt.Sprintf("%v", value))
			}
			continue
		}
		q.Add(pk, fmt.Sprintf("%v", paramV))
	}
	reqest.URL.RawQuery = q.Encode()

	// 執行
	resp, err := client.Do(reqest)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "API_CONNECT_ERROR", err.Error(), param)
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}
	if resp.StatusCode != 200 {
		errMsg := fmt.Sprintf("Url: %v, Status: %d ", apiURL, resp.StatusCode)
		apiErr = helper.ErrorHandle(global.WarnLog, "API_STATUS_ERROR", errMsg, param)
		c.send.HandleServiceError(global.RD3, apiURL, header, param, resp)
		return nil, apiErr
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "CURL_GET_FAIL", err.Error(), param)
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}

	return body, apiErr
}

// sendPost CURL POST
func (c *Curl) SendPost(apiURL string, header map[string]string, param map[string]interface{}, rawData bool) (body []byte, apiErr errorcode.Error) {
	// 組參數 (FormData, RawData)
	var requestData string
	if rawData == true {
		if _, ok := param["onlyText"]; ok {
			requestData = fmt.Sprintf("%v", param["onlyText"])
		} else {
			byteData, err := json.Marshal(param)
			if err != nil {
				apiErr = helper.ErrorHandle(global.WarnLog, "JSON_MARSHAL_ERROR", err.Error())
				c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
				return
			}
			requestData = string(byteData)
		}
	} else {
		form := url.Values{}
		for pk, pv := range param {
			paramV := reflect.ValueOf(pv)
			if paramV.Kind() == reflect.Slice {
				for i := 0; i < paramV.Len(); i++ {
					value := paramV.Index(i)
					form.Add(pk, fmt.Sprintf("%v", value))
				}
				continue
			}
			form.Add(pk, fmt.Sprintf("%v", paramV))
		}
		requestData = form.Encode()
	}

	// 建立一個請求
	client := &http.Client{}
	reqest, err := http.NewRequest(http.MethodPost, apiURL, strings.NewReader(requestData))
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "CURL_CREATE_FAIL", err.Error(), param)
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}

	// 組Header
	for hk, hv := range header {
		reqest.Header.Add(hk, hv)
	}

	// 執行
	resp, err := client.Do(reqest)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "API_CONNECT_ERROR", err.Error(), param)
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}
	if resp.StatusCode != 200 {
		errMsg := fmt.Sprintf("Url: %v, Status: %d ", apiURL, resp.StatusCode)
		apiErr = helper.ErrorHandle(global.WarnLog, "API_STATUS_ERROR", errMsg)
		c.send.HandleServiceError(global.RD3, apiURL, header, param, resp)
		return nil, apiErr
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "CURL_POST_FAIL", err.Error(), param)
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}

	return body, apiErr
}

// sendPut CURL PUT
func (c *Curl) SendPut(apiURL string, header map[string]string, param map[string]interface{}, rawData bool) (body []byte, apiErr errorcode.Error) {
	// 組參數 (FormData, RawData)
	var requestData string
	if rawData == true {
		if _, ok := param["onlyText"]; ok {
			requestData = fmt.Sprintf("%v", param["onlyText"])
		} else {
			byteData, err := json.Marshal(param)
			if err != nil {
				apiErr = helper.ErrorHandle(global.WarnLog, "JSON_MARSHAL_ERROR", err.Error())
				c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
				return
			}
			requestData = string(byteData)
		}
	} else {
		form := url.Values{}
		for pk, pv := range param {
			paramV := reflect.ValueOf(pv)
			if paramV.Kind() == reflect.Slice {
				for i := 0; i < paramV.Len(); i++ {
					value := paramV.Index(i)
					form.Add(pk, fmt.Sprintf("%v", value))
				}
				continue
			}
			form.Add(pk, fmt.Sprintf("%v", paramV))
		}
		requestData = form.Encode()
	}

	// 建立一個請求
	client := &http.Client{}
	reqest, err := http.NewRequest(http.MethodPut, apiURL, strings.NewReader(requestData))
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "CURL_CREATE_FAIL", err.Error())
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}

	// 組Header
	for hk, hv := range header {
		reqest.Header.Add(hk, hv)
	}

	// 執行
	resp, err := client.Do(reqest)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "API_CONNECT_ERROR", err.Error())
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}

	if resp.StatusCode != 200 {
		errMsg := fmt.Sprintf("Url: %v, Status: %d ", apiURL, resp.StatusCode)
		apiErr = helper.ErrorHandle(global.WarnLog, "API_STATUS_ERROR", errMsg)
		c.send.HandleServiceError(global.RD3, apiURL, header, param, resp)
		return nil, apiErr
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "CURL_POST_FAIL", err.Error())
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}

	return body, apiErr
}

// sendDelete CURL Delete
func (c *Curl) SendDelete(apiURL string, header map[string]string, param map[string]interface{}, rawData bool) (body []byte, apiErr errorcode.Error) {
	// 組參數 (FormData, RawData)
	var requestData string
	if rawData == true {
		if _, ok := param["onlyText"]; ok {
			requestData = fmt.Sprintf("%v", param["onlyText"])
		} else {
			byteData, err := json.Marshal(param)
			if err != nil {
				apiErr = helper.ErrorHandle(global.WarnLog, "JSON_MARSHAL_ERROR", err.Error())
				c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
				return
			}
			requestData = string(byteData)
		}
	} else {
		form := url.Values{}
		for pk, pv := range param {
			paramV := reflect.ValueOf(pv)
			if paramV.Kind() == reflect.Slice {
				for i := 0; i < paramV.Len(); i++ {
					value := paramV.Index(i)
					form.Add(pk, fmt.Sprintf("%v", value))
				}
				continue
			}
			form.Add(pk, fmt.Sprintf("%v", paramV))
		}
		requestData = form.Encode()
	}

	// 建立一個請求
	client := &http.Client{}
	reqest, err := http.NewRequest(http.MethodDelete, apiURL, strings.NewReader(requestData))
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "CURL_CREATE_FAIL", err.Error())
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}

	// 組Header
	for hk, hv := range header {
		reqest.Header.Add(hk, hv)
	}

	// 執行
	resp, err := client.Do(reqest)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "API_CONNECT_ERROR", err.Error())
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}
	if resp.StatusCode != 200 {
		errMsg := fmt.Sprintf("Url: %v, Status: %d ", apiURL, resp.StatusCode)
		apiErr = helper.ErrorHandle(global.WarnLog, "API_STATUS_ERROR", errMsg)
		c.send.HandleServiceError(global.RD3, apiURL, header, param, resp)
		return nil, apiErr
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "CURL_POST_FAIL", err.Error())
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}

	return body, apiErr
}

func (c *Curl) PostFile(apiURL string, header map[string]string, param map[string]interface{}, filename string) (body []byte, apiErr errorcode.Error) {
	buffer := bytes.NewBufferString("")
	bodyWriter := multipart.NewWriter(buffer)

	// 組參數
	for pk, pv := range param {
		paramV := reflect.ValueOf(pv)
		if paramV.Kind() == reflect.Slice {
			for i := 0; i < paramV.Len(); i++ {
				value := paramV.Index(i)
				bodyWriter.WriteField(pk, fmt.Sprintf("%v", value))
			}
			continue
		}
		bodyWriter.WriteField(pk, fmt.Sprintf("%v", paramV))
	}

	// use the bodyWriter to write the Part headers to the buffer
	_, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "WRITE_BUFFER_ERROR", err.Error())
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}

	// the file data will be the second part of the body
	fh, err := os.Open(filename)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "OPEN_FILE_ERROR", err.Error())
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}
	// need to know the boundary to properly close the part myself.
	boundary := bodyWriter.Boundary()
	closeBuffer := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

	// use multi-reader to defer the reading of the file data until
	// writing to the socket buffer.
	requestReader := io.MultiReader(buffer, fh, closeBuffer)
	fi, err := fh.Stat()
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "START_FILE_ERROR", err.Error())
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}

	req, err := http.NewRequest("POST", apiURL, requestReader)
	if err != nil {
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}

	// Set headers for multipart, and Content Length
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)

	// 組Header
	for hk, hv := range header {
		req.Header.Add(hk, hv)
	}

	req.ContentLength = fi.Size() + int64(buffer.Len()) + int64(closeBuffer.Len())

	// 執行
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "API_CONNECT_ERROR", err.Error(), param)
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}
	if resp.StatusCode != 200 {
		errMsg := fmt.Sprintf("Url: %v, Status: %d ", apiURL, resp.StatusCode)
		apiErr = helper.ErrorHandle(global.WarnLog, "API_STATUS_ERROR", errMsg, param)
		c.send.HandleServiceError(global.RD3, apiURL, header, param, resp)
		return nil, apiErr
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "CURL_POST_FAIL", err.Error(), param)
		c.send.HandleServiceError(global.RD3, apiURL, header, param, err.Error())
		return nil, apiErr
	}

	return body, apiErr
}
