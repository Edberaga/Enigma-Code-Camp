package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func evenNames(names []string) []string {
    var evenLengthNames []string

    for _, name := range names {
        if len(name)%2 == 0 {
            evenLengthNames = append(evenLengthNames, name)
        }
    }

    return evenLengthNames
}

func main() {
    //Declare the Scanner
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()

	//Write the collection of strings
    x := scanner.Text()
    slice := strings.Split(x, " ")

    // Call the evenNames function to get names with even lengths
    names := evenNames(slice)

    // Print the result
    for _, name := range names {
		fmt.Print(name, " ")
	}
}