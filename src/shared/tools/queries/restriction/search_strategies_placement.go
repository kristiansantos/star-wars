package restriction

import (
	"encoding/json"
	"fmt"
)

func SearchStrategiesByPlacementNameQuery(placementId string) (queryResult map[string]interface{}) {
	query := fmt.Sprintf(`
		{
			"query": {
				"match": {
					"_id": "%s"
				}
			}
		}
	`, placementId)

	_ = json.Unmarshal([]byte(query), &queryResult)

	return
}
