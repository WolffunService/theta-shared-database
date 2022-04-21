package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

//NewHttpRequest Post request
func NewGetRequest(address string, method string) ([]byte, int, error) {

	urlString, err := url.Parse(address)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	urlString.Path = path.Join(urlString.Path, method)
	res, err := http.Get(urlString.String())
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	return body, http.StatusOK, nil
}

func NewGetRequestParams(address string, method string, mapQueryParams map[string]string) ([]byte, int, error) {

	base, err := url.Parse(address)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	base.Path = path.Join(base.Path, method)
	if mapQueryParams != nil && len(mapQueryParams) >= 1 {
		q := url.Values{}
		for key, value := range mapQueryParams {
			q.Add(key, value)
		}
		base.RawQuery = q.Encode()
	}

	res, err := http.Get(base.String())
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	return body, http.StatusOK, nil
}

func NewPutRequest(address string, method string) (int, error) {

	urlString, err := url.Parse(address)
	if err != nil {
		return http.StatusBadRequest, err
	}

	urlString.Path = path.Join(urlString.Path, method)
	req, err := http.NewRequest(http.MethodPut, urlString.String(), nil)
	if err != nil {
		return http.StatusBadRequest, err
	}
	client := http.Client{}
	res, errRes := client.Do(req)
	if errRes != nil {
		return http.StatusBadRequest, err
	}
	return res.StatusCode, nil
}

// NewPostRequest NewHttpRequest Post request
//TODO reorder return param to match with GetRequest ([]byte, int, error)
func NewPostRequest(address string, method string, reqData interface{}, token string) (int, []byte, error) {
	//parse to json
	jsonString, err := json.Marshal(reqData)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	urlString, err := url.Parse(address)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	urlString.Path = path.Join(urlString.Path, method)
	client := http.Client{}
	req, err := http.NewRequest("POST", urlString.String(), bytes.NewBuffer(jsonString))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if len(token) > 0 {
		req.Header.Set(
			"Authorization",
			"Bearer "+token)
	}
	res, err := client.Do(req)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	defer req.Body.Close()
	fmt.Printf("call api - %s with status %v\n", urlString.String(), res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, nil, err
	}
	return res.StatusCode, body, nil
}
