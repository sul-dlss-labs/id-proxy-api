package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	app "github.com/sul-dlss-labs/identifier-service"
	"github.com/sul-dlss-labs/identifier-service/generated/models"
	"github.com/sul-dlss-labs/identifier-service/generated/restapi/operations"
)

// IdentifiersInfo is the api endpoint for the types of identifiers
type IdentifiersInfo struct{}

// NewIdentifiersInfo makes a new instance of IdentifiersInfo
func NewIdentifiersInfo(rt *app.Runtime) *IdentifiersInfo {
	return &IdentifiersInfo{}
}

// Handle the HTTP response to GET /identifiers
func (d *IdentifiersInfo) Handle(params operations.GetIdentifiersInfoParams) middleware.Responder {
	druid := "DRUID"
	template := "r:zznnnzznnnn"
	payload := models.Sources{&models.Source{Name: &druid, Template: &template}}

	return operations.NewGetIdentifiersInfoOK().WithPayload(payload)
}
