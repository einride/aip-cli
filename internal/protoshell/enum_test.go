package protoshell

import (
	"testing"

	ltype "google.golang.org/genproto/googleapis/logging/type"
	"google.golang.org/protobuf/reflect/protoreflect"
	"gotest.tools/v3/assert"
)

func TestCompleteEnum(t *testing.T) {
	for _, tt := range []struct {
		name       string
		toComplete string
		enum       protoreflect.EnumValueDescriptors
		expected   []string
	}{
		{
			name:       "empty",
			toComplete: "",
			enum:       ltype.LogSeverity_DEFAULT.Descriptor().Values(),
			expected: []string{
				"DEFAULT", "DEBUG", "INFO", "NOTICE", "WARNING", "ERROR", "CRITICAL", "ALERT", "EMERGENCY",
			},
		},

		{
			name:       "prefix match",
			toComplete: "E",
			enum:       ltype.LogSeverity_DEFAULT.Descriptor().Values(),
			expected:   []string{"ERROR", "EMERGENCY"},
		},

		{
			name:       "no match",
			toComplete: "Z",
			enum:       ltype.LogSeverity_DEFAULT.Descriptor().Values(),
			expected:   nil,
		},

		{
			name:       "exact match",
			toComplete: "DEFAULT",
			enum:       ltype.LogSeverity_DEFAULT.Descriptor().Values(),
			expected:   []string{"DEFAULT"},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual := CompleteEnumValue(tt.toComplete, tt.enum)
			assert.DeepEqual(t, tt.expected, actual)
		})
	}
}
