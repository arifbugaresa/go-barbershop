package constant

type Dialect string

func (d Dialect) Dialect() string {
	return string(d)
}

const (
	Postgres Dialect = "postgres"
	Mysql    Dialect = "mysql"
)

type TableName string

func (t TableName) TableName() string {
	return string(t)
}

const (
	User                   TableName = "users"
	Role                   TableName = "roles"
	Permission             TableName = "permissions"
	RolePermission         TableName = "role_permissions"
	EmailTemplateTableName TableName = "email_templates"
	Service                TableName = "services"
)

type DateTimeFormat string

func (d DateTimeFormat) String() string {
	return string(d)
}

const (
	DateFormat DateTimeFormat = "2006-02-01"
)

type RegexFormat string

func (d RegexFormat) String() string {
	return string(d)
}

type EmailTemplate string

const (
	SignUpEmailTemplate EmailTemplate = "sign_up_email_template"
	LoginEmailTemplate  EmailTemplate = "login_email_template"
)

func (d EmailTemplate) String() string {
	return string(d)
}

type MqQueue string

func (r MqQueue) String() string {
	return string(r)
}

const (
	DefaultQueue MqQueue = "default"
	EmailQueue   MqQueue = "email"
)

type AppMode string

func (d AppMode) String() string {
	return string(d)
}

const (
	DevelopmentMode AppMode = "development"
	StagingMode     AppMode = "staging"
	ProductionMode  AppMode = "production"
)
