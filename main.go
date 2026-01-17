package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"monkey/eval"
	"monkey/lexer"
	"monkey/message"
	"monkey/parser"
	"monkey/repl"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

func runProgram(debug bool, filename string) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	f := filename

	content, err := os.ReadFile(f)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	input := string(content)

	l := lexer.New(filename, input)

	p := parser.New(l, wd)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			fmt.Println(err)
		}
		os.Exit(1)
	}
	scope := eval.NewScope(nil, os.Stdout)
	RegisterGoGlobals()

	if debug {
		eval.REPLColor = true
		eval.Dbg = eval.NewDebugger()
		eval.Dbg.SetFunctions(p.Functions)
		eval.Dbg.ShowBanner()

		dbgInfosArr := parser.SplitSlice(parser.DebugInfos) //[][]ast.Node
		eval.Dbg.SetDbgInfos(dbgInfosArr)
		// for idx, dbgInfos := range dbgInfosArr {
		// 	for _, dbgInfo := range dbgInfos {
		// 		fmt.Printf("idx:%d, Line:<%d-%d>, node.Type=%T, node=<%s>\n", idx, dbgInfo.Pos().Line, dbgInfo.End().Line, dbgInfo, dbgInfo.String())
		// 	}
		// }

		eval.MsgHandler = message.NewMessageHandler()
		eval.MsgHandler.AddListener(eval.Dbg)

	}

	result := eval.Eval(program, scope)
	if result.Type() == eval.ERROR_OBJ {
		fmt.Println(result.Inspect())
	}

	// e := eval.Eval(program, scope)
	//
	//	if e.Inspect() != "nil" {
	//		fmt.Println(e.Inspect())
	//	}
}

// Register go package methods/types
// Note here, we use 'gfmt', 'glog', 'gos' 'gtime', because in magpie
// we already have built in module 'fmt', 'log' 'os', 'time'.
// And Here we demonstrate the use of import go language's methods.
func RegisterGoGlobals() {
	eval.RegisterFunctions("gfmt", map[string]interface{}{
		"Errorf":   fmt.Errorf,
		"Println":  fmt.Println,
		"Print":    fmt.Print,
		"Printf":   fmt.Printf,
		"Fprint":   fmt.Fprint,
		"Fprintln": fmt.Fprintln,
		"Fscan":    fmt.Fscan,
		"Fscanf":   fmt.Fscanf,
		"Fscanln":  fmt.Fscanln,
		"Scan":     fmt.Scan,
		"Scanf":    fmt.Scanf,
		"Scanln":   fmt.Scanln,
		"Sscan":    fmt.Sscan,
		"Sscanf":   fmt.Sscanf,
		"Sscanln":  fmt.Sscanln,
		"Sprint":   fmt.Sprint,
		"Sprintf":  fmt.Sprintf,
		"Sprintln": fmt.Sprintln,
	})

	eval.RegisterFunctions("glog", map[string]interface{}{
		"Fatal":     log.Fatal,
		"Fatalf":    log.Fatalf,
		"Fatalln":   log.Fatalln,
		"Flags":     log.Flags,
		"Panic":     log.Panic,
		"Panicf":    log.Panicf,
		"Panicln":   log.Panicln,
		"Print":     log.Print,
		"Printf":    log.Printf,
		"Println":   log.Println,
		"SetFlags":  log.SetFlags,
		"SetOutput": log.SetOutput,
		"SetPrefix": log.SetPrefix,
	})

	eval.RegisterFunctions("gos", map[string]interface{}{
		"Chdir":    os.Chdir,
		"Chmod":    os.Chmod,
		"Chown":    os.Chown,
		"Exit":     os.Exit,
		"Getpid":   os.Getpid,
		"Hostname": os.Hostname,
		"Environ":  os.Environ,
		"Getenv":   os.Getenv,
		"Setenv":   os.Setenv,
		"Create":   os.Create,
		"Open":     os.Open,
	})

	argsStart := 1
	if len(os.Args) > 2 {
		argsStart = 2
	}
	eval.RegisterVars("gos", map[string]interface{}{
		"Args": os.Args[argsStart:],
	})

	eval.RegisterVars("runtime", map[string]interface{}{
		"GOOS":   runtime.GOOS,
		"GOARCH": runtime.GOARCH,
	})

	eval.RegisterVars("gtime", map[string]interface{}{
		"Duration": time.Duration(0),
		"Ticker":   time.Ticker{},
		"Time":     time.Time{},
	})
	eval.RegisterFunctions("gtime", map[string]interface{}{
		"After":           time.After,
		"Sleep":           time.Sleep,
		"Tick":            time.Tick,
		"Since":           time.Since,
		"FixedZone":       time.FixedZone,
		"LoadLocation":    time.LoadLocation,
		"NewTicker":       time.NewTicker,
		"Date":            time.Date,
		"Now":             time.Now,
		"Parse":           time.Parse,
		"ParseDuration":   time.ParseDuration,
		"ParseInLocation": time.ParseInLocation,
		"Unix":            time.Unix,
		"AfterFunc":       time.AfterFunc,
		"NewTimer":        time.NewTimer,
		"Nanosecond":      time.Nanosecond,
		"Microsecond":     time.Microsecond,
		"Millisecond":     time.Millisecond,
		"Second":          time.Second,
		"Minute":          time.Minute,
		"Hour":            time.Hour,
	})

	eval.RegisterFunctions("math/rand", map[string]interface{}{
		"New":         rand.New,
		"NewSource":   rand.NewSource,
		"Float64":     rand.Float64,
		"ExpFloat64":  rand.ExpFloat64,
		"Float32":     rand.Float32,
		"Int":         rand.Int,
		"Int31":       rand.Int31,
		"Int31n":      rand.Int31n,
		"Int63":       rand.Int63,
		"Int63n":      rand.Int63n,
		"Intn":        rand.Intn,
		"NormFloat64": rand.NormFloat64,
		"Perm":        rand.Perm,
		"Seed":        rand.Seed,
		"Uint32":      rand.Uint32,
	})

	eval.RegisterFunctions("io/ioutil", map[string]interface{}{
		"WriteFile": ioutil.WriteFile,
		"ReadFile":  ioutil.ReadFile,
		"TempDir":   ioutil.TempDir,
		"TempFile":  ioutil.TempFile,
		"ReadAll":   ioutil.ReadAll,
		"ReadDir":   ioutil.ReadDir,
		"NopCloser": ioutil.NopCloser,
	})

	eval.RegisterFunctions("bufio", map[string]interface{}{
		"NewWriter":     bufio.NewWriter,
		"NewReader":     bufio.NewReader,
		"NewReadWriter": bufio.NewReadWriter,
		"NewScanner":    bufio.NewScanner,
	})
	eval.RegisterFunctions("gregex", map[string]interface{}{
		"Match":            regexp.Match,
		"MatchReader":      regexp.MatchReader,
		"MatchString":      regexp.MatchString,
		"QuoteMeta":        regexp.QuoteMeta,
		"Compile":          regexp.Compile,
		"CompilePOSIX":     regexp.CompilePOSIX,
		"MustCompile":      regexp.MustCompile,
		"MustCompilePOSIX": regexp.MustCompilePOSIX,
	})
}

func main() {
	// Flags
	debug := flag.Bool("d", false, "enable debug mode")
	flag.BoolVar(debug, "debug", false, "enable debug mode")

	expr := flag.String("e", "", "execute code from string")

	help := flag.Bool("h", false, "display help")
	flag.BoolVar(help, "help", false, "display help")

	flag.Parse()

	// Help
	if *help {
		usage()
		return
	}

	// Execute code from string using -e
	if *expr != "" {
		if err := runExpr(*expr, *debug); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		return
	}

	// No args -> REPL
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("monkey programming language REPL\n")
		repl.Start(os.Stdout, true)
		return
	}

	// Run file
	runProgram(*debug, args[0])
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  monkey                 Start REPL")
	fmt.Println("  monkey file.mk         Run file")
	fmt.Println("  monkey -d file.mk      Run file in debug mode")
	fmt.Println("  monkey -e \"code\"       Execute code string")
	fmt.Println("  monkey -h              Show help")
}

// write code to temp file then run it
func runExpr(code string, debug bool) error {
	// Force system temp dir
	tmpDir := os.TempDir()

	// If os.TempDir returns a relative dir, make it absolute
	if !filepath.IsAbs(tmpDir) {
		tmpDir, _ = filepath.Abs(tmpDir)
	}

	// Create /tmp/monkey folder to be safe
	tmpDir = filepath.Join(tmpDir, "monkey")
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return err
	}

	// Create a temp file in that directory
	tmpFile, err := os.CreateTemp(tmpDir, "monkey-*.mk")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	// Write the code into it
	if _, err := tmpFile.WriteString(code); err != nil {
		return err
	}
	tmpFile.Close()

	// Debug: print file path
	// fmt.Println("Temp file created at:", tmpFile.Name())

	// Run the temp file
	// debug: println(tmpFile.Name())
	runProgram(debug, tmpFile.Name())
	return nil
}
