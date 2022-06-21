package logformat

//go:generate msgp
//go:generate easyjson -all

const (
	BufferedMeasurementSize    = 50
	BufferedMeasurementSizeIvt = 8 // BufferedMeasurementSizeIvt must be at least (BufferedMeasurementSize / 8) + 1
)

// check interface compatibility
var (
//_ Serializable = (*IvtBufferedMeasurement)(nil)
)

// IvtBufferedMeasurement is collection of multiple measurements from the IVT sensor
type IvtBufferedMeasurement struct {
	Voltages     [BufferedMeasurementSize]int32    `json:"u" msg:"u"` // Voltage in mV
	Currents     [BufferedMeasurementSize]int32    `json:"i" msg:"i"` // Current in mA
	Temperatures [BufferedMeasurementSizeIvt]int32 `json:"t" msg:"t"` // Temperature of the shunt in 1/10 °C
	SegmentStart Timespec                          `json:"s" msg:"s"` // SegmentStart is the time the first message of this segment was recorded
	SegmentEnd   Timespec                          `json:"e" msg:"e"` // SegmentEnd is the timestamp the last message of this sement was recorded
}
