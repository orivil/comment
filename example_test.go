package comment_test

import (
	"os"
	"fmt"
	"gopkg.in/orivil/comment.v0"
	"log"
)


// controller comment
type controller struct {

}

// @route {get}/
func (c *controller) index() {

}

// add two parameters
func add(a, b int) int {
	return a + b
}

func ExampleGetDirComment() {
	fileFilter := func(f os.FileInfo) bool {
		// match all file
		return true
	}

	structComment, methodComment, functionComment, err := comment.GetDirComment(fileFilter, "./")
	if err != nil {
		log.Fatal(err)
	}

	// get struct comment
	fmt.Println(structComment["comment_test.controller"] == "controller comment\n")

	// get method comment
	fmt.Println(methodComment["comment_test.controller"]["index"] == "@route {get}/\n")

	// get function comment
	fmt.Println(functionComment["comment_test.add"] == "add two parameters\n")

	// Output:
	// true
	// true
	// true
}
