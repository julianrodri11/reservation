package services

import (
	"fmt"
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
func (s *UserService) Register(userDTO dto.UserDTO) (*dto.UserDTO, error) {
	// Verificar si el usuario ya existe por correo electrónico
	if _, err := s.Repo.FindByEmail(userDTO.Email); err == nil {
		// Si no hay error, significa que el usuario ya existe
		return nil, fmt.Errorf("el usuario con correo %s ya está registrado", userDTO.Email)
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

	// Retornar el DTO del usuario creado
	return &createdUserDTO, err
}

// actualizar un usuario
func (s *UserService) Update(userDTO dto.UserDTO) error {

	var userEntity entity.Users
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	userDTO.Password = string(hashedPassword)
	utils.ConvertDTOtoEntity(&userDTO, &userEntity)

	// Validar si el usuario existe antes de proceder con la actualización
	exists, err := s.Repo.UserExistsById(int(userDTO.ID))
	if err != nil {
		return err // Devuelve el error si hubo problemas al consultar la existencia del usuario
	}

	if !exists {
		return fmt.Errorf("el usuario con ID %d no existe", userDTO.ID) // Retorna un error si el usuario no existe
	}

	return s.Repo.UpdateUser(userEntity)
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

	return userDTO, nil
}
