package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pretorian41/goaggregate/models"
	"github.com/pretorian41/goaggregate/utils"
)

func FetchFromAPI(source string, id string) models.ApiResult {
	resp, err := http.Get(fmt.Sprintf("http://%s?id=%s", source, id))

	if err != nil || 200 != resp.StatusCode{
		return models.ApiResult{
			Source: source,
			Data:   make(map[int]struct{}),
		}
	}

	defer resp.Body.Close() // Ensure the response body is closed

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	utils.InfoLog.Println("response: " + string(body))

	return models.ApiResult{
		Source: source,
		Data:   data,
	}
}

type Data map[string]string
type Priority map[string]int

func Reduce(results <-chan models.ApiResult) map[string]interface{} {
	merged := make(map[string]interface{})
	weights := make(map[string]int)
	utils.InfoLog.Println("Reduce started")

	count := 0
	for res := range results {
		utils.InfoLog.Printf("Reducing result from %s", res.Source)
		data, ok := res.Data.(map[string]interface{})
		if !ok {
			utils.InfoLog.Printf("Invalid data type from %s", res.Source)
			continue
		}
		for field, val := range data {
			weightCheck, ok := res.Priorities[field]
			if !ok {
				weightCheck = 100
			}

			_, exists := weights[field]

			if !exists || weights[field] > weightCheck {
				merged[field] = val
				weights[field] = weightCheck
			}
		}
		count++
	}

	utils.InfoLog.Printf("Reduce done, processed %d results, total %d fields", count, len(merged))

	return merged
}
