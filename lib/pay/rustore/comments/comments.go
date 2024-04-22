package comments

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
	"time"

	"github.com/caiknife/mp3lister/lib/pay/rustore/client"
)

const (
	GetCommentsFullURL         = client.BaseURL + "/%s/comment"
	AnswerCommentURL           = client.BaseURL + "/%s/feedback?commentId=%s"
	RedactOrDeleteCommentURL   = client.BaseURL + "/%s/feedback/%s"
	GetCommentsFeedbackFullURL = client.BaseURL + "/%s/feedback"
	GetCommentsFullCSVURL      = client.BaseURL + "/%s/comment/export?from=%s&to=%s"
	GetRaitingsURL             = client.BaseURL + "/%s/comment/statistic"
)

type Comments struct {
	client      *client.Client
	packageName string
}

func New(c *client.Client, packageName string) *Comments {
	return &Comments{
		client:      c,
		packageName: packageName,
	}
}

// nolint: tagliatelle
type GetCommentsFullResponse struct {
	Code                string         `json:"code"`
	Message             *string        `json:"message"`
	GetCommentsFullBody []CommentsBody `json:"body"`
	Timestamp           time.Time      `json:"timestamp"`
}

// nolint: tagliatelle
type CommentsBody struct {
	PackageName    string `json:"packageName"`
	AppID          int    `json:"appId"`
	CommentID      int64  `json:"commentId"`
	UserName       string `json:"userName"`
	AppRating      int    `json:"appRating"`
	CommentStatus  string `json:"commentStatus"`
	CommentDate    string `json:"commentDate"`
	CommentText    string `json:"commentText"`
	LikeCounter    int    `json:"likeCounter"`
	DislikeCounter int    `json:"dislikeCounter"`
	UpdatedAt      string `json:"updatedAt"`
	AppVersionName string `json:"appVersionName"`
	Edited         bool   `json:"edited"`
}

func (c *Comments) GetCommentsFull(ctx context.Context, params neturl.Values) (GetCommentsFullResponse, error) {
	url := fmt.Sprintf(
		GetCommentsFullURL,
		c.packageName,
	)
	if len(params) > 0 {
		url += "?" + params.Encode()
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return GetCommentsFullResponse{}, err
	}

	response, err := c.client.Do(
		request,
		client.RequestOpts{},
	)
	if err != nil {
		return GetCommentsFullResponse{}, err
	}
	defer response.Body.Close()

	getCommentsFull := GetCommentsFullResponse{}
	err = json.NewDecoder(response.Body).Decode(&getCommentsFull)

	return getCommentsFull, err
}

type AnswerCommentResponse struct {
	Code      string            `json:"code"`
	Message   *string           `json:"message"`
	Body      AnswerCommentBody `json:"body"`
	Timestamp time.Time         `json:"timestamp"`
}

type AnswerCommentBody struct {
	ID int `json:"id"`
}

func (c *Comments) AnswerComment(ctx context.Context, commentID string,
	feedbackText string) (AnswerCommentResponse, error) {
	feedbackTextJSON, err := json.Marshal(FeedbackText{Message: feedbackText})
	if err != nil {

		return AnswerCommentResponse{}, err
	}

	url := fmt.Sprintf(
		AnswerCommentURL,
		c.packageName,
		commentID,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bytes.NewBuffer(feedbackTextJSON),
	)
	if err != nil {

		return AnswerCommentResponse{}, err
	}

	response, err := c.client.Do(
		request,
		client.RequestOpts{},
	)
	if err != nil {

		return AnswerCommentResponse{}, err
	}

	defer response.Body.Close()

	answerComment := AnswerCommentResponse{}
	err = json.NewDecoder(response.Body).Decode(&answerComment)

	return answerComment, err
}

type FeedbackText struct {
	Message string `json:"message"`
}

type RedactFeedbackResponse struct {
	Code      string             `json:"code"`
	Message   *string            `json:"message"`
	Body      RedactFeedbackBody `json:"body"`
	Timestamp time.Time          `json:"timestamp"`
}

type RedactFeedbackBody struct {
	ID int64 `json:"id"`
}

func (c *Comments) RedactFeedback(ctx context.Context, feedbackID string, feedbackText string) (RedactFeedbackResponse, error) {
	feedbackTextJSON, err := json.Marshal(FeedbackText{Message: feedbackText})
	if err != nil {
		return RedactFeedbackResponse{}, err
	}

	url := fmt.Sprintf(
		RedactOrDeleteCommentURL,
		c.packageName,
		feedbackID,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bytes.NewBuffer(feedbackTextJSON),
	)
	if err != nil {

		return RedactFeedbackResponse{}, err
	}

	response, err := c.client.Do(request, client.RequestOpts{})
	if err != nil {

		return RedactFeedbackResponse{}, err
	}

	defer response.Body.Close()

	redactFeedback := RedactFeedbackResponse{}
	err = json.NewDecoder(response.Body).Decode(&redactFeedback)

	return redactFeedback, err
}

type DeleteFeedbackResponse struct {
	Code      string             `json:"code"`
	Message   *string            `json:"message"`
	Body      DeleteFeedbackBody `json:"body"`
	Timestamp time.Time          `json:"timestamp"`
}

type DeleteFeedbackBody struct {
	ID int64 `json:"id"`
}

func (c *Comments) DeleteFeedback(ctx context.Context, feedbackID string) (DeleteFeedbackResponse, error) {
	url := fmt.Sprintf(
		RedactOrDeleteCommentURL,
		c.packageName,
		feedbackID,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		url,
		nil,
	)
	if err != nil {

		return DeleteFeedbackResponse{}, err
	}

	response, err := c.client.Do(request, client.RequestOpts{})
	if err != nil {

		return DeleteFeedbackResponse{}, err
	}

	defer response.Body.Close()

	deleteFeedback := DeleteFeedbackResponse{}
	err = json.NewDecoder(response.Body).Decode(&deleteFeedback)

	return deleteFeedback, err
}

// nolint: tagliatelle
type GetRaitingsResponse struct {
	Code            string       `json:"code"`
	Message         *string      `json:"message"`
	GetRaitingsBody RaitingsBody `json:"body"`
	Timestamp       time.Time    `json:"timestamp"`
}

// nolint: tagliatelle
type RaitingsBody struct {
	GetRaitingsBodyRatings RaitingsBodyRatings `json:"ratings"`
	AverageUserRating      float64             `json:"averageUserRating"`
	TotalRatings           int                 `json:"totalRatings"`
	TotalResponses         int                 `json:"totalResponses"`
	RatingsNoComments      int                 `json:"ratingsNoComments"`
}

// nolint: tagliatelle
type RaitingsBodyRatings struct {
	AmountFive  int `json:"amountFive"`
	AmountFour  int `json:"amountFour"`
	AmountThree int `json:"amountThree"`
	AmountTwo   int `json:"amountTwo"`
	AmountOne   int `json:"amountOne"`
}

func (c *Comments) GetRaitings(ctx context.Context) (GetRaitingsResponse, error) {
	url := fmt.Sprintf(
		GetRaitingsURL,
		c.packageName,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {

		return GetRaitingsResponse{}, err
	}

	response, err := c.client.Do(request, client.RequestOpts{})
	if err != nil {

		return GetRaitingsResponse{}, err
	}

	defer response.Body.Close()

	raitings := GetRaitingsResponse{}
	err = json.NewDecoder(response.Body).Decode(&raitings)

	return raitings, err
}

type GetCommentFeedbackResponse struct {
	Code      string         `json:"code"`
	Message   *string        `json:"message"`
	Body      []FeedbackBody `json:"body"`
	Timestamp time.Time      `json:"timestamp"`
}

type FeedbackBody struct {
	ID        string    `json:"id"`
	CommentID string    `json:"commentId"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	Date      time.Time `json:"date"`
}

func (c *Comments) GetCommentFeedbackStatesFull(ctx context.Context, params neturl.Values) (GetCommentFeedbackResponse, error) {
	url := fmt.Sprintf(
		GetCommentsFeedbackFullURL,
		c.packageName,
	)
	if len(params) > 0 {
		url += "?" + params.Encode()
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {

		return GetCommentFeedbackResponse{}, err
	}

	response, err := c.client.Do(request, client.RequestOpts{})
	if err != nil {

		return GetCommentFeedbackResponse{}, err
	}

	defer response.Body.Close()

	commentFeedback := GetCommentFeedbackResponse{}
	err = json.NewDecoder(response.Body).Decode(&commentFeedback)

	return commentFeedback, err
}

func (c *Comments) GetCommentsCSVFile(ctx context.Context, dateFrom, dateTo string) ([]byte, error) {
	url := fmt.Sprintf(
		GetCommentsFullCSVURL,
		c.packageName,
		dateFrom,
		dateTo,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil)

	if err != nil {

		return []byte(""), err
	}

	response, err := c.client.Do(request, client.RequestOpts{})
	if err != nil {

		return []byte(""), err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	return body, err
}
