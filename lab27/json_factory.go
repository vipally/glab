package lab27

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// Factory impliments a factory that can create multi products by type name
type Factory struct {
	mp map[string]reflect.Type
}

// NewFactory create a new factory
func NewFactory() *Factory {
	return &Factory{mp: make(map[string]reflect.Type)}
}

// MustReg register the creator by name, it panic if name is duplicate
func (f *Factory) MustReg(name string, v interface{}) {
	err := f.Reg(name, v)
	if err != nil {
		panic(err)
	}
}

// MustReg register the creator by name
func (f *Factory) Reg(name string, v interface{}) error {
	if _, ok := f.mp[name]; ok {
		return fmt.Errorf("duplicate reg of %s,%#v", name, v)
	}
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	f.mp[name] = t
	return nil
}

// Create make product by name
func (f *Factory) Create(name string) (interface{}, error) {
	t, ok := f.mp[name]
	if !ok {
		return nil, fmt.Errorf("product %s cannot create from factory %#v", name, f)
	}
	return reflect.New(t).Interface(), nil
}

// UnmarshalJSONForFlexObj create the real object and ummarshal it for FlexObject
func (f *Factory) DelayedFlexObjectJSONUnmarshal(kind string, obj *json.FlexObject) error {
	b, ok := obj.D.([]byte)
	if !ok {
		return errors.New("FlexObject type isn't []byte")
	}
	p, err := f.Create(kind)
	if err != nil {
		return err
	}
	obj.D = p
	return json.Unmarshal(b, obj.D)
}
