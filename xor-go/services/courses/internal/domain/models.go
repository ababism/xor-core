package domain

import "github.com/google/uuid"

const (
	ServiceName = "xor-go/courses"
)

// Visibility represents the visibility status of an entity
type Visibility int

const (
	_ Visibility = iota
	Hidden
	Visible
	Unlocked
)

type RequestsStatus int

const (
	_ RequestsStatus = iota
	Unwatched
	Approved
	Rejected
)

type AccessStatus int

const (
	_ AccessStatus = iota
	Inaccessible
	Pending
	Accessible
)

// Actor represents a user of system with data taken from request
type Actor struct {
	ID   uuid.UUID
	role string
}

// Teacher represents a teacher entity
type Teacher struct {
	AccountID uuid.UUID
	//Courses   []uuid.UUID
}

// Student represents a student entity — registered user in this system
type Student struct {
	AccountID uuid.UUID
}

// Course represents a course entity
type Course struct {
	ID         uuid.UUID
	ProductID  uuid.UUID
	FeedbackID uuid.UUID
	Name       string
	Discipline string
	Landing    byte
	Visibility Visibility
	Sections   []Section
}

type Section struct {
	ID          uuid.UUID
	ProductID   uuid.UUID
	FeedbackID  uuid.UUID
	Heading     string
	Description string
	Visibility  Visibility
	Themes      []Theme
}

type Theme struct {
	ID         uuid.UUID
	ProductID  uuid.UUID
	FeedbackID uuid.UUID
	Heading    string
	Visibility Visibility
	Lessons    []Lesson
}

// Lesson represents a lesson entity
type Lesson struct {
	Product    Product
	ID         uuid.UUID
	Transcript string
	VideoURI   string
	Visibility Visibility
}

type Product struct {
	ID    uuid.UUID
	Owner uuid.UUID
	Price float64
	Item  uuid.UUID
}

type PublicationRequest struct {
	ID         uuid.UUID
	CourseID   uuid.UUID
	AssigneeID uuid.UUID
}
