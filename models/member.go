package models

import "time"

// 会员
type Member struct {
	/** 主键 */
	MemberID int `gorm:"column:member_id;primaryKey" json:"member_id"`
	/** 用户 nickname */
	MemberName string `gorm:"column:member_name" json:"member_name"`
	/** 组id */
	GroupID int `gorm:"column:group_id" json:"group_id"`
	/** 密码 */
	Password string `gorm:"column:password" json:"password"`
	/** 密码盐 */
	Salt string `gorm:"column:salt" json:"salt"`
	/** verify 免验证 */
	Verify int64 `gorm:"column:verify" json:"verify"`
	/** 微信号 */
	Wechat string `gorm:"column:wechat" json:"wechat"`
	/** 头像 */
	Avatar string `gorm:"column:avatar" json:"avatar"`
	/** 用户心情 */
	Heart string `gorm:"column:heart" json:"heart"`
	/** 邮箱地址 */
	Email string `gorm:"column:email" json:"email"`
	/** 区号 */
	AreaCode int `gorm:"column:area_code" json:"area_code"`
	/** 手机号 */
	Mobile string `gorm:"column:mobile" json:"mobile"`
	/** 真实姓名 */
	Name string `gorm:"column:name" json:"name"`
	/** 交易密码 */
	Paypass string `gorm:"column:paypass" json:"paypass"`
	/** 实名认证 */
	Prove int `gorm:"column:prove" json:"prove"`
	/** 认证时间 */
	ProveTime *time.Time `gorm:"column:prove_time" json:"prove_time"`
	/** 身份证号 */
	Idcard string `gorm:"column:idcard" json:"idcard"`
	/** 家庭住址 */
	Address string `gorm:"column:address" json:"address"`
	/** 邮政编码 */
	Zipcode string `gorm:"column:zipcode" json:"zipcode"`
	/** 修改次数 */
	ModifyName int `gorm:"column:modify_name" json:"modify_name"`
	/** 积分 */
	Score int `gorm:"column:score" json:"score"`
	/** 钱包 */
	Money float64 `gorm:"column:money" json:"money"`
	/** 密保问题 */
	Question string `gorm:"column:question" json:"question"`
	/** 答案 */
	Answer string `gorm:"column:answer" json:"answer"`
	/** 性别 */
	Gender int `gorm:"column:gender" json:"gender"`
	/** 状态 */
	Status int `gorm:"column:status" json:"status"`
	/** 用户appid */
	AppID int `gorm:"column:app_id" json:"app_id"`
	/** 用户appsecret */
	AppSecret string `gorm:"column:app_secret" json:"app_secret"`
	/** 点击量 */
	Hits int `gorm:"column:hits" json:"hits"`
	/** 日点击 */
	HitsDay int `gorm:"column:hits_day" json:"hits_day"`
	/** 周点击 */
	HitsWeek int `gorm:"column:hits_week" json:"hits_week"`
	/** 月点击 */
	HitsMonth int `gorm:"column:hits_month" json:"hits_month"`
	/** 点击时间 */
	HitsLasttime int `gorm:"column:hits_lasttime" json:"hits_lasttime"`
	/** 激活码 */
	Valicode string `gorm:"column:valicode" json:"valicode"`
	/** 邀请人 */
	InviteID int `gorm:"column:invite_id" json:"invite_id"`
	/** 邀请人父级 */
	InviteVid int `gorm:"column:invite_vid" json:"invite_vid"`
	/** 顶级邀请人 */
	InviteTid int `gorm:"column:invite_tid" json:"invite_tid"`
	/** 邀请层级 */
	InviteLevel int `gorm:"column:invite_level" json:"invite_level"`
	/** 邀请关系 */
	InviteList string `gorm:"column:invite_list" json:"invite_list"`
	/** 注册渠道 */
	Channel string `gorm:"column:channel" json:"channel"`
	/** lv1认证 */
	Lv1Status int `gorm:"column:lv1_status" json:"lv1_status"`
	/** lv2认证 */
	Lv2Status int `gorm:"column:lv2_status" json:"lv2_status"`
	/** 登录ip */
	LoginIp string `gorm:"column:login_ip" json:"login_ip"`
	/** 登录时间 */
	LoginTime int `gorm:"column:login_time" json:"login_time"`
	/** 登录次数 */
	LoginCount int `gorm:"column:login_count" json:"login_count"`
	/** 用户地址 */
	Url string `gorm:"column:url" json:"url"`
	/** photonpay持卡人id */
	HolderID string `gorm:"column:holder_id" json:"holder_id"`
	/** 注册IP */
	CreateIp string `gorm:"column:create_ip" json:"create_ip"`
	/** 自动提现 */
	AutoWithdraw int `gorm:"column:auto_withdraw" json:"auto_withdraw"`
	/** 默认可开卡数量 */
	PercentMax int `gorm:"column:percent_max" json:"percent_max"`
	/** 注册时间 */
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time"`
	/** 修改时间 */
	UpdateTime *time.Time `gorm:"column:update_time" json:"update_time"`
	/** 软删除标识 */
	DeleteTime *time.Time `gorm:"column:delete_time" json:"delete_time"`
	/** 旧ID */
	MemberOld int `gorm:"column:member_old" json:"member_old"`
}

func (Member) TableName() string {
	return "db_member"
}
