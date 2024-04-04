package domain

import (
	"github.com/google/uuid"
)

func (c *Actor) HasRole(role string) bool {
	return c.roles.Contains(role)
}

func (c *Course) ApplyVisibility() {
	if c.Visibility == Hidden {
		c.FeedbackID = uuid.Nil
		c.Landing = nil
		c.Sections = nil
	} else {
		for i := range c.Sections {
			c.Sections[i].ApplyVisibility()
		}
	}
}

func (c *Course) Validate() error {
	// TODO
	//apperror.New(http.StatusInternalServerError, message, "invalid course fields", err)
	return nil
}

func (s *Section) ApplyVisibility() {
	if s.Visibility == Hidden {
		s.Description = ""
		s.FeedbackID = uuid.Nil
		s.Themes = nil
	} else {
		for i := range s.Themes {
			s.Themes[i].ApplyVisibility()
		}
	}
}

func (t *Theme) ApplyVisibility() {
	if t.Visibility == Hidden {
		t.FeedbackID = uuid.Nil
		t.Lessons = nil
	} else {
		for i := range t.Lessons {
			t.Lessons[i].ApplyVisibility()
		}
	}
}

func (l *Lesson) ApplyVisibility() {
	if l.Visibility == Hidden {
		l.VideoURI = ""
		l.Transcript = ""
	}
}
