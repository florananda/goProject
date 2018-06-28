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
 
 ### Rest_Apis_FileOperation/FileServer.go
  This is an implementation of REST APIs for file operations.
  APIs are:
  #### POST /register
   This is POST operation that registers the user
  #### POST /login
   This is a POST login operation used to login the users and sends a token back. This token would be used in other opearations to          identify the user.
  #### PUT /files/<filename>
   This is a PUT operation to add a file to thhe user's personal storage. Storage would be identified by username. Token generated in      Login step is used to validate and identify the user. File content is passed in the request Body.
  #### GET /files/<filename>
   This is GET operation. Given the filename in the parameter this operation retrieves the file contents from the storage and send it in    the response Body. Storage would be identified by username. Token generated in Login step is used to validate and identify the          user. 
  #### GET /files
   This is a GET operation. it finds all the file names for the corresponding user and sends all the filenames in a JSON message.          Storage would be identified by username. Token generated in Login step is used to validate and identify the user. 
  #### DELETE /files/<filename>
   This is a DELETE operation. Given the filename in the parameter this operation deletes the file from the corresponding user's            storage.Storage would be identified by username. Token generated in Login step is used to validate and identify the user.
 
 ### Channels
  THis folder contains programs related to Concurrency in Golang. It includes the following:
  #### DistributedSearchEngine.go
  Uses channels to search from a database using distribued workers
  #### ChannelSumSquaresCubes.go
  uses channel to find sum of squares and cubes of the digits of a number
  #### Workerpool.go
  uses channel to handle in and out values using a pool of workers
  #### ChannelFanin.go
  uses multiple channels to fanin to a single channel
  #### ChannelFanout.go
  uses single channel to fanout to multiple channel
