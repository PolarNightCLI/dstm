package localization

import "github.com/nicksnyder/go-i18n/v2/i18n"

var msgAndCountList = map[string]func(int) *i18n.LocalizeConfig{
	"MSG002": func(c int) *i18n.LocalizeConfig {
		return &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "MyUnreadEmails",
				One:   "I have {{.PluralCount}} unread email.",
				Other: "I have {{.PluralCount}} unread emails.",
			},
			PluralCount: c,
		}
	},
}
