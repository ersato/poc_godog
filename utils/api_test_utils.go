package utils

import (
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
)

type ApiFeature struct {
	resp       *http.Response
	pathConfig string
}

func (a *ApiFeature) SetConfig(pathConfig string) {
	a.pathConfig = pathConfig
}

func (a *ApiFeature) CallService(typeScopePnrWrapper, pathUrlConfig, pathClaimState, pathRequestBody, headersKey string) (err error) {
	url := a.getUrl(typeScopePnrWrapper, pathUrlConfig)
	headerJson := GetConfigJsonFile(a.pathConfig, headersKey)
	var headers []map[string]string
	if err := json.Unmarshal([]byte(headerJson), &headers); err != nil {
		panic(err)
	}
	a.resp = CallWebServicePostWithHeaders(headers, pathRequestBody, url+pathClaimState)
	return
}

func (a *ApiFeature) getUrl(typeScopePnrWrapper, pathUrlConfig string) string {
	if hostTest := os.Getenv("HOST_TEST"); hostTest != "" {
		return hostTest + typeScopePnrWrapper + "/"
	}
	return GetConfigJsonFile(a.pathConfig, pathUrlConfig)
}

func (a *ApiFeature) TheResponseCodeShouldBe(code int) (err error) {
	if code != a.resp.StatusCode {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.StatusCode)
	}
	return nil
}

func (a *ApiFeature) TheResponseBodyEqual(jsonResponseBodyPath string) (err error) {
	var expected, actual interface{}
	bodyBytes, _ := io.ReadAll(a.resp.Body)
	if err = json.Unmarshal(GetJsonRead(jsonResponseBodyPath), &expected); err != nil {
		return
	}
	if err = json.Unmarshal(bodyBytes, &actual); err != nil {
		return
	}
	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected JSON does not match actual, %v vs. %v", expected, actual)
	}
	return nil
}

func (a *ApiFeature) TheResponseAttributeEqual(table *godog.Table) (err error) {
	var actual interface{}
	bodyBytes, _ := io.ReadAll(a.resp.Body)
	if err = json.Unmarshal(bodyBytes, &actual); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(table.Rows); i++ {
		var attribute = table.Rows[i].Cells[0].Value
		var expected = table.Rows[i].Cells[1].Value
		var actual = gjson.Get(string(bodyBytes), attribute).String()
		if actual != expected {
			return fmt.Errorf("attribute %s expected response value to be: %s, but actual is: %s",
				attribute, expected, actual)
		}
	}
	return
}
