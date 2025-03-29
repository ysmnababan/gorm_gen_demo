package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type Querier interface {
	// Get all roles and their navigation access using a CTE
	//
	// WITH RoleNav AS (
	//   SELECT r.role_id, r.role_name, rn.navigation_id, rn.allow_read, rn.allow_create
	//   FROM roles r
	//   JOIN roles_navigation rn ON r.role_id = rn.role_id
	// )
	// SELECT * FROM RoleNav WHERE role_name = @roleName
	GetRoleNavigation(roleName string) ([]*gen.T, error)
}

type Dummy struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"size:50;not null;unique"`
}

func main() {
	dsn := "host=localhost user=admin password=password dbname=mydatabase port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	// dsn := "host=206.189.144.36 user=developer password=xnuNiGE3EyYF9E9u dbname=cafetaria port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	gormdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// gormdb = gormdb.Debug()
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
		// OutPath: "../staging/repo",
		OutPath: "./query",

		Mode:             gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		WithUnitTest:     true,
		FieldNullable:    true,
		FieldCoverable:   true,
		FieldSignable:    true,
		FieldWithTypeTag: true,
	})
	_ = g

	g.UseDB(gormdb)

	// Generate models
	roles_navigation := g.GenerateModel("roles_navigation")
	roles := g.GenerateModel("roles", gen.FieldRelate(
		field.HasMany, // One Role can have many RoleNavigation entries
		"RolesNav",
		roles_navigation,
		&field.RelateConfig{
			RelateSlicePointer: true,
			GORMTag: field.GormTag{
				"foreignKey": []string{"role_id"},
				"references": []string{"role_id"},
			},
			JSONTag: "roles_nav,omitempty",
		},
	),
	)

	navigation := g.GenerateModel("navigation", gen.FieldRelate(
		field.HasMany,
		"RolesNav",
		roles_navigation,
		&field.RelateConfig{
			RelateSlicePointer: true,
			GORMTag: field.GormTag{
				"foreignKey": []string{"navigation_id"},
				"references": []string{"navigation_id"},
			},
		},
	))
	roles_navigation = g.GenerateModel("roles_navigation",
		gen.FieldRelate(field.BelongsTo, "Role", roles, &field.RelateConfig{
			RelatePointer: true,
			GORMTag: field.GormTag{
				"foreignKey": []string{"role_id"},
				"references": []string{"role_id"},
			},
		}),
		gen.FieldRelate(field.BelongsTo, "Navigation", navigation, &field.RelateConfig{
			RelatePointer: true,
			GORMTag: field.GormTag{
				"foreignKey": []string{"navigation_id"},
				"references": []string{"navigation_id"},
			},
		}),
	)

	// Apply models to generator
	g.ApplyBasic(roles, navigation, roles_navigation)

	// Apply the interface to the Role model
	g.ApplyInterface(func(Querier) {}, roles)
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
