package localization

import "github.com/nicksnyder/go-i18n/v2/i18n"

var msgOnlyList = map[string]func() *i18n.LocalizeConfig{
	"_version": func() *i18n.LocalizeConfig {
		return &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "_version",
				Other: "Version",
			},
		}
	},
	"_short_des": func() *i18n.LocalizeConfig {
		return &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "_short_des",
				Other: "Don't Strave Together Dedicated Server Manager For Linux",
			},
		}
	},
	"_long_des": func() *i18n.LocalizeConfig {
		return &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "_long_des",
				Other: "Don't Strave Together Dedicated Server Manager For Linux\n  Github: https://github.com/qaqland/dstm",
			},
		}
	},
}
