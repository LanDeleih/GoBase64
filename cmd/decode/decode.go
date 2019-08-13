package decode

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

//Base64 decode base64 function
func Base64() {

	read := os.Args[1:]

	for i := range read {

		decode, err := base64.StdEncoding.DecodeString(read[i])
		if err != nil {
			log.Fatalln("Cannot decode input", err)
		}

		fmt.Println(string(decode))
	}

}
