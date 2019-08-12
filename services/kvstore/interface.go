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
 */

// Code generated by gen_interface.go. DO NOT EDIT.

package kvstore

import (
	"net/http"
)

// Servicer represents the interface for implementing all endpoints for this service
type Servicer interface {
	/*
		CreateIndex - Creates an index on a collection.
		Parameters:
			collection: The name of the collection.
			indexDefinition
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateIndex(collection string, indexDefinition IndexDefinition, resp ...*http.Response) (*IndexDescription, error)
	/*
		DeleteIndex - Removes an index from a collection.
		Parameters:
			collection: The name of the collection.
			index: The name of the index.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteIndex(collection string, index string, resp ...*http.Response) error
	/*
		DeleteRecordByKey - Deletes a record with a given key.
		Parameters:
			collection: The name of the collection.
			key: The key of the record.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteRecordByKey(collection string, key string, resp ...*http.Response) error
	/*
		DeleteRecords - Removes records in a collection that match the query.
		Parameters:
			collection: The name of the collection.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteRecords(collection string, query *DeleteRecordsQueryParams, resp ...*http.Response) error
	/*
		GetRecordByKey - Returns a record with a given key.
		Parameters:
			collection: The name of the collection.
			key: The key of the record.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetRecordByKey(collection string, key string, resp ...*http.Response) (*map[string]interface{}, error)
	/*
		InsertRecord - Inserts a record into a collection.
		Parameters:
			collection: The name of the collection.
			body: Record to add to the collection, formatted as a JSON object.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	InsertRecord(collection string, body map[string]interface{}, resp ...*http.Response) (*Key, error)
	/*
		InsertRecords - Inserts multiple records in a single request.
		Parameters:
			collection: The name of the collection.
			requestBody: Array of records to insert.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	InsertRecords(collection string, requestBody []map[string]interface{}, resp ...*http.Response) ([]string, error)
	/*
		ListIndexes - Returns a list of all indexes on a collection.
		Parameters:
			collection: The name of the collection.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListIndexes(collection string, resp ...*http.Response) ([]IndexDefinition, error)
	/*
		ListRecords - Returns a list of records in a collection with basic filtering, sorting, pagination and field projection.
		Use key-value query parameters to filter fields. Fields are implicitly ANDed and values for the same field are implicitly ORed.
		Parameters:
			collection: The name of the collection.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListRecords(collection string, query *ListRecordsQueryParams, resp ...*http.Response) ([]map[string]interface{}, error)
	/*
		Ping - Returns the health status from the database.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	Ping(resp ...*http.Response) (*PingResponse, error)
	/*
		PutRecord - Updates the record with a given key, either by inserting or replacing the record.
		Parameters:
			collection: The name of the collection.
			key: The key of the record.
			ifMatch: Record version identifier.
			body: Record to add to the collection, formatted as a JSON object.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	PutRecord(collection string, key string, ifMatch *string, body map[string]interface{}, resp ...*http.Response) (*Key, error)
	/*
		QueryRecords - Returns a list of query records in a collection.
		Parameters:
			collection: The name of the collection.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	QueryRecords(collection string, query *QueryRecordsQueryParams, resp ...*http.Response) ([]map[string]interface{}, error)
}
