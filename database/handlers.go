package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetUserData gets user data from database, if user is not found, creates new one
func (db *Database) GetUserData(userName, userId string) (bson.M, error) {
	collection := db.client.Database("farmsDb").Collection("userFarm")
	var opicak bson.M
	err := collection.FindOne(context.TODO(), bson.M{"userId": userId}).Decode(&opicak)
	if errors.Is(err, mongo.ErrNoDocuments) {
		//No user found, creates new one
		_, err = collection.InsertOne(context.TODO(), bson.D{{"userId", userId}, {"userName", userName}, {"bananas", 0}, {"xp", 0}, {"hovna", 0}})
	} else if err != nil {
		return nil, err
	}

	return opicak, err
}

func (db *Database) GetTopUsers() ([]bson.M, error) {
	collection := db.client.Database("farmsDb").Collection("userFarm")
	// Sort by `price` field descending
	findOptions := options.Find().
		SetSort(bson.D{{"bananas", -1}}).
		SetLimit(10)

	//Does the query
	documents, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}

	var monkeys []bson.M
	err = documents.All(context.TODO(), &monkeys)

	return monkeys, err
}

func (db *Database) AddBananas(userId string, banans int) error {
	collection := db.client.Database("farmsDb").Collection("userFarm")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"userId": userId},
		bson.D{
			{Key: "$inc", Value: bson.D{{Key: "bananas", Value: banans}}},
		},
	)

	return err
}

func (db *Database) ResetBananas(userId string, bananas int) error {
	collection := db.client.Database("farmsDb").Collection("userFarm")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"userId": userId},
		bson.D{
			{Key: "$inc", Value: bson.D{{Key: "bananas", Value: -bananas}}},
		},
	)

	return err
}

func (db *Database) AddHovno(userId string) error {
	collection := db.client.Database("farmsDb").Collection("userFarm")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"userId": userId},
		bson.D{
			{Key: "$inc", Value: bson.D{{Key: "hovna", Value: 1}}},
		},
	)

	return err
}

func (db *Database) SubHovno(username, userId string) (bool, error) {
	user, _ := db.GetUserData(username, userId)
	hovna, ok := user["hovna"].(int32)
	if !ok || hovna <= 0 {
		err := db.addField(userId, "hovna", 0)
		return false, err
	}

	collection := db.client.Database("farmsDb").Collection("userFarm")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"userId": userId},
		bson.D{
			{Key: "$inc", Value: bson.D{{Key: "hovna", Value: -1}}},
		},
	)

	return true, err
}

func (db *Database) AddMoney(userId string, money int) error {
	collection := db.client.Database("farmsDb").Collection("userFarm")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"userId": userId},
		bson.D{
			{Key: "$inc", Value: bson.D{{Key: "money", Value: money}}},
		},
	)

	return err
}

func (db *Database) addField(userId, fieldName string, value int) error {
	collection := db.client.Database("farmsDb").Collection("userFarm")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"userId": userId},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: fieldName, Value: value}}},
		},
	)

	return err
}
