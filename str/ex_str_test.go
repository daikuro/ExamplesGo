package str

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestLength(t *testing.T) {
	assert.Equal(t, 2, len("AB"))
}

func TestCount(t *testing.T) {
	c := strings.Count("1.2.3", ".")
	assert.Equal(t, 2, c)
}