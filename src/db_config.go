package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type Dummy struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"size:50;not null;unique"`
}

func main() {
	dsn := "host=localhost user=admin password=password dbname=mydatabase port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	gormdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Create tables from models
	// err = gormdb.AutoMigrate(&Dummy{})
	// if err != nil {
	// 	log.Fatal("Failed to migrate tables:", err)
	// }
	// fmt.Println("Tables created successfully")

	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	_ = g

	g.UseDB(gormdb)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}

// var tables []struct {
// 	Schemaname string
// 	Tablename  string
// }
// var currentDB string
// // gormdb.Exec("SET search_path TO public")

// gormdb.Raw("SELECT current_database()").Scan(&currentDB)
// fmt.Println("Connected to database:", currentDB)

// var searchPath string
// gormdb.Raw("SHOW search_path").Scan(&searchPath)
// fmt.Println("Current search_path:", searchPath)
// result := gormdb.Raw("SELECT tablename FROM pg_tables WHERE schemaname = 'public'").Scan(&tables)

// if result.Error != nil {
// 	log.Fatal("Error fetching tables:", result.Error)
// }

// for _, table := range tables {
// 	fmt.Println("Found table ->", table.Schemaname+"."+table.Tablename)
// }

// fmt.Println("heelo")
// for _, table := range []string{`navigation`, "roles", "roles_navigation"} {
// 	exists := gormdb.Migrator().HasTable(table)
// 	fmt.Printf("Table %s exists: %v\n", table, exists)
// }
