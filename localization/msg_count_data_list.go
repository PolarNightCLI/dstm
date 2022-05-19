package localization

import "github.com/nicksnyder/go-i18n/v2/i18n"

var msgCountDataList = map[string]func(int, map[string]string) *i18n.LocalizeConfig{
	"MSG004": func(c int, d map[string]string) *i18n.LocalizeConfig {
		return &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "PersonUnreadEmails",
				One:   "{{.Name}} has {{.UnreadEmailCount}} unread email.",
				Other: "{{.Name}} has {{.UnreadEmailCount}} unread emails.",
			},
			PluralCount:  c,
			TemplateData: d,
		}
	},
}
