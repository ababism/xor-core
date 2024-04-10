package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO Review
//
//// Course represents a course entity
//type Course struct {
//	ID         uuid.UUID         `bson:"_id"`
//	ProductID  uuid.UUID         `bson:"product_id"`
//	FeedbackID uuid.UUID         `bson:"feedback_id"`
//	Name       string            `bson:"name"`
//	Discipline string            `bson:"discipline"`
//	Landing    string            `bson:"landing"`
//	Visibility domain.Visibility `bson:"visibility"`
//	Sections   []Section         `bson:"sections"`
//}
//
//// Section represents a section entity
//type Section struct {
//	ID          uuid.UUID         `bson:"_id"`
//	ProductID   uuid.UUID         `bson:"product_id"`
//	FeedbackID  uuid.UUID         `bson:"feedback_id"`
//	Heading     string            `bson:"heading"`
//	Description string            `bson:"description"`
//	Visibility  domain.Visibility `bson:"visibility"`
//	Themes      []Theme           `bson:"themes"`
//}
//
//// Theme represents a theme entity
//type Theme struct {
//	ID         uuid.UUID         `bson:"_id"`
//	ProductID  uuid.UUID         `bson:"product_id"`
//	FeedbackID uuid.UUID         `bson:"feedback_id"`
//	Heading    string            `bson:"heading"`
//	Visibility domain.Visibility `bson:"visibility"`
//	LessonIDs  []uuid.UUID
//}
//
//// Lesson represents a lesson entity
//type Lesson struct {
//	ID         uuid.UUID         `bson:"_id"`
//	Product    Product           `bson:"product"`
//	Visibility domain.Visibility `bson:"visibility"`
//	Transcript string            `bson:"transcript"`
//	VideoURI   string            `bson:"video_uri"`
//}
//
//// Product represents a product entity
//type Product struct {
//	ID    uuid.UUID `bson:"_id"`
//	Owner uuid.UUID `bson:"owner"`
//	Price float32   `bson:"price"`
//	Item  uuid.UUID `bson:"item"`
//}
//
//// PublicationRequest represents a publication request entity
//type PublicationRequest struct {
//	ID         uuid.UUID `bson:"_id"`
//	CourseID   uuid.UUID `bson:"course_id"`
//	AssigneeID uuid.UUID `bson:"assignee_id"`
//}
//
//// Teacher represents a teacher entity
//type Teacher struct {
//	AccountID uuid.UUID `bson:"account_id"`
//	//Courses   []uuid.UUID `bson:"courses"`
//}
//
//// Student represents a student entity
//type Student struct {
//	AccountID uuid.UUID `bson:"account_id"`
//}

type Teacher struct {
	mID       primitive.ObjectID `bson:"_id,omitempty"`
	AccountID string             `bson:"account_id"`
}

type Student struct {
	mID       primitive.ObjectID `bson:"_id,omitempty"`
	AccountID string             `bson:"account_id"`
}

type Course struct {
	mID        primitive.ObjectID `bson:"_id,omitempty"`
	ID         string             `bson:"course_id,omitempty"`
	FeedbackID string             `bson:"feedback_id"`
	TeacherID  string             `bson:"teacher_id"`
	Name       string             `bson:"name"`
	Discipline string             `bson:"discipline"`
	Landing    string             `bson:"landing"`
	Visibility int                `bson:"visibility"`
	Sections   []Section          `bson:"sections"`
}

type Section struct {
	mID         primitive.ObjectID `bson:"_id,omitempty"`
	ID          string             `bson:"section_id,omitempty"`
	FeedbackID  string             `bson:"feedback_id"`
	Heading     string             `bson:"heading"`
	Description string             `bson:"description"`
	Visibility  int                `bson:"visibility"`
	Themes      []Theme            `bson:"themes"`
}

type Theme struct {
	mID        primitive.ObjectID `bson:"_id,omitempty"`
	ID         string             `bson:"theme_id,omitempty"`
	FeedbackID string             `bson:"feedback_id"`
	Heading    string             `bson:"heading"`
	Visibility int                `bson:"visibility"`
	LessonIDs  []string           `bson:"lesson_ids"`
}

type Lesson struct {
	mID        primitive.ObjectID `bson:"_id,omitempty"`
	ID         string             `bson:"lesson_id,omitempty"`
	CourseID   string             `bson:"course_id"`
	TeacherID  string             `bson:"teacher_id"`
	Product    Product            `bson:"product"`
	Visibility int                `bson:"visibility"`
	Transcript string             `bson:"transcript"`
	VideoURI   string             `bson:"video_uri"`
}

type Product struct {
	mID    primitive.ObjectID `bson:"_id,omitempty"`
	ID     string             `bson:"product_id,omitempty"`
	Owner  string             `bson:"owner"`
	Price  float32            `bson:"price"`
	ItemID string             `bson:"item"`
}

type PublicationRequest struct {
	mID           primitive.ObjectID `bson:"_id,omitempty"`
	ID            string             `bson:"publication_request_id,omitempty"`
	CourseID      string             `bson:"course_id"`
	AssigneeID    string             `bson:"assignee_id"`
	RequestStatus int                `bson:"request_status"`
	Comment       *string            `bson:"comment,omitempty"`
	UpdatedAt     primitive.DateTime `bson:"updated_at"`
}

type LessonAccess struct {
	mID          primitive.ObjectID `bson:"_id,omitempty"`
	ID           string             `bson:"lesson_access_id,omitempty"`
	LessonID     string             `bson:"lesson_id"`
	StudentID    string             `bson:"student_id"`
	AccessStatus int                `bson:"access_status"`
	UpdatedAt    primitive.DateTime `bson:"updated_at"`
}
