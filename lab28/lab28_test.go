package main

import (
	"encoding/json"
	"fmt"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func _TestDecodeToNil(t *testing.T) {
	if true {
		txt := []byte(`[123, "foo"]`)
		var d interface{}
		err := json.Unmarshal(txt, d)
		fmt.Printf("json:d=%#v err=%v\n", d, err)
	}
	// output:
	// json:d=<nil> err=json: Unmarshal(nil)

	if true {
		txt := []byte(`[123, "foo"]`)
		var d interface{}
		err := yaml.Unmarshal(txt, d)
		fmt.Printf("yaml:d=%#v err=%v\n", d, err)
	}
	// output:
	// panic: reflect: call of reflect.Value.Type on zero Value [recovered]
	// panic: reflect: call of reflect.Value.Type on zero Value [recovered]
	// panic: reflect: call of reflect.Value.Type on zero Value
}

func _TestDecodeToValue(t *testing.T) {
	if true {
		txt := []byte(`"foo"`)
		var d string
		err := json.Unmarshal(txt, d)
		fmt.Printf("json:d=%#v err=%v\n", d, err)
	}
	// output:
	// json:d="" err=json: Unmarshal(non-pointer string)

	if true {
		txt := []byte(`"foo"`)
		var d string
		err := yaml.Unmarshal(txt, d)
		fmt.Printf("json:d=%#v err=%v\n", d, err)
	}
	// output:
	// panic: reflect: reflect.Value.Set using unaddressable value [recovered]
	// panic: reflect: reflect.Value.Set using unaddressable value [recovered]
	// panic: reflect: reflect.Value.Set using unaddressable value
}

func TestDecodeToMap(t *testing.T) {
	if true {
		txt := []byte(`{"a":123, "b":"foo"}`)
		var d = map[string]int{}
		err := json.Unmarshal(txt, d)
		fmt.Printf("json:d=%#v err=%v\n", d, err)
	}
	// output:
	// json:d=map[string]int{} err=json: Unmarshal(non-pointer map[string]int)

	if true {
		txt := []byte(`{"a":123, "b":"foo"}`)
		var d = map[string]interface{}{}
		err := yaml.Unmarshal(txt, d)
		fmt.Printf("json:d=%#v err=%v\n", d, err)
	}
	// output:
	// json:d=map[string]interface {}{"a":123, "b":"foo"} err=<nil>
}

func TestDecodeToInterface(t *testing.T) {
	if true {
		txt := `
[
  123, 
  "foo"
]
`
		var d []interface{}

		s := fmt.Sprintf(`[%s]`, txt)
		err := json.Unmarshal([]byte(s), &d)
		fmt.Printf("json:d=%#v err=%v\n", d, err)
	}
	// output:
	// json:d=[]interface {}{[]interface {}{123, "foo"}} err=<nil>

	if true {
		txt := `
- 123
- "foo"
`
		var d []interface{}
		s := fmt.Sprintf(`[%s]`, txt)
		err := yaml.Unmarshal([]byte(s), &d)
		fmt.Printf("yaml:d=%#v err=%v\n", d, err)
	}
	// output:
	// yaml:d=[]interface {}(nil) err=yaml: line 1: did not find expected node content
}
