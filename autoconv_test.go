package json

import (
	"fmt"
	"strings"
	"testing"
)

type QuotedVals struct {
	IntQ  int
	BoolQ bool
}

func TestUnmarshalQuoted(t *testing.T) {
	var quotedVals QuotedVals
	jsonDec := NewDecoder(strings.NewReader(`{"IntQ":"1","BoolQ":"true"}`))
	jsonDec.UseAutoConvert()
	if err := jsonDec.Decode(&quotedVals); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if quotedVals.IntQ != 1 || quotedVals.BoolQ != true {
		t.Fatalf("Did not read QuotedVals")
	}
}

type SliceVals struct {
	Sub  SliceValSub
	Val  []int
	Vals []int
	Bw   []int
	FV   fancyVal
	NI   newInt
}

type SliceValSub struct {
	Dat string
}

func TestUnmarshalSliceAutoConvertCustomType(t *testing.T) {
	var Vals SliceVals
	jsonDec := NewDecoder(strings.NewReader(`{"Sub":{"Dat":"hi"},"Val":"1","Vals":[1,2],"Bw":["1000000","1000000"],"FV":"lang","NI":"45"}`))
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	if err := jsonDec.Decode(&Vals); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if Vals.Sub.Dat != "hi" || Vals.Val[0] != 1 || Vals.Vals[1] != 2 ||
		Vals.Bw[1] != 1000000 || Vals.FV.dat != "lang" {
		t.Fatalf("Did not read correctly")
	}

	data, _ := Marshal(Vals)
	if string(data) != `{"Sub":{"Dat":"hi"},"Val":[1],"Vals":[1,2],"Bw":[1000000,1000000],"FV":"lang","NI":"test"}` {
		t.Fatalf("Did not marshal correctly")
	}
}

type fancyVal struct {
	dat string
}

func (f *fancyVal) UnmarshalText(text []byte) error {
	f.dat = string(text)
	return nil
}
func (f fancyVal) MarshalText() (text []byte, err error) {
	return []byte(f.dat), nil
}

type newInt int

func (i newInt) UnmarshalText(text []byte) error {
	i = 42
	return nil
}
func (i newInt) MarshalText() (text []byte, err error) {
	return []byte(fmt.Sprintf("test")), nil
}
