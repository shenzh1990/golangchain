package gorm

import (
	"github.com/gotoeasy/glang/cmn"
	"github.com/pgvector/pgvector-go"
	"golangchain/pkg/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"math"
	"reflect"
)

var Db *gorm.DB
var DbConfig = settings.AppConfig.Db

func Start() {
	Db, err := gorm.Open(postgres.Open(DbConfig.DBUrl), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   true, // skip the snake_casing of names
		},
	})
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&GormItem{})
}

type GormItem struct {
	gorm.Model
	Embedding pgvector.Vector `gorm:"type:vector(3)"`
}

func CreateDocs() {

}
func CreateGormItems(db *gorm.DB) {
	items := []GormItem{
		GormItem{Embedding: pgvector.NewVector([]float32{1, 1, 1})},
		GormItem{Embedding: pgvector.NewVector([]float32{2, 2, 2})},
		GormItem{Embedding: pgvector.NewVector([]float32{1, 1, 2})},
	}

	result := db.Create(items)

	if result.Error != nil {
		panic(result.Error)
	}
}

func TestGorm() {
	db, err := gorm.Open(postgres.Open(DbConfig.DBUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS vector")
	db.Exec("DROP TABLE IF EXISTS gorm_items")

	db.AutoMigrate(&GormItem{})

	db.Exec("CREATE INDEX ON gorm_items USING hnsw (embedding vector_l2_ops)")

	CreateGormItems(db)

	var items []GormItem
	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "embedding <-> ?", Vars: []interface{}{pgvector.NewVector([]float32{1, 1, 1})}},
	}).Limit(5).Find(&items)
	if items[0].ID != 1 || items[1].ID != 3 || items[2].ID != 2 {
		cmn.Error("Bad ids")
	}
	if !reflect.DeepEqual(items[1].Embedding.Slice(), []float32{1, 1, 2}) {
		cmn.Error("Bad embedding")
	}

	var distances []float64
	db.Model(&GormItem{}).Select("embedding <-> ?", pgvector.NewVector([]float32{1, 1, 1})).Order("id").Find(&distances)
	if distances[0] != 0 || distances[1] != math.Sqrt(3) || distances[2] != 1 {
		cmn.Error("Bad distances")
	}
}
