package model_test

import (
	"testing"

	"github.com/huangsam/go-trial/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestUserAccount(t *testing.T) {
	t.Run("create UserAccount", func(t *testing.T) {
		user := model.UserAccount{
			Username: "testuser",
			Password: "password123",
		}

		assert.Equal(t, "testuser", user.Username)
		assert.Equal(t, "password123", user.Password)
	})

	t.Run("update UserAccount", func(t *testing.T) {
		user := model.UserAccount{
			Username: "testuser",
			Password: "password123",
		}

		user.Username = "newuser"
		user.Password = "newpassword123"

		assert.Equal(t, "newuser", user.Username)
		assert.Equal(t, "newpassword123", user.Password)
	})
}
