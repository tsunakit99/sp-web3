package task

import (
	"errors"
	"strings"
)

const (
	MaxTitleLength       = 100
	MaxDescriptionLength = 1000
)

func ValidateTaskAndDescription(title, description string) error {
	if strings.TrimSpace(title) == "" {
		return errors.New("title cannot be empty")
	}
	if len(title) > MaxTitleLength {
		return errors.New("title exceeds maximum length")
	}
	if strings.TrimSpace(description) == "" {
		return errors.New("description cannot be empty")
	}
	if len(description) > MaxDescriptionLength {
		return errors.New("description exceeds maximum length")
	}
	return nil
}
