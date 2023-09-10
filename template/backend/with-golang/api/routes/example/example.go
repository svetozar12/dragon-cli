package content

import (
	"strconv"
	"svetozar12/headless-cms-be/db"
	"svetozar12/headless-cms-be/models"

	"github.com/gofiber/fiber/v2"
)

type Body struct {
	Message int `json:"message" binding:"required"`
}

type Content struct {
	models.Model
	Body
	ContentModel contentmodel.ContentModel `json:"contentModel" binding:"required" gorm:"foreignKey:ModelId"`
}

func ContentRoutes(app fiber.Router) {
	content := app.Group("/content")
	content.Get("/", getContent)
	content.Post("/", createContent)
}

// Content godoc
// @Summary      Get all content
// @Tags         content
// @Accept       json
// @Param        page     query     int  false  "page"   default(1)
// @Param        limit    query     int  false  "limit"  default(10)
// @Param        userId  query     string  true   "userId"
// @Success      200  {object} models.PaginationModel[[]content.Content]
// @Router       /v1/content [get]
func getContent(c *fiber.Ctx) error {
	var content []Content
	var total int64
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	userId := c.Query("userId")
	offSet := (page - 1) * limit
	db.DB.Where("user_id = ?", userId).Preload("ContentModel").Find(&content).Count(&total)
	db.DB.Where("user_id = ?", userId).Preload("ContentModel").Offset(offSet).Limit(limit).Find(&content)
	return c.Status(fiber.StatusOK).JSON(models.PaginationModel[[]Content]{Pagination: models.Pagination{Total: total, Offset: page, Limit: limit}, Data: content})
}

// Content godoc
// @Summary      Create content
// @Tags         content
// @Accept       json
// @Param request body content.Body true "query params""
// @Success      201  {object} content.Content
// @Router       /v1/content [post]
func createContent(c *fiber.Ctx) error {
	content := new(Content)
	var fieldTypes []fieldtype.FieldType
	err := c.BodyParser(content)
	if err != nil {
		return c.SendStatus(fiber.ErrUnprocessableEntity.Code)
	}

	db.DB.Where("content_model_id = ?", content.ModelId).Find(&fieldTypes)
	db.DB.Create(&content)
	db.DB.Preload("ContentModel").First(&content, content.ID)

	var fields []field.Field
	for i := 0; i < len(fieldTypes); i++ {
		fieldType := fieldTypes[i]
		field := field.Field{FieldType: fieldType, Body: field.Body{Name: fieldType.Name, Value: "", TypeId: int(fieldType.ID), ContentId: int(content.ID)}}
		fields = append(fields, field)
	}
	db.DB.Create(&fields)
	return c.Status(fiber.StatusCreated).JSON(content)
}
