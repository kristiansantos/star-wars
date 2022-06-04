package placement

import (
	"encoding/json"
	"fmt"
)

func SearchByNameQuery(name, pageType string) (queryResult map[string]interface{}) {
	query := fmt.Sprintf(`
		{
			"query": {
				"bool": {
					"filter": [
						{"match": {"name": "%s"}},
						{"match": {"page.type": "%s"}}
					]
				}
			}
		}
	`, name, pageType)

	_ = json.Unmarshal([]byte(query), &queryResult)

	return
}
