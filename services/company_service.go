package services

import (
	"fmt"
	"log"
	"reservation-system/models/dto"
	"reservation-system/models/entity"
	"reservation-system/repository"

	"reservation-system/utils"
)

type CompanyService struct {
	Repo *repository.CompanyRepository
}

// crear un empresa
func (s *CompanyService) Register(companyDTO dto.CompanyDTO) (*dto.CompanyDTO, error) {

	// Verificar si el empresa ya existe por correo electrónico
	if _, err := s.Repo.FindByEmail(companyDTO.Email); err == nil {
		// Si no hay error, significa que el empresa ya existe
		return nil, fmt.Errorf("La empresa con correo %s ya está registrado", companyDTO.Email)
	}

	var companyEntity entity.Company
	// Convertir el DTO a una entidad
	utils.ConvertDTOtoEntity(&companyDTO, &companyEntity)

	// Crear el empresa en la base de datos y devolver el resultado
	createdCompany, err := s.Repo.CreateCompany(companyEntity)
	if err != nil {
		return nil, fmt.Errorf("error al crear el empresa: %v", err)
	}

	// Convertir la entidad creada nuevamente a DTO para la respuesta
	var createdCompanyDTO dto.CompanyDTO
	utils.ConvertDTOtoEntity(createdCompany, &createdCompanyDTO)

	// Retornar el DTO del empresa creado
	return &createdCompanyDTO, err
}

// actualizar un empresa
func (s *CompanyService) Update(companyDTO dto.CompanyDTO) (*dto.CompanyDTO, error) {

	// Validar si el empresa existe antes de proceder con la actualización
	existe, err := s.Repo.CompanyExistsById(int(companyDTO.ID))
	if err != nil {
		return nil, fmt.Errorf("error al consultar el ID %d : ", companyDTO.ID) // Devuelve el error si hubo problemas al consultar la existencia del empresa
	}
	if !existe {
		return nil, fmt.Errorf("el empresa con ID %d no existe", companyDTO.ID)
	}

	var companyEntity entity.Company
	utils.ConvertDTOtoEntity(&companyDTO, &companyEntity)

	// Actualiza el empresa en la base de datos y devolver el resultado
	updatedCompany, err := s.Repo.UpdateCompany(companyEntity)
	if err != nil {
		return nil, fmt.Errorf("error al actualizar el empresa: %v", err)
	}

	// Convertir la entidad creada nuevamente a DTO para la respuesta
	var updatedCompanyDTO dto.CompanyDTO
	utils.ConvertDTOtoEntity(updatedCompany, &updatedCompanyDTO)

	return &updatedCompanyDTO, err
}

// Consultar todos los usuarios
func (s *CompanyService) GetAll() ([]dto.CompanyDTO, error) {

	companies, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}

	var companyDTOs []dto.CompanyDTO
	for _, company := range companies {
		var companyDTO dto.CompanyDTO
		utils.GenericMapper(&company, &companyDTO)

		companyDTOs = append(companyDTOs, companyDTO)
	}

	return companyDTOs, nil
}

// Consultar un empresa por correo
func (s *CompanyService) GetCompanyByEmail(email string) (dto.CompanyDTO, error) {

	company, err := s.Repo.FindByEmail(email)
	if err != nil {
		return dto.CompanyDTO{}, err
	}

	var companyDTO dto.CompanyDTO
	err = utils.GenericMapper(&company, &companyDTO)
	if err != nil {
		log.Println("Error mapping User entity to CompanyDTO:", err)
		return dto.CompanyDTO{}, err
	}
	return companyDTO, nil
}

// Eliminar un empresa por id
func (s *CompanyService) DeleteCompanyById(id int) (dto.CompanyDTO, error) {

	company, err := s.Repo.DeleteCompany(id)
	if err != nil {
		return dto.CompanyDTO{}, err
	}

	var companyDTO dto.CompanyDTO
	err = utils.GenericMapper(&company, &companyDTO)
	if err != nil {
		log.Println("Error mapping User entity to CompanyDTO:", err)
		return dto.CompanyDTO{}, err
	}
	return companyDTO, nil
}
