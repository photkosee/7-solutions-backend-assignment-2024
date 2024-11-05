package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Function to calculate the maximum path sum from top to bottom of the tree
func maxPathSum(tree [][]int) int {
	if len(tree) == 0 {
		return 0
	}

	// Starting from the second last row up to the top
	// update each value to be the max path sum
	// from that node down to the bottom
	for row := len(tree) - 2; row >= 0; row-- {
		for col := 0; col < len(tree[row]); col++ {
			// Choose the maximum path of the two possible child nodes
			tree[row][col] += max(tree[row+1][col], tree[row+1][col+1])
		}
	}

	return tree[0][0]
}

func main() {
	// Get the JSON data
	fileData, err := os.ReadFile("../files/hard.json")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}

	// Parse the JSON data
	var jsonTree [][]int
	json.Unmarshal(fileData, &jsonTree)

	manualTree := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}
	
	resultManual := maxPathSum(manualTree)
	resultFetch := maxPathSum(jsonTree)

	fmt.Printf("The maximum path sum from the manual input is: %d\n", resultManual)
	fmt.Printf("The maximum path sum from the json file is: %d\n", resultFetch)
}
