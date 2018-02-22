package druid

import (
	"github.com/sul-dlss-labs/identifier-service/db"
)

// Minter is an interface for objects that mint identifiers
type Minter interface {
	Mint() Druid
}

// PersistingMinter is the object responsible for minting ids
type PersistingMinter struct {
	db db.Repository
}

// NewMinter create a new persting minter instance
func NewMinter(db db.Repository) Minter {
	return &PersistingMinter{db: db}
}

// Mint creates a never been issued random DRUID
func (m *PersistingMinter) Mint() Druid {
	candidate := Generate()
	for !m.isUnique(candidate) {
		candidate = Generate()
	}

	if err := m.store(candidate); err == nil {
		return candidate
	}
	// This condition arises if the database encounters a unique key constraint violation
	return m.Mint()
}

// Persist the druid in the datastore
func (m *PersistingMinter) store(d Druid) error {
	return m.db.CreateItem(d.Persistable())
}

// Check the datastore to see if this druid has already been issued
func (m *PersistingMinter) isUnique(d Druid) bool {
	exists, err := m.db.Exists(d.String())
	if err != nil {
		panic(err)
	}
	return !exists
}
