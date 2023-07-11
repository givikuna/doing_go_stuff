package main

import (
	"fmt"
	"os/exec"
	"reflect"
	"strconv"
)

type person struct {
	name    string
	age     int
	country string
	married bool
}

func main() {
	var helloWorld string = "Hello World!"
	fmt.Println(helloWorld)

	// you do not need the type to be given, it can be inferred, not even a need to write var.
	// this is not recommended for production code as it may cause unknown type mismatch errors
	helloWorld2 := "Hello World!"
	fmt.Println(helloWorld2)

	// to see if the strings the same, you can compare them with ==
	if helloWorld == helloWorld2 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// you can also assign this to a boolean variable
	var areTheyEqual bool = helloWorld == helloWorld2
	fmt.Println(areTheyEqual)

	// in go you can make arrays just like every other language
	var integerArray [5]int = [5]int{1, 2, 3, 4, 5} // notice how the length here is 5? and how that had to be specified?
	// you can avoid this problem by simply using slices
	var integerArray2 []int = []int{1, 2, 3, 4, 5} // now it can be as long as you want and it'll work just fine

	// you can append other integers into integerArray2 by doing this:
	integerArray2 = append(integerArray2, 6)
	// you need to do it like this, as integerArray2 won't just get a new element added by append but rather the append function literally creates a new array
	// in other words, arrays with []int not having a length is just an abstraction taken care of by the Go devs, because they're super cool

	// you can remove the last element like this:
	integerArray2 = integerArray2[:len(integerArray2)-1]
	// again we are making a new array

	// if you want to remove any index at any element you can do this
	superDuperCoolArray := []int{1, 2, 3}
	fmt.Println("Before: ", superDuperCoolArray)
	index1 := 1 // the index at which we want to remove the element, in this case, it'll remove the number 2, because in go the counting in arrays starts at 0
	superDuperCoolArray = append(superDuperCoolArray[:index1], superDuperCoolArray[index1+1:]...)
	// this looks complicated no? yeah it does
	// needlessly complicated if you ask me
	// this works like this:
	// superDuperCoolArray[:index1] creates an array of all elements up to index1
	// superDuperCoolArray[index1+1:] creates another array starting after index1
	// this leaves the element at index1 out of the array
	// the ... (the ellipsis) is used to expand superDuperCoolArray[index1+1], allowing each of the elements within it to be appended into superDuperCoolArray[:index1] one by one
	// Go arrays work similarly to Python arrays as you can see, outside of the lack of cool built in functions
	fmt.Println("After: ", superDuperCoolArray)

	// now you might wonder if there is an easier way of doing this,
	// and there is, make this into a function, that's all I know

	/*
		beforehand we crated two arrays:
		var integerArray [5]int = [5]int{1, 2, 3, 4, 5} // notice how the length here is 5? and how that had to be specified?
		var integerArray2 []int = []int{1, 2, 3, 4, 5} // now it can be as long as you want and it'll work just fine

		these two arrays CANNOT be compared using == as they are of different types: [5]int and []int
		but they can be compared using a function called DeepEqual from a package called reflect
		this function can compare these two, but integerArray2 and integerArray must be made into slices
		this can  be achieved by writing this: integerArray2[:] and integerArray[:]
		without this, integerArray2 and integerArray will be made out to be of different types meaning they'll simply not be equal
		this gives us a boolean value (true), which can later be made into a string using a function called FormatBool in strconv package
		Go has a LOT of packages, and you don't need to memorize all of them, just google it up if you don't remember which one you need
	*/
	fmt.Println("arrays are equal: " + strconv.FormatBool(reflect.DeepEqual(integerArray[:], integerArray2[:])))

	// this begs the question, can two arrays be compared with ==
	// yes of course, but they must be of the same type.
	// that is why I used reflect.DeepEqual instead of == in the above example
	// but if you'd rather not import a package then you can just make the types the same by typecasting
	// for example
	fmt.Println(integerArray == [5]int(integerArray2)) // true

	// go is great at inferring stuff
	// meaning if you need an array to stay a fixed length, but are to lazy to assign a length or count it yourself do it like this
	someRandomIntegerArray1 := [...]int{1, 2, 3, 4, 5} // this array will never be more than 5, and the length stays fixed, unlike an array with a tyoe []int
	fmt.Println(someRandomIntegerArray1)

	// you can get the length of an array like this
	fmt.Print("length is: ")
	fmt.Println(len(integerArray)) // just like in python

	// in go to print out an array there are two main options (that I know of)
	// ether use a loop to iterate through them one at a time
	// or just print it (this may not always work)

	// print simply:
	fmt.Println(integerArray)

	// or iterate over it
	// every programming language has loops
	// to do loops in go you can do this
	fmt.Print("[")
	for i := 0; i < len(integerArray); i++ {
		fmt.Print(integerArray[i])
		if i != len(integerArray)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")

	fmt.Println()

	// lets make another integer array
	var integerArray3 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("are the two arrays equal?: " + strconv.FormatBool(integerArray3 == integerArray))

	// you can typecast many different types into other tings like this by just putting type(variable)
	// for example
	var someRandomInteger int = 5
	var someRandomFloat float64 = float64(someRandomInteger)
	fmt.Println(someRandomFloat) // gives us 5
	// this can be done in the opposite direction too
	/*
		to note: float64 in a float data type, and the 64 represents the size of the float (64 bits)
		(just in case you don't know, float means a floating-point variable, basically a decimal)
		float64 is not the only way to have a float, we can also have float and float32 in go
		data type of float is just an alias for float64
		float32 on the other hand means a 32 bit floating-point variable
		the only real difference that I know of is that float32 takes up less space
		typically just use float or float64 (as they're the same thing) because they'll get the job done either way
	*/
	someRandomFloat = 6.1234567890

	// something cool you might not know is that you can increment variables easily in go
	// like this
	someRandomInteger++
	// this makes this integer one higher

	someRandomFloat++
	// this works on floats too, and is nearly identical to someRandomFloat = someRandomFloat + 1 with some quirks
	someRandomFloat += float64(someRandomInteger) // this can also work, this is the same as someRandomFloat = someRandomFloat + float64(someRandomInteger)

	fmt.Println("our new float: " + strconv.FormatFloat(someRandomFloat, 'f', -1, 64))
	/*
		notice the usage of strconv again?
		using FormatFloat function we made someRandomFloat into a string variable
		but look at how many arguments we pass into the function, lets look at them one by one:
		first: someRandomFloat passes the floating point variable we want to be parsed into a string
		second: 'f' specifies the format, here the f stands for floating-pont number
		third: -1 means to use the smallest number of digits necessary to accurately represent the float
		fourth: 64 indicates the precision of the float, specifically for float64

		weirdly enough, this doesn't print out 13.123456789 (which is what it is equal to mathematically)
		but rather it prints out 13.123456788999999
		be careful with these weirdness when working with floats always use ranges, as in, if a float is approximately equal to something,
		never exactly.
		this is for the simple reason that they're not too predictable at times (this isn't a go thing, this is a thing for every programming language that I know of)
	*/

	// go has multiple different data types

	// strings are arrays of characters
	// this is a normal string
	var exampleString string = "random string!"
	fmt.Println(exampleString)
	//this can also be written as an array of characters like this:
	var charArray []byte = []byte{'R', 'a', 'n', 'd', 'o', 'm', ' ', 's', 't', 'r', 'i', 'n', 'g', '!'} // byte is basically character, no need to do too much a deep dive
	var stringFromACharArray string = string(charArray[:])
	fmt.Println(stringFromACharArray)

	// you can also get the type of a variable as a string in go like this:
	fmt.Println("type of someRandomInteger is " + reflect.TypeOf(someRandomInteger).String())

	// side note by me: rust compiler is better than Go's tbh, rust is just better

	// in go you can skip a line using \n in a string
	// for example
	for i := 0; i < 5; i++ {
		fmt.Print("\n") // skips 5 lines
	}

	// in every programming language we have while loops.
	// in go its a bit cooler
	// just use a for loop with a condition
	var randomBoolean bool = true
	for randomBoolean {
		break
	}

	// for conditionals in go you can use || and && the same way as in other languages with || representing or and && representing and
	// (just like in English)

	// you can also use for loops ot iterate over arrays
	var stringArray []string = []string{"a", "b", "c"}
	for _, el := range stringArray {
		fmt.Println(el)
	}
	// now you may have noticed the _ variable
	// that's simply a placeholder for the index of the element
	// for example:
	for index, element := range stringArray {
		fmt.Println("index", index, "element", element)
	}

	//

	// FUNCTIONS!!!

	//

	// in go you can call functions that you've created
	var sumOfTwoIntegers int = sum(10, 11)         // 21
	fmt.Println("the sum was: ", sumOfTwoIntegers) // the sum was: 21
	/*
		notice how I didn't convert the sumOfTwoIntegers into a string?
		you don't really have to with fmt.Println, and many other functions in go,
		it does it for you, just put it in using commas like shown above and it'll work
	*/

	// lets sort an array using the sort() method (made by me, scroll down, it's in this file)
	myArrayToSort := []int{543, 52, 15, 54, 45, 13}
	fmt.Println(myArrayToSort)
	mySortingFunction(myArrayToSort) // sorts myArrayToSort in ascending order
	fmt.Println(myArrayToSort)
	// notice how I didn't have to assign the sorted array to my array?
	// as I passed the array into the sort function as an argument, the changes happening to the paramter arr []int also applied to myArrayToSort []int
	// but this might not always be preferred behavior
	// for this you can use the copy function
	myArrayToSort2 := []int{543, 52, 15, 54, 45, 13}
	copiedArray := make([]int, len(myArrayToSort2)) // makes copiedArray to be the same length and type as myArrayToSort2
	copy(copiedArray, myArrayToSort2)               // makes copiedArray to be identical to myArrayToSort2
	mySortingFunction(copiedArray)
	fmt.Println(myArrayToSort2) // this won't be sorted
	// this can be simplified to this
	fmt.Println(mySortingFunction(func() []int {
		a := make([]int, len(myArrayToSort2))
		copy(a, myArrayToSort2)
		return a
	}()))

	/*
		here I used a lambda,
		lambdas in go are functions that can be assigned to variables or just be sent as arguments to a function
		they are also known as anonymous functions or function literals and are not given names
		they are a way to define inline functions, very cool
		this can help lessen the code length, and in some cases make it more readable

		in our example this was the lambda:
		func() []int {
			a := make([]int, len(myArrayToSort2))
			copy(a, myArrayToSort2)
			return a
		} // and here I added "()" to call the lambda to run, just like I would a function
	*/

	// scroll down to see a function called runALambda, it simply takes in a lambda as an argument and runs it
	runALambda(func() {
		fmt.Println("\nthis is printed from a lambda!!\n\n")
	})

	// these can come in handy at different situations

	// the changing of the value of the array only applies to arrays (and a few others) but
	// in go, if you want your variables to change as you send them to a function there are two things you can do
	// you can either make a function that'll change the variable and then return the changed value and then assign it to the original variable
	// or use pointers
	// pointers are a lot simpler than they appear
	// a pointer is a variable that holds the memory address of another variable, these're used to indirectly access and modify the value of a referenced variable
	// in other words, a pointer points at the location of a variable in memory
	// lets take an example
	i := 7
	increment(i)   // this is a function at the bottom of the file, it just adds 1 to i
	fmt.Println(i) // still 7
	/*
		increment(i) doesn't actually change i at all
		the increment function looks like this:
		func increment(i int) {
			i = i + 1
		}

		we could change the increment function to return the changed value instead of nothing and then assign it to i like this:
			i = increment(i)

			func increment(i int) int {
				i = i + 1
				return i
			}
		but this can also be done with pointers

		let's make a function like this (scroll down to see the actual function):
			func increment2(i *int) {
				*i++
			}
		and we can do this:
	*/
	increment2(&i)
	fmt.Println(i) // now this is 8

	// the use of & is to get the memory address of i and send it to the function
	// the use of * in the type is to clarify the pointer type (yes pointers can have their own types)
	// we could have written this like this:
	pointer := &i
	increment2(pointer)
	fmt.Println(i) // now this is 9

	fmt.Print("\n\n\n")

	// constants in Go are similar to other languages
	const ZERO = 0
	// we don't need to assign types nor write := for constants
	/*
		one downside of constants in go is that they often have ot be hard coded
		the value can not determined at runtime or else it won't be a constant and you'll simply have to make a variable instead of a constant
		in other words, you can not create a function that, depending on the values of other variables, will dynamically determine the value of the constant
		"this is not a bug, this is a feature" - some Go dev at Google probably
	*/

	/*
		in go you can make your own types
		for example:
			type person struct {
				name string
				age int
				country string
				married bool
			}
			/// this is actually a piece of code if you scroll up
	*/
	jake := person{
		name:    "Jake Doe",
		age:     69,
		country: "USA",
		married: true,
	}
	// in the example above the type is inferred
	// we can also hard code the type
	var john person = person{
		name:    "John Doe",
		age:     69,
		country: "USA",
		married: true,
	}
	// how do we compare if jake and john are the same person?
	// we can write a function for this
	// and to avoid going outside of the main function I can make a variable and simply assign a lambda to it
	isSamePerson := func(p1 person, p2 person) bool {
		return p1.name == p2.name && p1.age == p2.age && p1.country == p2.country && p1.married == p2.married
	}

	fmt.Println(isSamePerson(jake, john)) // false, not the same person (their names are different)

	// in go you can also run shell commands

	cmd := exec.Command("ls", "-l") // if you're on an unix based system this'll simply print all the files and directories within the directory you are in
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output)) // prints out the output
}

// you can also make functions in go
func sum(x int, y int) int { // takes in two ints, returns an int
	return x + y // will return the sum of the two elements
}

func mySortingFunction(arr []int) []int { // takes in an integer array, returns an integer array (sorted)
	for i := 0; i < len(arr); i++ {
		swapped := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}

func runALambda(lambda func()) {
	lambda()
}

func increment(i int) {
	i = i + 1
}

func increment2(i *int) {
	*i++
}
