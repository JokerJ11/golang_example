package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	token := "6195801993:AAGKpwK5nfYScQE9Z8hzfBJ-P2kUV-0QFpk"

	URL := fmt.Sprintf("https://api.telegram.org/bot%s", token)

	url := fmt.Sprintf("%s/sendMessage", URL)

	body, _ := json.Marshal(map[string]string{
		"chat_id": "5436770582",
		"text":    "this is leeeeeeeee ????????????",
	})

	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return
	}
	defer response.Body.Close()

	bData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	fmt.Println(string(bData))

	return
}
