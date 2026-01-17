

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


gfmt.Printf("Hello %s!\n", "go function");
gfmt.Printf("%v\n", gos.Args)

let f, err = gos.Open("./mdoc.go")
if (err != nil) { // or 'if err {' 
    gfmt.Println(f, err)
} else {
    gfmt.Println(f, err)
}

let name,_ = gos.Hostname
gfmt.Println(name)

i = math_rand.Int
gfmt.Println(i)

let files, err = io_ioutil.ReadDir(".")
if err != nil {
	glog.Fatal(err)
}

for file in files {
	if file.Name() == ".git" {
		continue
	}
	gfmt.Printf("Name=%s, Size=%d\n", file.Name(), file.Size())
}

for s in gregex.MustCompile(``[\s_]``).Split("foo_bar_baz", -1) {
  println(s)
}

println(runtime.GOOS)
println(runtime.GOARCH)