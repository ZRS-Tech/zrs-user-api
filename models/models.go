package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Create Struct
type Book struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Isbn   string             `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Title  string             `json:"title" bson:"title,omitempty"`
	Author *Author            `json:"author" bson:"author,omitempty"`
}

type Author struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

type Subject struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Code string             `json:"code,omitempty" bson:"code,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
}

type QuestionItem struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title          string             `json:"title,omitempty" bson:"title,omitempty"`
	Description    string             `json:"description,omitempty" bson:"description,omitempty"`
	Subject        *Subject           `json:"subject" bson:"subject,omitempty"`
	AnswerOptions  []string           `json:"answerOptions" bson:"answerOptions,omitempty"`
	CurrentAnswers []string           `json:"currentAnswers" bson:"currentAnswers,omitempty"`
}

type Customer struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DisplayName string `json:"displayName,omitempty" bson:"displayName,omitempty"`
	EmailAddress string `json:"emailaddress,omitempty" bson:"emailaddress,omitempty"`
	PrimaryMobileNumber string `json:"primaryMobileNumber,omitempty" bson:"primaryMobileNumber,omitempty"`
	AlternateMobileNumber string `json:"alternateMobileNumber,omitempty" bson:"alternateMobileNumber,omitempty"`
	Source string `json:"source,omitempty" bson:"source,omitempty"`
	AuthProvider string `json:"authProvider,omitempty" bson:"uuthProvider,omitempty"`
	UserId  string `json:"userid,omitempty" bson:"userid,omitempty"`
	Grade   string `json:"grade" bson:"grade,omitempty"`
	Session string `json:"session,omitempty" bson:"session,omitempty"`
	Stream  string `json:"stream,omitempty" bson:"stream,omitempty"`
	School  string `json:"school,omitempty" bson:"school,omitempty"`
	City    string `json:"city,omitempty" bson:"city,omitempty"`
	State   string `json:"state,omitempty" bson:"state,omitempty"`
	Pincode string `json:"pincode,omitempty" bson:"pincode,omitempty"`
	LastLoginAt string `json:"lastLoginAt,omitempty" bson:"lastLoginAt,omitempty"`
	CreatedAt string `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	LastUpdatedAt string `json:"lastUpdatedAt,omitempty" bson:"lastUpdatedAt,omitempty"`
	SecurityQuestion string `json:"securityQuestion,omitempty" bson:"securityQuestion,omitempty"`
	SecurityAnswer string `json:"securityAnswer,omitempty" bson:"securityAnswer,omitempty"`
}
