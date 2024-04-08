package collections

type CollectionName string

func (cn CollectionName) String() string {
	return string(cn)
}

const (
	CourseCollectionName              CollectionName = "courses"
	CourseEditorCollectionName        CollectionName = "courses_edit"
	LessonCollectionName              CollectionName = "lessons"
	LessonEditorCollectionName        CollectionName = "lessons_edit"
	StudentCollectionName             CollectionName = "students"
	TeacherCollectionName             CollectionName = "teachers"
	PublicationRequestsCollectionName CollectionName = "publication_requests"
	LessonAccessCollectionName        CollectionName = "lesson_accesses"
)
