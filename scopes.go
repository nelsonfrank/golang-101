package main

import "fmt"

// SCOPE
// Scope of the variable
// is the way you can access and use that variable
//
// Level of Scope
// i) Block - variable declared inside block(such as function) can be accessed
//            inside that block only and not outside.
// ii) Package - variable declared in package level (i.e outside the function/blocks)
//              can be accessed throughout package(i.e it can be access to all file
//              under that package).
// iii) Universe - variable declared into different package can be access into another
//                package's file by importing that package and exporting that variable
//                variable can be declared as exported by using capitalize variable name
//                example: var Students int = 20,
//                NB: variable should be declared at package level
//               import custom package path structure - "host/user/project/(dir)/package"
//
// Structure of golang project
//   - Separate packages into different folder with name of the package declared in it's files
//     i.e Structure
//      - root
//           - main
//                - (all files under this package)
//            - not-main
//                - (all files under this package)
	
func scopes() {

    fmt.Println("Hello World")

}