/*
 * Universe.sh houston API
 *
 * Universe.sh houston API
 *
 * API version: 1.0.0
 * Contact: oss@universe.sh
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Account struct for Account
type Account struct {
	Id string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Email string `json:"email,omitempty"`
	TwoFactorAuthentication bool `json:"two_factor_authentication,omitempty"`
	Verified bool `json:"verified,omitempty"`
	DefaultTeam AccountDefaultTeam `json:"default_team,omitempty"`
	LastLogin string `json:"last_login,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
