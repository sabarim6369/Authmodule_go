package models
import "go.mongodb.org/mongo-driver/bson/primitive"
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`   
	Name      string             `bson:"name"`             
	Email     string             `bson:"email"`
	CreatedAt int64              `bson:"created_at"`
}
