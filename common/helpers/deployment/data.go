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
interface ColorScheme {
  primary: string;
  primaryBg: string;
  primaryBgHover: string;
}

export const COLOR_SCHEME_DEFAULT: ColorScheme = {
  primary: "{{ .Primary }}",
  primaryBg: "{{ .PrimaryBg }}",
  primaryBgHover: "{{ .PrimaryBgHover }}",
};

export const COLOR_SCHEME = COLOR_SCHEME_DEFAULT;

export const COLOR_SCHEMES = [
  COLOR_SCHEME_DEFAULT,
];
`

	// Kubernetes templates
	kubernetesWebsiteDomainNameTemplateContent = `{{ .WebsiteDomainName }}`

	// SMTP templates
	smtpDomainNameDeploymentTemplateContent = `{{ .SmtpDomainName }}`
	smtpSelectorDeploymentTemplateContent   = `{{ .SmtpSelector }}`
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

type KubernetesWebsiteDomainNameData struct {
	WebsiteDomainName string
}

type SmtpDomainNameData struct {
	SmtpDomainName string
}

type SmtpSelectorData struct {
	SmtpSelector string
}
