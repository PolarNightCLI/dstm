package dstini

type ClusterConfig struct {
	Network  cNetwork
	Gameplay gameplay
	Misc     misc
	Shard    cShard
	Steam    cSteam
}

type cNetwork struct {
	Name        string `ini:"cluster_name"`
	Description string `ini:"cluster_description"`
	Password    string `ini:"cluster_password"`
	Lang        string `ini:"cluster_language"`
	Intention   string `ini:"cluster_intention"`

	//OfflineCluster   bool `ini:"offline_cluster"`
	//TickRate         int  `ini:"tick_rate"`
	//WhitelistSlots   int  `ini:"whitelist_slots"`
	//LanOnlyCluster   bool `ini:"lan_only_cluster"`
	//AutosaverEnabled bool `ini:"autosaver_enabled"`
}

type gameplay struct {
	GameMode       string `ini:"game_mode"`
	MaxPlayers     uint   `ini:"max_players"`
	PVP            bool   `ini:"pvp"`
	PauseWhenEmpty bool   `ini:"pause_when_empty"`
	VoteEnabled    bool   `ini:"vote_enabled"`

	//VoteKickEnabled bool `ini:"vote_kick_enabled"`
}

type misc struct {
	ConsoleEnabled bool `ini:"console_enabled"`
	MaxSnapshots   uint `ini:"max_snapshots"`
}

type cShard struct {
	ShardEnabled bool   `ini:"shard_enabled"`
	BindIP       string `ini:"bind_ip"`
	MasterIP     string `ini:"master_ip"`
	MasterPort   uint   `ini:"master_port"`
	ClusterKey   string `ini:"cluster_key"`
}

type cSteam struct {
	SteamGroupOnly bool `ini:"steam_group_only"`

	//SteamGroupID     int  `ini:"steam_group_id"`
	//SteamGroupAdmins bool `ini:"steam_group_admins"`
}
