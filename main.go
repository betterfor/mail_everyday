package main

import (
	_ "embed"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

const (
	// 墨迹天气api
	weatherApi = "https://tianqi.moji.com/weather/china/"
	// one api
	oneApi = "http://m.wufazhuce.com/one"
)

var (
	// 天气地区
	local = "anhui/yaohai-district"
)

//go:embed mail.tpl
var htmlTpl string

const outputFile = "output.html"

func main() {
	c := GetMailContent()
	t, err := template.New("").Parse(htmlTpl)
	if err != nil {
		log.Fatal("cannot parse html template error: ", err)
	}

	fi, err := os.OpenFile("output/"+outputFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal("cannot open file output/output.html", err)
	}

	err = t.Execute(fi, c)
	if err != nil {
		log.Fatal("cannot execute html template error: ", err)
	}
}

func GetMailContent() MailContent {
	oneData, err := getOneData()
	if err != nil {
		fmt.Println(err)
	}
	tips, days, err := getWeather()
	if err != nil {
		fmt.Println(err)
	}
	return MailContent{
		OneData:   *oneData,
		Tips:      tips,
		ThreeDays: days,
	}
}

type MailContent struct {
	OneData   OneData
	Tips      string
	ThreeDays []DayData
}

type OneData struct {
	ImgUrl  string // 图片地址
	Type    string // 类型
	Content string // 一段话
	Date    string // 日期
}

// 获取ONE内容
func getOneData() (*OneData, error) {
	resp, err := http.Get(oneApi)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var token string
	dom.Find("script").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "One.token") {
			tokenStart := strings.Index(s.Text(), "'")
			tokenEnd := strings.Index(s.Text()[tokenStart+1:], "'")
			token = s.Text()[tokenStart+1 : tokenStart+1+tokenEnd]
		}
	})

	cookies := resp.Cookies()

	req, err := http.NewRequest(http.MethodGet, "http://m.wufazhuce.com/one/ajaxlist/0?_token="+token, nil)
	if err != nil {
		return nil, err
	}

	for _, c := range cookies {
		req.AddCookie(c)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ret := gjson.GetBytes(bts, "data.0").Map()
	return &OneData{
		ImgUrl:  ret["img_url"].String(),
		Type:    ret["picture_author"].String()[:strings.Index(ret["picture_author"].String(), "|")],
		Content: ret["content"].String(),
		Date:    ret["date"].String(),
	}, nil
}

// 获取天气信息
func getWeather() (string, []DayData, error) {
	api := weatherApi + local
	resp, err := http.Get(api)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", nil, errors.New(resp.Status)
	}

	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", nil, err
	}
	tips := getWeatherTips(dom)
	datas := getWeatherData(dom)
	return tips, datas, nil
}

// 获取天气提醒
func getWeatherTips(dom *goquery.Document) string {
	return dom.Find(".wea_tips").Find("em").Text()
}

type DayData struct {
	Day            string // 日期
	WeatherImgUrl  string // 天气图
	WeatherText    string // 天气描述
	Temperature    string // 温度
	WindDirection  string // 风向
	WindLevel      string // 风力
	Pollution      string // 空气质量
	PollutionLevel string // 空气质量等级
}

// 获取天气预报
func getWeatherData(dom *goquery.Document) (threeDaysData []DayData) {
	dom.Find(".forecast .days").Each(func(i int, s *goquery.Selection) {
		var dayData DayData
		singleDay := s.Find("li")
		singleDay.Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				dayData.Day = strings.TrimSpace(s.Text())
			case 1:
				w := s.Find("img")
				dayData.WeatherImgUrl, _ = w.Attr("src")
				dayData.WeatherText, _ = w.Attr("alt")
			case 2:
				dayData.Temperature = s.Text()
			case 3:
				dayData.WindDirection = s.Find("em").Text()
				dayData.WindLevel = s.Find("b").Text()
			case 4:
				dayData.Pollution = strings.TrimSpace(s.Text())
				dayData.PollutionLevel, _ = s.Find("strong").Attr("class")
			}

		})
		threeDaysData = append(threeDaysData, dayData)
	})
	return
}
