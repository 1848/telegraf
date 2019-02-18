package mvalue

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/metric"
)

type MValueParser struct {
	MetricName  string
	DataType    string
	DefaultTags map[string]string
	IgnoreBegin []string
	Separator   string
}

func (v *MValueParser) Parse(buf []byte) ([]telegraf.Metric, error) {
	vStr := string(bytes.TrimSpace(bytes.Trim(buf, "\x00")))

	metrics := make([]telegraf.Metric, 0)

	values := strings.Split(vStr, "\n")
	for _, value := range values {
		var cont = false

		for _, ign := range v.IgnoreBegin {
			if strings.HasPrefix(value, ign) {
				cont = true
				break
			}
		}
		if cont {
			continue
		}

		kv := strings.Split(value, v.Separator)
		if len(kv) > 1 {
			var m telegraf.Metric

			cv, err := strconv.ParseFloat(kv[len(kv)-1], 64)
			if err == nil {
				m, _ = metric.New(v.MetricName, v.DefaultTags, map[string]interface{}{kv[0]: cv}, time.Now().UTC())
			} else {
				m, _ = metric.New(v.MetricName, v.DefaultTags, map[string]interface{}{kv[0]: kv[len(kv)-1]}, time.Now().UTC())
			}

			metrics = append(metrics, m)
		}
	}

	return metrics, nil
}

func (v *MValueParser) ParseLine(line string) (telegraf.Metric, error) {
	metrics, err := v.Parse([]byte(line))

	if err != nil {
		return nil, err
	}

	if len(metrics) < 1 {
		return nil, fmt.Errorf("Can not parse the line: %s, for data format: value", line)
	}

	return metrics[0], nil
}

func (v *MValueParser) SetDefaultTags(tags map[string]string) {
	v.DefaultTags = tags
}
