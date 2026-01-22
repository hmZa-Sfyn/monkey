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

class @Test {
  property Enabled
}

//marker annotation
class @Demo {}


class TestExample {
  @Demo
  @Test{Enabled = true}
  fn TestA() {
    printf("TestA is called\n")
  }

  @Demo
  @Test{Enabled = false}
  fn TestB() {
    printf("TestB is called\n")
  }

  @Demo
  @Test{Enabled = false}
  fn TestC() {
    printf("TestC is called\n")
  }

  @Demo
  @Test{Enabled = true}
  fn TestD() {
    printf("TestD is called\n")
  }
}

testObj = new TestExample()
for method in testObj.getMethods() {
  printf("\nmethodName=%s\n", method.name)
  annos = method.getAnnotations()
  for anno in annos {
    //println()
    //printf("ANNO NAME=%s, enabled=%t\n", anno, anno.Enabled)

    if anno.instanceOf(Test) {
      printf("ANNO NAME=%s, enabled=%t\n", anno, anno.Enabled)
      if anno.Enabled {
        method.invoke()
      }
    } elif anno.instanceOf(Demo) {
      printf("ANNO NAME=%s, \n", anno)
    }
  }
}
