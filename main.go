package main

import (
	"fmt"

	"github.com/jon-castro/gator-golang/internal/config"
)

func main() {
	read, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = read.SetUser("jon")
	if err != nil {
		return
	}

	setFile, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(setFile)
}
