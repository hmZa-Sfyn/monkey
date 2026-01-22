<p align="center">
    <img alt="Monkey language logo" src="./logo.png" width="310">
</p>

# Monkey Programming Language

## Overview

Monkey is a toy programming language and interpreter written in **Go**, originaly based on heifenghuang's magpie(credits at the end of this file).  
- It features a clean, C-style syntax and is heavily inspired by **Ruby**, **Python**, **Perl**, **C#**, and **JavaScript**.
- Revolves around `oop` features, and monkey is `not a functional language`.

Although designed primarily for learning and experimentation, Monkey provides a rich and modern feature set, along with built-in tooling that makes it feel like a production-ready language.

## Official website
### [monkeylang.vercel.app](https://monkeylang.vercel.app/)

## Summary

- Implemented as an interpreter in **Go**
- C-style syntax with influences from multiple modern languages
- Supports imperative, functional, and object-oriented programming paradigms
- Can import and interoperate with Go modules
- Includes a built-in HTML documentation generator (**mdoc**)
- Provides a simple but effective debugger
- Comes with an interactive REPL featuring real-time syntax highlighting
- Has been used to implement another programming language written entirely in Monkey
- Most Monkey scripts can run in a web browser due to JavaScript-like semantics

## Documentation

A complete language tutorial and reference can be found in the  
[`docs`](docs) directory.

## Features

### Language Features

- Class system with:
  - Properties
  - Indexers
  - Operator overloading
- First-class functions
- Functions with:
  - Variadic parameters
  - Default values
  - Multiple return values
- Built-in asynchronous programming with `async` / `await`
- Exception handling with `try` / `catch` / `finally`
- Optional type support (similar to Java 8)
- Simple macro processing
- Elixir-style pipe operator (`|>`)

### Data Types

- `int`, `uint`, `float`, `bool`
- `array`, `tuple`, `hash`
- All data types:
  - Support JSON marshal & unmarshal
  - Are extensible

### Built-in Capabilities

- LINQ-style query support
- Native datetime literals
- `using` statement (C#-like resource management)
- Integrated service processing
- Built-in macro system

### Tooling

- Syntax-highlighted REPL
- Simple built-in debugger
- Go package integration via:
  - `RegisterFunctions`
  - `RegisterVars`

## Purpose

### Monkey is ideal for:
-

###More
- Learning how interpreters and languages are built
- Experimenting with language features and design ideas
- Exploring how tooling (REPLs, debuggers, doc generators) can be integrated into a language ecosystem

## Example1(Linq)

```csharp
// async/await
async fn add(a, b) { a + b }

result = await add(1, 2)
println(result)

// linq example
class Linq {
    static fn TestSimpleLinq() {
        //Prepare Data Source
        let ingredients = [
            {Name: "Sugar",  Calories: 500},
            {Name: "Egg",    Calories: 100},
            {Name: "Milk",   Calories: 150},
            {Name: "Flour",  Calories: 50},
            {Name: "Butter", Calories: 200},
        ]

        //Query Data Source
        ingredient = from i in ingredients where i.Calories >= 150 orderby i.Name select i

        //Display
        for item in ingredient => println(item)
    }

    static fn TestFileLinq() {
        //Read Data Source from file.
        file = newFile("./examples/linqSample.csv", "r")

        //Query Data Source
        result = from field in file where int(field[1]) > 300000 select field[0]

        //Display
        for item in result => printf("item = %s\n", item)

        file.close()
    }

    /* Code from https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/let-clause */
    static fn TestComplexLinq() {
        //Data Source
        stringList = [
            "A penny saved is a penny earned.",
            "The early bird catches the worm.",
            "The pen is mightier than the sword."
        ]

        //Query Data Source
        earlyBirdQuery =
            from sentence in stringList
            let words = sentence.split(" ")
            from word in words
            let w = word.lower()
            where w[0] == "a" || w[0] == "e" ||
                  w[0] == "i" || w[0] == "o" ||
                  w[0] == "u"
            select word

        //Display
        for v in earlyBirdQuery => printf("'%s' starts with a vowel\n", v)
    }
}

Linq.TestSimpleLinq()
println("======================================")
Linq.TestFileLinq()
println("======================================")
Linq.TestComplexLinq()
```

## Example2(Rest Service)

```csharp
//service Hello on "0.0.0.0:8090" {
service Hello on "0.0.0.0:8090:debug" { //':debug': for debugging request
  //In '@route', you could use 'url(must), methods, host, schemes, headers, queries'
  @route(url="/authentication/login", methods=["POST"])
  fn login(writer, request) {
    //writer.writeJson({ sessionId: "3d5bd2cA15ef047689" })
    //writer.writeJson({ sessionId: "3d5bd2cA15ef047689" }), 200 # same as above
    //return { sessionId: "3d5bd2cA15ef047689" }, 200 # same as above
    return { sessionId: "3d5bd2cA15ef047689" } // same as above
  }

  @route(url="/authentication/logout", methods=["POST"])
  fn logout(writer, request) {
    // writer.writeHeader(http.STATUS_CREATED) # return http status code 201
    return http.STATUS_CREATED // same as above
  }

  @route(url="/meters/setting-result/{acceptNo}", methods=["GET"])
  fn load_survey_result(writer, request) {
    //using 'vars' dictionary to access the url parameters
    //writer.writeJson({ acceptNo: vars["acceptNo"], resultCode: "1"})
    return { acceptNo: vars["acceptNo"], resultCode: "1"} // same as above
  }

  @route(url="/articles/{category}/{id:[0-9]+}", methods=["GET"])
  fn getArticle(writer, request) {
    //using 'vars' dictionary to access the url parameters
    //writer.writeJson({ category: vars["category"], id: vars["id"]})
    return { category: vars["category"], id: vars["id"]} // same as above
  }
}
```

## Getting started

### Learn my examples:
[Examples.md](./docs/EXAMPLES.md)

### Playgroung using wasm:


## Contributing

Contributions are highly welcome! 🎉  
If you make any improvements or changes to the language, please let me know so I can acknowledge your work in the **Credits** section.

Bug fixes, feature additions, documentation improvements, and suggestions are all appreciated.

## Credits

- **haifenghuang**  
  This repository was forked from his original project  
  [magpie](https://github.com/haifenghuang/magpie)

- **mayoms**  
  This project is based on mayoms’s  
  [monkey](https://github.com/mayoms/monkey) interpreter

- **ahmetb**  
  The LINQ module is based on  
  [go-linq](https://github.com/ahmetb/go-linq)

- **shopspring**  
  The decimal module is based on  
  [decimal](https://github.com/shopspring/decimal)

- **gorilla**  
  The service module is based on  
  [mux](https://github.com/gorilla/mux)

## Installation

1. Download or clone the repository.
2. Run the following command:

```bash
./setup.sh
```

## License

- This project is licensed under the MIT License.