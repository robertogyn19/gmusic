package gmusic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/google/go-querystring/query"
)

type SettingsData struct {
	Settings Settings `json:"settings"`
}

type Settings struct {
	EntitlementInfo struct {
		ExpirationMillis int  `json:"expirationMillis"`
		IsCanceled       bool `json:"isCanceled"`
		IsSubscription   bool `json:"isSubscription"`
		IsTrial          bool `json:"isTrial"`
	} `json:"entitlementInfo"`
	Lab []struct {
		Description    string `json:"description"`
		DisplayName    string `json:"displayName"`
		Enabled        bool   `json:"enabled"`
		ExperimentName string `json:"experimentName"`
	} `json:"lab"`
	MaxUploadedTracks      int  `json:"maxUploadedTracks"`
	SubscriptionNewsletter bool `json:"subscriptionNewsletter"`
	UploadDevice           []struct {
		DeviceType             int    `json:"deviceType"`
		ID                     string `json:"id"`
		LastAccessedFormatted  string `json:"lastAccessedFormatted"`
		LastAccessedTimeMillis int    `json:"lastAccessedTimeMillis"`
		LastEventTimeMillis    int    `json:"lastEventTimeMillis"`
		Name                   string `json:"name"`
	} `json:"uploadDevice"`
}

func (g *GMusic) request(method, url string, data interface{}, client *http.Client) (*http.Response, error) {
	var body io.Reader
	if data != nil {
		str, isString := data.(string)
		barray, isByteArray := data.([]byte)

		buf := new(bytes.Buffer)

		switch {
		case isString:
			buf = bytes.NewBufferString(str)
		case isByteArray:
			buf = bytes.NewBuffer(barray)
		default:
			if err := json.NewEncoder(buf).Encode(data); err != nil {
				return nil, err
			}
		}

		if method == "GET" {
			params, _ := query.Values(data)
			url = fmt.Sprintf("%s?%s", url, params.Encode())
		} else {
			body = buf
		}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	auth := fmt.Sprintf("GoogleLogin auth=%s", g.Auth)
	req.Header.Add("Authorization", auth)
	req.Header.Add("Content-Type", "application/json")

	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		return nil, fmt.Errorf("gmusic: %s", resp.Status)
	}
	return resp, nil
}

func (g *GMusic) sjRequest(method, path string, data interface{}) (*http.Response, error) {
	return g.request(method, sjURL+path, data, nil)
}

func (g *GMusic) setDeviceID() error {
	const phoneDevice = 2
	req, err := http.NewRequest("HEAD", googlePlayMusicEndpoint+"/listen", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "GoogleLogin auth="+g.Auth)
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	xt := make(url.Values)
	u, _ := url.Parse(googlePlayMusicEndpoint)
	for _, c := range jar.Cookies(u) {
		if c.Name == "xt" {
			xt.Set("xt", c.Value)
		}
	}
	settings, err := g.settings(xt, client)
	if err != nil {
		return err
	}
	for _, d := range settings.UploadDevice {
		if d.DeviceType != phoneDevice || len(d.ID) != 18 {
			continue
		}
		g.DeviceID = d.ID[2:]
		break
	}
	if g.DeviceID == "" {
		return fmt.Errorf("no valid devices")
	}
	return nil
}

func (g *GMusic) settings(xtData url.Values, jarClient *http.Client) (*Settings, error) {
	resp, err := g.request("POST", googlePlayMusicEndpoint+"/services/fetchsettings?"+xtData.Encode(), nil, jarClient)
	if err != nil {
		return nil, err
	}
	var data SettingsData
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data.Settings, nil
}
