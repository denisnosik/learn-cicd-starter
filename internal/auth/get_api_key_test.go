package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headerKey string
		headerVal string
		wantVal   string
		wantErr   bool
	}{
		"valid simple": {headerKey: "Authorization", headerVal: "ApiKey 123456", wantVal: "123456", wantErr: false},
		"empty header": {headerKey: "", headerVal: "", wantVal: "", wantErr: true},
	}

	for name, tc := range tests {
		headers := http.Header{}
		headers.Set(tc.headerKey, tc.headerVal)
		got, err := GetAPIKey(headers)
		if err != nil && !tc.wantErr {
			t.Fatalf("%s: expected no err, but got one: %v", name, err)
		}
		if !reflect.DeepEqual(tc.wantVal, got) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.wantVal, got)
		}
	}
}
