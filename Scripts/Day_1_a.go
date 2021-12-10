package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("Source Files\\Day_1_a_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)
	l := list.New()
	var mySlice = []int{}
	for _, line := range strings.Split(strings.TrimSuffix(text, "\r\n"), "\r\n") {
		lineInt, err := strconv.ParseInt(line, 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		var myInt = int(lineInt)
		l.PushBack(myInt)
		mySlice = append(mySlice, myInt)
	}
	var previousIndicator = 999999999
	var numIncreases = 0
	for i := 0; i < len(mySlice); i++ {
		var currentIndicator = mySlice[i]
		if currentIndicator > previousIndicator {
			numIncreases += 1
		}
		previousIndicator = currentIndicator
	}
	fmt.Println(mySlice)
	fmt.Println(strconv.Itoa(numIncreases))
}
