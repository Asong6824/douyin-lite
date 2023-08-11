package setting

import(
    "time"
    
)

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
	UploadVideoMaxSize   int
	UploadVideoAllowExts []string
}

type MysqlSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type RedisSettingS struct {
	Host          string
  	Port          string
  	Password      string
  	DB            int
  	MaxIdleConns  int
  	axActiveConns int
  	IdleTimeout   time.Duration
}

type JWTSettingS struct {
	Secret string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}