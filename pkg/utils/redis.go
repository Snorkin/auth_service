package utils

import "fmt"

func CreateKey(prefix string, sessionId string) string {
	return fmt.Sprintf("%s: %s", prefix, sessionId)
}
