package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func encodeBase64() string {

	if len(os.Args) < 2 {
		fmt.Println("You missed file name")
	}

	file, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		fmt.Println(err)
	}

	encode := base64.StdEncoding.EncodeToString(file)

	return encode

}

func decodeBase64() {

	// if len(os.Args) < 2 {
	// 	fmt.Println("You missed file name")
	// }

	// file, err := ioutil.ReadFile(os.Args[1])

	// if err != nil {
	// 	fmt.Println("Can't read file:", os.Args[1])
	// 	fmt.Println(err)
	// }

	s, err := io.Reader(os.Args[1])
	
	decode, _ := base64.StdEncoding.DecodeString(s)

	fmt.Println(string(decode))
}

func main() {

	fmt.Println(decodeBase64)

}
