/**
 * Copyright 2015 @ 56x.net.
 * name : profilemanager.go
 * author : jarryliu
 * date : 2016-05-26 21:19
 * description :
 * history :
 */
package merchant

type (

	// Authenticate 商户认证信息
	Authenticate struct {
		// Id
		Id int `db:"id" pk:"yes" auto:"yes" json:"id" bson:"id"`
		// 商户编号
		MchId int `db:"mch_id" json:"mchId" bson:"mchId"`
		// 公司名称
		OrgName string `db:"org_name" json:"orgName" bson:"orgName"`
		// 营业执照编号
		OrgNo string `db:"org_no" json:"orgNo" bson:"orgNo"`
		// 营业执照照片
		OrgPic string `db:"org_pic" json:"orgPic" bson:"orgPic"`
		// 办公地
		WorkCity int `db:"work_city" json:"workCity" bson:"workCity"`
		// 资质图片
		QualificationPic string `db:"qualification_pic" json:"qualificationPic" bson:"qualificationPic"`
		// 法人身份证号
		PersonId string `db:"person_id" json:"personId" bson:"personId"`
		// 法人姓名
		PersonName string `db:"person_name" json:"personName" bson:"personName"`
		// 法人身份证照片
		PersonPic string `db:"person_pic" json:"personPic" bson:"personPic"`
		// 联系人手机
		PersonPhone string `db:"person_phone" json:"personPhone" bson:"personPhone"`
		// 授权书
		AuthorityPic string `db:"authority_pic" json:"authorityPic" bson:"authorityPic"`
		// 开户银行
		BankName string `db:"bank_name" json:"bankName" bson:"bankName"`
		// 开户户名
		BankAccount string `db:"bank_account" json:"bankAccount" bson:"bankAccount"`
		// 开户账号
		BankNo string `db:"bank_no" json:"bankNo" bson:"bankNo"`
		// 扩展数据
		ExtraData string `db:"extra_data" json:"extraData" bson:"extraData"`
		// 审核状态
		ReviewStatus int `db:"review_status" json:"reviewStatus" bson:"reviewStatus"`
		// 审核备注
		ReviewRemark string `db:"review_remark" json:"reviewRemark" bson:"reviewRemark"`
		// 审核时间
		ReviewTime int `db:"review_time" json:"reviewTime" bson:"reviewTime"`
		// 版本号: 0: 待审核 1: 已审核
		Version int `db:"version" json:"reviewTime" bson:"reviewTime"`
		// 更新时间
		UpdateTime int `db:"update_time" json:"updateTime" bson:"updateTime"`
	}

	// 企业信息
	EnterpriseInfo struct {
		// 编号
		ID int32 `db:"id" pk:"yes" auto:"yes"`
		// 商户编号
		MchId int64 `db:"mch_id"`
		// 公司名称
		CompanyName string `db:"company_name"`
		// 公司营业执照编号
		CompanyNo string `db:"company_no"`
		// 法人
		PersonName string `db:"person_name"`
		// 法人身份证编号
		PersonIdNo string `db:"person_id"`
		// 身份证验证图片(人捧身份证照相)
		PersonImage string `db:"person_image"`
		// 公司电话
		Tel string `db:"tel"`
		// 省
		Province int32 `db:"province"`
		// 市
		City int32 `db:"city"`
		// 区
		District int32 `db:"district"`
		// 省+市+区字符串表示
		Location string `db:"location"`
		// 公司地址
		Address string `db:"address"`
		// 营业执照图片
		CompanyImage string `db:"company_image"`
		// 授权书
		AuthDoc string `db:"auth_doc"`
		//是否已审核
		Reviewed int32 `db:"review_status"`
		// 审核时间
		ReviewTime int64 `db:"review_time"`
		// 审核备注
		ReviewRemark string `db:"review_remark"`
		//更新时间
		UpdateTime int64 `db:"update_time"`
	}

	// 基本资料管理器
	IProfileManager interface {
		// SaveAuthenticate 保存商户认证信息
		SaveAuthenticate(v *Authenticate) (int, error)

		// 标记企业为审核通过
		ReviewAuthenticate(reviewed bool, message string) error

		// 修改密码
		ChangePassword(newPassword, oldPwd string) error
	}
)
