package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDateTime_TheDayIsFirstSundayOfMonth(t *testing.T) {
	v := NewDateTime("2023-12-03 10:00:00")
	assert.True(t, v.TheDayIsFirstSundayOfMonth())
	v = NewDateTime("2023-12-02 10:00:00")
	assert.False(t, v.TheDayIsFirstSundayOfMonth())
}
