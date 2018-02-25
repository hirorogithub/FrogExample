package util

import "encoding/json"

//help me simply reflect from map[string]interface{} to struct
//dst should be ptr to struct
func MapToStruct(src map[string]interface{}, dst interface{}) bool {
	jdata, err := json.Marshal(src)
	if err != nil {
		return false
	}

	err = json.Unmarshal(jdata, &dst)
	if err != nil {
		return false
	}
	return true
}
