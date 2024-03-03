package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	var priorities = map[int] string{
		1: "low",
		2: "medium",
		3: "high",
	}
	reader := bufio.NewReader(os.Stdin);
	fmt.Print("Enter a priority number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inputNum, err := strconv.Atoi(input)
	if err!=nil{
		fmt.Println("Error occured: ", err)
	}
	fmt.Println(priorities[inputNum])
}
