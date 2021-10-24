package logformat

import (
	"fmt"
)

//go:generate msgp
//go:generate easyjson -all

// check interface compatibility
var (
	_ Serializable = (*Timespec)(nil)
	_ Serializable = (*CANFrame)(nil)
)

const (
	FrameLength = 16 // FrameLength is total length of the frame in the kernel.
)

// Timespec is the unix.Timespec struct with msgp tags
type Timespec struct {
	Sec  int32 `json:"s" msg:"s"`   // Seconds
	Nsec int32 `json:"ns" msg:"ns"` // Nanoseconds
}

// CANFrame is a CAN frame
type CANFrame struct {
	TimestampSoftware Timespec `json:"tss" msg:"tss"` // Software (kernel) timestamp of the frame
	TimestampHardware Timespec `json:"tsh" msg:"tsh"` // Hardware timestamp of the frame
	ID                uint32   `json:"id" msg:"id"`   // CAN ID
	DLC               uint8    `json:"l" msg:"l"`     // CAN DLC (data length code)
	Data              []byte   `json:"d" msg:"d"`     // CAN Data (0-8 bytes)
}

// String formats a CAN message in human-readable format
func (f *CANFrame) String() string {
	var dataString string
	for i := 0; i < int(f.DLC); i++ {
		dataString += fmt.Sprintf("%02X ", f.Data[i])
	}
	return fmt.Sprintf("%d.%d: %08X [%d] %s", f.TimestampHardware.Sec, f.TimestampHardware.Nsec, f.ID, f.DLC, dataString)
}
