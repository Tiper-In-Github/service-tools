package snowflakeid

import (
	"fmt"
	"testing"
	"time"
)

func TestNew_DefaultValues(t *testing.T) {
	sf := New()
	if sf.dataCenterId != 0 {
		t.Errorf("expected dataCenterId to be 0, got %d", sf.dataCenterId)
	}
	if sf.machineId != 1 {
		t.Errorf("expected machineId to be 6001, got %d", sf.machineId)
	}
	if sf.mode != Mode_Standard {
		t.Errorf("expected mode to be Mode_Standard, got %v", sf.mode)
	}
}

func TestSetDataCenterId(t *testing.T) {
	sf := New()

	err := sf.SetDataCenterId(31)
	if err != nil {
		t.Error("expected valid data center id 31 to be allowed")
	}

	err = sf.SetDataCenterId(32)
	if err == nil {
		t.Error("expected error when setting dataCenterId > maxDatacenterNum")
	}
}

func TestSetMachineId(t *testing.T) {
	sf := New()

	err := sf.SetMachineId(31)
	if err != nil {
		t.Error("expected valid machine id 31 to be allowed")
	}

	err = sf.SetMachineId(32)
	if err == nil {
		t.Error("expected error when setting machineId > maxMachineNum")
	}
}

func TestNextId_UniqueIDs(t *testing.T) {
	sf := New()
	sf.SetMode(Mode_Standard)

	ids := make(map[int64]bool)
	for i := 0; i < 1000; i++ {
		id, err := sf.NextId()
		if err != nil {
			t.Fatalf("failed to generate id: %v", err)
		}
		if ids[id] {
			t.Errorf("duplicate id generated: %d", id)
		}
		ids[id] = true
	}
	t.Log(ids)
}

func TestNextId_SequenceIncrement(t *testing.T) {
	sf := New()
	sf.lastTimestamp = time.Now().UnixNano() / 1e6

	// Generate same timestamp multiple times
	for i := int64(0); i < 4095; i++ {
		id, err := sf.NextId()
		if err != nil {
			t.Fatalf("failed to generate id: %v", err)
		}
		seq := id & maxSequenceNum
		expectedSeq := i + 1 // 因为第一次是 0 + 1
		if seq != expectedSeq {
			t.Errorf("expected sequence %d, got %d", expectedSeq, seq)
		}
	}
}

func TestNextId_TimeBackwards(t *testing.T) {
	sf := New()
	sf.lastTimestamp = time.Now().UnixNano()/1e6 + 1000 // 移动到未来

	// 模拟时间倒退
	// now := (time.Now().UnixNano() / 1e6) - 1000
	id, err := sf.NextId()
	if id != 0 || err == nil {
		t.Errorf("expected error when clock is moved backwards")
	}
}

func TestSetModeAndGenerate(t *testing.T) {
	sf := New()

	modes := []GenMode{Mode_Standard, Mode_IP, Mode_PseudoRandom}
	for _, mode := range modes {
		sf.SetMode(mode)
		id, err := sf.NextId()
		if err != nil {
			t.Errorf("mode %v failed to generate ID: %v", mode, err)
		}
		fmt.Printf("Mode %v generated ID: %d\n", mode, id)
	}
}
