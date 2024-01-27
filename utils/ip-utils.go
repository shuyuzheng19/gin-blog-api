package utils

import (
	"common-web-framework/helper"
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"net/http"
	"strings"
)

func GetIPAddress(request *http.Request) string {
	ipAddress := request.Header.Get("X-Forwarded-For")
	if ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = request.Header.Get("Proxy-Client-IP")
	}
	if ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = request.Header.Get("WL-Proxy-Client-IP")
	}
	if ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = request.RemoteAddr
	}
	return ipAddress
}

func IsCnIp(ip string) bool {
	var city = GetIpCity(ip)

	if strings.Contains(city, "中国") && !strings.Contains(city, "香港") && !strings.Contains(city, "台湾") {
		return true
	}

	return false
}

var searcher *xdb.Searcher

func LoadIpDB(dbPath string) {
	var err error

	searcher, err = xdb.NewWithFileOnly(dbPath)

	if err != nil {
		helper.ErrorPanicAndMessage(err, "加载IP数据库失败")
	}
}

func GetIpCity(ip string) string {
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		return "未知"
	}

	var split = strings.Split(region, "|")

	return strings.ReplaceAll(split[0]+" "+split[2]+" "+split[3], "0", "")
}

// GetClientPlatformInfo 获取客户端平台信息
func GetClientPlatformInfo(userAgent string) string {
	if userAgent == "" {
		return ""
	}

	userAgent = strings.ToLower(userAgent)

	var os, browser string
	// 匹配操作系统
	switch {
	case strings.Contains(userAgent, "windows"):
		os = "Windows"
	case strings.Contains(userAgent, "mac"):
		os = "Mac"
	case strings.Contains(userAgent, "android"):
		os = "Android"
	case strings.Contains(userAgent, "iphone") || strings.Contains(userAgent, "ipad"):
		os = "iOS"
	}
	// 匹配浏览器
	switch {
	case strings.Contains(userAgent, "micromessenger"):
		browser = "微信客户端"
	case strings.Contains(userAgent, "edg"):
		browser = "Edge"
	case strings.Contains(userAgent, "chrome"):
		browser = "Chrome"
	case strings.Contains(userAgent, "firefox"):
		browser = "Firefox"
	case strings.Contains(userAgent, "safari"):
		browser = "Safari"
	}

	if os != "" && browser != "" {
		return fmt.Sprintf("%s %s", os, browser)
	} else {
		return userAgent
	}
}
