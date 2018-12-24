package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jessevdk/go-flags"
	"gopkg.in/cheggaaa/pb.v1"
	"log"
	"os"
)

type options struct {
	TableName   string `short:"t" long:"table-name" description:"Name of the dynamo db table" required:"true"`
	EndpointUrl string `short:"e" long:"endpoint-url" description:"Endpoint url of destination dynamodb instance" required:"true"`
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(0)
	}
	sess, err := session.NewSession()
	srcDynamoDB := getDynamoDbSession(sess)
	destDynamoDB := getDynamoDbSessionUsingEndpointUrl(sess, opts.EndpointUrl)
	output, err := srcDynamoDB.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(opts.TableName),
	})
	checkError("unable to fetch total count of records", err)
	bar := pb.StartNew(int(*output.Table.ItemCount))
	result := doFirstScan(opts.TableName, srcDynamoDB, destDynamoDB, bar)
	doSubSequentScan(result, opts.TableName, srcDynamoDB, destDynamoDB, bar)
	bar.Finish()
}

func doFirstScan(tableName string, srcDynamoDB *dynamodb.DynamoDB, destDynamoDB *dynamodb.DynamoDB, pb *pb.ProgressBar) *dynamodb.ScanOutput {
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	result, err := srcDynamoDB.Scan(params)
	checkError("Query API call failed", err)
	writeItem(result, tableName, destDynamoDB, func() {
		pb.Increment()
	})
	return result
}

func doSubSequentScan(firstScanResult *dynamodb.ScanOutput, tableName string, srcDynamoDB *dynamodb.DynamoDB, destDynamoDB *dynamodb.DynamoDB, pb *pb.ProgressBar) {
	lastEvaluatedKey := firstScanResult.LastEvaluatedKey
	for lastEvaluatedKey != nil {
		params := &dynamodb.ScanInput{
			TableName:         aws.String(tableName),
			ExclusiveStartKey: lastEvaluatedKey,
		}
		result, err := srcDynamoDB.Scan(params)
		checkError("Query API call failed", err)
		lastEvaluatedKey = result.LastEvaluatedKey
		writeItem(result, tableName, destDynamoDB, func() {
			pb.Increment()
		})
	}
}

func writeItem(result *dynamodb.ScanOutput, tableName string, destDynamoDB *dynamodb.DynamoDB, onItemPutSuccess func()) {
	for i, item := range result.Items {
		putItemInput := &dynamodb.PutItemInput{
			Item:      item,
			TableName: aws.String(tableName),
		}
		_, err := destDynamoDB.PutItem(putItemInput)
		checkError(fmt.Sprintf("put item failed for item %v at %v position", item, i), err)
		onItemPutSuccess()
	}
}

func getDynamoDbSession(sess *session.Session) *dynamodb.DynamoDB {
	return dynamodb.New(sess)
}

func getDynamoDbSessionUsingEndpointUrl(sess *session.Session, endpointURL string) *dynamodb.DynamoDB {
	return dynamodb.New(sess, aws.NewConfig().WithEndpoint(endpointURL))
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, ": ", err.Error())
	}
}
