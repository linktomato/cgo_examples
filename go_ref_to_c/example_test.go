package go_ref_to_c

import "testing"
import "github.com/stretchr/testify/assert"

func TestCAdd(t *testing.T) {
	assert.Equal(t, 3, CAdd(1, 2))
}
