package handlers

import (
	"net/http"

	app "github.com/sul-dlss-labs/identifier-service"
	"github.com/sul-dlss-labs/identifier-service/druid"
)

func setupFakeRuntime() *TestEnv {
	return &TestEnv{}
}

type TestEnv struct {
}

func (d *TestEnv) Handler() http.Handler {
	// For testing we can inject a fake config/services here.
	rt, _ := app.NewRuntime(nil)
	rt = rt.WithMinter(mockMinter())
	return BuildAPI(rt).Serve(nil)
}

func mockMinter() druid.Minter {
	return &fakeMinter{}
}

type fakeMinter struct {
}

func (d fakeMinter) Mint() druid.Druid {
	return druid.Generate()
}
