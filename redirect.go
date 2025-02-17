package main

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

const (
	// Default value for HTTP status code
	DefaultStatusCode int = http.StatusFound
	// Default value for redirection target
	DefaultRedirectionTarget string = "/"
	// Key for loading desired destination
	destinationKey string = "destination"
	// Key for loading desired HTTP status code
	statusCodeKey string = "statuscode"
	// Key to enable adding original path in redirect
	includePath = "include_path"
)

func init() {
	r := NewSpinRedirect()
	spinhttp.Handle(r.handleFunc)
}

func main() {
}

// SpinRedirect is a struct that provides a handleFunc
// for redirecting to a destination URL using configurable HTTP status code.
type SpinRedirect struct {
	cfg ConfigReader
}

// NewSpinRedirect returns a new SpinRedirect
func NewSpinRedirect() SpinRedirect {
	return SpinRedirect{
		cfg: NewDefaultConfigReader(),
	}
}

func (s SpinRedirect) handleFunc(w http.ResponseWriter, r *http.Request) {
	dest, _ := s.getDestination()
	code, _ := s.getStatusCode(r.Method)

	w.Header().Set("Location", s.WithPath(dest, r))
	w.WriteHeader(code)
}

// getDestination returns the destination URL and a boolean indicating if the user provided a destination URL.
// If no destination is found, DefaultRedirectionTarget (/) is returned.
func (s SpinRedirect) getDestination() (string, bool) {
	d := s.cfg.Get(destinationKey)
	if len(d) == 0 {
		return DefaultRedirectionTarget, false
	}
	return d, true
}

func (s SpinRedirect) WithPath(dest string, r *http.Request) string {
	d := s.cfg.Get(includePath)
	if d != "true" {
		return dest
	}

	u, err := url.Parse(dest)
	if err != nil {
		return dest
	}

	// Combine paths if request has a non-empty path
	if r.URL.Path != "" && r.URL.Path != "/" {
		// Ensure we don't double-slash when combining paths
		u.Path = strings.TrimSuffix(u.Path, "/") + "/" + strings.TrimPrefix(r.URL.Path, "/")
	}

	return u.String()
}

// getStatusCode returns the HTTP status code
// If no status code is found, or if the provided value is invalid,
// DefaultStatusCode is returned along with a boolean indicating if the user provided a valid status code.
func (s SpinRedirect) getStatusCode(method string) (int, bool) {
	str := s.cfg.Get(statusCodeKey)
	code, err := strconv.Atoi(str)
	if err != nil {
		return DefaultStatusCode, false
	}
	if !isValidRedirectStatusCode(code, method) {
		return DefaultStatusCode, false
	}
	return code, true
}

// isValidRedirectStatusCode returns true if the provided status code is valid for redirection
func isValidRedirectStatusCode(code int, method string) bool {
	if code == http.StatusSeeOther &&
		(method == http.MethodPut || method == http.MethodPost) {
		return true
	}
	validCodes := []int{
		http.StatusMovedPermanently,
		http.StatusFound,
		http.StatusTemporaryRedirect,
		http.StatusPermanentRedirect,
	}

	for _, c := range validCodes {
		if c == code {
			return true
		}
	}
	return false
}
