// Package mongo provide a wrapper around the MongoDB driver.
package mongo

import (
	"context"
	"errors"

	"github.com/JustWorking42/go-password-manager/internal/server/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoStorage is a wrapper around the MongoDB driver.
type MongoStorage struct {
	client *mongo.Client
}

// NewMongoStorage creates a new MongoStorage.
func NewMongoStorage(ctx context.Context, config storage.Config) (storage.Storage, error) {
	clientOptions := options.Client().ApplyURI(config.Uri())
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &MongoStorage{client: client}, nil
}

// AddPassword adds a new password by userID to the database.
func (m *MongoStorage) AddPassword(ctx context.Context, id primitive.ObjectID, data storage.PasswordData) error {
	collection := m.client.Database("service").Collection("users")

	filter := bson.M{
		"_id": id,
		"passwords": bson.M{
			"$not": bson.M{
				"$elemMatch": bson.M{
					"serviceName": data.ServiceName,
				},
			},
		},
	}

	update := bson.M{
		"$push": bson.M{
			"passwords": data,
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("password with the same ServiceName already exists")
	}
	return nil
}

// GetPassword gets a password by userID and tittle from the database.
func (m *MongoStorage) GetPassword(ctx context.Context, id primitive.ObjectID, serviceTitle string) (storage.PasswordData, error) {
	collection := m.client.Database("service").Collection("users")
	var data struct {
		Passwords []storage.PasswordData `bson:"passwords"`
	}
	filter := bson.M{"_id": id, "passwords": bson.M{"$elemMatch": bson.M{"serviceName": serviceTitle}}}
	projection := bson.M{"passwords.$": 1}
	result := collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection))
	err := result.Decode(&data)
	if err != nil {
		return storage.PasswordData{}, err
	}
	if len(data.Passwords) == 0 {
		return storage.PasswordData{}, errors.New("password not found")
	}

	return data.Passwords[0], nil
}

// AddUser adds a new user to the database.
func (m *MongoStorage) AddUser(ctx context.Context, user storage.User) (primitive.ObjectID, error) {
	collection := m.client.Database("service").Collection("users")
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	userId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.ObjectID{}, errors.New("cannot convert id to object id")
	}
	return userId, nil
}

// IsLoginEnabled checks if login is enabled.
func (m *MongoStorage) IsLoginEnabled(ctx context.Context, login string) (bool, error) {
	collection := m.client.Database("service").Collection("users")
	filter := bson.M{"login": login}

	var user storage.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

// GetUser gets a user from the database.
func (m *MongoStorage) GetUser(ctx context.Context, login string) (storage.User, error) {
	collection := m.client.Database("service").Collection("users")
	filter := bson.M{"login": login}

	var user storage.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return storage.User{}, err
	}
	return user, nil
}

// AddCard adds a new card by userID to the database.
func (m *MongoStorage) AddCard(ctx context.Context, id primitive.ObjectID, card storage.CardData) error {
	collection := m.client.Database("service").Collection("users")

	filter := bson.M{
		"_id": id,
		"cards": bson.M{
			"$not": bson.M{
				"$elemMatch": bson.M{
					"cardName": card.CardName,
				},
			},
		},
	}

	update := bson.M{
		"$push": bson.M{
			"cards": card,
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("card with the same CardName already exists")
	}
	return nil
}

// GetCard gets a card  by userID and tittle from the database.
func (m *MongoStorage) GetCard(ctx context.Context, id primitive.ObjectID, cardName string) (storage.CardData, error) {
	collection := m.client.Database("service").Collection("users")
	var data struct {
		Cards []storage.CardData `bson:"cards"`
	}
	filter := bson.M{"_id": id, "cards": bson.M{"$elemMatch": bson.M{"cardName": cardName}}}
	projection := bson.M{"cards.$": 1}
	result := collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection))
	err := result.Decode(&data)
	if err != nil {
		return storage.CardData{}, err
	}
	if len(data.Cards) == 0 {
		return storage.CardData{}, errors.New("card not found")
	}

	return data.Cards[0], nil
}

// AddNote adds a new note by userID to the database.
func (m *MongoStorage) AddNote(ctx context.Context, id primitive.ObjectID, note storage.Note) error {
	collection := m.client.Database("service").Collection("users")

	filter := bson.M{
		"_id": id,
		"notes": bson.M{
			"$not": bson.M{
				"$elemMatch": bson.M{
					"name": note.Name,
				},
			},
		},
	}

	update := bson.M{
		"$push": bson.M{
			"notes": note,
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("note with the same Name already exists")
	}
	return nil
}

// GetNote gets a note  by userID and tittle from the database.
func (m *MongoStorage) GetNote(ctx context.Context, id primitive.ObjectID, noteName string) (storage.Note, error) {
	collection := m.client.Database("service").Collection("users")
	var data struct {
		Notes []storage.Note `bson:"notes"`
	}
	filter := bson.M{"_id": id, "notes": bson.M{"$elemMatch": bson.M{"name": noteName}}}
	projection := bson.M{"notes.$": 1}
	result := collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection))
	err := result.Decode(&data)
	if err != nil {
		return storage.Note{}, err
	}
	if len(data.Notes) == 0 {
		return storage.Note{}, errors.New("note not found")
	}

	return data.Notes[0], nil
}

// AddBytes adds a new bytes by userID to the database.
func (m *MongoStorage) AddBytes(ctx context.Context, id primitive.ObjectID, binaryData storage.BinaryData) error {
	collection := m.client.Database("service").Collection("users")

	filter := bson.M{
		"_id": id,
		"binaryDatas": bson.M{
			"$not": bson.M{
				"$elemMatch": bson.M{
					"name": binaryData.Name,
				},
			},
		},
	}

	update := bson.M{
		"$push": bson.M{
			"binaryDatas": binaryData,
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("binary data with the same Name already exists")
	}
	return nil
}

// GetBytes gets a bytes  by userID and tittle from the database.
func (m *MongoStorage) GetBytes(ctx context.Context, id primitive.ObjectID, binaryName string) (storage.BinaryData, error) {
	collection := m.client.Database("service").Collection("users")
	var data struct {
		BinaryDatas []storage.BinaryData `bson:"binaryDatas"`
	}
	filter := bson.M{"_id": id, "binaryDatas": bson.M{"$elemMatch": bson.M{"name": binaryName}}}
	projection := bson.M{"binaryDatas.$": 1}
	result := collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection))
	err := result.Decode(&data)
	if err != nil {
		return storage.BinaryData{}, err
	}
	if len(data.BinaryDatas) == 0 {
		return storage.BinaryData{}, errors.New("binary data not found")
	}

	return data.BinaryDatas[0], nil
}

// Close closes the database connection.
func (m *MongoStorage) Close(ctx context.Context) {
	m.client.Disconnect(ctx)
}
