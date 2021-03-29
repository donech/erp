package common

import (
	"context"
	"flag"
	"fmt"
	"github.com/donech/tool/xlog"
	"testing"
)

func init()  {
	testing.Init()
	flag.Parse()
	xlog.New(xlog.Config{
		ServiceName:     "testing",
		Stdout:          true,
	})
}

func TestGenToken(t *testing.T) {
	field := CustomField{
		UserID: 1,
		Name:   "solar",
	}
	token, err := GenToken(context.Background(), field)
	if err != nil {
		t.Fatal("GenToken error", err)
	}
	fmt.Println("token: ", token)
}

func TestValidToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIk5hbWUiOiJzb2xhciIsImV4cCI6MTYxNzAwMjMwMywibmJmIjoxNjE2OTk4NzAzfQ.otZaDG0MvNl8d4UG33xXQXQUWGW-gjmLkiyjjhasl6w"
	result, err := ValidToken(context.Background(), token)
	if err != nil {
		t.Fatal("ValidToken error", err)
	}
	fmt.Println(result)
}

