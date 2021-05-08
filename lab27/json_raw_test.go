package lab27

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeRaw(t *testing.T) {
	type AnimalRaw struct {
		Kind string          `json:"kind"`
		Attr json.RawMessage `json:"attr"`
	}
	var factory = NewFactory()
	factory.MustReg("dog", (*DogAttr)(nil))
	factory.MustReg("duck", (*DuckAttr)(nil))

	var animals []AnimalRaw
	json.Unmarshal(sampleJson, &animals)
	for i, v := range animals {
		d, _ := factory.Create(v.Kind)
		json.Unmarshal(v.Attr, d)
		fmt.Printf("index %d, kind=%s attr=%#v\n", i, v.Kind, d)
	}
}

func TestEncodeRaw(t *testing.T) {
	type AnimalRaw struct {
		Kind string          `json:"kind"`
		Attr json.RawMessage `json:"attr"`
	}
	var animals = []AnimalRaw{
		AnimalRaw{
			Kind: "dog",
			Attr: []byte(`{"type": "Collie","color": "white"}`),
		},
		AnimalRaw{
			Kind: "duck",
			Attr: []byte(`{"Weight": 2.34}`),
		},
	}
	b, _ := json.MarshalIndent(animals, "", "  ")
	fmt.Println(string(b))
}
