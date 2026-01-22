

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


if ().empty() { //() is a tuple
    println("TRUE")
} else {
    println("FALSE")
}

// (1) is 1, not a tuple. you must use (1,) if you really want a tuple
//if (1).empty() {
//    println("TRUE")
//} else {
//    println("FALSE")
//}

let x,y,z,w = (),[1,2,3], 4, (1+2) * 3
println("x=", x)
println("y=", y)
println("z=", z)
println("w=", w)

while () {  //'()' is a tuple
  println("empty")
}
println("END")

let m,n = () => 10, {v:k for k,v in {"key1":"value1", "key2":"value2"}}
println(m())
println(n)
println(m)