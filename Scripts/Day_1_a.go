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
	myArray := [...]int{}
	var mySlice = []int{}
	// l.PushBack(list_item)
	for _, line := range strings.Split(strings.TrimSuffix(text, "\n"), "\n") {
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
	for i := 0; i < len(l); i++{
		var currentIndicator = l[i]
		currentIndicator, err := strconv.ParseInt(nth i l, 0, 64)
	}
	fmt.Println(text)
	fmt.Print(mySlice)
}