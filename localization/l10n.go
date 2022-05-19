package localization

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var Locale = language.English

type localization struct {
	bundle *i18n.Bundle
}

func NewLocalizer() localization {
	// en, zh, ja
	b := i18n.NewBundle(language.English)
	b.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	b.MustLoadMessageFile("./localization/active.zh.toml")
	b.MustLoadMessageFile("./localization/active.ja.toml")
	return localization{
		bundle: b,
	}
}

const (
	MsgOnly      = 0
	MsgAndCount  = 1
	MsgAndData   = 2
	MsgCountData = 3
)

func (l localization) String(key string, confType int, count int, data map[string]string) string {
	localizer := i18n.NewLocalizer(l.bundle, Locale.String())
	var refer *i18n.LocalizeConfig
	switch confType {
	case MsgOnly:
		f, ok := msgOnlyList[key]
		if !ok {
			return key
		}
		refer = f()
	case MsgAndCount:
		f, ok := msgAndCountList[key]
		if !ok {
			return key
		}
		refer = f(count)
	case MsgAndData:
		f, ok := msgAndDataList[key]
		if !ok {
			return key
		}
		refer = f(data)
	case MsgCountData:
		f, ok := msgCountDataList[key]
		if !ok {
			return key
		}
		refer = f(count, data)
	default:
		panic("Error: wrong config type in localization.String()")
	}
	return localizer.MustLocalize(refer)
}
