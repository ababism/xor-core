package models

import (
	"xor-go/services/courses/internal/domain"
)

func (cm *Course) ToDomain() domain.Course {
	var sections []domain.Section
	for _, sec := range cm.Sections {
		sections = append(sections, sec.ToDomain())
	}
	return domain.Course{
		ID:         cm.ID,
		FeedbackID: cm.FeedbackID,
		Name:       cm.Name,
		Discipline: cm.Discipline,
		Landing:    cm.Landing,
		Visibility: cm.Visibility,
		Sections:   sections,
	}
}

func (sm *Section) ToDomain() domain.Section {
	var themes []domain.Theme
	for _, theme := range sm.Themes {
		themes = append(themes, theme.ToDomain())
	}
	return domain.Section{
		ID:          sm.ID,
		FeedbackID:  sm.FeedbackID,
		Heading:     sm.Heading,
		Description: sm.Description,
		Visibility:  sm.Visibility,
		Themes:      themes,
	}
}

func (tm *Theme) ToDomain() domain.Theme {

	return domain.Theme{
		ID:         tm.ID,
		FeedbackID: tm.FeedbackID,
		Heading:    tm.Heading,
		Visibility: tm.Visibility,
		LessonIDs:  nil,
	}
}

func (tm *Theme) ToLessonDomain() domain.Theme {
	return domain.Theme{
		ID:         tm.ID,
		FeedbackID: tm.FeedbackID,
		Heading:    tm.Heading,
		Visibility: tm.Visibility,
		LessonIDs:  tm.LessonIDs,
	}
}

func (lm *Lesson) ToDomain() domain.Lesson {
	return domain.Lesson{
		Product:    lm.Product.ToDomain(),
		ID:         lm.ID,
		Transcript: lm.Transcript,
		VideoURI:   lm.VideoURI,
		Visibility: lm.Visibility,
	}
}
func (pm *Product) ToDomain() domain.Product {
	return domain.Product{
		ID:    pm.ID,
		Owner: pm.Owner,
		Price: pm.Price,
		Item:  pm.Item,
	}
}
func (prm *PublicationRequest) ToDomain() domain.PublicationRequest {
	return domain.PublicationRequest{
		ID:         prm.ID,
		CourseID:   prm.CourseID,
		AssigneeID: prm.AssigneeID,
	}
}
func (tm *Teacher) ToDomain() domain.Teacher {
	return domain.Teacher{
		AccountID: tm.AccountID,
	}
}

func (sm *Student) ToDomain() domain.Student {
	return domain.Student{
		AccountID: sm.AccountID,
	}
}
