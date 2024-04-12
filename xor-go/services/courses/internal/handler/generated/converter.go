package generated

import (
	"github.com/google/uuid"
	"time"
	"xor-go/services/courses/internal/domain"
)

func (a OptionalActor) ToDomainWithRoles(roles []string) domain.Actor {
	if a.Roles == nil {
		s := []string{domain.UnregisteredRole}
		aID := uuid.Nil
		if a.ID != nil {
			aID = *a.ID
		}
		return domain.NewActor(aID, s)
	}
	roles = append(roles, *a.Roles)
	aID := uuid.Nil
	if a.ID != nil {
		aID = *a.ID

	}
	return domain.NewActor(aID, roles)
}

func (a OptionalActor) ToDomain() domain.Actor {
	if a.Roles == nil {
		s := []string{domain.UnregisteredRole}
		aID := uuid.Nil
		if a.ID != nil {
			aID = *a.ID
		}
		return domain.NewActor(aID, s)
	}
	s := []string{*a.Roles}
	aID := uuid.Nil
	if a.ID != nil {
		aID = *a.ID

	}
	return domain.NewActor(aID, s)
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
	cID := uuid.Nil
	if c.ID != nil {
		cID = *c.ID

	}
	cFeedbackID := uuid.Nil
	if c.FeedbackID != nil {
		cFeedbackID = *c.FeedbackID
	}
	cTeacherID := uuid.Nil
	if c.TeacherID != nil {
		cTeacherID = *c.TeacherID

	}
	return &domain.Course{
		ID:         cID,
		FeedbackID: cFeedbackID,
		TeacherID:  cTeacherID,
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
	sID := uuid.Nil
	if s.ID != nil {
		sID = *s.ID
	}
	return domain.Section{
		Description: s.Description,
		//FeedbackID:  s.FeedbackID,
		Heading:    s.Heading,
		ID:         sID,
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
	tID := uuid.Nil
	if t.ID != nil {
		tID = *t.ID

	}
	return domain.Theme{
		//FeedbackID: t.FeedbackID,
		Heading:    t.Heading,
		ID:         tID,
		LessonIDs:  *t.LessonIDs,
		Visibility: t.Visibility.ToDomain(),
	}
}

func (l Lesson) ToDomain() *domain.Lesson {
	pr := domain.Product{}
	if l.Product != nil {
		pr = l.Product.ToDomain()
	}
	uri := ""
	if l.VideoURI != nil {
		uri = *l.VideoURI
	}
	tID := uuid.Nil
	if l.TeacherID != nil {
		tID = *l.TeacherID
	}
	lID := uuid.Nil
	if l.ID != nil {
		lID = *l.ID
	}
	return &domain.Lesson{
		CourseID:   l.CourseID,
		ID:         lID,
		Product:    pr,
		TeacherID:  tID,
		Transcript: l.Transcript,
		VideoURI:   uri,
		Visibility: l.Visibility.ToDomain(),
	}
}

func (p Product) ToDomain() domain.Product {
	pID := uuid.Nil
	if p.ID != nil {
		pID = *p.ID
	}
	pItem := uuid.Nil
	if p.Item != nil {
		pItem = *p.Item
	}
	pOwner := uuid.Nil
	if p.Owner != nil {
		pOwner = *p.Owner
	}
	return domain.Product{
		ID:    pID,
		Item:  pItem,
		Owner: pOwner,
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
	pID := uuid.Nil
	if p.ID != nil {
		pID = *p.ID
	}

	return domain.PublicationRequest{
		AssigneeID:    p.AssigneeID,
		Comment:       p.Comment,
		CourseID:      p.CourseID,
		ID:            pID,
		RequestStatus: ToDomainRequestStatus(p.RequestStatus),
		UpdatedAt:     ToTimeDomain(p.UpdatedAt),
	}
}
func (pr Teacher) ToDomain() domain.Teacher {
	return domain.Teacher{
		AccountID: pr.AccountID,
	}
}

func ToDomainRequestStatus(s *PublicationRequestRequestStatus) domain.RequestsStatus {
	if s == nil {
		return domain.Unwatched
	}
	return (*s).ToDomain()

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
	lID := uuid.Nil
	if l.ID != nil {
		lID = *l.ID

	}
	return domain.LessonAccess{
		ID:           lID,
		LessonID:     l.LessonID,
		StudentID:    l.StudentID,
		AccessStatus: ToAccessStatus(l.AccessStatus),
		UpdatedAt:    ToTimeDomain(updatedAt),
	}
}

func ToAccessStatus(accessStatus LessonAccessAccessStatus) domain.AccessStatus {
	switch accessStatus {
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
