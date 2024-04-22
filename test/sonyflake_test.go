package test

import (
	"testing"

	"github.com/sony/sonyflake"
)

func TestSonyflake(t *testing.T) {
	var st sonyflake.Settings
	st.MachineID = func() (uint16, error) {
		return snowflakeMachineID(), nil
	}
	st.CheckMachineID = func(u uint16) bool {
		return u == snowflakeMachineID()
	}

	s, err := sonyflake.New(st)
	if err != nil {
		t.Error(err)
		return
	}

	for range 10 {
		t.Log(s.NextID())
	}
}
