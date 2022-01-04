package provider

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"strings"
)

type MoeDict struct {
	Heteronyms []Heteronym
}

type Heteronym struct {
	Definitions []string
}

const (
	dictionary = `https://www.moedict.tw/a/%s.json`
)

func Query(word string) (dict MoeDict, err error) {
	response, err := http.Get(fmt.Sprintf(dictionary, word))

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	return convertJsonToStruct(string(body))
}

func convertJsonToStruct(json string) (dict MoeDict, err error) {
	m := make(map[interface{}]interface{})
	yaml.Unmarshal([]byte(json), &m)

	heteronyms := convertHeteronyms(m["h"])

	return MoeDict{heteronyms}, err
}

func convertHeteronyms(in interface{}) (out []Heteronym) {
	heteronyms := convertSliceMap(in.([]interface{}))

	for _, heteronym := range heteronyms {
		definitions := convertDefinitions(heteronym["d"])

		out = append(out, Heteronym{
			definitions,
		})
	}

	return
}

func convertSliceMap(sliceMap []interface{}) (sm []map[interface{}]interface{}) {
	for _, m := range sliceMap {
		sm = append(sm, m.(map[interface{}]interface{})) // type transfer m to map[interface{}]interface{}
	}

	return sm
}

func convertDefinitions(in interface{}) (out []string) {
	definitions := convertSliceMap(in.([]interface{}))

	for _, definition := range definitions {
		def := convertDef(definition["f"].(string))
		out = append(out, def)
	}

	return out
}

func convertDef(in string) string {
	def := in

	def = strings.Replace(def, "~", "", -1)
	def = strings.Replace(def, "`", "", -1)
	return def
}
