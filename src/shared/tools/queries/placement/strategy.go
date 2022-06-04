package placement

import (
	"encoding/json"
	"fmt"
)

func StrategyQuery(minItem, maxItem int) (queryResult map[string]interface{}) {
	query := fmt.Sprintf(`
		{
			"placement": {
				"minItem": %d,
				"maxItem": %d
			}
		}
	`, minItem, maxItem)

	_ = json.Unmarshal([]byte(query), &queryResult)

	return
}
