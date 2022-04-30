package utility

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func Get(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return ""
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Print(fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body))
		return ""
	}
	if err != nil {
		log.Print(err)
		return ""
	}
	//fmt.Printf("%s", body)
	return string(body)
}
func Post(url string, info interface{}) string {
	infoJson, _ := json.Marshal(info)
	res, err := http.Post(url, "application/json", strings.NewReader(string(infoJson)))
	if err != nil {
		log.Print(err)
		return ""
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Print(fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body))
		return ""
	}
	if err != nil {
		log.Print(err)
		return ""
	}
	//fmt.Printf("%s", body)
	return string(body)
}

func PostWithHeader(url string, info interface{}, header map[string]string) string {
	infoJson, _ := json.Marshal(info)
	requst, _ := http.NewRequest("POST", url, strings.NewReader(string(infoJson)))
	for key, value := range header {
		requst.Header.Set(key, value)
	}
	response, err := (&http.Client{}).Do(requst)
	if err != nil {
		log.Print(err)
	}
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if response.StatusCode > 299 {
		log.Print(fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", response.StatusCode, body))
		return ""
	}
	if err != nil {
		log.Print(err)
		return ""
	}
	//fmt.Printf("%s", body)
	return string(body)
}
