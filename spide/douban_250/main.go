package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MovieInfo struct {
	Index             int    `json:"index"`
	Title             string `json:"title"`
	Rating_num        string `json:"rating_num"`
	Rating_people_num string `json:"rating_people_num"`
}

func main() {
	var (
		client     = InitMogDB()
		collection *mongo.Collection
	)
	collection = client.Database("my_db").Collection("douban_250")

	for i := 0; i < 10; i++ {
		fmt.Printf("正在爬取第 %d 页", i)
		Spider(strconv.Itoa(i*25), collection)
	}
}

func InitMogDB() *mongo.Client {
	var (
		client *mongo.Client
		err    error
	)
	//1.建立连接
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return client
}

func Spider(pageNumber string, collection *mongo.Collection) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://movie.douban.com/top250?start="+pageNumber, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://movie.douban.com/chart")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	doc.Find("#wrapper #content .article .grid_view .item ").Each(func(i int, s *goquery.Selection) {
		pageNumberInt, _ := strconv.Atoi(pageNumber)
		index := pageNumberInt + i + 1
		title := s.Find(".hd .title:nth-child(1)").Text()
		rating_num := s.Find(".bd .star .rating_num").Text()
		rating_people_num := s.Find(".bd .star span:nth-child(4)").Text()

		movieInfo := MovieInfo{
			Index:             index,
			Title:             title,
			Rating_num:        rating_num,
			Rating_people_num: rating_people_num,
		}
		bsonMovieInfo := bson.M{"index": movieInfo.Index, "title": movieInfo.Title, "rating_num": movieInfo.Rating_num, "rating_people_num": movieInfo.Rating_people_num}
		_, err := collection.InsertOne(context.TODO(), bsonMovieInfo)
		if err != nil {
			fmt.Print(err)
			return
		}
		// fmt.Println(result.InsertedID)
	})
}
