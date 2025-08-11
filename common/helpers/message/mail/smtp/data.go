package smtpHelper

import (
	"api/common/constants"
	"api/common/helpers"
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

type EmailData struct {
	HomePageLink string
	Logo         string
	Title        string
	Message      string
}

type EmailDataCheckCode struct {
	EmailData
	Code            string
	DurationMinutes int
}

func loadTemplate(templateFileName string, data any) (body []byte, err error) {
	if data == nil {
		errMsg := "No provided data!"
		err = fmt.Errorf("%s", errMsg)
		return
	}
	helpers.Logger.Info("Loading templates...")
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	AssetMailPath := filepath.Join(exeDir, constants.AssetMailPath)
	helpers.Logger.Info("Directory:", zap.String("Base DIR", exeDir))

	tmpl, err := template.ParseFiles(
		filepath.Join(AssetMailPath, "layout/base.html"),
		filepath.Join(AssetMailPath, templateFileName),
	)
	if err != nil {
		errMsg := "Error loading templates!"
		err = fmt.Errorf("%s: %s %w", errMsg, err.Error(), err)
		return
	}

	var buf bytes.Buffer
	if err = tmpl.ExecuteTemplate(&buf, "layout", data); err != nil {
		errMsg := "Error executing template!"
		err = fmt.Errorf("%s: %s %w", errMsg, err.Error(), err)
		return
	}

	return buf.Bytes(), nil
}

func (data *EmailData) LoadTemplate() (body []byte, err error) {
	body, err = loadTemplate("default.html", data)
	return
}

func (data *EmailDataCheckCode) LoadTemplate() (body []byte, err error) {
	body, err = loadTemplate("check-code.html", data)
	return
}
