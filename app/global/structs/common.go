package structs

// EnvConfig dev.yaml格式
type EnvConfig struct {
	DBMaster    DbMaster    `yaml:"master"`
	DbSlave     DbSlave     `yaml:"slave"`
	API         API         `yaml:"api"`
	ImagePath   ImagePath   `yaml:"img_path"`
	ImageServer ImageServer `yaml:"img_server"`
	Log         Log         `yaml:"log"`
	DB          DB          `yaml:"db"`
	Redis       Redis       `yaml:"redis"`
}

// DbMaster 載入db的master環境設定
type DbMaster struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// DbSlave 載入db的slave環境設定
type DbSlave struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// API 載入各單位api環境設定
type API struct {
	CypressURL   string `yaml:"cypress_url"`
	CypressToken string `yaml:"cypress_token"`
	RD1URL       string `yaml:"rd1_url"`
}

// ImagePath 載入各單位other環境設定
type ImagePath struct {
	ImgPathRotate string `yaml:"img_path_rotate"`
}

// ImageServer 載入各單位other環境設定
type ImageServer struct {
	ImgServerRotate string `yaml:"img_server_rotate"`
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
