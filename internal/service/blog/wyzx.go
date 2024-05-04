package blog

import (
	"bing16Service/internal/interface/file"
	"strings"

	"github.com/spf13/viper"
)

type WyzxData struct {
	Title string `json:"title"`

	Pic string `json:"pic"`

	Intro string `json:"intro"`

	Href string `json:"href"`

	Tags []string `json:"tags"`

	Type string `json:"type"`
}

func AddWyzx(blog Blog) {
	var blogFilePath = viper.GetString("project.blog_root") + "/" + blog.Path + "/" + blog.Title
	err := file.WriteFile(blogFilePath, blog.Context)

	if err != nil {
		return
	}

	prefix := "export const DATA = "

	var componetnInfoPath = viper.GetString("project.component_path") + "/" + blog.Path + "/" + blog.Path + ".js"
	wyzxDatas, err := file.ReadJson[WyzxData](componetnInfoPath, prefix)

	if err != nil {
		return
	}

	title := blog.Title[:strings.Index(blog.Title, ".")]
	suffix := blog.Title[strings.Index(blog.Title, ".")+1:]

	var newWyzx WyzxData

	for _, wyzxData := range wyzxDatas {
		if title == wyzxData.Title {
			newWyzx = wyzxData
			newWyzx.Intro = blog.Intro
			newWyzx.Type = suffix
			newWyzx.Href = "wyzx" + "/" + title
			break
		}
	}

	if newWyzx.Title == "" {
		newWyzx.Href = "wyzx" + "/" + title
		newWyzx.Title = title
		newWyzx.Type = suffix
		newWyzx.Intro = blog.Intro
		wyzxDatas = append(wyzxDatas, newWyzx)
	}

	file.WriteJson(componetnInfoPath, prefix, wyzxDatas)

	// err1 := git.CommitAndPush("add new blog")

	// if err1 != nil {
	// 	return
	// }
}
