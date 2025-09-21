package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestNewUserAccount_ValidCreation(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		password string
	}{
		{
			name:     "Valid user with standard email and password",
			email:    "user@example.com",
			password: "strongPassword123",
		},
		{
			name:     "Valid user with complex email",
			email:    "user.name+tag@example.co.uk",
			password: "mySecretPassword!@#",
		},
		{
			name:     "Valid user with short password",
			email:    "test@test.com",
			password: "123456",
		},
		{
			name:     "Valid user with long password",
			email:    "long@example.com",
			password: "thisIsAVeryLongPasswordWithManyCharactersAndSpecialSymbols!@#$%^&*()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userAccount, err := NewUserAccount(tt.email, tt.password)

			// Verify no error occurred
			assert.NoError(t, err, "NewUserAccount should not return error for valid input")
			assert.NotNil(t, userAccount, "UserAccount should not be nil")

			// Verify fields are set correctly
			assert.Equal(t, tt.email, userAccount.Email(), "Email should match input")
			assert.NotEmpty(t, userAccount.PasswordHash(), "Password hash should not be empty")
			assert.NotEqual(t, tt.password, userAccount.PasswordHash(), "Password hash should not equal plain password")

			// Verify password is properly hashed
			err = bcrypt.CompareHashAndPassword([]byte(userAccount.PasswordHash()), []byte(tt.password))
			assert.NoError(t, err, "Password should be correctly hashed")

			// Verify ID is initialized (should be 0 for new entities)
			assert.Equal(t, int64(0), userAccount.ID(), "ID should be 0 for new entities")
		})
	}
}

func TestNewUserAccount_PasswordHashing(t *testing.T) {
	email := "test@example.com"
	password := "testPassword123"

	userAccount, err := NewUserAccount(email, password)
	require.NoError(t, err)
	require.NotNil(t, userAccount)

	t.Run("Password is properly hashed", func(t *testing.T) {
		// Verify password hash is not the plain password
		assert.NotEqual(t, password, userAccount.PasswordHash())

		// Verify password hash is not empty
		assert.NotEmpty(t, userAccount.PasswordHash())

		// Verify hash length (bcrypt hashes are 60 characters)
		assert.Len(t, userAccount.PasswordHash(), 60)

		// Verify hash starts with bcrypt identifier
		assert.True(t, len(userAccount.PasswordHash()) > 3)
		assert.Equal(t, "$2a", userAccount.PasswordHash()[:3])
	})

	t.Run("Password can be verified against hash", func(t *testing.T) {
		// Correct password should verify
		err := bcrypt.CompareHashAndPassword([]byte(userAccount.PasswordHash()), []byte(password))
		assert.NoError(t, err, "Correct password should verify against hash")

		// Incorrect password should fail verification
		err = bcrypt.CompareHashAndPassword([]byte(userAccount.PasswordHash()), []byte("wrongPassword"))
		assert.Error(t, err, "Incorrect password should fail verification")
	})

	t.Run("Multiple users with same password have different hashes", func(t *testing.T) {
		user1, err1 := NewUserAccount("user1@example.com", password)
		user2, err2 := NewUserAccount("user2@example.com", password)

		require.NoError(t, err1)
		require.NoError(t, err2)
		require.NotNil(t, user1)
		require.NotNil(t, user2)

		// Different users with same password should have different hashes
		assert.NotEqual(t, user1.PasswordHash(), user2.PasswordHash(),
			"Different users with same password should have different hashes")
	})
}

func TestNewUserAccount_EdgeCases(t *testing.T) {
	t.Run("Empty email", func(t *testing.T) {
		userAccount, err := NewUserAccount("", "password123")

		assert.Error(t, err, "NewUserAccount should return error for empty email")
		assert.Nil(t, userAccount, "UserAccount should be nil for empty email")
	})

	t.Run("Empty password", func(t *testing.T) {
		userAccount, err := NewUserAccount("test@example.com", "")

		assert.Error(t, err, "NewUserAccount should return error for empty password")
		assert.Nil(t, userAccount, "UserAccount should be nil for empty password")
	})

	t.Run("Very long email", func(t *testing.T) {
		longEmail := "very.long.email.address.with.many.dots@very-long-domain-name.example.com"
		userAccount, err := NewUserAccount(longEmail, "password123")

		assert.NoError(t, err)
		assert.NotNil(t, userAccount)
		assert.Equal(t, longEmail, userAccount.Email())
	})
}

func TestUserAccount_Getters(t *testing.T) {
	email := "test@example.com"
	password := "testPassword123"

	userAccount, err := NewUserAccount(email, password)
	require.NoError(t, err)
	require.NotNil(t, userAccount)

	t.Run("ID getter", func(t *testing.T) {
		id := userAccount.ID()
		assert.Equal(t, int64(0), id, "ID should return the internal id field")
		assert.IsType(t, int64(0), id, "ID should return int64 type")
	})

	t.Run("Email getter", func(t *testing.T) {
		emailResult := userAccount.Email()
		assert.Equal(t, email, emailResult, "Email should return the internal email field")
		assert.IsType(t, "", emailResult, "Email should return string type")
	})

	t.Run("PasswordHash getter", func(t *testing.T) {
		hashResult := userAccount.PasswordHash()
		assert.NotEmpty(t, hashResult, "PasswordHash should not be empty")
		assert.IsType(t, "", hashResult, "PasswordHash should return string type")
		assert.NotEqual(t, password, hashResult, "PasswordHash should not return plain password")
	})
}
