package main


import (
	"context"
	"log"
	"net"
	pb "GRPC_Totality_Corp/grpc" 
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/reflection"
)

type Server struct{}

var Users = map[int32]pb.UserDetails{
	1: {Id:1, Fname:"Steve",City:"LA",Phone:1234567890,Height:5.8,Married:true},
	2: {Id:2, Fname:"Praneetha",City:"HYD",Phone:2222222222,Height:5.1,Married:true},
	3: {Id:3, Fname:"Sreeja",City:"BNGLR",Phone:3333333333,Height:5.2,Married:true},
}

func (*Server) GetUserDetails(ctx context.Context, req *pb.Request) (*pb.UserDetails, error) {
	if req.UserId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "user_id must be a positive integer")
	}
	if req.UserId>int32(len(Users)){
		return nil, status.Errorf(codes.InvalidArgument, "user_id is not available")
	}
	

	user, ok := Users[req.UserId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return &user, nil
}

func (*Server) GetUsersDetails(ctx context.Context, req *pb.UserIds) (*pb.UserDetailsList, error) {

    var userDetailsList pb.UserDetailsList
    for _, userID := range req.UserId {
		if userID<=0{
			return nil, status.Errorf(codes.InvalidArgument, "user_id must be a positive integer")
		}
		if userID>int32(len(Users)){
			return nil, status.Errorf(codes.InvalidArgument, "user_id is not available")
		}
        user, ok := Users[userID]
        if ok {
            userDetailsList.UserDetails = append(userDetailsList.UserDetails, &user)
        }
    }

    return &userDetailsList, nil
}


func (*Server) MustEmbedUnimplementedUserServiceServer(){
	return 
}


func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &Server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


/* TestCases
1. grpcurl -plaintext -d '{"user_id": 2}' localhost:8080 UserService/GetUserDetails
Output:
{
  "id": 2,
  "fname": "Praneetha",
  "city": "HYD",
  "phone": "2222222222",
  "height": 5.1,
  "Married": true
}
2. $ grpcurl -plaintext -d '{"user_id": [1,2,3]}' localhost:8080 UserService/GetUsersDetails
Output:
{
  "userDetails": [
    {
      "id": 1,
      "fname": "Steve",
      "city": "LA",
      "phone": "1234567890",
      "height": 5.8,
      "Married": true
    },
    {
      "id": 2,
      "fname": "Praneetha",
      "city": "HYD",
      "phone": "2222222222",
      "height": 5.1,
      "Married": true
    },
    {
      "id": 3,
      "fname": "Sreeja",
      "city": "BNGLR",
      "phone": "3333333333",
      "height": 5.2,
      "Married": true
    }
  ]
}
3. grpcurl -plaintext -d '{"user_id": [0,2,3]}' localhost:8080 UserService/GetUsersDetails
ERROR:
  Code: InvalidArgument
  Message: user_id must be a positive integer

4.  grpcurl -plaintext -d '{"user_id": [1,-2,3]}' localhost:8080 UserService/GetUsersDetails
ERROR:
  Code: InvalidArgument
  Message: user_id must be a positive integer

5.  grpcurl -plaintext -d '{"user_id": [10,-2,3]}' localhost:8080 UserService/GetUsersDetails
ERROR:
  Code: InvalidArgument
  Message: user_id is not available

*/