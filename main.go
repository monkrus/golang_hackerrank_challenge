// 0. To complete this challenge, 
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

s := bufio.NewScanner(os.Stdin)
    fmt.Println("Hello, World.")
    for s.Scan() {
        fmt.Println(s.Text())
    }
}
	
// Variable `s`  is assigned to the  stdin input
// Print out the first line
// Print out the content of the inputString variable

********

To complete the code in the editor below. The variables i, d, and s  are already declared and initialized for you. 
You must:

/* 1.
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
/* 2.
Given the meal price (base cost of a meal), tip percent (the percentage of the meal price being added as tip),
and tax percent (the percentage of the meal price being added as tax) for a meal, find and print the meal's total cost.

Note: Be sure to use precise values for your calculations, or you may end up with an incorrectly rounded result!*/

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"itHub, Inc.
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

/*
3. 
Given an integer, , perform the following conditional actions:
If  is odd, print Weird
If  is even and in the inclusive range of  to , print Not Weird
If  is even and in the inclusive range of  to , print Weird
If  is even and greater than , print Not Weird
Complete the stub code provided in your editor to print whether or not  is weird.

*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)
	if N%2 != 0 {
		fmt.Printf("Weird")
	} else {
		if N >= 2 && N <= 5 {
			fmt.Printf("Not Weird")
		} else if N >= 6 && N <= 20 {
			fmt.Printf("Weird")
		}
		if N > 20 {
			fmt.Printf("Not Weird")
		}
	}
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

*******************8
/* 4.
Write a Person class with an instance variable, aga, and a constructor that takes an integer, initialAge, as a parameter.

The constructor must assign initialAge to age after confirmaing the argument passed as initialAge is not negative.
If a negative argument is passed initialAge, the constructor should set age to 0 and print `age is not valid`, setting age to 0.

In addition, you should perform the following actions:
- yearPasses() should increase the age variable  by 1
- amIOld() should perfomr the following actions

If age<13 , print You are young..

If  and age>13, print You are a teenager..

Otherwise, print You are old..

*/
package main
import "fmt"

type person struct {
	age int
}


func (p person) NewPerson(initialAge int) person {
    //Add some more code to run some checks on initialAge
    if initialAge < 0 {
        fmt.Println("Age is not valid, setting age to 0.")
        p.age = 0
    }

    return p
}

func (p person) amIOld() {
    //Do some computation in here and print out the correct statement to the console
    if p.age < 13 {
        fmt.Println("You are young.")
    } else if p.age >= 13 && p.age < 18 {
        fmt.Println("You are a teenager.")
    } else {
        fmt.Println("You are old.")
    }
}

func (p person) yearPasses() person {
    //Increment the age of the person in here
    p.age++

    return p
}


func main() {
    var T,age int

    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        fmt.Scan(&age)
        p := person{age: age}
        p = p.NewPerson(age)
        p.amIOld()
        for j := 0; j < 3; j++ {
            p = p.yearPasses()
        }
        p.amIOld()
        fmt.Println()
    }
}


/* 5.
Given an integer, n, print its first 10 multiples. 
Each multiple n x i (where 1<= i <=10) should be printed on a new line in the form: n x i = result.
*/

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)
    for result := 1; result <= 10; result++ {
    fmt.Printf("%d x %d = %d\n", n, result, n*int32(result))
}
    
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


/*  6.

Task
Given a string,S, of length N that is indexed from 0 to N-1 , 
print its even-indexed and odd-indexed characters as  space-separated strings 
on a single line (see the Sample below for more detail).

Note: 0 is considered to be an even index.
*/
package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    // set variables for  AMOUNT if TESTS and LENGTH of repetitions
    var tests, length int
    // set variables for INPUT, also for EVEN and ODD characters
    var input, even, odd string
    //  scans(reads) the input 
    scanner := bufio.NewScanner(os.Stdin)
    //  reads the amount of TESTS
    fmt.Scanf("%d", &tests)
    //starts a cycle 
    for i :=0; i<tests; i++{
    //sets even/odd to empty value
        even = ""
        odd = ""
    //scans 
        scanner.Scan()
    // input equals to scanned text
        input = scanner.Text()
    // set a length
        length = len(input)
    //split all the chars
        chars := strings.Split(input, "")
     // start next cycle to find odd/even characters
        for j:=0; j<length; j++{
            if j%2 == 0{
                even += chars[j]
            }else{
                odd += chars[j]
            }
        }
     // final print out
        fmt.Println(even, odd)
    }
}

/* 7.

Given the and array A of N integers, print A`s elements in reverse order as a single line of space-separated numbers
*/

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }
   
    for i := len(arr); i > 0; i-- {
		fmt.Printf("%d ", arr[i-1])
    }
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



/* 8.
Given n names and phone numbers, assemble a phone book that maps friends' names to their respective phone numbers. 
You will then be given an unknown number of names to query your phone book for. For each  `name` queried, print the associated
entry from your phone book on a new line in the form name=phoneNumber; 
if an entry for `name` is not found, print `Not found` instead.

Note: phone book should be a Dictionart/Map/HashMap data structure
*/
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func main() {
    phoneBook := make(map[string]string)
    index := 0
    howMany := 0

    // Grab all inputs in loop
    // Handle each input based on what type of input it is.
    // First one is expected # of entries, then the entries, then the lookups
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        text := scanner.Text()

        switch { // true switch. always run
        case index == 0:
            // first loop. collect number of inputs
            howManyTemp, err := strconv.ParseInt(text, 10, 64)
            checkErr(err)
            howMany = int(howManyTemp)
        case index <= howMany:
            // This is an entry for the phonebook
            // Save input to the phonebook map
            split := strings.Split(text, " ")
            phoneBook[split[0]] = split[1]
        default:
            // All other inputs after the last entry
            // These are the lookup requests
            // Check each one against the phonebook
            number, test := phoneBook[text]
            if test == true {
                fmt.Printf("%s=%s\n", text, number)
            } else {
                fmt.Printf("Not found\n")
            }
        }
        // Increment the index
        index++
    }
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }
    return strings.TrimRight(string(str), "\r\n")
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

/* 9.
// Complete the factorial function below.
*/
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the factorial function below.

var result int32 =0
func factorial(n int32) int32 {
 if (n > 0) {
        result = n * factorial(n-1)
     return result
 }
    
    return 1

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    result := factorial(n)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
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

/* 10.
Given a base- integer, n , convert it to binary (base-). 
Then find and print the base-2 integer denoting the maximum number of consecutive 1's in n's binary representation.
*/
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    
    count := 0
    max := 0
    bi := strconv.FormatInt(nTemp, 2)
    for _, s := range bi {
        if string(s) == "1" {
            count++
        } else {
            if count > max {
                max = count
            }
            count = 0
        }

        if count > max {
            max = count
        }
    }
    fmt.Println(max)
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
/*   11.
Calculate the hourglass sum for every hourglass in A,then print the maximum hourglass sum.
*/
package main

import "fmt"

func main() {
    //create a 6x6 array, read standard input, write to our array
    A := [6][6]int{}
    for i := 0; i < 6; i++ {
        for j := 0; j < 6; j++ {
            fmt.Scanf("%d", &A[i][j])
        }
    }

    //looking for the maximum value of hourglass in our array
    var maximumHourglass, currentHourglass int
    maximumHourglass = -63 //set minimum possible value of hourglass with all elements -9
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            currentHourglass = A[i][j] + A[i][j+1] + A[i][j+2] + A[i+1][j+1] + A[i+2][j] + A[i+2][j+1] + A[i+2][j+2]
            if currentHourglass > maximumHourglass {
                maximumHourglass = currentHourglass
            }
        }
    }

    //print our maximum value to STDOUT
    fmt.Println(maximumHourglass)
}

/*
Please see the execution order for the abovementioned code

i+2   [5] [6] [7]
i+1       [4]
i     [3] [2] [1]
       j  j+1 j+2
*/

/* 12.
You are given two classes, Person and Student, where Person is the base class and Student is the derived class. 
Completed code for Person and a declaration for Student are provided for you in the editor.
Observe that Student inherits all the properties of Person.

Complete the Student class by writing the following:

A Student class constructor, which has 4 parameters:
A string,firstName .
A string,lastName .
An integer,id .
An integer array (or vector) of test scores, scores.
A char calculate() method that calculates a Student object's average and returns the grade character
representative of their calculated average:


Letter Average
 O    90<=a<=100
 E    80<=a<90
 A    70<=a<80
 P    55<=a<70
 D    40<=a<55
 T      a<40
*/
'use strict';

var _input = '';
var _index = 0;
process.stdin.on('data', (data) => { _input += data; });
process.stdin.on('end', () => {
    _input = _input.split(new RegExp('[ \n]+'));
    main();    
});
function read() { return _input[_index++]; }

/**** Ignore above this line. ****/

class Person {
    constructor(firstName, lastName, identification) {
        this.firstName = firstName;
        this.lastName = lastName;
        this.idNumber = identification;
    }
    
    printPerson() {
        console.log(
            "Name: " + this.lastName + ", " + this.firstName 
            + "\nID: " + this.idNumber
        )
    }
}

class Student extends Person {
    /*	
    *   Class Constructor
    *   
    *   @param firstName - A string denoting the Person's first name.
    *   @param lastName - A string denoting the Person's last name.
    *   @param id - An integer denoting the Person's ID number.
    *   @param scores - An array of integers denoting the Person's test scores.
    */
    // Write your constructor here
constructor(firstName, lastName, id, scores) {
    super(firstName, lastName, id);
    this.scores = scores;
  }
    /*	
    *   Method Name: calculate
    *   @return A character denoting the grade.
    */
    // Write your method here
     calculate() { // using arrow function as class property gives an error because it needs babel-eslint parser instead
    const average =
      this.scores.reduce((acc, score) => acc = acc + score, 0) / this.scores.length;

    if (average >= 90 && average <= 100) {
      return 'O';
    } else if (average >= 80) {
      return 'E';
    } else if (average >= 70) {
      return 'A';
    } else if (average >= 55) {
      return 'P';
    } else if (average >= 40) {
      return 'D';
    }
    return 'T';
  }
}

function main() {
    let firstName = read()
    let lastName = read()
    let id = +read()
    let numScores = +read()
    let testScores = new Array(numScores)
    
    for (var i = 0; i < numScores; i++) {
        testScores[i] = +read()  
    }

    let s = new Student(firstName, lastName, id, testScores)
    s.printPerson()
    s.calculate()
    console.log('Grade: ' + s.calculate())
}

/*
13. 
Given a Book class and a Solution class, write a MyBook class that does the following:

Inherits from Book
Has a parameterized constructor taking these  parameters:
string title
string author
int price
Implements the Book class' abstract display() method so it prints these  lines:
, a space, and then the current instance's title .
, a space, and then the current instance's author.
, a space, and then the current instance's proce.
*/
