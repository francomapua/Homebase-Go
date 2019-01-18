package main

/*	WORKSPACES
- One per Go Installation, under GOPATH
	- Find via command line: go env GOPATH
- all projects live here
- Grants commands:
	go run <file>
	go build
	go install
- Folder Structure

go/
	src/
		*packages go in here
	bin/
		* installs go in here

*/

/*	PACKAGES
- defn. Collections of code
- Executable projects must have one .go file with "package main" to run
- all folders under GOPATH are considered packages
*/

/*	IMPORTS
- add other packages into your scope
*/
import (
	"math"
	"crypto"
	"fmt"
)

// MAIN FUNCTION
func main(){
	structs()
}
/*	EXPORTED NAMES
- Capitalize in order to export something
math.pi <- not exportable
math.Pi <- exportable
- not considered as Public and Static for some reason
*/

/*	VARIABLES	*/
func declareVariables(){
	var d int 							// Var d, an int
	var a, b, c int 					// Var a, b, and c, ints 
	var i, j int = 1 ,2 				// i = 1, j = 2
	i, j = 3, 4							// i = 3, j = 4							<- Not allowed at package level
	var x, y, z = false, 2, "no!"		// x = false, y = 2, z = "no!"
	k := 3								// Inferred.  Variables must be new		<- Not allowed at package level
}



/*	FUNCTIONS	*/
func add(x int, y int) int { // Function add with arguments x (an int)
	return x + y
}
// Multiple Return
func addAndSubtract(x int, y int) (int, int){ 
	return x + y, x - y
}
// Multiple Named Return; Used as x, y := addAndSubstractAgain
func addAndSubtractAgain(x, y int) (sum, difference int){ 
	sum = x + y
	difference = x - y
	return sum, difference
}
// Defering things to the termination of the function
func deferFunc(){
	defer fmt.Println(" coming") // is added to the defer stack
	defer fmt.Println(" is")
	fmt.Println("Winter") // Winter
} // is coming
// Functions are values too
func funcsAsVals(){
	hypot := func(x, y float64) float64{
		return math.Sqrt(x * x + y * y)
	}
	fmt.Println(hypot(5, 12))
}
// Package Level -- No Logic here
//func(){
	// Function Level
//}()

/*	TYPES
- will default to Zero Values if not initialized
*/
func valueTypes(){
	var boolean bool // Defaults to false
	//true && true
	//true || true
	//!false
	var sstring string // Defaults to ""
	var integer int // Defaults to 0
	var uinteger uint
	var bbyte byte
	var rrune rune
	var float float32 // Or float64
	var complex complex64 // or complex 128
	var newString  = string(integer) // Type Conversion, Explicit
	//nil Is null
}

/*	CONSTANTS	*/
const yourname string = "Franco"

/*	CONDITIONALS */
func conditionals(){
	// Standard Structure
	if 1 == 1 {

	} else if 2 == 2{

	} else {

	}

	// Wit Short Statement before condition
	y := 1
	if x:=1; x < y{

	}

	// Switch
	switch os := runtime.GOOS; os {
		case "linux" :
			fmt.Println("Hey I'm")
		case "windows" :
			fmt.Println("Not consistent")
		default:
			fmt.Println("Or nots")
	}
}

/*	ARRAYS	
- fixed size
*/
func arrays(){
	var arrayOfFiveInts [5]int // a 0-index array; "arrayOfFiveInts :=" works too
	fmt.Println(arrayOfFiveInts) // 0 0 0 0 0
	thisWorksToo := [5]int
	var initializedArray = [5]int {0, 1, 2, 3, 4}
	fmt.Println(initializedArray) // 0 1 2 3 4
}
/*	SLICES
- inherits from arrays
- can be appended
*/
func slices(){
	var sliceOfInt []int
	var initializedSlice = []int {5, 4, 3, 2, 1}
	initializedSlice = append(initializedSlice, 0)
	fmt.Println(initializedSlice) // 5 4 3 2 1 0
}
/*	MAPS	*/	
func maps(){
	mapOfStringsToInts := make(map[string] int)
	mapOfStringsToInts["two"] = 2
	fmt.Println(mapOfStringsToInts["two"]) // = 2
}

/*	LOOPS	*/
func loops(){
	// Basic Loops
	for i := 0; i < 5; i++ {
		fmt.Println(i) // 0, 1, 2 ,3 ,4
	}
	
	// Iterating over a collection
	var initializedSlice = []int {5, 4, 3, 2, 1}
	for index, value := range initializedSlice {
		fmt.Println(index) // 0, 1 ,2, 3, 4
		fmt.Println(value) // 5, 4 ,3 ,2 ,1
	}
	
	// "While"
	for ; limit < 5; {
		limit = limit + 1
	}
	// OR
	for limit < 5 {
		limit = limit + 1
	}
}

/*	ERRORS */
func errors(){
	someError := errors.New("Error")
}
func errorReportingFunction()(int, error){
	return 5, nil // OR nil, new Error
}

/*	STRUCTS
- The classes of Go
- are declared package level
*/
type person struct{
	// Attributes
	name string	
	age int		
	address string
}
func (this person) BarkNameAndAge() string{
	return this.name + this.age
}
func (this person) singAddress(){
	fmt.Println(this.address)
}

func structs(){
	george := person{"George", 12}
	fmt.Println(george.age) // 12
	fmt.Println(george.BarkNameAndAge())
	george.singAddress()

	alice := person{name : "Alice", age : 12}
}

/*	POINTERS
- Behave like they do in C
*/
func pointers(){
	i = 1;
	// Continue later
}

/* METHOD RECEIVERS
- attaches functions to structs to make them methods
- can be pointer or non-pointer
	- pointer functions can change the struct they are attached to
*/
type Mutatable struct{
	a int
	b int
}
func (m Mutatable) StayTheSame() {
    m.a = 5
    m.b = 7
}
func (m *Mutatable) Mutate() {
    m.a = 5
    m.b = 7
}

func methodReceivers(){
	m := &Mutatable{0, 0}
    fmt.Println(m) // {0,0}
    m.StayTheSame()
    fmt.Println(m) // {0,0}
    m.Mutate()
    fmt.Println(m) // {5, 7}
}



