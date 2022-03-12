package database

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"regexp"
)

func RemoveFirstChar(input string) string {
	if len(input) <= 1 {
		return ""
	}
	return input[1:]
}

func BoolToInt(input bool) int {
	if input {
		return 1
	}
	return 0
}

func RemoveSpecialChars(input string) string {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	return re.ReplaceAllString(input, " ")
}

func ContextError(ctx context.Context) error {

	switch ctx.Err() {
	case context.Canceled:
		log.Println("Request Canceled")
		return status.Error(codes.DeadlineExceeded, "Request Canceled")
	case context.DeadlineExceeded:
		log.Println("DeadLine Exceeded")
		return status.Error(codes.DeadlineExceeded, "DeadLine Exceeded")
	default:
		return nil
	}
}
