package model

import (
	"encoding/json"
	"fmt"
)

func StrategyQuery(modelId string) (queryResult map[string]interface{}) {
	query := fmt.Sprintf(`
		{
			"query": {
				"match": {
					"models.modelRef": "%s"
				}
			}
		}
	`, modelId)

	_ = json.Unmarshal([]byte(query), &queryResult)

	return
}
