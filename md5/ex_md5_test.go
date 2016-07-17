package md5

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBase64MD5(t *testing.T) {
	s := Base64MD5("a")
	assert.Equal(t, "DMF1ucDxtqgxw5niaXcmYQ==", s)
}
