package repository

import (
	"reservation-system/models/entity"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	DB *gorm.DB
}

// Crear companias y devolver el companias creado
func (r *CompanyRepository) CreateCompany(company entity.Company) (*entity.Company, error) {
	err := r.DB.Create(&company).Error
	if err != nil {
		// Si hay un error al crear el companias, devolver nil y el error
		return nil, err
	}
	// Si la creación fue exitosa, devolver el companias y nil como error
	return &company, nil
}

// Actualizar companias existente
func (r *CompanyRepository) UpdateCompany(company entity.Company) (*entity.Company, error) {
	err := r.DB.Save(&company).Error
	if err != nil {
		// Si hay un error al actualziar el companias, devolver nil y el error
		return nil, err
	}
	// Si la actualización fue exitosa, devolver el companias y nil como error
	return &company, nil
}

// Consultar todos los companias
func (r *CompanyRepository) FindAll() ([]entity.Company, error) {
	var companies []entity.Company
	err := r.DB.Find(&companies).Error
	return companies, err
}

// consultar por correo
func (r *CompanyRepository) FindByEmail(email string) (entity.Company, error) {
	var company entity.Company
	err := r.DB.Where("email = ?", email).First(&company).Error
	return company, err
}

// eliminar por id
func (r *CompanyRepository) DeleteCompany(id int) (entity.Company, error) {
	var company entity.Company

	// Buscar el companias por ID
	err := r.DB.Where("id = ?", id).First(&company).Error
	if err != nil {
		return company, err
	}

	// Si lo encuentra, lo elimina
	err = r.DB.Delete(&company).Error
	return company, err
}

// consultar por id
func (r *CompanyRepository) FindById(id int) (entity.Company, error) {
	var company entity.Company
	err := r.DB.Where("id = ?", id).First(&company).Error
	return company, err
}

// Verificar si existe un companias por ID
func (r *CompanyRepository) CompanyExistsById(id int) (bool, error) {
	var company entity.Company
	err := r.DB.Where("id = ?", id).First(&company).Error
	if err != nil {
		// Si ocurre un error de tipo RecordNotFound, devolver false
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		// Si es otro tipo de error, devolver false y el error
		return false, err
	}
	// Si el companias es encontrado, devolver true
	return true, nil
}
