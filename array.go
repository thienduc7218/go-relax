package main

import "fmt"

func main() {
	// ----------Arrays------------
	// Define an array of size 4 that store deployment options
	var myArr = [4]string{"R-pi", "AWS", "GCP", "Azure"}

	// Define an array and let the compiler count its size
	myArr2 := [...]string{"R-pi", "AWS", "GCP", "Azure"}

	// Loop through the options array
	for i := 0; i < len(myArr); i++ {
		option := myArr[i]
		fmt.Println(i, option)
	}

	for index, option := range myArr2 {
		fmt.Println(index, option)
	}
}
