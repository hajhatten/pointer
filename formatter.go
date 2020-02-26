package pointer

import "fmt"

// StringPtr formats a string pointer to either "<nil>" or the actual value
func StringPtr(input *string) string {
	if input == nil {
		return "<nil>"
	}
	return *input
}

// Int64Ptr formats a int64 pointer to either "<nil>" or the actual value
func Int64Ptr(input *int64) string {
	if input == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%d", *input)
}
