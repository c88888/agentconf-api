package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/didi/nightingale/v5/src/server/common/conv"
	"github.com/prometheus/common/model"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func InstantQuery(querType, promql, start, end, step string) (map[string]float64, error) {
	ret := make(map[string]float64)
	val, warnings, err := executePrometheusQuery(querType, promql, start, end, step)
	if err != nil {
		return ret, err
	}

	if len(warnings) > 0 {
		return ret, fmt.Errorf("instant query occur warnings, promql: %s, warnings: %v", promql, warnings)
	}

	vectors := conv.ConvertVectors(val)
	for i := range vectors {
		ident, has := vectors[i].Labels["ident"]
		if has {
			ret[string(ident)] = vectors[i].Value
		}
	}
	return ret, nil
}

/*
	{
	    "status": "success",
	    "isPartial": false,
	    "data": {
	        "resultType": "vector",
	        "result": [
	            {
	                "metric": {
	                    "ident": "zy-js-pro-yiqiyoo-db-01"
	                },
	                "value": [1729671950, "57.1178208009"]
	            },
	            {
	                "metric": {
	                    "ident": "zy-wh-dev-yiqiyoo-game-01"
	                },
	                "value": [1729671950, "19.322362526099994"]
	            },
	            {
	                "metric": {
	                    "ident": "zy-wy-dev-yiqiyoo-db-01"
	                },
	                "value": [1729671950, "84.155066996228"]
	            }
	        ]
	    }
	}
*/
type PromQueryResult struct {
	Status    string   `json:"status"`
	IsPartial bool     `json:"isPartial"`
	Warnings  []string `json:"warnings,omitempty"`
	Data      struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"` // 包含各个指标的标签
			Value  []interface{}     `json:"value"`  // 用于 vector 类型的结果
			Values [][]interface{}   `json:"values"` // 用于 matrix 类型的结果
		} `json:"result"` // 结果集，包含多个metric和其值
	} `json:"data"`
}

func executePrometheusQuery(queryType, query, start, end, step string) (model.Value, []string, error) {
	targetURL := global.GVA_CONFIG.Prom.Host + fmt.Sprintf("/api/v1/%s?query=%s&start=%s&end=%s&step=%s", queryType, url.QueryEscape(query), start, end, step)
	resp, err := http.Get(targetURL)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var result PromQueryResult
	if err := json.Unmarshal(responseBody, &result); err != nil {
		return nil, nil, err
	}

	if result.Status != "success" {
		return nil, result.Warnings, fmt.Errorf("query failed with status: %s", result.Status)
	}

	var value model.Value

	switch result.Data.ResultType {
	case "vector":
		var vector model.Vector
		for _, r := range result.Data.Result {
			sample := model.Sample{
				Metric: make(model.Metric),
			}
			for k, v := range r.Metric {
				sample.Metric[model.LabelName(k)] = model.LabelValue(v)
			}

			if len(r.Value) == 2 {
				//r.Value 的第二个元素是字符串形式的数值，需转换为 float64
				timestamp := model.Time(int64(r.Value[0].(float64)) * 1000) // 将秒转换为毫秒
				valueFloat, err := strconv.ParseFloat(r.Value[1].(string), 64)
				if err != nil {
					return nil, nil, fmt.Errorf("failed to parse value: %v", err)
				}
				sample.Value = model.SampleValue(valueFloat)
				sample.Timestamp = timestamp
			}
			vector = append(vector, &sample)
		}
		value = vector

	case "matrix":
		var matrix model.Matrix
		for _, r := range result.Data.Result {
			stream := model.SampleStream{
				Metric: make(model.Metric),
			}
			for k, v := range r.Metric {
				stream.Metric[model.LabelName(k)] = model.LabelValue(v)
			}

			for _, val := range r.Values {
				if len(val) == 2 {
					timestamp := model.Time(int64(val[0].(float64)) * 1000)
					valueFloat, err := strconv.ParseFloat(val[1].(string), 64)
					if err != nil {
						return nil, nil, fmt.Errorf("failed to parse value: %v", err)
					}
					stream.Values = append(stream.Values, model.SamplePair{
						Timestamp: timestamp,
						Value:     model.SampleValue(valueFloat),
					})
				}
			}
			matrix = append(matrix, &stream)
		}
		value = matrix

	default:
		return nil, result.Warnings, fmt.Errorf("unsupported result type: %s", result.Data.ResultType)
	}
	return value, result.Warnings, nil
}
