package services

import (
	"testing"

	"github.com/Reskill-2022/hoarder/config"
)

func TestUUIDFromURI(t *testing.T) {
	service := NewCalendlyService(config.New())

	testCases := []struct {
		name     string
		uri      string
		expected string
	}{
		{
			name:     "valid uri",
			uri:      "https://calendly.com/hoarder/30min/12345678-1234-1234-1234-123456789012",
			expected: "12345678-1234-1234-1234-123456789012",
		},
		{
			name:     "invalid uri",
			uri:      "https://calendly.com/hoarder/30min/12345678-1234-1234-1234-123456789012/",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := service.UUIDFromURI(tc.uri)

			if actual != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, actual)
			}
		})
	}
}
