package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/sul-dlss-labs/identifier-service/config"
	baloo "gopkg.in/h2non/baloo.v3"
)

func setupTest() *baloo.Client {
	remoteHost, ok := os.LookupEnv("TEST_REMOTE_ENDPOINT")
	if !ok {
		port := config.NewConfig().Port
		remoteHost = fmt.Sprintf("localhost:%v", port)
	}
	return baloo.New(fmt.Sprintf("http://%s", remoteHost))
}

func TestCreateDruid(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	const schema = `{ "type": "array" }`
	setupTest().Post("/v1/identifiers/druids").
		Body(nil). // https://github.com/h2non/baloo/issues/21, https://github.com/go-swagger/go-swagger/issues/1430
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(schema).
		Done()
}

func TestIdentifiers(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	setupTest().Get("/v1/identifiers").
		Expect(t).
		Status(200).
		Type("json").
		Done()
}
