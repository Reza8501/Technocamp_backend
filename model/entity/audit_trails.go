package entity

type AuditTrail struct {
	User        string `json:"user" gorm:"column:user"`
	IpAddress   string `json:"ip_address" gorm:"column:ip_address"`
	Service     string `json:"service" gorm:"column:service"`
	Method      string `json:"method" gorm:"column:method"`
	Endpoint    string `json:"endpoint" gorm:"column:endpoint"`
	Status      string `json:"status" gorm:"column:status"`
	RequestBody string `json:"request_body" gorm:"column:request_body"`
	InquiryKey  string `json:"inquiry_key" gorm:"column:inquiry_key"`
	Remark      string `json:"remark" gorm:"column:remark"`
	Time        string `json:"time" gorm:"column:time"`
}

type AuditTrail_v2 struct {
	User        string `json:"user" gorm:"column:user"`
	IpAddress   string `json:"ip_address" gorm:"column:ip_address"`
	Service     string `json:"service" gorm:"column:service"`
	Method      string `json:"method" gorm:"column:method"`
	Endpoint    string `json:"endpoint" gorm:"column:endpoint"`
	HttpStatus      string `json:"http_status" gorm:"column:http_status"`
	RespondStatus      string `json:"respond_status" gorm:"column:respond_status"`
	RequestBody string `json:"request_body" gorm:"column:request_body"`
	RespondBody string `json:"respond_body" gorm:"column:respond_body"`
	InquiryKey  string `json:"inquiry_key" gorm:"column:inquiry_key"`
	Remark      string `json:"remark" gorm:"column:remark"`
	Time        string `json:"time" gorm:"column:time"`
}
