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
	"os/user"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
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

		dbgInfosArr := parser.SplitSlice(parser.DebugInfos)
		eval.Dbg.SetDbgInfos(dbgInfosArr)

		eval.MsgHandler = message.NewMessageHandler()
		eval.MsgHandler.AddListener(eval.Dbg)
	}

	result := eval.Eval(program, scope)
	if result.Type() == eval.ERROR_OBJ {
		fmt.Println(result.Inspect())
	}
}

// Register go package methods/types
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

// ──────────────────────────────────────────────────────────────────────────────
// Package config (flat key: value format)
// ──────────────────────────────────────────────────────────────────────────────

type PackageConfig struct {
	Packages map[string]string
}

const (
	configDirName  = "mk"
	configFileName = "PATH.yaml"
)

func getConfigPath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	var configBase string
	if runtime.GOOS == "windows" {
		appData := os.Getenv("APPDATA")
		if appData == "" {
			appData = filepath.Join(u.HomeDir, "AppData", "Roaming")
		}
		configBase = filepath.Join(appData, configDirName, "cfg")
	} else {
		configBase = filepath.Join(u.HomeDir, "."+configDirName, "cfg")
	}

	return filepath.Join(configBase, configFileName), nil
}

func loadPackageConfig() (*PackageConfig, error) {
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &PackageConfig{Packages: make(map[string]string)}, nil
		}
		return nil, err
	}

	packages := make(map[string]string)
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// ~ expansion
		if strings.HasPrefix(value, "~/") {
			home, _ := os.UserHomeDir()
			value = filepath.Join(home, value[2:])
		} else if value == "~" {
			home, _ := os.UserHomeDir()
			value = home
		}

		if key != "" {
			packages[key] = filepath.Clean(value)
		}
	}

	return &PackageConfig{Packages: packages}, nil
}

func savePackageConfig(cfg *PackageConfig) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	var sb strings.Builder

	keys := make([]string, 0, len(cfg.Packages))
	for k := range cfg.Packages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, name := range keys {
		p := cfg.Packages[name]

		// Try to use ~ when possible
		home, _ := os.UserHomeDir()
		if strings.HasPrefix(p, home) {
			rel := strings.TrimPrefix(p, home)
			if rel == "" {
				p = "~"
			} else if strings.HasPrefix(rel, string(os.PathSeparator)) {
				p = "~" + rel
			}
		}

		sb.WriteString(fmt.Sprintf("%s: %s\n", name, p))
	}

	return os.WriteFile(path, []byte(sb.String()), 0644)
}

// ──────────────────────────────────────────────────────────────────────────────
// Package management commands
// ──────────────────────────────────────────────────────────────────────────────

func handlePkgCommand(args []string) error {
	if len(args) == 0 {
		fmt.Println("Usage:")
		fmt.Println("  monkey pkg -init <name> [-dir <path>]")
		fmt.Println("  monkey pkg -list")
		return nil
	}

	switch args[0] {
	case "-list", "--list":
		return cmdPkgList()

	case "-init", "--init":
		return cmdPkgInit(args[1:])

	default:
		return fmt.Errorf("unknown pkg command: %s\nUse one of: -init, -list", args[0])
	}
}

func cmdPkgList() error {
	cfg, err := loadPackageConfig()
	if err != nil {
		return err
	}

	if len(cfg.Packages) == 0 {
		fmt.Println("No packages registered yet.")
		p, _ := getConfigPath()
		fmt.Printf("Config file: %s\n", p)
		return nil
	}

	fmt.Println("Known Monkey packages:")
	fmt.Println("")

	for _, name := range sortedKeys(cfg.Packages) {
		fmt.Printf("%-12s %s\n", name+":", cfg.Packages[name])
	}

	return nil
}

func cmdPkgInit(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("package name is required\nExample: monkey pkg -init mylib")
	}

	pkgName := args[0]
	dirPath := "."

	for i := 1; i < len(args)-1; i++ {
		if args[i] == "-dir" || args[i] == "--dir" {
			dirPath = args[i+1]
			break
		}
	}

	absDir, err := filepath.Abs(dirPath)
	if err != nil {
		return err
	}

	if _, err := os.Stat(absDir); os.IsNotExist(err) {
		if err := os.MkdirAll(absDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}

	filesToCreate := []struct {
		path    string
		content string
	}{}

	for _, f := range filesToCreate {
		full := filepath.Join(absDir, f.path)
		if _, err := os.Stat(full); err == nil {
			fmt.Printf("Warning: %s already exists, skipping...\n", full)
			continue
		}
		if err := os.WriteFile(full, []byte(f.content), 0644); err != nil {
			return fmt.Errorf("failed to create %s: %v", full, err)
		}
	}

	cfg, err := loadPackageConfig()
	if err != nil {
		return err
	}

	cfg.Packages[pkgName] = absDir

	if err := savePackageConfig(cfg); err != nil {
		return fmt.Errorf("failed to save config: %v", err)
	}

	fmt.Printf("Package '%s' created at:\n  %s\n", pkgName, absDir)
	fmt.Println("Registered in package config.")
	return nil
}

func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func main() {
	debug := flag.Bool("d", false, "enable debug mode")
	flag.BoolVar(debug, "debug", false, "enable debug mode")

	expr := flag.String("e", "", "execute code from string")

	help := flag.Bool("h", false, "display help")
	flag.BoolVar(help, "help", false, "display help")

	flag.Parse()

	args := flag.Args()

	// Package commands
	if len(args) >= 1 && args[0] == "pkg" {
		if err := handlePkgCommand(args[1:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if *help {
		usage()
		return
	}

	if *expr != "" {
		if err := runExpr(*expr, *debug); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		return
	}

	if len(args) == 0 {
		fmt.Println("monkey programming language REPL\n")
		repl.Start(os.Stdout, true)
		return
	}

	runProgram(*debug, args[0])
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  monkey                       Start REPL")
	fmt.Println("  monkey file.mk               Run file")
	fmt.Println("  monkey -d file.mk            Run file in debug mode")
	fmt.Println("  monkey -e \"code\"             Execute code string")
	fmt.Println("  monkey pkg -init <name> [-dir <path>]   Create new package")
	fmt.Println("  monkey pkg -list             List registered packages")
	fmt.Println("  monkey -h                    Show help")
}

func runExpr(code string, debug bool) error {
	tmpDir := filepath.Join(os.TempDir(), "monkey")
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return err
	}

	tmpFile, err := os.CreateTemp(tmpDir, "monkey-*.mk")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(code); err != nil {
		return err
	}
	tmpFile.Close()

	runProgram(debug, tmpFile.Name())
	return nil
}
