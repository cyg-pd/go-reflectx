package reflectx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mock struct{}

func TestName(t *testing.T) {
	assert.Equal(t, "mock", Name(mock{}))
	assert.Equal(t, "mock", Name(&mock{}))
}

func TestFullName(t *testing.T) {
	assert.Equal(t, "github.com/cyg-pd/go-reflectx.mock", FullName(mock{}))
	assert.Equal(t, "github.com/cyg-pd/go-reflectx.mock", FullName(&mock{}))
}
