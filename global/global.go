package global

import "os"

// DefaultDateFormat - Date default format
const DefaultDateFormat = "2006-01-02 15:04"

// GetEnv - return env w/ fallback string
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// TrimQuotes - get of quotes of string
func TrimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}
