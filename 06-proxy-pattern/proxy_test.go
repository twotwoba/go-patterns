package proxy

import "testing"

func TestProxyObject(t *testing.T) {
	proxyObject := new(ProxyObject)
	proxyObject.ObjDo("run")
	// 代理控制了，不会执行
	proxyObject.ObjDo("forbidden")
}
