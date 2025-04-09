package redirect

import (
	"fmt"
	"net/http"
	"testing"
)

type testConfigReader struct {
	includePath string
	trimPrefix  string
}

func (c *testConfigReader) Get(key string) string {
	switch key {
	case includePathKey:
		return c.includePath
	case trimPrefixKey:
		return c.trimPrefix
	default:
		return ""
	}
}

func TestWithPath(t *testing.T) {
	type test struct {
		name    string
		cfg     testConfigReader
		reqPath string
		wantURL string
	}

	tests := []test{
		{
			"include_path false",
			testConfigReader{},
			"/foo/bar",
			"http://localhost",
		},
		{
			"include_path true, trim_prefix empty",
			testConfigReader{includePath: "true"},
			"/foo/bar",
			"http://localhost/foo/bar",
		},
		{
			"include_path true, trim_prefix is foo",
			testConfigReader{includePath: "true", trimPrefix: "/foo"},
			"/foo/bar",
			"http://localhost/bar",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := &SpinRedirect{
				cfg: &tc.cfg,
			}

			dest := "http://localhost"
			req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost%s", tc.reqPath), nil)
			if err != nil {
				t.Fatalf("failed to create new http request: %s", err)
			}

			gotURL := r.WithPath(dest, req)
			if gotURL != tc.wantURL {
				t.Fatalf("failed: got '%s', want '%s'", gotURL, tc.wantURL)
			}
		})
	}
}
