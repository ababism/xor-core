package models

import (
	"github.com/google/uuid"
	"xor-go/services/courses/internal/domain"
)

// TODO Review

// Course represents a course entity
type Course struct {
	ID         uuid.UUID         `bson:"_id"`
	ProductID  uuid.UUID         `bson:"product_id"`
	FeedbackID uuid.UUID         `bson:"feedback_id"`
	Name       string            `bson:"name"`
	Discipline string            `bson:"discipline"`
	Landing    []byte            `bson:"landing"`
	Visibility domain.Visibility `bson:"visibility"`
	Sections   []Section         `bson:"sections"`
}

// Section represents a section entity
type Section struct {
	ID          uuid.UUID         `bson:"_id"`
	ProductID   uuid.UUID         `bson:"product_id"`
	FeedbackID  uuid.UUID         `bson:"feedback_id"`
	Heading     string            `bson:"heading"`
	Description string            `bson:"description"`
	Visibility  domain.Visibility `bson:"visibility"`
	Themes      []Theme           `bson:"themes"`
}

// Theme represents a theme entity
type Theme struct {
	ID         uuid.UUID         `bson:"_id"`
	ProductID  uuid.UUID         `bson:"product_id"`
	FeedbackID uuid.UUID         `bson:"feedback_id"`
	Heading    string            `bson:"heading"`
	Visibility domain.Visibility `bson:"visibility"`
	LessonIDs  []uuid.UUID
}

// Lesson represents a lesson entity
type Lesson struct {
	ID         uuid.UUID         `bson:"_id"`
	Product    Product           `bson:"product"`
	Visibility domain.Visibility `bson:"visibility"`
	Transcript string            `bson:"transcript"`
	VideoURI   string            `bson:"video_uri"`
}

// Product represents a product entity
type Product struct {
	ID    uuid.UUID `bson:"_id"`
	Owner uuid.UUID `bson:"owner"`
	Price float64   `bson:"price"`
	Item  uuid.UUID `bson:"item"`
}

// PublicationRequest represents a publication request entity
type PublicationRequest struct {
	ID         uuid.UUID `bson:"_id"`
	CourseID   uuid.UUID `bson:"course_id"`
	AssigneeID uuid.UUID `bson:"assignee_id"`
}

// Teacher represents a teacher entity
type Teacher struct {
	AccountID uuid.UUID `bson:"account_id"`
	//Courses   []uuid.UUID `bson:"courses"`
}

// Student represents a student entity
type Student struct {
	AccountID uuid.UUID `bson:"account_id"`
}
