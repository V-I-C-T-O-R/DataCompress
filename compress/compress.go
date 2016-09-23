package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
)

var FilePath string

type MesMark struct {
	mark  bool
	key   interface{}
	value interface{}
}

var markMap map[string][]string

func init() {
	markMap = make(map[string][]string)
	flag.StringVar(&FilePath, "FilePath", "example.json", "read to compress's file")
}
func main() {
	flag.Parse()
	var content interface{}
	s := readF()
	json.Unmarshal(s, &content)
	iter := (content).(map[string]interface{})
	m := &MesMark{mark: true, value: iter}
	fmt.Println("baseMap(m)", baseMap(m))
}
func baseMap(m *MesMark) (x interface{}) {
	if m.mark {
		x = make(map[interface{}]interface{})
		for k, v := range m.value.(map[string]interface{}) {
			switch v.(type) {
			case bool, byte, int, int8, int16, int32, int64, uint16, uint32, uint64, float32, float64, string:
				x.(map[interface{}]interface{})[k] = v
				continue
			case []interface{}:
				x.(map[interface{}]interface{})[k] = baseMap(&MesMark{mark: false, value: v, key: k})
			default:
				m.value.(map[string]interface{})[k] = baseMap(&MesMark{mark: true, value: v, key: k})
			}
		}
	} else {
		listCount := make(map[interface{}]int)
		var listMap []interface{}
		var flag bool
		for _, v := range m.value.([]interface{}) {
			switch v.(type) {
			case bool, byte, int, int8, int16, int32, int64, uint16, uint32, uint64, float32, float64, string:
				if value, ok := listCount[v]; !ok {
					listCount[v] = 1
				} else {
					listCount[v] = value + 1
				}
				flag = true
				continue
			case []interface{}:
				baseMap(&MesMark{mark: false, value: v})
			default:
				listMap = append(listMap, baseMap(&MesMark{mark: true, value: v}))
			}
		}
		if flag {
			slice := []string{}
			for k, v := range listCount {
				str := toString(k) + ":" + strconv.Itoa(v)
				slice = append(slice, str)
			}
			markMap[m.key.(string)] = slice
			x = slice
		} else {
			x = listMap
		}
	}
	return
}
func toString(v interface{}) (s string) {
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Bool:
		s = strconv.FormatBool(value.Bool()) + ":" + "1"
	case reflect.Int:
		s = strconv.Itoa(int(value.Int())) + ":" + "2"
	case reflect.Int8:
		s = strconv.Itoa(int(value.Int())) + ":" + "3"
	case reflect.Int16:
		s = strconv.Itoa(int(value.Int())) + ":" + "4"
	case reflect.Int32:
		s = strconv.Itoa(int(value.Int())) + ":" + "5"
	case reflect.Int64:
		s = strconv.Itoa(int(value.Int())) + ":" + "6"
	case reflect.Uint8:
		s = strconv.Itoa(int(value.Int())) + ":" + "7"
	case reflect.Uint16:
		s = strconv.Itoa(int(value.Int())) + ":" + "8"
	case reflect.Uint32:
		s = strconv.Itoa(int(value.Int())) + ":" + "9"
	case reflect.Uint64:
		s = strconv.Itoa(int(value.Int())) + ":" + "10"
	case reflect.Float32:
		s = strconv.FormatFloat(float64(value.Float()), 'g', 3, 32) + ":" + "11"
	case reflect.Float64:
		s = strconv.FormatFloat(float64(value.Float()), 'g', 3, 32) + ":" + "12"
	case reflect.String:
		s = value.String() + ":" + "13"
	default:
		s = strconv.Itoa(int(value.Int())) + ":" + "14"
	}
	return
}
func readF() []byte {
	_, filename, _, _ := runtime.Caller(1)
	basePath := path.Join(path.Dir(filename), "..", "example")
	file, err := os.Open(basePath + "/" + FilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return content
}