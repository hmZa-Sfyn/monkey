

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


//#!~/.mk/bin/magpie
const a = 3.14159
//a = 3.14 //error
printf("PI=%f\n", a)

fn Test() {
  let a = 15 //good
  printf("In Test(), a=%d\n", a)
}

Test()

const b = fn(x) {
  println(x)
}

//b = a //error
b("Hello")