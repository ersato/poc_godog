package utils

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func CallWebServicePostWithHeaders(headers []map[string]string, jsonPath, url string) (resp *http.Response) {
	fmt.Printf("URL: %s\n", url)
	var jsonStr = GetJsonRead(jsonPath)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonStr))
	for _, mapValue := range headers {
		req.Header.Add(mapValue["key"], mapValue["value"])
	}
	if tigerToken := os.Getenv("TIGER_TOKEN"); tigerToken != "" {
		req.Header.Add("x-tiger-token", tigerToken)
	}
	client := &http.Client{}
	resp, _ = client.Do(req)
	return resp
}
