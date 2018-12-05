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

// Complete account details needed for the user.
type Account struct {
	// Your username
	Username string `json:"username"`
	// Your password
	Password string `json:"password"`
	// Your phone number in E.164 format.
	UserPhone string `json:"user_phone"`
	// Your email
	UserEmail string `json:"user_email"`
	// Your first name
	UserFirstName string `json:"user_first_name"`
	// Your last name
	UserLastName string `json:"user_last_name"`
	// Your delivery to value.
	AccountName string `json:"account_name"`
	// Your country
	Country string `json:"country"`
}
