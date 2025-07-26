package abstraction_test

import (
	"testing"

	"github.com/huangsam/go-trial/lesson/abstraction"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserByID(t *testing.T) {
	mockRepo := new(abstraction.MockUserRepository)
	user := &abstraction.User{ID: 1, Name: "John Doe", Email: "john.doe@example.com"}

	mockRepo.On("GetUserByID", 1).Return(user, nil)

	result, err := mockRepo.GetUserByID(1)
	require.NoError(t, err)
	assert.Equal(t, user, result)

	mockRepo.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(abstraction.MockUserRepository)
	user := &abstraction.User{ID: 2, Name: "Jane Doe", Email: "jane.doe@example.com"}

	mockRepo.On("CreateUser", user).Return(nil)

	require.NoError(t, mockRepo.CreateUser(user))

	mockRepo.AssertExpectations(t)
}
