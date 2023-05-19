package cyberscope

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func CheckBytescode(code string) ([]BytescodeInfo, error) {

	url := "https://www.cyberscope.io/api/signaturescan"

	payload := strings.NewReader(fmt.Sprintf("q=%v", code))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))
	var respond []BytescodeInfo
	err = json.Unmarshal(body, &respond)
	if err != nil {
		var errorRespond ErrorRespond
		err = json.Unmarshal(body, &errorRespond)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(errorRespond.Message)
	}

	return respond, nil
}

type ErrorRespond struct {
	Message string `json:"message"`
}

type BytescodeInfo struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Types     []string `json:"types"`
	Signature string   `json:"signature"`
	Hash      string   `json:"hash"`
	Total     int      `json:"total"`
}
