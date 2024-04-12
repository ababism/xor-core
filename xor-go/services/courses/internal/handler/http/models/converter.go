package models

import (
	"time"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/handler/generated"
)

const (
	UnknownVisibility = "unknown"
	UnknownRequestStatus
	UnknownAccessStatus
)

func ToCourseListResponse(courses []*domain.Course) []generated.Course {
	courseResponses := make([]generated.Course, len(courses))
	for i, course := range courses {
		if course != nil {
			courseResponses[i] = ToCourseResponse(*course)
		}
	}
	return courseResponses
}

func ToCourseResponse(course domain.Course) generated.Course {
	return generated.Course{
		Discipline: course.Discipline,
		FeedbackID: &course.FeedbackID,
		ID:         &course.ID,
		Landing:    course.Landing,
		Name:       course.Name,
		Sections:   ToSectionSliceResponse(course.Sections),
		TeacherID:  &course.TeacherID,
		Visibility: ToVisibilityCourse(course.Visibility),
	}
}

func ToSectionSliceResponse(sections []domain.Section) *[]generated.Section {
	if sections == nil {
		return nil
	}

	sectionResponses := make([]generated.Section, len(sections))
	for i, section := range sections {
		sectionResponses[i] = ToSectionResponse(section)
	}
	return &sectionResponses
}

func ToSectionResponse(section domain.Section) generated.Section {
	return generated.Section{
		Description: section.Description,
		//FeedbackID:  section.FeedbackID,
		Heading:    section.Heading,
		ID:         &section.ID,
		Themes:     ToThemeResponseSlice(section.Themes),
		Visibility: ToVisibilitySection(section.Visibility),
	}
}

func ToThemeResponseSlice(themes []domain.Theme) *[]generated.Theme {
	if themes == nil {
		return nil
	}

	themeResponses := make([]generated.Theme, len(themes))
	for i, theme := range themes {
		themeResponses[i] = ToThemeResponse(theme)
	}
	return &themeResponses
}

func ToThemeResponse(theme domain.Theme) generated.Theme {
	return generated.Theme{
		//FeedbackID: theme.FeedbackID,
		Heading:    theme.Heading,
		ID:         &theme.ID,
		LessonIDs:  &theme.LessonIDs,
		Visibility: ToVisibilityTheme(theme.Visibility),
	}
}

func ToLessonResponse(lesson domain.Lesson) generated.Lesson {
	pr := ToProductResponse(lesson.Product)
	return generated.Lesson{
		CourseID:   lesson.CourseID,
		ID:         &lesson.ID,
		Product:    &pr,
		TeacherID:  &lesson.TeacherID,
		Transcript: lesson.Transcript,
		VideoURI:   &lesson.VideoURI,
		Visibility: ToVisibilityLesson(lesson.Visibility),
	}
}

func ToProductResponse(product domain.Product) generated.Product {
	return generated.Product{
		ID:    &product.ID,
		Item:  &product.Item,
		Owner: &product.Owner,
		Price: product.Price,
	}
}

func ToPaymentRedirectResponse(paymentRedirect domain.PaymentRedirect) generated.PaymentRedirect {
	return generated.PaymentRedirect{
		Response: &paymentRedirect.Response,
	}
}

func ToPublicationRequestResponse(publicationRequest domain.PublicationRequest) generated.PublicationRequest {
	rs := ToRequestStatus(publicationRequest.RequestStatus)
	return generated.PublicationRequest{
		AssigneeID:    publicationRequest.AssigneeID,
		Comment:       publicationRequest.Comment,
		CourseID:      publicationRequest.CourseID,
		ID:            &publicationRequest.ID,
		RequestStatus: &rs,
		UpdatedAt:     ToRequestTime(publicationRequest.UpdatedAt),
	}
}

func ToRequestTime(t time.Time) *time.Time {
	var resT *time.Time
	var defaultTime time.Time
	if t == defaultTime {
		resT = nil
	} else {
		resT = &t
	}
	return resT
}
func ToAccessStatus(accessStatus domain.AccessStatus) generated.LessonAccessAccessStatus {

	switch accessStatus {
	case domain.Accessible:
		return generated.Accessible
	case domain.Inaccessible:
		return generated.Inaccessible
	case domain.Pending:
		return generated.Pending
	default:
		return UnknownAccessStatus
	}
}

func ToLessonAccessResponse(lessonAccess domain.LessonAccess) generated.LessonAccess {
	status := ToAccessStatus(lessonAccess.AccessStatus)
	return generated.LessonAccess{
		ID:           &lessonAccess.ID,
		LessonID:     lessonAccess.LessonID,
		StudentID:    lessonAccess.StudentID,
		AccessStatus: status,
		UpdatedAt:    ToRequestTime(lessonAccess.UpdatedAt),
	}
}

func ToRequestStatus(status domain.RequestsStatus) generated.PublicationRequestRequestStatus {
	switch status {
	case domain.Approved:
		return generated.Approved
	case domain.Rejected:
		return generated.Rejected
	case domain.Unwatched:
		return generated.Unwatched
	default:
		return UnknownRequestStatus
	}
}
func ToVisibilityCourse(v domain.Visibility) generated.CourseVisibility {
	switch v {
	case domain.Visible:
		return generated.CourseVisibilityVisible
	case domain.Hidden:
		return generated.CourseVisibilityVisible
	default:
		return UnknownVisibility
	}
}

func ToVisibilitySection(v domain.Visibility) generated.SectionVisibility {
	switch v {
	case domain.Visible:
		return generated.SectionVisibilityVisible
	case domain.Hidden:
		return generated.SectionVisibilityHidden
	default:
		return UnknownVisibility
	}
}

func ToVisibilityTheme(v domain.Visibility) generated.ThemeVisibility {
	switch v {
	case domain.Visible:
		return generated.ThemeVisibilityVisible
	case domain.Hidden:
		return generated.ThemeVisibilityHidden
	default:
		return UnknownVisibility
	}
}

func ToVisibilityLesson(v domain.Visibility) generated.LessonVisibility {
	switch v {
	case domain.Visible:
		return generated.LessonVisibilityVisible
	case domain.Hidden:
		return generated.LessonVisibilityHidden
	default:
		return UnknownVisibility
	}
}
