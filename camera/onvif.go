package camera

import (
	"encoding/xml"
	"github.com/zgwit/gonvif"
	"github.com/zgwit/gonvif/media"
	"io/ioutil"
	"net/http"
)

func readResponse(resp *http.Response) string {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func GetOnvifStream(addr, username, password string) (string, error) {
	dev, err := gonvif.NewDevice(
		gonvif.DeviceParams{
			Xaddr:      addr,
			Username:   username,
			Password:   password,
			HttpClient: nil,
		})
	if err != nil {
		return "", err
	}
	su := media.GetStreamUri{
		ProfileToken: "Profile_101",
	}
	resp, err := dev.CallMethod(su)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var msu media.GetStreamUriResponse
	err = xml.Unmarshal(b, &msu)
	if err != nil {
		return "", err
	}

	return string(msu.MediaUri.Uri), nil
}
