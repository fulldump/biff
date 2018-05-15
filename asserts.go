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

func (t *A) AssertDistinct(obtained, expected interface{}) bool {
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

func (t *A) AssertEqual(obtained, expected interface{}) bool {

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

func (t *A) AssertEqualJson(obtained, expected interface{}) bool {

	b, _ := json.Marshal(expected)

	e := interface{}(nil)
	json.Unmarshal(b, &e)

	if reflect.DeepEqual(e, obtained) {
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

func (t *A) AssertNil(obtained interface{}) bool {

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

func (t *A) AssertNotNil(obtained interface{}) bool {

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

		a_func, ok := a.(*ast.CallExpr)
		if !ok {
			return
		}

		a_func_ident, ok := a_func.Args[0].(*ast.Ident)
		if ok {
			variable = a_func_ident.String()
			return
		}

		a_func_expr, ok := a_func.Args[0].(*ast.IndexExpr)
		if ok {
			variable = l[a_func_expr.Pos()-1 : a_func_expr.End()-1]
			return
		}

	}()

	fmt.Printf("    %s is %#v\n", variable, value)
}
