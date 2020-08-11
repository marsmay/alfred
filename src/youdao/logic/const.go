package logic

import "regexp"

const AppName = "cn.ac.may.youdao"
const ApiUri = "https://openapi.youdao.com/api"
const QuicklookUrl = "http://youdao.com/w/"

const LangAuto = "auto"
const CodeSuccess = "0"

const (
	ConvertCn2En = "zh-CHS2en"
	ConvertEn2Cn = "en2zh-CHS"
)

const (
	IconTranslateSay = "icon/translate-say.png"
	IconTranslate    = "icon/translate.png"
)

var SetKeyRegex = regexp.MustCompile(`^key (\w{16}) (\w{32})$`)
