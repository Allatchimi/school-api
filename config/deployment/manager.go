package configDeploy

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"api/common/helpers"
	"api/common/utils"
	"api/common/utils/security"
	"api/config"
	"api/services/school/common/school/model"

	"go.uber.org/zap"
)

const (
	// App env templates
	envTemplateContent = `
NEXT_PUBLIC_APP_NAME="{{ .Name }}"
NEXT_PUBLIC_APP_DESCRIPTION="{{ .Description }}"
NEXT_PUBLIC_WEBSITE_URL="{{ .WebsiteURL }}"

API_BASE_URL="{{ .ApiBaseUrl }}"
CDN_URL="{{ .CdnUrl }}"
CDN_KEY="{{ .CdnKey }}"

NEXT_AUTH_URL="{{ .NextAuthUrl }}"
NEXT_AUTH_SECRET="{{ .NextAuthSecret }}"

GOOGLE_CLIENT_ID="459098306223-7b9ln9s1s6ccv67r9mr2vp2o52j0f7hv.apps.googleusercontent.com"
GOOGLE_CLIENT_SECRET="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"

SCHOOL_ID={{ .SchoolID }}
SCHOOL_TYPE="{{ .SchoolType }}"
SCHOOL_API_KEY="{{ .SchoolApiKey }}"
`

	// App color templates
	colorTemplateContent = `
export const COLOR_PRIMARY = "{{ .Primary }}";
export const COLOR_PRIMARY_BG = "{{ .PrimaryBg }}";
export const COLOR_PRIMARY_BG_HOVER = "{{ .PrimaryBgHover }}";

export const COLORS = [COLOR_PRIMARY, COLOR_PRIMARY_BG, COLOR_PRIMARY_BG_HOVER];
`

	// Deployment templates
	domainNameDeploymentTemplateContent = `{{ .DomainName }}`
	protocolDeploymentTemplateContent   = `{{ .Protocol }}`
	domainCertDeploymentTemplateContent = `{{ .DomainCert }}`
	domainKeyDeploymentTemplateContent  = `{{ .DomainKey }}`
)

type AppEnvData struct {
	Name        string
	Description string
	WebsiteURL  string

	ApiBaseUrl string
	CdnUrl     string
	CdnKey     string

	NextAuthUrl    string
	NextAuthSecret string

	SchoolID     int64
	SchoolType   string
	SchoolApiKey string
}

type AppColorData struct {
	Primary        string
	PrimaryBg      string
	PrimaryBgHover string
}

type DeploymentData struct {
	DomainName string
	Protocol   string
	DomainCert string
	DomainKey  string
}

const (
	deploymentsDir = ".deployment-"
)

// DeploySchool generates configuration files and pushes them to the repository.
func DeploySchool(school *model.School) (ok bool, err error) {
	if school == nil || school.ID < 1 || school.Config == nil || len(school.Config.DomainName) < 1 {
		errMsg := "School is nil or have invalid fields!"
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Create temp directory
	tempDir, err := os.MkdirTemp("", deploymentsDir)
	if err != nil {
		errMsg := "Failed to create temp directory!"
		err = fmt.Errorf("%s: %s %w", errMsg, err.Error(), err)
		return
	}
	defer os.RemoveAll(tempDir)

	// Generate website URL
	websiteURL := fmt.Sprintf("%s://%s", school.Config.Protocol, school.Config.DomainName)
	// Generate API key using HMAC SHA256
	apiKey, err := security.GenerateHMAC_SHA256_Base64URL(
		fmt.Sprintf("%d", school.ID),
		config.Env.SchoolApiSecret,
	)
	if err != nil {
		errMsg := "Failed to generate API key!"
		err = fmt.Errorf("%s! Error: %s", errMsg, err.Error())
		return
	}
	if len(apiKey) < 1 {
		errMsg := "API key is empty!"
		err = fmt.Errorf("%s", errMsg)
		return
	}
	// Generate app environment variables
	envData := AppEnvData{
		Name:           school.Config.WebsiteTitle,
		Description:    school.Config.WebsiteDescription,
		WebsiteURL:     websiteURL,
		ApiBaseUrl:     config.Env.SchoolApiBaseURL,
		CdnUrl:         config.Env.SchoolCdnUrl,
		CdnKey:         config.Env.SchoolCdnKey,
		NextAuthUrl:    websiteURL,
		NextAuthSecret: security.GenerateRandomBase64(32),
		SchoolID:       school.ID,
		SchoolType:     school.Type,
		SchoolApiKey:   apiKey,
	}
	// Generate app color data
	colorData := AppColorData{
		Primary:        school.Config.ColorPrimary,
		PrimaryBg:      school.Config.ColorPrimaryBg,
		PrimaryBgHover: school.Config.ColorPrimaryBgHover,
	}
	// Generate deployment data
	deploymentData := DeploymentData{
		DomainName: school.Config.DomainName,
		Protocol:   school.Config.Protocol,
		DomainCert: strings.ReplaceAll(school.Config.DomainCert, `\n`, "\n"),
		DomainKey:  strings.ReplaceAll(school.Config.DomainKey, `\n`, "\n"),
	}
	// Define output directory structure
	outputDir := filepath.Join(tempDir, "schools", fmt.Sprintf("%d", school.ID))
	colorDir := filepath.Join(outputDir, "src", "lib", "api", "constants", "common")
	faviconDir := filepath.Join(outputDir, "src", "app")
	logosDir := filepath.Join(outputDir, "public", "images", "logos")
	deploymentDir := filepath.Join(outputDir, "conf")
	// Create required directories
	for _, dir := range []string{outputDir, colorDir, faviconDir, logosDir, deploymentDir} {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			errMsg := "Failed to create directory!"
			err = fmt.Errorf("%s: %s %w", errMsg, dir, err)
			return
		}
	}
	// Generate app files
	if err = helpers.RenderTemplate(filepath.Join(outputDir, ".env"), envTemplateContent, envData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(outputDir, ".env"), envTemplateContent, err)
		return
	}
	if err = helpers.RenderTemplate(filepath.Join(colorDir, "color.ts"), colorTemplateContent, colorData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(colorDir, "color.ts"), colorTemplateContent, err)
		return
	}
	// Generate deployment files
	if err = helpers.RenderTemplate(filepath.Join(deploymentDir, "domain.txt"), domainNameDeploymentTemplateContent, deploymentData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(deploymentDir, "domain.txt"), domainNameDeploymentTemplateContent, err)
		return
	}
	if err = helpers.RenderTemplate(filepath.Join(deploymentDir, "protocol.txt"), protocolDeploymentTemplateContent, deploymentData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(deploymentDir, "protocol.txt"), protocolDeploymentTemplateContent, err)
		return
	}
	if err = helpers.RenderTemplate(filepath.Join(deploymentDir, "domain.cert"), domainCertDeploymentTemplateContent, deploymentData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(deploymentDir, "domain.cert"), domainCertDeploymentTemplateContent, err)
		return
	}
	if err = helpers.RenderTemplate(filepath.Join(deploymentDir, "domain.key"), domainKeyDeploymentTemplateContent, deploymentData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(deploymentDir, "domain.key"), domainKeyDeploymentTemplateContent, err)
		return
	}
	// Download favicon if available
	if len(school.Favicon) > 0 {
		if err = utils.HttpDownloadFile(school.Favicon, filepath.Join(faviconDir, "favicon.ico")); err != nil {
			helpers.Logger.Error("Failed to download favicon!", zap.String("URL", school.Favicon), zap.String("Error", err.Error()))
		}
	}
	// Download logos if available
	if len(school.Logo) > 0 {
		if err = utils.HttpDownloadFile(school.Logo, filepath.Join(logosDir, "logo.png")); err != nil {
			helpers.Logger.Error("Failed to download logo!", zap.String("URL", school.Logo), zap.String("Error", err.Error()))
		}
	}
	if len(school.LogoWhite) > 0 {
		if err = utils.HttpDownloadFile(school.LogoWhite, filepath.Join(logosDir, "logo-white.png")); err != nil {
			helpers.Logger.Error("Failed to download logo!", zap.String("URL", school.LogoWhite), zap.String("Error", err.Error()))
		}
	}

	// Setup SSH key
	if runtime.GOOS != "windows" {
		err = helpers.GitSetupSSHKey()
		if err != nil {
			helpers.Logger.Error("Failed to setup SSH key!", zap.String("Error", err.Error()))
			return
		}
		helpers.Logger.Info("SSH key setup successful.")
	}

	// Push deployment
	ok, err = helpers.GitPushSchoolDeployment(fmt.Sprintf("%d", school.ID), tempDir, outputDir)
	if err != nil {
		helpers.Logger.Error("Failed to push school deployment!", zap.String("Error", err.Error()))
		return false, err
	}
	if !ok {
		helpers.Logger.Warn(fmt.Sprintf("School deployment skipped! No changes detected for school %d!", school.ID))
	} else {
		helpers.Logger.Info(fmt.Sprintf("School deployment successfully added for school %d!", school.ID))
	}
	return
}

// DeleteSchoolDeployment deletes a school deployment from the repository.
func DeleteSchoolDeployment(schoolID int64) (ok bool, err error) {
	if schoolID < 1 {
		errMsg := "School ID is invalid!"
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Create temp directory
	tempDir, err := os.MkdirTemp("", deploymentsDir)
	if err != nil {
		errMsg := "Failed to create temp directory!"
		err = fmt.Errorf("%s: %s %w", errMsg, err.Error(), err)
		return
	}
	defer os.RemoveAll(tempDir)

	// Setup SSH key
	if runtime.GOOS != "windows" {
		err = helpers.GitSetupSSHKey()
		if err != nil {
			helpers.Logger.Error("Failed to setup SSH key!", zap.String("Error", err.Error()))
			return
		}
		helpers.Logger.Info("SSH key setup successful.")
	}

	// Push deployment to delete school
	ok, err = helpers.GitPushDeletedSchoolDeployment(fmt.Sprintf("%d", schoolID), tempDir)
	if err != nil {
		helpers.Logger.Error("Failed to push deleted school deployment!", zap.String("Error", err.Error()))
		return false, err
	}
	if !ok {
		helpers.Logger.Warn(fmt.Sprintf("School deployment deletion skipped! No changes detected for school %d!", schoolID))
	} else {
		helpers.Logger.Info(fmt.Sprintf("School deployment successfully deleted for school %d!", schoolID))
	}
	return
}
