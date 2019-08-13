package encode

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Base64 this function will encode files to base64
func Base64() string {

	if len(os.Args) < 2 {
		fmt.Println("You missed file name")
	}

	file, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		log.Fatalln("Can't read file:", os.Args[1])
	}

	encode := base64.StdEncoding.EncodeToString(file)

	return encode

}
