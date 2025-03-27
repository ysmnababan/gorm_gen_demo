package main

import (
	"fmt"
	"gorm_demo/src/query"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type NavigationItem struct {
	NavigationID       string  `json:"navigation_id" gorm:"column:navigation_id"`
	NavigationName     string  `json:"navigation_name" gorm:"column:navigation_name"`
	ParentNavigationID *string `json:"parent_navigation_id" gorm:"column:parent_navigation_id"`
	URLPath            string  `json:"url_path" gorm:"column:url_path"`
	SortOrder          int     `json:"sort_order" gorm:"column:sort_order"`
	AllowRead          bool    `json:"allow_read" gorm:"column:allow_read"`
	AllowCreate        bool    `json:"allow_create" gorm:"column:allow_create"`
	AllowUpdate        bool    `json:"allow_update" gorm:"column:allow_update"`
	AllowDelete        bool    `json:"allow_delete" gorm:"column:allow_delete"`
	AllowApproval      bool    `json:"allow_approval" gorm:"column:allow_approval"`
}

func main() {
	dsn := "host=localhost user=admin password=password dbname=mydatabase port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	query.SetDefault(db)
	r := query.Role
	role, err := r.Debug().
		Select(r.RoleID, r.RoleName).
		Where(r.DeletedAt.IsNull()).
		Where(r.ModifiedBy.Eq("")).
		Find()
	if err != nil {
		log.Fatal(err)
	}
	for _, val := range role {
		fmt.Println(*val)
	}

	n := query.Navigation
	navigations, _ := n.Debug().
		Select(n.NavigationID, n.NavigationName, n.ParentNavigationID, n.URLPath, n.SortOrder).
		Where(n.IsActive.Is(true)).
		Where(n.DeletedAt.IsNull()).
		Find()
	for _, val := range navigations {
		fmt.Println(val)
	}

	rn := query.RolesNavigation
	navs := []*NavigationItem{}
	n.Debug().
		Select(n.NavigationID, n.NavigationName, n.ParentNavigationID, n.URLPath, n.SortOrder, rn.AllowRead, rn.AllowCreate, rn.AllowUpdate, rn.AllowDelete, rn.AllowApproval).
		Join(rn, rn.NavigationID.EqCol(n.NavigationID), rn.AllowRead.Is(true)).Join(r, r.RoleID.EqCol(rn.RoleID)).
		Where(rn.RoleID.Neq("")).Where(n.IsActive.Is(true)).Where(r.DeletedAt.IsNull()).
		Order(n.SortOrder).Scan(&navs)
	for _, n := range navs {
		fmt.Println(n)
	}
}
