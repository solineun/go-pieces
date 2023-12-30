package tracker

import (
	"context"
	"fmt"
	"io"
	"mymodule/golang_book/idioms/ch12_context/values/tracker"
	"net/http"

	"github.com/google/uuid"
)

type guidKey int

const key guidKey = 1

func contextWithGUID(ctx context.Context, guid string) context.Context {
	return context.WithValue(ctx, key, guid)
}

func guidFromContext(ctx context.Context) (string, bool) {
	guid, ok := ctx.Value(key).(string)
	return guid, ok
}

func MiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if guid := r.Header.Get("X-GUID"); guid != "" {
			ctx = contextWithGUID(ctx, guid)
		} else {
			ctx = contextWithGUID(ctx, uuid.NewString())
		}
		r = r.WithContext(ctx)		
		h.ServeHTTP(w, r)
	})
}

type LoggerImpl struct{}

func (l LoggerImpl) Log(ctx context.Context, msg string) {
	if guid, ok := guidFromContext(ctx); ok {
		msg = fmt.Sprintf("GUID: %s - %s", guid, msg)
	}
	fmt.Println(msg)
}

func RequestWithGUID(r *http.Request) *http.Request {
	ctx := r.Context()
	if guid, ok := guidFromContext(ctx); ok {
		r.Header.Add("X-GUID", guid)
	}
	return r
}

type Logger interface{
	Log(context.Context, string)
}

type RequestDecorator func (*http.Request) *http.Request

type BusinessLogic struct {
	RequestDecorator
	Logger Logger
	Remote string
}

func (l BusinessLogic) businessLogic(ctx context.Context, user, data string) (string, error) {
	l.Logger.Log(ctx, "starting business logic for " + user + " with " + data)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, l.Remote+"?querry="+data, nil)
	if err != nil {
		l.Logger.Log(ctx, "error building remote request:" + err.Error())
		return "", err
	}
	req = l.RequestDecorator(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		l.Logger.Log(ctx, "error requesting remote:" + err.Error())
		return "", err
	}
	return resp.Status, nil
}

func Run() {
	bl := BusinessLogic{
		RequestDecorator: tracker.RequestWithGUID,
		Logger: tracker.LoggerImpl,
		Remote: "http://www.example.com/query",
	}
	// some business logic usage
}