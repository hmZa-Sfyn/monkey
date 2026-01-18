/*
 * lib: monkey/bridge/verylong.mk
 * author: hmza-sfyn
 * github: https://github.com/hmza-sfyn/monkey
 * license: MIT
 *
 * This is a monkey library for handling very long integers.
 * It uses python3's built-in arbitrary-precision integer capabilities.
 * It is intended to be used as a bridge between monkey and python3 for very long integer operations.
 *
 */

class pyint {
    let pynum = "0";
    fn init(num="0") {
        if (`python3 -c "print(int($num))"`) == str(num) {
            this.pynum = num
        }
        else {
            println("Invalid integer format: " + num)
        }
    }

    fn +(other) {
        if type(other) == type(new pyint(2)) {
            th = this.pynum
            ot = other.pynum
            return new pyint(`python3 -c "print($th + $ot)"`)
        }
        else {
            th = this.pynum
            ot = other
            return new pyint(`python3 -c "print($th + $ot)"`)     
        }
    }

    fn -(other) {
        if other.is_a(pyint) {
            th = this.pynum
            ot = other.pynum
            return new pyint(`python3 -c "print($th - $ot)"`)
        }
        else {
            th = this.pynum
            ot = other
            return new pyint(`python3 -c "print($th - $ot)"`)     
        }
    }
    
    fn *(other) {
        if other.is_a(pyint) {
            th = this.pynum
            ot = other.pynum
            return new pyint(`python3 -c "print($th * $ot)"`)
        }
        else {
            th = this.pynum
            ot = other
            return new pyint(`python3 -c "print($th * $ot)"`)     
        }
    }

    fn /(other) {
        if other.is_a(pyint) {
            th = this.pynum
            ot = other.pynum
            return new pyint(`python3 -c "print($th // $ot)"`)
        }
        else {
            th = this.pynum
            ot = other
            return new pyint(`python3 -c "print($th // $ot)"`)     
        }
    }

    fn ++() {
        th = this.pynum
        return new pyint(`python3 -c "print($th + 1)"`)
    }

    fn --() {
        th = this.pynum
        return new pyint(`python3 -c "print($th - 1)"`)
    }

    fn %(other) {
        if other.is_a(pyint) {
            th = this.pynum
            ot = other.pynum
            return new pyint(`python3 -c "print($th % $ot)"`)
        }
        else {
            th = this.pynum
            ot = other
            return new pyint(`python3 -c "print($th % $ot)"`)     
        }
    }   
}

fn Max(a,b) {
    anum = a.pynum
    bnum = b.pynum
    if anum == `python3 -c "print(max($anum, $bnum))"` {
        return a
    }
    else {
        return b
    }
}

fn Min(a,b) {
    anum = a.pynum
    bnum = b.pynum
    if anum == `python3 -c "print(min($anum, $bnum))"` {
        return a
    }
    else {
        return b
    }
}

fn Abs(a) {
    anum = a.pynum
    if anum == `python3 -c "print(abs($anum))"` {
        return a
    }
    else {
        return new pyint(`python3 -c "print(-$anum)"`)
    }
}

fn Sqrt(a) {
    anum = a.pynum
    if anum == `python3 -c "print($anum ** 0.5)"` {
        return a
    }
    else {
        return new pyint(`python3 -c "print(int($anum ** 0.5))"`)
    }
}   

fn Pow(a,b) {
    anum = a.pynum
    bnum = b.pynum
    if anum == `python3 -c "print($anum ** $bnum)"` {
        return a
    }
    else {
        return new pyint(`python3 -c "print($anum ** $bnum)"`)
    }
}

fn Mod(a,b) {
    anum = a.pynum
    bnum = b.pynum
    if anum == `python3 -c "print($anum % $bnum)"` {
        return a
    }
    else {
        return new pyint(`python3 -c "print($anum % $bnum)"`)
    }
}


fn Pyint(ops) {
    return new pyint(`python3 -c "print($ops)"`)
}

fn Demo_DO_NOT_USE() {
    // VeryLongInteger demo using ALL functions from the library
    // January 2026 - let's play with big numbers! 🚀

    println("=== Very Long Integer Demo using Python bridge ===\n")

    // Different ways to create pyint objects
    let a = new pyint("123456789012345678901234567890")          // constructor
    let b = Pyint("987654321098765432109876543210")             // helper function
    let c = Pyint("1000000000000000000000000000000")           // 10³⁰
    let d = new pyint("3141592653589793238462643383279")       // pi × 10²⁷ approx

    println("a =", a.pynum)
    println("b =", b.pynum)
    println("c =", c.pynum)
    println("d ≈ π×10²⁷ =", d.pynum)
    println("")

    // ───────────────────────────────────────────────
    // All arithmetic operations
    // ───────────────────────────────────────────────
    println("Basic operations:")
    let sum     = a + b
    let diff    = b - a
    let product = a * c
    let div     = product / b          // integer division
    let mod     = product % b

    println("  a + b          =", sum.pynum)
    println("  b - a          =", diff.pynum)
    println("  a × c          =", product.pynum)
    println("  (a×c) // b     =", div.pynum)
    println("  (a×c) % b      =", mod.pynum)
    println("")

    // ───────────────────────────────────────────────
    // Increment / Decrement
    // ───────────────────────────────────────────────
    println("In/De-crement:")
    let x = new pyint("999999999999999999")
    println("  x             =", x.pynum)

    x = x++     // same as x = x + 1
    println("  x++           =", x.pynum)

    x = x--     // same as x = x - 1
    println("  x--           =", x.pynum)
    println("")

    // ───────────────────────────────────────────────
    // Comparison helpers: Max / Min
    // ───────────────────────────────────────────────
    println("Comparisons (Max/Min):")
    println("  Max(a, b)      =", Max(a, b).pynum)
    println("  Max(c, d)      =", Max(c, d).pynum)
    println("  Min(a, c)      =", Min(a, c).pynum)
    println("  Min(b, d)      =", Min(b, d).pynum)
    println("")

    // ───────────────────────────────────────────────
    // Other math functions
    // ───────────────────────────────────────────────
    println("Other math functions:")
    let neg_d   = d * new pyint("-1")
    let abs_neg = Abs(neg_d)
    let sqrt_c  = Sqrt(c)
    let pow_2_20 = Pow(new pyint("2"), new pyint("20"))
    let pow_3_12 = Pow(Pyint("3"), Pyint("12"))

    println("  d             =", d.pynum)
    println("  -d            =", neg_d.pynum)
    println("  abs(-d)       =", abs_neg.pynum)
    println("  sqrt(10³⁰)    ≈", sqrt_c.pynum, "  (≈ 10¹⁵)")
    println("  2²⁰           =", pow_2_20.pynum)
    println("  3¹²           =", pow_3_12.pynum)
    println("")

    println("=== Demo finished - all functions used! ===")
    println("Big integer math in Monkey via Python — pretty cool combination!  (•̀ᴗ•́)و ̑̑")
}