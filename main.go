package main

import "fmt"

func countItems(items []string) map[string]int {
	result := map[string]int{}

	for _, item := range items {
		result[item]++
	}

	return result
}

func main() {
	items := []string{"milk", "eggs", "milk", "bread", "eggs", "milk"}
	fmt.Println(countItems(items))
}
