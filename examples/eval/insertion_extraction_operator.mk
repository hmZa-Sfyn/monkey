

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


// Output to stdout by using insertion operator("<<")
// 'endl' is a predefined object, which is "\n".
stdout << "hello " << "world!" << " How are you?" << endl;




// Read from stdin by using extraction operator(">>")
let name;
stdout << "Your name please: ";
stdin >> name;
printf("Welcome, name=%v\n", name)




//Read file by using extraction operator(">>")
infile = newFile("./file.demo", "r")
if (infile == nil) {
	println("opening 'file.demo' for reading failed, error:", infile.message())
	os.exit(1)
}

let line;
let num = 0
while ( infile>>line != nil) {
	num++
	printf("%d	%s\n", num, line)
}
infile.close()




//Writing to file by using inserttion operator("<<")
outfile = newFile("./outfile.demo", "w")
if (outfile == nil) {
	println("opening 'outfile.demo' for writing failed, error:", outfile.message())
	os.exit(1)
}

outfile << "Hello" << endl
outfile << "world" << endl
outfile.close()