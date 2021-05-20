package lab27

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
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

func TestCase5901(t *testing.T) {
	type AnimalRaw struct {
		Kind string      `json:"kind"`
		Attr interface{} `json:"attr"`
	}
	var animals []AnimalRaw
	var dec = json.NewDecoder(bytes.NewReader(sampleJson))
	dec.Register(func(b []byte, v *interface{}) error {
		*v = append(json.RawMessage(nil), b...)
		return nil
	})
	dec.Decode(&animals)
	fmt.Printf("%#v\n", animals)
}

func TestCase5901_default(t *testing.T) {
	type AnimalRaw struct {
		Kind string      `json:"kind"`
		Attr interface{} `json:"attr"`
	}
	var animals []AnimalRaw
	json.Unmarshal(sampleJson, &animals)
	fmt.Printf("%#v\n", animals)
}

func TestCase5901_3rdParty(t *testing.T) {
	type AnimalRaw struct {
		Kind string      `json:"kind"`
		Attr interface{} `json:"attr"`
	}
	var animals []AnimalRaw
	gin.Context.BindJSON(&animals)

	fmt.Printf("%#v\n", animals)
	/*
		func decodeJSON(r io.Reader, obj interface{}) error {
			decoder := json.NewDecoder(r)
			if EnableDecoderUseNumber {
				decoder.UseNumber()
			}
			if EnableDecoderDisallowUnknownFields {
				decoder.DisallowUnknownFields()
			}
			if err := decoder.Decode(obj); err != nil {
				return err
			}
			return validate(obj)
		}
	*/
}
