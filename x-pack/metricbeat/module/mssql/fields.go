// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package mssql

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "mssql", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJyslEGPmzAQhe/8iqe97x/gUKnSXiq1hyq9R4MZghtjs55ho/TXVzghCwGiZJXcMn5+/ubNiFfs+ZijEXl3GaBWHed4Sf9fMqBkMdG2aoPP8WuDze+faELZOc6AyI5JOEfBShlQWXal5BkAvMJTw5/O/U+PLefYxdC158rE/tu5OLYZW5XFpbTkNfMrCzSs0RphxcE6h4rV1LC+CrGhXgQqQqdgMnUvr2JoQJB3B+H4wXFkfU01ISOlgoQnh2uMM863820cao4MrXngRk0fjILZn9C5TIhXZktkYzpbzo6uCP7UfOkBP94W5KdW9nw8hFguZuDCbistGd52QrsvRrHpDZAMJmOqQky5aCQvZFLNhd2DQWhQcqvNLcHNk+otIJ+YCxfWMMYoxVFZFhULoxH7jxGqlMC862kXLvi5YHi1E17ahUf6rxnBmK61XN7B9fQo1t6G9TeNboQzZmmN3kPyHS1Hw17p7wVilYwENOiHetrEdNqrbzJXLtAcagAW6w1vHYluCzL7+wZITehOLMMiJ2xvTh+f3i2xrTreXpinD/0x3gdW4X8AAAD//45S2Xs="
}