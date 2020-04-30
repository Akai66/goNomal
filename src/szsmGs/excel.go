package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"strconv"
	"time"
)

func main() {
	//读取php获取的工时report数据
	reportMap := getReportData()
	f, err := excelize.OpenFile("./gshz.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	for name, report := range reportMap {
		for _, detail := range report {
			for taskName, data := range detail {
				if taskName == "面向大项目部销售的综合支持" || taskName == "HBH03" || taskName == "HUN01" || taskName == "HBH161" || taskName == "HD127" {
					continue
				}
				for date, item := range data {
					//定位需要修改的cell的列坐标
					tm, _ := time.Parse("2006/01/02", date)
					formatTm := int(tm.Unix()/(24*3600) + 25569)
					axisOne, err := f.SearchSheet(name, strconv.Itoa(formatTm))
					if err != nil {
						fmt.Println(err.Error())
					}
					if len(axisOne) <= 0 {
						fmt.Println("获取纵坐标失败:", name, taskName, date, item["content"], item["spentTime"])
					}
					col, _, _ := excelize.SplitCellName(axisOne[0])
					//定位需要修改的cell的行坐标
					axisTwo, err := f.SearchSheet(name, taskName)
					if err != nil {
						fmt.Println(err.Error())
					}
					if len(axisTwo) <= 0 {
						fmt.Println("获取行坐标失败:", name, taskName, date, item["content"], item["spentTime"])
					}
					_, row, _ := excelize.SplitCellName(axisTwo[0])
					//记录日志
					fmt.Println(name, taskName, date, item["content"], item["spentTime"], col, row)
					//填写工时
					axis, _ := excelize.JoinCellName(col, row)
					f.SetCellValue(name, axis, item["spentTime"])
					if item["content"] != "" {
						comment := fmt.Sprintf(`{"author":"作者: ","text":"%v"}`, item["content"])
						f.AddComment(name, axis, comment)
					}
				}
			}
		}
	}

	if err = f.Save(); err != nil {
		println(err.Error())
	}
}

func getReportData() map[string]map[string]map[string]map[string]map[string]interface{} {
	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		fmt.Println("read file error:", err.Error())
	}
	reportMap := make(map[string]map[string]map[string]map[string]map[string]interface{})
	json.Unmarshal(data, &reportMap)
	if err != nil {
		fmt.Println(err)
	}
	return reportMap
}
