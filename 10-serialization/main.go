package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

func main() {
	exampleItem()
}

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
