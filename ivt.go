package logformat

//go:generate msgp
//go:generate easyjson -all

// check interface compatibility
var (
	_ Serializable = (*IVTMessage)(nil)
)

// IVTMessage is a message received from the IVT sensor, containing the power measurements
type IVTMessage struct {
	Timestamp   Timespec `json:"ts" msg:"ts"` // Timestamp of the message
	Voltage     int32    `json:"u" msg:"u"`   // Voltage in mV
	Current     int32    `json:"i" msg:"i"`   // Current in mA
	Temperature int32    `json:"t" msg:"t"`   // Temperature of the shunt in 1/10 °C
}
