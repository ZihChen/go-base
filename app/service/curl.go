package service

import (
	"GoFormat/app/global/errorcode"
	"GoFormat/app/global/helper"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// sendGet CURL GET
func sendGet(apiURL string, header map[string]string, param map[string]interface{}) (body []byte, apiErr errorcode.Error) {
	client := &http.Client{}
	// 建立一個請求
	reqest, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		go helper.WarnLog(fmt.Sprintf("CURL_CREATE_FAIL: %v", err))
		apiErr = errorcode.GetAPIError("CURL_CREATE_FAIL")
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
		go helper.WarnLog(fmt.Sprintf("API_CONNECT_ERROR: ErrorMsg: %v ", err))
		apiErr = errorcode.GetAPIError("API_CONNECT_ERROR")
		return nil, apiErr
	}
	if resp.StatusCode != 200 {
		go helper.WarnLog(fmt.Sprintf("API_STATUS_ERROR: Status: %d, ErrorMsg: %v ", resp.StatusCode, err))
		apiErr = errorcode.GetAPIError("API_STATUS_ERROR")
		return nil, apiErr
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		go helper.WarnLog(fmt.Sprintf("CURL_GET_ERROR: %v", err2))
		apiErr = errorcode.GetAPIError("CURL_GET_ERROR")
		return nil, apiErr
	}

	return body, apiErr
}

// sendPost CURL POST
func sendPost(apiURL string, header map[string]string, param map[string]interface{}) (body []byte, apiErr errorcode.Error) {
	// 組參數
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

	// 建立一個請求
	client := &http.Client{}
	reqest, err := http.NewRequest(http.MethodPost, apiURL, strings.NewReader(form.Encode()))
	if err != nil {
		go helper.WarnLog(fmt.Sprintf("CURL_CREATE_FAIL: %v", err))
		apiErr = errorcode.GetAPIError("CURL_CREATE_FAIL")
		return nil, apiErr
	}

	// 組Header
	for hk, hv := range header {
		reqest.Header.Add(hk, hv)
	}

	// 執行
	resp, err := client.Do(reqest)
	if err != nil {
		go helper.WarnLog(fmt.Sprintf("API_CONNECT_ERROR: ErrorMsg: %v ", err))
		apiErr = errorcode.GetAPIError("API_CONNECT_ERROR")
		return nil, apiErr
	}
	if resp.StatusCode != 200 {
		go helper.WarnLog(fmt.Sprintf("API_STATUS_ERROR: Status: %d, ErrorMsg: %v ", resp.StatusCode, err))
		apiErr = errorcode.GetAPIError("API_STATUS_ERROR")
		return nil, apiErr
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		go helper.WarnLog(fmt.Sprintf("CURL_POST_FAIL: %v", err2))
		apiErr = errorcode.GetAPIError("CURL_POST_FAIL")
		return nil, apiErr
	}

	return body, apiErr
}

// sendPut CURL PUT
func sendPut(apiURL string, header map[string]string, param map[string]interface{}) (body []byte, apiErr errorcode.Error) {
	// 組參數
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

	// 建立一個請求
	client := &http.Client{}
	reqest, err := http.NewRequest(http.MethodPut, apiURL, strings.NewReader(form.Encode()))
	if err != nil {
		go helper.WarnLog(fmt.Sprintf("CURL_CREATE_FAIL: %v", err))
		apiErr = errorcode.GetAPIError("CURL_CREATE_FAIL")
		return nil, apiErr
	}

	// 組Header
	for hk, hv := range header {
		reqest.Header.Add(hk, hv)
	}

	// 執行
	resp, err := client.Do(reqest)
	if err != nil {
		go helper.WarnLog(fmt.Sprintf("API_CONNECT_ERROR: ErrorMsg: %v ", err))
		apiErr = errorcode.GetAPIError("API_CONNECT_ERROR")
		return nil, apiErr
	}
	if resp.StatusCode != 200 {
		go helper.WarnLog(fmt.Sprintf("API_STATUS_ERROR: Status: %d, ErrorMsg: %v ", resp.StatusCode, err))
		apiErr = errorcode.GetAPIError("API_STATUS_ERROR")
		return nil, apiErr
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		go helper.WarnLog(fmt.Sprintf("CURL_POST_FAIL: %v", err2))
		apiErr = errorcode.GetAPIError("CURL_POST_FAIL")
		return nil, apiErr
	}

	return body, apiErr
}
