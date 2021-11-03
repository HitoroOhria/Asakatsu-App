package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/asakatsu-app/domain/api_io"
	"example.com/asakatsu-app/infrastructure"
	"example.com/asakatsu-app/injector"
)

const location = "Asia/Tokyo"

func init() {
	log.Print("run: handler.init()")

	initTimeLocation()
}

// initTimeLocation は、time パッケージの TimeZone を JST で初期化します。
func initTimeLocation() {
	log.Print("run: handler.initTimeLocation()")

	loc, err := time.LoadLocation(location)
	if err != nil {
		log.Printf("time.LoadLocation fialed.(err=%+v)", err)
		loc = time.FixedZone(location, 9*60*60)
	}

	time.Local = loc
}

// FetchActivitiesFromSlackBatch は、FetchActivitiesFromSlackBatch を実行します。
func FetchActivitiesFromSlackBatch(ctx context.Context) error {
	log.Print("run: handler.FetchActivitiesFromSlackBatch()")

	firebaseHander := injector.InjectFirebaseHandler(ctx)
	firestoreDBConn := infrastructure.GetFirestoreDBConn(ctx, firebaseHander)
	defer firestoreDBConn.Close()

	usecase := injector.InjectFetchActivitiesFromSlackBatchUsecase(ctx)
	if err := usecase.Exec(); err != nil {
		fmt.Printf(err.Error())
		return err
	}

	return nil
}

// GetActivitiesFromSlackUidFunction は、GetActivitiesFromSlackUidFunction を実行します。
func GetActivitiesFromSlackUidFunction(
	input *api_io.GetActivitiesFromSlackUidInput,
) *api_io.GetActivitiesFromSlackUidOutput {
	log.Print("run: handler.GetActivitiesFromSlackUidUsecase()")

	ctx := context.Background()
	firebaseHander := injector.InjectFirebaseHandler(ctx)
	firestoreDBConn := infrastructure.GetFirestoreDBConn(ctx, firebaseHander)
	defer firestoreDBConn.Close()

	usecase := injector.InjectGetActivitiesFromSlackUidUsecase(ctx)
	activityFields, err := usecase.Exec(input.SlackUID)
	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}

	return api_io.NewGetActivitiesFromSlackUidOutput(activityFields)
}
