package gofakeit

import "strconv"

// Generate Random User Agent
func UserAgent() string {
	randNum := randIntRange(0, 3)
	switch randNum {
	case 0:
		return ChromeUserAgent()
	case 1:
		return FirefoxUserAgent()
	case 2:
		return SafariUserAgent()
	case 3:
		return OperaUserAgent()
	}

	return ChromeUserAgent()
}

// Generate Random Chrome User Agent
func ChromeUserAgent() string {
	randNum1 := strconv.Itoa(randIntRange(531, 536)) + strconv.Itoa(randIntRange(0, 2))
	randNum2 := strconv.Itoa(randIntRange(36, 40))
	randNum3 := strconv.Itoa(randIntRange(800, 899))
	return "Mozilla/5.0 " + "(" + randomPlatform() + ") AppleWebKit/" + randNum1 + " (KHTML, like Gecko) Chrome/" + randNum2 + ".0." + randNum3 + ".0 Mobile Safari/" + randNum1
}

// Generate Random Firefox User Agent
func FirefoxUserAgent() string {
	ver := "Gecko/" + Date().Format("2006-02-01") + " Firefox/" + strconv.Itoa(randIntRange(35, 37)) + ".0"
	platforms := []string{
		"(" + windowsPlatformToken() + "; " + "en-US" + "; rv:1.9." + strconv.Itoa(randIntRange(0, 3)) + ".20) " + ver,
		"(" + linuxPlatformToken() + "; rv:" + strconv.Itoa(randIntRange(5, 8)) + ".0) " + ver,
		"(" + macPlatformToken() + " rv:" + strconv.Itoa(randIntRange(2, 7)) + ".0) " + ver,
	}

	return "Mozilla/5.0 " + ShuffleStrings(platforms)[0]
}

// Generate Random Safari User Agent
func SafariUserAgent() string {
	randNum := strconv.Itoa(randIntRange(531, 536)) + "." + strconv.Itoa(randIntRange(1, 51)) + "." + strconv.Itoa(randIntRange(1, 8))
	ver := strconv.Itoa(randIntRange(4, 6)) + "." + strconv.Itoa(randIntRange(0, 2))

	mobileDevices := []string{
		"iPhone; CPU iPhone OS",
		"iPad; CPU OS",
	}

	platforms := []string{
		"(Windows; U; " + windowsPlatformToken() + ") AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + ver + " Safari/" + randNum,
		"(" + macPlatformToken() + " rv:" + strconv.Itoa(randIntRange(4, 7)) + ".0; en-US) AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + ver + " Safari/" + randNum,
		"(" + ShuffleStrings(mobileDevices)[0] + " " + strconv.Itoa(randIntRange(7, 9)) + "_" + strconv.Itoa(randIntRange(0, 3)) + "_" + strconv.Itoa(randIntRange(1, 3)) + " like Mac OS X; " + "en-US" + ") AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + strconv.Itoa(randIntRange(3, 5)) + ".0.5 Mobile/8B" + strconv.Itoa(randIntRange(111, 120)) + " Safari/6" + randNum,
	}

	return "Mozilla/5.0 " + ShuffleStrings(platforms)[0]
}

// Generate Random Opera User Agent
func OperaUserAgent() string {
	platform := "(" + randomPlatform() + "; en-US) Presto/2." + strconv.Itoa(randIntRange(8, 13)) + "." + strconv.Itoa(randIntRange(160, 355)) + " Version/" + strconv.Itoa(randIntRange(10, 13)) + ".00"

	return "Opera/" + strconv.Itoa(randIntRange(8, 10)) + "." + strconv.Itoa(randIntRange(10, 99)) + " " + platform
}

// Random linux platform
func linuxPlatformToken() string {
	return "X11; Linux " + getRandValue([]string{"useragent", "linux_processor"})
}

// Random mac platform
func macPlatformToken() string {
	return "Macintosh; " + getRandValue([]string{"useragent", "mac_processor"}) + " Mac OS X 10_" + strconv.Itoa(randIntRange(5, 9)) + "_" + strconv.Itoa(randIntRange(0, 10))
}

// Random windows platform
func windowsPlatformToken() string {
	return getRandValue([]string{"useragent", "windows_platform"})
}

func randomPlatform() string {
	platforms := []string{
		linuxPlatformToken(),
		macPlatformToken(),
		windowsPlatformToken(),
	}

	return ShuffleStrings(platforms)[0]
}
