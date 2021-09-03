// +build e2e

package test

import (
	"testing"
	"github.com/go-resty/resty/v2"
)

func TestHealthEndpoint(t *testing.T) {
	client := resty.New()
}