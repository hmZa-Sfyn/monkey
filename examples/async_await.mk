

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


//Three kind of async
//    Function Statement:              async fn add(x, y) { x + y }
//    Function Literal:      let add = async fn (x,y)     { x + y }
//    Lambda:                let add = async (x, y) =>    { x + y } 


//==========================
//async function statement
//==========================
async fn add(a, b) { a + b }

result = await add(1, 2)
println(result)


//==========================
//async function literal
//==========================
let addFn = async fn(a, b) { a + b }

result = await addFn(3, 4)
println(result)


//==========================
//async lambda expression
//==========================
let lambdaFn = async (a, b) => { a + b }

result = await lambdaFn(5, 6)
println(result)