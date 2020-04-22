package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	callBack := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {

			request := map[string]interface{}{
				"message": args[0].String(),
				"source":  args[1].String(),
				"lineno":  args[2].Int(),
				"colno":   args[3].Int(),
				"error":   args[4].Get("stack").String(),
			}
			fmt.Println(request)
			jsonValue, _ := json.Marshal(request)
			res, err := http.Post("http://localhost:8100/bug-catcher", "application/json; charset=utf-8", bytes.NewBuffer(jsonValue))
			if err != nil {
				fmt.Printf("Request faild with: %s", err)
			} else {
				data, _ := ioutil.ReadAll(res.Body)
				fmt.Println(string(data))
			}
		}()
		return nil
	})

	js.Global().Get("window").Set("onerror", callBack)

	<-c
}
