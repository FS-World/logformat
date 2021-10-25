package logformat

import (
	"bytes"
	"math/rand"
	"os"
	"reflect"
	"testing"
)

const (
	testSize = 1000
)

func TestCANLogWriterReader(t *testing.T) {
	testFile := "test.bin"
	w, err := NewSerializedLogWriter(testFile)
	if err != nil {
		t.Fatal(err)
	}
	m := CANFrame{
		TimestampSoftware: Timespec{0, 0},
		TimestampHardware: Timespec{0, 0},
		ID:                0,
		DLC:               8,
		Data:              []byte{0, 1, 2, 3, 4, 5, 6, 7},
	}
	expected := make([]Serializable, testSize)
	for i := 0; i < testSize; i++ {
		// fill data with semi-random values
		m.ID = uint32(i)
		m.TimestampHardware.Sec = int32(i)
		m.TimestampSoftware.Sec = rand.Int31()
		m.TimestampSoftware.Nsec = rand.Int31()
		rand.Read(m.Data)

		// save expected value
		expected[i] = &m

		// write to log
		err = w.Write(&m)
		if err != nil {
			t.Fatal(err)
		}
	}
	expHash, err := w.Close()
	if err != nil {
		t.Fatal(err)
	}

	r, err := NewSerializedLogReader(testFile, &CANFrame{})
	if err != nil {
		t.Fatal(err)
	}
	frames, gotHash, err := r.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(frames) != testSize {
		t.Fatal("error reading messages")
	}
	if !bytes.Equal(expHash, gotHash) {
		t.Fatalf("error: hashes don't match\nEXP %v != %v", expHash, gotHash)
	}
	for i := 0; i < testSize; i++ {
		if !reflect.DeepEqual(frames[i], expected[i]) {
			t.Fatalf("error: messages don't match\nEXP %v != %v", expected[i], frames[i])
		}
	}
	err = r.Close()
	if err != nil {
		t.Fatal(err)
	}
	err = os.Remove(testFile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIVTLogWriterReader(t *testing.T) {
	testFile := "test_ivt.bin"
	w, err := NewSerializedLogWriter(testFile)
	if err != nil {
		t.Fatal(err)
	}
	var m IVTMessage
	expected := make([]Serializable, testSize)
	for i := 0; i < testSize; i++ {
		// fill data with semi-random values
		m.Timestamp.Sec = int32(i)
		m.Voltage = rand.Int31()
		m.Current = rand.Int31()
		m.Temperature = rand.Int31()

		// save expected data
		expected[i] = &m

		// write to log
		err = w.Write(&m)
		if err != nil {
			t.Fatal(err)
		}
	}
	expHash, err := w.Close()
	if err != nil {
		t.Fatal(err)
	}

	r, err := NewSerializedLogReader(testFile, &m)
	if err != nil {
		t.Fatal(err)
	}
	frames, gotHash, err := r.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(frames) != testSize {
		t.Fatal("error reading messages")
	}
	if !bytes.Equal(expHash, gotHash) {
		t.Fatalf("error: hashes don't match\nEXP %v != %v", expHash, gotHash)
	}
	for i := 0; i < testSize; i++ {
		if !reflect.DeepEqual(frames[i], expected[i]) {
			t.Fatalf("error: messages don't match\nEXP %v != %v", expected[i], frames[i])
		}
	}
	err = r.Close()
	if err != nil {
		t.Fatal(err)
	}
	err = os.Remove(testFile)
	if err != nil {
		t.Fatal(err)
	}
}
