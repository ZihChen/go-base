package structs

// EnvConfig dev.yaml格式
type EnvConfig struct {
	DBMaster    DBMaster    `yaml:"master"`
	DBSlave     DBSlave     `yaml:"slave"`
	API         API         `yaml:"api"`
	Log         Log         `yaml:"log"`
	DB          DB          `yaml:"db"`
	Redis       Redis       `yaml:"redis"`
	GrpcSetting GrpcSetting `yaml:"grpc_setting"`
}

// DBMaster 載入db的master環境設定
type DBMaster struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// DBSlave 載入db的slave環境設定
type DBSlave struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// API 載入各單位api環境設定
type API struct {
	PitayaGrpcServer string `yaml:"pitaya_grpc_server"`
	LemonGrpcServer  string `yaml:"lemon_grpc_server"`
}

// Log 載入Log設定檔規則
type Log struct {
	LogDir    string `yaml:"log_dir"`
	AccessLog string `yaml:"access_log"`
	ErrorLog  string `yaml:"error_log"`
}

// DB 對DB其他操作的設定
type DB struct {
	Debug bool `yaml:"debug"`
}

// Redis 載入redis設定
type Redis struct {
	RedisHost string `yaml:"redis_host"`
	RedisPort string `yaml:"redis_port"`
	RedisPwd  string `yaml:"redis_password"`
}

// APIResult 回傳API格式
type APIResult struct {
	ErrorCode   int         `json:"error_code"`
	ErrorMsg    string      `json:"error_msg"`
	LogIDentity string      `json:"log_id"`
	Result      interface{} `json:"result"`
}

// GrpcSetting grpc 自動註冊服務設定
type GrpcSetting struct {
	Code string `yaml:"code"`
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}
