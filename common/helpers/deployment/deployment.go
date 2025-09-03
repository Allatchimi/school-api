package deploymentHelper

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"api/common/helpers"
	htmlHelper "api/common/helpers/html"
	httpHelper "api/common/helpers/http"
	securityUtil "api/common/utils/security"
	"api/config"
	"api/services/school/common/school/model"

	"go.uber.org/zap"
)

const (
	baseDir = ".deployment-"
)

// DeploySchool generates configuration files and pushes them to the repository.
func DeploySchool(school *model.School) (err error) {
	if school == nil || school.ID < 1 || school.Config == nil || len(school.Config.WebsiteDomainName) < 1 {
		errMsg := "School is nil or have invalid fields!"
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Create temp directory
	tempDir, err := os.MkdirTemp("", baseDir)
	if err != nil {
		errMsg := "Failed to create temp directory!"
		err = fmt.Errorf("%s: %s %w", errMsg, err.Error(), err)
		return
	}
	defer os.RemoveAll(tempDir)

	// Generate website URL
	websiteURL := fmt.Sprintf("%s://%s", httpHelper.DefaultProtocol(), school.Config.WebsiteDomainName)
	// Generate API key using HMAC SHA256
	apiKey, err := securityUtil.GenerateHMAC_SHA256_Base64URL(
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
	helpers.Logger.Info("API key generated successfully.", zap.String("API Key", apiKey))
	// Generate app environment variables
	envData := AppEnvData{
		AppName:            school.Name,
		WebsiteTitle:       school.Config.WebsiteTitle,
		WebsiteDescription: school.Config.WebsiteDescription,
		WebsiteURL:         websiteURL,
		ApiBaseUrl:         config.Env.SchoolApiBaseURL,
		CdnUrl:             config.Env.SchoolCdnUrl,
		CdnKey:             config.Env.SchoolCdnKey,
		NextAuthUrl:        websiteURL,
		NextAuthSecret:     securityUtil.GenerateRandomBase64(32),
		SchoolID:           school.ID,
		SchoolType:         school.Type,
		SchoolApiKey:       apiKey,
	}
	// Generate app color data
	colorData := AppColorData{
		Primary:        school.Config.ColorPrimary,
		PrimaryBg:      school.Config.ColorPrimaryBg,
		PrimaryBgHover: school.Config.ColorPrimaryBgHover,
	}
	// Generate deployment status data
	deploymentStatusSchoolApiData := DeploymentStatusSchoolApiKeyData{
		SchoolApiKey: apiKey,
	}
	deploymentStatusApiUrlData := DeploymentStatusApiUrlData{
		ApiUrl: fmt.Sprintf("%s%s/schools/deployment/status", config.Env.ApiBaseURL, config.Env.ApiGroup),
	}
	// Generate deployment data
	kubernetesDeploymentData := KubernetesWebsiteDomainNameData{
		WebsiteDomainName: school.Config.WebsiteDomainName,
	}
	smtpDomainNameData := SmtpDomainNameData{
		SmtpDomainName: school.Config.UserEmailDomainName,
	}
	var selector string = strings.ReplaceAll(strings.ToLower(strings.TrimSpace(school.Config.UserEmailDomainName)), ".", "")
	if len(selector) > 100 {
		selector = fmt.Sprintf("school%dselector", school.ID)
	}
	selector = fmt.Sprintf("%s%d", selector, time.Now().Year())
	smtpSelectorData := SmtpSelectorData{
		SmtpSelector: selector,
	}
	// Define output directory structure
	outputDir := filepath.Join(tempDir, fmt.Sprintf("%d", school.ID))
	deploymentDir := filepath.Join(outputDir, "deployment")
	deploymentStatusDir := filepath.Join(deploymentDir, "status")
	deploymentKubernetesDir := filepath.Join(deploymentDir, "kubernetes")
	deploymentSmtpDir := filepath.Join(deploymentDir, "smtp")
	websiteDir := filepath.Join(outputDir, "website")
	colorDir := filepath.Join(websiteDir, "src", "lib", "constants", "others")
	faviconDir := filepath.Join(websiteDir, "src", "app")
	logosDir := filepath.Join(websiteDir, "public", "assets", "images", "logos")
	// Create required directories
	for _, dir := range []string{outputDir, deploymentDir, deploymentStatusDir, deploymentKubernetesDir, deploymentSmtpDir, websiteDir, colorDir, faviconDir, logosDir} {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			errMsg := "Failed to create directory!"
			err = fmt.Errorf("%s: %s %w", errMsg, dir, err)
			return
		}
	}
	// Generate website files
	if err = htmlHelper.RenderTemplate(filepath.Join(websiteDir, ".env"), envTemplateContent, envData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(websiteDir, ".env"), envTemplateContent, err)
		return
	}
	if err = htmlHelper.RenderTemplate(filepath.Join(colorDir, "color.ts"), colorTemplateContent, colorData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(colorDir, "color.ts"), colorTemplateContent, err)
		return
	}
	// Generate deployment status files
	if err = htmlHelper.RenderTemplate(filepath.Join(deploymentStatusDir, "schoolapikey.txt"), deploymentStatusSchoolApiKeyTemplateContent, deploymentStatusSchoolApiData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(deploymentStatusDir, "schoolapikey.txt"), deploymentStatusSchoolApiKeyTemplateContent, err)
		return
	}
	if err = htmlHelper.RenderTemplate(filepath.Join(deploymentStatusDir, "apiurl.txt"), deploymentStatusApiUrlTemplateContent, deploymentStatusApiUrlData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(deploymentStatusDir, "apiurl.txt"), deploymentStatusApiUrlTemplateContent, err)
		return
	}
	// Generate kubernetes deployment files
	if err = htmlHelper.RenderTemplate(filepath.Join(deploymentKubernetesDir, "domainname.txt"), kubernetesWebsiteDomainNameTemplateContent, kubernetesDeploymentData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(deploymentKubernetesDir, "domainname.txt"), kubernetesWebsiteDomainNameTemplateContent, err)
		return
	}
	// Generate smtp deployment files
	if err = htmlHelper.RenderTemplate(filepath.Join(deploymentSmtpDir, "domainname.txt"), smtpDomainNameDeploymentTemplateContent, smtpDomainNameData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(deploymentSmtpDir, "domainname.txt"), smtpDomainNameDeploymentTemplateContent, err)
		return
	}
	if err = htmlHelper.RenderTemplate(filepath.Join(deploymentSmtpDir, "selector.txt"), smtpSelectorDeploymentTemplateContent, smtpSelectorData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(deploymentSmtpDir, "selector.txt"), smtpSelectorDeploymentTemplateContent, err)
		return
	}
	// Download favicon if available
	if len(school.Favicon) > 0 {
		if err = httpHelper.HttpDownloadFile(school.Favicon, filepath.Join(faviconDir, "favicon.ico")); err != nil {
			helpers.Logger.Error("Failed to download favicon!", zap.String("URL", school.Favicon), zap.String("Error", err.Error()))
		}
	}
	// Download logos if available
	if len(school.Logo) > 0 {
		if err = httpHelper.HttpDownloadFile(school.Logo, filepath.Join(logosDir, "logo.png")); err != nil {
			helpers.Logger.Error("Failed to download logo!", zap.String("URL", school.Logo), zap.String("Error", err.Error()))
		}
	}
	if len(school.LogoWhite) > 0 {
		if err = httpHelper.HttpDownloadFile(school.LogoWhite, filepath.Join(logosDir, "logo-white.png")); err != nil {
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
	err = helpers.GitPushSchoolDeployment(fmt.Sprintf("%d", school.ID), tempDir, outputDir)
	if err != nil {
		helpers.Logger.Error("Failed to push school deployment!", zap.String("Error", err.Error()))
		return
	}
	helpers.Logger.Info(fmt.Sprintf("School deployment successfully added for school %d!", school.ID))
	return
}

// DeleteSchoolDeployment deletes a school deployment from the repository.
func DeleteSchoolDeployment(schoolID int64) (err error) {
	if schoolID < 1 {
		errMsg := "School ID is invalid!"
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Generate API key using HMAC SHA256
	apiKey, err := securityUtil.GenerateHMAC_SHA256_Base64URL(
		fmt.Sprintf("%d", schoolID),
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
	helpers.Logger.Info("API key generated successfully.", zap.String("API Key", apiKey))

	// Create temp directory
	tempDir, err := os.MkdirTemp("", baseDir)
	if err != nil {
		errMsg := "Failed to create temp directory!"
		err = fmt.Errorf("%s: %s %w", errMsg, err.Error(), err)
		return
	}
	defer os.RemoveAll(tempDir)

	// Define output directory structure
	deploymentStatusDir := filepath.Join(tempDir, "status")
	// Create required directories
	for _, dir := range []string{deploymentStatusDir} {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			errMsg := "Failed to create directory!"
			err = fmt.Errorf("%s: %s %w", errMsg, dir, err)
			return
		}
	}
	// Generate deployment status data
	deploymentStatusSchoolApiData := DeploymentStatusSchoolApiKeyData{
		SchoolApiKey: apiKey,
	}
	deploymentStatusApiUrlData := DeploymentStatusApiUrlData{
		ApiUrl: fmt.Sprintf("%s%s/schools/deployment/status", config.Env.ApiBaseURL, config.Env.ApiGroup),
	}
	// Generate deployment status files
	if err = htmlHelper.RenderTemplate(filepath.Join(deploymentStatusDir, "schoolapikey.txt"), deploymentStatusSchoolApiKeyTemplateContent, deploymentStatusSchoolApiData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(deploymentStatusDir, "schoolapikey.txt"), deploymentStatusSchoolApiKeyTemplateContent, err)
		return
	}
	if err = htmlHelper.RenderTemplate(filepath.Join(deploymentStatusDir, "apiurl.txt"), deploymentStatusApiUrlTemplateContent, deploymentStatusApiUrlData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(deploymentStatusDir, "apiurl.txt"), deploymentStatusApiUrlTemplateContent, err)
		return
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

	// Push deployment to delete school
	err = helpers.GitPushDeletedSchoolDeployment(fmt.Sprintf("%d", schoolID), tempDir, deploymentStatusDir)
	if err != nil {
		helpers.Logger.Error("Failed to push deleted school deployment!", zap.String("Error", err.Error()))
		return
	}
	helpers.Logger.Info(fmt.Sprintf("School deployment successfully deleted for school %d!", schoolID))
	return
}
