package main

import (
	"github.com/nguli-team/bakalo/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
