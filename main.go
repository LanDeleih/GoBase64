package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func encodeBase64() string {

	if len(os.Args) < 2 {
		fmt.Println("You missed file name")
	}

	file, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		log.Fatalf("Can't read file:", os.Args[1])
	}

	encode := base64.StdEncoding.EncodeToString(file)

	return encode

}

func decodeBase64() {

	read, err := ioutil.Read(os.Args[1:])
	if err != nil {
		log.Fatalf("Cannot read programm args")
	}

	scanner := bufio.NewScanner(read)

	scanner.Scan()
	s := scanner.Text()

	decode, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatalf("Cannot decode input", err)
	}

	fmt.Println(string(decode))
}

func main() {

}
