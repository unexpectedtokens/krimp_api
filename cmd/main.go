package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"

	"google.golang.org/api/option"
)

func initateFirebaseAdminSDK() (*firebase.App, error) {
	opt := option.WithCredentialsFile("./krimp-43829-firebase-adminsdk-gc8wp-705d896b03.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return app, nil
}

func main() {
	app, err := initateFirebaseAdminSDK()
	if err != nil {
		panic(err)
	}
	client, _ := app.Auth(context.Background())
	client.VerifyIDToken(context.Background(), "asdasdasdw1231aqsd")
}
