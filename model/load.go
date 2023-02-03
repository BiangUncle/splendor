package model

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

var CsvFilePath = "../csv/"

// ReadCsv 读取 csv 文件
func ReadCsv(filepath string) ([]string, [][]string) {
	//打开文件(只读模式)，创建io.read接口实例
	opencast, err := os.Open(filepath)
	if err != nil {
		log.Println("csv文件打开失败！")
	}
	defer opencast.Close()

	//创建csv读取接口实例
	ReadCsv := csv.NewReader(opencast)

	//获取一行内容，一般为第一行内容
	read, _ := ReadCsv.Read() //返回切片类型：[chen  hai wei]
	//log.Println(read)

	//读取所有内容
	ReadAll, _ := ReadCsv.ReadAll() //返回切片类型：[[s s ds] [a a a]]
	//log.Println(ReadAll)

	/*
	  说明：
	   1、读取csv文件返回的内容为切片类型，可以通过遍历的方式使用或Slicer[0]方式获取具体的值。
	   2、同一个函数或线程内，两次调用Read()方法时，第二次调用时得到的值为每二行数据，依此类推。
	   3、大文件时使用逐行读取，小文件直接读取所有然后遍历，两者应用场景不一样，需要注意。
	*/

	return read, ReadAll
}

// ToInt 字符串转数字
func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return i
}

// LoadDefaultDevelopmentCard 加载默认发展卡
func LoadDefaultDevelopmentCard() error {
	_, rows := ReadCsv(CsvFilePath + "dev_card.csv")
	defaultDevelopmentCardStacks = &DevelopmentCardStacks{}

	for _, row := range rows {

		card := &DevelopmentCard{
			Idx:       ToInt(row[0]),
			Level:     ToInt(row[1]),
			BonusType: ToInt(row[2]),
			Prestige:  ToInt(row[3]),
			Acquires: TokenStack{
				ToInt(row[4]),
				ToInt(row[5]),
				ToInt(row[6]),
				ToInt(row[7]),
				ToInt(row[8]),
				0,
			},
		}

		DevelopmentCardMap[card.Idx] = card

		switch card.Level {
		case DevelopmentCardLevelBottom:
			defaultDevelopmentCardStacks.BottomStack = append(defaultDevelopmentCardStacks.BottomStack, card)
		case DevelopmentCardLevelMiddle:
			defaultDevelopmentCardStacks.MiddleStack = append(defaultDevelopmentCardStacks.MiddleStack, card)
		case DevelopmentCardLevelTop:
			defaultDevelopmentCardStacks.TopStack = append(defaultDevelopmentCardStacks.TopStack, card)
		}
	}

	return nil
}

// LoadDefaultNobleTiles 加载默认贵族卡
func LoadDefaultNobleTiles() error {
	_, rows := ReadCsv(CsvFilePath + "noble_tile.csv")

	for _, row := range rows {

		noble := &NobleTile{
			Idx:      ToInt(row[0]),
			Prestige: ToInt(row[1]),
			Acquires: TokenStack{
				ToInt(row[2]),
				ToInt(row[3]),
				ToInt(row[4]),
				ToInt(row[5]),
				ToInt(row[6]),
				0,
			},
		}

		defaultNobleTilesStack = append(defaultNobleTilesStack, noble)
		NobleTilesMap[noble.Idx] = noble
	}

	return nil
}
