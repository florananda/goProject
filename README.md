# This is a list of Small go programs: 


 ### HttpMicroService.go 
  it is amicroservice in go.
  A http server and the request details are saved in mongodb cloud. 
  Change the "db server address", "dbusername", and "dbpassword" according to the corresponding database information.
  Corresponding test files are in test folder.
  
  ```
  go run HttpMicroService.go  
  ```
  For test method run the following command in the test folder
  
  ```
  go test -run=TestHandler
  ```
 For mongodb cloud create an account and a free database in mlab.
 
 ```
 https://mlab.com/login
 ```
