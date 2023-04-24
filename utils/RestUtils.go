package utils

import (
	"bytes"
	"net/http"
)

func CallWebServicePostWithHeaders(headers []map[string]string, jsonPath, url string) (resp *http.Response) {
	var jsonStr = GetJsonRead(jsonPath)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonStr))
	for _, mapValue := range headers {
		req.Header.Add(mapValue["key"], mapValue["value"])
	}
	client := &http.Client{}
	resp, _ = client.Do(req)
	if err != nil {
		return
	}
	return resp
}
