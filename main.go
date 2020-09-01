package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var nbool = flag.Bool("n", true, "paragraph")

func onParagraph(array []string) {
	for i := range array {
		f, err := os.Open(array[i])
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		count := 0
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			count++
			fmt.Printf("%v: %v\n", strconv.Itoa(count), scanner.Text())
		}
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func offParagraph(array []string) {
	for i := range array {
		text, err := ioutil.ReadFile(array[i])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(text))
	}
}

func main() {

	flag.Parse()
	if *nbool == true {
		onParagraph(flag.Args())
	} else {
		offParagraph(flag.Args())
	}
}
