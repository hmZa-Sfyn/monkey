

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


//import "examples/sub_package/calc"
import "sub_package/calc"

a, b = 10, 2
printf("10 + 2 = %d\n", calc.Add(a, b))
printf("10 - 2 = %d\n", calc.Minus(a, b))

printf("2 * 2 = %d\n", calc.SQUARE_OF_TWO)

c, d = calc.Add_and_Minus(a, b)
printf("c = %d, d = %d\n", c, d)

e, f = calc.Add_and_Minus(c, d)
printf("e = %d, f = %d\n", e, f)

//multiply is not exported, so below line will report runtime error.
//printf("10*2=%d\n", calc.multiply(a, b))

//will report error: 'unexported name'
//printf("4*4 = %d\n", calc.sQUARE_OF_FOUR)

//will report error: 'undefined method'
//printf("10 + 2 = %d\n", calc.Adds(a, b))

//will report error: 'unexported name'
//printf("10 + 2 = %d\n", calc._add(a, b))

//import std.io // => ~/.mk/monkey/lib/std/io.mk
//import !.httpclient // => ./httpclient.mk

