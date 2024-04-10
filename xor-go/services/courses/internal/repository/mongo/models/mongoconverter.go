package models

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"xor-go/pkg/xapperror"
	"xor-go/services/courses/internal/domain"
)

func (cm *Course) ToDomain() domain.Course {
	return domain.Course{}
}
func (l *Lesson) ToDomain() (*domain.Lesson, error) {
	ID, err := ParseUUID(l.ID)
	if err != nil {
		return nil, err
	}
	tID, err := ParseUUID(l.TeacherID)
	if err != nil {
		return nil, err
	}
	cID, err := ParseUUID(l.CourseID)
	if err != nil {
		return nil, err
	}

	return &domain.Lesson{
		ID:         ID,
		CourseID:   cID,
		TeacherID:  tID,
		Product:    domain.Product{},
		Visibility: domain.Visibility(l.Visibility),
		Transcript: l.Transcript,
		VideoURI:   l.VideoURI,
	}, nil
}
func ParseUUID(uuidStr string) (uuid.UUID, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return uuid.Nil, xapperror.New(http.StatusInternalServerError,
			"can't parses uuid from db",
			"can't parses uuid from db", err)
	}

	return id, nil
}
func ParseUUIDWithInfo(uuidStr, fieldName, structName string) (uuid.UUID, error) {

	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return uuid.Nil, xapperror.New(http.StatusInternalServerError,
			fmt.Sprintf("uuid string %s field of %s", fieldName, structName),
			fmt.Sprintf("uuid string %s field of %s", fieldName, structName), err)
	}

	return id, nil
}

//func ParseUUIDs(uuidStrings []string) ([]uuid.UUID, error) {
//	var uuids []uuid.UUID
//
//	for _, uuidStr := range uuidStrings {
//		id, err := uuid.Parse(uuidStr)
//		if err != nil {
//			return nil, err // Return early if any UUID string is invalid
//		}
//		uuids = append(uuids, id)
//	}
//
//	return uuids, nil
//}

//func (cm *Course) ToDomain() domain.Course {
//	var sections []domain.Section
//	for _, sec := range cm.Sections {
//		sections = append(sections, sec.ToDomain())
//	}
//	return domain.Course{
//		ID:         cm.ID,
//		FeedbackID: cm.FeedbackID,
//		Name:       cm.Name,
//		Discipline: cm.Discipline,
//		Landing:    cm.Landing,
//		Visibility: cm.Visibility,
//		Sections:   sections,
//	}
//}
//
//func (sm *Section) ToDomain() domain.Section {
//	var themes []domain.Theme
//	for _, theme := range sm.Themes {
//		themes = append(themes, theme.ToDomain())
//	}
//	return domain.Section{
//		ID:          sm.ID,
//		FeedbackID:  sm.FeedbackID,
//		Heading:     sm.Heading,
//		Description: sm.Description,
//		Visibility:  sm.Visibility,
//		Themes:      themes,
//	}
//}
//
//func (tm *Theme) ToDomain() domain.Theme {
//
//	return domain.Theme{
//		ID:         tm.ID,
//		FeedbackID: tm.FeedbackID,
//		Heading:    tm.Heading,
//		Visibility: tm.Visibility,
//		LessonIDs:  nil,
//	}
//}
//
//func (tm *Theme) ToLessonDomain() domain.Theme {
//	return domain.Theme{
//		ID:         tm.ID,
//		FeedbackID: tm.FeedbackID,
//		Heading:    tm.Heading,
//		Visibility: tm.Visibility,
//		LessonIDs:  tm.LessonIDs,
//	}
//}
//
//func (lm *Lesson) ToDomain() domain.Lesson {
//	return domain.Lesson{
//		Product:    lm.Product.ToDomain(),
//		ID:         lm.ID,
//		Transcript: lm.Transcript,
//		VideoURI:   lm.VideoURI,
//		Visibility: lm.Visibility,
//	}
//}
//func (pm *Product) ToDomain() domain.Product {
//	return domain.Product{
//		ID:    pm.ID,
//		Owner: pm.Owner,
//		Price: pm.Price,
//		Item:  pm.Item,
//	}
//}
//func (prm *PublicationRequest) ToDomain() domain.PublicationRequest {
//	return domain.PublicationRequest{
//		ID:         prm.ID,
//		CourseID:   prm.CourseID,
//		AssigneeID: prm.AssigneeID,
//	}
//}
//func (tm *Teacher) ToDomain() domain.Teacher {
//	return domain.Teacher{
//		AccountID: tm.AccountID,
//	}
//}
//
//func (sm *Student) ToDomain() domain.Student {
//	return domain.Student{
//		AccountID: sm.AccountID,
//	}
//}
