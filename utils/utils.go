package utils

import (
	"encoding/json"
	"log"
)

func GetStringFromMapInterface(in map[string]interface{}, key, def string) string {
	if len(in) == 0 {
		return def
	}

	v, ok := in[key]
	if !ok {
		return def
	}

	out, ok := v.(string)
	if !ok {
		return def
	}

	return out
}

func DumpObject(c interface{}) {
	b, _ := json.MarshalIndent(c, "", "    ")
	log.Printf("%s\n", string(b))
}
