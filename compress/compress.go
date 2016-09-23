package compress

import (
	"encoding/json"
	"github.com/V-I-C-T-O-R/DataCompress/utils"
	"reflect"
	"strconv"
)

var MarkMap map[string][]string

type MesMark struct {
	mark  bool
	key   interface{}
	value interface{}
}

func init() {
	MarkMap = make(map[string][]string)
}
func DoCompress(file string, output string) {
	var content interface{}
	s := utils.ReadF(file)
	json.Unmarshal(s, &content)
	data, _ := json.Marshal(baseMap(&MesMark{mark: true, value: (content).(map[string]interface{})}))
	utils.WriteFile(data, output)
}
func baseMap(m *MesMark) (x interface{}) {
	if m.mark {
		x = make(map[string]interface{})
		for k, v := range m.value.(map[string]interface{}) {
			switch v.(type) {
			case bool, byte, int, int8, int16, int32, int64, uint16, uint32, uint64, float32, float64, string:
				x.(map[string]interface{})[k] = v
				continue
			case []interface{}:
				x.(map[string]interface{})[k] = baseMap(&MesMark{mark: false, value: v, key: k})
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
			MarkMap[m.key.(string)] = slice
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
