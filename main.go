package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetInput(misaka chan string) (input string) {

	fmt.Printf("Add your cc list: ")
	fmt.Scan(&input)
	resp, err := http.Get(input)

	if err != nil {
		log.Fatal(err)
		misaka <- fmt.Sprintf("Error: %v", err)
		return
	}

	defer resp.Body.Close()

	var buf bytes.Buffer

	_, err = io.Copy(&buf, resp.Body)

	if err != nil {
		log.Fatal(err)
		misaka <- fmt.Sprintf("Error reading response body: %v", err)
		return
	}

	responseString := buf.String()

	scanner := bufio.NewScanner(&buf)

	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	misaka <- fmt.Sprintf("%v\n\nTotal CC: %v", responseString, lineCount)

	close(misaka)

	return

}

func main() {

	misaka := make(chan string)

	go GetInput(misaka)

	message := <-misaka

	fmt.Println(message)

}
