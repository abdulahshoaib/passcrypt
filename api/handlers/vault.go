package handlers

import ()

func AddToValut() {
	// should recive the passcode as bearer
	// takeout the uploaded body
	// the body can accept:
	//
	//	"service":	{
	//
	// 			"email": "string",
	// 			"username": "string",
	// 			"name": "string",
	// 			"main_password": "string",
	// 			"secondary_password": "string",
	// 			"mobile": "string",
	//
	// 			"2fa_enabled": true,
	// 			"2fa_type": "TOTP/SMS/Email",
	//
	// 			"recovery-codes": ["string"],
	//
	//			"access_token": "string",
	//			"refresh_token": "string",
	//			"api_key": "string",
	//
	//			"client_id": "string",
	//			"client_id": "string",
	//			"client_secret": "string",
	//
	//			"token_type": "Bearer | JWT | Custom",
  //			"scopes": ["read", "write"],
  //			"token_issued_at": "timestamp",
  //			"token_expires_at": "timestamp",
  //			"session_cookie": "string",
	//
  //			"custom_headers": {
  //			  "X-Api-Token": "abcdef",
  //			  "X-Custom": "val"
  //			},
	//
	//			"security_questions": [{
	//	      "question": "string",
	//	      "answer": "string"
	//	  	}],
	//
	//			"tags": ["tags"],
	//
	//			"custom_fields": {
  //			    "key1": "value1",
  //			    "key2": "value2"
  //			}
	//  }
}

func GetFromVault()  {
	// should recive the passcode as bearer
	// the server should return in the form:
	// {[
	// 		"service1": { ... },
	//		"service2": { ... },
	// ]}
	// this should return just the fields that are present in the valut
	// we'll see about the interface later on...
}

func UpdateValut() {
	// should recive the passcode as Bearer
	// the body should be of the following structure
	// {
	//		"service": { "field", ["string"] }
	// }
}
