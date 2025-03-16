package anotherpackage

import "fmt"

// if var name begins with capital letter - it is public and available in other packages
var MyExportedVar = "This var is in anotherpackage and available for export"

var myPrivateVar = "This var name starts with small letter and scoped only to this package"

//the same logic applies to functions

func MyExportedPrintFunc(var1, var2 string) {
	fmt.Println("This is printed from anotherpackage function")
	fmt.Println("First var is passed as parameter to the function: ", var1)
	fmt.Println("Second var is passed as parameter to the function: ", var2)
	fmt.Print("And this is package scoped var from anotherpackage, available for export: ", MyExportedVar)
	fmt.Print("And this is package scoped var from anotherpackage, NOT available for export: ", myPrivateVar)

}

func myPrivateFunc() {

	fmt.Println("This func is private in anotherpackage and cannot be used in another package")
}
