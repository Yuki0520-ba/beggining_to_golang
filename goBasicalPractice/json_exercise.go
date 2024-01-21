package goBasicalPractice

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Country string `json:"country"`
	State   string `json:"state"`
}
type PersonInfo struct {
	Name      string  `json:"name"`
	BitrthDay string  `json:"birthday"`
	Address   Address `json:"address"`
}

func parseJson() {
	var myData = `
	{
		"name": "John",
		"birthday": "1998-04-12",
		"address": {
			"country": "America",
			"state": "NewYork"
		}
	}
	`
	var persedData PersonInfo
	err := json.Unmarshal([]byte(myData), &persedData) // シンプルな文字列をパースする場合json.Unmarshalを利用
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(persedData)
}

func readJsonFile(filename string) {
	fmt.Println(os.Getwd())
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	var persedData PersonInfo
	err = json.NewDecoder(f).Decode(&persedData) // Jsonファイルを読み込む場合デコーダーを利用してパースする
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(persedData)
	fmt.Println(persedData.Address.Country)

}

func jsonExercise() {
	parseJson()
	// readJsonFile("goBasicalPractice/sample-data.json")
}
