package appconfig

import "os"

var PORT = "8009"
var STATIC_ROUTE = "/public"
var STATIC_DIR = "./public"
var SECRET_KEY = "SECRET_KEY"

func InitAppConfig() {
	portEnv := os.Getenv("APP_PORT")
	if portEnv != "" {
		PORT = portEnv
	}

	staticRouteEnv := os.Getenv("STATIC_ROUTE")
	if staticRouteEnv != "" {
		STATIC_ROUTE = staticRouteEnv
	}

	staticDirEnv := os.Getenv("STATIC_DIR")
	if staticDirEnv != "" {
		STATIC_DIR = staticDirEnv
	}

	secretKeyEnv := os.Getenv("SECRET_KEY")
	if secretKeyEnv != "" {
		SECRET_KEY = secretKeyEnv
	}
}