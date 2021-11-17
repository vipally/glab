package lab30

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestReflect(t *testing.T) {
	var d testCase
	d.init()
	fmt.Println(d.accessAll(d.data))
	fmt.Println("------------")
	fmt.Println(d.accessAllInter(d.data))
	var x []interface{} = []interface{}{
		map[string]interface{}{
			"a": "foo",
			"b": 123,
		},
	}
	b, _ := json.MarshalIndent(x, "", "  ")
	var xx interface{}
	json.Unmarshal(b, &xx)
	fmt.Println(reflect.TypeOf(x[0]).Kind().String(), reflect.ValueOf(x[0]).Kind().String())
	valueOfValue := reflect.ValueOf(x)
	x0 := valueOfValue.Index(0).Interface()
	fmt.Println(reflect.TypeOf(x0).Kind().String(), reflect.ValueOf(x0).Kind().String())
}

type testCase struct {
	data interface{}
	buf  *bytes.Buffer
}

func (c *testCase) init() {
	d := map[string]interface{}{
		"a": "f",
		"b": []string{"foo", "bar"},
		"c": map[string]interface{}{
			"c1": 1,
			"c2": map[string]interface{}{
				"c21": "foo",
				"c22": []interface{}{
					[]interface{}{1, "foo"},
					map[string]interface{}{
						"c221": "foo",
						"c222": true,
					},
				},
			},
		},
		"d": []interface{}{
			map[string]interface{}{
				"c21": "foo",
				"c22": []interface{}{
					[]interface{}{1, "foo"},
					map[string]interface{}{
						"c221": "foo",
						"c222": true,
					},
				},
			},
			[]interface{}{
				[]interface{}{1, "foo"},
				map[string]interface{}{
					"c221": []string{"foo", "bar"},
					"c222": true,
				},
			},
		},
	}
	b, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		panic(err)
	}
	c.data = nil
	if err := json.Unmarshal(b, &c.data); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", c.data)
}
func (c *testCase) accessAll(value interface{}) string {
	c.buf = bytes.NewBuffer(nil)
	c.access("", value, 0)
	return c.buf.String()
}

func (c *testCase) access(name string, value interface{}, depth int) error {

	switch reflect.TypeOf(value).Kind() {
	case reflect.Map:
		return c.accessMap(name, value, depth)
	case reflect.Array, reflect.Slice:
		return c.accessArray(name, value, depth)
	default:
	}

	return c.accessSingle(name, value, depth)

}

func (c *testCase) accessMap(key string, value interface{}, depth int) error {
	key = division(key)

	//var keys []string
	valueOfValue := reflect.ValueOf(value)
	//numField := valueOfValue.Len()
	//keys = make([]string, 0, numField)
	//vals := make(map[string]string, numField)

	iter := valueOfValue.MapRange()
	for iter.Next() {
		if !iter.Value().CanInterface() {
			continue
		}
		k := iter.Key().String()
		//keys = append(keys, k)

		val := iter.Value().Interface()
		c.access(key+k, val, depth+1)
	}

	//	sort.Strings(keys)

	return nil
}
func (c *testCase) accessArray(key string, value interface{}, depth int) error {
	key = division(key)

	valueOfValue := reflect.ValueOf(value)
	numField := valueOfValue.Len()

	for i := 0; i < numField; i++ {
		item := valueOfValue.Index(i)
		if !item.CanInterface() {
			continue
		}
		c.access(key+strconv.Itoa(i+1), item, depth+1)
	}
	return nil
}
func (c *testCase) accessSingle(name string, value interface{}, depth int) error {
	c.buf.WriteString(fmt.Sprintf("%s=%v\n", name, value))
	return nil
}

func division(key string) string {
	if key != "" {
		key = key + "."
	}
	return key
}

//------------------------------------------------------------------------------

func (c *testCase) accessAllInter(data interface{}) string {
	c.buf.Reset()
	c.accessInterface("", data, 0)
	return c.buf.String()
}

func (c *testCase) accessInterface(name string, d interface{}, depth int) error {
	if depth >= 10 {
		return errors.New("buildQuery out of recursion")
	}
	writeSingle := func(v interface{}) {
		c.buf.WriteString(fmt.Sprintf("%s=%v\n", name, d))
	}
	fromBytes := func(v []byte) error {
		var x interface{}
		if err := json.Unmarshal(v, &x); err != nil {
			return err
		}
		return c.accessInterface(name, x, depth+1)
	}
	switch v := d.(type) {
	case string:
		writeSingle(v)
	case float64:
		writeSingle(v)
	case bool:
		writeSingle(v)
	case map[string]interface{}:
		names := make([]string, 0, len(v))
		for k := range v {
			names = append(names, k)
		}
		sort.Strings(names)
		for _, k := range names {
			n := k
			if name != "" {
				n = fmt.Sprintf("%s.%s", name, k)
			}
			if val := v[k]; val != nil {
				if err := c.accessInterface(n, v[k], depth+1); err != nil {
					return err
				}
			}
		}
	case []interface{}:
		if name != "" {
			for i, vv := range v {
				n := fmt.Sprintf("%s.%d", name, i+1)
				c.accessInterface(n, vv, depth+1)
			}
		}
	//--------------------------------------------------------------------------
	default:
		b, err := json.Marshal(d)
		if err != nil {
			return err
		}
		return fromBytes(b)
	case []byte:
		return fromBytes(v)
	case json.RawMessage:
		return fromBytes([]byte(v))
	}
	return nil
	return nil
}
