package logformat

import (
	"bytes"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"
)

const (
	testSize = 1000
)

func TestCANLogWriterReader(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testFile := "test.bin"
	w, err := NewSerializedLogWriter(testFile)
	if err != nil {
		t.Fatal(err)
	}
	var expected []CANFrame
	for i := 0; i < testSize; i++ {
		m := CANFrame64{
			TimestampSoftware: Timespec64{0, 0},
			TimestampHardware: Timespec64{0, 0},
			ID:                0,
			DLC:               8,
			Data:              []byte{0, 1, 2, 3, 4, 5, 6, 7},
		}
		// fill data with semi-random values
		m.ID = uint32(i)
		m.TimestampHardware.Sec = uint64(i)
		m.TimestampSoftware.Sec = uint64(rand.Int63())
		m.TimestampSoftware.Nsec = uint64(rand.Int63())
		rand.Read(m.Data)

		// save expected value
		expected = append(expected, *m.ToOld())

		// write to log
		err = w.Write(m.ToOld())
		if err != nil {
			t.Fatal(err)
		}
	}
	expHash, err := w.Close()
	if err != nil {
		t.Fatal(err)
	}

	r, err := NewSerializedLogReader(testFile)
	if err != nil {
		t.Fatal(err)
	}
	frames, gotHash, err := r.ReadAllCanFrames()
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
			t.Fatalf("error: messages don't match at index %d \nEXP %s != %s", i, expected[i].String(), frames[i].String())
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
