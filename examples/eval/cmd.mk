

//
// Monkey language project
// Author: hmza sfyn
// License: MIT
// Repository: https://github.com/hmZa-Sfyn/monkey
// 
// This file is a part of the Monkey language project.
// Learn from this example and contribute to the project!
// This file is one of examples provided in the Monkey language project.
// 
// Jan 2026
// - hmzasfyn
//


if (RUNTIME_OS == "linux") {
	var = "~"
	out = `ls -la $var`
	println(out)
}
elif (RUNTIME_OS == "windows") {
	out = `dir`
	println(out)

	println("")
	println("")
	//test command not exists
	out = `dirs`
	if (!out.ok) {
		printf("Error: %s\n", out)
	}
}



println("\n\n=========================")
println("Double backtick is raw string")

let rawString =``abc
def
lmn
``
println(rawString)

