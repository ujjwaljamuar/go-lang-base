package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// exampleItem()
	getCartItem()
}

// gob
type GoItem struct {
	SKU   string
	Name  string
	Price int
	Added time.Time
	Image []byte
}

// gob is go only serializer

func exampleItem() {
	var item GoItem = GoItem{
		SKU:   "BTN-FF0000",
		Name:  "Red Button",
		Price: 75,
		Added: time.Date(2026, time.March, 7, 14, 48, 50, 50, time.UTC),
		Image: nil,
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(item); err != nil {
		fmt.Println("error (encode): ", err)
		return
	}

	dec := gob.NewDecoder(&buf)
	var item2 GoItem
	if err := dec.Decode(&item2); err != nil {
		fmt.Println("error (decode): ", err)
		return
	}

	fmt.Printf("%s: %s\n", item2.SKU, item2.Name)
}

// JSON->GO  []byte  json.Unmarshal
// GO->JSON  []byte  json.Marshal
// JSON->GO  io.Reader  json.Decoder
// GO->JSON  io.Writer  json.Encoder

// JSON
type CartItem struct {
	Id           string  `json:"id"`
	Index        int     `json:"index"`
	Name         string  `json:"name"`
	Image        string  `json:"image"`
	CurrentPrice float64 `json:"current_price"`
}

func getCartItem() {
	const url string = "https://ujjwaljamuar.github.io/JSONsAPIs/ProductItems.json"
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)

	resp, err := client.Do(req)
	checkError(err)

	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	checkError(err)
	// content := string(respBytes)

	// carItemsArr := make([]CartItem, 0)
	// decoder := json.NewDecoder(strings.NewReader(content))

	// _ , err = decoder.Token()
	// checkError(err)

	// var ct CartItem
	// for decoder.More() {
	// 	err := decoder.Decode(&ct)
	// 	checkError(err)
	// 	carItemsArr = append(carItemsArr, ct)
	// }
	var carItemsArr []CartItem
	err = json.Unmarshal(respBytes, &carItemsArr)
	checkError(err)

	for _, ct := range carItemsArr {
		id := ct.Id
		index := ct.Index
		name := ct.Name
		image := ct.Image
		currentPrice := ct.CurrentPrice

		fmt.Printf("===================================\n")
		fmt.Printf("id: %v\n", id)
		fmt.Printf("index: %v\n", index)
		fmt.Printf("name: %v\n", name)
		fmt.Printf("image: %v\n", image)
		fmt.Printf("currentPrice: %v\n", currentPrice)
		fmt.Printf("===================================\n")
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
