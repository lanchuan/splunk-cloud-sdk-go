// Package identity -- generated by scloudgen
// !! DO NOT EDIT !!
//
package identity

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/flags"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/jsonx"
	model "github.com/splunk/splunk-cloud-sdk-go/services/identity"
)

// AddGroupMember Adds a member to a given group.
func AddGroupMember(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}
	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.AddGroupMemberBody{

		Name: name,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.AddGroupMember(group, generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// AddGroupRole Adds a role to a given group.
func AddGroupRole(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}
	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.AddGroupRoleBody{

		Name: name,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.AddGroupRole(group, generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// AddMember Adds a member to a given tenant.
func AddMember(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.AddMemberBody{

		Name: name,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.AddMember(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// AddPrincipalPublicKey Add service principal public key
func AddPrincipalPublicKey(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}
	// Parse all flags

	var algDefault model.EcJwkAlg
	alg := &algDefault
	err = flags.ParseFlag(cmd.Flags(), "alg", &alg)
	if err != nil {
		return fmt.Errorf(`error parsing "alg": ` + err.Error())
	}
	var crvDefault string
	crv := &crvDefault
	err = flags.ParseFlag(cmd.Flags(), "crv", &crv)
	if err != nil {
		return fmt.Errorf(`error parsing "crv": ` + err.Error())
	}
	var dDefault string
	d := &dDefault
	err = flags.ParseFlag(cmd.Flags(), "d", &d)
	if err != nil {
		return fmt.Errorf(`error parsing "d": ` + err.Error())
	}
	var kidDefault string
	kid := &kidDefault
	err = flags.ParseFlag(cmd.Flags(), "kid", &kid)
	if err != nil {
		return fmt.Errorf(`error parsing "kid": ` + err.Error())
	}
	var ktyDefault model.EcJwkKty
	kty := &ktyDefault
	err = flags.ParseFlag(cmd.Flags(), "kty", &kty)
	if err != nil {
		return fmt.Errorf(`error parsing "kty": ` + err.Error())
	}
	var principal string
	err = flags.ParseFlag(cmd.Flags(), "principal", &principal)
	if err != nil {
		return fmt.Errorf(`error parsing "principal": ` + err.Error())
	}
	var xDefault string
	x := &xDefault
	err = flags.ParseFlag(cmd.Flags(), "x", &x)
	if err != nil {
		return fmt.Errorf(`error parsing "x": ` + err.Error())
	}
	var yDefault string
	y := &yDefault
	err = flags.ParseFlag(cmd.Flags(), "y", &y)
	if err != nil {
		return fmt.Errorf(`error parsing "y": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.EcJwk{

		Alg: alg,
		Crv: crv,
		D:   d,
		Kid: kid,
		Kty: kty,
		X:   x,
		Y:   y,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.AddPrincipalPublicKey(principal, generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// AddRolePermission Adds permissions to a role in a given tenant.
func AddRolePermission(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var permission string
	err = flags.ParseFlag(cmd.Flags(), "permission", &permission)
	if err != nil {
		return fmt.Errorf(`error parsing "permission": ` + err.Error())
	}
	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.AddRolePermissionBody{

		Permission: permission,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.AddRolePermission(role, generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// CreateGroup Creates a new group in a given tenant.
func CreateGroup(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.CreateGroupBody{

		Name: name,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.CreateGroup(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// CreatePrincipal Create a new principal
func CreatePrincipal(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}
	// Parse all flags

	var algDefault model.EcJwkAlg
	alg := &algDefault
	err = flags.ParseFlag(cmd.Flags(), "alg", &alg)
	if err != nil {
		return fmt.Errorf(`error parsing "alg": ` + err.Error())
	}
	var credentials []model.Credential
	err = flags.ParseFlag(cmd.Flags(), "credentials", &credentials)
	if err != nil {
		return fmt.Errorf(`error parsing "credentials": ` + err.Error())
	}
	var crvDefault string
	crv := &crvDefault
	err = flags.ParseFlag(cmd.Flags(), "crv", &crv)
	if err != nil {
		return fmt.Errorf(`error parsing "crv": ` + err.Error())
	}
	var dDefault string
	d := &dDefault
	err = flags.ParseFlag(cmd.Flags(), "d", &d)
	if err != nil {
		return fmt.Errorf(`error parsing "d": ` + err.Error())
	}
	var email string
	err = flags.ParseFlag(cmd.Flags(), "email", &email)
	if err != nil {
		return fmt.Errorf(`error parsing "email": ` + err.Error())
	}
	var enabledDefault bool
	enabled := &enabledDefault
	err = flags.ParseFlag(cmd.Flags(), "enabled", &enabled)
	if err != nil {
		return fmt.Errorf(`error parsing "enabled": ` + err.Error())
	}
	var firstName string
	err = flags.ParseFlag(cmd.Flags(), "first-name", &firstName)
	if err != nil {
		return fmt.Errorf(`error parsing "first-name": ` + err.Error())
	}
	var inviteId string
	err = flags.ParseFlag(cmd.Flags(), "invite-id", &inviteId)
	if err != nil {
		return fmt.Errorf(`error parsing "invite-id": ` + err.Error())
	}
	var kidDefault string
	kid := &kidDefault
	err = flags.ParseFlag(cmd.Flags(), "kid", &kid)
	if err != nil {
		return fmt.Errorf(`error parsing "kid": ` + err.Error())
	}
	var kind model.PrincipalKind
	err = flags.ParseFlag(cmd.Flags(), "kind", &kind)
	if err != nil {
		return fmt.Errorf(`error parsing "kind": ` + err.Error())
	}
	var ktyDefault model.EcJwkKty
	kty := &ktyDefault
	err = flags.ParseFlag(cmd.Flags(), "kty", &kty)
	if err != nil {
		return fmt.Errorf(`error parsing "kty": ` + err.Error())
	}
	var lastName string
	err = flags.ParseFlag(cmd.Flags(), "last-name", &lastName)
	if err != nil {
		return fmt.Errorf(`error parsing "last-name": ` + err.Error())
	}
	var nameDefault string
	name := &nameDefault
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var xDefault string
	x := &xDefault
	err = flags.ParseFlag(cmd.Flags(), "x", &x)
	if err != nil {
		return fmt.Errorf(`error parsing "x": ` + err.Error())
	}
	var yDefault string
	y := &yDefault
	err = flags.ParseFlag(cmd.Flags(), "y", &y)
	if err != nil {
		return fmt.Errorf(`error parsing "y": ` + err.Error())
	}
	// Form query params
	generated_query := model.CreatePrincipalQueryParams{}
	generated_query.InviteId = inviteId

	// Form the request body
	generated_request_body := model.CreatePrincipalBody{

		Credentials: credentials,
		Enabled:     enabled,
		Key: &model.EcJwk{
			Alg: alg,
			Crv: crv,
			D:   d,
			Kid: kid,
			Kty: kty,
			X:   x,
			Y:   y,
		},
		Kind: kind,
		Name: name,
		Profile: &model.CreatePrincipalProfile{
			Email:     email,
			FirstName: firstName,
			LastName:  lastName,
		},
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.CreatePrincipal(generated_request_body, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// CreateRole Creates a new authorization role in a given tenant.
func CreateRole(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.CreateRoleBody{

		Name: name,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.CreateRole(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// DeleteGroup Deletes a group in a given tenant.
func DeleteGroup(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	err = client.IdentityService.DeleteGroup(group)
	if err != nil {
		return err
	}

	return nil
}

// DeletePrincipalPublicKey Deletes principal public key
func DeletePrincipalPublicKey(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}
	// Parse all flags

	var keyId string
	err = flags.ParseFlag(cmd.Flags(), "key-id", &keyId)
	if err != nil {
		return fmt.Errorf(`error parsing "key-id": ` + err.Error())
	}
	var principal string
	err = flags.ParseFlag(cmd.Flags(), "principal", &principal)
	if err != nil {
		return fmt.Errorf(`error parsing "principal": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	err = client.IdentityService.DeletePrincipalPublicKey(principal, keyId)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRole Deletes a defined role for a given tenant.
func DeleteRole(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	err = client.IdentityService.DeleteRole(role)
	if err != nil {
		return err
	}

	return nil
}

// GetGroup Returns information about a given group within a tenant.
func GetGroup(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.GetGroup(group)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetGroupMember Returns information about a given member within a given group.
func GetGroupMember(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}
	var member string
	err = flags.ParseFlag(cmd.Flags(), "member", &member)
	if err != nil {
		return fmt.Errorf(`error parsing "member": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.GetGroupMember(group, member)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetGroupRole Returns information about a given role within a given group.
func GetGroupRole(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}
	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.GetGroupRole(group, role)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetMember Returns a member of a given tenant.
func GetMember(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var member string
	err = flags.ParseFlag(cmd.Flags(), "member", &member)
	if err != nil {
		return fmt.Errorf(`error parsing "member": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.GetMember(member)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetPrincipal Returns the details of a principal, including its tenant membership and any relevant profile information.

func GetPrincipal(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}
	// Parse all flags

	var principal string
	err = flags.ParseFlag(cmd.Flags(), "principal", &principal)
	if err != nil {
		return fmt.Errorf(`error parsing "principal": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.GetPrincipal(principal)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetPrincipalPublicKey Returns principal public key
func GetPrincipalPublicKey(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}
	// Parse all flags

	var keyId string
	err = flags.ParseFlag(cmd.Flags(), "key-id", &keyId)
	if err != nil {
		return fmt.Errorf(`error parsing "key-id": ` + err.Error())
	}
	var principal string
	err = flags.ParseFlag(cmd.Flags(), "principal", &principal)
	if err != nil {
		return fmt.Errorf(`error parsing "principal": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.GetPrincipalPublicKey(principal, keyId)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetPrincipalPublicKeys Returns principal public keys
func GetPrincipalPublicKeys(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}
	// Parse all flags

	var principal string
	err = flags.ParseFlag(cmd.Flags(), "principal", &principal)
	if err != nil {
		return fmt.Errorf(`error parsing "principal": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.GetPrincipalPublicKeys(principal)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetRole Returns a role for a given tenant.
func GetRole(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.GetRole(role)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetRolePermission Gets a permission for the specified role.
func GetRolePermission(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var permission string
	err = flags.ParseFlag(cmd.Flags(), "permission", &permission)
	if err != nil {
		return fmt.Errorf(`error parsing "permission": ` + err.Error())
	}
	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.GetRolePermission(role, permission)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListGroupMembers Returns a list of the members within a given group.
func ListGroupMembers(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}
	var orderby string
	err = flags.ParseFlag(cmd.Flags(), "orderby", &orderby)
	if err != nil {
		return fmt.Errorf(`error parsing "orderby": ` + err.Error())
	}
	var page_sizeDefault int32
	page_size := &page_sizeDefault
	err = flags.ParseFlag(cmd.Flags(), "page-size", &page_size)
	if err != nil {
		return fmt.Errorf(`error parsing "page-size": ` + err.Error())
	}
	var page_token string
	err = flags.ParseFlag(cmd.Flags(), "page-token", &page_token)
	if err != nil {
		return fmt.Errorf(`error parsing "page-token": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListGroupMembersQueryParams{}
	generated_query.Orderby = orderby
	generated_query.PageSize = page_size
	generated_query.PageToken = page_token

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.ListGroupMembers(group, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListGroupRoles Returns a list of the roles that are attached to a group within a given tenant.
func ListGroupRoles(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}
	var orderby string
	err = flags.ParseFlag(cmd.Flags(), "orderby", &orderby)
	if err != nil {
		return fmt.Errorf(`error parsing "orderby": ` + err.Error())
	}
	var page_sizeDefault int32
	page_size := &page_sizeDefault
	err = flags.ParseFlag(cmd.Flags(), "page-size", &page_size)
	if err != nil {
		return fmt.Errorf(`error parsing "page-size": ` + err.Error())
	}
	var page_token string
	err = flags.ParseFlag(cmd.Flags(), "page-token", &page_token)
	if err != nil {
		return fmt.Errorf(`error parsing "page-token": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListGroupRolesQueryParams{}
	generated_query.Orderby = orderby
	generated_query.PageSize = page_size
	generated_query.PageToken = page_token

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.ListGroupRoles(group, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListGroups List the groups that exist in a given tenant.
func ListGroups(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var access model.ListGroupsaccess
	err = flags.ParseFlag(cmd.Flags(), "access", &access)
	if err != nil {
		return fmt.Errorf(`error parsing "access": ` + err.Error())
	}
	var orderby string
	err = flags.ParseFlag(cmd.Flags(), "orderby", &orderby)
	if err != nil {
		return fmt.Errorf(`error parsing "orderby": ` + err.Error())
	}
	var page_sizeDefault int32
	page_size := &page_sizeDefault
	err = flags.ParseFlag(cmd.Flags(), "page-size", &page_size)
	if err != nil {
		return fmt.Errorf(`error parsing "page-size": ` + err.Error())
	}
	var page_token string
	err = flags.ParseFlag(cmd.Flags(), "page-token", &page_token)
	if err != nil {
		return fmt.Errorf(`error parsing "page-token": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListGroupsQueryParams{}
	generated_query.Access = access
	generated_query.Orderby = orderby
	generated_query.PageSize = page_size
	generated_query.PageToken = page_token

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.ListGroups(&generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListMemberGroups Returns a list of groups that a member belongs to within a tenant.
func ListMemberGroups(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var member string
	err = flags.ParseFlag(cmd.Flags(), "member", &member)
	if err != nil {
		return fmt.Errorf(`error parsing "member": ` + err.Error())
	}
	var orderby string
	err = flags.ParseFlag(cmd.Flags(), "orderby", &orderby)
	if err != nil {
		return fmt.Errorf(`error parsing "orderby": ` + err.Error())
	}
	var page_sizeDefault int32
	page_size := &page_sizeDefault
	err = flags.ParseFlag(cmd.Flags(), "page-size", &page_size)
	if err != nil {
		return fmt.Errorf(`error parsing "page-size": ` + err.Error())
	}
	var page_token string
	err = flags.ParseFlag(cmd.Flags(), "page-token", &page_token)
	if err != nil {
		return fmt.Errorf(`error parsing "page-token": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListMemberGroupsQueryParams{}
	generated_query.Orderby = orderby
	generated_query.PageSize = page_size
	generated_query.PageToken = page_token

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.ListMemberGroups(member, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListMemberPermissions Returns a set of permissions granted to the member within the tenant.

func ListMemberPermissions(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var member string
	err = flags.ParseFlag(cmd.Flags(), "member", &member)
	if err != nil {
		return fmt.Errorf(`error parsing "member": ` + err.Error())
	}
	var orderby string
	err = flags.ParseFlag(cmd.Flags(), "orderby", &orderby)
	if err != nil {
		return fmt.Errorf(`error parsing "orderby": ` + err.Error())
	}
	var page_sizeDefault int32
	page_size := &page_sizeDefault
	err = flags.ParseFlag(cmd.Flags(), "page-size", &page_size)
	if err != nil {
		return fmt.Errorf(`error parsing "page-size": ` + err.Error())
	}
	var page_token string
	err = flags.ParseFlag(cmd.Flags(), "page-token", &page_token)
	if err != nil {
		return fmt.Errorf(`error parsing "page-token": ` + err.Error())
	}
	var scope_filter string
	err = flags.ParseFlag(cmd.Flags(), "scope-filter", &scope_filter)
	if err != nil {
		return fmt.Errorf(`error parsing "scope-filter": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListMemberPermissionsQueryParams{}
	generated_query.Orderby = orderby
	generated_query.PageSize = page_size
	generated_query.PageToken = page_token
	generated_query.ScopeFilter = scope_filter

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.ListMemberPermissions(member, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListMemberRoles Returns a set of roles that a given member holds within the tenant.

func ListMemberRoles(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var member string
	err = flags.ParseFlag(cmd.Flags(), "member", &member)
	if err != nil {
		return fmt.Errorf(`error parsing "member": ` + err.Error())
	}
	var orderby string
	err = flags.ParseFlag(cmd.Flags(), "orderby", &orderby)
	if err != nil {
		return fmt.Errorf(`error parsing "orderby": ` + err.Error())
	}
	var page_sizeDefault int32
	page_size := &page_sizeDefault
	err = flags.ParseFlag(cmd.Flags(), "page-size", &page_size)
	if err != nil {
		return fmt.Errorf(`error parsing "page-size": ` + err.Error())
	}
	var page_token string
	err = flags.ParseFlag(cmd.Flags(), "page-token", &page_token)
	if err != nil {
		return fmt.Errorf(`error parsing "page-token": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListMemberRolesQueryParams{}
	generated_query.Orderby = orderby
	generated_query.PageSize = page_size
	generated_query.PageToken = page_token

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.ListMemberRoles(member, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListMembers Returns a list of members in a given tenant.
func ListMembers(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var kind model.ListMemberskind
	err = flags.ParseFlag(cmd.Flags(), "kind", &kind)
	if err != nil {
		return fmt.Errorf(`error parsing "kind": ` + err.Error())
	}
	var orderby string
	err = flags.ParseFlag(cmd.Flags(), "orderby", &orderby)
	if err != nil {
		return fmt.Errorf(`error parsing "orderby": ` + err.Error())
	}
	var page_sizeDefault int32
	page_size := &page_sizeDefault
	err = flags.ParseFlag(cmd.Flags(), "page-size", &page_size)
	if err != nil {
		return fmt.Errorf(`error parsing "page-size": ` + err.Error())
	}
	var page_token string
	err = flags.ParseFlag(cmd.Flags(), "page-token", &page_token)
	if err != nil {
		return fmt.Errorf(`error parsing "page-token": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListMembersQueryParams{}
	generated_query.Kind = kind
	generated_query.Orderby = orderby
	generated_query.PageSize = page_size
	generated_query.PageToken = page_token

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.ListMembers(&generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListPrincipals Returns the list of principals that the Identity service knows about.

func ListPrincipals(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}
	// Parse all flags

	var orderby string
	err = flags.ParseFlag(cmd.Flags(), "orderby", &orderby)
	if err != nil {
		return fmt.Errorf(`error parsing "orderby": ` + err.Error())
	}
	var page_sizeDefault int32
	page_size := &page_sizeDefault
	err = flags.ParseFlag(cmd.Flags(), "page-size", &page_size)
	if err != nil {
		return fmt.Errorf(`error parsing "page-size": ` + err.Error())
	}
	var page_token string
	err = flags.ParseFlag(cmd.Flags(), "page-token", &page_token)
	if err != nil {
		return fmt.Errorf(`error parsing "page-token": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListPrincipalsQueryParams{}
	generated_query.Orderby = orderby
	generated_query.PageSize = page_size
	generated_query.PageToken = page_token

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.ListPrincipals(&generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListRoleGroups Gets a list of groups for a role in a given tenant.
func ListRoleGroups(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var orderby string
	err = flags.ParseFlag(cmd.Flags(), "orderby", &orderby)
	if err != nil {
		return fmt.Errorf(`error parsing "orderby": ` + err.Error())
	}
	var page_sizeDefault int32
	page_size := &page_sizeDefault
	err = flags.ParseFlag(cmd.Flags(), "page-size", &page_size)
	if err != nil {
		return fmt.Errorf(`error parsing "page-size": ` + err.Error())
	}
	var page_token string
	err = flags.ParseFlag(cmd.Flags(), "page-token", &page_token)
	if err != nil {
		return fmt.Errorf(`error parsing "page-token": ` + err.Error())
	}
	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListRoleGroupsQueryParams{}
	generated_query.Orderby = orderby
	generated_query.PageSize = page_size
	generated_query.PageToken = page_token

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.ListRoleGroups(role, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListRolePermissions Gets the permissions for a role in a given tenant.
func ListRolePermissions(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var orderby string
	err = flags.ParseFlag(cmd.Flags(), "orderby", &orderby)
	if err != nil {
		return fmt.Errorf(`error parsing "orderby": ` + err.Error())
	}
	var page_sizeDefault int32
	page_size := &page_sizeDefault
	err = flags.ParseFlag(cmd.Flags(), "page-size", &page_size)
	if err != nil {
		return fmt.Errorf(`error parsing "page-size": ` + err.Error())
	}
	var page_token string
	err = flags.ParseFlag(cmd.Flags(), "page-token", &page_token)
	if err != nil {
		return fmt.Errorf(`error parsing "page-token": ` + err.Error())
	}
	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListRolePermissionsQueryParams{}
	generated_query.Orderby = orderby
	generated_query.PageSize = page_size
	generated_query.PageToken = page_token

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.ListRolePermissions(role, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListRoles Returns all roles for a given tenant.
func ListRoles(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var orderby string
	err = flags.ParseFlag(cmd.Flags(), "orderby", &orderby)
	if err != nil {
		return fmt.Errorf(`error parsing "orderby": ` + err.Error())
	}
	var page_sizeDefault int32
	page_size := &page_sizeDefault
	err = flags.ParseFlag(cmd.Flags(), "page-size", &page_size)
	if err != nil {
		return fmt.Errorf(`error parsing "page-size": ` + err.Error())
	}
	var page_token string
	err = flags.ParseFlag(cmd.Flags(), "page-token", &page_token)
	if err != nil {
		return fmt.Errorf(`error parsing "page-token": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListRolesQueryParams{}
	generated_query.Orderby = orderby
	generated_query.PageSize = page_size
	generated_query.PageToken = page_token

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.ListRoles(&generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// RemoveGroupMember Removes the member from a given group.
func RemoveGroupMember(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}
	var member string
	err = flags.ParseFlag(cmd.Flags(), "member", &member)
	if err != nil {
		return fmt.Errorf(`error parsing "member": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	err = client.IdentityService.RemoveGroupMember(group, member)
	if err != nil {
		return err
	}

	return nil
}

// RemoveGroupRole Removes a role from a given group.
func RemoveGroupRole(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}
	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	err = client.IdentityService.RemoveGroupRole(group, role)
	if err != nil {
		return err
	}

	return nil
}

// RemoveMember Removes a member from a given tenant
func RemoveMember(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var member string
	err = flags.ParseFlag(cmd.Flags(), "member", &member)
	if err != nil {
		return fmt.Errorf(`error parsing "member": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	err = client.IdentityService.RemoveMember(member)
	if err != nil {
		return err
	}

	return nil
}

// RemoveRolePermission Removes a permission from the role.
func RemoveRolePermission(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var permission string
	err = flags.ParseFlag(cmd.Flags(), "permission", &permission)
	if err != nil {
		return fmt.Errorf(`error parsing "permission": ` + err.Error())
	}
	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	err = client.IdentityService.RemoveRolePermission(role, permission)
	if err != nil {
		return err
	}

	return nil
}

// RevokePrincipalAuthTokens Revoke all existing access tokens issued to a principal. Principals can reset their password by visiting https://login.splunk.com/en_us/page/lost_password

func RevokePrincipalAuthTokens(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}
	// Parse all flags

	var principal string
	err = flags.ParseFlag(cmd.Flags(), "principal", &principal)
	if err != nil {
		return fmt.Errorf(`error parsing "principal": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	err = client.IdentityService.RevokePrincipalAuthTokens(principal)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePrincipalPublicKey Update principal public key
func UpdatePrincipalPublicKey(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}
	// Parse all flags

	var keyId string
	err = flags.ParseFlag(cmd.Flags(), "key-id", &keyId)
	if err != nil {
		return fmt.Errorf(`error parsing "key-id": ` + err.Error())
	}
	var principal string
	err = flags.ParseFlag(cmd.Flags(), "principal", &principal)
	if err != nil {
		return fmt.Errorf(`error parsing "principal": ` + err.Error())
	}
	var status model.PrincipalPublicKeyStatusBodyStatus
	err = flags.ParseFlag(cmd.Flags(), "status", &status)
	if err != nil {
		return fmt.Errorf(`error parsing "status": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.PrincipalPublicKeyStatusBody{

		Status: status,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.UpdatePrincipalPublicKey(principal, keyId, generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ValidateToken Validates the access token obtained from the authorization header and returns the principal name and tenant memberships.

func ValidateToken(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var include model.ValidateTokenincludeEnum
	err = flags.ParseFlag(cmd.Flags(), "include", &include)
	if err != nil {
		return fmt.Errorf(`error parsing "include": ` + err.Error())
	}
	// Form query params
	generated_query := model.ValidateTokenQueryParams{}

	generated_query.Include = []model.ValidateTokenincludeEnum{include}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.IdentityService.ValidateToken(&generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}
