package request

import "net/http"

func Queries(r *http.Request) map[string]string {
	query := r.URL.Query()
	result := make(map[string]string)
	for key, values := range query {
		result[key] = values[0]
	}
	return result
}

func Query(r *http.Request, key, defaultValue string) string {
	query := r.URL.Query()
	if value, ok := query[key]; ok {
		return value[0]
	}
	return defaultValue
}
