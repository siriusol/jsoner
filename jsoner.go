package jsoner

import (
	"encoding/json"
	"sort"
)

type Object struct {
	Type string
}

type JSON struct {
	m map[string]interface{}
}

type Value struct {
	Type string
	data []byte
}

func New(data []byte) (*JSON, error) {
	m := map[string]interface{}{}

	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	return &JSON{
		m: m,
	}, nil
}

func parseJSON(data []byte) error {
	return nil
}

func (j *JSON) Keys() []string {
	return nil
}

func (j *JSON) KeysInOrder() []string {
	keys := make([]string, 0, len(j.m))
	for k := range j.m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}

func (j *JSON) GetValue(key string) {

}
