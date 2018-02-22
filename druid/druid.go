package druid

import "github.com/sul-dlss-labs/identifier-service/db"

// Druid is Digital Resource Unique ID
type Druid string

func (d Druid) String() string {
	return string(d)
}

// Persistable convert to a persistable representation
func (d Druid) Persistable() *db.Druid {
	return &db.Druid{ID: d.String()}
}
