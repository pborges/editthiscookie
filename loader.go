package editthiscookie

import (
	"net/url"
	"net/http"
	"encoding/json"
	"time"
	"net/http/cookiejar"
	"io"
)

type cookie struct {
	Domain         string  `json:"domain"`
	ExpirationDate float64 `json:"expirationDate"`
	HostOnly       bool    `json:"hostOnly"`
	HTTPOnly       bool    `json:"httpOnly"`
	Name           string  `json:"name"`
	Path           string  `json:"path"`
	SameSite       string  `json:"sameSite"`
	Secure         bool    `json:"secure"`
	Session        bool    `json:"session"`
	StoreID        string  `json:"storeId"`
	Value          string  `json:"value"`
	ID             int     `json:"id"`
}

func Load(client *http.Client, reader io.Reader, u *url.URL) error {
	etc := make([]cookie, 0)
	err := json.NewDecoder(reader).Decode(&etc)
	if err != nil {
		return err
	}
	cookies := make([]*http.Cookie, len(etc))
	for i, c := range etc {
		cookies[i] = new(http.Cookie)
		cookies[i].Name = c.Name
		cookies[i].Value = c.Value
		cookies[i].Domain = c.Domain
		cookies[i].Expires = time.Unix(int64(c.ExpirationDate), 0)
		cookies[i].HttpOnly = c.HostOnly
		cookies[i].Path = c.Path
		cookies[i].Secure = c.Secure
	}
	client.Jar, err = cookiejar.New(nil)
	if err != nil {
		return err
	}
	client.Jar.SetCookies(u, cookies)
	return nil
}
