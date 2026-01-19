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
	PREFIXOP:           "unsupported operator for prefix expression:'%s' and type: %s",
	INFIXOP:            "unsupported operator for infix expression: %s '%s' %s",
	POSTFIXOP:          "unsupported operator for postfix expression:'%s' and type: %s",
	MOD_ASSIGNOP:       "unsupported operator for modulor assignment:'%s'",
	UNKNOWNIDENT:       "unknown identifier: '%s' is not defined",
	UNKNOWNIDENTEX:     "identifier '%s' not found. \n\nDid you mean one of: \n\n  %s\n",
	NOMETHODERROR:      "undefined method '%s' for object %s",
	NOMETHODERROREX:    "undefined method '%s' for object '%s'. \n\nDid you mean one of: \n\n  %s\n",
	NOINDEXERROR:       "index error: type %s is not indexable",
	KEYERROR:           "key error: type %s is not hashable",
	INDEXERROR:         "index error: '%d' out of range",
	SLICEERROR:         "index error: slice '%d:%d' out of range",
	ARGUMENTERROR:      "wrong number of arguments. expected=%s, got=%d",
	INPUTERROR:         "unsupported input type '%s' for function or method: %s",
	RTERROR:            "return type should be %s",
	PARAMTYPEERROR:     "%s argument for '%s' should be type %s. got=%s",
	INLENERR:           "function %s takes input with max length %s. got=%s",
	INVALIDARG:         "invalid argument supplied",
	DIVIDEBYZERO:       "divide by zero",
	THROWERROR:         "throw object must be a string",
	THROWNOTHANDLED:    "throw object '%s' not handled",
	GREPMAPNOTITERABLE: "grep/map's operating type must be iterable",
	NOTITERABLE:        "foreach's operating type must be iterable",
	RANGETYPEERROR:     "range(..) type should be %s type, got='%s'",
	DEFERERROR:         "defer outside function or defer statement not a function",
	SPAWNERROR:         "spawn must be followed by a function",
	ASSERTIONERROR:     "assertion failed",
	//	STDLIBERROR:     "calling '%s' failed",
	NULLABLEERROR:       "%s is null",
	JSONERROR:           "json error: maybe unsupported type or invalid data",
	DBSCANERROR:         "scan type not supported",
	FUNCCALLBACKERROR:   "callback error: must be '%d' parameter(s), got '%d'",
	FILEMODEERROR:       "known file mode supplied",
	FILEOPENERROR:       "file open failed, reason: %s",
	NOTCLASSERROR:       "Identifier %s is not a class",
	PARENTNOTDECL:       "Parent class %s not declared",
	CLSNOTDEFINE:        "Class %s not defined",
	CLSMEMBERPRIVATE:    "Variable(%s) of class(%s) is private",
	CLSCALLPRIVATE:      "Method (%s) of class(%s) is private",
	PROPERTYUSEERROR:    "Invalid use of Property(%s) of class(%s)",
	MEMBERUSEERROR:      "Invalid use of member(%s) of class(%s)",
	INDEXERUSEERROR:     "Invalid use of Indexer of class(%s)",
	INDEXERTYPEERROR:    "Invalid use of Indexer of class(%s), Only interger type of Indexer is supported",
	INDEXERSTATICERROR:  "Invalid use of Indexer of class(%s), Indexer cannot declared as static",
	INDEXNOTFOUNDERROR:  "Indexer not found for class(%s)",
	CALLNONSTATICERROR:  "Could not call non-static",
	CLASSCATEGORYERROR:  "No class(%s) found for category(%s)",
	CLASSCREATEERROR:    "You must use 'new' to create class('%s')",
	PARENTNOTANNOTATION: "Annotation(%s)'s Parent(%s) is not annotation",
	OVERRIDEERROR:       "Method(%s) of class(%s) must override a superclass method",
	METAOPERATORERROR:   "Meta-Operators' item must be Numbers|String",
	SERVICENOURLERROR:   "Service(%s)'s function('%s') must have url",
	CONSTNOTASSIGNERROR: "Const variable '%s' cannot be modified",
	DIAMONDOPERERROR:    "Diamond operator must be followed by a file object, but got '%s'",
	NAMENOTEXPORTED:     "Cannot refer to unexported name '%s.%s'",
	IMPORTERROR:         "Import error: %s",
	GENERICERROR:        "%s",
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

	// Parse file and line
	rePos := regexp.MustCompile(`<(/[^:>]+):(\d+)>`)
	matches := rePos.FindStringSubmatch(e.PosMarker)
	if matches == nil {
		// Fallback if position not parseable
		sb.WriteString("\033[1;31merror:\033[0m ")
		sb.WriteString(e.Message)
		sb.WriteString("\n")
		return sb.String()
	}

	filePath := matches[1]
	lineNum, _ := strconv.Atoi(matches[2])
	if lineNum <= 0 {
		lineNum = 1
	}

	// Read source line
	var sourceLine string
	if content, err := ioutil.ReadFile(filePath); err == nil {
		lines := strings.Split(string(content), "\n")
		if lineNum <= len(lines) {
			sourceLine = strings.TrimRight(lines[lineNum-1], "\r\n")
		}
	}

	// Try to find the problematic identifier (e.g., 'y')
	identifier := e.extractIdentifier()

	// Build underline
	var underline string
	if identifier != "" && sourceLine != "" {
		if idx := strings.Index(sourceLine, identifier); idx != -1 {
			prefix := strings.Repeat(" ", idx)
			mark := strings.Repeat("^", len(identifier))
			underline = prefix + "\033[1;31m" + mark + "\033[0m"
		} else {
			// Fallback: underline first non-whitespace
			trimmed := len(sourceLine) - len(strings.TrimLeft(sourceLine, " \t"))
			underline = strings.Repeat(" ", trimmed) + "\033[1;31m^\033[0m"
		}
	} else {
		// Just point to start
		underline = "\033[1;31m^\033[0m"
	}

	// Format like Rust:
	// --> FILE:LINE
	//   |
	// L | code
	//   | ^~~~
	//   |
	// error: message

	sb.WriteString(fmt.Sprintf("\033[1;90m-->\033[0m \033[1m%s:%d\033[0m\n", filePath, lineNum))
	sb.WriteString("\033[1;90m |\033[0m\n")

	// Line number formatting (right-aligned to 2 digits, but flexible)
	lineNumStr := strconv.Itoa(lineNum)
	sb.WriteString(fmt.Sprintf("\033[1;37m%2s\033[0m\033[1;90m |\033[0m \033[90m%s\033[0m\n", lineNumStr, sourceLine))

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
