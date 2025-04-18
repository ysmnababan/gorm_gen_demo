// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package repo

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                 = new(Query)
	ActorAdmin        *actorAdmin
	Banner            *banner
	BannerAd          *bannerAd
	Cart              *cart
	CartItem          *cartItem
	Channel           *channel
	Config            *config
	Cuisine           *cuisine
	Customer          *customer
	DataEnv           *dataEnv
	Inbox             *inbox
	InboxCustomer     *inboxCustomer
	Location          *location
	LocationCode      *locationCode
	LogKeyword        *logKeyword
	LogRecent         *logRecent
	MenuFavorite      *menuFavorite
	MenuStock         *menuStock
	Merchant          *merchant
	MerchantKyc       *merchantKyc
	Navigation        *navigation
	Order             *order
	OrderFlow         *orderFlow
	OrderItem         *orderItem
	OrderRating       *orderRating
	OrderRefund       *orderRefund
	Outlet            *outlet
	OutletChannel     *outletChannel
	OutletCode        *outletCode
	OutletCuisine     *outletCuisine
	OutletFavorite    *outletFavorite
	OutletLogName     *outletLogName
	OutletMenu        *outletMenu
	OutletOp          *outletOp
	OutletQri         *outletQri
	OutletSection     *outletSection
	OutletSectionMenu *outletSectionMenu
	Payment           *payment
	Role              *role
	RolesNavigation   *rolesNavigation
	UserEvent         *userEvent
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	ActorAdmin = &Q.ActorAdmin
	Banner = &Q.Banner
	BannerAd = &Q.BannerAd
	Cart = &Q.Cart
	CartItem = &Q.CartItem
	Channel = &Q.Channel
	Config = &Q.Config
	Cuisine = &Q.Cuisine
	Customer = &Q.Customer
	DataEnv = &Q.DataEnv
	Inbox = &Q.Inbox
	InboxCustomer = &Q.InboxCustomer
	Location = &Q.Location
	LocationCode = &Q.LocationCode
	LogKeyword = &Q.LogKeyword
	LogRecent = &Q.LogRecent
	MenuFavorite = &Q.MenuFavorite
	MenuStock = &Q.MenuStock
	Merchant = &Q.Merchant
	MerchantKyc = &Q.MerchantKyc
	Navigation = &Q.Navigation
	Order = &Q.Order
	OrderFlow = &Q.OrderFlow
	OrderItem = &Q.OrderItem
	OrderRating = &Q.OrderRating
	OrderRefund = &Q.OrderRefund
	Outlet = &Q.Outlet
	OutletChannel = &Q.OutletChannel
	OutletCode = &Q.OutletCode
	OutletCuisine = &Q.OutletCuisine
	OutletFavorite = &Q.OutletFavorite
	OutletLogName = &Q.OutletLogName
	OutletMenu = &Q.OutletMenu
	OutletOp = &Q.OutletOp
	OutletQri = &Q.OutletQri
	OutletSection = &Q.OutletSection
	OutletSectionMenu = &Q.OutletSectionMenu
	Payment = &Q.Payment
	Role = &Q.Role
	RolesNavigation = &Q.RolesNavigation
	UserEvent = &Q.UserEvent
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                db,
		ActorAdmin:        newActorAdmin(db, opts...),
		Banner:            newBanner(db, opts...),
		BannerAd:          newBannerAd(db, opts...),
		Cart:              newCart(db, opts...),
		CartItem:          newCartItem(db, opts...),
		Channel:           newChannel(db, opts...),
		Config:            newConfig(db, opts...),
		Cuisine:           newCuisine(db, opts...),
		Customer:          newCustomer(db, opts...),
		DataEnv:           newDataEnv(db, opts...),
		Inbox:             newInbox(db, opts...),
		InboxCustomer:     newInboxCustomer(db, opts...),
		Location:          newLocation(db, opts...),
		LocationCode:      newLocationCode(db, opts...),
		LogKeyword:        newLogKeyword(db, opts...),
		LogRecent:         newLogRecent(db, opts...),
		MenuFavorite:      newMenuFavorite(db, opts...),
		MenuStock:         newMenuStock(db, opts...),
		Merchant:          newMerchant(db, opts...),
		MerchantKyc:       newMerchantKyc(db, opts...),
		Navigation:        newNavigation(db, opts...),
		Order:             newOrder(db, opts...),
		OrderFlow:         newOrderFlow(db, opts...),
		OrderItem:         newOrderItem(db, opts...),
		OrderRating:       newOrderRating(db, opts...),
		OrderRefund:       newOrderRefund(db, opts...),
		Outlet:            newOutlet(db, opts...),
		OutletChannel:     newOutletChannel(db, opts...),
		OutletCode:        newOutletCode(db, opts...),
		OutletCuisine:     newOutletCuisine(db, opts...),
		OutletFavorite:    newOutletFavorite(db, opts...),
		OutletLogName:     newOutletLogName(db, opts...),
		OutletMenu:        newOutletMenu(db, opts...),
		OutletOp:          newOutletOp(db, opts...),
		OutletQri:         newOutletQri(db, opts...),
		OutletSection:     newOutletSection(db, opts...),
		OutletSectionMenu: newOutletSectionMenu(db, opts...),
		Payment:           newPayment(db, opts...),
		Role:              newRole(db, opts...),
		RolesNavigation:   newRolesNavigation(db, opts...),
		UserEvent:         newUserEvent(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	ActorAdmin        actorAdmin
	Banner            banner
	BannerAd          bannerAd
	Cart              cart
	CartItem          cartItem
	Channel           channel
	Config            config
	Cuisine           cuisine
	Customer          customer
	DataEnv           dataEnv
	Inbox             inbox
	InboxCustomer     inboxCustomer
	Location          location
	LocationCode      locationCode
	LogKeyword        logKeyword
	LogRecent         logRecent
	MenuFavorite      menuFavorite
	MenuStock         menuStock
	Merchant          merchant
	MerchantKyc       merchantKyc
	Navigation        navigation
	Order             order
	OrderFlow         orderFlow
	OrderItem         orderItem
	OrderRating       orderRating
	OrderRefund       orderRefund
	Outlet            outlet
	OutletChannel     outletChannel
	OutletCode        outletCode
	OutletCuisine     outletCuisine
	OutletFavorite    outletFavorite
	OutletLogName     outletLogName
	OutletMenu        outletMenu
	OutletOp          outletOp
	OutletQri         outletQri
	OutletSection     outletSection
	OutletSectionMenu outletSectionMenu
	Payment           payment
	Role              role
	RolesNavigation   rolesNavigation
	UserEvent         userEvent
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                db,
		ActorAdmin:        q.ActorAdmin.clone(db),
		Banner:            q.Banner.clone(db),
		BannerAd:          q.BannerAd.clone(db),
		Cart:              q.Cart.clone(db),
		CartItem:          q.CartItem.clone(db),
		Channel:           q.Channel.clone(db),
		Config:            q.Config.clone(db),
		Cuisine:           q.Cuisine.clone(db),
		Customer:          q.Customer.clone(db),
		DataEnv:           q.DataEnv.clone(db),
		Inbox:             q.Inbox.clone(db),
		InboxCustomer:     q.InboxCustomer.clone(db),
		Location:          q.Location.clone(db),
		LocationCode:      q.LocationCode.clone(db),
		LogKeyword:        q.LogKeyword.clone(db),
		LogRecent:         q.LogRecent.clone(db),
		MenuFavorite:      q.MenuFavorite.clone(db),
		MenuStock:         q.MenuStock.clone(db),
		Merchant:          q.Merchant.clone(db),
		MerchantKyc:       q.MerchantKyc.clone(db),
		Navigation:        q.Navigation.clone(db),
		Order:             q.Order.clone(db),
		OrderFlow:         q.OrderFlow.clone(db),
		OrderItem:         q.OrderItem.clone(db),
		OrderRating:       q.OrderRating.clone(db),
		OrderRefund:       q.OrderRefund.clone(db),
		Outlet:            q.Outlet.clone(db),
		OutletChannel:     q.OutletChannel.clone(db),
		OutletCode:        q.OutletCode.clone(db),
		OutletCuisine:     q.OutletCuisine.clone(db),
		OutletFavorite:    q.OutletFavorite.clone(db),
		OutletLogName:     q.OutletLogName.clone(db),
		OutletMenu:        q.OutletMenu.clone(db),
		OutletOp:          q.OutletOp.clone(db),
		OutletQri:         q.OutletQri.clone(db),
		OutletSection:     q.OutletSection.clone(db),
		OutletSectionMenu: q.OutletSectionMenu.clone(db),
		Payment:           q.Payment.clone(db),
		Role:              q.Role.clone(db),
		RolesNavigation:   q.RolesNavigation.clone(db),
		UserEvent:         q.UserEvent.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                db,
		ActorAdmin:        q.ActorAdmin.replaceDB(db),
		Banner:            q.Banner.replaceDB(db),
		BannerAd:          q.BannerAd.replaceDB(db),
		Cart:              q.Cart.replaceDB(db),
		CartItem:          q.CartItem.replaceDB(db),
		Channel:           q.Channel.replaceDB(db),
		Config:            q.Config.replaceDB(db),
		Cuisine:           q.Cuisine.replaceDB(db),
		Customer:          q.Customer.replaceDB(db),
		DataEnv:           q.DataEnv.replaceDB(db),
		Inbox:             q.Inbox.replaceDB(db),
		InboxCustomer:     q.InboxCustomer.replaceDB(db),
		Location:          q.Location.replaceDB(db),
		LocationCode:      q.LocationCode.replaceDB(db),
		LogKeyword:        q.LogKeyword.replaceDB(db),
		LogRecent:         q.LogRecent.replaceDB(db),
		MenuFavorite:      q.MenuFavorite.replaceDB(db),
		MenuStock:         q.MenuStock.replaceDB(db),
		Merchant:          q.Merchant.replaceDB(db),
		MerchantKyc:       q.MerchantKyc.replaceDB(db),
		Navigation:        q.Navigation.replaceDB(db),
		Order:             q.Order.replaceDB(db),
		OrderFlow:         q.OrderFlow.replaceDB(db),
		OrderItem:         q.OrderItem.replaceDB(db),
		OrderRating:       q.OrderRating.replaceDB(db),
		OrderRefund:       q.OrderRefund.replaceDB(db),
		Outlet:            q.Outlet.replaceDB(db),
		OutletChannel:     q.OutletChannel.replaceDB(db),
		OutletCode:        q.OutletCode.replaceDB(db),
		OutletCuisine:     q.OutletCuisine.replaceDB(db),
		OutletFavorite:    q.OutletFavorite.replaceDB(db),
		OutletLogName:     q.OutletLogName.replaceDB(db),
		OutletMenu:        q.OutletMenu.replaceDB(db),
		OutletOp:          q.OutletOp.replaceDB(db),
		OutletQri:         q.OutletQri.replaceDB(db),
		OutletSection:     q.OutletSection.replaceDB(db),
		OutletSectionMenu: q.OutletSectionMenu.replaceDB(db),
		Payment:           q.Payment.replaceDB(db),
		Role:              q.Role.replaceDB(db),
		RolesNavigation:   q.RolesNavigation.replaceDB(db),
		UserEvent:         q.UserEvent.replaceDB(db),
	}
}

type queryCtx struct {
	ActorAdmin        IActorAdminDo
	Banner            IBannerDo
	BannerAd          IBannerAdDo
	Cart              ICartDo
	CartItem          ICartItemDo
	Channel           IChannelDo
	Config            IConfigDo
	Cuisine           ICuisineDo
	Customer          ICustomerDo
	DataEnv           IDataEnvDo
	Inbox             IInboxDo
	InboxCustomer     IInboxCustomerDo
	Location          ILocationDo
	LocationCode      ILocationCodeDo
	LogKeyword        ILogKeywordDo
	LogRecent         ILogRecentDo
	MenuFavorite      IMenuFavoriteDo
	MenuStock         IMenuStockDo
	Merchant          IMerchantDo
	MerchantKyc       IMerchantKycDo
	Navigation        INavigationDo
	Order             IOrderDo
	OrderFlow         IOrderFlowDo
	OrderItem         IOrderItemDo
	OrderRating       IOrderRatingDo
	OrderRefund       IOrderRefundDo
	Outlet            IOutletDo
	OutletChannel     IOutletChannelDo
	OutletCode        IOutletCodeDo
	OutletCuisine     IOutletCuisineDo
	OutletFavorite    IOutletFavoriteDo
	OutletLogName     IOutletLogNameDo
	OutletMenu        IOutletMenuDo
	OutletOp          IOutletOpDo
	OutletQri         IOutletQriDo
	OutletSection     IOutletSectionDo
	OutletSectionMenu IOutletSectionMenuDo
	Payment           IPaymentDo
	Role              IRoleDo
	RolesNavigation   IRolesNavigationDo
	UserEvent         IUserEventDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		ActorAdmin:        q.ActorAdmin.WithContext(ctx),
		Banner:            q.Banner.WithContext(ctx),
		BannerAd:          q.BannerAd.WithContext(ctx),
		Cart:              q.Cart.WithContext(ctx),
		CartItem:          q.CartItem.WithContext(ctx),
		Channel:           q.Channel.WithContext(ctx),
		Config:            q.Config.WithContext(ctx),
		Cuisine:           q.Cuisine.WithContext(ctx),
		Customer:          q.Customer.WithContext(ctx),
		DataEnv:           q.DataEnv.WithContext(ctx),
		Inbox:             q.Inbox.WithContext(ctx),
		InboxCustomer:     q.InboxCustomer.WithContext(ctx),
		Location:          q.Location.WithContext(ctx),
		LocationCode:      q.LocationCode.WithContext(ctx),
		LogKeyword:        q.LogKeyword.WithContext(ctx),
		LogRecent:         q.LogRecent.WithContext(ctx),
		MenuFavorite:      q.MenuFavorite.WithContext(ctx),
		MenuStock:         q.MenuStock.WithContext(ctx),
		Merchant:          q.Merchant.WithContext(ctx),
		MerchantKyc:       q.MerchantKyc.WithContext(ctx),
		Navigation:        q.Navigation.WithContext(ctx),
		Order:             q.Order.WithContext(ctx),
		OrderFlow:         q.OrderFlow.WithContext(ctx),
		OrderItem:         q.OrderItem.WithContext(ctx),
		OrderRating:       q.OrderRating.WithContext(ctx),
		OrderRefund:       q.OrderRefund.WithContext(ctx),
		Outlet:            q.Outlet.WithContext(ctx),
		OutletChannel:     q.OutletChannel.WithContext(ctx),
		OutletCode:        q.OutletCode.WithContext(ctx),
		OutletCuisine:     q.OutletCuisine.WithContext(ctx),
		OutletFavorite:    q.OutletFavorite.WithContext(ctx),
		OutletLogName:     q.OutletLogName.WithContext(ctx),
		OutletMenu:        q.OutletMenu.WithContext(ctx),
		OutletOp:          q.OutletOp.WithContext(ctx),
		OutletQri:         q.OutletQri.WithContext(ctx),
		OutletSection:     q.OutletSection.WithContext(ctx),
		OutletSectionMenu: q.OutletSectionMenu.WithContext(ctx),
		Payment:           q.Payment.WithContext(ctx),
		Role:              q.Role.WithContext(ctx),
		RolesNavigation:   q.RolesNavigation.WithContext(ctx),
		UserEvent:         q.UserEvent.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
