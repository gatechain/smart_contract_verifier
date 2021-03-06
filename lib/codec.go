package lib

import (
	"github.com/tendermint/go-amino"
)

// amino codec to marshal/unmarshal
type Codec = amino.Codec

func NewCdc() *Codec {
	return amino.NewCodec()
}

// generic sealed codec to be used throughout framework
var Cdc *Codec

func init() {
	cdc := NewCdc()
	Cdc = cdc.Seal()
}
