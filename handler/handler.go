package handler

import (
	"context"
	"gorm_demo/src/model"
	"gorm_demo/src/query"
)

type handler struct {
	Q *query.Query
}

func NewHandler(Q *query.Query) *handler {
	return &handler{
		Q: Q,
	}
}

func (h *handler) ListNav(ctx context.Context) ([]*model.Navigation, error) {
	n := h.Q.Navigation

	return n.Debug().
		Select(n.NavigationID, n.NavigationName, n.ParentNavigationID, n.URLPath, n.SortOrder).
		Where(n.IsActive.Is(true)).
		Where(n.DeletedAt.IsNull()).
		Find()
}
