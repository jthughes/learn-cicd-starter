package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// func GetAPIKey(headers http.Header) (string, error) {
	tests := []struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		{
			expectErr: "no authorization header",
		},
		{
			key:       "Authorization",
			expectErr: "no authorization header",
		},
		{
			key:       "Authorization",
			value:     "blah",
			expectErr: "malformed authorization header",
		},
		{
			key:       "Authorization",
			value:     "Bearer xxxxxx",
			expectErr: "malformed authorization header",
		},
		{
			key:       "Authorization",
			value:     "ApiKey xxxxxx",
			expect:    "xxxxxx",
			expectErr: "not expecting an error",
		},
	}

	for ti, tc := range tests {
		t.Run((fmt.Sprintf("TestGetApIKey Test #%v: ", ti)), func(t *testing.T) {
			header := http.Header{}
			header.Add(tc.key, tc.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), tc.expectErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			}
			if output != tc.expect {
				t.Errorf("Unexpected: TestGetAPIKey:%s", output)
				return
			}
		})
	}

}
