package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	app "github.com/sul-dlss-labs/id-proxy-api"
	"github.com/sul-dlss-labs/id-proxy-api/generated/models"
	"github.com/sul-dlss-labs/id-proxy-api/generated/restapi/operations"
)

// CreateDruid is the api endpoint for creating DRUIDs
type CreateDruid struct{}

// NewCreateDruid makes a new instance of CreateDruid
func NewCreateDruid(rt *app.Runtime) *CreateDruid {
	return &CreateDruid{}
}

// Handle the HTTP response to /identifiers/druids
func (d *CreateDruid) Handle(params operations.MintNewDRUIDSParams) middleware.Responder {
	id := "testing"
	identifier := &models.Identifier{ID: &id}
	identifiers := []*models.Identifier{identifier}
	return operations.NewMintNewDRUIDSOK().WithPayload(identifiers)
}
