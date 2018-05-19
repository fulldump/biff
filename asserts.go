package biff

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
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

	printExpectedObtained(expected, obtained)

	return false
}

// AssertEqual return true if `obtained` is equal to `expected` otherwise it
// will print trace and exit.
func (a *A) AssertEqual(obtained, expected interface{}) bool {

	if reflect.DeepEqual(expected, obtained) {
		printShould(expected)
		return true
	}

	printExpectedObtained(expected, obtained)

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

	printExpectedObtained(e, o)

	return false
}

// AssertNil return true if `obtained` is nil, otherwise it will print trace and
// exit.
func (a *A) AssertNil(obtained interface{}) bool {

	if nil == obtained || reflect.ValueOf(obtained).IsNil() {
		printShould(nil)
		return true
	}

	printExpectedObtained(nil, obtained)

	return false
}

// AssertNotNil return true if `obtained` is NOT nil, otherwise it will print trace
// and exit.
func (a *A) AssertNotNil(obtained interface{}) bool {

	if nil == obtained || reflect.ValueOf(obtained).IsNil() {
		line := getStackLine(2)
		fmt.Printf(""+
			"    Expected: not nil\n"+
			"    Obtained: %#v\n"+
			"    at %s\n", obtained, line)
		return false
	}

	printShould(obtained)

	return true
}

// AssertTrue return true if `obtained` is true, otherwise it will print trace
// and exit.
func (a *A) AssertTrue(obtained interface{}) bool {

	if reflect.DeepEqual(true, obtained) {
		printShould(true)
		return true
	}

	printExpectedObtained(true, obtained)

	return false
}

// AssertFalse return true if `obtained` is false, otherwise it will print trace
// and exit.
func (a *A) AssertFalse(obtained interface{}) bool {

	if reflect.DeepEqual(false, obtained) {
		printShould(false)
		return true
	}

	printExpectedObtained(true, obtained)

	return false
}

// AssertInArray return true if `item` match at least with one element of the
// array. Otherwise it will print trace and exit.
func (a *A) AssertInArray(item, array interface{}) bool {

	v := reflect.ValueOf(array)
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		line := getStackLine(2)
		fmt.Printf("Expected second argument to be array:\n"+
			"    Obtained: %#v\n"+
			"    at %s\n", array, line)
		os.Exit(1)
	}

	l := v.Len()
	for i := 0; i < l; i++ {
		e := v.Index(i)
		if reflect.DeepEqual(e.Interface(), item) {
			printShould(item)
			return true
		}
	}

	line := getStackLine(2)
	fmt.Printf(""+
		"    Expected item to be in array.\n"+
		"    Item: %#v\n"+
		"    Array: %#v\n"+
		"    at %s\n", item, array, line)

	os.Exit(1)

	return false
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

func printExpectedObtained(expected, obtained interface{}) {

	line := getStackLine(3)
	fmt.Printf(""+
		"    Expected: %#v\n"+
		"    Obtained: %#v\n"+
		"    at %s\n", expected, obtained, line)

	os.Exit(1)

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
		if ok && aFuncExpr.Pos() > 0 && (aFuncExpr.End()-1) < token.Pos(len(l)) {
			variable = l[aFuncExpr.Pos()-1 : aFuncExpr.End()-1]
			return
		}

	}()

	fmt.Printf("    %s is %#v\n", variable, value)
}
