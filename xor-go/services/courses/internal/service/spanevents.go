package service

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/domain/keys"
)

func ToSpan(span *trace.Span, a domain.Actor) {
	if span == nil {
		return
	}
	(*span).SetAttributes(attribute.String(keys.ActorIDAttributeKey, a.ID.String()))
	(*span).SetAttributes(attribute.StringSlice(keys.ActorRolesAttributeKey, a.GetRoles()))

	(*span).AddEvent("actor logged in for service method")
}
