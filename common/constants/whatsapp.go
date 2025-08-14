package constants

const (
	// 2 params: {{examName subject}}, {{examDate}}, {{schoolName}}
	WHATSAPP_TEMPLATE_COURSE_PUBLISHED        = "course_published"
	WHATSAPP_TEMPLATE_COURSE_PUBLISHED_PARENT = "course_published_parent"

	// 3 params: {{examName subject}}, {{examDate}}, {{schoolName}}
	WHATSAPP_TEMPLATE_EXAM_PUBLISHED = "exam_published"
	// 4 params: {{examName subject}}, {{childName}}, {{examDate}}, {{schoolName}}
	WHATSAPP_TEMPLATE_EXAM_PUBLISHED_PARENT = "exam_published_parent"

	// 2 params: {{subject}}, {{schoolName}}
	WHATSAPP_TEMPLATE_RESULT_PUBLISHED = "result_published"
	// 2 params: {{subject (className)}}, {{schoolName}}
	WHATSAPP_TEMPLATE_RESULT_PUBLISHED_PARENT = "result_published_parent"

	// 2 params: {{periodName}}, {{schoolName}}
	WHATSAPP_TEMPLATE_REPORT_PUBLISHED = "report_published"
	// 2 params: {{periodName (className)}}, {{schoolName}}
	WHATSAPP_TEMPLATE_REPORT_PUBLISHED_PARENT = "report_published_parent"

	// 3 params: {{requestTitle}}, {{status}}, {{schoolName}}
	WHATSAPP_TEMPLATE_REQUEST_STATUS = "request_status"

	// 1 params: {{schoolName}}
	WHATSAPP_TEMPLATE_WELCOME = "welcome"
)
