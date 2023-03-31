package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/staurran/messengerKR.git/internal/app/ds"
)

type Repository struct {
	db *gorm.DB
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetChats(userID uint, k int) ([]ds.Chat, error) {
	var chats []ds.Chat
	subQuery := r.db.Model(ds.Message{}).Select("messages.id, messages.context, messages.user_from, messages.chat_id").Where("message.chat_id = chat.id")
	result := r.db.Find(&chats, "id_user = ?", userID).
		Joins("Join messages on chat.id = messages.chat_id").
		Where("messages.time_created = Max(?)", subQuery).
		Order("message.time_created")
	if result.Error != nil {
		return nil, result.Error
	}
	return chats, nil
}

func (r *Repository) GetLastMes(chatId uint) (lastMessage LastMessage, err error) {
	var message ds.Message
	err = r.db.First(&message, "chat_id = ?", chatId).Error
	if err != nil {
		return
	}

	user, err := r.GetUserByID(message.UserFromID)
	if err != nil {
		return
	}

	lastMessage.Content = message.Context
	lastMessage.UserName = user.Username

	return lastMessage, nil
}

func (r *Repository) GetQuantityShownMes(chatId, userId uint) (int64, error) {
	var count int64
	subQuery := r.db.Model(&ds.Message{}).Where("chat_is = ?", chatId)
	err := r.db.Model(&ds.Shown{}).Where("user_id = ?, message_id IN (?)", userId, subQuery).Count(&count).Error
	return count, err
}

//func (r *Repository) GetAllProducts() ([]ds.Goods, error) {
//	var products []ds.Goods
//	result := r.db.Find(&products)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//	return products, nil
//}
//
//func (r *Repository) GetProductByID(id uint) (*ds.Goods, error) {
//	product := &ds.Goods{}
//	err := r.db.First(product, "id_good = ?", id).Error // find product with code D42
//	if err != nil {
//		return nil, err
//	}
//	return product, nil
//}
//
//func (r *Repository) CreateProduct(product *ds.Goods) error {
//	err := r.db.Create(product).Error
//	return err
//}
//
//func (r *Repository) ChangeProduct(product ds.Goods) error {
//	db_product := &ds.Goods{}
//	err := r.db.First(db_product, "id_good = ?", product.Id_good).Error // find product with code D42
//	if err != nil {
//		return err
//	}
//	if product.Price != 0 {
//		db_product.Price = product.Price
//	}
//	if product.Description != "" {
//		db_product.Description = product.Description
//	}
//	if product.Type != "" {
//		db_product.Type = product.Type
//	}
//	if product.Color != "" {
//		db_product.Color = product.Color
//	}
//	if product.Image != "" {
//		db_product.Image = product.Image
//	}
//	if product.Company != "" {
//		db_product.Company = product.Company
//	}
//	err = r.db.Save(&db_product).Error
//	return err
//}
//
//func (r *Repository) DeleteProduct(id uint) error {
//	err := r.db.First(&ds.Goods{}, "id_good = ?", id).Error
//	if err != nil {
//		return err
//	}
//	err = r.db.Delete(&ds.Goods{}, "id_good = ?", id).Error
//	return err
//}

func (r *Repository) CreateChat(chat *ds.Chat) error {
	err := r.db.Create(chat).Error
	return err
}

func (r *Repository) DeleteChat(chatId uint, userId uint) error {
	err := r.db.First(&ds.Chat{}, "id_chat = ?", chatId).Error

	return err
}

//Users

func (r *Repository) CreateChatUser(uchat *ds.ChatUser) error {
	err := r.db.Create(uchat).Error
	return err
}

//
//func (r *Repository) CreateBasketRow(basket_row *ds.Basket) error {
//	err := r.db.Create(basket_row).Error
//	return err
//}
//
//func (r *Repository) GetBasket(id_user uint) ([]ds.Basket, error) {
//	var basket []ds.Basket
//	result := r.db.Find(&basket, "id_user = ?", id_user)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//	return basket, nil
//}
//
//func (r *Repository) GetBasketById(id uint) (ds.Basket, error) {
//	var basket ds.Basket
//	result := r.db.Find(&basket, "id_row = ?", id)
//	if result.Error != nil {
//		return basket, result.Error
//	}
//	return basket, nil
//}
//
//func (r *Repository) DeleteBasketRow(basket_row *ds.Basket) error {
//	err := r.db.Model(&ds.Basket{}).Where("id_good = ?", basket_row.Id_good, "id_user = ?", basket_row.Id_user).Take(&basket_row).Error
//	if err != nil {
//		return err
//	}
//	err = r.db.Delete(&ds.Basket{}, "id_row = ?", basket_row.Id_row).Error
//	return err
//}
//
//func (r *Repository) ChangeQuantity(basket_row *ds.Basket, quantity int) error {
//	err := r.db.Model(&ds.Basket{}).Where("id_good = ?", basket_row.Id_good, "id_user = ?", basket_row.Id_user).Take(&basket_row).Error
//	if err != nil {
//		return err
//	}
//	err = r.db.Model(&ds.Basket{}).Where("id_good = ?", basket_row.Id_row).Update("quantity", quantity).Error
//	return err
//}
//
//func (r *Repository) CreateOrder(params *ds.Orders) error {
//	err := r.db.Create(params).Error
//
//	return err
//}
//
//func (r *Repository) CreateGoodOrder(orderGood *ds.GoodOrder) error {
//	err := r.db.Create(orderGood).Error
//	return err
//}
//
//type ordersStatus struct {
//	Id_order    uint
//	Date        string
//	Name        string
//	Total       int
//	Description string
//	Login       string
//}
//
//func (r *Repository) GetOrder(id_user uint) ([]ordersStatus, error) {
//	var order []ordersStatus
//	err := r.db.Table("orders").Select("*").Joins("JOIN statuses on statuses.id_status = orders.status").Find(&order, "id_user = ?", id_user).Error
//
//	if err != nil {
//		return nil, err
//	}
//	return order, nil
//}
//
//func (r *Repository) GetAllOrders() ([]ordersStatus, error) {
//	var order []ordersStatus
//	err := r.db.Table("orders").Select("*").Joins("JOIN statuses on statuses.id_status = orders.status").Joins("JOIN users on users.id_user = orders.id_user").Find(&order).Error
//
//	if err != nil {
//		return nil, err
//	}
//	return order, nil
//}
//
//type GoodQuantity struct {
//	Id_good     uint   `json:"Id_good"`
//	Type        string `json:"type"`
//	Company     string `json:"company"`
//	Color       string `json:"color"`
//	Price       uint   `json:"price"`
//	Description string `json:"description"`
//	Image       string `json:"image"`
//	Quantity    uint   `json:"quantity"`
//}
//
//func (r *Repository) GetGoodOrder(id_order uint) ([]GoodQuantity, error) {
//	var goods []GoodQuantity
//	err := r.db.Table("good_orders").Select("*").Joins("JOIN goods on goods.id_good = good_orders.id_good").Find(&goods, "id_order = ?", id_order).Error
//
//	if err != nil {
//		return goods, err
//	}
//	return goods, nil
//}
//
//func (r *Repository) GetAllStatuses() ([]ds.Statuses, error) {
//	var products []ds.Statuses
//	result := r.db.Find(&products)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//	return products, nil
//}
//
//func (r *Repository) ChangeStatus(id_order uint, id_status uint) error {
//	err := r.db.Model(&ds.Orders{}).Where("id_order = ?", id_order).Update("status", id_status).Error
//	return err
//}
