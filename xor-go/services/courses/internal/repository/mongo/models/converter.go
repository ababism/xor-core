package models

import (
	"xor-go/services/courses/internal/domain"
)

func ToMongoModelCourse(course *domain.Course) Course {
	var sections []Section
	for _, sec := range course.Sections {
		sections = append(sections, ToMongoModelSection(&sec))
	}
	return Course{
		ID:         course.ID,
		ProductID:  course.ProductID,
		FeedbackID: course.FeedbackID,
		Name:       course.Name,
		Discipline: course.Discipline,
		Landing:    course.Landing,
		Visibility: course.Visibility,
		Sections:   sections,
	}
}

func ToMongoModelSection(section *domain.Section) Section {
	var themes []Theme
	for _, theme := range section.Themes {
		themes = append(themes, ToMongoModelTheme(theme))
	}
	return Section{
		ID:          section.ID,
		ProductID:   section.ProductID,
		FeedbackID:  section.FeedbackID,
		Heading:     section.Heading,
		Description: section.Description,
		Visibility:  section.Visibility,
		Themes:      themes,
	}
}

func ToMongoModelTheme(theme domain.Theme) Theme {
	var lessons []Lesson
	for _, lesson := range theme.Lessons {
		lessons = append(lessons, ToMongoModelLesson(lesson))
	}
	return Theme{
		ID:         theme.ID,
		ProductID:  theme.ProductID,
		FeedbackID: theme.FeedbackID,
		Heading:    theme.Heading,
		Visibility: theme.Visibility,
		Lessons:    lessons,
	}
}

func ToMongoModelLesson(lesson domain.Lesson) Lesson {
	return Lesson{
		Product:    ToMongoModelProduct(lesson.Product),
		ID:         lesson.ID,
		Transcript: lesson.Transcript,
		VideoURI:   lesson.VideoURI,
		Visibility: lesson.Visibility,
	}
}

func ToMongoModelProduct(product domain.Product) Product {
	return Product{
		ID:    product.ID,
		Owner: product.Owner,
		Price: product.Price,
		Item:  product.Item,
	}
}

func ToMongoModelPublicationRequest(publicationRequest domain.PublicationRequest) PublicationRequest {
	return PublicationRequest{
		ID:         publicationRequest.ID,
		CourseID:   publicationRequest.CourseID,
		AssigneeID: publicationRequest.AssigneeID,
	}
}

func ToMongoModelTeacher(teacher domain.Teacher) Teacher {
	return Teacher{
		AccountID: teacher.AccountID,
	}
}

func ToMongoModelStudent(student Student) Student {
	return Student{
		AccountID: student.AccountID,
	}
}
