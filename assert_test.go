package goldsert

import (
	"testing"
)

func TestAssert_JSONMarshaling(t *testing.T) {
	for _, tt := range marhalingTestCases {
		t.Run(tt.name, func(t *testing.T) {
			gs := New()

			gs.JSONMarshaling(t, tt.v)
		})
	}
}

func TestAssert_YAMLMarshaling(t *testing.T) {
	for _, tt := range marhalingTestCases {
		t.Run(tt.name, func(t *testing.T) {
			gs := New()

			gs.YAMLMarshaling(t, tt.v)
		})
	}
}

func TestAssert_XMLMarshaling(t *testing.T) {
	for _, tt := range marhalingTestCases {
		t.Run(tt.name, func(t *testing.T) {
			gs := New()

			gs.XMLMarshaling(t, tt.v)
		})
	}
}

func TestAssert_JSONMarshalingP(t *testing.T) {
	for _, tt := range marshalingPTestCases {
		t.Run(tt.name, func(t *testing.T) {
			gs := New()

			gs.JSONMarshalingP(t, tt.v, tt.want)
		})
	}
}

func TestAssert_YAMLMarshalingP(t *testing.T) {
	for _, tt := range marshalingPTestCases {
		t.Run(tt.name, func(t *testing.T) {
			gs := New()

			gs.YAMLMarshalingP(t, tt.v, tt.want)
		})
	}
}

func TestAssert_XMLMarshalingP(t *testing.T) {
	for _, tt := range marshalingPTestCases {
		t.Run(tt.name, func(t *testing.T) {
			gs := New()

			gs.XMLMarshalingP(t, tt.v, tt.want)
		})
	}
}
