package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"

	"github.com/jszwec/csvutil"
)

type Cars struct {
	Model    string  `csv:"model"`
	Mpg      float32 `csv:"mpg,omitempty"`
	Disp     float32 `csv:"disp,omitempty"`
	Cylinder float32 `csv:"cyl,omitempty"`
	HP       float32 `csv:"hp,omitempty"`
}

// type Cars struct {
// 	Model    string
// 	Mpg      float32
// 	Disp     float32
// 	Cylinder float32
// 	HP       float32
// }

type User struct {
	Name    string
	Address string
	Cars
}

func main() {
	var csvInput = []byte(`model,mpg,cyl,disp,hp 
Mazda RX4,21,6,160,110
Mazda RX4 Wag,21,6,160,110
Datsun 710,22.8,4,108,93
Hornet 4 Drive,21.4,6,258,110
Hornet Sportabout,18.7,8,360,175`)

	fmt.Println("================= Unmarshall result ===================")
	unmarshal(csvInput)
	// marshal()
	fmt.Println("================= Decoder result ===================")
	separatorOptions(csvInput)

}

func unmarshal(csv []byte) {
	var cars []Cars

	if err := csvutil.Unmarshal(csv, &cars); err != nil {
		fmt.Println("error", err)
	}

	for _, c := range cars {
		fmt.Printf("%+v \n", c)
	}
}

func marshal() {
	users := []User{
		{
			Name:    "Budi",
			Address: "Bekasi",
			Cars:    Cars{"Mazda RX4", 21, 6, 160, 110},
		},
		{
			Name:    "Joe",
			Address: "NYC",
			Cars:    Cars{"Mazda RX4", 21, 6, 160, 110},
		},
	}

	fmt.Printf("%+v \n", users)

	b, err := csvutil.Marshal(users)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

}

func separatorOptions(r []byte) {

	csvReader := csv.NewReader(bytes.NewReader(r))
	// csvReader.Comma = '|'

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Fatal(err)
	}

	var cars []Cars

	for {
		var c Cars
		if err := dec.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(c)
		cars = append(cars, c)
	}

	fmt.Printf("%+v \n", cars)
}
