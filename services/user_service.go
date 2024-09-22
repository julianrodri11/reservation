package services

import (
	"log"
	"reservation-system/models/dto"
	"reservation-system/models/entity"
	"reservation-system/repository"

	"reservation-system/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

// crear un usuario
func (s *UserService) Register(userDTO dto.UserDTO) error {

	var userEntity entity.User

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	userDTO.Password = string(hashedPassword)

	// Usar la función genérica para mapear el DTO a la entidad
	utils.ConvertDTOtoEntity(&userDTO, &userEntity)

	return s.Repo.Create(userEntity)
}

// Consultar todos los usuarios
func (s *UserService) GetAllUsers() ([]dto.UserDTO, error) {
	users, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}

	var userDTOs []dto.UserDTO
	for _, user := range users {
		var userDTO dto.UserDTO
		utils.GenericMapper(&user, &userDTO)
		userDTOs = append(userDTOs, userDTO)
	}

	return userDTOs, nil
}

// Consultar un usuario por correo
func (s *UserService) GetUserByEmail(email string) (dto.UserDTO, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return dto.UserDTO{}, err
	}

	var userDTO dto.UserDTO
	err = utils.GenericMapper(&user, &userDTO)
	if err != nil {
		log.Println("Error mapping User entity to UserDTO:", err)
		return dto.UserDTO{}, err
	}

	return userDTO, nil
}

// Consultar un usuario por correo
func (s *UserService) DeleteUserById(id int) (dto.UserDTO, error) {
	user, err := s.Repo.DeleteUser(id)
	if err != nil {
		return dto.UserDTO{}, err
	}

	var userDTO dto.UserDTO
	err = utils.GenericMapper(&user, &userDTO)
	if err != nil {
		log.Println("Error mapping User entity to UserDTO:", err)
		return dto.UserDTO{}, err
	}

	return userDTO, nil
}
