package repository

import (
	"cloud-service/entity"
	"database/sql"
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Equals(one entity.ResourceEntity, other entity.ResourceEntity) bool {
	return one.ID == other.ID &&
		one.Name == other.Name &&
		one.Key == other.Key &&
		one.ResourceType == other.ResourceType &&
		one.Size == other.Size &&
		one.ParentId == other.ParentId &&
		compareChilds(one.Childs, other.Childs)
}

func compareChilds(one []entity.ResourceEntity, two []entity.ResourceEntity) bool {
	if len(one) != len(two) {
		return false
	}

	for i, v := range one {
		if !Equals(v, two[i]) {
			return false
		}
	}

	return true
}

var (
	mock       sqlmock.Sqlmock
	db         *gorm.DB
	repository *ResourceRepository
)

func TestMain(m *testing.M) {
	var mockDb *sql.DB
	mockDb, mock, _ = sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})

	db, _ = gorm.Open(dialector, &gorm.Config{})
	repository = NewResourceRepository(db)

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestCreate(t *testing.T) {
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "resources" ("created_at","updated_at","deleted_at","name","key","resource_type","size","parent_id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8) ON CONFLICT ("id") DO UPDATE SET "parent_id"="excluded"."parent_id" RETURNING "id","parent_id"`,
	)).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "node4", "", 0, 0, sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"id", "parent_id"}).AddRow(1, nil))
	mock.ExpectCommit()
	expectedEntityNode1 := &entity.ResourceEntity{
		Name:   "node1",
		Childs: make([]entity.ResourceEntity, 0),
	}
	expectedEntityNode2 := &entity.ResourceEntity{
		Name:   "node2",
		Childs: []entity.ResourceEntity{*expectedEntityNode1},
	}
	expectedEntityNode3 := &entity.ResourceEntity{
		Name:   "node3",
		Childs: []entity.ResourceEntity{*expectedEntityNode2},
	}
	expectedEntityNode4 := &entity.ResourceEntity{
		Name:   "node4",
		Childs: []entity.ResourceEntity{*expectedEntityNode3},
	}

	createdEntity, err := repository.CreateNewResource(*expectedEntityNode4)
	if err != nil {
		t.Errorf("CreateNewResource error: %s", err)
	}

	if Equals(createdEntity, *expectedEntityNode4) {
		t.Errorf("created and founded are not equal")
	}

}
