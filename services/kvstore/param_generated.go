/*
 * Copyright © 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 *
 * KV Store API
 *
 * With the Splunk Cloud KV store service, you can save and retrieve data within your Splunk apps, enabling you to manage and maintain the state of the application.
 *
 * API version: v1beta1.2 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package kvstore

// DeleteRecordsQueryParams represents valid query parameters for the DeleteRecords operation
// For convenience DeleteRecordsQueryParams can be formed in a single statement, for example:
//     `v := DeleteRecordsQueryParams{}.SetQuery(...)`
type DeleteRecordsQueryParams struct {
	// Query : Query JSON expression.
	Query string `key:"query"`
}

func (q DeleteRecordsQueryParams) SetQuery(v string) DeleteRecordsQueryParams {
	q.Query = v
	return q
}

// ListRecordsQueryParams represents valid query parameters for the ListRecords operation
// For convenience ListRecordsQueryParams can be formed in a single statement, for example:
//     `v := ListRecordsQueryParams{}.SetCount(...).SetFields(...).SetFilters(...).SetOffset(...).SetOrderby(...)`
type ListRecordsQueryParams struct {
	// Count : Maximum number of records to return.
	Count *int32 `key:"count"`
	// Fields : Comma-separated list of fields to include or exclude.
	Fields  []string               `key:"fields"`
	Filters map[string]interface{} `key:"filters"`
	// Offset : Number of records to skip from the start.
	Offset *int32 `key:"offset"`
	// Orderby : Sort order. Format is &#x60;&lt;field&gt;:&lt;sort order&gt;&#x60;. Valid sort orders are 1 for ascending, -1 for descending.
	Orderby []string `key:"orderby"`
}

func (q ListRecordsQueryParams) SetCount(v int32) ListRecordsQueryParams {
	q.Count = &v
	return q
}

func (q ListRecordsQueryParams) SetFields(v []string) ListRecordsQueryParams {
	q.Fields = v
	return q
}

func (q ListRecordsQueryParams) SetFilters(v map[string]interface{}) ListRecordsQueryParams {
	q.Filters = v
	return q
}

func (q ListRecordsQueryParams) SetOffset(v int32) ListRecordsQueryParams {
	q.Offset = &v
	return q
}

func (q ListRecordsQueryParams) SetOrderby(v []string) ListRecordsQueryParams {
	q.Orderby = v
	return q
}

// QueryRecordsQueryParams represents valid query parameters for the QueryRecords operation
// For convenience QueryRecordsQueryParams can be formed in a single statement, for example:
//     `v := QueryRecordsQueryParams{}.SetCount(...).SetFields(...).SetOffset(...).SetOrderby(...).SetQuery(...)`
type QueryRecordsQueryParams struct {
	// Count : Maximum number of records to return.
	Count *int32 `key:"count"`
	// Fields : Comma-separated list of fields to include or exclude.
	Fields []string `key:"fields"`
	// Offset : Number of records to skip from the start.
	Offset *int32 `key:"offset"`
	// Orderby : Sort order. Format is &#x60;&lt;field&gt;:&lt;sort order&gt;&#x60;. Valid sort orders are 1 for ascending, -1 for descending.
	Orderby []string `key:"orderby"`
	// Query : Query JSON expression.
	Query string `key:"query"`
}

func (q QueryRecordsQueryParams) SetCount(v int32) QueryRecordsQueryParams {
	q.Count = &v
	return q
}

func (q QueryRecordsQueryParams) SetFields(v []string) QueryRecordsQueryParams {
	q.Fields = v
	return q
}

func (q QueryRecordsQueryParams) SetOffset(v int32) QueryRecordsQueryParams {
	q.Offset = &v
	return q
}

func (q QueryRecordsQueryParams) SetOrderby(v []string) QueryRecordsQueryParams {
	q.Orderby = v
	return q
}

func (q QueryRecordsQueryParams) SetQuery(v string) QueryRecordsQueryParams {
	q.Query = v
	return q
}
