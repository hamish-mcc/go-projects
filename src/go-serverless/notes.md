1. Create a Lambda function in AWS Console
   - Assign a function name
   - Runtime - Go 1.x
   - Execution role
     - Create a new role from AWS policy templates -> Simple microservices permissions
     - Assign a name
   - Create function
   - Runtime settings -> change Handler to 'main'
   - Code source -> upload .zip file
2. Create a Dynamo DB table in AWS Console
   - Create table
   - Assign name
   - Partion key (PK) -> email (String)
   - Create table
3. Configure API gateway in AWS console
   - REST API -> Build
   - Protocol = REST, Create new API = New API
   - Assign name -> Create API
   - Actions -> Create Method
     - ANY
     - Type - Lambda Function
     - Use Lambda Proxy integration
     - Assign Lambda function name
     - Default timeout - checked
   - Actions -> Deploy API
     - Deployment stage -> New Stage
     - Name -> staging
     - Deploy
4. Test
   - https://oq5yg6ngxd.execute-api.us-west-1.amazonaws.com/staging
