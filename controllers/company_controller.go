package controllers

import (
	"reservation-system/models/dto"
	"reservation-system/services"
	"reservation-system/utils"
	"strconv"

	"github.com/kataras/iris/v12"
)

// Crear una instancia del validador
//var validate = validator.New()

type CompanyController struct {
	Service *services.CompanyService
}

func init() {
	// Registrar la validación personalizada para el correo
	validate.RegisterValidation("customEmail", utils.ValidateEmail)
}

func (c *CompanyController) RegisterCompany(ctx iris.Context) {
	var company dto.CompanyDTO
	err := ctx.ReadJSON(&company)
	if err != nil {
		utils.HandleBadRequest(ctx, err)
		return
	}
	// Validar el DTO usando la función de utilidades
	err = validate.Struct(company)
	if utils.HandleValidationError(ctx, err) {
		return
	}
	// Intentar registrar al company usando el servicio
	createdCompany, err := c.Service.Register(company)
	if err != nil {
		utils.HandleFound(ctx, err)
		return
	}
	// Si no hay errores, retornar el company creado con un código de éxito
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "Register saved successfully", "company": createdCompany})
}

func (c *CompanyController) UpdateCompany(ctx iris.Context) {

	//leer Json
	var company dto.CompanyDTO
	err := ctx.ReadJSON(&company)
	if err != nil {
		utils.HandleBadRequest(ctx, err)
		return
	}
	// Validar el DTO usando la función de utilidades
	err = validate.Struct(company)
	if utils.HandleValidationError(ctx, err) {
		// Si hubo errores de validación, ya se manejaron, simplemente retornar
		return
	}
	// Intentar actualizar al company usando el servicio
	updateCompany, err := c.Service.Update(company)
	if err != nil {
		utils.HandleNotFound(ctx, err)
		return
	}
	// Si no hay errores, retornar el company creado con un código de éxito
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Register updated successfully", "company": updateCompany})
}

func (c *CompanyController) GetAllCompanies(ctx iris.Context) {
	companies, err := c.Service.GetAll()
	utils.HandleInternalServerError(ctx, err)
	ctx.JSON(companies)
}

func (c *CompanyController) GetCompanyByEmail(ctx iris.Context) {
	email := ctx.Params().Get("email")
	company, err := c.Service.GetCompanyByEmail(email)
	if err != nil {
		utils.HandleNotFound(ctx, err)
		return
	}
	ctx.StatusCode(iris.StatusFound)
	ctx.JSON(iris.Map{"message": "Register found successfully", "company": company})

}

func (c *CompanyController) DeleteCompany(ctx iris.Context) {
	id := ctx.Params().Get("id")
	companyID, err := strconv.Atoi(id)
	company, err := c.Service.DeleteCompanyById(companyID)
	if err != nil {
		utils.HandleNotFound(ctx, err)
		return
	}
	ctx.StatusCode(iris.StatusFound)
	ctx.JSON(iris.Map{"message": "Register deleted successfully", "company": company})
}
