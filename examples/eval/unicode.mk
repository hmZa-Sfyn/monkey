

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


// constant with mixed type runes
mixed = "\b5Ὂg̀9! ℃ᾭG"
for c in mixed {
    fmt.printf("For %q:\n", c)
    if unicode.isControl(c) {
        fmt.println("\tis control rune")
    }
    if unicode.isDigit(c) {
        fmt.println("\tis digit rune")
    }
    if unicode.isGraphic(c) {
        fmt.println("\tis graphic rune")
    }
    if unicode.isLetter(c) {
        fmt.println("\tis letter rune")
    }
    if unicode.isLower(c) {
        fmt.println("\tis lower case rune")
    }
    if unicode.isMark(c) {
        fmt.println("\tis mark rune")
    }
    if unicode.isNumber(c) {
        fmt.println("\tis number rune")
    }
    if unicode.isPrint(c) {
        fmt.println("\tis printable rune")
    }
    if !unicode.isPrint(c) {
        fmt.println("\tis not printable rune")
    }
    if unicode.isPunct(c) {
        fmt.println("\tis punct rune")
    }
    if unicode.isSpace(c) {
        fmt.println("\tis space rune")
    }
    if unicode.isSymbol(c) {
        fmt.println("\tis symbol rune")
    }
    if unicode.isTitle(c) {
        fmt.println("\tis title case rune")
    }
    if unicode.isUpper(c) {
        fmt.println("\tis upper case rune")
    }
}