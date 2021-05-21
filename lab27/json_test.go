package lab27

import (
	"encoding/json"
	"fmt"
	"testing"
)

var sampleJson = []byte(`
[
    {
        "kind":"dog",
        "attr":{
            "type":"Collie",
            "color":"black"
        }
    },
    {
        "kind":"duck",
        "attr":{
            "weight":1.2
        }
    }
]
`)

type DogAttr struct {
	Type  string `json:"type"`
	Color string `json:"color"`
}

type DuckAttr struct {
	Weight float64
}

func TestFlexObjectFactory(t *testing.T) {
	var factory = NewFactory()
	factory.MustReg("dog", (*DogAttr)(nil))
	factory.MustReg("duck", (*DuckAttr)(nil))

	type Animal struct {
		Kind string          `json:"kind"`
		Attr json.FlexObject `json:"attr"`
	}
	var animals []Animal
	json.Unmarshal(sampleJson, &animals)
	for i, v := range animals {
		factory.DelayedFlexObjectJSONUnmarshal(v.Kind, &v.Attr)
		fmt.Printf("index %d, kind=%s attr=%#v\n", i, v.Kind, v.Attr.D)
	}
	// Output:
	// index 0, kind=dog attr=&lab27.DogAttr{Type:"Collie", Color:"black"}
	// index 1, kind=duck attr=&lab27.DuckAttr{Weight:1.2}
}

func TestGenerateJsonByFlexObject(t *testing.T) {
	type Animal struct {
		Kind string          `json:"kind"`
		Attr json.FlexObject `json:"attr"`
	}
	var animals = []Animal{
		Animal{
			Kind: "dog",
			Attr: json.FlexObject{
				D: DogAttr{
					Type:  "Collie",
					Color: "white",
				},
			},
		},
		Animal{
			Kind: "duck",
			Attr: json.FlexObject{
				D: DuckAttr{
					Weight: 2.34,
				},
			},
		},
	}
	b, _ := json.MarshalIndent(animals, "", "  ")
	fmt.Println(string(b))
	// Ooutput:
	// [
	//  {
	//    "kind": "dog",
	//    "attr": {
	//      "type": "Collie",
	//      "color": "white"
	//    }
	//  },
	//  {
	//    "kind": "duck",
	//    "attr": {
	//      "Weight": 2.34
	//    }
	//  }
	// ]
}
