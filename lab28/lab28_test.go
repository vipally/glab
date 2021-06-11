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

func TestDecodeToInterface(t *testing.T) {
	if true {
		txt := `[123, "foo"]`
		var d []interface{}

		s := fmt.Sprintf(`[%s]`, txt)
		err := json.Unmarshal([]byte(s), &d)
		fmt.Printf("json:d=%#v err=%v\n", d, err)
	}
	// output:
	// json:d=[]interface {}{[]interface {}{123, "foo"}} err=<nil>

	if true {
		txt := `[123, "foo"]`
		var d []interface{}

		s := fmt.Sprintf(`- %s`, txt)
		err := yaml.Unmarshal([]byte(s), &d)
		fmt.Printf("yaml:d=%#v err=%v\n", d, err)
	}
	// output:
	// yaml:d=[]interface {}{[]interface {}{123, "foo"}} err=<nil>

}
