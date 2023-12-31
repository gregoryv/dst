package dst

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"
)

// SetContent updates dst with the file content of value.
func SetContent(dst *[]byte, value string) error {
	tmp, err := os.ReadFile(value)
	if err != nil {
		return fmt.Errorf("SetContent: %w", err)
	}
	*dst = tmp
	return nil
}

// SetFloat updates dst if the value is a valid float64
func SetFloat(dst *float64, value string) error {
	tmp, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("SetFloat: %w", err)
	}
	*dst = tmp
	return nil
}

// SetInt updates dst if the value is a valid int
func SetInt(dst *int, value string) error {
	tmp, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("SetInt: %w", err)
	}
	*dst = tmp
	return nil
}

// SetURL updates dst if the value is a valid url.URL
func SetURL(dst *url.URL, value string) error {
	tmp, err := url.Parse(value)
	if err != nil {
		return fmt.Errorf("SetURL: %w", err)
	}
	*dst = *tmp
	return nil
}

// SetDuration updates dst if the value is a valid time.Duration
func SetDuration(dst *time.Duration, value string) error {
	tmp, err := time.ParseDuration(value)
	if err != nil {
		return fmt.Errorf("SetDuration: %w", err)
	}
	*dst = tmp
	return nil
}

// SetUint32 updates dst if the value is a valid uint32
func SetUint32(dst *uint32, value string) error {
	tmp, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return fmt.Errorf("SetUint32: %w", err)
	}
	*dst = uint32(tmp)
	return nil
}

// SetBool updates dst if the value is a valid boolean string, empty
// value means true.
func SetBool(dst *bool, value string) error {
	if value == "" {
		*dst = true
		return nil
	}
	tmp, err := strconv.ParseBool(value)
	if err != nil {
		return fmt.Errorf("SetBool: %w", err)
	}
	*dst = tmp
	return nil
}
