package goldsert_test

import (
	"testing"

	"github.com/jimeh/go-goldsert"
)

type MyStruct struct {
	FooBar string `json:"foo_bar,omitempty" yaml:"fooBar,omitempty" xml:"Foo_Bar,omitempty"`
	Bar    string `json:"-" yaml:"-" xml:"-"`
	baz    string
}

// TestExampleMyStructMarshaling reads/writes the following golden files:
//
//  testdata/TestExampleMyStructMarshaling/goldsert_json.golden
//  testdata/TestExampleMyStructMarshaling/goldsert_yaml.golden
//  testdata/TestExampleMyStructMarshaling/goldsert_xml.golden
//
func TestExampleMyStructMarshaling(t *testing.T) {
	myStruct := &MyStruct{FooBar: "Hello World!"}

	goldsert.JSONMarshaling(t, myStruct)
	goldsert.YAMLMarshaling(t, myStruct)
	goldsert.XMLMarshaling(t, myStruct)
}

// TestExampleMyStructMarshalingP reads/writes the following golden files:
//
//  testdata/TestExampleMyStructMarshalingP/goldsert_json.golden
//  testdata/TestExampleMyStructMarshalingP/goldsert_yaml.golden
//  testdata/TestExampleMyStructMarshalingP/goldsert_xml.golden
//
func TestExampleMyStructMarshalingP(t *testing.T) {
	myStruct := &MyStruct{FooBar: "Hello World!", Bar: "Oops", baz: "nope!"}
	want := &MyStruct{FooBar: "Hello World!"}

	goldsert.JSONMarshalingP(t, myStruct, want)
	goldsert.YAMLMarshalingP(t, myStruct, want)
	goldsert.XMLMarshalingP(t, myStruct, want)
}

// TestExampleMyStructMarshalingTabular reads/writes the following golden files:
//
//  testdata/TestExampleMyStructMarshalingTabular/empty/goldsert_json.golden
//  testdata/TestExampleMyStructMarshalingTabular/empty/goldsert_yaml.golden
//  testdata/TestExampleMyStructMarshalingTabular/empty/goldsert_xml.golden
//  testdata/TestExampleMyStructMarshalingTabular/full/goldsert_json.golden
//  testdata/TestExampleMyStructMarshalingTabular/full/goldsert_yaml.golden
//  testdata/TestExampleMyStructMarshalingTabular/full/goldsert_xml.golden
//
func TestExampleMyStructMarshalingTabular(t *testing.T) {
	tests := []struct {
		name string
		obj  *MyStruct
	}{
		{name: "empty", obj: &MyStruct{}},
		{name: "full", obj: &MyStruct{FooBar: "Hello World!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			goldsert.JSONMarshaling(t, tt.obj)
			goldsert.YAMLMarshaling(t, tt.obj)
			goldsert.XMLMarshaling(t, tt.obj)
		})
	}
}
