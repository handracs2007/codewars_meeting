package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func Meeting(s string) string {
	// Make it uppercase
	s = strings.ToUpper(s)

	// separate by ; to get all the names
	names := strings.Split(s, ";")

	// Loop through all and make a string of lastname, firstname
	var allNames = make([]string, len(names))
	for _, name := range names {
		nameParts := strings.Split(name, ":")
		firstName := nameParts[0]
		lastName := nameParts[1]

		allNames = append(allNames, "("+lastName+", "+firstName+")")
	}

	sort.Strings(allNames)

	var buffer bytes.Buffer
	for _, name := range allNames {
		buffer.WriteString(name)
	}

	return buffer.String()
}

func main() {
	fmt.Println(Meeting("Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill"))
}
