package services

import (
    "sync"
    "github.com/app8izer/go-gin-ng6-starter/backend/repositories"
    "github.com/app8izer/go-gin-ng6-starter/backend/models"
    "log"
)

// Singleton setup
var singleton *UserService
var once sync.Once

type UserService struct {
    userRepository *repositories.UserRepository
}

func GetUserServiceInstance() (*UserService) {
    once.Do(func() {
        singleton = &UserService{}
        singleton.userRepository = repositories.GetUserRepositoryInstance()
    })

    // create initial admin user if not existent in db
    if user := singleton.GetByUsername("admin"); user == nil {

        admin := &models.User{
            Username: "admin",
            Password: "1234",
            Role: models.User_ROLE_ADMIN,
        }

        singleton.Create(admin)
    }

    return singleton
}

// Service methods
func (us *UserService) GetByUsername(username string) (*models.User) {
    return us.userRepository.GetByUsername(username)
}

func (us *UserService) GetById(id string) (*models.User) {
    return us.userRepository.GetById(id)
}

func (us *UserService) Create(user *models.User) (*models.User) {
    hash, err := HashPassword(user.Password)
    if err != nil {
        log.Fatal(err)
        return user
    }

    user.Password = hash


    return us.userRepository.Create(user)
}

func (us *UserService) GetAll() ([]*models.User) {
    return us.userRepository.GetAll()
}


