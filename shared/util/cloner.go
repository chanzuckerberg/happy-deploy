package util

import (
	"encoding/json"
	"reflect"
)

func DeepClone(dst, src interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, &dst); err != nil {
		return err
	}
	return nil
}

func DeepMerge(dst, src map[string]interface{}) map[string]interface{} {
	for k, v := range src {
		t := reflect.TypeOf(v)
		if t == nil {
			continue
		}
		switch v.(type) {
		case map[string]interface{}:
			dst[k] = DeepMerge(v.(map[string]interface{}), dst[k].(map[string]interface{}))
		default:
			if _, ok := dst[k]; ok {
				if v != nil {
					dst[k] = v
				}
			}
		}
	}
	return dst
}
