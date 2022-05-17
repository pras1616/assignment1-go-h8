package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	params := os.Args[1]
	intParams, _ := strconv.Atoi(params)

	type Biodata struct {
		Name      string `json:"Nama"`
		Alamat    string `json:"Alamat"`
		Pekerjaan string `json:"Pekerjaan"`
		Alasan    string `json:"Alasan Memilih Kelas Golang"`
	}

	filePath := "./biodata.json"
	file, err1 := ioutil.ReadFile(filePath)

	if err1 != nil {
		fmt.Printf("// error while reading file %s\n", filePath)
		fmt.Printf("File error: %v\n", err1)
		os.Exit(1)
	}

	var biodatas []Biodata

	err2 := json.Unmarshal(file, &biodatas)

	if err2 != nil {
		fmt.Println("error:", err2)
		os.Exit(1)
	}

	for k := range biodatas {
		if k == intParams-1 {
			jsonData := biodatas[k]
			// var showData Biodata
			// _ = json.Unmarshal([]byte(jsonData), &showData)
			fmt.Println(jsonData)
		}
	}
}

// func ChangesExample(params int, data string) string {
// 	// fmt.Println("contoh function", name)
// 	jsonData := nil
// 	for k := range data {
// 		if k == params-1 {
// 			jsonData = data[k]
// 			// var showData Biodata
// 			// _ = json.Unmarshal([]byte(jsonData), &showData)
// 			fmt.Println(jsonData)
// 		}
// 	}
// 	return jsonData
// }
