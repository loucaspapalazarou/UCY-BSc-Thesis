package modules

import (
	"2-Atomic-Adds/config"
	"fmt"
	"strings"
)

func isMessageValid(msg string) bool {
	if msg == "" {
		return false
	}
	if strings.Contains(msg, " ") {
		return false
	}
	if strings.Contains(msg, ".") {
		return false
	}
	if strings.Contains(msg, "{") {
		return false
	}
	if strings.Contains(msg, "}") {
		return false
	}
	if strings.Contains(msg, ";") {
		return false
	}
	return true
}

func isAtomicMessageValid(msg string) bool {
	if msg == "" {
		return false
	}
	if strings.Contains(msg, " ") {
		return false
	}
	if strings.Contains(msg, ".") {
		return false
	}
	if strings.Contains(msg, "{") {
		return false
	}
	if strings.Contains(msg, "}") {
		return false
	}
	parts := strings.Split(msg, ";")
	if len(parts) != 4 {
		return false
	}
	for _, p := range parts {
		if len(p) < 1 {
			return false
		}
	}
	return config.NetworkExists(strings.Split(msg, ";")[1])
}

func truncateResponse(r string, maxCount int) string {
	if maxCount < 0 {
		return r
	}
	records := strings.Split(strings.TrimSpace(r), " ")
	count := len(records)
	if count <= maxCount {
		return r
	}
	truncatedRecords := records[:maxCount]
	omittedCount := count - maxCount
	return strings.Join(truncatedRecords, " ") + fmt.Sprintf(" and %d other records", omittedCount)
}
