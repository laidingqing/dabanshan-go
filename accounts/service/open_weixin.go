package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Jscode2Session success code
type Jscode2Session struct {
	ExpiresIn  int64  `json:"expiresIn"`
	OpenID     string `json:"openID"`
	SessionKey string `json:"sessionKey"`
	UserID     string `json:"userId,omitempty"`
	Name       string `json:"name,omitempty"`
}

//fetch wx session by oauth code
func getWxSession(code string) (Jscode2Session, int64, error) {
	var ret Jscode2Session
	if code == "" {
		return ret, 500001, fmt.Errorf("Code为空")
	}
	urls := url.Values{}
	urls.Add("appid", "wxf5a6ca5f27f3d5cc")
	urls.Add("secret", "ef76ee8e04486cd4c4ba81188b5b9ccd")
	urls.Add("js_code", code)
	urls.Add("grant_type", "authorization_code")
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?" + urls.Encode())
	res, err := http.Get(url)
	if err != nil {
		return ret, 500002, err
	}
	info, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ret, 500003, err
	}
	err = json.Unmarshal(info, &ret)
	if err != nil || ret.OpenID == "" {
		return ret, 500004, fmt.Errorf(string(info))
	}
	return ret, 0, nil
}
