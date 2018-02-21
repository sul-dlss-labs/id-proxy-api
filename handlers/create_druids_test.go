package handlers

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/assert"
)

func TestCreateDruidsHappyPath(t *testing.T) {
	r := gofight.New()
	r.POST("/v1/identifiers/druids").
		Run(setupFakeRuntime().Handler(),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
				stat, _ := jsonparser.GetString(r.Body.Bytes(), "[0]")
				assert.Regexp(t, "^\\w{2}\\d{3}\\w{2}\\d{4}$", stat)
			})
}
