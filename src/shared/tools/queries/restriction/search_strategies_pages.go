package restriction

import (
	"encoding/json"
	"fmt"
)

func SearchStrategiesByPageNameQuery(page string) (queryResult map[string]interface{}) {
	query := fmt.Sprintf(`
		{
			"query": {
				"match": {
					"placementType": "%s"
				}
			}
		}
	`, page)

	_ = json.Unmarshal([]byte(query), &queryResult)

	return
}
