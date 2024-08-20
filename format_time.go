package main

import "time"

func formatDate(dateString string) string {
	t, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		return dateString
	}
	return t.Format("02-01-2006")
}
