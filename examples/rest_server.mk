

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


/*
    Rest Service Example
    For more advanced example, please see 'service.mk'
*/

let login = fn(writer, request) {
    if request.method == http.METHOD_POST {
        writer.header.add("Content-Type", "application/json; charset=utf-8")
        writer.write(json.toJson({ "sessionId": "3d5bd2cA15ef0476" }))
    }
}

let logout = fn(writer, request) {
    if request.method == http.METHOD_POST {
        writer.writeHeader(http.STATUS_CREATED) // return http status code 201
    }
}

http.handleFunc("/authentication/login", login)
http.handleFunc("/authentication/logout", logout)

addr = "0.0.0.0:8080"
printf("* Running on %s\n", addr)
http.listenAndServe(addr)
