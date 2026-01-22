

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


//This file demonstrate extension methods of class

class Animal {
	fn Walk() {
		println("Animal Walk!")
	}
}

//extension methods like objective-c
class Animal (Run) {
	fn Run() {
		println("Animal Run!")
		this.Walk() //call Walk() method of Animal class.
	}
}

animal = new Animal()
animal.Walk()

println()
animal.Run()