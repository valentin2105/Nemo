package global

import "os"

const DefaultDateFormat = "Mon Jan _2 2006 - 15:04"

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func TrimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func ChooseStatusFaIcon(s string) string {
	if s == "Running" {
		e := "fa-circle success"
		return e
	} else if s == "Succeeded" {
		e := "fa-circle"
		return e
	} else if s == "Failed" {
		e := "fa-circle"
		return e
	} else if s == "Pending" {
		e := "fa-circle"
		return e
	} else if s == "Error" {
		e := "fa-circle"
		return e
	} else if s == "CrashLoopBackOff" {
		e := "fa-circle failed"
		return e
	} else {
		e := "fa-circle"
		return e
	}
}
