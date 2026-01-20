package eval

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// constants for error types
const (
	_ int = iota
	PREFIXOP
	INFIXOP
	POSTFIXOP
	MOD_ASSIGNOP
	UNKNOWNIDENT
	UNKNOWNIDENTEX
	NOMETHODERROR
	NOMETHODERROREX
	NOINDEXERROR
	KEYERROR
	INDEXERROR
	SLICEERROR
	ARGUMENTERROR
	INPUTERROR
	RTERROR
	PARAMTYPEERROR
	INLENERR
	INVALIDARG
	DIVIDEBYZERO
	THROWERROR
	THROWNOTHANDLED
	GREPMAPNOTITERABLE
	NOTITERABLE
	RANGETYPEERROR
	DEFERERROR
	SPAWNERROR
	ASSERTIONERROR
	//	STDLIBERROR
	NULLABLEERROR
	JSONERROR
	DBSCANERROR
	FUNCCALLBACKERROR
	FILEMODEERROR
	FILEOPENERROR
	NOTCLASSERROR
	PARENTNOTDECL
	CLSNOTDEFINE
	CLSMEMBERPRIVATE
	CLSCALLPRIVATE
	PROPERTYUSEERROR
	MEMBERUSEERROR
	INDEXERUSEERROR
	INDEXERTYPEERROR
	INDEXERSTATICERROR
	INDEXNOTFOUNDERROR
	CALLNONSTATICERROR
	CLASSCATEGORYERROR
	CLASSCREATEERROR
	PARENTNOTANNOTATION
	OVERRIDEERROR
	METAOPERATORERROR
	SERVICENOURLERROR
	CONSTNOTASSIGNERROR
	DIAMONDOPERERROR
	NAMENOTEXPORTED
	IMPORTERROR
	GENERICERROR
)

var errorType = map[int]string{
	PREFIXOP:            "\x1b[31munsupported prefix operator\x1b[0m: \x1b[36m'%s'\x1b[0m on value of type \x1b[36m%s\x1b[0m",
	INFIXOP:             "\x1b[31munsupported infix operator\x1b[0m: \x1b[36m%s '%s' %s\x1b[0m",
	POSTFIXOP:           "\x1b[31munsupported postfix operator\x1b[0m: \x1b[36m'%s'\x1b[0m on value of type \x1b[36m%s\x1b[0m",
	MOD_ASSIGNOP:        "\x1b[31mmodulus assignment operator not supported\x1b[0m for \x1b[36m'%s'\x1b[0m",
	UNKNOWNIDENT:        "\x1b[31munknown identifier\x1b[0m: \x1b[36m'%s'\x1b[0m is not defined",
	UNKNOWNIDENTEX:      "\x1b[31midentifier not found\x1b[0m: \x1b[36m'%s'\x1b[0m — \x1b[33mdid you mean one of: %s?\x1b[0m",
	NOMETHODERROR:       "\x1b[31mundefined method\x1b[0m \x1b[36m'%s'\x1b[0m for object of type \x1b[36m%s\x1b[0m",
	NOMETHODERROREX:     "\x1b[31mundefined method\x1b[0m \x1b[36m'%s'\x1b[0m for object \x1b[36m'%s'\x1b[0m — \x1b[33mdid you mean one of: %s?\x1b[0m",
	NOINDEXERROR:        "\x1b[31mtype is not indexable\x1b[0m: \x1b[36m%s\x1b[0m cannot be indexed",
	KEYERROR:            "\x1b[31mtype is not hashable\x1b[0m: \x1b[36m%s\x1b[0m cannot be used as a map key",
	INDEXERROR:          "\x1b[31mindex out of range\x1b[0m: \x1b[36m%d\x1b[0m is invalid",
	SLICEERROR:          "\x1b[31mslice indices out of range\x1b[0m: \x1b[36m%d:%d\x1b[0m is invalid",
	ARGUMENTERROR:       "\x1b[31mwrong number of arguments\x1b[0m — expected \x1b[36m%s\x1b[0m, got \x1b[36m%d\x1b[0m",
	INPUTERROR:          "\x1b[31munsupported input type\x1b[0m \x1b[36m'%s'\x1b[0m for function/method \x1b[36m%s\x1b[0m",
	RTERROR:             "\x1b[31mreturn type mismatch\x1b[0m — expected \x1b[36m%s\x1b[0m",
	PARAMTYPEERROR:      "\x1b[31mincorrect argument type\x1b[0m — parameter \x1b[36m%s\x1b[0m of \x1b[36m'%s'\x1b[0m should be \x1b[36m%s\x1b[0m, got \x1b[36m%s\x1b[0m",
	INLENERR:            "\x1b[31minput too long\x1b[0m — function \x1b[36m%s\x1b[0m accepts max length \x1b[36m%s\x1b[0m, got \x1b[36m%s\x1b[0m",
	INVALIDARG:          "\x1b[31minvalid argument supplied\x1b[0m",
	DIVIDEBYZERO:        "\x1b[31mattempt to divide by zero\x1b[0m",
	THROWERROR:          "\x1b[31mthrow operand must be a string\x1b[0m",
	THROWNOTHANDLED:     "\x1b[31munhandled throw\x1b[0m: \x1b[36m'%s'\x1b[0m was not caught",
	GREPMAPNOTITERABLE:  "\x1b[31mgrep/map requires an iterable value\x1b[0m — got non-iterable type",
	NOTITERABLE:         "\x1b[31mforeach requires an iterable value\x1b[0m — got non-iterable type",
	RANGETYPEERROR:      "\x1b[31mrange() expects type\x1b[0m \x1b[36m%s\x1b[0m, got \x1b[36m'%s'\x1b[0m",
	DEFERERROR:          "\x1b[31mdefer used outside a function or not followed by a callable\x1b[0m",
	SPAWNERROR:          "\x1b[31mspawn must be followed by a function call\x1b[0m",
	ASSERTIONERROR:      "\x1b[31massertion failed\x1b[0m",
	NULLABLEERROR:       "\x1b[31mnull reference error\x1b[0m: \x1b[36m%s\x1b[0m is null",
	JSONERROR:           "\x1b[31mJSON error\x1b[0m — unsupported type or malformed data",
	DBSCANERROR:         "\x1b[31mscan failed\x1b[0m — unsupported target type",
	FUNCCALLBACKERROR:   "\x1b[31mcallback signature mismatch\x1b[0m — expected \x1b[36m%d\x1b[0m parameter(s), got \x1b[36m%d\x1b[0m",
	FILEMODEERROR:       "\x1b[31munknown file mode specified\x1b[0m",
	FILEOPENERROR:       "\x1b[31mfailed to open file\x1b[0m — reason: \x1b[36m%s\x1b[0m",
	NOTCLASSERROR:       "\x1b[31midentifier is not a class\x1b[0m: \x1b[36m%s\x1b[0m",
	PARENTNOTDECL:       "\x1b[31mparent class not declared\x1b[0m: \x1b[36m%s\x1b[0m",
	CLSNOTDEFINE:        "\x1b[31mclass not defined\x1b[0m: \x1b[36m%s\x1b[0m",
	CLSMEMBERPRIVATE:    "\x1b[31mprivate member access denied\x1b[0m — variable \x1b[36m%s\x1b[0m in class \x1b[36m%s\x1b[0m is private",
	CLSCALLPRIVATE:      "\x1b[31mprivate method access denied\x1b[0m — method \x1b[36m%s\x1b[0m in class \x1b[36m%s\x1b[0m is private",
	PROPERTYUSEERROR:    "\x1b[31minvalid property usage\x1b[0m — property \x1b[36m%s\x1b[0m of class \x1b[36m%s\x1b[0m cannot be used this way",
	MEMBERUSEERROR:      "\x1b[31minvalid member usage\x1b[0m — member \x1b[36m%s\x1b[0m of class \x1b[36m%s\x1b[0m is not accessible here",
	INDEXERUSEERROR:     "\x1b[31minvalid indexer usage\x1b[0m — class \x1b[36m%s\x1b[0m does not support indexing in this context",
	INDEXERTYPEERROR:    "\x1b[31minvalid indexer type\x1b[0m — class \x1b[36m%s\x1b[0m only supports integer-based indexing",
	INDEXERSTATICERROR:  "\x1b[31mindexers cannot be static\x1b[0m — class \x1b[36m%s\x1b[0m declared indexer as static (not allowed)",
	INDEXNOTFOUNDERROR:  "\x1b[31mno indexer defined\x1b[0m for class \x1b[36m%s\x1b[0m",
	CALLNONSTATICERROR:  "\x1b[31mcannot call non-static member without an instance\x1b[0m",
	CLASSCATEGORYERROR:  "\x1b[31mno class found\x1b[0m — class \x1b[36m%s\x1b[0m not found in category \x1b[36m%s\x1b[0m",
	CLASSCREATEERROR:    "\x1b[31muse 'new' to instantiate classes\x1b[0m — class \x1b[36m'%s'\x1b[0m requires constructor call",
	PARENTNOTANNOTATION: "\x1b[31mparent of annotation must also be an annotation\x1b[0m — parent \x1b[36m%s\x1b[0m of \x1b[36m%s\x1b[0m is not an annotation",
	OVERRIDEERROR:       "\x1b[31mmethod must override a superclass method\x1b[0m — \x1b[36m%s\x1b[0m in class \x1b[36m%s\x1b[0m has no matching parent method",
	METAOPERATORERROR:   "\x1b[31mmeta-operators only accept numbers or strings\x1b[0m",
	SERVICENOURLERROR:   "\x1b[31mservice endpoint missing URL\x1b[0m — function \x1b[36m'%s'\x1b[0m in service \x1b[36m%s\x1b[0m requires a URL annotation",
	CONSTNOTASSIGNERROR: "\x1b[31mcannot assign to constant\x1b[0m — variable \x1b[36m'%s'\x1b[0m is read-only",
	DIAMONDOPERERROR:    "\x1b[31mdiamond operator must be followed by a file object\x1b[0m, got \x1b[36m'%s'\x1b[0m",
	NAMENOTEXPORTED:     "\x1b[31munexported name not accessible\x1b[0m — \x1b[36m%s.%s\x1b[0m is not exported",
	IMPORTERROR:         "\x1b[31mimport failed\x1b[0m: \x1b[36m%s\x1b[0m",
	GENERICERROR:        "\x1b[31m%s\x1b[0m",
}

func NewError(line string, t int, args ...interface{}) Object {
	msg := fmt.Sprintf(errorType[t], args...)
	return &Error{Kind: t, Message: msg, PosMarker: line}
}

type Error struct {
	Kind      int
	Message   string
	PosMarker string // e.g., "</path/to/file.mk:5>"
}

func (e Error) Error() string {
	return e.Message + " (at line) " + strings.TrimLeft(e.PosMarker, " \t")
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }

// extractIdentifier tries to pull the quoted identifier from common error messages
func (e *Error) extractIdentifier() string {
	// Match 'identifier' inside single quotes
	re := regexp.MustCompile(`'([^']*)'`)
	matches := re.FindStringSubmatch(e.Message)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func (e *Error) Inspect() string {
	var sb strings.Builder

	// Parse PosMarker: supports both </file:line> and </file:line:column>
	rePos := regexp.MustCompile(`<(/[^:>]+):(\d+)(?::(\d+))?>`)
	matches := rePos.FindStringSubmatch(e.PosMarker)
	if matches == nil {
		sb.WriteString("\033[1;31merror:\033[0m ")
		sb.WriteString(e.Message)
		sb.WriteString("\n")
		return sb.String()
	}

	filePath := matches[1]
	lineNum, _ := strconv.Atoi(matches[2])
	colNum := 1 // 1-based, default
	if matches[3] != "" {
		if c, err := strconv.Atoi(matches[3]); err == nil && c > 0 {
			colNum = c
		}
	}
	if lineNum <= 0 {
		lineNum = 1
	}

	// Read the source line
	var sourceLine string
	if content, err := ioutil.ReadFile(filePath); err == nil {
		lines := strings.Split(string(content), "\n")
		if lineNum <= len(lines) {
			sourceLine = strings.TrimRight(lines[lineNum-1], "\r\n")
		}
	}

	// Try to extract quoted identifier (e.g., 'hfjdk')
	identifier := e.extractIdentifier()

	var underline string

	if identifier != "" && sourceLine != "" {
		// Find the identifier in the source line
		/*if idx := strings.Index(sourceLine, identifier); idx != -1 {
			// Build underline exactly under the identifier
			prefix := strings.Repeat(" ", idx)
			mark := strings.Repeat("^", len(identifier))
			underline = prefix + "\033[1;31m" + mark + "\033[0m"
		} else {*/
		// Identifier not found in source — fall back to column
		colIdx := colNum - 1 // to 0-based
		if colIdx < 0 {
			colIdx = 0
		}
		if colIdx >= len(sourceLine) {
			colIdx = len(sourceLine)
		}
		prefix := strings.Repeat(" ", colIdx)
		// Underline at least one char
		underline = prefix + "\033[1;31m^\033[0m"
		//}
	} else {
		// No identifier: use column-based underline
		colIdx := colNum - 1 // 0-based
		if colIdx < 0 {
			colIdx = 0
		}
		if colIdx >= len(sourceLine) {
			colIdx = len(sourceLine)
		}
		prefix := strings.Repeat(" ", colIdx)
		underline = prefix + "\033[1;31m^\033[0m"
	}

	// Format output
	sb.WriteString(fmt.Sprintf("\033[1;90m-->\033[0m \033[1m%s:%d\033[0m\n", filePath, lineNum))
	sb.WriteString("\033[1;90m |\033[0m\n")

	lineNumStr := strconv.Itoa(lineNum)
	sb.WriteString(fmt.Sprintf("\033[1;37m%2s\033[0m\033[1;90m\033[0m \033[90m%s\033[0m\n", lineNumStr, sourceLine))

	if underline != "" {
		sb.WriteString(fmt.Sprintf("\033[1;90m |\033[0m   %s\n", underline))
	}

	sb.WriteString("\033[1;90m |\033[0m\n")
	sb.WriteString(fmt.Sprintf("\033[1;31merror:\033[0m %s\n", e.Message))

	return sb.String()
}

func (e *Error) CallMethod(line string, scope *Scope, method string, args ...Object) Object {
	return NewError(line, GENERICERROR, e.Message)
}
