// 1. To complete this challenge, 
// you must save a line of input from stdin to a variable, 
// print Hello, World. on a single line, and finally print the value of your variable on a second line.`

package main

import "fmt"

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    // String is defined and later printed out.Note the difference between the Println and Printf.
    
    // line of input from STDIN is saved to var s
    s := bufio.NewScanner(os.Stdin)
    // printing "Hello World"
    fmt.Println("Hello, World.")
    //print the value of the variable
    for s.Scan() {
        fmt.Println(s.Text())
    }
}

//Optionally,you can do:
func main() {
	fmt.Println("Hello, World.")
	io.Copy(os.Stdout, os.Stdin)
}

// Or 
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("data/a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
	}
	file.Close()

}

********

To complete the code in the editor below. The variables i, d, and s  are already declared and initialized for you. 
You must:

/* 2.
- Declare  variables: one of type int, one of type double, and one of type String.
- Read  lines of input from stdin (according to the sequence given in the Input Format section below) and initialize your  variables.
- Use the  operator to perform the following operations:
- Print the sum of  plus your int variable on a new line.
- Print the sum of  plus your double variable to a scale of one decimal place on a new line.
- Concatenate  with the string you read as input and print the result on a new line.
*/
package main


import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    // Declare second integer, double, and String variables.
    var x uint64
    var y float64
    var z string
    
    // Read and save an integer, double, and String to your variables.
    fmt.Scan(&x)
    fmt.Scan(&y)
	// scan everything
    scanner.Scan()
    z =scanner.Text()
    //for scanner.Scan() {
       // fmt.Println(scanner.Text()) // Println will add back the final '\n'
    //}
    // Print the sum of both integer variables on a new line.
   fmt.Printf("%d", i+x)
     // Print the sum of the double variables on a new line.
   
    fmt.Printf("\n%.1f", d+y)
    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.

fmt.Printf("\n%s%s", s, z)
}
********************
/* 3.
Given the meal price (base cost of a meal), tip percent (the percentage of the meal price being added as tip), and tax percent (the percentage of the meal price being added as tax) for a meal, find and print the meal's total cost.

/*Note: Be sure to use precise values for your calculations, or you may end up with an incorrectly rounded result!

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

// Complete the solve function below.
func solve(meal_cost float64, tip_percent int32, tax_percent int32) {

tipCost := (meal_cost * float64(tip_percent)) / 100
    taxCost := (meal_cost * float64(tax_percent)) / 100
    totalCost := meal_cost + float64(tipCost) + float64(taxCost)
    fmt.Printf("%d", int32(math.Round(totalCost)))
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    meal_cost, err := strconv.ParseFloat(readLine(reader), 64)
    checkError(err)

    tip_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    tip_percent := int32(tip_percentTemp)

    tax_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    tax_percent := int32(tax_percentTemp)

    solve(meal_cost, tip_percent, tax_percent)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}


**************************
