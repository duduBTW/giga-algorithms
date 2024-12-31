//go:build js && wasm
// +build js,wasm

package jslayer

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"syscall/js"
)

func GetElementById(id string) (js.Value, error) {
	element := js.Global().Get("document").Call("getElementById", id)
	if IsNil(element) {
		return element, errors.New("Element not found")
	}

	return element, nil
}

func QuerySelector(selector string) (js.Value, error) {
	element := js.Global().Get("document").Call("querySelector", selector)
	if IsNil(element) {
		return element, errors.New("Element not found")
	}

	return element, nil
}

func QuerySelectorAll(selector string) ([]js.Value, error) {
	nodeList := js.Global().Get("document").Call("querySelectorAll", selector)
	if IsNil(nodeList) {
		return nil, errors.New("Element not found")
	}

	var elements []js.Value
	length := nodeList.Get("length").Int()
	for i := 0; i < length; i++ {
		element := nodeList.Index(i)
		if IsNil(element) {
			continue
		}

		elements = append(elements, element)
	}

	return elements, nil
}

func Id(id string) string {
	return "#" + id
}

func GetStringAttr(selector string, attr string) (string, error) {
	element := js.Global().Get("document").Call("querySelector", selector)
	if IsNil(element) {
		return "", errors.New("Element not found")
	}
	return element.Get("dataset").Get(attr).String(), nil
}

func removeFirstAndLastChar(s string) string {
	if len(s) <= 1 {
		return s
	}
	return s[1 : len(s)-2]
}

func GetJsonData[T any](id string) (T, error) {
	var data T

	element := js.Global().Get("document").Call("getElementById", id)
	if IsNil(element) {
		return data, errors.New("Element not found")
	}
	jsonData := element.Get("innerText").String()
	jsonData = strings.Trim(strings.ReplaceAll(removeFirstAndLastChar(jsonData), `\"`, `"`), " ")
	fmt.Println(jsonData)

	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func SetInnerText(selector string, text string) {
	element := js.Global().Get("document").Call("querySelector", selector)
	if IsNil(element) {
		fmt.Println("Element not found", selector)
		return
	}
	element.Set("innerText", text)
}

func CopyToClipboard(text string) {
	clipboard := js.Global().Get("navigator").Get("clipboard")
	if IsNil(clipboard) {
		fmt.Println("Clipboard API not supported")
		return
	}
	clipboard.Call("writeText", text)
}

func SetInnerHTML(selector string, html string) error {
	element := js.Global().Get("document").Call("querySelector", selector)
	if IsNil(element) {
		return errors.New("Element not found")
	}
	element.Set("innerHTML", html)
	return nil
}

func ReplaceWithHTML(selector string, html string) error {
	element := js.Global().Get("document").Call("querySelector", selector)
	if IsNil(element) {
		return errors.New("Element not found")
	}
	element.Call("replaceWith", js.Global().Get("document").Call("createRange").Call("createContextualFragment", html))
	return nil
}
