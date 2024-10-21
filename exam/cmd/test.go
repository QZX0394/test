/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"exam/client"
	"exam/client/request"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

var addr string
var port int

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("test called")
		addr, _ := cmd.Flags().GetString("addr")
		port1, _ := cmd.Flags().GetInt("port1")
		port2, _ := cmd.Flags().GetInt("port2")

		c1 := client.NewClient(addr, port1)
		c2 := client.NewClient(addr, port2)

		// 注册第一个用户
		//registerReq1 := &request.RegisterUserReq{
		//	Email:    "WPS1",
		//	Password: "123",
		//}
		//registerResp1, err := c.Register(registerReq1)
		//if err != nil {
		//	fmt.Println("Register error:", err)
		//	return
		//}
		//fmt.Println("Register response:", *registerResp1)
		//
		////注册第二个用户
		//registerReq2 := &request.RegisterUserReq{
		//	Email:    "WPS2",
		//	Password: "123",
		//}
		//registerResp2, err := c.Register(registerReq2)
		//if err != nil {
		//	fmt.Println("Register error:", err)
		//	return
		//}
		//fmt.Println("Register response:", *registerResp2)

		// 登录用户1
		loginReq1 := &request.LoginUserReq{
			Email:    "WPS1",
			Password: "123",
		}
		loginResp1, err := c1.Login(loginReq1)
		if err != nil {
			fmt.Println("Login error:", err)
			return
		}
		fmt.Println("Login response:", *loginResp1)

		// 登录用户2
		loginReq2 := &request.LoginUserReq{
			Email:    "WPS1",
			Password: "123",
		}
		loginResp2, err := c1.Login(loginReq2)
		if err != nil {
			fmt.Println("Login error:", err)
			return
		}
		fmt.Println("Login response:", *loginResp2)

		// 将 token 塞到 cookie 里面
		token := loginResp1.Data.Token
		cookie := &http.Cookie{
			Name:    "wps_id",
			Value:   token,
			Expires: time.Now().Add(24 * time.Hour),
		}
		var Cookies []*http.Cookie
		Cookies = append(Cookies, cookie)

		// 创建短链接请求
		createReq := &request.CreateShortLinkReq{
			OriginalURL: "http://www.baidu.com",
			Title:       "example",
			ExpiresAt:   time.Now().Add(7 * 24 * time.Hour).UnixMilli(), // 过期时间设置为 7 天后
		}

		createResp, err := c2.CreateShortLink(createReq, Cookies...)
		if err != nil {
			fmt.Println("CreateShortLink error:", err)
			return
		}
		fmt.Println("CreateShortLink response:", *createResp)

		// Share short link
		shareReq := &request.ShareShortLinkReq{
			LinkId:  createResp.Data.LinkID,
			UserIds: []string{"WPS2"},
		}
		shareResp, err := c2.ShareShortLink(shareReq, Cookies...)
		if err != nil {
			fmt.Println("ShareShortLink error:", err)
			return
		}
		fmt.Println("ShareShortLink response:", *shareResp)

		// Update short link
		updateReq := &request.UpdateShortLinkReq{
			LinkId:      createResp.Data.LinkID,
			Title:       "hello",
			OriginalUrl: createResp.Data.OriginalURL,
			ExpiredAt:   createResp.Data.ExpiredAt,
		}
		updateResp, err := c2.UpdateShortLink(updateReq, Cookies...)
		if err != nil {
			fmt.Println("UpdateShortLink error:", err)
			return
		}
		fmt.Println("UpdateShortLink response:", *updateResp)

		// List short links
		listReq := &request.ListShortLinksReq{
			Page:     1,
			PageSize: 10,
		}
		listResp, err := c2.ListShortLinks(listReq, Cookies...)
		if err != nil {
			fmt.Println("ListShortLinks error:", err)
			return
		}
		fmt.Println("ListShortLinks response:", *listResp)

		// Search short links by link ID
		searchReq := &request.ListShortLinksByLinkIDReq{
			Keyword:  "example",
			Page:     1,
			PageSize: 10,
		}
		searchResp, err := c2.SearchShortLinksByLinkID(searchReq, Cookies...)
		if err != nil {
			fmt.Println("SearchShortLinksByLinkID error:", err)
			return
		}
		fmt.Println("SearchShortLinksByLinkID response:", *searchResp)

		// Summarize
		summarizeReq := &request.SummarizeReq{
			LinkId: createResp.Data.LinkID,
		}
		summarizeResp, err := c2.Summarize(summarizeReq, Cookies...)
		if err != nil {
			fmt.Println("Summarize error:", err)
			return
		}
		fmt.Println("Summarize response:", *summarizeResp)

		// Delete short link
		deleteReq := &request.DeleteShortLinkReq{
			LinkId: createResp.Data.LinkID,
		}
		deleteResp, err := c2.DeleteShortLink(deleteReq, Cookies...)
		if err != nil {
			fmt.Println("DeleteShortLink error:", err)
			return
		}
		fmt.Println("DeleteShortLink successful:", *deleteResp)
	},
}

func init() {
	testCmd.Flags().StringP("addr", "u", "127.0.0.1", "url")
	testCmd.Flags().Int("port1", 8080, "")
	testCmd.Flags().Int("port2", 8082, "")

}
