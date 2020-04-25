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
// Invoice struct for Invoice
type Invoice struct {
	Id string `json:"id,omitempty"`
	ChargesTotal float32 `json:"charges_total,omitempty"`
	CreditsTotal float32 `json:"credits_total,omitempty"`
	Total float32 `json:"total,omitempty"`
	PeriodStart string `json:"period_start,omitempty"`
	PeriodEnd string `json:"period_end,omitempty"`
	Payed bool `json:"payed,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
