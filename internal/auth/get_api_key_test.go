package auth

import (
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  error
	}{
		"simple": {input: map[string][]string{"Authorization": []string{"ApiKey mockapikey"}}, want: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, got := GetAPIKey(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
