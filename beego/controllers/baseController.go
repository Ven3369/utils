package controllers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

//@作者：Ven
//@类型详情：基础控制器类型用于“继承”
//@备注信息：无
//@时间：2018-07-08
type BaseController struct {
	beego.Controller
}

func init() {
}

//@作者：Ven
//@函数详情：获取客户端方法
//@备注信息：无
//@参数：无
//@返回值：string
//@时间：2018-07-08
func (that *BaseController) GetMethod() string {
	return that.Ctx.Input.Method()
}

//@作者：Ven
//@函数详情：获取客户端ip
//@备注信息：无
//@参数：无
//@返回值：string
//@时间：2018-07-08
func (that *BaseController) QueryToInt64(key string) (int64, error) {
	k := that.Ctx.Input.Query(key)
	if k == "" {
		return 0, fmt.Errorf("not_key")
	}

	return strconv.ParseInt(k, 10, 64)

}

//@作者：Ven
//@函数详情：获取客户端ip
//@备注信息：无
//@参数：无
//@返回值：string
//@时间：2018-07-08
func (that *BaseController) GetIp() string {
	return that.Ctx.Input.IP()
}

//@作者：Ven
//@函数详情：获取请求头头部为tk的token 信息
//@备注信息：无
//@参数：无
//@返回值：string
//@时间：2018-07-08
func (that *BaseController) GetToken() string {
	return that.GetHeaderByKey("tk")
}

//@作者：Ven
//@函数详情：ajax请求返回数据统一函数
//@备注信息：该函数会终结Controller 的运行
//@参数：[data]{interface{}类型传入返回客户端的数据体,[status]{int类型 为状态码}
//@返回值：无
//@时间：2018-07-08
func (that *BaseController) ajaxData(data interface{}, status int) {
	out := make(map[string]interface{})
	out["status"] = status
	out["data"] = data
	that.Data["json"] = out
	that.ServeJSON()
	that.StopRun()
}

/**
 * @作者: Ven
 * @详情: 无
 * @备注信息: 无
 * @param {interface{}} data
 * @param {int} status
 * @param {string} msg
 * @返回值: {*}
 * @时间: 2021-01-14 15:17:48
 */
func (that *BaseController) WriteMsg(data interface{}, status int, msg string) {
	out := make(map[string]interface{})
	out["status"] = status
	out["data"] = data
	out["msg"] = msg
	that.Data["json"] = out
	that.ServeJSON()
	that.StopRun()
}

/**
 * @作者: Ven
 * @详情: 无
 * @备注信息: 无
 * @param {interface{}} data
 * @param {int} status
 * @param {string} msg
 * @返回值: {*}
 * @时间: 2021-01-14 15:17:51
 */
func (that *BaseController) WriteList(data interface{}, status int, msg string) {
	out := make(map[string]interface{})
	out["status"] = status
	out["data"] = data
	out["msg"] = msg
	that.Data["json"] = out
	that.ServeJSON()
	that.StopRun()
}

//@作者：Ven
//@函数详情：当前系统时间戳毫秒数
//@备注信息：无
//@参数：
//@返回值：当前系统时间戳毫秒数
//@时间：2018-07-08
func (that *BaseController) NowUnixTime() int64 {
	return time.Now().Unix()
}

//@作者：Ven
//@函数详情：当前系统时间戳纳秒数
//@备注信息：无
//@参数：
//@返回值：当前系统时间戳毫秒数
//@时间：2018-07-08
func (that *BaseController) NowUnixNano() int64 {
	return time.Now().UnixNano()
}

//@作者：Ven
//@函数详情：获取当前请求的Body体
//@备注信息：无
//@参数：
//@返回值：包体字节数组
//@时间：2018-07-08
func (that *BaseController) GetRequestBody() []byte {
	return that.Ctx.Input.RequestBody
}

//@作者：Ven
//@函数详情：传入数据容器将请求体内容以json形式解析
//@备注信息：无
//@参数：object为传入传出参数
//@返回值：error
//@时间：2018-07-08
func (that *BaseController) GetRequestBodyToJson(object interface{}) error {
	return json.Unmarshal(that.Ctx.Input.RequestBody, &object)
}

//@作者：LXW
//@函数详情：传入数据容器将请求体内容以json形式解析，并不需使用科学计数法
//@备注信息：无
//@参数：object为传入传出参数
//@返回值：error
//@时间：2018-07-08
func (that *BaseController) GetRequestBodyToJsonUseNumber(object interface{}) error {
	bytes := that.Ctx.Input.RequestBody
	d := json.NewDecoder(strings.NewReader(string(bytes[:])))
	d.UseNumber()
	return d.Decode(object)
}

//@作者：Ven
//@函数详情：传入数据容器将请求体内容以xml形式解析
//@备注信息：无
//@参数：object为传入传出参数
//@返回值：error
//@时间：2018-07-08
func (that *BaseController) GetRequestBodyToXml(object interface{}) error {
	return xml.Unmarshal(that.Ctx.Input.RequestBody, &object)
}

//@作者：Ven
//@函数详情：将请求体内容以json形式解析返回map[string]interface{}数据
//@备注信息：无
//@参数：无
//@返回值：map 和 error
//@时间：2018-07-08
func (that *BaseController) GetRequestJsonToMap() (map[string]interface{}, error) {
	m := make(map[string]interface{}, 0)
	err := json.Unmarshal(that.Ctx.Input.RequestBody, &m)
	return m, err
}

//@作者：Ven
//@函数详情：
//@备注信息：无
//@参数：
//@返回值：
//@时间：2018-07-08
func (that *BaseController) GetRequestParam(key string) string {
	return that.Ctx.Input.Param(key)
}

//@作者：Ven
//@函数详情：
//@备注信息：无
//@参数：
//@返回值：
//@时间：2018-07-08
func (that *BaseController) GetRequestParams() map[string]string {
	return that.Ctx.Input.Params()
}

//@作者：Ven
//@函数详情：
//@备注信息：无
//@参数：
//@返回值：
//@时间：2018-07-08
func (that *BaseController) Query(key string) string {
	return that.Ctx.Input.Query(key)
}

//@作者：Ven
//@函数详情：获取header里制定的key值
//@备注信息：无
//@参数：[key]{string}
//@返回值：header内key对应的value
//@时间：2018-07-08
func (that *BaseController) GetHeaderByKey(key string) string {
	return that.Ctx.Request.Header.Get(key)
}

//@作者：Ven
//@函数详情：获取请求头的 Header结构
//@备注信息：无
//@参数：无
//@返回值：http.Header
//@时间：2018-07-08
func (that *BaseController) GetHeader() http.Header {
	return that.Ctx.Request.Header
}

//@作者：Ven
//@函数详情：获取请求的url
//@备注信息：无
//@参数：
//@返回值：
//@时间：2018-07-08
func (that *BaseController) GetRequestUrl() string {
	return that.Ctx.Input.URL()
}
