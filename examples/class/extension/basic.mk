

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


//This file demonstrate extension methods of class

class Module {
	fn __() {
		println("<<Wrapper for Parent-Class Module>>")
	}

    fn Pwd() {
        return `pwd`
    }
}


//extension methods like objective-c
class Module (Config) {
    let _debug_user = "root"
    let _debug_pass = "toor"
	fn Config(data) {
        self = {}
		self._debug = false

        self._debug_user = "" 
        self._debug_pass = ""

        debugStat = data["debug"]["true"]
        if debugStat != false {
            self._debug = true
            self._debug_user = data["debug"]["uname"]
            self._debug_pass = data["debug"]["pass"]
        }
        _debug = self._debug
        _debug_user = self._debug_user
        _debug_pass = self._debug_pass

        self.pwd = this.Pwd() // calling function defined in Module class
        self.mode = () => { 
            if _debug {
                "dev:debug"
            } else {
                "prod"
            }
        }

        return self
	}
}

moduleX = new Module()

// call extension
x = moduleX.Config({"debug":{"true":"true","uname":"root","pass":"toor"}})

println(moduleX.Pwd())
println(x.mode())