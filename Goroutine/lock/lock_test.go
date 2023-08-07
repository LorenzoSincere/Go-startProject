package main

import (
	"os"
	"testing"
)

func TestLock(t *testing.T) {

}

func TestMain(m *testing.M) {
	//测试前：数据装载，配置初始化等前置工作
	code := m.Run()
	//测试后： 释放资源等收尾工作
	os.Exit(code)
}
