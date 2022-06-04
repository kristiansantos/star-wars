package restriction

import (
	"encoding/json"
)

func SearchStrategiesSitewideQuery() (queryResult map[string]interface{}) {
	query := `
		{
			"query": {
				"match_all": {}
			}
		}
	`

	_ = json.Unmarshal([]byte(query), &queryResult)

	return
}
