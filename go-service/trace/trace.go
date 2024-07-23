package trace

import "github.com/devkhatri523/ecom-go/config/config"

func GetTraceKey() string {
	return defaultString(config.Default().GetString("logger.traceKey"), "debugId")
}

func defaultString(str string, defaultStr string) string {
	if str == "" {
		return defaultStr
	} else {
		return str
	}
}
