/*
 * Copyright © 2020 Splunk, Inc.
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
 * Data Stream Processing REST API
 *
 * Use the Streams service to perform create, read, update, and delete (CRUD) operations on your data pipeline. The Streams service also has metrics and preview session endpoints and gives you full control over your data pipeline.
 *
 * API version: v2beta1.4 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package streams

// GetPipelineQueryParams represents valid query parameters for the GetPipeline operation
// For convenience GetPipelineQueryParams can be formed in a single statement, for example:
//     `v := GetPipelineQueryParams{}.SetVersion(...)`
type GetPipelineQueryParams struct {
	// Version : version
	Version string `key:"version"`
}

func (q GetPipelineQueryParams) SetVersion(v string) GetPipelineQueryParams {
	q.Version = v
	return q
}

// GetPipelinesStatusQueryParams represents valid query parameters for the GetPipelinesStatus operation
// For convenience GetPipelinesStatusQueryParams can be formed in a single statement, for example:
//     `v := GetPipelinesStatusQueryParams{}.SetActivated(...).SetCreateUserId(...).SetName(...).SetOffset(...).SetPageSize(...).SetSortDir(...).SetSortField(...)`
type GetPipelinesStatusQueryParams struct {
	// Activated : activated
	Activated *bool `key:"activated"`
	// CreateUserId : createUserId
	CreateUserId string `key:"createUserId"`
	// Name : name
	Name string `key:"name"`
	// Offset : offset
	Offset *int32 `key:"offset"`
	// PageSize : pageSize
	PageSize *int32 `key:"pageSize"`
	// SortDir : sortDir
	SortDir string `key:"sortDir"`
	// SortField : sortField
	SortField string `key:"sortField"`
}

func (q GetPipelinesStatusQueryParams) SetActivated(v bool) GetPipelinesStatusQueryParams {
	q.Activated = &v
	return q
}

func (q GetPipelinesStatusQueryParams) SetCreateUserId(v string) GetPipelinesStatusQueryParams {
	q.CreateUserId = v
	return q
}

func (q GetPipelinesStatusQueryParams) SetName(v string) GetPipelinesStatusQueryParams {
	q.Name = v
	return q
}

func (q GetPipelinesStatusQueryParams) SetOffset(v int32) GetPipelinesStatusQueryParams {
	q.Offset = &v
	return q
}

func (q GetPipelinesStatusQueryParams) SetPageSize(v int32) GetPipelinesStatusQueryParams {
	q.PageSize = &v
	return q
}

func (q GetPipelinesStatusQueryParams) SetSortDir(v string) GetPipelinesStatusQueryParams {
	q.SortDir = v
	return q
}

func (q GetPipelinesStatusQueryParams) SetSortField(v string) GetPipelinesStatusQueryParams {
	q.SortField = v
	return q
}

// GetRegistryQueryParams represents valid query parameters for the GetRegistry operation
// For convenience GetRegistryQueryParams can be formed in a single statement, for example:
//     `v := GetRegistryQueryParams{}.SetLocal(...)`
type GetRegistryQueryParams struct {
	// Local : local
	Local *bool `key:"local"`
}

func (q GetRegistryQueryParams) SetLocal(v bool) GetRegistryQueryParams {
	q.Local = &v
	return q
}

// GetTemplateQueryParams represents valid query parameters for the GetTemplate operation
// For convenience GetTemplateQueryParams can be formed in a single statement, for example:
//     `v := GetTemplateQueryParams{}.SetVersion(...)`
type GetTemplateQueryParams struct {
	// Version : version of the template
	Version *int64 `key:"version"`
}

func (q GetTemplateQueryParams) SetVersion(v int64) GetTemplateQueryParams {
	q.Version = &v
	return q
}

// ListConnectionsQueryParams represents valid query parameters for the ListConnections operation
// For convenience ListConnectionsQueryParams can be formed in a single statement, for example:
//     `v := ListConnectionsQueryParams{}.SetConnectorId(...).SetCreateUserId(...).SetFunctionId(...).SetName(...).SetOffset(...).SetPageSize(...).SetShowSecretNames(...).SetSortDir(...).SetSortField(...)`
type ListConnectionsQueryParams struct {
	ConnectorId     string `key:"connectorId"`
	CreateUserId    string `key:"createUserId"`
	FunctionId      string `key:"functionId"`
	Name            string `key:"name"`
	Offset          *int32 `key:"offset"`
	PageSize        *int32 `key:"pageSize"`
	ShowSecretNames string `key:"showSecretNames"`
	// SortDir : Specify either ascending (&#39;asc&#39;) or descending (&#39;desc&#39;) sort order for a given field (sortField), which must be set for sortDir to apply. Defaults to &#39;asc&#39;.
	SortDir   string `key:"sortDir"`
	SortField string `key:"sortField"`
}

func (q ListConnectionsQueryParams) SetConnectorId(v string) ListConnectionsQueryParams {
	q.ConnectorId = v
	return q
}

func (q ListConnectionsQueryParams) SetCreateUserId(v string) ListConnectionsQueryParams {
	q.CreateUserId = v
	return q
}

func (q ListConnectionsQueryParams) SetFunctionId(v string) ListConnectionsQueryParams {
	q.FunctionId = v
	return q
}

func (q ListConnectionsQueryParams) SetName(v string) ListConnectionsQueryParams {
	q.Name = v
	return q
}

func (q ListConnectionsQueryParams) SetOffset(v int32) ListConnectionsQueryParams {
	q.Offset = &v
	return q
}

func (q ListConnectionsQueryParams) SetPageSize(v int32) ListConnectionsQueryParams {
	q.PageSize = &v
	return q
}

func (q ListConnectionsQueryParams) SetShowSecretNames(v string) ListConnectionsQueryParams {
	q.ShowSecretNames = v
	return q
}

func (q ListConnectionsQueryParams) SetSortDir(v string) ListConnectionsQueryParams {
	q.SortDir = v
	return q
}

func (q ListConnectionsQueryParams) SetSortField(v string) ListConnectionsQueryParams {
	q.SortField = v
	return q
}

// ListPipelinesQueryParams represents valid query parameters for the ListPipelines operation
// For convenience ListPipelinesQueryParams can be formed in a single statement, for example:
//     `v := ListPipelinesQueryParams{}.SetActivated(...).SetCreateUserId(...).SetIncludeData(...).SetName(...).SetOffset(...).SetPageSize(...).SetSortDir(...).SetSortField(...)`
type ListPipelinesQueryParams struct {
	// Activated : activated
	Activated *bool `key:"activated"`
	// CreateUserId : createUserId
	CreateUserId string `key:"createUserId"`
	// IncludeData : includeData
	IncludeData *bool `key:"includeData"`
	// Name : name
	Name string `key:"name"`
	// Offset : offset
	Offset *int32 `key:"offset"`
	// PageSize : pageSize
	PageSize *int32 `key:"pageSize"`
	// SortDir : sortDir
	SortDir string `key:"sortDir"`
	// SortField : sortField
	SortField string `key:"sortField"`
}

func (q ListPipelinesQueryParams) SetActivated(v bool) ListPipelinesQueryParams {
	q.Activated = &v
	return q
}

func (q ListPipelinesQueryParams) SetCreateUserId(v string) ListPipelinesQueryParams {
	q.CreateUserId = v
	return q
}

func (q ListPipelinesQueryParams) SetIncludeData(v bool) ListPipelinesQueryParams {
	q.IncludeData = &v
	return q
}

func (q ListPipelinesQueryParams) SetName(v string) ListPipelinesQueryParams {
	q.Name = v
	return q
}

func (q ListPipelinesQueryParams) SetOffset(v int32) ListPipelinesQueryParams {
	q.Offset = &v
	return q
}

func (q ListPipelinesQueryParams) SetPageSize(v int32) ListPipelinesQueryParams {
	q.PageSize = &v
	return q
}

func (q ListPipelinesQueryParams) SetSortDir(v string) ListPipelinesQueryParams {
	q.SortDir = v
	return q
}

func (q ListPipelinesQueryParams) SetSortField(v string) ListPipelinesQueryParams {
	q.SortField = v
	return q
}

// ListTemplatesQueryParams represents valid query parameters for the ListTemplates operation
// For convenience ListTemplatesQueryParams can be formed in a single statement, for example:
//     `v := ListTemplatesQueryParams{}.SetOffset(...).SetPageSize(...).SetSortDir(...).SetSortField(...)`
type ListTemplatesQueryParams struct {
	// Offset : offset
	Offset *int32 `key:"offset"`
	// PageSize : pageSize
	PageSize *int32 `key:"pageSize"`
	// SortDir : sortDir
	SortDir string `key:"sortDir"`
	// SortField : sortField
	SortField string `key:"sortField"`
}

func (q ListTemplatesQueryParams) SetOffset(v int32) ListTemplatesQueryParams {
	q.Offset = &v
	return q
}

func (q ListTemplatesQueryParams) SetPageSize(v int32) ListTemplatesQueryParams {
	q.PageSize = &v
	return q
}

func (q ListTemplatesQueryParams) SetSortDir(v string) ListTemplatesQueryParams {
	q.SortDir = v
	return q
}

func (q ListTemplatesQueryParams) SetSortField(v string) ListTemplatesQueryParams {
	q.SortField = v
	return q
}
