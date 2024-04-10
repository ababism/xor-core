package domain

import (
	"github.com/google/uuid"
	"net/http"
	"xor-go/pkg/xapperror"
	"xor-go/pkg/xstringset"
)

func NewActor(ID uuid.UUID, roles []string) Actor {
	a := Actor{ID: ID,
		roles: xstringset.New()}
	a.initRoles(roles)
	return a
}

func (c *Actor) GetRoles() []string {
	// TODO - does it necessary to return make([]string, 0)
	if c == nil || c.roles == nil {
		return make([]string, 0)
	}
	copyRoles := make([]string, 0, c.roles.Size())
	for r, _ := range c.roles {
		copyRoles = append(copyRoles, r)
	}
	return copyRoles
}

func (c *Actor) HasRole(role string) bool {
	return c.roles.Contains(role)
}

func (c *Actor) HasOneOfRoles(roles ...string) bool {
	if roles == nil || len(roles) == 0 {
		return true
	}

	for _, role := range roles {
		if c.roles.Contains(role) {
			return true
		}
	}
	return false
}

func (c *Actor) HasAllRoles(roles ...string) bool {
	if roles == nil || len(roles) == 0 {
		return true
	}

	for _, role := range roles {
		if !c.roles.Contains(role) {
			return false
		}
	}
	return true
}

func (c *Actor) initRoles(roles []string) {
	c.roles.AddItems(roles)
}

func (c *Course) ApplyVisibility() {
	if c.Visibility == Hidden {
		c.FeedbackID = uuid.Nil
		c.Landing = ""
		c.Sections = nil
	} else {
		for i := range c.Sections {
			c.Sections[i].ApplyVisibility()
		}
	}
}

func (s *Section) ApplyVisibility() {
	if s.Visibility == Hidden {
		s.Description = ""
		s.Themes = nil
	} else {
		for i := range s.Themes {
			s.Themes[i].ApplyVisibility()
		}
	}
}

func (t *Theme) ApplyVisibility() {
	if t.Visibility == Hidden {

		t.LessonIDs = nil
	}
}

func (l *Lesson) ApplyVisibility() {
	if l.Visibility == Hidden {
		l.VideoURI = ""
		l.Transcript = ""
	}
}

func (l *Lesson) ApplyPaywall() {
	l.VideoURI = ""
	l.Transcript = ""

}
func (c *Course) FillEmptyUUIDs() {
	if c.ID == uuid.Nil || (c.ID == uuid.UUID{}) {
		c.ID = uuid.New()
	}
	for i := range c.Sections {
		if c.Sections[i].ID == uuid.Nil || (c.Sections[i].ID == uuid.UUID{}) {
			c.Sections[i].FillEmptyUUIDs()
		}
	}
}
func (s *Section) FillEmptyUUIDs() {
	if s.ID == uuid.Nil || (s.ID == uuid.UUID{}) {
		s.ID = uuid.New()
	}
	for i := range s.Themes {
		if s.Themes[i].ID == uuid.Nil || (s.Themes[i].ID == uuid.UUID{}) {
			s.Themes[i].FillEmptyUUIDs()
		}
	}
}
func (t *Theme) FillEmptyUUIDs() {
	if t.ID == uuid.Nil || (t.ID == uuid.UUID{}) {
		t.ID = uuid.New()
	}
}

func (c *Course) Validate() error {
	if c.ID == uuid.Nil || (c.ID == uuid.UUID{}) {
		return xapperror.New(http.StatusInternalServerError, "courses id shouldn't be empty", "validate courses id shouldn't be empty", nil)
	}
	return nil
}

func (s *Section) Validate() error {
	// TODO
	//apperror.New(http.StatusInternalServerError, message, "invalid section fields", err)
	return nil
}

func (t *Theme) Validate() error {
	// TODO
	//apperror.New(http.StatusInternalServerError, message, "invalid theme fields", err)
	return nil
}

func (l *Lesson) Validate() error {
	if l.ID == uuid.Nil || (l.ID == uuid.UUID{}) {
		return xapperror.New(http.StatusInternalServerError, "lesson id shouldn't be empty", "validate lesson id shouldn't be empty", nil)
	}
	return nil
}

func (m PublicationRequest) Validate() error {
	if m.ID == uuid.Nil || (m.ID == uuid.UUID{}) {
		return xapperror.New(http.StatusInternalServerError, "publication request id shouldn't be empty", "validate publication request id shouldn't be empty", nil)
	}
	return nil
}

func (s Student) Validate() error {
	// TODO
	//apperror.New(http.StatusInternalServerError, message, "invalid student profile fields", err)
	return nil
}

func (t Teacher) Validate() error {
	// TODO
	//apperror.New(http.StatusInternalServerError, message, "invalid teacher profile fields", err)
	return nil
}

func (p Product) Validate() error {
	// TODO
	//apperror.New(http.StatusInternalServerError, message, "invalid teacher profile fields", err)
	return nil
}
