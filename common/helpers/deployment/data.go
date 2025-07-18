package deploymentHelper

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

SCHOOL_ID={{ .SchoolID }}
SCHOOL_TYPE="{{ .SchoolType }}"
SCHOOL_API_KEY="{{ .SchoolApiKey }}"
`

	// App colors templates
	colorTemplateContent = `
export const COLOR_PRIMARY = "{{ .Primary }}";
export const COLOR_PRIMARY_BG = "{{ .PrimaryBg }}";
export const COLOR_PRIMARY_BG_HOVER = "{{ .PrimaryBgHover }}";

export const COLORS = [COLOR_PRIMARY, COLOR_PRIMARY_BG, COLOR_PRIMARY_BG_HOVER];
`

	// Deployment templates
	domainNameDeploymentTemplateContent = `{{ .DomainName }}`
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
}
