package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/luoyanke/gitlab-webhook-tool/internal"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func main() {

	var feishuWebhook string
	//开启的端口
	var port int

	// 解析命令行参数
	flag.StringVar(&feishuWebhook, "feishuWebhook", "", "")
	flag.IntVar(&port, "port", 6710, "6710")
	flag.Parse()

	http.HandleFunc("/web-hook", func(writer http.ResponseWriter, request *http.Request) {

		var bodyBytes, _ = ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		//body := string(bodyBytes)
		//log.Print(body)

		var baseBody internal.BaseBody
		err := json.Unmarshal(bodyBytes, &baseBody)
		if err != nil {
			log.Fatal(err)
			return
		}
		if baseBody.ObjectKind == "merge_request" {
			mergeRequestNotify(bodyBytes, feishuWebhook)
		} else if baseBody.ObjectKind == "push" {
			pushNotify(bodyBytes, feishuWebhook)
		}
	})

	// 启动 HTTP 服务器
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		os.Exit(1)
	}
}

func mergeRequestNotify(bodyBytes []byte, feishuWebhook string) {
	var body internal.MergeRequestBody
	var writer bytes.Buffer
	err := json.Unmarshal(bodyBytes, &body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var title string
	if body.ObjectAttributes.State == "opened" {
		title = body.Project.Name + " 合并请求提交事件"
	} else if body.ObjectAttributes.State == "merged" {
		title = body.Project.Name + " 合并请求完成事件"
	} else {
		title = body.Project.Name + " 合并请求事件"
	}

	tmpl, _ := template.New("index").Parse(internal.MergeRequestFeishuCardTmpl())
	tmpl.Execute(&writer, map[string]interface{}{
		"projectName":  body.Project.Name,
		"userName":     body.User.Username + "(" + body.User.Name + ")",
		"sourceBranch": body.ObjectAttributes.SourceBranch,
		"targetBranch": body.ObjectAttributes.TargetBranch,
		"webUrl":       body.Project.WebURL + "/merge_requests",
		"title":        title,
	})
	var cardBody internal.FeishuCard
	cardBody.MsgType = "interactive"
	cardBody.Card = writer.String()
	//log.Print(cardBody.Card)
	marshal, err := json.Marshal(cardBody)
	req, err := http.NewRequest("POST", feishuWebhook, bytes.NewBuffer(marshal))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	} else {
		//var bodyBytes, _ = ioutil.ReadAll(req.Body)
		//s := string(bodyBytes)
		//log.Print(s)
	}
	defer resp.Body.Close()
}

func pushNotify(bodyBytes []byte, feishuWebhook string) {
	var body internal.PushRequestBody
	var writer bytes.Buffer
	err := json.Unmarshal(bodyBytes, &body)
	if err != nil {
		log.Fatal(err)
		return
	}
	var commits string
	for index, obj := range body.Commits {
		msg := strings.ReplaceAll(obj.Message, "\n", "")
		commits += "- " + obj.Author.Name + "< " + obj.Author.Email + " >:   " + msg + "\\n"
		if index > 8 {
			i := len(body.Commits) - index
			commits += "-  other " + strconv.Itoa(i) + " commit  ...\\n"
			break
		}
	}
	branch := strings.Replace(body.Ref, "refs/heads/", "", 1)
	tmpl, _ := template.New("index").Parse(internal.PushFeishuCardTmpl())
	tmpl.Execute(&writer, map[string]interface{}{
		"projectName": body.Project.Name,
		"userName":    body.UserName,
		"ref":         body.Ref,
		"webUrl":      body.Project.WebURL + "/commits/" + branch,
		"commit":      commits,
		"title":       body.Project.Name + " 代码推送事件",
	})

	var cardBody internal.FeishuCard
	cardBody.MsgType = "interactive"
	cardBody.Card = writer.String()
	//log.Print(cardBody.Card)
	marshal, err := json.Marshal(cardBody)
	req, err := http.NewRequest("POST", feishuWebhook, bytes.NewBuffer(marshal))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	} else {
		//var bodyBytes, _ = ioutil.ReadAll(req.Body)
		//s := string(bodyBytes)
		//log.Print(s)
	}
	defer resp.Body.Close()
}
