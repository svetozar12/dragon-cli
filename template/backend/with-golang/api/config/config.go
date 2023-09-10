package config

import (
	"@dragon-cli-template/db"
	"@dragon-cli-template/routes/content"
	contentmodel "@dragon-cli-template/routes/contentModel"
	"@dragon-cli-template/routes/field"
	fieldtype "@dragon-cli-template/routes/fieldType"
)

func Init() {
	db.DB.AutoMigrate(&contentmodel.ContentModel{})
	db.DB.AutoMigrate(&fieldtype.FieldType{}, &field.Field{}, &content.Content{})
}
