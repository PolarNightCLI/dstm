package config

// ref: https://github.com/spf13/viper/issues/258

type Config struct {
	Common commonConf
	Path   pathConf
	Url    urlConf
	Color  colorConf
}

type commonConf struct {
	Lang string
}

type pathConf struct {
	ConfDir      string `mapstructure:"conf_dir"`
	DSTRootDir   string `mapstructure:"dst_root_dir"`
	UGCDir       string `mapstructure:"ugc_dir"`
	V1ModDir     string `mapstructure:"v1_mod_dir"`
	V2ModDir     string `mapstructure:"v2_mod_dir"`
	KleiRootDir  string `mapstructure:"klei_root_dir"`
	WorldDirName string `mapstructure:"world_dir_name"`
}

type urlConf struct {
	GithubURL string `mapstructure:"github_url"`
	GiteeURL  string `mapstructure:"gitee_url"`
}

type colorConf struct {
	MainColor string `mapstructure:"main_color"`
}
