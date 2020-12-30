package user

import (
	context "context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Server - user service server
type Server struct{}

type user struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Age      string `json:"age,omitempty" bson:"age,omitempty"`
	MobileNo string `json:"mobile_no,omitempty" bson:"mobile_no,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
}

const host = "localhost"
const port = "27017"
const dbname = "sofikuldb"

// SaveUser - this is the SaveUser method on server
func (s *Server) SaveUser(ctx context.Context, userReq *UserRequest) (*UserReply, error) {
	log.Printf("Received: %+v", userReq)

	user01 := user{Name: userReq.Name, Age: userReq.Age, Email: userReq.Email, MobileNo: userReq.MobileNo}
	connectionString := fmt.Sprintf("mongodb://%s:%s/%s", host, port, dbname)
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	dbConnectMsg := fmt.Sprintf("Connected to DB %s", connectionString)
	fmt.Println(dbConnectMsg)
	DB := client.Database(dbname)
	res, err := DB.Collection("user_grpc").InsertOne(context.TODO(), user01)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("Document inserted with id %s", res.InsertedID))
	}
	msg := fmt.Sprintf("Inserted id : %v", res.InsertedID)
	return &UserReply{Message: msg, Error: "false", Payload: ""}, nil
}
