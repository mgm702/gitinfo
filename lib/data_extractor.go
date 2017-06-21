package lib

import (
	_ "encoding/json"
	"fmt"
	_ "os"
)

func SayHi() error {
	fmt.Println("Hello from data extractor")
	return nil
}
