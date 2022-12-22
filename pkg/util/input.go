package util

import "os"

func GetInputContent() string {
	data, err := os.ReadFile("input.txt")
	check(err)
	return string(data)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
