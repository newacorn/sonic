package sonic

import (
	"github.com/bytedance/sonic/internal/decoder/api"
	"github.com/bytedance/sonic/internal/encoder"
)

type EncoderOptions = encoder.Options

const (
	// SortMapKeys indicates that the keys of a map needs to be sorted
	// before serializing into JSON.
	// WARNING: This hurts performance A LOT, USE WITH CARE.
	SortMapKeys EncoderOptions = encoder.SortMapKeys

	// EscapeHTML indicates encoder to escape all HTML characters
	// after serializing into JSON (see https://pkg.go.dev/encoding/json#HTMLEscape).
	// WARNING: This hurts performance A LOT, USE WITH CARE.
	EscapeHTML EncoderOptions = encoder.EscapeHTML

	// CompactMarshaler indicates that the output JSON from json.Marshaler
	// is always compact and needs no validation
	CompactMarshaler EncoderOptions = encoder.CompactMarshaler

	// NoQuoteTextMarshaler indicates that the output text from encoding.TextMarshaler
	// is always escaped string and needs no quoting
	NoQuoteTextMarshaler EncoderOptions = encoder.NoQuoteTextMarshaler

	// NoNullSliceOrMap indicates all empty Array or Object are encoded as '[]' or '{}',
	// instead of 'null'
	NoNullSliceOrMap EncoderOptions = encoder.NoNullSliceOrMap

	// ValidateString indicates that encoder should validate the input string
	// before encoding it into JSON.
	ValidateString EncoderOptions = encoder.ValidateString

	// NoValidateJSONMarshaler indicates that the encoder should not validate the output string
	// after encoding the JSONMarshaler to JSON.
	NoValidateJSONMarshaler EncoderOptions = encoder.NoValidateJSONMarshaler

	// NoEncoderNewline indicates that the encoder should not add a newline after every message
	NoEncoderNewline EncoderOptions = encoder.NoEncoderNewline

	// CompatibleWithStd is used to be compatible with std encoder.
	CompatibleWithStd EncoderOptions = SortMapKeys | EscapeHTML | CompactMarshaler
)

type DecoderOptions = api.Options

const (
	OptionUseInt64         DecoderOptions = api.OptionUseInt64
	OptionUseNumber        DecoderOptions = api.OptionUseNumber
	OptionUseUnicodeErrors DecoderOptions = api.OptionUseUnicodeErrors
	OptionDisableUnknown   DecoderOptions = api.OptionDisableUnknown
	OptionCopyString       DecoderOptions = api.OptionCopyString
	OptionValidateString   DecoderOptions = api.OptionValidateString
)
const (
	StdEncoderOpts     = encoder.EscapeHTML | encoder.SortMapKeys | encoder.CompactMarshaler
	StdDecoderOpts     = api.OptionCopyString | api.OptionValidateString
	FastEncoderOpts    = encoder.NoQuoteTextMarshaler | encoder.NoValidateJSONMarshaler
	FastDecoderOpts    = 0
	DefaultEncoderOpts = 0
	DefaultDecoderOpts = 0
)

func EncodeInto(buf *[]byte, val interface{}, opts encoder.Options) error {
	return encoder.EncodeInto(buf, val, opts)
}
func DecodeString(buf string, val interface{}, opts api.Options) error {
	dec := api.NewDecoderCopy(buf)
	dec.SetOptions(opts)
	err := dec.Decode(val)
	/* check for errors */
	if err != nil {
		return err
	}

	return dec.CheckTrailings()
}
