package util

import (
	"crypto/md5"
	"encoding/hex"
	"net"
	"net/http"
)
//md5加密算法
func MD5Encry(s string) string{
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

func RemoteIp(req *http.Request)string{
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("XRealIP") ; ip != "" {
		remoteAddr = ip
	}else if ip = req.Header.Get("XForwardedFor");ip != "" {
		remoteAddr = ip
	}else{
		remoteAddr,_,_ = net.SplitHostPort(remoteAddr)
	}
	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}
