package domain

import (
	"github.com/google/uuid"
	"time"
	"xor-go/pkg/xstringset"
)

const (
	ServiceName = "xor-go/courses"
)

// Visibility represents the visibility status of an entity: Hidden, Visible
type Visibility int

const (
	_ Visibility = iota
	Hidden
	Visible
)

// RequestsStatus Unwatched, Approved or Rejected
type RequestsStatus int

const (
	_ RequestsStatus = iota
	Unwatched
	Approved
	Rejected
)

// AccessStatus for user to course Inaccessible, Pending or Accessible
type AccessStatus int

const (
	_ AccessStatus = iota
	Inaccessible
	Pending
	Accessible
)

// Actor represents a user of system with data taken from request
type Actor struct {
	ID    uuid.UUID
	roles xstringset.Set
}

func NewActor(ID uuid.UUID, roles []string) *Actor {
	a := &Actor{ID: ID,
		roles: xstringset.New()}
	a.initRoles(roles)
	return a
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
	ID uuid.UUID
	//ProductID  uuid.UUID
	FeedbackID uuid.UUID
	TeacherID  uuid.UUID
	Name       string
	Discipline string
	Landing    []byte
	Visibility Visibility
	Sections   []Section
}

type Section struct {
	ID uuid.UUID
	//ProductID   uuid.UUID
	FeedbackID  uuid.UUID
	Heading     string
	Description string
	Visibility  Visibility
	Themes      []Theme
}

type Theme struct {
	ID uuid.UUID
	//ProductID  uuid.UUID
	FeedbackID uuid.UUID
	Heading    string
	Visibility Visibility
	// TODO MB Make uuid.UUID?
	// Lessons in most cases nil
	Lessons []Lesson
}

// Lesson represents a lesson entity
type Lesson struct {
	ID         uuid.UUID
	SectionID  uuid.UUID
	TeacherID  uuid.UUID
	Product    Product
	Visibility Visibility
	Transcript string
	VideoURI   string
}

type Product struct {
	ID    uuid.UUID
	Owner uuid.UUID
	Price float64
	Item  uuid.UUID
}

type PublicationRequest struct {
	ID            uuid.UUID
	CourseID      uuid.UUID
	AssigneeID    uuid.UUID
	RequestStatus RequestsStatus
	Comment       *string
	UpdatedAt     time.Time
}

type LessonAccess struct {
	ID           uuid.UUID
	LessonID     uuid.UUID
	StudentID    uuid.UUID
	AccessStatus AccessStatus
	UpdatedAt    time.Time
}
type PaymentRedirect struct {
	response []byte
}
