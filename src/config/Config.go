package config

type Config struct {
	Mysql    Mysql       `mapstructure:"mysql"`
	Port     int         `mapstructure:"port"`
	Name     string      `mapstructure:"name"`
	LogsPath string      `mapstructure:"logsPath"`
	Minio    MinioConfig `mapstructure:"minio"`
}

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

type MinioConfig struct {
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyID     string `mapstructure:"accessKeyID"`
	SecretAccessKey string `mapstructure:"secretAccessKey"`
	BucketName      string `mapstructure:"bucketName"`
	Path            string `mapstructure:"path"`
}

type EnvConfig struct {
	Env string `mapstructure:"env"`
}
