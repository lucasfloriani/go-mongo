package dao

import (
	"context"

	"github.com/lucasfloriani/go-mongo/model"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
)

// UserDAO persists user data in database, contains methods for each CRUD actions.
type UserDAO struct {
	db *mongo.Collection
}

// NewUserDAO creates a new UserDAO
func NewUserDAO(db *mongo.Database) *UserDAO {
	return &UserDAO{db.Collection("user")}
}

func (dao *UserDAO) filter(offset, limit int) []findopt.Find {
	var elems []findopt.Find

	elems = append(elems, findopt.Limit(int64(limit)), findopt.Skip(int64(offset)))

	return elems
}

// All retrieves the user records with the specified offset and limit from the database.
func (dao *UserDAO) All(offset, limit int) (elements []model.User, err error) {
	cur, err := dao.db.Find(context.Background(), nil, dao.filter(offset, limit)...)
	if err != nil {
		return
	}
	defer cur.Close(context.Background())

	var elem model.User
	for cur.Next(context.Background()) {
		if err = cur.Decode(&elem); err != nil {
			return
		}
		elements = append(elements, elem)
	}

	return
}

// Count returns the number of the user records in the database.
func (dao *UserDAO) Count() (int, error) {
	count, err := dao.db.Count(context.Background(), nil)
	return int(count), err
}

// Get reads the user with the specified ID from the database.
func (dao *UserDAO) Get(id string) (*model.User, error) {
	objID, err := objectid.FromHex(id)
	if err != nil {
		return nil, err
	}
	u := model.NewUser()
	err = dao.db.FindOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.ObjectID("_id", objID),
		),
	).Decode(u)
	return u, err
}

// Create saves a new user record in the database.
// The User.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *UserDAO) Create(u *model.User) error {
	res, err := dao.db.InsertOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("name", u.Name),
			bson.EC.Int32("age", int32(u.Age)),
			bson.EC.SubDocumentFromElements("address",
				bson.EC.String("name", u.Address.Name),
			),
			bson.EC.ArrayFromElements("phones", dao.getPhones(u)...),
			bson.EC.ArrayFromElements("courses", dao.getCourses(u)...),
		),
	)
	if err != nil {
		return err
	}
	u.ID = res.InsertedID.(objectid.ObjectID)
	return nil
}

// Update saves the changes to an user in the database.
func (dao *UserDAO) Update(u *model.User) error {
	_, err := dao.db.UpdateOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.ObjectID("_id", u.ID),
		),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.String("name", u.Name),
				bson.EC.Int32("age", int32(u.Age)),
				bson.EC.SubDocumentFromElements("address",
					bson.EC.String("name", u.Address.Name),
				),
				bson.EC.ArrayFromElements("phones", dao.getPhones(u)...),
				bson.EC.ArrayFromElements("courses", dao.getCourses(u)...),
			),
		),
	)
	return err
}

// Delete deletes an user with the specified ID from the database.
func (dao *UserDAO) Delete(u *model.User) error {
	_, err := dao.db.DeleteOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.ObjectID("_id", u.ID),
		),
	)
	return err
}

func (dao *UserDAO) getPhones(u *model.User) (elems []*bson.Value) {
	for _, phone := range u.Phones {
		elems = append(elems,
			bson.VC.DocumentFromElements(
				bson.EC.String("number", phone.Number),
			),
		)
	}

	return
}

func (dao *UserDAO) getCourses(u *model.User) (elems []*bson.Value) {
	for _, course := range u.Courses {
		elems = append(elems,
			bson.VC.DocumentFromElements(
				bson.EC.ObjectID("_id", course.ID),
				bson.EC.String("name", course.Name),
				bson.EC.String("link", course.Link),
			),
		)
	}

	return
}
