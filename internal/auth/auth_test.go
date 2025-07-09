package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		key     string
		value   string
		want    string
		wantErr string
	}{
		{key: "", want: "", wantErr: "no authorization header"},
		{key: "Author", want: "", wantErr: "no authorization header"},
		{key: "Authorization", value: "key", want: "", wantErr: "malformed authorization header"},
		{key: "Authorization", value: "Key 123456", want: "", wantErr: "malformed authorization header"},
		{key: "Authorization", value: "ApiKey 123456", want: "123456", wantErr: "not expecting an error"},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("TestGetAPIKey Case #%v:", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(tc.key, tc.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), tc.wantErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			}

			if output != tc.want {
				t.Errorf("Unexpected: TestGetAPIKey:%s", output)
				return
			}
		})
	}
}
