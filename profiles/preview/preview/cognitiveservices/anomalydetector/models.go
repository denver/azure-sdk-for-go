// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package anomalydetector

import original "github.com/Azure/azure-sdk-for-go/services/preview/cognitiveservices/v1.0/anomalydetector"

type Granularity = original.Granularity

const (
	Daily    Granularity = original.Daily
	Hourly   Granularity = original.Hourly
	Minutely Granularity = original.Minutely
	Monthly  Granularity = original.Monthly
	Weekly   Granularity = original.Weekly
	Yearly   Granularity = original.Yearly
)

type APIError = original.APIError
type BaseClient = original.BaseClient
type ChangePointDetectRequest = original.ChangePointDetectRequest
type ChangePointDetectResponse = original.ChangePointDetectResponse
type EntireDetectResponse = original.EntireDetectResponse
type LastDetectResponse = original.LastDetectResponse
type Point = original.Point
type Request = original.Request

func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleGranularityValues() []Granularity {
	return original.PossibleGranularityValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
