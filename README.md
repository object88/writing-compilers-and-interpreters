# Writing Compilers And Interpreters

This project is an attempt to work through [Ronald Mak](http://www.cs.sjsu.edu/~mak/)'s book, ["Writing Compilers and Interpreters" (3rd ed., 2009)](https://www.wiley.com/en-us/Writing+Compilers+and+Interpreters%3A+A+Software+Engineering+Approach%2C+3rd+Edition-p-9780470177075), using the Go programming language.

Like the book's original content, this project is designed to compile and interpret the Pascal programming language.  Or, more pointedly, to learn the mechanics behind compiling and interpreting a programming language.

Certain changes have been made, which should not have any impact on the content of the work.  In particular, Steve Francia's [cobra package](https://github.com/spf13/cobra) is used to handle command line parsing.

Other changes are natural occurances due to the differences between Go and Java.  Go has no exceptions, for example, while Java depends on them extensively.  The Go code in this project depends heavily on returned `error` instances, which changes the signature of functions.  We have to be mindful of how this effects the logical flow of the functions themselves.
