package logformat

//go:generate msgp
//go:generate easyjson -all

const (
	BufferedMeasurementSize = 50
)

// check interface compatibility
var (
	_ Serializable = (*IVTMessage)(nil)
	_ Serializable = (*IvtBufferedMeasurement)(nil)
)

// IVTMessage is a message received from the IVT sensor, containing the power measurements
type IVTMessage struct {
	Timestamp   Timespec `json:"ts" msg:"ts"` // Timestamp of the message
	Voltage     int32    `json:"u" msg:"u"`   // Voltage in mV
	Current     int32    `json:"i" msg:"i"`   // Current in mA
	Temperature int32    `json:"t" msg:"t"`   // Temperature of the shunt in 1/10 °C
}

// IvtBufferedMeasurement is collection of multiple measurements from the IVT sensor
type IvtBufferedMeasurement struct {
	Voltages     [BufferedMeasurementSize]int32           `json:"u" msg:"u"` // Voltage in mV
	Currents     [BufferedMeasurementSize]int32           `json:"i" msg:"i"` // Current in mA
	Temperatures [(BufferedMeasurementSize / 8) + 2]int32 `json:"t" msg:"t"` // Temperature of the shunt in 1/10 °C
	SegmentStart Timespec                                 `json:"s" msg:"s"` // SegmentStart is the time the first message of this segment was recorded
	SegmentEnd   Timespec                                 `json:"e" msg:"e"` // SegmentEnd is the timestamp the last message of this sement was recorded
}
