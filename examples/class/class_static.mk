

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


//This file demonstrate static usage

//////////////////////////////////////////////////////////////////////////////////
//                                       First
//////////////////////////////////////////////////////////////////////////////////
//class IndexedNames
//{
//    static let SIZE = 10
//    static fn getSize() {
//        printf("thi.type=%_, size=%d\n", this, this.SIZE)
//    }
//}
//
//fn Main()
//{
//    println("In main, Class Call, IndexedNames.SIZE=", IndexedNames.SIZE)
//    IndexedNames.getSize()
//
////    obj = new IndexedNames()
////    println("In main, Instance Call,IndexedNames.SIZE=", obj.SIZE)
////    obj.getSize()
//
//}
//
//Main()


//////////////////////////////////////////////////////////////////////////////////
//                                       Second
//////////////////////////////////////////////////////////////////////////////////
class Test
{
   static let x = 0;
   static let y = 5;

   static fn Main()
   {
      println(Test.x);
      println(Test.y);

      Test.x = 99;
      println(Test.x);
   }
}

Test.Main()


//////////////////////////////////////////////////////////////////////////////////
//                                       Third
//////////////////////////////////////////////////////////////////////////////////
//class Test
//{
//    static let name;
//    static let age;
//
//    static fn init()
//    {
//        println("Using static constructor to initialize static data members");
//        name = "John Sena";
//        age = 23;
//    }
//    static fn display()
//    {
//
//        println("Using static function");
//        println(name);
//        println(age);
//    }
//
//    static fn Main()
//    {
//        Test.init();
//        Test.display()
//    
//    }
//}
//
//Test.Main()


//////////////////////////////////////////////////////////////////////////////////
//                                       Fourth
//////////////////////////////////////////////////////////////////////////////////
//class MyStaticClass
//{
//    static let myStaticVariable = 0;
//
//    static fn MyStaticMethod()
//    {
//        println("This is a static method.");
//    }
//
//    static property MyStaticProperty { get; set; }
//}
//
//class Program
//{
//
//    static fn Main()
//    {
//        println(MyStaticClass.myStaticVariable);
//
//        MyStaticClass.MyStaticMethod();
//
//        MyStaticClass.MyStaticProperty = 100;
//
//        println(MyStaticClass.MyStaticProperty);
//    }
//
//    fn otherInfo() {
//
//       Program.Main()
//    }
//
//}
//
//p = new Program()
//p.otherInfo()
//Program.Main()