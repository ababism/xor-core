package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestActor(t *testing.T) {
	t.Parallel()
	// Arrange a new actor for testing
	const (
		role1        = "role1"
		role2        = "role2"
		missingRole1 = "missingRole1"
	)
	actorID := uuid.New()
	roles := []string{role1, role2}
	actor := NewActor(actorID, roles)

	t.Run("HasOneOfRoles", func(t *testing.T) {
		t.Parallel()

		// Test if the actor has one of the specified roles
		assert.True(t, actor.HasOneOfRoles(role1))
		assert.True(t, actor.HasOneOfRoles(role2))
		assert.False(t, actor.HasOneOfRoles(missingRole1))
	})

	t.Run("HasOneOfRolesEmptyRoles", func(t *testing.T) {
		t.Parallel()

		// Test bad scenario
		assert.True(t, actor.HasOneOfRoles())
	})

	t.Run("HasOneOfRolesEmptyActor", func(t *testing.T) {
		t.Parallel()

		emptyActor := NewActor(actorID, make([]string, 0))
		// Test if the actor has one of the specified roles
		assert.False(t, emptyActor.HasOneOfRoles(role1))
		assert.False(t, emptyActor.HasOneOfRoles(missingRole1))
	})

	t.Run("HasAllRoles", func(t *testing.T) {
		t.Parallel()

		// Test if the actor has all the specified roles
		assert.True(t, actor.HasAllRoles(role1, role2))
		assert.False(t, actor.HasAllRoles(role1, missingRole1))
		assert.False(t, actor.HasAllRoles(missingRole1))
	})

	t.Run("InitRoles", func(t *testing.T) {
		t.Parallel()

		// Test if the actor's roles are initialized correctly
		require.Equal(t, 2, actor.roles.Size(), "unexpected number of roles")
		assert.True(t, actor.roles.Contains(role1))
		assert.True(t, actor.roles.Contains(role2))
	})
}
