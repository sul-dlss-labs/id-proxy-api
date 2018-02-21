package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	app "github.com/sul-dlss-labs/identifier-service"
	"github.com/sul-dlss-labs/identifier-service/druid"
	"github.com/sul-dlss-labs/identifier-service/generated/models"
	"github.com/sul-dlss-labs/identifier-service/generated/restapi/operations"
)

// CreateDruid is the api endpoint for creating DRUIDs
type CreateDruid struct{}

// NewCreateDruid makes a new instance of CreateDruid
func NewCreateDruid(rt *app.Runtime) *CreateDruid {
	return &CreateDruid{}
}

// Handle the HTTP response to /identifiers/druids
func (d *CreateDruid) Handle(params operations.MintNewDRUIDSParams) middleware.Responder {
	identifier := models.Identifier(druid.Generate())
	identifiers := []models.Identifier{identifier}
	return operations.NewMintNewDRUIDSOK().WithPayload(identifiers)
}
