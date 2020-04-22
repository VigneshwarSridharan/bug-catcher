package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	price := 54655
	book := map[string]interface{}{"title": "sdfsdf", "price": price}
	jsonValue, _ := json.Marshal(book)
	res, err := http.Post("http://localhost:8100/test", "application/json; charset=utf-8", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("Request faild with: %s", err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}

}
