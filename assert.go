package goldsert

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"reflect"
	"testing"

	"github.com/jimeh/go-golden"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// Assert is the core of goldsert, holding both configuration and implementation
// of all assertion methods.
//
// You can customize serialization by setting any of the encoder/decoder
// function fields to a custom function that returns an encoder/decoder
// configured as you need.
//
// You can also customize golden file generation by setting the Golden field to
// a custom *golden.Golden instance. See the github.com/jimeh/go-golden package
// for details about what can be configured.
type Assert struct {
	JSONEncoderFunc func(io.Writer) *json.Encoder
	JSONDecoderFunc func(io.Reader) *json.Decoder
	YAMLEncoderFunc func(io.Writer) *yaml.Encoder
	YAMLDecoderFunc func(io.Reader) *yaml.Decoder
	XMLEncoderFunc  func(io.Writer) *xml.Encoder
	XMLDecoderFunc  func(io.Reader) *xml.Decoder
	Golden          *golden.Golden

	// NormalizeLineBreaks enables line-break normalization which replaces
	// Windows' CRLF (\r\n) and Mac Classic CR (\r) line breaks with Unix's LF
	// (\n) line breaks.
	NormalizeLineBreaks bool
}

// New returns a new *Assert instance configured with default settings.
//
// The default encoders all specify indentation of two spaces, essentially
// enforcing pretty formatting for JSON and XML.
//
// The default decoders for JSON and YAML prohibit unknown fields which are not
// present on the provided struct.
func New() *Assert {
	return &Assert{
		JSONEncoderFunc:     newJSONEncoder,
		JSONDecoderFunc:     newJSONDecoder,
		YAMLEncoderFunc:     newYAMLEncoder,
		YAMLDecoderFunc:     newYAMLDecoder,
		XMLEncoderFunc:      newXMLEncoder,
		XMLDecoderFunc:      newXMLDecoder,
		Golden:              golden.New(),
		NormalizeLineBreaks: true,
	}
}

// JSONMarshaling asserts that the given "v" value JSON marshals to an expected
// value fetched from a golden file on disk, and then verifies that the
// marshaled result produces a value that is equal to "v" when unmarshaled.
//
// Used for objects that do NOT change when they are marshaled and unmarshaled.
func (s *Assert) JSONMarshaling(t *testing.T, v interface{}) {
	t.Helper()

	s.JSONMarshalingP(t, v, v)
}

// JSONMarshalingP asserts that the given "v" value JSON marshals to an expected
// value fetched from a golden file on disk, and then verifies that the
// marshaled result produces a value that is equal to "want" when unmarshaled.
//
// Used for objects that change when they are marshaled and unmarshaled.
func (s *Assert) JSONMarshalingP(t *testing.T, v, want interface{}) {
	t.Helper()

	var buf bytes.Buffer
	err := s.JSONEncoderFunc(&buf).Encode(v)
	require.NoErrorf(t, err, "failed to JSON marshal %T: %+v", v, v)

	marshaled := buf.Bytes()
	if s.NormalizeLineBreaks {
		marshaled = normalizeLineBreaks(marshaled)
	}

	if s.Golden.Update() {
		s.Golden.SetP(t, "goldsert_json", marshaled)
	}

	gold := s.Golden.GetP(t, "goldsert_json")
	if s.NormalizeLineBreaks {
		gold = normalizeLineBreaks(gold)
	}
	assert.JSONEq(t, string(gold), string(marshaled))

	if reflect.ValueOf(want).Kind() != reflect.Ptr {
		require.FailNowf(t,
			"only pointer types can be asserted",
			"%T is not a pointer type", want,
		)
	}

	got := reflect.New(reflect.TypeOf(want).Elem()).Interface()
	err = s.JSONDecoderFunc(bytes.NewBuffer(gold)).Decode(got)
	require.NoErrorf(t, err,
		"failed to JSON unmarshal %T from %s",
		got, s.Golden.FileP(t, "goldsert_json"),
	)
	assert.Equal(t, want, got,
		"unmarshaling from golden file does not match expected object",
	)
}

// YAMLMarshaling asserts that the given "v" value YAML marshals to an expected
// value fetched from a golden file on disk, and then verifies that the
// marshaled result produces a value that is equal to "v" when unmarshaled.
//
// Used for objects that do NOT change when they are marshaled and unmarshaled.
func (s *Assert) YAMLMarshaling(t *testing.T, v interface{}) {
	t.Helper()

	s.YAMLMarshalingP(t, v, v)
}

// YAMLMarshalingP asserts that the given "v" value YAML marshals to an expected
// value fetched from a golden file on disk, and then verifies that the
// marshaled result produces a value that is equal to "want" when unmarshaled.
//
// Used for objects that change when they are marshaled and unmarshaled.
func (s *Assert) YAMLMarshalingP(t *testing.T, v, want interface{}) {
	t.Helper()

	var buf bytes.Buffer
	err := s.YAMLEncoderFunc(&buf).Encode(v)
	require.NoErrorf(t, err, "failed to YAML marshal %T: %+v", v, v)

	marshaled := buf.Bytes()
	if s.NormalizeLineBreaks {
		marshaled = normalizeLineBreaks(marshaled)
	}

	if s.Golden.Update() {
		s.Golden.SetP(t, "goldsert_yaml", marshaled)
	}

	gold := s.Golden.GetP(t, "goldsert_yaml")
	if s.NormalizeLineBreaks {
		gold = normalizeLineBreaks(gold)
	}
	assert.YAMLEq(t, string(gold), string(marshaled))

	if reflect.ValueOf(want).Kind() != reflect.Ptr {
		require.FailNowf(t,
			"only pointer types can be asserted",
			"%T is not a pointer type", want,
		)
	}

	got := reflect.New(reflect.TypeOf(want).Elem()).Interface()
	err = s.YAMLDecoderFunc(bytes.NewBuffer(gold)).Decode(got)
	require.NoErrorf(t, err,
		"failed to YAML unmarshal %T from %s",
		got, s.Golden.FileP(t, "goldsert_yaml"),
	)
	assert.Equal(t, want, got,
		"unmarshaling from golden file does not match expected object",
	)
}

// XMLMarshaling asserts that the given "v" value XML marshals to an expected
// value fetched from a golden file on disk, and then verifies that the
// marshaled result produces a value that is equal to "v" when unmarshaled.
//
// Used for objects that do NOT change when they are marshaled and unmarshaled.
func (s *Assert) XMLMarshaling(t *testing.T, v interface{}) {
	t.Helper()

	s.XMLMarshalingP(t, v, v)
}

// XMLMarshalingP asserts that the given "v" value XML marshals to an expected
// value fetched from a golden file on disk, and then verifies that the
// marshaled result produces a value that is equal to "want" when unmarshaled.
//
// Used for objects that change when they are marshaled and unmarshaled.
func (s *Assert) XMLMarshalingP(t *testing.T, v, want interface{}) {
	t.Helper()

	var buf bytes.Buffer
	err := s.XMLEncoderFunc(&buf).Encode(v)
	require.NoErrorf(t, err, "failed to XML marshal %T: %+v", v, v)

	marshaled := buf.Bytes()
	if s.NormalizeLineBreaks {
		marshaled = normalizeLineBreaks(marshaled)
	}

	if s.Golden.Update() {
		s.Golden.SetP(t, "goldsert_xml", marshaled)
	}

	gold := s.Golden.GetP(t, "goldsert_xml")
	if s.NormalizeLineBreaks {
		gold = normalizeLineBreaks(gold)
	}
	assert.Equal(t, string(gold), string(marshaled))

	if reflect.ValueOf(want).Kind() != reflect.Ptr {
		require.FailNowf(t,
			"only pointer types can be asserted",
			"%T is not a pointer type", want,
		)
	}

	got := reflect.New(reflect.TypeOf(want).Elem()).Interface()
	err = s.XMLDecoderFunc(bytes.NewBuffer(gold)).Decode(got)
	require.NoErrorf(t, err,
		"failed to XML unmarshal %T from %s",
		got, s.Golden.FileP(t, "goldsert_xml"),
	)
	assert.Equal(t, want, got,
		"unmarshaling from golden file does not match expected object",
	)
}

// newJSONEncoder is the default JSONEncoderFunc used by Assert. It returns a
// *json.Encoder which is set to indent with two spaces.
func newJSONEncoder(w io.Writer) *json.Encoder {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	return enc
}

// newJSONDecoder is the default JSONDecoderFunc used by Assert. It returns a
// *json.Decoder which disallows unknown fields.
func newJSONDecoder(r io.Reader) *json.Decoder {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec
}

// newYAMLEncoder is the default YAMLEncoderFunc used by Assert. It returns a
// *yaml.Encoder which is set to indent with two spaces.
func newYAMLEncoder(w io.Writer) *yaml.Encoder {
	enc := yaml.NewEncoder(w)
	enc.SetIndent(2)

	return enc
}

// newYAMLDecoder is the default YAMLDecoderFunc used by Assert. It returns a
// *yaml.Decoder which disallows unknown fields.
func newYAMLDecoder(r io.Reader) *yaml.Decoder {
	dec := yaml.NewDecoder(r)
	dec.KnownFields(true)

	return dec
}

// newXMLEncoder is the default XMLEncoderFunc used by Assert. It returns a
// *xml.Encoder which is set to indent with two spaces.
func newXMLEncoder(w io.Writer) *xml.Encoder {
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")

	return enc
}

// newXMLDecoder is the default XMLDecoderFunc used by Assert.
func newXMLDecoder(r io.Reader) *xml.Decoder {
	return xml.NewDecoder(r)
}

func normalizeLineBreaks(data []byte) []byte {
	// Replace CRLF (\r\n, windows) with LF (\n, unix)
	result := bytes.ReplaceAll(data, []byte{13, 10}, []byte{10})
	// Replace CR (\r, mac) with LF (\n, unix)
	result = bytes.ReplaceAll(result, []byte{13}, []byte{10})

	return result
}
