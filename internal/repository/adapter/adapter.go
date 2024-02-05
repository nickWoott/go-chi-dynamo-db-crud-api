package adapter

// import(
// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/service/dynamodb"
"github.com/aws/aws-sdk-go/dynamodb/dynamodbattribute"
"github.com/aws/aws-sdk-go/dynamodb/expresison"

// )

type Database struct {
	connection *dynamodb.DynamoDB, 
	logMode: bool
}

type Interface interface{
Health() bool
FindAll(condition expression.Expression, tablename string)(response *dynamodb.ScanOutput, err error)
FindOne(condition map[string]interface{}, tableName string)
CreateOrUpdate(entity interface{}, tablename string)(response *dynamodb.PutItemOutput, err, error)
Delete(condition map[string]interface{}, tablename string)(respose *dynamodb.DeleteItemOutput, erro error)





}

func NewAdapter(con *dynamodb.Dynamodb) Interface {
return &Database {
	connection: con, 
	logMode: false,
}
}
//the database functions
//find the documntation for these and look through it
// check check, lets have a look here at what is happening
// do you know what it is that is happening 
func (db *Database) Health() bool {
_, err := db,.connection.ListTables(&dynamodb.ListTablesInput{})
return err == nil 
}

// what the hell is this then Akhil?
//he said you need to actually need yto read the documentation
// why has he done routes that are not actually in the controller,
func(db *Database)FindAll {condition expression.Expression, tableName string)(response *dynamodb.ScanOutput, err error )

	input = &dynamodb.ScanInput{
		ExpressionAttributeNames: condition.Name(), 
		ExpressionAttributeValyes: condition.Values(),
		FilterExpression: condition,Filter(), 
		ProjecftionExpression: condition.Projection(),
		Tablename : aws.String(tablename),
	}
}


func(db *Database)FindOne(condition map[string]interface{}, tableName)(response *dynamodb.GetItemOutput, err error  {

if err != nil {
	return nil, err
}
	
	conditionParsed, errr := dynamodbattribute.MarshalMap(condition)
input := &dynamodb.GetItemInput {
	TableName : aws.String(tableName), 
	Key : conditionParsed, 
}
return db.connection.GetItem(input)
}


func(db *Database)CreateOrUpdate (entity interface(), tableName string) (reponse *dynamodb.PutItemOutput, err error) {

entityParsed, err := dynamodbattribute.MarshalMap(entity)

if err != nil {
	return nil err 
}

	input := &dynamodb.PutItemInput{
		Item: entityParsed,
		TableName: aws.String(tableName)
	}
	return db.connection.PutItem(input)
}


func(db *Database)Delete(condition map[string] interafce{}, tableName string) (response *dynamodb.DeleteItemOutput, err error) {
	
	conditionParsed, err := dynamodb.MarshalMap(condition)

	if err != nil {
		return nil, err
	}

	input:= &dynamodb.DeleteItemInput {
		Key: conditionParsed, 
		TableName: aws.String(tablename), 

	}

	return db.connection.DeleteItem(input)
}