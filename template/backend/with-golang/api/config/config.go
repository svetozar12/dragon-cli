package config

import (
	"svetozar12/headless-cms-be/db"
	"svetozar12/headless-cms-be/routes/content"
	contentmodel "svetozar12/headless-cms-be/routes/contentModel"
	"svetozar12/headless-cms-be/routes/field"
	fieldtype "svetozar12/headless-cms-be/routes/fieldType"
)

func Init() {
	db.DB.AutoMigrate(&contentmodel.ContentModel{})
	db.DB.AutoMigrate(&fieldtype.FieldType{}, &field.Field{}, &content.Content{})
}
