package types

import (
	"testing"
	"time"

	"github.com/dromara/carbon/v2"
	"github.com/stretchr/testify/assert"
)

func TestDateTime_TheDayIsFirstSundayOfMonth(t *testing.T) {
	v := NewDateTime("2023-12-03 10:00:00")
	assert.True(t, v.TheDayIsFirstSundayOfMonth())
	v = NewDateTime("2023-12-02 10:00:00")
	assert.False(t, v.TheDayIsFirstSundayOfMonth())
}

func TestDateTime_FromStdTime(t *testing.T) {
	n := time.Now()
	dateTime := &DateTime{Carbon: carbon.CreateFromStdTime(n)}
	t.Log(dateTime)

	newDateTime := NewDateTime(n.Format(time.RFC3339))
	t.Log(newDateTime)
}
