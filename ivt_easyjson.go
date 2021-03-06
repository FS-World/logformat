// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package logformat

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson3af0d39DecodeGithubComFsWorldLogformat(in *jlexer.Lexer, out *IvtBufferedMeasurement) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "u":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v1 := 0
				for !in.IsDelim(']') {
					if v1 < 50 {
						(out.Voltages)[v1] = int32(in.Int32())
						v1++
					} else {
						in.SkipRecursive()
					}
					in.WantComma()
				}
				in.Delim(']')
			}
		case "i":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v2 := 0
				for !in.IsDelim(']') {
					if v2 < 50 {
						(out.Currents)[v2] = int32(in.Int32())
						v2++
					} else {
						in.SkipRecursive()
					}
					in.WantComma()
				}
				in.Delim(']')
			}
		case "t":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v3 := 0
				for !in.IsDelim(']') {
					if v3 < 8 {
						(out.Temperatures)[v3] = int32(in.Int32())
						v3++
					} else {
						in.SkipRecursive()
					}
					in.WantComma()
				}
				in.Delim(']')
			}
		case "us":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v4 := 0
				for !in.IsDelim(']') {
					if v4 < 11 {
						(out.LVSupplyVoltage)[v4] = float32(in.Float32())
						v4++
					} else {
						in.SkipRecursive()
					}
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ui":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v5 := 0
				for !in.IsDelim(']') {
					if v5 < 11 {
						(out.IVTSupplyVoltage)[v5] = float32(in.Float32())
						v5++
					} else {
						in.SkipRecursive()
					}
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pg5":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v6 := 0
				for !in.IsDelim(']') {
					if v6 < 11 {
						(out.PGood5V)[v6] = bool(in.Bool())
						v6++
					} else {
						in.SkipRecursive()
					}
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pg3v3":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v7 := 0
				for !in.IsDelim(']') {
					if v7 < 11 {
						(out.PGood3V3)[v7] = bool(in.Bool())
						v7++
					} else {
						in.SkipRecursive()
					}
					in.WantComma()
				}
				in.Delim(']')
			}
		case "s":
			(out.SegmentStart).UnmarshalEasyJSON(in)
		case "e":
			(out.SegmentEnd).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3af0d39EncodeGithubComFsWorldLogformat(out *jwriter.Writer, in IvtBufferedMeasurement) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"u\":"
		out.RawString(prefix[1:])
		out.RawByte('[')
		for v8 := range in.Voltages {
			if v8 > 0 {
				out.RawByte(',')
			}
			out.Int32(int32((in.Voltages)[v8]))
		}
		out.RawByte(']')
	}
	{
		const prefix string = ",\"i\":"
		out.RawString(prefix)
		out.RawByte('[')
		for v9 := range in.Currents {
			if v9 > 0 {
				out.RawByte(',')
			}
			out.Int32(int32((in.Currents)[v9]))
		}
		out.RawByte(']')
	}
	{
		const prefix string = ",\"t\":"
		out.RawString(prefix)
		out.RawByte('[')
		for v10 := range in.Temperatures {
			if v10 > 0 {
				out.RawByte(',')
			}
			out.Int32(int32((in.Temperatures)[v10]))
		}
		out.RawByte(']')
	}
	{
		const prefix string = ",\"us\":"
		out.RawString(prefix)
		out.RawByte('[')
		for v11 := range in.LVSupplyVoltage {
			if v11 > 0 {
				out.RawByte(',')
			}
			out.Float32(float32((in.LVSupplyVoltage)[v11]))
		}
		out.RawByte(']')
	}
	{
		const prefix string = ",\"ui\":"
		out.RawString(prefix)
		out.RawByte('[')
		for v12 := range in.IVTSupplyVoltage {
			if v12 > 0 {
				out.RawByte(',')
			}
			out.Float32(float32((in.IVTSupplyVoltage)[v12]))
		}
		out.RawByte(']')
	}
	{
		const prefix string = ",\"pg5\":"
		out.RawString(prefix)
		out.RawByte('[')
		for v13 := range in.PGood5V {
			if v13 > 0 {
				out.RawByte(',')
			}
			out.Bool(bool((in.PGood5V)[v13]))
		}
		out.RawByte(']')
	}
	{
		const prefix string = ",\"pg3v3\":"
		out.RawString(prefix)
		out.RawByte('[')
		for v14 := range in.PGood3V3 {
			if v14 > 0 {
				out.RawByte(',')
			}
			out.Bool(bool((in.PGood3V3)[v14]))
		}
		out.RawByte(']')
	}
	{
		const prefix string = ",\"s\":"
		out.RawString(prefix)
		(in.SegmentStart).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"e\":"
		out.RawString(prefix)
		(in.SegmentEnd).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v IvtBufferedMeasurement) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3af0d39EncodeGithubComFsWorldLogformat(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IvtBufferedMeasurement) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3af0d39EncodeGithubComFsWorldLogformat(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IvtBufferedMeasurement) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3af0d39DecodeGithubComFsWorldLogformat(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IvtBufferedMeasurement) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3af0d39DecodeGithubComFsWorldLogformat(l, v)
}
