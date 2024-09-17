package aipcli

import (
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"
	"gotest.tools/v3/assert"
)

func Test_qualifiedServiceUse(t *testing.T) {
	for _, tt := range []struct {
		name     string
		expected []string
	}{
		{
			name: "google.iam.v1.IAMPolicy",
			expected: []string{
				"iam-policy",
				"iam-policy-v1",
				"iam-policy-v1-iam",
				"iam-policy-v1-iam-google",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual := make([]string, 0, len(tt.expected))
			var i int
			for {
				result, ok := qualifiedServiceUse(protoreflect.FullName(tt.name), i)
				if !ok {
					break
				}
				actual = append(actual, result)
				i++
			}
			assert.DeepEqual(t, tt.expected, actual)
		})
	}
}
