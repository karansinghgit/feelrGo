package db

//Mapping represents the structure of an Index in ES
var Mapping = `{
	"mappings": {
	  "properties": {
		"feelr": {
		  "properties": {
			"feelrId": {
			  "type": "text"
			},
			"question": {
			  "type": "text"
			},
			"timestamp": {
			  "type": "text"
			},
			"topic": {
			  "type": "text"
			}
		  }
		},
		"message": {
		  "properties": {
			"chatID": {
			  "type": "text"
			},
			"senderID": {
			  "type": "text"
			},
			"text": {
			  "type": "text"
			},
			"feelrID": {
			  "type": "text"
			},
			"senderAnswer": {
			  "type": "text"
			},
			"receiverAnswer": {
			  "type": "text"
			},
			"timestamp": {
			  "type": "text"
			}
		  }
		},
		"user": {
		  "properties": {
			"userID": {
			  "type": "text"
			},
			"name": {
		      "type": "text"
			}
		  }
		},
		"chat": {
		  "properties": {
			"chatID": {
			  "type": "text"
			},
			"senderID": {
			  "type": "text"
			},
			"receiverID": {
			  "type": "text"
			}
		  }
		}
	  }
	}
  }`
