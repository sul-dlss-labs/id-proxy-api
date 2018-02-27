package handlers

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/assert"
)

func TestCreateDruidHappyPath(t *testing.T) {
	r := gofight.New()
	r.POST("/v1/identifiers/druids").
		Run(setupFakeRuntime().Handler(),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
				stat, _ := jsonparser.GetString(r.Body.Bytes(), "[0]")
				assert.Regexp(t, "^\\w{2}\\d{3}\\w{2}\\d{4}$", stat)
			})
}

func TestCreateMultipleDruidsHappyPath(t *testing.T) {
	r := gofight.New()
	r.POST("/v1/identifiers/druids?quantity=5").
		Run(setupFakeRuntime().Handler(),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
				first, _ := jsonparser.GetString(r.Body.Bytes(), "[0]")
				assert.Regexp(t, "^\\w{2}\\d{3}\\w{2}\\d{4}$", first)
				second, _ := jsonparser.GetString(r.Body.Bytes(), "[1]")
				assert.Regexp(t, "^\\w{2}\\d{3}\\w{2}\\d{4}$", second)
			})
}

func TestCreateDruidsZero(t *testing.T) {
	r := gofight.New()
	r.POST("/v1/identifiers/druids?quantity=0").
		Run(setupFakeRuntime().Handler(),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
				assert.Equal(t, "[]\n", r.Body.String())
			})
}

func TestCreateDruidsInvalidQuantiy(t *testing.T) {
	r := gofight.New()
	r.POST("/v1/identifiers/druids?quantity=-1").
		Run(setupFakeRuntime().Handler(),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
				assert.Equal(t, "[]\n", r.Body.String())
			})
}

func TestCreateDruidsTooMany(t *testing.T) {
	r := gofight.New()
	r.POST("/v1/identifiers/druids?quantity=10000").
		Run(setupFakeRuntime().Handler(),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
				count := 0
				jsonparser.ArrayEach(r.Body.Bytes(), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
					count++
				})
				// Only return 1000
				assert.Equal(t, int(MaxQuantity), count)
			})
}
