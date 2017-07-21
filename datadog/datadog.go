package datadog

import (
	"os"
	"time"

	"github.com/teddyking/cfbench/bench"
)

type Point [2]int64

type MetricSeries struct {
	Metric string   `json:"metric"`         // phase.ShortName
	Points []Point  `json:"points"`         // single point for this run: time.Now(), duration
	Type   string   `json:"type"`           // gauge
	Host   string   `json:"host,omitempty"` // empty
	Tags   []string `json:"tags,omitempty"` // format key:value, e.g. sha:abcd   or version:1.19.0-rc22
}

type JsonResult struct {
	Series []MetricSeries `json:"series"`
}

func BuildJSONOutput(phases bench.Phases) JsonResult {
	timeOfTest := time.Now().Unix()
	result := JsonResult{}
	for _, phase := range phases {
		if !phase.IsValid() {
			continue
		}
		newSeries := MetricSeries{
			Metric: "cfbench." + phase.ShortName,
			Points: []Point{
				Point{timeOfTest, int64(phase.Duration())},
			},
			Type: "gauge",
			Tags: []string{os.Getenv("CF_TAG")},
		}
		result.Series = append(result.Series, newSeries)
	}
	return result
}
