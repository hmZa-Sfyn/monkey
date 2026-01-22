

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



class @ParentMinMaxValidator {
  property MaxLength default 10
}

class @MinMaxValidator : ParentMinMaxValidator {
  property MinLength
  property MaxLength default 10
}

class @NoSpaceValidator {}

class @DepartmentValidator {
  property Department
}

class Request {
  @MinMaxValidator(MinLength=1, MaxLength=10)
  @MinMaxValidator(MinLength=1)
  property FirstName;

  @NoSpaceValidator
  property LastName;

  @DepartmentValidator(Department=["Department of Education", "Department of Labors", "Department of Justice"])
  property Dept;
}

class RequestHandler {
  static fn handle(o) {
    props = o.getProperties()
    for p in props {
      annos = p.getAnnotations()
      for anno in annos {
        if anno.instanceOf(MinMaxValidator) {
          if len(p.value) > anno.MaxLength || len(p.value) < anno.MinLength {
            printf("Property '%s' is not valid!\n", p.name)
          }
        } elif anno.instanceOf(NoSpaceValidator) {
          for c in p.value {
            if c == " " || c == "\t" {
              printf("Property '%s' is not valid!\n", p.name)
              break
            }
          }
        } elif anno.instanceOf(DepartmentValidator) {
          found = false
          for d in anno.Department {
            if p.value == d {
              found = true
            }
          }
          if !found {
            printf("Property '%s' is not valid!\n", p.name)
          }
        }
      }
    }
  }
}

class RequestMain {
  static fn main() {
    request = new Request()
    //request.FirstName = "Haifeng123456789" not valid
    //request.LastName = "Huang    "  not valid
    request.Dept = "Department of Justice"
    RequestHandler.handle(request)
  }
}

RequestMain.main()