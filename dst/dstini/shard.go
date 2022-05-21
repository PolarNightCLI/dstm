package dstini

type ShardConfig struct {
	Network sNetwork
	Shard   sShard
	Account account
	Steam   sSteam
}

type sNetwork struct {
	ServerPort uint `ini:"server_port"`
}

type sShard struct {
	IsMaster bool   `ini:"is_master"`
	Name     string `ini:"name"`
	Id       uint   `ini:"id"`
}

type account struct {
	EncodeUserPath bool `ini:"encode_user_path"`
}

type sSteam struct {
	MasterServerPort   uint `ini:"master_server_port"`
	AuthenticationPort uint `ini:"authentication_port"`
}
