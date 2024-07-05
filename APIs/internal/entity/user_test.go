package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Jon Snow", "jonjon@got.com>", "123456")
	assert.Nil(t, err)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jon Snow", user.Name)
	assert.Equal(t, "jonjon@got.com>", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Jon Snow", "jonjon@got.com>", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("<PASSWORD>"))
	assert.NotEqual(t, "123456", user.Password)
}
