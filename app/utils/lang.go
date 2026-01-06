package utils

import (
	"encoding/json"
	"log"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func NewLang() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	files, err := filepath.Glob("lang/*.json")
	if err != nil {
		log.Fatalf("failed to list lang files: %v", err)
	}

	for _, file := range files {
		_, err := bundle.LoadMessageFile(file)
		if err != nil {
			log.Fatalf("failed to load lang file %s: %v", file, err)
		}
	}
}

func GetLang(c echo.Context) string {
	lang := c.Request().Header.Get("lang")
	if lang == "" {
		lang = "en"
	}
	return lang
}

func ReloadAll() error {
	files, err := filepath.Glob("lang/*.json")
	if err != nil {
		return err
	}

	for _, file := range files {
		_, err := bundle.LoadMessageFile(file)
		if err != nil {
			log.Printf("Reload error for %s: %v", file, err)
		}
	}

	return nil
}
