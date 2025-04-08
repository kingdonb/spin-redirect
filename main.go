package main

import (
	spinhttp "github.com/spinframework/spin-go-sdk/http"

	"github.com/fermyon/spin-redirect/redirect"
)

func init() {
	r := redirect.NewSpinRedirect()
	spinhttp.Handle(r.HandleFunc)
}

func main() {
}
