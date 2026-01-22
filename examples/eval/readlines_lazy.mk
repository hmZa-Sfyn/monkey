

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


// Readlines is an iterator that returns one line of a file at a time. 
fn Readlines(file) {
	ch = chan()
	spawn fn() {
		while ((l = file.readLine()) != nil) {
			ch.send(l)
		}
		ch.close()
	}()

	return ch
}

file = newFile("./file.demo", "r")
if (file == nil) {
	println("opening 'file.demo' for reading failed, error:", file.message())
	os.exit(1)
}


let reader = Readlines(file)
for line in reader {
    fmt.printf("%d\t%s\n", $_ + 1, line) // $_ is the index
}

file.close()