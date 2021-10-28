<h1 align="center">
  go-goldsert
</h1>

<p align="center">
  <strong>
    A suite of Go test helpers which uses golden files to assert marshaling and unmarshaling of given objects.
  </strong>
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/jimeh/go-goldsert">
    <img src="https://img.shields.io/badge/%E2%80%8B-reference-387b97.svg?logo=go&logoColor=white"
  alt="Go Reference">
  </a>
  <a href="https://github.com/jimeh/go-goldsert/actions">
    <img src="https://img.shields.io/github/workflow/status/jimeh/go-goldsert/CI.svg?logo=github" alt="Actions Status">
  </a>
  <a href="https://codeclimate.com/github/jimeh/go-goldsert">
    <img src="https://img.shields.io/codeclimate/coverage/jimeh/go-goldsert.svg?logo=code%20climate" alt="Coverage">
  </a>
  <a href="https://github.com/jimeh/go-goldsert/issues">
    <img src="https://img.shields.io/github/issues-raw/jimeh/go-goldsert.svg?style=flat&logo=github&logoColor=white"
alt="GitHub issues">
  </a>
  <a href="https://github.com/jimeh/go-goldsert/pulls">
    <img src="https://img.shields.io/github/issues-pr-raw/jimeh/go-goldsert.svg?style=flat&logo=github&logoColor=white" alt="GitHub pull requests">
  </a>
  <a href="https://github.com/jimeh/go-goldsert/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/jimeh/go-goldsert.svg?style=flat" alt="License Status">
  </a>
</p>

Each test helper operates in two stages:

1. Marshal the provided object to a byte slice and reads the corresponding
   golden file from disk, follow by verifying both byte slices are identical.
2. Unmarshal the content from the golden file and verify that the result is
   identical to the original object.

## Import

```go
import "github.com/jimeh/go-goldsert"
```

## Usage

Typical usage should look something like this in a tabular test:

```go
type MyStruct struct {
    FooBar string `json:"foo_bar" yaml:"fooBar" xml:"Foo_Bar"`
}

func TestMyStructMarshaling(t *testing.T) {
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
```

The above example will read from the following golden files:

- `testdata/TestMyStructMarshaling/empty/goldsert_json.golden`
- `testdata/TestMyStructMarshaling/empty/goldsert_yaml.golden`
- `testdata/TestMyStructMarshaling/empty/goldsert_xml.golden`
- `testdata/TestMyStructMarshaling/full/goldsert_json.golden`
- `testdata/TestMyStructMarshaling/full/goldsert_yaml.golden`
- `testdata/TestMyStructMarshaling/full/goldsert_xml.golden`

If a corresponding golden file cannot be found on disk, the test will fail. To
create/update golden files, simply set the `GOLDEN_UPDATE` environment variable
to one of `1`, `y`, `t`, `yes`, `on`, or `true` when running tests.

It is highly recommended that golden files are committed to source control, as
it allow tests to fail when the marshal results for an object changes.

## Documentation

Please see the
[Go Reference](https://pkg.go.dev/github.com/jimeh/go-goldsert#section-documentation)
for documentation and examples.

## License

[MIT](https://github.com/jimeh/go-goldsert/blob/main/LICENSE)
