package localization

import "github.com/nicksnyder/go-i18n/v2/i18n"

var msgAndDataList = map[string]func(map[string]string) *i18n.LocalizeConfig{
	"MSG003": func(d map[string]string) *i18n.LocalizeConfig {
		return &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "HelloPerson",
				Other: "Hello {{.Name}}",
			},
			TemplateData: d,
		}
	},
}
