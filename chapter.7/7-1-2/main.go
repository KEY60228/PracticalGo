package main

import "context"

type T interface {
	Do(context.Context)
	CreateTweet(context.Context, int64) *T
	UpdateTweet(context.Context, int64) *T
	CreateReply(context.Context, int64, int64) *T
	Retweet(context.Context, int64, int64) *T
}

type Notifier struct{}

// 以下自動生成

func (n Notifier) Do(_ context.Context) {
	panic("not implemented") // TODO: Implement
}

func (n Notifier) CreateTweet(_ context.Context, _ int64) *T {
	panic("not implemented") // TODO: Implement
}

func (n Notifier) UpdateTweet(_ context.Context, _ int64) *T {
	panic("not implemented") // TODO: Implement
}

func (n Notifier) CreateReply(_ context.Context, _ int64, _ int64) *T {
	panic("not implemented") // TODO: Implement
}

func (n Notifier) Retweet(_ context.Context, _ int64, _ int64) *T {
	panic("not implemented") // TODO: Implement
}
