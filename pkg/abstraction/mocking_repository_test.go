package abstraction_test

import (
	"testing"

	abstraction_mock "github.com/huangsam/go-trial/mock/abstraction"
	"github.com/huangsam/go-trial/pkg/abstraction"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserByID(t *testing.T) {
	mockRepo := new(abstraction_mock.MockUserRepository)
	user := &abstraction.User{ID: 1, Name: "John Doe", Email: "john.doe@example.com"}

	mockRepo.On("GetUserByID", 1).Return(user, nil)

	result, err := mockRepo.GetUserByID(1)
	require.NoError(t, err)
	assert.Equal(t, user, result)

	mockRepo.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(abstraction_mock.MockUserRepository)
	user := &abstraction.User{ID: 2, Name: "Jane Doe", Email: "jane.doe@example.com"}

	mockRepo.On("CreateUser", user).Return(nil)

	err := mockRepo.CreateUser(user)
	require.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
