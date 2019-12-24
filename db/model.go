package db

// StudentUpdateRequest struct
type StudentUpdateRequest struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	ClassName string `json:"class_name" bson:"class_name"`
	Email     string `json:"email"`
	Age       int    `json:"age" bson:"age"`
}

type StudentSearchRequest struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty" bson:"first_name"`
	LastName  string `json:"last_name,omitempty" bson:"last_name"`
	ClassName string `json:"class_name,omitempty" bson:"class_name"`
	Email     string `json:"email,omitempty"`
	Age       int    `json:"age,omitempty" bson:"age"`
}

// Error struct
type Error struct {
	Code int
	Msg  string
}

// Student struct
type Student struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	ClassName string `json:"class_name" bson:"class_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}
