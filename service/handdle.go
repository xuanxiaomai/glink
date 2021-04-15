package service

import (
	"Baiyuetribe/glink/api"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// 接入，创建API及方法，批量引入
func ApiHandler(c *fiber.Ctx) error {
	uri := c.Request().URI().String()
	resultList := strings.Split(uri, "/api/")
	if len(resultList) != 2 {
		return c.JSON(fiber.Map{"msg": "请输入正确的url地址"})
	}
	url := resultList[1]
	// 格局url匹配函数
	var r string
	if strings.Contains(url, "douyin") { // 没有匹配时，值为-1
		r = api.DouYin(url)
	} else if strings.Contains(url, "weishi") {
		r = api.WeiShi(url)
	} else if strings.Contains(url, "pipix") {
		r = api.PiPix(url)
	} else if strings.Contains(url, "kuaishou") {
		r = api.KuaiShou(url)
	} else if strings.Contains(url, "kg3.qq.com") {
		r = api.Kg3(url)
	} else if strings.Contains(url, "ixigua.com") {
		r = "截至2021-04-04，仅能通过下载合并的方式获取无无水印视频"
	} else if strings.Contains(url, "eyepetizer") {
		r = api.EyePetizer(url)
	} else if strings.Contains(url, "vuevideo") {
		r = api.VueVlog(url)
	} else if strings.Contains(url, "xiaokaxiu") {
		r = "暂未支持该接口，请提交issue"
	} else if strings.Contains(url, "ippzone") {
		r = api.IppZone(url)
	} else if strings.Contains(url, "weibo.com") {
		r = api.WeiBo(url)
	} else if strings.Contains(url, "zuiyou") {
		r = api.ZuiYou(url)
	} else if strings.Contains(url, "bbq.bilibili") {
		r = api.QingShi(url)
	} else if strings.Contains(url, "bilibili.com") {
		r = api.Bilibili(url)
	} else if strings.Contains(url, "immomo") {
		r = api.MonMo(url)
	} else {
		r = "暂未支持该接口，请提交issue"
	}
	return c.SendString(r)
	// return c.JSON(&fiber.Map{"msg": r})

}

func PrintLogo() {
	fmt.Print(`

	██████╗  █████╗ ██╗██╗   ██╗██╗   ██╗███████╗
	██╔══██╗██╔══██╗██║╚██╗ ██╔╝██║   ██║██╔════╝
	██████╔╝███████║██║ ╚████╔╝ ██║   ██║█████╗  
	██╔══██╗██╔══██║██║  ╚██╔╝  ██║   ██║██╔══╝  
	██████╔╝██║  ██║██║   ██║   ╚██████╔╝███████╗
	╚═════╝ ╚═╝  ╚═╝╚═╝   ╚═╝    ╚═════╝ ╚══════╝
			欢迎使用Glink短视频去水印软件
			作者QQ：2894049053

`)
}
