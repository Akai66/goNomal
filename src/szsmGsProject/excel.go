package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	DOMAIN_ID         string = "ad585628-76bf-49eb-a433-c85613981461"
	URL_GETUSERLIST   string = "http://szzb-api.shengxunwei.com/api/User/GetUserList"
	URL_GETREPORTLIST string = "http://szzb-api.shengxunwei.com/api/WeeklyReport/GetWeeklyReportList"
	URL_GETREPORT     string = "http://szzb-api.shengxunwei.com/api/WeeklyReport/GetWeeklyReport?id="
	FILE_PATH         string = "./gshz.xlsx"
)

var (
	passUser = map[string]int{
		"姜锋":  1,
		"李晗蕾": 1,
		"刘俊杰": 1,
		"胡云":  1,
		"鲍晓宇": 1,
		"梁铮":  1,
		"王祖欣": 1,
		"黄旭伟": 1,
		"管理员": 1,
		"吴海波": 1,
		"李文东": 1,
	}
	token     string
	startDate string
)

func main() {
	//校验参数 token,startdate,例子：4942733d-5511-4a79-a4c1-a1098657ead1 2020/05/04
	if len(os.Args[1:]) != 2 {
		fmt.Println("参数错误")
		os.Exit(1)
	}
	token, startDate = os.Args[1], os.Args[2]

	//获取用户列表
	userList := getUserInfo()

	//获取工时信息
	reportData := getReportData(userList)

	//修改excel数据
	saveExcel(reportData)
}

/**
获取用户列表信息
*/
func getUserInfo() (userList []map[string]interface{}) {
	var postData = map[string]interface{}{
		"OrderBy": "",
		"PagingInfo": map[string]int{
			"currentPage": 1,
			"pageSize":    50,
			"totalPage":   0,
			"totalCount":  0,
		},
		"Parameters": map[string]string{
			"domainId":           DOMAIN_ID,
			"organizationId":     "",
			"organizationIdPath": "",
			"keyword":            "",
		},
	}
	jsonData, err := json.Marshal(postData)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	header := map[string]string{
		"Content-Type":   "application/json;charset=UTF-8",
		"Content-Length": string(len(jsonData)),
	}
	respBytes := postReq(jsonData, header, URL_GETUSERLIST)
	result := make(map[string]map[string][]map[string]interface{})
	json.Unmarshal(respBytes, &result)
	userList = result["data"]["data"]
	return userList
}

func getReportData(userList []map[string]interface{}) (reportData map[string][]map[string]interface{}) {
	reportData = make(map[string][]map[string]interface{})
	dateArr := strings.Split(startDate, "/")
	startMon, _ := strconv.Atoi(dateArr[1])
	startYear, _ := strconv.Atoi(dateArr[0])
	endYear, endMon := startYear, startMon
	for _, user := range userList {
		userId := user["id"].(string)
		userName := user["name"].(string)
		if _, ok := passUser[userName]; ok {
			continue
		}
		var postData = map[string]interface{}{
			"OrderBy": "",
			"PagingInfo": map[string]int{
				"currentPage": 1,
				"pageSize":    10,
				"totalPage":   4,
				"totalCount":  0,
			},
			"Parameters": map[string]interface{}{
				"startYear":  startYear,
				"endYear":    endYear,
				"startMonth": startMon,
				"endMonth":   endMon,
				"userId":     userId,
				"userName":   userName,
			},
		}
		jsonData, err := json.Marshal(postData)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		header := map[string]string{
			"Content-Type":   "application/json;charset=UTF-8",
			"Content-Length": string(len(jsonData)),
		}
		respBytes := postReq(jsonData, header, URL_GETREPORTLIST)
		result := make(map[string]map[string][]map[string]interface{})
		json.Unmarshal(respBytes, &result)
		reportList := result["data"]["data"]
		if len(reportList) <= 0 {
			fmt.Printf("获取工时列表失败,name:%s\n", userName)
			continue
		}
		var reqId string
		for _, report := range reportList {
			if strings.Index(report["monday"].(string), startDate) != -1 {
				reqId = report["id"].(string)
				break
			}
		}
		if reqId == "" {
			fmt.Printf("获取本周工时report失败,name:%s\n", userName)
			continue
		}
		url := URL_GETREPORT + reqId
		header = map[string]string{
			"Content-Length": "0",
		}
		reportResBytes := postReq([]byte{}, header, url)
		reportResult := make(map[string]map[string][]map[string]interface{})
		json.Unmarshal(reportResBytes, &reportResult)
		if len(reportResult["data"]["weeklyReportItem_BusinessOpportunity"]) <= 0 && len(reportResult["data"]["weeklyReportItem_Task"]) <= 0 {
			fmt.Printf("获取工时报告失败,name:%s\n", userName)
			continue
		}
		for _, data := range reportResult["data"]["weeklyReportItem_BusinessOpportunity"] {
			if data["spentTime"] != nil && data["spentTime"].(float64) > 0 {
				dateArr := strings.Split(data["date"].(string), " ")
				date := dateArr[0]
				taskData := map[string]interface{}{
					"task":      data["businessOpportunitySerialNumber"],
					"date":      date,
					"spentTime": data["spentTime"],
					"content":   data["content"],
				}
				reportData[userName] = append(reportData[userName], taskData)
			}
		}
		for _, data := range reportResult["data"]["weeklyReportItem_Task"] {
			if data["spentTime"] != nil && data["spentTime"].(float64) > 0 {
				dateArr := strings.Split(data["date"].(string), " ")
				date := dateArr[0]
				taskData := map[string]interface{}{
					"task":      data["taskName"],
					"date":      date,
					"spentTime": data["spentTime"],
					"content":   data["content"],
				}
				reportData[userName] = append(reportData[userName], taskData)
			}
		}
	}
	return reportData
}

/**
同步excel文件
*/
func saveExcel(reportData map[string][]map[string]interface{}) {
	f, err := excelize.OpenFile(FILE_PATH)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for name, datas := range reportData {
		for _, data := range datas {
			//定位需要修改的cell的列坐标
			date := data["date"].(string)
			spendTime := data["spentTime"].(float64)
			task := data["task"].(string)
			content := data["content"].(string)
			tm, _ := time.Parse("2006/01/02", date)
			formatTm := int(tm.Unix()/(24*3600) + 25569)
			axisOne, err := f.SearchSheet(name, strconv.Itoa(formatTm))
			if err != nil {
				fmt.Println(err.Error())
			}
			if len(axisOne) <= 0 {
				fmt.Println("获取纵坐标失败:", name, task, date)
				continue
			}
			col, _, _ := excelize.SplitCellName(axisOne[0])
			//定位需要修改的cell的行坐标
			axisTwo, err := f.SearchSheet(name, task)
			if err != nil {
				fmt.Println(err.Error())
			}
			if len(axisTwo) <= 0 {
				fmt.Println("获取行坐标失败:", name, task, date)
				continue
			}
			_, row, _ := excelize.SplitCellName(axisTwo[0])
			//记录日志
			fmt.Println(name, task, date, content, spendTime, col, row)
			//填写工时
			axis, _ := excelize.JoinCellName(col, row)
			f.SetCellValue(name, axis, spendTime)
			if content != "" {
				comment := fmt.Sprintf(`{"author":"作者: ","text":"%s"}`, content)
				f.AddComment(name, axis, comment)
			}
		}
	}
	if err = f.Save(); err != nil {
		println(err.Error())
	}
}

/**
发送post请求
*/
func postReq(postData []byte, header map[string]string, url string) []byte {
	reader := bytes.NewReader(postData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.113 Safari/537.36")
	request.Header.Set("Host", "szzb-api.shengxunwei.com")
	request.Header.Set("Origin", "http://szzb.shengxunwei.com")
	request.Header.Set("Referer", "http://szzb.shengxunwei.com/")
	request.Header.Set("token", token)
	for k, v := range header {
		request.Header.Set(k, v)
	}
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	resp.Body.Close()
	return respBytes
}
