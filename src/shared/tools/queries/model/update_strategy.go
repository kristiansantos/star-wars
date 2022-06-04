package model

import (
	"encoding/json"
	"fmt"
)

func UpdateStrategyModelQuery(index int, modelId, indexKey, esIndex, id string) (queryResult map[string]interface{}, err error) {
	sourceModelId := fmt.Sprintf("ctx._source.models[%d].modelId = params.modelId", index)
	sourceIndexKey := fmt.Sprintf("ctx._source.models[%d].indexKey = params.indexKey", index)
	sourceEsIndex := fmt.Sprintf("ctx._source.models[%d].esIndex = params.esIndex", index)

	source := fmt.Sprintf("%s;%s;%s;", sourceModelId, sourceIndexKey, sourceEsIndex)

	query := fmt.Sprintf(`
		{
			"script": {
				"lang": "painless",
				"source": "%s",
				"params": {
					"modelId": "%s",
					"indexKey": "%s",
					"esIndex": "%s"
				}
			},
			"query": {
				"match": {
					"_id": "%s"
				}
			}
		}
	`, source, modelId, indexKey, esIndex, id)

	err = json.Unmarshal([]byte(query), &queryResult)

	return
}

func UpdateStrategyFallbackQuery(modelId, indexKey, esIndex, id string) (queryResult map[string]interface{}, err error) {
	sourceModelId := "ctx._source.fallback.modelId = params.modelId"
	sourceIndexKey := "ctx._source.fallback.indexKey = params.indexKey"
	sourceEsIndex := "ctx._source.fallback.esIndex = params.esIndex"

	source := fmt.Sprintf("%s;%s;%s;", sourceModelId, sourceIndexKey, sourceEsIndex)

	query := fmt.Sprintf(`
		{
			"script": {
				"lang": "painless",
				"source": "%s",
				"params": {
					"modelId": "%s",
					"indexKey": "%s",
					"esIndex": "%s"
				}
			},
			"query": {
				"match": {
					"fallback.modelRef": "%s"
				}
			}
		}
	`, source, modelId, indexKey, esIndex, id)

	err = json.Unmarshal([]byte(query), &queryResult)

	return
}
