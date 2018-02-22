# Localstack

We use [Localstack](https://github.com/localstack/localstack) as a stub implementation of many AWS services that we can run locally for development and testing of Taco.

## Installing
`localstack`'s dependencies are found here:
https://github.com/localstack/localstack#requirements

Follow the directions on installing `localstack` here: https://github.com/localstack/localstack#installing.


## Running

Start localstack by running the command:
```shell
$ SERVICES=dynamodb localstack start
```

Next we have to set up the services.

#### Create the druids table in DynamoDB:
```shell
$ awslocal dynamodb create-table --table-name druids \
  --attribute-definitions "AttributeName=id,AttributeType=S" \
  --key-schema "AttributeName=id,KeyType=HASH" \
  --provisioned-throughput=ReadCapacityUnits=10,WriteCapacityUnits=10
```
