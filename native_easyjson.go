// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package openrtb

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

func easyjson3eeb1cddDecodeGithubComBsmOpenrtb(in *jlexer.Lexer, out *Native) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "request":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Request).UnmarshalJSON(data))
			}
		case "ver":
			out.Ver = string(in.String())
		case "api":
			if in.IsNull() {
				in.Skip()
				out.API = nil
			} else {
				in.Delim('[')
				if out.API == nil {
					if !in.IsDelim(']') {
						out.API = make([]int, 0, 8)
					} else {
						out.API = []int{}
					}
				} else {
					out.API = (out.API)[:0]
				}
				for !in.IsDelim(']') {
					var v1 int
					v1 = int(in.Int())
					out.API = append(out.API, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "battr":
			if in.IsNull() {
				in.Skip()
				out.BAttr = nil
			} else {
				in.Delim('[')
				if out.BAttr == nil {
					if !in.IsDelim(']') {
						out.BAttr = make([]int, 0, 8)
					} else {
						out.BAttr = []int{}
					}
				} else {
					out.BAttr = (out.BAttr)[:0]
				}
				for !in.IsDelim(']') {
					var v2 int
					v2 = int(in.Int())
					out.BAttr = append(out.BAttr, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ext":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Ext).UnmarshalJSON(data))
			}
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
func easyjson3eeb1cddEncodeGithubComBsmOpenrtb(out *jwriter.Writer, in Native) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"request\":")
	out.Raw((in.Request).MarshalJSON())
	if in.Ver != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ver\":")
		out.String(string(in.Ver))
	}
	if len(in.API) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"api\":")
		if in.API == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.API {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v4))
			}
			out.RawByte(']')
		}
	}
	if len(in.BAttr) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"battr\":")
		if in.BAttr == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.BAttr {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v6))
			}
			out.RawByte(']')
		}
	}
	if len(in.Ext) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ext\":")
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Native) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3eeb1cddEncodeGithubComBsmOpenrtb(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Native) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3eeb1cddEncodeGithubComBsmOpenrtb(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Native) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3eeb1cddDecodeGithubComBsmOpenrtb(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Native) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3eeb1cddDecodeGithubComBsmOpenrtb(l, v)
}
