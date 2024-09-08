## CloudBees task

```
Background: All API referenced are gRPC APIs, not REST ones. 
* I want to board a train from London to France. The train ticket will cost $20.  
* Create API where you can submit a purchase for a ticket.  Details included in the receipt are: 
    a. From, To, User , price paid.
    b. User should include first and last name, email address
* The user is allocated a seat in the train.  Assume the train has only 2 sections, section A and section B.
* An API that shows the details of the receipt for the user
* An API that lets you view the users and seat they are allocated by the requested section
* An API to remove a user from the train
* An API to modify a user's seat

* Code must compile with some effort on unit tests, doesn't have to be 100%, but it shouldn't be 0%.
* Please code this with Golang and gRPC
* Adding a persistence layer will be cumbersome, so just store the data in your current session/in memory.
* The results can be in the console output from your grpc-server and grpc-client
```

### Solution:
```
* gRPC is used for the communication between the client and the server.
* A hashmap as in-memory storage is used for storing the data, so on server restart, all data will be lost.
* 2D array is used to store the seat allocations for each section. Section A is denoted with 0 index and Section B is denoted with index 1.
* UserEmail is used as a unique identifier for each user.
```


### Pre-requisites:
```
1. go 1.23.0
2. protoc 3.21.12
```

### Run tests:
```
go test ./...
```

### Build & run server:
```    
go build -o booking-server && ./booking-server
```

### Build & run client:
```    
go build -o booking-client  ./client &&./booking-client
```



### API:
```    
rpc Create (BookingRequest) returns (Receipt) {}
rpc GetReceipt (UserEmail) returns (Receipt) {}
rpc UpdateSeat (UpdateSeatRequest) returns (Receipt) {}
rpc Cancel (UserEmail) returns (Empty) {}
rpc GetSeatAllocations (Section) return (AllocationList) {}
```