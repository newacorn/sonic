package decoder

// NewDecoderCopy creates a  decoder instance prevent escape.
func NewDecoderCopy(s string) Decoder {
	return Decoder{s: s}
}
