package common

type ServiceConfiguration struct {
	SERVICE_PORT                string `mapstructure:"SERVICE_PORT"`
	SERVICE_NAME                string `mapstructure:"SERVICE_NAME"`
	SERVICE_DOMAIN              string `mapstructure:"SERVICE_DOMAIN"`
	SERVICE_BASEPATH            string `mapstructure:"SERVICE_BASEPATH"`
	SERVICE_VERSION             string `mapstructure:"SERVICE_VERSION"`
	COURSE_COLLECTION 			string `mapstructure:"COURSE_COLLECTION"`
	CHAPTER_COLLECTION 			string `mapstructure:"CHAPTER_COLLECTION"`
	MONGO_DATABASE              string `mapstructure:"MONGO_DATABASE"`
	SERVICE_DESCRIPTION         string `mapstructure:"SERVICE_DESCRIPTION"`
	SERVICE_HOST                string `mapstructure:"SERVICE_HOST"`
	JAEGER_URL                  string `mapstructure:"JAEGER_URL"`
	TRACER_ENABLED              string `mapstructure:"TRACER_ENABLED"`
	GCP_PROJECT_ID              string `mapstructure:"GCP_PROJECT_ID"`
	GCP_PRIVATE_KEY             string `mapstructure:"GCP_PRIVATE_KEY"`
	GCP_CLIENT_EMAIL            string `mapstructure:"GCP_CLIENT_EMAIL"`
	GCP_AUTHORIZATION_TYPE      string `mapstructure:"GCP_AUTHORIZATION_TYPE"`
	PUBSUB_SUBSCRIPTION         string `mapstructure:"PUBSUB_SUBSCRIPTION"`
	PUBSUB_TOPIC                string `mapstructure:"PUBSUB_TOPIC"`
	GATEWAY_PUBLIC_KEY          string `mapstructure:"GATEWAY_PUBLIC_KEY"`
	GATEWAY_PRIVATE_KEY         string `mapstructure:"GATEWAY_PRIVATE_KEY"`
	AUDIT_SERVICE               string `mapstructure:"AUDIT_SERVICE"`
	FANOUT_SERVICE              string `mapstructure:"FANOUT_SERVICE"`
	REDIS_HOST                  string `mapstructure:"REDIS_HOST"`
	REDIS_PORT                  string `mapstructure:"REDIS_PORT"`
	REDIS_NAME                  string `mapstructure:"REDIS_NAME"`
	REDIS_PASSWORD              string `mapstructure:"REDIS_PASSWORD"`
	REDIS_TOKEN_COLLECTION      string `mapstructure:"REDIS_TOKEN_COLLECTION"`
	REDIS_AUTHMECHANISAM        string `mapstructure:"REDIS_AUTHMECHANISAM"`
	LOOKUP_SERVICE              string `mapstructure:"LOOKUP_SERVICE"`
	DEFAULT_SENDER_SUBCRIBER_ID string `mapstructure:"DEFAULT_SENDER_SUBCRIBER_ID"`
	DEFAULT_SIGNATURE           string `mapstructure:"DEFAULT_SIGNATURE"`
	DEFAULT_REQUEST_ID          string `mapstructure:"DEFAULT_REQUEST_ID"`
	VLOOKUP_URL                 string `mapstructure:"VLOOKUP_URL"`
	DEFAULT_COUNTRY             string `mapstructure:"DEFAULT_COUNTRY"`
	VLOOKUP_SIGNATURE_URL       string `mapstructure:"VLOOKUP_SIGNATURE_URL"`
	VLOOKUP_PRIVATE_KEY         string `mapstructure:"VLOOKUP_PRIVATE_KEY"`
}