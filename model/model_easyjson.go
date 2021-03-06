// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

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

func easyjsonC80ae7adDecodeTaskModel(in *jlexer.Lexer, out *Price) {
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
		case "symbol":
			out.Symbol = string(in.String())
		case "price":
			out.Price = string(in.String())
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
func easyjsonC80ae7adEncodeTaskModel(out *jwriter.Writer, in Price) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"symbol\":"
		out.RawString(prefix[1:])
		out.String(string(in.Symbol))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.String(string(in.Price))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Price) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeTaskModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Price) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeTaskModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Price) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeTaskModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Price) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeTaskModel(l, v)
}
