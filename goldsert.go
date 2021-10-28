// Package goldsert provides a suite of test helpers which uses golden files to
// assert marshaling and unmarshaling of given objects.
//
// Each test helper operates in two stages:
//
// Firstly they marshal the provided object to a byte slice and reads the
// corresponding golden file from disk, follow by verifying both byte slices are
// identical.
//
// Secondly they unmarshal the content from the golden file, and verifies that
// the result is identical to the original object.
//
// Usage
//
// Typical usage should look something like this in a tabular test:
//
//  type MyStruct struct {
//      FooBar string `json:"foo_bar" yaml:"fooBar" xml:"Foo_Bar"`
//  }
//
//  func TestMyStructMarshaling(t *testing.T) {
//      tests := []struct {
//          name string
//          obj  *MyStruct
//      }{
//          {name: "empty", obj: &MyStruct{}},
//          {name: "full", obj: &MyStruct{FooBar: "Hello World!"}},
//      }
//      for _, tt := range tests {
//          t.Run(tt.name, func(t *testing.T) {
//              goldsert.JSONMarshaling(t, tt.obj)
//              goldsert.YAMLMarshaling(t, tt.obj)
//              goldsert.XMLMarshaling(t, tt.obj)
//          })
//      }
//  }
//
// The above example will read from the following golden files:
//
//  testdata/TestMyStructMarshaling/empty/goldsert_json.golden
//  testdata/TestMyStructMarshaling/empty/goldsert_yaml.golden
//  testdata/TestMyStructMarshaling/empty/goldsert_xml.golden
//  testdata/TestMyStructMarshaling/full/goldsert_json.golden
//  testdata/TestMyStructMarshaling/full/goldsert_yaml.golden
//  testdata/TestMyStructMarshaling/full/goldsert_xml.golden
//
// If a corresponding golden file cannot be found on disk, the test will fail.
// To create/update golden files, simply set the GOLDEN_UPDATE environment
// variable to one of "1", "y", "t", "yes", "on", or "true" when running tests.
//
// It is highly recommended that golden files are committed to source control,
// as it allow tests to fail when the marshal results for an object changes.
package goldsert

import (
	"testing"
)

var global = New()

// JSONMarshaling asserts that the given "v" value JSON marshals to an expected
// value fetched from a golden file on disk, and then verifies that the
// marshaled result produces a value that is equal to "v" when unmarshaled.
//
// Used for objects that do NOT change when they are marshaled and unmarshaled.
func JSONMarshaling(t *testing.T, v interface{}) {
	t.Helper()

	global.JSONMarshaling(t, v)
}

// JSONMarshalingP asserts that the given "v" value JSON marshals to an expected
// value fetched from a golden file on disk, and then verifies that the
// marshaled result produces a value that is equal to "want" when unmarshaled.
//
// Used for objects that change when they are marshaled and unmarshaled.
func JSONMarshalingP(t *testing.T, v, want interface{}) {
	t.Helper()

	global.JSONMarshalingP(t, v, want)
}

// YAMLMarshaling asserts that the given "v" value YAML marshals to an expected
// value fetched from a golden file on disk, and then verifies that the
// marshaled result produces a value that is equal to "v" when unmarshaled.
//
// Used for objects that do NOT change when they are marshaled and unmarshaled.
func YAMLMarshaling(t *testing.T, v interface{}) {
	t.Helper()

	global.YAMLMarshaling(t, v)
}

// YAMLMarshalingP asserts that the given "v" value YAML marshals to an expected
// value fetched from a golden file on disk, and then verifies that the
// marshaled result produces a value that is equal to "want" when unmarshaled.
//
// Used for objects that change when they are marshaled and unmarshaled.
func YAMLMarshalingP(t *testing.T, v, want interface{}) {
	t.Helper()

	global.YAMLMarshalingP(t, v, want)
}

// XMLMarshaling asserts that the given "v" value XML marshals to an expected
// value fetched from a golden file on disk, and then verifies that the
// marshaled result produces a value that is equal to "v" when unmarshaled.
//
// Used for objects that do NOT change when they are marshaled and unmarshaled.
func XMLMarshaling(t *testing.T, v interface{}) {
	t.Helper()

	global.XMLMarshaling(t, v)
}

// XMLMarshalingP asserts that the given "v" value XML marshals to an expected
// value fetched from a golden file on disk, and then verifies that the
// marshaled result produces a value that is equal to "want" when unmarshaled.
//
// Used for objects that change when they are marshaled and unmarshaled.
func XMLMarshalingP(t *testing.T, v, want interface{}) {
	t.Helper()

	global.XMLMarshalingP(t, v, want)
}
