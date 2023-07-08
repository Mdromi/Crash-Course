# AWS CLI Commands

This repository contains AWS CLI commands to set up and invoke an AWS Lambda function.

## Prerequisites

- AWS CLI should be installed on your system.
- AWS credentials should be configured using `aws configure`.

## Commands

1. Create IAM Role:
    ```
    aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json
    ```

2. Attach IAM Policy:
    ```
    aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
    ```

3. Create Function ZIP:
    ```
    zip function.zip main
    ```

4. Create Lambda Function:
    ```
    aws lambda create-function --function-name 7-lambda-yt --zip-file fileb://function.zip --handler main --runtime go1.x --role arn:aws:iam::AWS_ACCOUNT_ID:role/lambda-ex
    ```

5. Invoke Lambda Function:
    ```
    aws lambda invoke --function-name 7-lambda-yt --payload '{"What is your name?": "jim", "How old are you?": 33}' output.txt
    ```

Alternatively, you can use the following simplified command without `--cli-binary-format`:
    ```
    aws lambda invoke --function-name 7-lambda-yt --payload '{"What is your name?": "jim", "How old are you?": 33}' output.txt
    ```

## Note

Make sure to replace `AWS-account-ID` in the commands with your actual AWS account ID.

