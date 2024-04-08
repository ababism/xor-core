package domain

import (
	"github.com/google/uuid"
	"xor-go/pkg/xstringset"
)

func NewActor(ID uuid.UUID, roles []string) Actor {
	a := Actor{ID: ID,
		roles: xstringset.New()}
	a.initRoles(roles)
	return a
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

func (c *Course) Validate() error {
	// TODO
	//apperror.New(http.StatusInternalServerError, message, "invalid course fields", err)
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
	// TODO
	//apperror.New(http.StatusInternalServerError, message, "invalid lesson fields", err)
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
