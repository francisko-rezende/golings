// slices2
// Make me compile!

package main

import "fmt"

func main() {
	names := [4]string{"John", "Maria", "Carl", "Peter"}
	totalNames := len(names)
	lastTwoNames := names[totalNames-2 : totalNames] // after figuring out the answer, try with other low/high bounds
	fmt.Println(lastTwoNames)
}
