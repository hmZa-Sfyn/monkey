

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


fn process(tasks) {
  ch = chan()
  spawn fn() {
    for idx, task in tasks {
        ch.send(task)
    }
    ch.close()
  }()
  return ch
}

tasks = ["foo", "bar", "baz", "hhf", "hht", "hy", "hjq1234567890"]
results = process(tasks)

for result in results {
  println(result)
}

// XRange is an iterator over all the numbers from 0 to the limit.
fn XRange(limit) {
	ch = chan()
	spawn fn() {
		for i in 0..limit {
			ch.send(i)
		}

		// Ensure that at the end of the loop we close the channel!
		ch.close()
	}()
	return ch
}

for i in XRange(10) {
    fmt.println(i)
}
