package request

type WebUser struct {
	ID			interface{}			`bson:"_id,omitempty"`	// mongodb创建时自动生成的ID
	Username 	string 				`bson:"username" binding:"required"`
	Password 	string 				`bson:"password" binding:"required"`
	Email 		string 				`bson:"email" binding:"required"`
}

type WebRegisterData struct {
	Username 	string 				`json:"username"`
	Password 	string 				`json:"password"`
	Email 		string 				`json:"email"`
}

type WebLoginData struct {
	Email 		string 				`json:"email" binding:"required"`
	Password 	string 				`json:"password" binding:"required"`
}

type WebTestData struct {
	Email 		string 				`json:"email" binding:"required"`
	Password 	string 				`json:"password" binding:"required"`
}

type WebLoginResponseData struct {
	ID			interface{}			`bson:"_id"`
	Email 		string				`bson:"email"`
	//Phone 		string 				`bson:"phone"`
}


