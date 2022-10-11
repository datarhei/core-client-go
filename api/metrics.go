package api

import (
	"time"
)

type MetricsDescription struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Labels      []string `json:"labels"`
}

type MetricsQueryMetric struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
}

type MetricsQuery struct {
	Timerange int64                `json:"timerange_sec"`
	Interval  int64                `json:"interval_sec"`
	Metrics   []MetricsQueryMetric `json:"metrics"`
}

type MetricsResponseMetric struct {
	Name   string                 `json:"name"`
	Labels map[string]string      `json:"labels"`
	Values []MetricsResponseValue `json:"values"`
}

type MetricsResponseValue struct {
	TS    time.Time `json:"ts"`
	Value float64   `json:"value"`
}

type MetricsResponse struct {
	Timerange int64                   `json:"timerange_sec"`
	Interval  int64                   `json:"interval_sec"`
	Metrics   []MetricsResponseMetric `json:"metrics"`
}
