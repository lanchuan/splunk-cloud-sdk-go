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

package provisioner

import (
	"net/http"
)

// Servicer represents the interface for implementing all endpoints for this service
type Servicer interface {
	/*
		CreateInvite - prov service endpoint
		Creates an invite to invite a person to the tenant using their email address
		Parameters:
			inviteBody
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateInvite(inviteBody InviteBody, resp ...*http.Response) (*InviteInfo, error)
	/*
		CreateProvisionJob - prov service endpoint
		Creates a new job that provisions a new tenant and subscribes apps to the tenant
		Parameters:
			createProvisionJobBody
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateProvisionJob(createProvisionJobBody CreateProvisionJobBody, resp ...*http.Response) (*ProvisionJobInfo, error)
	/*
		DeleteInvite - prov service endpoint
		Deletes an invite in the given tenant
		Parameters:
			inviteId
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteInvite(inviteId string, resp ...*http.Response) error
	/*
		GetInvite - prov service endpoint
		Gets an invite in the given tenant
		Parameters:
			inviteId
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetInvite(inviteId string, resp ...*http.Response) (*InviteInfo, error)
	/*
		GetProvisionJob - prov service endpoint
		Gets details of a specific provision job
		Parameters:
			jobId
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetProvisionJob(jobId string, resp ...*http.Response) (*ProvisionJobInfo, error)
	/*
		GetTenant - prov service endpoint
		Gets a specific tenant
		Parameters:
			tenantName
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetTenant(tenantName string, resp ...*http.Response) (*TenantInfo, error)
	/*
		ListInvites - prov service endpoint
		Lists the invites in a given tenant
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListInvites(resp ...*http.Response) (*Invites, error)
	/*
		ListProvisionJobs - prov service endpoint
		Lists all provision jobs created by the user
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListProvisionJobs(resp ...*http.Response) (*ProvisionJobs, error)
	/*
		ListTenants - prov service endpoint
		Lists all tenants that the user can read
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListTenants(resp ...*http.Response) (*Tenants, error)
	/*
		UpdateInvite - prov service endpoint
		Updates an invite in the given tenant
		Parameters:
			inviteId
			updateInviteBody
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateInvite(inviteId string, updateInviteBody UpdateInviteBody, resp ...*http.Response) (*InviteInfo, error)
}
