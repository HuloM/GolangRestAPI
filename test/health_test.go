// +build e2e

package test

import (
	"testing"
	"github.com/go-resty/resty/v2 v2.4.0"
)

func TestHealthEndpoint(t *testing.T) {
	client := resty.New()
}