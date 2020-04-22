package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js"
)

type Book struct {
	title  string
	author string
	price  float32
}

func onError(event []js.Value) {
	fmt.Println("okok")
}

func main() {
	c := make(chan struct{}, 0)
	books := []Book{
		Book{title: "Book 1", author: "Jams", price: 250},
		Book{title: "Book 2", author: "Jams", price: 352},
	}
	fmt.Println("Hello WASM!")
	// demo := js.Global().Get("document").Call("getElementById", "demo")
	elm := js.Global().Get("document").Call("getElementById", "text")

	html := ``
	for _, book := range books {
		html += `<div>`
		html += `<p> Title: ` + book.title + `</p>`
		html += `<p> Author: ` + book.author + `</p>`
		html += `<p> Price: ` + fmt.Sprint(book.price) + `</p>`
		html += `</div>`
	}
	elm.Set("innerHTML", html)
	fmt.Println(elm.Get("innerHTML"))

	// demo.Call("addEventListener", "click", onError)
	js.Global().Get("window").Set("onerror", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// message := args[0]
		// source := args[1]
		// lineno := args[2]
		// colno := args[3]
		error := args[4]
		fmt.Println(error.Get("stack").String())
		return nil
	}))

	req := map[string]interface{}{"title": "sdfsdf", "price": 525}
	jsonValue, _ := json.Marshal(req)
	res, err := http.Post("http://localhost:8100/test", "application/json; charset=utf-8", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("Request faild with: %s", err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}

	<-c
}
