package utils

import (
	"fmt"
	"strings"
	"time"
)

// ExtractDateFromLine extracts the date from a string like "From 1810898143257904491@xxx Sun Sep 22 12:10:48 +0000 2024".
func ExtractDateFromLine(line string) string {
	parts := strings.Split(line, " ")
	if len(parts) < 8 {
		return ""
	}

	// Construct the date string, including the timezone.
	dateStr := fmt.Sprintf("%s %s %s %s %s", parts[2], parts[3], parts[4], parts[5], strings.TrimSuffix(parts[7], "\r\n"))

	// Parse the date, handling the timezone.
	date, err := time.Parse("Mon Jan 02 15:04:05 2006", dateStr)
	if err != nil {
		fmt.Errorf("Error parsing date from line '%s': %w", line, err)
		return ""
	}

	// Return the date in "YYYY_MM_DD" format.
	return date.Format("2006_01_02")
}
