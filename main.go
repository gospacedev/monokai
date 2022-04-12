package main

import (
	"bufio"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "os"
)

type Account struct {
    Name string   `json:"name"`
	EmailAddress string `json:"emailAddress"`
	PhoneNumber string `json:"phoneNumber"`
    Address   *Address `json:"address"`
}

type Address struct {
    City string `json:"city"`
	Country string`json:"country"`
}

func main() {
	csvFile, _ := os.Open("data.csv")

    reader := csv.NewReader(bufio.NewReader(csvFile))
    var data []Account

    for {
        line, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
        data = append(data, Account{
            Name: line[0],
            EmailAddress: line[1],
			PhoneNumber: line[2],
            Address: &Address{
                City:  line[3],
                Country: line[4],
            },
        })
	}
		//MarshalIndent adds json newlines
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", b)
}