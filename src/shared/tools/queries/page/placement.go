package page

import (
	"encoding/json"
	"fmt"
)

func UpdatePlacementNameQuery(name, pageType string) (queryResult map[string]interface{}, err error) {
	source := fmt.Sprintf("ctx._source.page.name = '%s'", name)

	query := fmt.Sprintf(`
		{
			"script": {
				"lang": "painless",
				"source": "%s"
			},
			"query": {
				"match": {
					"page.type": "%s"
				}
			}
		}
	`, source, pageType)

	err = json.Unmarshal([]byte(query), &queryResult)

	return
}
