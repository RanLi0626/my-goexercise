package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Dictionary struct {
	AttrKey         string                 `json:"attrKey"`
	Decimal         int                    `json:"decimal"`
	Modifier        int                    `json:"modifier"`
	DisplayName     map[string]string      `json:"displayName"`
	DisplayType     string                 `json:"displayType"`
	Unit            map[string]string      `json:"unit"`
	ValidationRules map[string][]KeyParams `json:"validationRules"`
}

type KeyParams struct {
	Key    string        `json:"key"`
	Params []interface{} `json:"params"`
}

func main() {
	f, err := os.Create("result.csv")
	if err != nil {
		log.Fatalf("create csv error , %v", err)
	}
	defer f.Close()
	w := csv.NewWriter(f)

	file, err := ioutil.ReadFile("dictionary.json")
	if err != nil {
		log.Fatalf("open file error, %v", err)
	}
	dic := []Dictionary{}
	json.Unmarshal(file, &dic)
	var data [][]string
	for _, v := range dic {
		var dn string
		dn = v.DisplayName[v.AttrKey]

		var units [3]string
		index := 0
		for _, v := range v.Unit {
			units[index] = v
			index++
		}

		var vrs [5]string
		for k, vr := range v.ValidationRules {
			if k == v.AttrKey {
				vrs[0] = buildStr(vr)
			}
			if k == "volume_liter" {
				vrs[1] = buildStr(vr)
			}
			if k == "mass_kg" {
				vrs[2] = buildStr(vr)
			}
			if k == "target" {
				vrs[3] = buildStr(vr)
			}
			if k == "result" {
				vrs[4] = buildStr(vr)
			}
		}

		datarow := []string{
			v.AttrKey, strconv.Itoa(v.Decimal), strconv.Itoa(v.Modifier), dn, v.DisplayType, units[0], units[1], units[2], vrs[0], vrs[1], vrs[2], vrs[3], vrs[4],
		}
		data = append(data, datarow)
	}
	w.WriteAll(data)
}

func buildStr(kps []KeyParams) string {
	var kpsStr []string
	var str string
	for _, kp := range kps {
		var params []string
		var pStr string
		for _, v := range kp.Params {
			switch v.(type) {
			case float64:
				v = v.(float64)
				params = append(params, strconv.FormatFloat(v.(float64), 'f', 0, 64))
			case string:
				params = append(params, v.(string))
			case bool:
				params = append(params, strconv.FormatBool(v.(bool)))
			}
			pStr = strings.Join(params, ",")
		}
		kpsStr = append(kpsStr, kp.Key+":"+pStr)
		str = strings.Join(kpsStr, "|")
	}
	return str
}
