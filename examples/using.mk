

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


// Here we use 'using' statement, so we do not need to call infile.close().
// When finished running 'using' statement, it will automatically call infile.close().
using (infile = newFile("./file.demo", "r")) {
	if (infile == nil) {
		println("opening 'file.demo' for reading failed, error:", infile.message())
		os.exit(1)
	}

	let line;
	let num = 0
	//Read file by using extraction operator(">>")
	while (infile>>line != nil) {
		num++
		printf("%d	%s\n", num, line)
	}
}
