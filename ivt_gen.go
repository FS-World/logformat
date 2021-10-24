package logformat

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *IVTMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ts":
			err = z.Timestamp.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Timestamp")
				return
			}
		case "u":
			z.Voltage, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "Voltage")
				return
			}
		case "i":
			z.Current, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "Current")
				return
			}
		case "t":
			z.Temperature, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "Temperature")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *IVTMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "ts"
	err = en.Append(0x84, 0xa2, 0x74, 0x73)
	if err != nil {
		return
	}
	err = z.Timestamp.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Timestamp")
		return
	}
	// write "u"
	err = en.Append(0xa1, 0x75)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.Voltage)
	if err != nil {
		err = msgp.WrapError(err, "Voltage")
		return
	}
	// write "i"
	err = en.Append(0xa1, 0x69)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.Current)
	if err != nil {
		err = msgp.WrapError(err, "Current")
		return
	}
	// write "t"
	err = en.Append(0xa1, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.Temperature)
	if err != nil {
		err = msgp.WrapError(err, "Temperature")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *IVTMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "ts"
	o = append(o, 0x84, 0xa2, 0x74, 0x73)
	o, err = z.Timestamp.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Timestamp")
		return
	}
	// string "u"
	o = append(o, 0xa1, 0x75)
	o = msgp.AppendInt32(o, z.Voltage)
	// string "i"
	o = append(o, 0xa1, 0x69)
	o = msgp.AppendInt32(o, z.Current)
	// string "t"
	o = append(o, 0xa1, 0x74)
	o = msgp.AppendInt32(o, z.Temperature)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *IVTMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ts":
			bts, err = z.Timestamp.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Timestamp")
				return
			}
		case "u":
			z.Voltage, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Voltage")
				return
			}
		case "i":
			z.Current, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Current")
				return
			}
		case "t":
			z.Temperature, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Temperature")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *IVTMessage) Msgsize() (s int) {
	s = 1 + 3 + z.Timestamp.Msgsize() + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size
	return
}
