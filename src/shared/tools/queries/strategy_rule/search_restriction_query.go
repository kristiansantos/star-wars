package strategyRule

import (
	"encoding/json"
	"fmt"
)

func SearchRestrictionsByPageAndPlacementNameQuery(placement, page string) (queryResult map[string]interface{}) {
	query := fmt.Sprintf(`
		{
			"query": {
				"bool": {
					"should": [
						{ "match": { "context.rules.placements": "%s" } },
						{ "match": { "context.rules.pages": "%s" } },
						{ "match": { "context.type": "sitewide" } }
					]
				}
			},
			"post_filter": {
				"term": {
					"active": true
				}
			}
		}
	`, placement, page)

	_ = json.Unmarshal([]byte(query), &queryResult)

	return
}
