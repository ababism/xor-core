package keys

import (
	"fmt"
	"xor-go/services/courses/internal/domain"
)

const (
	CourseIDAttributeKey       = "course_id"
	LessonIDAttributeKey       = "lesson_id"
	LessonAccessIDAttributeKey = "lesson_access_id"
	ProductIDAttributeKey      = "product_id"
	ProductSliceAttributeKey   = "product_slice"
	StudentIDAttributeKey      = "student_id"
	TeacherIDAttributeKey      = "teacher_id"
	ActorIDAttributeKey        = "actor_id"
	ActorRolesAttributeKey     = "actor_roles"
)

func ProductToStrings(products []domain.Product) []string {
	stringProducts := make([]string, 0, len(products))
	for _, p := range products {
		stringProducts = append(stringProducts, fmt.Sprintf("product: %s of %s", p.ID.String(), p.Item.String()))
	}
	return stringProducts
}
