package smtpHelper

import (
	"api/common/constants"
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
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

var assetPath string

func checkAssetMailPath() {
	if len(assetPath) > 0 {
		return
	}
	// Get the absolute path of the executable
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	AssetMailPath := filepath.Join(exeDir, constants.AssetMailPath)

	// Parse files
	_, err := template.ParseFiles(
		filepath.Join(AssetMailPath, "layout/base.html"),
	)
	if err != nil {
		_, err2 := template.ParseFiles(
			filepath.Join(constants.AssetMailPath, "layout/base.html"),
		)
		if err2 != nil {
			return
		}
		assetPath = constants.AssetMailPath
		return
	}

	// Update the path
	assetPath = AssetMailPath
}

func loadTemplate(templateFileName string, data any) (body []byte, err error) {
	if data == nil {
		errMsg := "No provided data!"
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check the correct path
	checkAssetMailPath()
	if len(assetPath) < 1 {
		errMsg := "Invalid asset path!"
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Parse files
	tmpl, err := template.ParseFiles(
		filepath.Join(assetPath, "layout/base.html"),
		filepath.Join(assetPath, templateFileName),
	)
	if err != nil {
		errMsg := "Error loading templates!"
		err = fmt.Errorf("%s: %s %w", errMsg, err.Error(), err)
		return
	}

	// Load template
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
