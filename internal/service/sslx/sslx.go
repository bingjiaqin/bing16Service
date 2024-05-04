package sslx

import (
	"bing16Service/internal/interface/file"
	"bing16Service/internal/interface/git"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type SslxItem struct {
	Content string `json:"content"`

	Time string `json:"time"`

	Pic string `jspn:"pic"`
}

type SslxData struct {
	Year string `json:"year"`

	SubTitle string `json:"subTitle"`

	Lists []SslxItem `json:"lists"`
}

func Add(data string) error {
	var sslxDataPath = viper.GetString("project.component_path") + "/sslx/sslx.js"

	prefix := "export const DATA = "

	sslxDatas, err := file.ReadJson[SslxData](sslxDataPath, prefix)

	if err != nil {
		return err
	}

	var newItem SslxItem
	newItem.Content = "&emsp;&emsp;" + strings.ReplaceAll(data, "\n", "<br/>&emsp;&emsp;") + "<br/>"
	newItem.Time = time.Now().Format("2006-01-02 15:04")

	latestData := sslxDatas[len(sslxDatas)-1]
	sslxDatas[len(sslxDatas)-1].Lists = append(latestData.Lists, newItem)

	file.WriteJson(sslxDataPath, prefix, sslxDatas)

	err1 := git.CommitAndPush("add sslx content")

	if err1 != nil {
		return err1
	}

	return nil
}
