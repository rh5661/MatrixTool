// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package dbModify

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

func easyjson8da5bf1dDecodeGithubComRh5661MatrixToolPkgDbModify(in *jlexer.Lexer, out *QueryParameters) {
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
		case "startDate":
			out.StartDate = string(in.String())
		case "util":
			out.Util = string(in.String())
		case "dualBilling":
			out.DualBilling = bool(in.Bool())
		case "terms":
			if in.IsNull() {
				in.Skip()
				out.Terms = nil
			} else {
				in.Delim('[')
				if out.Terms == nil {
					if !in.IsDelim(']') {
						out.Terms = make([]int, 0, 8)
					} else {
						out.Terms = []int{}
					}
				} else {
					out.Terms = (out.Terms)[:0]
				}
				for !in.IsDelim(']') {
					var v1 int
					v1 = int(in.Int())
					out.Terms = append(out.Terms, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson8da5bf1dEncodeGithubComRh5661MatrixToolPkgDbModify(out *jwriter.Writer, in QueryParameters) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"startDate\":"
		out.RawString(prefix[1:])
		out.String(string(in.StartDate))
	}
	{
		const prefix string = ",\"util\":"
		out.RawString(prefix)
		out.String(string(in.Util))
	}
	{
		const prefix string = ",\"dualBilling\":"
		out.RawString(prefix)
		out.Bool(bool(in.DualBilling))
	}
	{
		const prefix string = ",\"terms\":"
		out.RawString(prefix)
		if in.Terms == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Terms {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v QueryParameters) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8da5bf1dEncodeGithubComRh5661MatrixToolPkgDbModify(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v QueryParameters) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8da5bf1dEncodeGithubComRh5661MatrixToolPkgDbModify(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *QueryParameters) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8da5bf1dDecodeGithubComRh5661MatrixToolPkgDbModify(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *QueryParameters) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8da5bf1dDecodeGithubComRh5661MatrixToolPkgDbModify(l, v)
}