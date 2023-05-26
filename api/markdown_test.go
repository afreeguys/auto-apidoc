package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMdMethod_GenerateTestDataFromJson1(t *testing.T) {
	m := &MdMethod{
		ParamsJson: "",
	}
	assert.Equal(t, "", m.GenerateTestDataFromJson())

	m.ParamsJson = "{\n\t\"id\":\"int //自增主键id\",\n\t\"accountId\":\"int //广告主id\",\n\t\"deviceId\":\"string //设备id\",\n\t\"convertName\":\"string //转化名称\",\n\t\"convertSource\":\"int //转化来源。1：移动应用，2：网页，3：h5api，4-小程序\",\n\t\"convertPlatform\":\"int //移动平台。1：Android，2：IOS\",\n\t\"convertType\":\"int //转化类型。1：激活，2：注册，3：付费，4：留存，11: 电话拨打，12: 表单提交，13: 按钮点击，14: 页面访问，15: 下载行为，16: 抽奖行为，17: 投票行为，18: 收藏行为，19: 购买行为，20: 分享行为，21: 微信复制，22: 二维码，23: 搜索，24: 咨询，25: 导航跳转，26: 其他\",\n\t\"convertScheme\":\"int //转化方案。1：API，2：SDK，3：JS\",\n\t\"callbackUrl\":\"string //转化数据回调URL\",\n\t\"landingPage\":\"string //api落地页地址\",\n\t\"eventUrl\":\"string //事件url\",\n\t\"status\":\"int //状态：0 正常 1 删除\",\n\t\"statusDebug\":\"int //联调状态：1 未联调 2 联调成功 3 联调失败\",\n\t\"debugFailReason\":\"int //联调失败原因： 1 连接无法访问 2 缺少必要参数\",\n\t\"appPackageName\":\"string //应用包名\",\n\t\"deviceTime\":\"date //设备变更时间\",\n\t\"convertDebugId\":\"long //二维码版本\",\n\t\"miniPlatform\":\"int //小程序平台 1-微信；\",\n\t\"createTime\":\"date //创建时间\",\n\t\"updateTime\":\"date //更新时间\",\n\t\"debugFailReasonNew\":\"string //失败原因汇总\",\n\t\"statusDeliver\":\"int //绑定状态\",\n\t\"miniSourceId\":\"string //小程序原始id\",\n\t\"miniPath\":\"string //小程序路径\"\n}"
	fmt.Println(m.GenerateTestDataFromJson())
}

func TestMdMethod_GenerateTestDataFromJson2(t *testing.T) {
	data := convertTypeToData("1", "int")
	_, ok := data.(int)
	assert.True(t, ok)
	assert.Equal(t, 0, data)

	data = convertTypeToData("1", "string")
	_, ok = data.(string)
	assert.True(t, ok)
	assert.Equal(t, "1_测试数据", data)

	data = convertTypeToData("1", "long")
	_, ok = data.(int64)
	assert.True(t, ok)
	assert.Equal(t, int64(0), data)

	data = convertTypeToData("1", "empty")
	assert.Equal(t, nil, data)

}
