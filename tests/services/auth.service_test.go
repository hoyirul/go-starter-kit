package tests

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hoyirul/go-starter-kit/internal/models"
	"github.com/hoyirul/go-starter-kit/internal/services"
	mocks "github.com/hoyirul/go-starter-kit/mocks/services"
	"github.com/hoyirul/go-starter-kit/utils"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAuthRepository(ctrl)
	authService := services.NewAuthService(mockRepo)

	user := &models.User{
		Name:     "John Doe",
		Email:    "john@mail.com",
		Password: "password",
	}

	t.Run("Successful registration", func(t *testing.T) {
		hashedPassword, _ := utils.HashPassword(user.Password)
		// buat copy user supaya tidak mengubah user asli
		userToCreate := user
		userToCreate.Password = hashedPassword

		mockRepo.EXPECT().
			Create(gomock.Any()).
			Return(nil)

		err := authService.Register(user)

		assert.NoError(t, err)
	})

	t.Run("Error during registration", func(t *testing.T) {
		hashedPassword, _ := utils.HashPassword(user.Password)
		userToCreate := user
		userToCreate.Password = hashedPassword

		mockRepo.EXPECT().
			Create(gomock.Any()).
			Return(errors.New("registration error"))

		err := authService.Register(user)

		assert.Error(t, err)
		assert.Equal(t, "registration error", err.Error())
	})
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAuthRepository(ctrl)
	authService := services.NewAuthService(mockRepo)

	plainPassword := "password"
	hashedPassword, _ := utils.HashPassword(plainPassword)
	user := &models.User{
		Email:    "john@mail.com",
		Password: hashedPassword,
	}

	t.Run("Successful login", func(t *testing.T) {
		mockRepo.EXPECT().
			FindByEmail(user.Email).
			Return(user, nil)

		result, err := authService.Login(user.Email, plainPassword)

		assert.NoError(t, err)
		assert.Equal(t, user.Email, result.Email)
	})

	t.Run("Invalid credentials", func(t *testing.T) {
		mockRepo.EXPECT().
			FindByEmail(user.Email).
			Return(user, nil)

		result, err := authService.Login(user.Email, "wrongpassword")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "invalid credentials", err.Error())
	})

	t.Run("User not found", func(t *testing.T) {
		mockRepo.EXPECT().
			FindByEmail(user.Email).
			Return(nil, errors.New("user not found"))

		result, err := authService.Login(user.Email, plainPassword)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "user not found", err.Error())
	})

	t.Run("Error finding user by email", func(t *testing.T) {
		mockRepo.EXPECT().
			FindByEmail(user.Email).
			Return(nil, errors.New("database error"))

		result, err := authService.Login(user.Email, plainPassword)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "database error", err.Error())
	})
}

func TestFindByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAuthRepository(ctrl)
	authService := services.NewAuthService(mockRepo)

	user := &models.User{
		Email: "john@mail.com",
	}

	t.Run("Found user by email", func(t *testing.T) {
		mockRepo.EXPECT().
			FindByEmail(user.Email).
			Return(user, nil)

		result, err := authService.FindByEmail(user.Email)

		assert.NoError(t, err)
		assert.Equal(t, user.Email, result.Email)
	})

	t.Run("User not found by email", func(t *testing.T) {
		mockRepo.EXPECT().
			FindByEmail(user.Email).
			Return(nil, errors.New("user not found"))

		result, err := authService.FindByEmail(user.Email)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "user not found", err.Error())
	})
}

func TestFindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAuthRepository(ctrl)
	authService := services.NewAuthService(mockRepo)

	user := &models.User{
		ID: "123",
	}

	t.Run("Found user by ID", func(t *testing.T) {
		mockRepo.EXPECT().
			FindByID(user.ID).
			Return(user, nil)

		result, err := authService.FindByID(user.ID)

		assert.NoError(t, err)
		assert.Equal(t, user.ID, result.ID)
	})

	t.Run("User not found by ID returns error", func(t *testing.T) {
		mockRepo.EXPECT().
			FindByID(user.ID).
			Return(nil, errors.New("user not found"))

		result, err := authService.FindByID(user.ID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "user not found", err.Error())
	})

	t.Run("User is nil but no error from repo", func(t *testing.T) {
		mockRepo.EXPECT().
			FindByID(user.ID).
			Return(nil, nil)

		result, err := authService.FindByID(user.ID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "user not found", err.Error())
	})
}
