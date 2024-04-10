package models

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"xor-go/services/courses/internal/domain"
)

func ToMongoModelLesson(l domain.Lesson) Lesson {
	return Lesson{
		mID:        primitive.NilObjectID,
		ID:         l.ID.String(),
		CourseID:   l.CourseID.String(),
		TeacherID:  l.TeacherID.String(),
		Product:    ToMongoModelProduct(l.Product),
		Visibility: int(l.Visibility),
		Transcript: l.Transcript,
		VideoURI:   l.VideoURI,
	}
}
func ToMongoModelProduct(p domain.Product) Product {
	return Product{
		mID:    primitive.NilObjectID,
		ID:     p.ID.String(),
		Owner:  p.Owner.String(),
		Price:  p.Price,
		ItemID: p.Item.String(),
	}
}

func ToMongoModelPublicationRequest(pr domain.PublicationRequest) PublicationRequest {
	return PublicationRequest{
		mID:        primitive.NilObjectID,
		ID:         pr.ID.String(),
		CourseID:   pr.CourseID.String(),
		AssigneeID: pr.AssigneeID.String(),
		UpdatedAt:  pr.UpdatedAt,
	}
}

func ToMongoModelTeacher(t domain.Teacher) Teacher {
	return Teacher{
		mID:       primitive.NilObjectID,
		AccountID: t.AccountID.String(),
	}
}

func ToMongoModelStudent(s domain.Student) Student {
	return Student{
		mID:       primitive.NilObjectID,
		AccountID: s.AccountID.String(),
	}
}

// TODO Rewrite
func ToMongoModelCourse(course *domain.Course) Course {
	return Course{
		mID:        primitive.NilObjectID,
		ID:         course.ID.String(),
		FeedbackID: course.FeedbackID.String(),
		TeacherID:  course.TeacherID.String(),
		Name:       course.Name,
		Discipline: course.Discipline,
		Landing:    course.Landing,
		Visibility: int(course.Visibility),
		Sections:   ToMongoModelSections(course.Sections),
	}
}

func ToMongoModelSections(sections []domain.Section) []Section {
	var result []Section
	for _, sec := range sections {
		result = append(result, ToMongoModelSection(sec))
	}
	return result
}

func ToMongoModelSection(section domain.Section) Section {
	return Section{
		mID:         primitive.NilObjectID,
		ID:          section.ID.String(),
		Heading:     section.Heading,
		Description: section.Description,
		Visibility:  int(section.Visibility),
		Themes:      ToMongoModelThemes(section.Themes),
	}
}

func ToMongoModelThemes(themes []domain.Theme) []Theme {
	var result []Theme
	for _, theme := range themes {
		result = append(result, ToMongoModelTheme(theme))
	}
	return result
}

func ToMongoModelTheme(theme domain.Theme) Theme {
	return Theme{
		mID:        primitive.NilObjectID,
		ID:         theme.ID.String(),
		Heading:    theme.Heading,
		Visibility: int(theme.Visibility),
		LessonIDs:  ToStringUUIDs(theme.LessonIDs),
	}
}

func ToStringUUIDs(uuids []uuid.UUID) []string {
	var result []string
	for _, id := range uuids {
		result = append(result, id.String())
	}
	return result
}

func ToMongoModelLessonAccess(course domain.LessonAccess) LessonAccess {
	return LessonAccess{
		mID:       primitive.NilObjectID,
		ID:        course.ID.String(),
		StudentID: course.StudentID.String(),
		LessonID:  course.LessonID.String(),
		UpdatedAt: course.UpdatedAt,
	}
}

//
//func ToMongoModelCourse(course *domain.Course) Course {
//	var sections []Section
//	for _, sec := range course.Sections {
//		sections = append(sections, ToMongoModelSection(&sec))
//	}
//	return Course{
//		ID:         course.ID,
//		FeedbackID: course.FeedbackID,
//		Name:       course.Name,
//		Discipline: course.Discipline,
//		Landing:    course.Landing,
//		Visibility: course.Visibility,
//		Sections:   sections,
//	}
//}
//
//func ToMongoModelSection(section *domain.Section) Section {
//	var themes []Theme
//	for _, theme := range section.Themes {
//		themes = append(themes, ToMongoModelTheme(theme))
//	}
//	return Section{
//		ID:          section.ID,
//		FeedbackID:  section.FeedbackID,
//		Heading:     section.Heading,
//		Description: section.Description,
//		Visibility:  section.Visibility,
//		Themes:      themes,
//	}
//}

//func ToMongoModelTheme(theme domain.Theme) Theme {
//
//	return Theme{
//		ID:         theme.ID,
//		FeedbackID: theme.FeedbackID,
//		Heading:    theme.Heading,
//		Visibility: theme.Visibility,
//		LessonIDs:  theme.LessonIDs,
//	}
//}
//
//func ToMongoModelLesson(lesson domain.Lesson) Lesson {
//	return Lesson{
//		Product:    ToMongoModelProduct(lesson.Product),
//		ID:         lesson.ID,
//		Transcript: lesson.Transcript,
//		VideoURI:   lesson.VideoURI,
//		Visibility: lesson.Visibility,
//	}
//}
//
//func ToMongoModelProduct(product domain.Product) Product {
//	return Product{
//		ID:    product.ID,
//		Owner: product.Owner,
//		Price: product.Price,
//		Item:  product.Item,
//	}
//}
//
//func ToMongoModelPublicationRequest(publicationRequest domain.PublicationRequest) PublicationRequest {
//	return PublicationRequest{
//		ID:         publicationRequest.ID,
//		CourseID:   publicationRequest.CourseID,
//		AssigneeID: publicationRequest.AssigneeID,
//	}
//}
//
//func ToMongoModelTeacher(teacher domain.Teacher) Teacher {
//	return Teacher{
//		AccountID: teacher.AccountID,
//	}
//}
//
//func ToMongoModelStudent(student Student) Student {
//	return Student{
//		AccountID: student.AccountID,
//	}
//}
