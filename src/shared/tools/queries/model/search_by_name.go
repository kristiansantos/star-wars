package model

import (
	"encoding/json"
	"fmt"
)

func SearchByNameQuery(name string) (queryResult map[string]interface{}) {
	query := fmt.Sprintf(`
		{
			"query": {
				"match_phrase": {
					"name": "%s"
				}
			}
		}
	`, name)

	_ = json.Unmarshal([]byte(query), &queryResult)

	return
}
