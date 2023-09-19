/*
 * @FilePath: /sxwz-bot/pkg/fans/fans.go
 * @Author: maggot-code
 * @Date: 2023-09-19 07:58:21
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-19 11:45:43
 * @Description:
 */
package fans

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type FansBody struct {
	Fans int32     `json:"fans"`
	Day  int32     `json:"rate1"`
	Week int32     `json:"rate7"`
	Date time.Time `json:"created_at"`
}

type FansUserStats struct {
	Total int32 `json:"fans"`
	Trend int32 `json:"rate1"`
}

type FnasUser struct {
	Name  string `json:"name"`
	Stats FansUserStats
}

type FansUseCase struct {
	uid   string
	total int
	user  FnasUser
	group []FansBody
}

var api = "https://api.zeroroku.com/bilibili/author"

func NewFansRepo(uid int32) (*FansUseCase, error) {
	fans := &FansUseCase{
		uid: fmt.Sprintf("%d", uid),
	}

	if err := fans.GetUser(); err != nil {
		return nil, err
	}

	if err := fans.GetFans(); err != nil {
		return nil, err
	}

	fans.Sort()

	return fans, nil
}

func (fc *FansUseCase) GetUser() error {
	query := url.Values{}
	query.Set("mid", fc.uid)

	uri, err := url.Parse(api)
	if err != nil {
		return err
	}

	uri.RawQuery = query.Encode()
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	res, err := client.Get(uri.String())
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&fc.user)
	if err != nil {
		return err
	}

	return nil
}

func (fc *FansUseCase) GetFans() error {
	query := url.Values{}
	query.Set("mid", fc.uid)

	uri, err := url.Parse(api + "/fans")
	if err != nil {
		return err
	}

	uri.RawQuery = query.Encode()
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	res, err := client.Get(uri.String())
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&fc.group)
	if err != nil {
		return err
	}

	fc.total = len(fc.group)
	return nil
}

func (fc *FansUseCase) Sort() {
	sort.Slice(fc.group, func(i, j int) bool {
		return fc.group[i].Date.Before(fc.group[j].Date)
	})
}

func (fc *FansUseCase) FormatTotal(str string) string {
	parts := strings.Split(str, ".")
	intPart := parts[0]
	fracPart := ""
	if len(parts) > 1 {
		fracPart = fmt.Sprintf(".%s", parts[1])
	}

	n := len(intPart)
	if n <= 3 {
		return fmt.Sprintf("%s%s", intPart, fracPart)
	}

	return fmt.Sprintf("%s,%s%s", fc.FormatTotal(intPart[:n-3]), intPart[n-3:], fracPart)
}

func (fc *FansUseCase) FormatTrend() string {
	var trend string

	if fc.user.Stats.Trend > 0 {
		trend = fmt.Sprintf("+%d", fc.user.Stats.Trend)
	} else {
		trend = fmt.Sprintf("%d", fc.user.Stats.Trend)
	}

	return trend
}

func (fc *FansUseCase) CountIncrease() int32 {
	return fc.group[fc.total-1].Fans - fc.group[fc.total-2].Fans
}

func (fc *FansUseCase) CountReduce() int32 {
	return fc.group[fc.total-2].Fans - fc.group[fc.total-1].Fans
}

func (fc *FansUseCase) ToUserTotal() string {
	total := strconv.FormatInt(int64(fc.user.Stats.Total), 10)
	return fmt.Sprintf("%s：%s (%s)", fc.user.Name, fc.FormatTotal(total), fc.FormatTrend())
}

func (fc *FansUseCase) ToUserTrue() string {
	return fmt.Sprintf("%s：实际增加：%d | 取关：%d", fc.user.Name, fc.CountIncrease(), fc.CountReduce())
}
