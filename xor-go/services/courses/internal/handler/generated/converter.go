package generated

import (
	"time"
	"xor-go/services/courses/internal/domain"
)

func (a OptionalActor) ToDomainWithRoles(roles []string) domain.Actor {
	if a.Roles == nil {
		s := []string{domain.UnregisteredRole}
		return domain.NewActor(a.ID, s)
	}
	roles = append(roles, *a.Roles)
	return domain.NewActor(a.ID, roles)
}

func (a OptionalActor) ToDomain() domain.Actor {
	if a.Roles == nil {
		s := []string{domain.UnregisteredRole}
		return domain.NewActor(a.ID, s)
	}
	s := []string{*a.Roles}
	return domain.NewActor(a.ID, s)
}

func (a Actor) ToDomainWithRoles(roles []string) domain.Actor {
	roles = append(roles, a.Roles)
	return domain.NewActor(a.ID, roles)
}

func (a Actor) ToDomain() domain.Actor {
	s := []string{a.Roles}
	return domain.NewActor(a.ID, s)
}

func (c Course) ToDomain() *domain.Course {
	return &domain.Course{
		ID:         c.ID,
		FeedbackID: c.FeedbackID,
		TeacherID:  c.TeacherID,
		Name:       c.Name,
		Discipline: c.Discipline,
		Landing:    c.Landing,
		Visibility: c.Visibility.ToDomain(),
		// generic
		Sections: ToSectionSliceToDomain(*c.Sections),
	}
}

func ToSectionSliceToDomain(sections []Section) []domain.Section {
	if sections == nil {
		return nil
	}

	sectionResponses := make([]domain.Section, len(sections))
	for i, section := range sections {
		sectionResponses[i] = section.ToDomain()
	}
	return sectionResponses
}

func (s Section) ToDomain() domain.Section {
	return domain.Section{
		Description: s.Description,
		//FeedbackID:  s.FeedbackID,
		Heading:    s.Heading,
		ID:         s.ID,
		Themes:     ToThemeSliceToDomain(*s.Themes),
		Visibility: s.Visibility.ToDomain(),
	}
}
func ToThemeSliceToDomain(themes []Theme) []domain.Theme {
	if themes == nil {
		return nil
	}

	themeResponses := make([]domain.Theme, len(themes))
	for i, theme := range themes {
		themeResponses[i] = theme.ToDomain()
	}
	return themeResponses
}

func (t Theme) ToDomain() domain.Theme {
	return domain.Theme{
		//FeedbackID: t.FeedbackID,
		Heading:    t.Heading,
		ID:         t.ID,
		LessonIDs:  *t.LessonIDs,
		Visibility: t.Visibility.ToDomain(),
	}
}

func (l Lesson) ToDomain() *domain.Lesson {
	pr := domain.Product{}
	if l.Product != nil {
		pr = l.Product.ToDomain()
	}
	return &domain.Lesson{
		CourseID:   l.CourseID,
		ID:         l.ID,
		Product:    pr,
		TeacherID:  l.TeacherID,
		Transcript: l.Transcript,
		VideoURI:   *l.VideoURI,
		Visibility: l.Visibility.ToDomain(),
	}
}

func (p Product) ToDomain() domain.Product {
	return domain.Product{
		ID:    p.ID,
		Item:  p.Item,
		Owner: p.Owner,
		Price: p.Price,
	}
}
func ToProductSliceDomain(products []Product) []domain.Product {
	if products == nil {
		return nil
	}

	result := make([]domain.Product, len(products))
	for i, p := range products {
		result[i] = p.ToDomain()
	}
	return result
}

func (pr PaymentRedirect) ToDomain() domain.PaymentRedirect {
	if pr.Response == nil {
		return domain.PaymentRedirect{}
	}
	return domain.PaymentRedirect{
		Response: *pr.Response,
	}
}

func (p PublicationRequest) ToDomain() domain.PublicationRequest {
	return domain.PublicationRequest{
		AssigneeID:    p.AssigneeID,
		Comment:       p.Comment,
		CourseID:      p.CourseID,
		ID:            p.ID,
		RequestStatus: p.RequestStatus.ToDomain(),
		UpdatedAt:     ToTimeDomain(p.UpdatedAt),
	}
}
func (pr Teacher) ToDomain() domain.Teacher {
	return domain.Teacher{
		AccountID: pr.AccountID,
	}
}
func (pr Student) ToDomain() domain.Student {
	return domain.Student{
		AccountID: pr.AccountID,
	}
}

func ToTimeDomain(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

func (l LessonAccess) ToDomain() domain.LessonAccess {
	var updatedAt *time.Time
	if l.UpdatedAt != nil {
		updatedAt = l.UpdatedAt
	}

	return domain.LessonAccess{
		ID:           l.ID,
		LessonID:     l.LessonID,
		StudentID:    l.StudentID,
		AccessStatus: ToAccessStatus(l.AccessStatus),
		UpdatedAt:    ToTimeDomain(updatedAt),
	}
}

func ToAccessStatus(accessStatus *LessonAccessAccessStatus) domain.AccessStatus {
	if accessStatus == nil {
		return 0
	}
	switch *accessStatus {
	case Accessible:
		return domain.Accessible
	case Inaccessible:
		return domain.Inaccessible
	case Pending:
		return domain.Pending
	default:
		return 0
	}
}

func (t ThemeVisibility) ToDomain() domain.Visibility {
	switch t {
	case ThemeVisibilityVisible:
		return domain.Visible
	case ThemeVisibilityHidden:
		return domain.Hidden
	default:
		return 0
	}
}

func (s SectionVisibility) ToDomain() domain.Visibility {
	switch s {
	case SectionVisibilityVisible:
		return domain.Visible
	case SectionVisibilityHidden:
		return domain.Hidden
	default:
		return 0
	}
}

func (l LessonVisibility) ToDomain() domain.Visibility {
	switch l {
	case LessonVisibilityVisible:
		return domain.Visible
	case LessonVisibilityHidden:
		return domain.Hidden
	default:
		return 0
	}
}

func (c CourseVisibility) ToDomain() domain.Visibility {
	switch c {
	case CourseVisibilityVisible:
		return domain.Visible
	case CourseVisibilityHidden:
		return domain.Hidden
	default:
		return 0
	}
}

func (r PublicationRequestRequestStatus) ToDomain() domain.RequestsStatus {
	switch r {
	case Approved:
		return domain.Approved
	case Rejected:
		return domain.Rejected
	case Unwatched:
		return domain.Unwatched
	default:
		return 0
	}
}