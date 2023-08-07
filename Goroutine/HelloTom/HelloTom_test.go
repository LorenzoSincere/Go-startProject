package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestHelloTom(t *testing.T) {
//	outPuts := HelloTom()
//	expectOutPut := "Tom"
//
//	if outPuts != expectOutPut {
//		t.Errorf("expected %s do not match actual %s", expectOutPut, outPuts)
//	}
//
//	assert.Equal(t, expectOutPut, outPuts)
//
//}

func TestJudgePassLineTrue(t *testing.T) {
	isPass := JudgePassLine(70)
	assert.Equal(t, true, isPass)
}

func TestJudgePassLineFalse(t *testing.T) {
	isPass := JudgePassLine(50)
	assert.Equal(t, false, isPass)
}
