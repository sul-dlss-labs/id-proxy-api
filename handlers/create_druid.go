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

// MaxQuantity is the maximum number of DRUIDs you can get per request.
// This prevents denial of service by requesting a huge number
const MaxQuantity int64 = 1000

// NewCreateDruid makes a new instance of CreateDruid
func NewCreateDruid(rt *app.Runtime) *CreateDruid {
	return &CreateDruid{}
}

// Handle the HTTP response to /identifiers/druids
func (d *CreateDruid) Handle(params operations.MintNewDRUIDSParams) middleware.Responder {
	identifiers := []models.Identifier{}
	quantity := *params.Quantity
	if quantity > MaxQuantity {
		quantity = MaxQuantity
	}
	for i := 0; i < int(quantity); i++ {
		identifiers = append(identifiers, models.Identifier(druid.Generate()))
	}
	return operations.NewMintNewDRUIDSOK().WithPayload(identifiers)
}
