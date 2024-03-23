// Package mongo provide a wrapper around the MongoDB driver.
package mongo

import (
	"context"
	"errors"

	"github.com/JustWorking42/go-password-manager/internal/server/storage"
	"github.com/sirupsen/logrus"
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
	clientOptions := options.Client().ApplyURI(config.Set)
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
	session, err := m.client.StartSession()
	if err != nil {
		logrus.WithField("userId", id).Errorf("Failed to start session: %v", err)
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
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

		result, err := collection.UpdateOne(sessCtx, filter, update)
		if err != nil {
			logrus.WithField("userId", id).Errorf("Failed to add password for user ID: %v, error: %v", id, err)
			return nil, err
		}
		if result.ModifiedCount == 0 {
			logrus.WithField("userId", id).WithField("serviceName", data.ServiceName).Errorf("Password with the same ServiceName already exists for user ID: %v", id)
			return nil, errors.New("password with the same ServiceName already exists")
		}
		logrus.WithField("userId", id).Infof("Password added successfully for user ID: %v", id)
		return nil, nil
	}

	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		logrus.WithField("userId", id).Errorf("Failed to add password in transaction for user ID: %v, error: %v", id, err)
		return err
	}
	return nil
}

// GetPassword gets a password by userID and tittle from the database.
func (m *MongoStorage) GetPassword(ctx context.Context, id primitive.ObjectID, serviceTitle string) (storage.PasswordData, error) {
	session, err := m.client.StartSession()
	if err != nil {
		logrus.WithField("userId", id).Errorf("Failed to start session: %v", err)
		return storage.PasswordData{}, err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		collection := m.client.Database("service").Collection("users")
		var data struct {
			Passwords []storage.PasswordData `bson:"passwords"`
		}
		filter := bson.M{"_id": id, "passwords": bson.M{"$elemMatch": bson.M{"serviceName": serviceTitle}}}
		projection := bson.M{"passwords.$": 1}
		result := collection.FindOne(sessCtx, filter, options.FindOne().SetProjection(projection))
		err := result.Decode(&data)
		if err != nil {
			logrus.WithField("userId", id).Errorf("Failed to get password for user ID: %v, error: %v", id, err)
			return nil, err
		}
		if len(data.Passwords) == 0 {
			logrus.WithField("userId", id).Errorf("Password not found for user ID: %v", id)
			return nil, errors.New("password not found")
		}

		logrus.WithField("userId", id).Infof("Password retrieved successfully for user ID: %v", id)
		return data.Passwords[0], nil
	}

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		logrus.WithField("userId", id).Errorf("Failed to get password in transaction for user ID: %v, error: %v", id, err)
		return storage.PasswordData{}, err
	}

	passwordData, ok := result.(storage.PasswordData)
	if !ok {
		logrus.WithField("userId", id).Error("Unexpected result type")
		return storage.PasswordData{}, errors.New("unexpected result type")
	}

	return passwordData, nil
}

// AddUser adds a new user to the database.
func (m *MongoStorage) AddUser(ctx context.Context, user storage.User) (primitive.ObjectID, error) {
	session, err := m.client.StartSession()
	if err != nil {
		logrus.Errorf("Failed to start session: %v", err)
		return primitive.ObjectID{}, err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		collection := m.client.Database("service").Collection("users")
		result, err := collection.InsertOne(sessCtx, user)
		if err != nil {
			logrus.Errorf("Failed to insert user: %v", err)
			return nil, err
		}
		userId, ok := result.InsertedID.(primitive.ObjectID)
		if !ok {
			logrus.Error("Failed to convert id to object id")
			return nil, errors.New("cannot convert id to object id")
		}
		logrus.Infof("User added successfully with ID: %v", userId)
		return userId, nil
	}

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		logrus.Errorf("Failed to add user in transaction: %v", err)
		return primitive.ObjectID{}, err
	}

	userId, ok := result.(primitive.ObjectID)
	if !ok {
		logrus.Error("Unexpected result type")
		return primitive.ObjectID{}, errors.New("unexpected result type")
	}

	return userId, nil
}

// IsLoginEnabled checks if login is enabled.
func (m *MongoStorage) IsLoginEnabled(ctx context.Context, login string) (bool, error) {
	session, err := m.client.StartSession()
	if err != nil {
		logrus.WithField("login", login).Errorf("Failed to start session: %v", err)
		return false, err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		collection := m.client.Database("service").Collection("users")
		filter := bson.M{"login": login}

		var user storage.User
		err := collection.FindOne(sessCtx, filter).Decode(&user)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				logrus.WithField("login", login).Info("Login is enabled")
				return true, nil
			}
			logrus.WithField("login", login).Errorf("is not enabled error: %v", err)
			return false, err
		}
		logrus.WithField("login", login).Error("Login is not enabled")
		return false, nil
	}

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		logrus.WithField("login", login).Errorf("Failed to check if login is enabled in transaction: %v", err)
		return false, err
	}

	isEnabled, ok := result.(bool)
	if !ok {
		logrus.WithField("login", login).Error("Unexpected result type")
		return false, errors.New("unexpected result type")
	}

	return isEnabled, nil
}

// GetUser gets a user from the database.
func (m *MongoStorage) GetUser(ctx context.Context, login string) (storage.User, error) {
	session, err := m.client.StartSession()
	if err != nil {
		logrus.WithField("login", login).Errorf("Failed to start session: %v", err)
		return storage.User{}, err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		collection := m.client.Database("service").Collection("users")
		filter := bson.M{"login": login}

		var user storage.User
		err := collection.FindOne(sessCtx, filter).Decode(&user)
		if err != nil {
			logrus.WithField("login", login).Errorf("Failed to get user: %v", err)
			return nil, err
		}
		logrus.Info("User found")
		return user, nil
	}

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		logrus.WithField("login", login).Errorf("Failed to get user in transaction: %v", err)
		return storage.User{}, err
	}

	user, ok := result.(storage.User)
	if !ok {
		logrus.WithField("login", login).Error("Unexpected result type")
		return storage.User{}, errors.New("unexpected result type")
	}

	return user, nil
}

// AddCard adds a new card by userID to the database.
func (m *MongoStorage) AddCard(ctx context.Context, id primitive.ObjectID, card storage.CardData) error {
	session, err := m.client.StartSession()
	if err != nil {
		logrus.WithField("id", id).Errorf("Failed to start session: %v", err)
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
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
		result, err := collection.UpdateOne(sessCtx, filter, update)
		if err != nil {
			logrus.WithField("id", id).Errorf("Failed to add card: %v", err)
			return nil, err
		}
		if result.ModifiedCount == 0 {
			logrus.WithField("id", id).WithField("cardName", card.CardName).Error("Failed to add card: card already exists")
			return nil, errors.New("card already exists")
		}
		logrus.WithField("id", id).WithField("cardName", card.CardName).Info("Card added")
		return nil, nil
	}

	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		logrus.WithField("id", id).WithField("cardName", card.CardName).Errorf("Failed to add card in transaction: %v", err)
		return err
	}
	return nil
}

// GetCard gets a card  by userID and tittle from the database.
func (m *MongoStorage) GetCard(ctx context.Context, id primitive.ObjectID, cardName string) (storage.CardData, error) {
	session, err := m.client.StartSession()
	if err != nil {
		logrus.WithField("id", id).Errorf("Failed to start session: %v", err)
		return storage.CardData{}, err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		collection := m.client.Database("service").Collection("users")
		var data struct {
			Cards []storage.CardData `bson:"cards"`
		}
		filter := bson.M{"_id": id, "cards": bson.M{"$elemMatch": bson.M{"cardName": cardName}}}
		projection := bson.M{"cards.$": 1}
		result := collection.FindOne(sessCtx, filter, options.FindOne().SetProjection(projection))
		err := result.Decode(&data)
		if err != nil {
			logrus.WithField("id", id).WithField("cardName", cardName).Errorf("Failed to get card: %v", err)
			return nil, err
		}
		if len(data.Cards) == 0 {
			logrus.WithField("id", id).WithField("cardName", cardName).Error("Failed to get card: card not found")
			return storage.CardData{}, errors.New("card not found")
		}
		logrus.WithField("id", id).WithField("cardName", cardName).Info("Card retrieved")
		return data.Cards[0], nil
	}

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		logrus.WithField("id", id).Errorf("Failed to get card: %v", err)
		return storage.CardData{}, err
	}

	cardData, ok := result.(storage.CardData)
	if !ok {
		logrus.WithField("id", id).WithField("cardName", cardName).Error("Failed to get card: unexpected result type")
		return storage.CardData{}, errors.New("unexpected result type")
	}

	return cardData, nil
}

// AddNote adds a new note by userID to the database.
func (m *MongoStorage) AddNote(ctx context.Context, id primitive.ObjectID, note storage.Note) error {
	session, err := m.client.StartSession()
	if err != nil {
		logrus.WithField("id", id).Errorf("Failed to start session: %v", err)
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
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
		result, err := collection.UpdateOne(sessCtx, filter, update)
		if err != nil {
			logrus.WithField("id", id).Errorf("Failed to add note: %v", err)
			return nil, err
		}
		if result.ModifiedCount == 0 {
			logrus.WithField("id", id).WithField("noteName", note.Name).Errorf("Failed to add note: note already exists")
			return nil, errors.New("note with the same Name already exists")
		}
		logrus.WithField("id", id).WithField("noteName", note.Name).Info("Note added")
		return nil, nil
	}

	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		logrus.WithField("id", id).Errorf("Failed to add note: %v", err)
		return err
	}
	return nil
}

// GetNote gets a note  by userID and tittle from the database.
func (m *MongoStorage) GetNote(ctx context.Context, id primitive.ObjectID, noteName string) (storage.Note, error) {
	session, err := m.client.StartSession()
	if err != nil {
		logrus.WithField("id", id).Errorf("Failed to start session: %v", err)
		return storage.Note{}, err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		collection := m.client.Database("service").Collection("users")
		var data struct {
			Notes []storage.Note `bson:"notes"`
		}
		filter := bson.M{"_id": id, "notes": bson.M{"$elemMatch": bson.M{"name": noteName}}}
		projection := bson.M{"notes.$": 1}
		result := collection.FindOne(sessCtx, filter, options.FindOne().SetProjection(projection))
		err := result.Decode(&data)
		if err != nil {
			logrus.WithField("id", id).Errorf("Failed to get note: %v", err)
			return nil, err
		}
		if len(data.Notes) == 0 {
			logrus.WithField("id", id).WithField("noteName", noteName).Errorf("Failed to get note: note not found")
			return storage.Note{}, errors.New("note not found")
		}
		logrus.WithField("id", id).WithField("noteName", noteName).Info("Note retrieved")
		return data.Notes[0], nil
	}

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		logrus.WithField("id", id).Errorf("Failed to get note: %v", err)
		return storage.Note{}, err
	}

	note, ok := result.(storage.Note)
	if !ok {
		logrus.WithField("id", id).Errorf("unexpected result type")
		return storage.Note{}, errors.New("unexpected result type")
	}

	return note, nil
}

// AddBytes adds a new bytes by userID to the database.
func (m *MongoStorage) AddBytes(ctx context.Context, id primitive.ObjectID, binaryData storage.BinaryData) error {
	session, err := m.client.StartSession()
	if err != nil {
		logrus.WithField("id", id).Errorf("Failed to start session: %v", err)
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
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
		result, err := collection.UpdateOne(sessCtx, filter, update)
		if err != nil {
			logrus.WithField("id", id).Errorf("Failed to add bytes: %v", err)
			return nil, err
		}
		if result.ModifiedCount == 0 {
			logrus.WithField("id", id).WithField("binaryDataName", binaryData.Name).Errorf("Failed to add bytes: binary data already exists")
			return nil, errors.New("binary data already exists")
		}
		return nil, nil
	}

	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		logrus.WithField("id", id).Errorf("Failed to add bytes: %v", err)
		return err
	}
	return nil
}

// GetBytes gets a bytes  by userID and tittle from the database.
func (m *MongoStorage) GetBytes(ctx context.Context, id primitive.ObjectID, binaryName string) (storage.BinaryData, error) {
	session, err := m.client.StartSession()
	if err != nil {
		logrus.WithField("id", id).Errorf("Failed to start session: %v", err)
		return storage.BinaryData{}, err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		collection := m.client.Database("service").Collection("users")
		var data struct {
			BinaryDatas []storage.BinaryData `bson:"binaryDatas"`
		}
		filter := bson.M{"_id": id, "binaryDatas": bson.M{"$elemMatch": bson.M{"name": binaryName}}}
		projection := bson.M{"binaryDatas.$": 1}
		result := collection.FindOne(sessCtx, filter, options.FindOne().SetProjection(projection))
		err := result.Decode(&data)
		if err != nil {
			logrus.WithField("id", id).Errorf("Failed to get bytes: %v", err)
			return nil, err
		}
		if len(data.BinaryDatas) == 0 {
			logrus.WithField("id", id).WithField("binaryName", binaryName).Errorf("Failed to get bytes: binary data not found")
			return nil, errors.New("binary data not found")
		}
		logrus.WithField("id", id).WithField("binaryName", binaryName).WithField("binaryData", data.BinaryDatas[0]).Infof("Got bytes")
		return data.BinaryDatas[0], nil
	}

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		logrus.WithField("id", id).Errorf("Failed to get bytes: %v", err)
		return storage.BinaryData{}, err
	}

	binaryData, ok := result.(storage.BinaryData)
	if !ok {
		logrus.WithField("id", id).Errorf("Failed to get bytes: unexpected result type")
		return storage.BinaryData{}, errors.New("unexpected result type")
	}

	return binaryData, nil
}

// Close closes the database connection.
func (m *MongoStorage) Close(ctx context.Context) error {
	err := m.client.Disconnect(ctx)
	if err != nil {
		logrus.Errorf("error closing mongo client: %v", err)
		return err
	}
	return nil
}
