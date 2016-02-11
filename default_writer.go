package log

import "log"

type DefaultWriter struct{}

func (dw DefaultWriter) Write(p []byte) (int, error) {
	log.Println(string(p))
	return len(p), nil
}
