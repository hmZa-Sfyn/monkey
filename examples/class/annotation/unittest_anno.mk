

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


//Annotations
class @Test {}
class @TestSetup {}
class @TestTearDown {}

//Calculator class
class Calculator{
  @TestSetup
  static fn Setup() {
    println("TestSetup called!\n")
  }

  @TestTearDown
  static fn TearDown() {
    println("\nTestTearDown called!\n")
  }

  @Test
  fn add(x, y) {
    return x + y
  }

  @Test
  fn sub(x, y) {
    return x - y
  }

  @Test
  fn mul(x, y) {
    return x * y
  }

  @Test
  fn div(x, y) {
    return x / y
  }
}

class Handler {
  static fn handle(o) {
    methods = o.getMethods()

    //run test setup
    for m in methods {
      testSetupAnno = m.getAnnotation(TestSetup)
      if testSetupAnno != nil {
        m.invoke()
        break
      }
    }

    //run test
    for m in methods {
      testAnno = m.getAnnotation(Test)
      if testAnno != nil {
        printf("%s(20, 10) = %v\n", m.name, m.invoke(20, 10))
      }
    }

    //run test teardown
    for m in methods {
      testTearDownAnno = m.getAnnotation(TestTearDown)
      if testTearDownAnno != nil {
        m.invoke()
        break
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