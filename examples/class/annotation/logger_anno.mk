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


//This file demonstrate annotation
class @MethodTrace { //Annotation class
  property LogLevel
}

enum Logger {
  INFO,
  DEBUG
}

class Calculator {
  @MethodTrace{LogLevel = Logger.INFO}
  fn calculatePow(x, y) {
    result = math.pow(x, y)
    return result
  }
  
  @MethodTrace{LogLevel = Logger.DEBUG}
  fn multiply(a, b) {
    return a * b
  }
  
  @MethodTrace{LogLevel = Logger.DEBUG}
  fn add(a, b) {
    return a + b
  }
}

class Handler {
  static fn handle(o) {
    for m in o.getMethods() {
      for a in m.getAnnotations() {
        if (a.is_a(MethodTrace)) {
          if a.LogLevel == Logger.INFO {
            println("Performing INFO log for \"" + m.getName() + "\" method: ")
            printf("calculation result = %v\n\n", m.invoke(2,3))
          } elif (a.LogLevel == Logger.DEBUG) {
            println("Performing DEBUG log for \"" + m.getName() + "\" method")
            printf("calculation result = %d\n\n", m.invoke(10,5))
          }
        }
      }
    }
  }
}

class Main {
  static fn main() {
    Calculator calculator = new Calculator()
    Handler.handle(calculator)
  }
}

Main.main()