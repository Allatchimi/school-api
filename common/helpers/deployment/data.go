package deploymentHelper

const (
	// App env templates
	envTemplateContent = `
NEXT_PUBLIC_APP_NAME="{{ .AppName }}"
NEXT_PUBLIC_WEBSITE_TITLE="{{ .WebsiteTitle }}"
NEXT_PUBLIC_WEBSITE_DESCRIPTION="{{ .WebsiteDescription }}"
NEXT_PUBLIC_WEBSITE_URL="{{ .WebsiteURL }}"

API_BASE_URL="{{ .ApiBaseUrl }}"
CDN_URL="{{ .CdnUrl }}"
CDN_KEY="{{ .CdnKey }}"

NEXT_AUTH_URL="{{ .NextAuthUrl }}"
NEXT_AUTH_SECRET="{{ .NextAuthSecret }}"

SCHOOL_ID={{ .SchoolID }}
SCHOOL_API_KEY="{{ .SchoolApiKey }}"
`

	// App colors templates
	colorTemplateContent = `
interface ColorScheme {
  label?: string;
  primary: string;
  primaryBg: string;
  primaryBgHover: string;
}

export const COLOR_SCHEME_BLACK: ColorScheme = {
  label: "Black 1",
  primary: "#3a3a3a",
  primaryBg: "#e7e7e7",
  primaryBgHover: "#bbbbbb",
};

export const COLOR_SCHEME_BLUE: ColorScheme = {
  label: "Blue 1",
  primary: "#3c6989",
  primaryBg: "#e1e8ec",
  primaryBgHover: "#a3bfd3",
};

export const COLOR_SCHEME_GREEN: ColorScheme = {
  label: "Green 1",
  primary: "#3c8971",
  primaryBg: "#e9f0ed",
  primaryBgHover: "#a3c6b9",
};

export const COLOR_SCHEME_ORANGE: ColorScheme = {
  label: "Orange",
  primary: "#895f3c",
  primaryBg: "#ebe5e1",
  primaryBgHover: "#c6b4a3",
};

export const COLOR_SCHEME_PINK: ColorScheme = {
  label: "Pink",
  primary: "#b63bac",
  primaryBg: "#ebe1e8",
  primaryBgHover: "#c6a3c4",
};

export const COLOR_SCHEME_RED: ColorScheme = {
  label: "Red",
  primary: "#b63b3b",
  primaryBg: "#ebe1e1",
  primaryBgHover: "#c6a3a3",
};

export const COLOR_SCHEME_YELLOW: ColorScheme = {
  label: "Yellow",
  primary: "#b6b43b",
  primaryBg: "#ebebe1",
  primaryBgHover: "#c5c6a3",
};

export const COLOR_SCHEME_DEFAULT: ColorScheme = {
  label: "Default",
  primary: "{{ .Primary }}",
  primaryBg: "{{ .PrimaryBg }}",
  primaryBgHover: "{{ .PrimaryBgHover }}",
};

export const COLOR_SCHEME = COLOR_SCHEME_DEFAULT;

export const COLOR_SCHEMES = [
  COLOR_SCHEME_BLACK,
  COLOR_SCHEME_BLUE,
  COLOR_SCHEME_GREEN,
  COLOR_SCHEME_ORANGE,
  COLOR_SCHEME_PINK,
  COLOR_SCHEME_RED,
  COLOR_SCHEME_YELLOW,
];
`

	// Deployment status templates
	deploymentStatusSchoolApiKeyTemplateContent = `{{ .SchoolApiKey }}`
	deploymentStatusApiUrlTemplateContent       = `{{ .ApiUrl }}`

	// Kubernetes templates
	kubernetesWebsiteDomainNameTemplateContent = `{{ .WebsiteDomainName }}`

	// SMTP templates
	smtpDomainNameDeploymentTemplateContent = `{{ .SmtpDomainName }}`
	smtpSelectorDeploymentTemplateContent   = `{{ .SmtpSelector }}`
)

type AppEnvData struct {
	AppName            string
	WebsiteTitle       string
	WebsiteDescription string
	WebsiteURL         string

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

type DeploymentStatusSchoolApiKeyData struct {
	SchoolApiKey string
}

type DeploymentStatusApiUrlData struct {
	ApiUrl string
}
