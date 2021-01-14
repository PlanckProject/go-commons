package mongo

type Config struct {
	Hosts            []string `mapstructure:"hosts"`
	ReplicaSet       string   `mapstructure:"replica_set"`
	Username         string   `mapstructure:"username"`
	Password         string   `mapstructure:"password"`
	SSL              bool     `mapstructure:"ssl"`
	RetryWrites      bool     `mapstructure:"retry_writes"`
	ConnectionString string   `mapstructure:"connection_string"`
	Database         string   `mapstructure:"database"`
	Collection       string   `mapstructure:"collection"`
}
