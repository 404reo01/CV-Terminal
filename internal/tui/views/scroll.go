package views

import "strings"

func ApplyScroll(content string, height, offset int) string {
	lines := strings.Split(content, "\n")
	total := len(lines)

	if offset < 0 {
		offset = 0
	}
	if offset >= total {
		offset = total - 1
	}

	end := offset + height
	if end > total {
		end = total
	}

	return strings.Join(lines[offset:end], "\n")
}

func MaxScroll(content string, height int) int {
	lines := strings.Split(content, "\n")
	max := len(lines) - height
	if max < 0 {
		return 0
	}
	return max
}
