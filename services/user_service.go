package services

import (
	"fmt"
	"log"
	"reservation-system/models/dto"
	"reservation-system/models/entity"
	"reservation-system/repository"
	"strings"

	"reservation-system/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

// crear un usuario
func (s *UserService) Register(userDTO dto.UserDTO) (*dto.UserDTO, error) {

	// Verificar si el usuario ya existe por correo electrónico
	if _, err := s.Repo.FindByEmail(strings.ToLower(userDTO.Email)); err == nil {
		// Si no hay error, significa que el usuario ya existe
		return nil, fmt.Errorf("el usuario con correo %s ya está registrado", strings.ToLower(userDTO.Email))
	}

	var userEntity entity.Users
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error al cifrar la contraseña: %v", err)
	}

	userDTO.Password = string(hashedPassword)
	// Convertir el DTO a una entidad
	utils.ConvertDTOtoEntity(&userDTO, &userEntity)

	// Crear el usuario en la base de datos y devolver el resultado
	createdUser, err := s.Repo.CreateUser(userEntity)
	if err != nil {
		return nil, fmt.Errorf("error al crear el usuario: %v", err)
	}

	// Convertir la entidad creada nuevamente a DTO para la respuesta
	var createdUserDTO dto.UserDTO
	utils.ConvertDTOtoEntity(createdUser, &createdUserDTO)
	createdUserDTO.Password = ""

	// Retornar el DTO del usuario creado
	return &createdUserDTO, err
}

// actualizar un usuario
func (s *UserService) Update(userDTO dto.UserDTO) (*dto.UserDTO, error) {

	// Validar si el usuario existe antes de proceder con la actualización
	existe, err := s.Repo.UserExistsById(int(userDTO.ID))
	if err != nil {
		return nil, fmt.Errorf("error al consultar el ID %d : ", userDTO.ID) // Devuelve el error si hubo problemas al consultar la existencia del usuario
	}
	if !existe {
		return nil, fmt.Errorf("el usuario con ID %d no existe", userDTO.ID)
	}

	var userEntity entity.Users
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	userDTO.Password = string(hashedPassword)

	utils.ConvertDTOtoEntity(&userDTO, &userEntity)

	// Actualiza el usuario en la base de datos y devolver el resultado
	updatedUser, err := s.Repo.UpdateUser(userEntity)
	if err != nil {
		return nil, fmt.Errorf("error al actualizar el usuario: %v", err)
	}

	// Convertir la entidad creada nuevamente a DTO para la respuesta
	var updatedUserDTO dto.UserDTO
	utils.ConvertDTOtoEntity(updatedUser, &updatedUserDTO)
	updatedUserDTO.Password = ""

	return &updatedUserDTO, err
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
		userDTO.Password = ""
		userDTOs = append(userDTOs, userDTO)
	}

	return userDTOs, nil
}

// Consultar un usuario por correo
func (s *UserService) GetUserByEmail(email string) (dto.UserDTO, error) {

	user, err := s.Repo.FindByEmail(strings.ToLower(email))
	if err != nil {
		return dto.UserDTO{}, err
	}

	var userDTO dto.UserDTO
	err = utils.GenericMapper(&user, &userDTO)
	if err != nil {
		log.Println("Error mapping User entity to UserDTO:", err)
		return dto.UserDTO{}, err
	}
	userDTO.Password = ""
	return userDTO, nil
}

// Eliminar un usuario por id
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
	userDTO.Password = ""
	return userDTO, nil
}
