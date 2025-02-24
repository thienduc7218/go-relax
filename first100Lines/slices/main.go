package main

import (
	"fmt"
	"reflect"
)

func main() {
	//------------Slices-----------
	// Define an array containing programming languages
	languages := [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust", // Must include the trailing comma
	}

	// Define slices
	classics := languages[0:3]  // or languages[:3]
	modern := make([]string, 4) // len(modern) = 4
	modern = languages[3:7]     // include 3 exclude 7
	new := languages[7:9]       // or languages[7:]

	fmt.Printf("classic languagues: %v\n", classics)
	fmt.Printf("modern languages: %v\n", modern)
	fmt.Printf("new languages: %v\n", new)

	//-----snip-------
	langArr := languages[:] //copy of the arr
	fmt.Println(reflect.TypeOf(langArr).Kind())

	// When define with [] if we leave it blank then it results into a slice, if we pass a number in, it results into an array
	frameworks := []string{
		"React", "Vue", "Angular", "Svelte",
		"Laravel", "Django", "Flask", "Fiber",
	}

	jsFrWs := frameworks[0:4:4]               // length 4 capacity 4
	frameworks = append(frameworks, "Meteor") // not possible with arrays

	fmt.Printf("all frameworks: %v\n", frameworks)
	fmt.Printf("js frameworks: %v\n", jsFrWs)
}
