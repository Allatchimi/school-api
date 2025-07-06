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
	baseDir = ".deployment-"
)

// DeploySchool generates configuration files and pushes them to the repository.
func DeploySchool(school *model.School) (ok bool, err error) {
	if school == nil || school.ID < 1 || school.Config == nil || len(school.Config.DomainName) < 1 {
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
	outputDir := filepath.Join(tempDir, fmt.Sprintf("%d", school.ID))
	deploymentDir := filepath.Join(outputDir, "deployment")
	websiteDir := filepath.Join(outputDir, "website")
	colorDir := filepath.Join(websiteDir, "src", "lib", "api", "constants", "common")
	faviconDir := filepath.Join(websiteDir, "src", "app")
	logosDir := filepath.Join(websiteDir, "public", "images", "logos")
	// Create required directories
	for _, dir := range []string{outputDir, deploymentDir, websiteDir, colorDir, faviconDir, logosDir} {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			errMsg := "Failed to create directory!"
			err = fmt.Errorf("%s: %s %w", errMsg, dir, err)
			return
		}
	}
	// Generate website files
	if err = helpers.RenderTemplate(filepath.Join(websiteDir, ".env"), envTemplateContent, envData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(websiteDir, ".env"), envTemplateContent, err)
		return
	}
	if err = helpers.RenderTemplate(filepath.Join(colorDir, "color.ts"), colorTemplateContent, colorData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(colorDir, "color.ts"), colorTemplateContent, err)
		return
	}
	// Generate deployment files
	if err = helpers.RenderTemplate(filepath.Join(deploymentDir, "domainname.txt"), domainNameDeploymentTemplateContent, deploymentData); err != nil {
		errMsg := "Failed to render template!"
		err = fmt.Errorf("%s: %s %s %w", errMsg, filepath.Join(deploymentDir, "domainname.txt"), domainNameDeploymentTemplateContent, err)
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
	tempDir, err := os.MkdirTemp("", baseDir)
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
