package cmds

import (
	"fmt"
	"net/http"

	"github.com/antchfx/htmlquery"

	"github.com/sirupsen/logrus"
)

func wnacgCheck(no string) (string, error) {
	var retrunString string
	url := fmt.Sprintf("https://www.wnacg.com/photos-index-aid-%s.html", no)
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		h2Node := htmlquery.FindOne(doc, "//h2")
		if h2Node != nil {
			logrus.Info(fmt.Sprintf("指定本子(%s)確認成功", no))
			retrunString = url
		} else {
			logrus.Errorf(fmt.Sprintf("找不到指定本子(%s)", no))
			retrunString = fmt.Sprintf("找不到指定本子(%s)", no)
		}
	case http.StatusNotFound:
		logrus.Errorf(fmt.Sprintf("找不到指定本子(%s)", no))
		retrunString = fmt.Sprintf("找不到指定本子(%s)", no)
	default:
		logrus.Errorf("伺服器端錯誤,請聯繫管理員")
		retrunString = "伺服器端錯誤,請聯繫管理員"
	}
	return retrunString, nil
}
