

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


#define DEBUG

//normal statement
println(".........start.........")

// only support two below format:
//    1. #ifdef xxx { body }
//    2. #ifdef xxx { body } #else { body }, here only one '#else' is supported'.
#ifdef DEBUG2
{
    add = fn(x, y) { x + y }
    printf("add = %d\n", add(1, 2))
}
#else
{
    sub = fn(x, y) { x - y }
    printf("sub = %d\n", sub(3, 1))
}

#define TESTING
#ifdef TESTING
{
    add = fn(x, y) { x + y }
    printf("add = %d\n", add(1, 2))
}

//normal statement
println(".........end...........")