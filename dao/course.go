package dao

import (
	"context"

	"github.com/lucasfloriani/go-mongo/model"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
)

// CourseDAO persists course data in database, contains methods for each CRUD actions.
type CourseDAO struct {
	db *mongo.Collection
}

// NewCourseDAO creates a new CourseDAO
func NewCourseDAO(db *mongo.Database) *CourseDAO {
	return &CourseDAO{db.Collection("course")}
}

func (dao *CourseDAO) filter(offset, limit int) []findopt.Find {
	var elems []findopt.Find

	elems = append(elems, findopt.Limit(int64(limit)), findopt.Skip(int64(offset)))

	return elems
}

// All retrieves the course records with the specified offset and limit from the database.
func (dao *CourseDAO) All(offset, limit int) (elements []model.Course, err error) {
	cur, err := dao.db.Find(context.Background(), nil, dao.filter(offset, limit)...)
	if err != nil {
		return
	}
	defer cur.Close(context.Background())

	var elem model.Course
	for cur.Next(context.Background()) {
		if err = cur.Decode(&elem); err != nil {
			return
		}
		elements = append(elements, elem)
	}

	return
}

// Count returns the number of the course records in the database.
func (dao *CourseDAO) Count() (int, error) {
	count, err := dao.db.Count(context.Background(), nil)
	return int(count), err
}

// Get reads the course with the specified ID from the database.
func (dao *CourseDAO) Get(id string) (*model.Course, error) {
	objID, err := objectid.FromHex(id)
	if err != nil {
		return nil, err
	}
	c := model.NewCourse()
	err = dao.db.FindOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.ObjectID("_id", objID),
		),
	).Decode(c)
	return c, err
}

// Create saves a new course record in the database.
// The Course.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *CourseDAO) Create(c *model.Course) error {
	res, err := dao.db.InsertOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("name", c.Name),
			bson.EC.String("link", c.Link),
		),
	)
	if err != nil {
		return err
	}
	c.ID = res.InsertedID.(objectid.ObjectID)
	return nil
}

// Update saves the changes to an course in the database.
func (dao *CourseDAO) Update(c *model.Course) error {
	_, err := dao.db.UpdateOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.ObjectID("_id", c.ID),
		),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.String("name", c.Name),
				bson.EC.String("link", c.Link),
			),
		),
	)
	return err
}

// Delete deletes an course with the specified ID from the database.
func (dao *CourseDAO) Delete(c *model.Course) error {
	_, err := dao.db.DeleteOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.ObjectID("_id", c.ID),
		),
	)
	return err
}
