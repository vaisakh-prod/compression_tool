package main

import (
	"fmt"
	"os"
)

var stringText string

type characterNode struct {
	character rune
	count     int
}

var characterTree []characterNode

func init() {
	filename := os.Args[1]
	fmt.Printf("the filename mentioned is %v\n", filename)
	if filename == "" {
		panic("Filename not provided")
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("error in reading file: %v", err.Error()))
	}
	stringText = string(data)
}

func main() {
	characterMap := make(map[rune]int)
	if stringText != "" {
		for _, char := range stringText {
			characterMap[char]++
		}
		for char, count := range characterMap {
			characterTree = append(characterTree, characterNode{character: char, count: count})
		}
	}
	for _, node := range characterTree {
		fmt.Printf("%c %v\n", node.character, node.count)
	}
}
