package paa

const (
	TaggSignature = "TAGG"
	AVG           = "AVGC"
	MAX           = "MAXC"
	FLAG          = "FLAG"
	SWIZ          = "SWIZ"
	PROC          = "PROC"
	OFFS          = "OFFS"
)

type Tagg struct {
	signature, name string
	dataLength      int32
	data            []int32
}

func (t Tagg) Signature() string {
	return t.signature
}

func (t Tagg) Name() string {
	return t.name
}

func (t Tagg) DataLength() int32 {
	return t.dataLength
}

func (t Tagg) Data() []int32 {
	return t.data
}
