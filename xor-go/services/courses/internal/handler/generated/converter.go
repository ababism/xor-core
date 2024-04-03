package generated

import (
	"xor-go/services/courses/internal/domain"
)

func ToVisibilityDomain(cv CourseVisibility) domain.Visibility {
	switch cv {
	case CourseVisibilityHidden:
		return domain.Hidden
	case CourseVisibilityUnlocked:
		return domain.Unlocked
	case CourseVisibilityVisible:
		return domain.Visible
	}
	return domain.Hidden
}
