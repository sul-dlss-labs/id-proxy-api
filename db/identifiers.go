package db

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/sul-dlss-labs/identifier-service/config"
)

// NewDynamoIdentifiers -- a connection to the identifiers table in Dynamo
func NewDynamoIdentifiers(config *config.Config, db *dynamodb.DynamoDB) *DynamoRepository {
	tableName := aws.String(config.IdentifierTableName)
	return &DynamoRepository{db: db,
		tableName: tableName}
}

// Repository the interface for the metadata repository
type Repository interface {
	Exists(string) (bool, error)
	CreateItem(*Druid) error
}

// DynamoRepository -- The fetching object
type DynamoRepository struct {
	db        *dynamodb.DynamoDB
	tableName *string
}

// CreateItem perist the resource in dynamo db
func (h DynamoRepository) CreateItem(d *Druid) error {
	row, err := dynamodbattribute.MarshalMap(d)

	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      row,
		TableName: h.tableName,
	}

	_, err = h.db.PutItem(input)

	if err != nil {
		return err
	}
	log.Printf("Saved %s to dynamodb", d.ID)
	return nil
}

// Exists -- given an identifier, find the resource
func (h DynamoRepository) Exists(id string) (bool, error) {
	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName:      h.tableName,
		ConsistentRead: aws.Bool(true),
	}
	resp, err := h.db.GetItem(params)
	if err != nil {
		log.Println(err)
		return false, err
	}
	var resource *Druid
	if err := dynamodbattribute.UnmarshalMap(resp.Item, &resource); err != nil {
		log.Println(err)
		return false, err
	}

	if resource.ID == "" {
		return false, nil
	}
	return true, nil
}
