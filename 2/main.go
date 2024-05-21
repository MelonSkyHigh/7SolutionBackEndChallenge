package main

import (
	"fmt"
	"strings"
)

func main() {
	var data string
	fmt.Scan(&data, '\n')
	// fmt.Print(data)
	var arr_list []int
	for i := 0; i <= len(data); i++ {
		arr_list = append(arr_list, 0)
	}
	// fmt.Println(arr_list)

	for {
		done_flag := true
		// fmt.Println("looping ", arr_list)
		for i := 0; i < len(data); i++ {
			// fmt.Println(arr_list[i])
			if string(data[i]) == "L" {
				if !(arr_list[i] > arr_list[i+1]) {
					arr_list[i]++
					done_flag = false
				}
			} else if string(data[i]) == "R" {
				if !(arr_list[i+1] > arr_list[i]) {
					arr_list[i+1]++
					done_flag = false
				}
			} else {
				if !(arr_list[i] == arr_list[i+1]) {
					arr_list[i] = max(arr_list[i], arr_list[i+1])
					arr_list[i+1] = max(arr_list[i], arr_list[i+1])
					done_flag = false
				}
			}
		}

		if done_flag {
			break
		}
	}
	// fmt.Println("finally ", arr_list)
	var result string
	result = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr_list)), result), "[]")
	fmt.Println(result)
}
