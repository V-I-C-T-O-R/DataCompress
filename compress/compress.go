package compress

import (
	"encoding/json"
	"errors"
	"github.com/V-I-C-T-O-R/DataCompress/utils"
	"log"
	"reflect"
	"strconv"
)

type MesMark struct {
	Mark  bool
	Key   interface{}
	Value interface{}
}

func DoCompress(file string, output string) error {
	if !utils.CheckFileIsExist(file) {
		log.Println("file no exist")
		return errors.New("file no exist")
	}
	s := utils.ReadF(file)
	var content interface{}
	err := json.Unmarshal(s, &content)
	if err != nil {
		log.Println("json file unmarshal failed")
		return err
	}
	data, _ := json.Marshal(baseMap(&MesMark{Mark: true, Value: (content).(map[string]interface{})}))
	utils.WriteFile(data, output)
	log.Println("file compress complete")
	return nil
}
func DoStreamCompress(b []byte, output string) error {
	var content interface{}
	err := json.Unmarshal(b, &content)
	if err != nil {
		log.Println("json file unmarshal failed")
		return err
	}
	data, _ := json.Marshal(baseMap(&MesMark{Mark: true, Value: (content).(map[string]interface{})}))
	utils.WriteFile(data, output)
	log.Println("data compress complete")
	return nil
}
func baseMap(m *MesMark) (x interface{}) {
	if m.Mark {
		x = make(map[string]interface{})
		for k, v := range m.Value.(map[string]interface{}) {
			switch v.(type) {
			case bool, byte, int, int8, int16, int32, int64, uint16, uint32, uint64, float32, float64, string:
				x.(map[string]interface{})[k] = v
				continue
			case []interface{}:
				x.(map[string]interface{})[k] = baseMap(&MesMark{Mark: false, Value: v, Key: k})
			default:
				m.Value.(map[string]interface{})[k] = baseMap(&MesMark{Mark: true, Value: v, Key: k})
			}
		}
	} else {
		listCount := make(map[interface{}]int)
		var listMap []interface{}
		var flag bool
		for _, v := range m.Value.([]interface{}) {
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
				baseMap(&MesMark{Mark: false, Value: v})
			default:
				listMap = append(listMap, baseMap(&MesMark{Mark: true, Value: v}))
			}
		}
		if flag {
			slice := []string{}
			for k, v := range listCount {
				str := toString(k) + "::" + strconv.Itoa(v)
				slice = append(slice, str)
			}
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
		s = strconv.FormatBool(value.Bool()) + "::" + "1"
	case reflect.Int:
		s = strconv.Itoa(int(value.Int())) + "::" + "2"
	case reflect.Int8:
		s = strconv.Itoa(int(value.Int())) + "::" + "3"
	case reflect.Int16:
		s = strconv.Itoa(int(value.Int())) + "::" + "4"
	case reflect.Int32:
		s = strconv.Itoa(int(value.Int())) + "::" + "5"
	case reflect.Int64:
		s = strconv.Itoa(int(value.Int())) + "::" + "6"
	case reflect.Uint8:
		s = strconv.Itoa(int(value.Int())) + "::" + "7"
	case reflect.Uint16:
		s = strconv.Itoa(int(value.Int())) + "::" + "8"
	case reflect.Uint32:
		s = strconv.Itoa(int(value.Int())) + "::" + "9"
	case reflect.Uint64:
		s = strconv.Itoa(int(value.Int())) + "::" + "10"
	case reflect.Float32:
		s = strconv.FormatFloat(float64(value.Float()), 'g', 8, 32) + "::" + "11"
	case reflect.Float64:
		s = strconv.FormatFloat(float64(value.Float()), 'g', 8, 32) + "::" + "12"
	case reflect.String:
		s = value.String() + "::" + "13"
	default:
		s = strconv.Itoa(int(value.Int())) + "::" + "14"
	}
	return
}
