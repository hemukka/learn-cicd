package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAuth(t *testing.T) {
	tests := []struct {
		input   http.Header
		want    string
		wantErr bool
	}{
		{input: http.Header{}, want: "", wantErr: true},
		{input: http.Header{"Authorization": []string{"ApiKey 123456"}}, want: "123456", wantErr: false},
		{input: http.Header{"Authorization": []string{"ApiKey"}}, want: "", wantErr: true},
		{input: http.Header{"Authorization": []string{"apikey 123456"}}, want: "", wantErr: true},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
		if tc.wantErr && err == nil {
			t.Fatalf("expected non-nil error, got nil error")
		}
	}
}
