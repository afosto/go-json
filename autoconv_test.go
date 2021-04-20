package json

import (
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
}

type SliceValSub struct {
	Dat string
}

func TestUnmarshalSliceAndAutoConvert(t *testing.T) {
	var Vals SliceVals
	jsonDec := NewDecoder(strings.NewReader(`{"Sub":{"Dat":"hi"},"Val":"1","Vals":[1,2],"Bw":["1000000","1000000"]}`))
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	if err := jsonDec.Decode(&Vals); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if Vals.Sub.Dat != "hi" || Vals.Val[0] != 1 || Vals.Vals[1] != 2 || Vals.Bw[1] != 1000000 {
		t.Fatalf("Did not read QuotedVals")
	}
}
