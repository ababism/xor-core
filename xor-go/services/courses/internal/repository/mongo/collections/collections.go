package collections

type CollectionName string

func (cn CollectionName) String() string {
	return string(cn)
}

const (
	CourseCollectionName              CollectionName = "course"
	CourseEditorCollectionName        CollectionName = "course_edit"
	LessonCollectionName              CollectionName = "lesson"
	LessonEditorCollectionName        CollectionName = "lesson_edit"
	StudentCollectionName             CollectionName = "student"
	TeacherCollectionName             CollectionName = "teacher"
	PublicationRequestsCollectionName CollectionName = "publication_request"
)
