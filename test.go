package main

import (
	"fmt"
)

func main0() {
	//var fileName = "poisson"
	var poissonData []string
	//if fileName != "" {
	//	data, err := ioutil.ReadFile(fileName)
	//	if err != nil {
	//		fmt.Println(fileName, "File reading error", err)
	//	}
	//	poissonData = strings.Split(string(data), ",")
	//}
	fmt.Printf("len:%d\n", len(poissonData))
	for _, s := range poissonData {
		fmt.Println(s)
	}
}
