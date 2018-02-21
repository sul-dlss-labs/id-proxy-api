package handlers

import (
	"net/http"
	"strings"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestIdentifiersInfo(t *testing.T) {
	r := gofight.New()
	r.GET("/v1/identifiers").
		Run(setupFakeRuntime().Handler(),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
				payload := strings.TrimSpace(r.Body.String())
				assert.Equal(t, `[{"name":"DRUID","template":"r:zznnnzznnnn"}]`, payload)
			})
}
