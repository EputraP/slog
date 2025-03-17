package main

import (
	"fmt"
	"io"
	"log/slog"
	"math/rand"
	"os"
)

type User struct {
	ID       int
	Username string
	Password string
}

func (u *User) LogValue() slog.Value {
	return slog.StringValue(fmt.Sprintf(
		"id=%d, username=%s, password=%s",
		u.ID, u.Username, u.Password,
	))
}

func main() {

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	handlerOpts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	logger := slog.New(slog.NewJSONHandler(multiWriter, handlerOpts))
	slog.SetDefault(logger)

	logger.Debug("This is a debug message", "key1", "value1", "key2", "value2")
	logger.Info("This is an info message", slog.String("key1", "value1"), slog.Int("key2", 2))
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")

	//slog group
	userGrop := slog.Group("users",
		"id", rand.Intn(1000),
		"username", "test")

	reqGroup := slog.Group("req",
		"method", "GET")

	logger.Info("Request", userGrop, reqGroup)

	//slog with

	requestLogger := logger.With(reqGroup)
	requestLogger.Info("Request")

	user := &User{
		ID:       1,
		Username: "test",
		Password: "password",
	}

	requestLogger.Info("User", "user data", user)

}
