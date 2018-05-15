// Copyright 2014 The genny Authors. All rights reserved.
// Use of this source code is governed by the terms of the MIT license that can
// be found in the LICENSE file.

/*
genny - Generics for Go, a command line tool meant for go generate.
(pron. Jenny)

Made by Mat Ryer (Twitter: @matryer) and Tyler Bunnell (Twitter: @TylerJBunnell).

Until the Go core team include support for generics in Go
http://golang.org/doc/faq#generics,
genny is a code-generation generics solution. It allows you write normal
buildable and testable Go code which, when processed by the 'genny gen' tool,
will replace the generics with specific types.

  - Generic code is valid Go code
  - Generic code compiles and can be tested
  - Use 'stdin' and 'stdout' or specify in and out files
  - Supports Go 1.4's go generate (http://tip.golang.org/doc/go1.4#gogenerate)
  - Multiple specific types will generate every permutation
  - Use 'BUILTINS' and 'NUMBERS' wildtype to generate specific code for all
    built-in (and number) Go types
  - Function names and comments also get updated

Library

We have started building a library of common things
https://github.com/cheekybits/gennylib,
and you can use 'genny get' to generate the specific versions you need.
For example:

	genny get maps/concurrentmap.go "KeyType=BUILTINS ValueType=BUILTINS"

will print out generated code for all types for a concurrent map. Any file in
the library may be generated locally in this way using all the same options
given to 'genny gen'.

Usage

The genny command line usage is documented here.

	genny [{flags}] gen "{types}"

	gen - generates type specific code from generic code.
	get <package/file> - fetch a generic template from the online library and gen it.

	{types}  - (optional) Command line flags (see below)
	{types}  - (required) Specific types for each generic type in the source
	{types} format:  {generic}={specific}[,another][ {generic2}={specific2}]

	Examples:

	Generic=Specific
	Generic1=Specific1 Generic2=Specific2
	Generic1=Specific1,Specific2 Generic2=Specific3,Specific4

	Flags:

	-in="": file to parse instead of stdin
	-out="": file to save output to instead of stdout
	-pkg="": package name for generated files

	* Comma separated type lists will generate code for each type

Flags

	-in  - specify the input file  (rather than using stdin)
	-out - specify the output file (rather than using stdout)

go generate

To use Go 1.4's 'go generate' capability, insert the following comment in your
source code file:

	//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "KeyType=string,int ValueType=string,int"

	- Start the line with '//go:generate '
	- Use the '-in' and '-out' flags to specify the files to work on
	- Use the 'genny' command as usual after the flags

Now, running 'go generate' (in a shell) for the package will cause the generic
versions of the files to be generated.

  - The output file will be overwritten, so it's safe to call 'go generate' many times
  - Use '$GOFILE' to refer to the current file
  - The '//go:generate' line will be removed from the output

How it works

Define your generic types using the special 'generic.Type' placeholder type:

	type KeyType generic.Type
	type ValueType generic.Type

	- You can use as many as you like
	- Give them meaningful names

Then write the generic code referencing the types as your normally would:

	func SetValueTypeForKeyType(key KeyType, value ValueType) { ... }

	- Generic type names will also be replaced in comments and function names (see Real example below)

Since 'generic.Type' is a real Go type, your code will compile, and you can
even write unit tests against your generic code.

Generating specific versions

Pass the file through the 'genny gen' tool with the specific types as the
argument:

	cat generic.go | genny gen "KeyType=string ValueType=interface{}"

The output will be the complete Go source file with the generic types replaced
with the types specified in the arguments.

Additional Examples

You can find examples in the example*_test.go files, as well as under the
'examples' directory and in the README file. Please also look in the generic
and parse packages.

Contributions

  - See the API documentation for the parse package (http://godoc.org/github.com/cheekybits/genny/parse)
  - Please do TDD
  - All input welcome
*/
package main
