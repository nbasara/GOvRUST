package main

import (
	"bufio"
	"fmt"
	"os"
)

// run using go run ch1.go arg1 arg2
// exercise change this to print the name of the file in which the ducplicate
// name occurs

func main() {
    counts := make(map[string]int)
    locations := make(map[string][]string)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts, locations, "stdin")
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup1: %v\n", err)
                continue
            }
            countLines(f, counts, locations, arg)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\t%v\n", n, line, locations[line])
        }
    }
    //printArgs()
}

func countLines(f *os.File, counts map[string]int,
locations map[string][]string, loc string) {
    // reads from file f
    input := bufio.NewScanner(f)
    // iterates through the file
    for input.Scan() {
        // increments map entry by 1
        entry := input.Text()
        counts[entry]++
        inArray := false
        for _, name := range locations[entry] {
            if name == loc {
                inArray = true
                break
            }
        }
        if !inArray {
            locations[entry] = append(locations[entry], loc)
        }
    }
}

func printArgs() {
    // initialize new variables to empty strings
    // the variables get the string type through short variable declaration
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}
