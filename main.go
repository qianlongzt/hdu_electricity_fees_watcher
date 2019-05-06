package main

import (
	"fmt"
	"os"
)

func main() {
	room := os.Args[1]
	fmt.Println(getElectInfoFromWeb(GetIDs(room)))
}
