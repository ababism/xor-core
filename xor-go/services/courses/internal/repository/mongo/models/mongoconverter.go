package models

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"xor-go/pkg/xapperror"
	"xor-go/services/courses/internal/domain"
)

func ParseUUID(uuidStr string) (uuid.UUID, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return uuid.Nil, xapperror.New(http.StatusInternalServerError,
			"can't parses uuid from db",
			"can't parses uuid from db", err)
	}

	return id, nil
}

func (c *Course) ToDomain() (*domain.Course, error) {
	ID, err := ParseUUID(c.ID)
	if err != nil {
		return nil, err
	}
	tID, err := ParseUUID(c.TeacherID)
	if err != nil {
		return nil, err
	}
	fID, err := ParseUUID(c.FeedbackID)
	if err != nil {
		return nil, err
	}
	dSections, err := ToSectionsDomain(c.Sections)
	if err != nil {
		return nil, err
	}
	return &domain.Course{
		ID:         ID,
		FeedbackID: fID,
		TeacherID:  tID,
		Name:       c.Name,
		Discipline: c.Discipline,
		Landing:    c.Landing,
		Visibility: domain.Visibility(c.Visibility),
		Sections:   dSections,
	}, nil
}

func ToSectionsDomain(sections []Section) ([]domain.Section, error) {
	var result []domain.Section
	for _, section := range sections {
		s, err := section.ToDomain()
		if err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	return result, nil
}

func (s Section) ToDomain() (domain.Section, error) {
	ID, err := ParseUUID(s.ID)
	if err != nil {
		return domain.Section{}, err
	}
	themes := make([]domain.Theme, 0)
	for _, theme := range s.Themes {
		t, err := theme.ToDomain()
		if err != nil {
			return domain.Section{}, err
		}
		themes = append(themes, t)
	}
	return domain.Section{
		ID:          ID,
		Heading:     s.Heading,
		Description: s.Description,
		Visibility:  domain.Visibility(s.Visibility),
		Themes:      themes,
	}, nil

}

func (s Theme) ToDomain() (domain.Theme, error) {
	ID, err := ParseUUID(s.ID)
	if err != nil {
		return domain.Theme{}, err
	}
	lessonIDs := make([]uuid.UUID, 0)
	for _, lessonID := range s.LessonIDs {
		id, err := ParseUUID(lessonID)
		if err != nil {
			return domain.Theme{}, err
		}
		lessonIDs = append(lessonIDs, id)
	}
	return domain.Theme{
		ID:         ID,
		Heading:    s.Heading,
		Visibility: domain.Visibility(s.Visibility),
		LessonIDs:  lessonIDs,
	}, nil
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

	lProduct, err := l.Product.ToDomain()
	if err != nil {
		return nil, err
	}

	return &domain.Lesson{
		ID:         ID,
		CourseID:   cID,
		TeacherID:  tID,
		Product:    lProduct,
		Visibility: domain.Visibility(l.Visibility),
		Transcript: l.Transcript,
		VideoURI:   l.VideoURI,
	}, nil
}
func (l *Product) ToDomain() (domain.Product, error) {
	ID, err := ParseUUID(l.ID)
	if err != nil {
		return domain.Product{}, err
	}
	oID, err := ParseUUID(l.Owner)
	if err != nil {
		return domain.Product{}, err

	}
	iID, err := ParseUUID(l.ItemID)
	if err != nil {
		return domain.Product{}, err

	}
	return domain.Product{
		ID:    ID,
		Owner: oID,
		Price: l.Price,
		Item:  iID,
	}, nil
}

func (p *PublicationRequest) ToDomain() (*domain.PublicationRequest, error) {
	ID, err := ParseUUID(p.ID)
	if err != nil {
		return nil, err
	}
	cID, err := ParseUUID(p.CourseID)
	if err != nil {
		return nil, err
	}
	aID, err := ParseUUID(p.AssigneeID)
	if err != nil {
		return nil, err
	}
	return &domain.PublicationRequest{
		ID:         ID,
		CourseID:   cID,
		AssigneeID: aID,
		UpdatedAt:  p.UpdatedAt,
	}, nil
}

func (t *Teacher) ToDomain() (*domain.Teacher, error) {
	aID, err := ParseUUID(t.AccountID)
	if err != nil {
		return nil, err
	}
	return &domain.Teacher{
		AccountID: aID,
	}, nil
}

func (s *Student) ToDomain() (*domain.Student, error) {
	aID, err := ParseUUID(s.AccountID)
	if err != nil {
		return nil, err
	}
	return &domain.Student{
		AccountID: aID,
	}, nil
}

func (la *LessonAccess) ToDomain() (domain.LessonAccess, error) {
	ID, err := ParseUUID(la.ID)
	if err != nil {
		return domain.LessonAccess{}, err
	}
	lID, err := ParseUUID(la.LessonID)
	if err != nil {
		return domain.LessonAccess{}, err
	}
	return domain.LessonAccess{
		ID:           ID,
		LessonID:     lID,
		AccessStatus: domain.AccessStatus(la.AccessStatus),
		UpdatedAt:    la.UpdatedAt,
	}, nil
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
