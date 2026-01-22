

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


let conn = dialTCP("tcp", "127.0.0.1:9090")
if (conn == nil) {
    println("dailTCP failed, error:", conn.message())
    os.exit(1)
}

let n = conn.write("Hello server, I'm client")
if (n == nil) {
    println("conn write failed, error:", n.message())
    os.exit(1)
}

let ret = conn.close()
if (ret == false) {
    println("Server close failed, error:", ret.message())
}