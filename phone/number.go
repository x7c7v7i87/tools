package phone

//过滤中国手机号码
func ChFilter(str string) string {
	if len(str) <= 10 {
		return str
	}
	return str[:3] + "*****" + str[len(str)-3:]
}