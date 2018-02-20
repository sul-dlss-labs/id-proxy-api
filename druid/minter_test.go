package druid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDruid(t *testing.T) {
	assert.Regexp(t, "^\\w{2}\\d{3}\\w{2}\\d{4}$", Generate())
}
