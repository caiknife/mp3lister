package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/caiknife/mp3lister/lib/pay/rustore/client"
	"github.com/caiknife/mp3lister/lib/pay/rustore/comments"
)

var (
	urlValues    = url.Values{}
	dateFrom     = "2024-01-01"
	dateTo       = "2024-02-06"
	feedbackText = "Спасибо вам большое!"
	feedbackID   = "123"         // id ответа, который вы оставили на отзыв
	commentID    = "321"         // id отзыва, на который вы хотите ответить
	packageName  = "app.example" // packageName вашего приложения
	file         = flag.String(
		"file",
		"",
		"file with options: privateKey, privateKeyID, companyID",
	)
	options Options
)

type Options struct {
	PrivateKey   string `json:"privateKey"`
	PrivateKeyID string `json:"privateKeyID"`
	CompanyID    string `json:"companyID"`
}

func main() {

	flag.Parse()
	data, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(data, &options); err != nil {
		panic(err)
	}

	// создали нового клиента для работы с отзывами

	cl := client.New(
		options.PrivateKeyID,
		options.PrivateKey,
		options.CompanyID,
	)
	err = cl.Auth()
	if err != nil {
		log.Fatal(err)
	}

	coms := comments.New(cl, packageName)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		client.TimeOutSeconds*time.Second,
	)

	defer cancel()

	// получение всех комментариев:

	comsGetCommentsFull, err := coms.GetCommentsFull(
		ctx,
		urlValues,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nкомменты: %+v\n", comsGetCommentsFull)

	// получение комментов по id:

	urlValues.Set("id", commentID)
	comsGetCommentsByID, err := coms.GetCommentsFull(
		ctx,
		urlValues,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nкомменты по id: %+v\n", comsGetCommentsByID)

	// получение комментов по странице и количеству отзывов:

	urlValues.Set("page", "0")
	urlValues.Add("size", "1")
	comsGetCommentsByPageAndSize, err := coms.GetCommentsFull(
		ctx,
		urlValues,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nкомменты по page and size: %+v\n", comsGetCommentsByPageAndSize)

	// получить все отзывы в csv:

	comsGetCommentsCSV, err := coms.GetCommentsCSVFile(
		ctx,
		dateFrom,
		dateTo,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nвсе отзывы в CSV: %+v\n", string(comsGetCommentsCSV))

	// получить все статусы ответов на отзывы

	urlValues = url.Values{}
	comsGetCommentFeedbackStatesFull, err := coms.GetCommentFeedbackStatesFull(
		ctx,
		urlValues,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nответы на отзывы: %+v\n", comsGetCommentFeedbackStatesFull)

	// получить конкретный статус ответа на отзыв по id:

	urlValues = url.Values{}
	urlValues.Set("id", feedbackID)
	comsGetCommentFeedbackStateByID, err := coms.GetCommentFeedbackStatesFull(
		ctx,
		urlValues,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nответ на конкретный отзыв: %+v\n", comsGetCommentFeedbackStateByID)

	// получить статусы ответов на отзывы по page and size:

	urlValues = url.Values{}
	urlValues.Set("page", "0")
	urlValues.Add("size", "2")
	comsGetCommentFeedbackStateByPageAndSize, err := coms.GetCommentFeedbackStatesFull(
		ctx,
		urlValues,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nответы на отзывы по page and size: %+v\n", comsGetCommentFeedbackStateByPageAndSize)

	// ответить на отзыв:

	answerComment, err := coms.AnswerComment(
		ctx,
		commentID,
		feedbackText,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nответили на отзыв: %+v\n", answerComment)

	// редактировать ответ на отзыв:

	redactFeedback, err := coms.RedactFeedback(
		ctx,
		feedbackID,
		feedbackText,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nотредактировали отзыв: %+v\n", redactFeedback)

	// удалить ответ на отзыв:

	urlValues = url.Values{}
	deleteFeedback, err := coms.DeleteFeedback(
		ctx,
		feedbackID,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nудалили ответ на отзыв: %+v\n", deleteFeedback)

	// получить рейтинг приложения:

	comsGetRaitings, err := coms.GetRaitings(
		ctx,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nрейтинг: %+v\n", comsGetRaitings)

}
