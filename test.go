package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	var fileName = "poisson"
	var poissonData []string
	if fileName != "" {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			print(err)
		}
		poissonData = strings.Split(string(data), ",")
	}
	print(poissonData[19])
}
