package instance

// import (
//
//	"github.com/aws/aws-sdk-go/aws/session"
//	"github.com/aws/aws-sdk-go/service/dynamodb"
//
// )
func GetConnection() *dynamodb.DynamoDB {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.sharedConfigEnable,
	}))

	return dynamodb.New(sess)
}

//return