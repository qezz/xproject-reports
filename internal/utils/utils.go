package utils

func GetProperty(jsonNode map[string]interface{}, path ...string) interface{} {
	if len(path) == 0 {
		return jsonNode
	}

	value := jsonNode[path[0]]

	switch value.(type) {
	case map[string]interface{}:
		return GetProperty(value.(map[string]interface{}), path[1:]...)
	case string:
		return value.(string)
	default:
		return nil
	}
}
