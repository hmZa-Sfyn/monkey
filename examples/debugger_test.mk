

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


let a = 10

fn printMe(msg) {
    msg = "** " + msg + " **"
    println(msg)
    demo(8)
}

let demo = fn(a) {
    printf("a=%d\n", a)
    if (a > 5) {
        println("a > 5")
        return
    } else {
        println("a<=5")
    }
}

for i in 1..5 {
    printf("i=%d\n", i)
}

demo(4)

if (a > 12) {
    printMe("a>12")
} else {
    printMe("a < 12")
}

println(len("Hello World"))

let x = 5 let y = "good"
let add = fn(x, y) {
    return x + y
}

sum = add(10, 5)
printf("x=%d, len(%s)=%d, 10+5=%d\n", x, y, strings.len(y), sum)

println("Program end.")
