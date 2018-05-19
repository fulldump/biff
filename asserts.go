package biff

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
)

// AssertNotEqual return true if `obtained` is not equal to `expected` otherwise
// it will print trace and exit.
func (a *A) AssertNotEqual(obtained, expected interface{}) bool {
	if !reflect.DeepEqual(expected, obtained) {
		printShould(expected)
		return true
	}

	line := getStackLine(2)
	fmt.Printf(""+
		"    Expected: %#v\n"+
		"    Obtained: %#v\n"+
		"    at %s\n", expected, obtained, line)

	os.Exit(1)

	return false
}

// AssertEqual return true if `obtained` is equal to `expected` otherwise it
// will print trace and exit.
func (a *A) AssertEqual(obtained, expected interface{}) bool {

	if reflect.DeepEqual(expected, obtained) {
		printShould(expected)
		return true
	}

	line := getStackLine(2)
	fmt.Printf(""+
		"    Expected: %#v\n"+
		"    Obtained: %#v\n"+
		"    at %s\n", expected, obtained, line)

	os.Exit(1)

	return false
}

func readFileLine(filename string, line int) string {

	data, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(data), "\n")

	return lines[line-1]
}

// AssertEqualJson return true if `obtained` is equal to `expected`. Prior to
// comparison, both values are JSON Marshaled/Unmarshaled to avoid JSON type
// issues like int vs float etc. Otherwise it will print trace and exit.
func (a *A) AssertEqualJson(obtained, expected interface{}) bool {

	e := interface{}(nil)
	{
		b, _ := json.Marshal(expected)
		json.Unmarshal(b, &e)
	}

	o := interface{}(nil)
	{
		b, _ := json.Marshal(obtained)
		json.Unmarshal(b, &o)
	}

	if reflect.DeepEqual(e, o) {
		printShould(expected)
		return true
	}

	line := getStackLine(2)
	fmt.Printf(""+
		"    Expected: %#v\n"+
		"    Obtained: %#v\n"+
		"    at %s\n", e, o, line)

	os.Exit(1)

	return false
}

// AssertNil return true if `obtained` is nil, otherwise it will print trace and
// exit.
func (a *A) AssertNil(obtained interface{}) bool {

	if nil == obtained || reflect.ValueOf(obtained).IsNil() {
		printShould(nil)
		return true
	}

	line := getStackLine(2)
	fmt.Printf(""+
		"    Expected: nil\n"+
		"    Obtained: %#v\n"+
		"    at %s\n", obtained, line)

	os.Exit(1)

	return false
}

// AssertNotNil return true if `obtained` is NOT nil, otherwise it will print trace
// and exit.
func (a *A) AssertNotNil(obtained interface{}) bool {

	line := getStackLine(2)
	if nil == obtained || reflect.ValueOf(obtained).IsNil() {
		fmt.Printf(""+
			"    Expected: not nil\n"+
			"    Obtained: %#v\n"+
			"    at %s\n", obtained, line)
		return false
	}

	printShould(obtained)

	return true
}

func getStackLine(linesToSkip int) string {

	stack := debug.Stack()
	lines := make([]string, 0)
	index := 0
	for i := 0; i < len(stack); i++ {
		if stack[i] == []byte("\n")[0] {
			lines = append(lines, string(stack[index:i-1]))
			index = i + 1
		}
	}
	return lines[linesToSkip*2+3] + " " + lines[linesToSkip*2+4]
}

func printShould(value interface{}) {
	variable := "It"

	func() {

		p := make([]runtime.StackRecord, 50)

		_, ok := runtime.GoroutineProfile(p)
		if !ok {
			return
		}

		frames := runtime.CallersFrames(p[0].Stack())

		frames.Next()
		frames.Next()
		frames.Next()
		frame, _ := frames.Next()

		l := readFileLine(frame.File, frame.Line)

		a, err := parser.ParseExpr(l)
		if nil != err {
			return
		}

		aFunc, ok := a.(*ast.CallExpr)
		if !ok {
			return
		}

		aFuncIdent, ok := aFunc.Args[0].(*ast.Ident)
		if ok {
			variable = aFuncIdent.String()
			return
		}

		aFuncExpr, ok := aFunc.Args[0].(*ast.IndexExpr)
		if ok {
			variable = l[aFuncExpr.Pos()-1 : aFuncExpr.End()-1]
			return
		}

	}()

	fmt.Printf("    %s is %#v\n", variable, value)
}
