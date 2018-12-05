/*
 * ClickSend v3 API
 *
 *  This is an official SDK for [ClickSend](https://clicksend.com)  Below you will find a current list of the available methods for clicksend.  *NOTE: You will need to create a free account to use the API. You can register [here](https://dashboard.clicksend.com/#/signup/step1/)..* 
 *
 * API version: 3.1
 * Contact: support@clicksend.com
 * Generated by: Swagger Codegen (https://github.com/clicksend-api/clicksend-codegen.git)
 */

package clicksend

// Base model for all address-related objects.
type Address struct {
	// Your address name.
	AddressName string `json:"address_name"`
	// Your address line 1
	AddressLine1 string `json:"address_line_1"`
	// Your city
	AddressCity string `json:"address_city"`
	// Your postal code
	AddressPostalCode string `json:"address_postal_code"`
	// Your country
	AddressCountry string `json:"address_country"`
	// Your address line 2
	AddressLine2 string `json:"address_line_2,omitempty"`
	// Your state
	AddressState string `json:"address_state,omitempty"`
}
