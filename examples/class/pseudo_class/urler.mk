

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


// A pseudo-class using function with closure
// scene: A web target accessor with closure functions too!
fn Target(url) {
    
    //init the self hash!
    self = {}
    self.url = url
    self.http = fn() { 
        strings.split(self.url, "//")[0]
        .replace(":","")
    }

    //basic func example
    self.address = fn() {
        webaddr = strings.split(self.url, "//")[1]
        webaddrs = strings.split(webaddr, ".")
    }

    //how to rev a array? trust me, I myself did'nt knew untill now!
    self._helper_rev_addr = fn() {
        addrs = self.address()
        newaddrs = []
        for x in len(addrs)-1..0 {
            newaddrs += (addrs[x])
        }

        newaddrs
    }
    self.webaddr_suffix = fn() {
        rev_addr = self._helper_rev_addr()

        rev_addr[0]
    }
    return self
}

t = Target("https://www.google.com")
println(t.http())
println(t.address())
println(t.webaddr_suffix())
