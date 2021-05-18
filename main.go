package main

import "bakalo.li/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
