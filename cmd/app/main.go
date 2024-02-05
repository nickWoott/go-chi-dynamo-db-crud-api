package main

//is this how we import stuff I am not sure, need to look at this
import (
	"fmt"
	"log"
	"net/http"

	"github.com/akhil/dynamo-go=yt/utils/logger"
)

// importing the config file
func main() {
	configs := Config.GetConfig()

	// this is from the repository file (package)
	connection := instance.getConnection()
	repository := adapter.NewAdapter(connection)

	loggger.INFO("waiting for the service to start....", nil)

	errors := Migrate(connection)

	if len(errors) > 0 {
		for _, err := range errors {
			logger.PANIC("Error on migration....", err)
		}
	}
	logger.PANIC("", checkTables(connection))

	port:= fmt.SprintF(":%v", configs.Port)

	router := routes.NewRouter().SetRouters(repository)

	logger.INFO("Service is running on port", port)

	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}

func Migrate(connection *dynamodb.DynamoDB) []error{
	var errors {} [] error
	callMigrateAnAppendError(&errors, connection, &RulesProduct.Rules{})
	return errors
}


func callMigrateAndAppendError(errors *[]error, connection *dynamodb.DynamoDB, rule rules.Interface){
	err:= rule.Migrate(connection)
	if err != nil{
		*errors = append (*errors, err)
	}
}



func checkTables(connection *dynamodb.DynamoDB) error {
response, err := connection.ListTables(&dynamodb.DynamoDB)
//either tablename found or not found
if response != nil {
	if len(response.TableNames) ==0 {
		logger.INFO("Tables not found", nil)
	}
	for _, tableName := range response.TableName {
		logger.INFO("Table found", *tablename)
	}
}
return err
}
