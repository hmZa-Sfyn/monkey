

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


let handler = fn(writer, request) {
	writer.write("Hello HTTP client, I'm server")
}

http.listenAndServe("127.0.0.1:8080", handler)
//Same as above
//http.listenAndServe("127.0.0.1:8080", fn(writer, request) {
//	writer.write("Hello HTTP client, I'm server")
//})