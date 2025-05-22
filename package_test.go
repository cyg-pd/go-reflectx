package reflectx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageName(t *testing.T) {
	assert.Equal(t, "github.com/cyg-pd/go-reflectx", PackageName())
	assert.Equal(t, "github.com/cyg-pd/go-reflectx", PackageName(0))
	assert.Equal(t, "github.com/cyg-pd/go-reflectx", PackageName(1))
	assert.Equal(t, "testing", PackageName(2))
	assert.Equal(t, "runtime", PackageName(3))
	assert.Equal(t, "", PackageName(99999))
}
