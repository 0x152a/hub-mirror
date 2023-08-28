package main

import (
	"context"
	"strings"
	"fmt"
	"os"
	"sync"
	"text/template"

	"github.com/spf13/pflag"
	"github.com/togettoyou/hub-mirror/pkg"
)

var (
	content    = pflag.StringP("content", "", "", "原始镜像，source|platform")
	maxContent = pflag.IntP("maxContent", "", 11, "原始镜像个数限制")
	repository = pflag.StringP("repository", "", "", "推送仓库地址，为空默认为 hub.docker.com")
	username   = pflag.StringP("username", "", "", "仓库用户名")
	password   = pflag.StringP("password", "", "", "仓库密码")
	outputPath = pflag.StringP("outputPath", "", "output.sh", "结果输出路径")
)

func main() {
	pflag.Parse()
	var platform string = ""
	var source string


	lines := strings.Split(*content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			index := strings.Index(source, "|")
			if index >= 0 {
				platform = strings.TrimSpace(source[index+1:])
				source = strings.TrimSpace(source[:index])
			} else {
				source = line
			}
			break
		}
	}

	fmt.Printf("mirrors: %s, platform: %s\n", source, platform)

	fmt.Println("初始化 Docker 客户端")
	cli, err := pkg.NewCli(context.Background(), *repository, *username, *password, os.Stdout)
	if err != nil {
		panic(err)
	}

	outputs := make([]*pkg.Output, 0)
	wg := sync.WaitGroup{}


	fmt.Println("开始转换镜像", source)
	wg.Add(1)
	go func() {
		defer wg.Done()

		output, err := cli.PullTagPushImage(context.Background(), source, platform)
		if err != nil {
			fmt.Println(source, "转换异常: ", err)
		} else {
			outputs = append(outputs, output)
		}
	}()

	wg.Wait()

	if len(outputs) == 0 {
		panic("没有转换成功的镜像")
	}

	tmpl, err := template.New("pull_images").Parse(
		`{{if .Repository}}# if your repository is private,please login...
# docker login {{ .Repository }} --username={your username}
{{end}}
{{- range .Outputs }}
docker pull {{ .Target }}
docker tag {{ .Target }} {{ .Source }}{{ end }}`)
	if err != nil {
		panic(err)
	}
	outputFile, err := os.Create(*outputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, map[string]interface{}{
		"Outputs":    outputs,
		"Repository": *repository,
	})
	if err != nil {
		panic(err)
	}
}
