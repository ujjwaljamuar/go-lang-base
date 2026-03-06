package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "./9-file-io/fromString.txt"
	file, err := os.Create(fileName)

	// wait till everything in this func gets executed
	defer file.Close()
	checkError(err)

	length, err := io.WriteString(file, "Hello from Go!")

	fmt.Printf("Wrote a file with %v characters\n", length)

	// readFileFromWeb()
	content := readFileFromWeb()
	tours := toursFromJson(content)

	printTours(tours)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// read a text file from web

func readFileFromWeb() string {
	const url = "http://services.explorecalifornia.org/json/tours.php"

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)

	req.Header.Set("User-Agent", "")
	resp, err := client.Do(req)
	checkError(err)

	// defer means do this at the end of everything else
	defer resp.Body.Close()

	fmt.Printf("Response type : %T\n", resp)

	bytes, err := io.ReadAll(resp.Body)
	checkError(err)

	content := string(bytes)

	// fmt.Print(content)
	return content
}

func printTours(tours []Tour) {
	for _, tour := range tours {
		price, _ := strconv.ParseFloat(tour.Price, 16)
		fmt.Printf("%v: ($%v)\n", tour.Name, price)
	}
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0)
	decoder := json.NewDecoder(strings.NewReader(content))

	_, err := decoder.Token()
	checkError(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}

	return tours
}

type Tour struct {
	Name, Price string
}
