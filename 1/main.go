package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// data_json := `[[59], [73, 41], [52, 40, 53], [26, 53, 6, 34]]`
	data_json, err := os.ReadFile("hard.json")
	if err != nil {
		panic(err)
	}
	var data [][]int
	json.Unmarshal([]byte(data_json), &data)

	for i := len(data) - 1; i >= 0; i-- {
		for j := 0; j < len(data[i])-1; j++ {
			// fmt.Println(data[i][j], data[i][j+1], " comp to ", data[i-1])
			data[i-1][j] += max(data[i][j], data[i][j+1])
			// fmt.Println("J is ", j, " with array ", data[i])
			// fmt.Println(data[i][j])
			// max(data[i][j], data[i][j+1])
		}
		// fmt.Println(data[i])
	}
	fmt.Println(data[0][0])
}
