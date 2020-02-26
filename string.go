package pointer

// String returns a pointer of given string
func String(input string) *string {
	return &input
}

// Int64 returns a pointer of given int64
func Int64(input int64) *int64 {
	return &input
}
