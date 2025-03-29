package main

import (
	"fmt"
	"gorm_demo/src/query"
	"log"

	// query "gorm_demo/staging/repo"

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
	// dsn := "host=206.189.144.36 user=developer password=xnuNiGE3EyYF9E9u dbname=cafetaria port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	query.SetDefault(db)
	r := query.Role
	// role, err := r.Debug().
	// 	Select(r.RoleID, r.RoleName).
	// 	Where(r.DeletedAt.IsNull()).
	// 	Where(r.ModifiedBy.Eq("")).
	// 	Find()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, val := range role {
	// 	fmt.Println(*val)
	// }

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
		Where(r.RoleID.Eq("7a22dda3-2272-45bb-9add-93c8763e7c38")).
		Order(n.SortOrder).Scan(&navs)
	for _, n := range navs {
		fmt.Println(n)
	}

	fmt.Println("here")
	role, _ := r.Debug().Preload(r.RolesNav).Where(r.RoleID.Eq("7a22dda3-2272-45bb-9add-93c8763e7c38")).First()
	fmt.Println(role.RoleID, role.RoleName)
	rolenav := role.RolesNav
	for _, val := range rolenav {
		fmt.Println(val)
	}

	fmt.Println("HERERERERERER")
	data, _ := rn.Debug().Preload(rn.Role.Where(rn.RoleID.Eq("7a22dda3-2272-45bb-9add-93c8763e7c38"))).Preload(rn.Navigation).Find()
	fmt.Println(data[0].Role)
	for _, val := range data {
		fmt.Println(val.Navigation.NavigationID, val.Navigation.NavigationName, val.Navigation.ParentNavigationID, val.Navigation.URLPath)
		fmt.Println(val.AllowApproval, val.AllowCreate, val.AllowDelete, val.AllowRead, val.AllowUpdate)
	}

	fmt.Println("hai")
	rolesNav, err := r.Debug().GetRoleNavigation("ADMIN")
	if err != nil {
		log.Fatal("Error:", err)
	}

	for _, rn := range rolesNav {
		fmt.Println("Role:", rn.RoleName, "Navigation ID:", rn)
	}
}
