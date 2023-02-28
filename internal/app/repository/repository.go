package repository

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lab3/internal/app/ds"
)

type Repository struct {
	db *gorm.DB
}
type P struct {
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

func (r *Repository) GetAllProducts() ([]ds.Goods, error) {
	var products []ds.Goods
	result := r.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *Repository) GetProductByID(id uint) (*ds.Goods, error) {
	product := &ds.Goods{}
	err := r.db.First(product, "id_good = ?", id).Error // find product with code D42
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *Repository) CreateProduct(product *ds.Goods) error {
	err := r.db.Create(product).Error
	return err
}

func (r *Repository) ChangeProduct(product ds.Goods) error {
	db_product := &ds.Goods{}
	err := r.db.First(db_product, "id_good = ?", product.Id_good).Error // find product with code D42
	if err != nil {
		return err
	}
	if product.Price != 0 {
		db_product.Price = product.Price
	}
	if product.Description != "" {
		db_product.Description = product.Description
	}
	if product.Type != "" {
		db_product.Type = product.Type
	}
	if product.Color != "" {
		db_product.Color = product.Color
	}
	if product.Image != "" {
		db_product.Image = product.Image
	}
	if product.Company != "" {
		db_product.Company = product.Company
	}
	err = r.db.Save(&db_product).Error
	return err
}

func (r *Repository) DeleteProduct(id uint) error {
	err := r.db.First(&ds.Goods{}, "id_good = ?", id).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&ds.Goods{}, "id_good = ?", id).Error
	return err
}

//Users

func (r *Repository) CreateUser(user *ds.Users) error {
	err := r.db.Create(user).Error
	return err
}

func (r *Repository) LoginCheck(user *ds.Users) error {
	user_db := ds.Users{}
	err := r.db.Model(&ds.Users{}).Where("login = ?", user.Login).Take(&user_db).Error
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user_db.Password), []byte(user.Password))
	if err != nil {
		return err
	}
	user.Id_user = user_db.Id_user
	user.Role = user_db.Role
	return nil
}

func (r *Repository) CheckLogin(login string) error {
	err := r.db.Model(&ds.Users{}).Where("login = ?", login).Error
	if err != nil {
		return nil
	}
	return err
}

func (r *Repository) GetUserByID(id uint) (*ds.Users, error) {
	user := &ds.Users{}
	err := r.db.First(user, "id_user = ?", id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetIdByLogin(login string) (uint, error) {
	user := &ds.Users{}
	err := r.db.First(user, "login = ?", login).Error
	if err != nil {
		return 0, err
	}
	return user.Id_user, nil
}

func (r *Repository) CreateBasketRow(basket_row *ds.Basket) error {
	err := r.db.Create(basket_row).Error
	return err
}

func (r *Repository) GetBasket(id_user uint) ([]ds.Basket, error) {
	var basket []ds.Basket
	result := r.db.Find(&basket, "id_user = ?", id_user)
	if result.Error != nil {
		return nil, result.Error
	}
	return basket, nil
}

func (r *Repository) GetBasketById(id uint) (ds.Basket, error) {
	var basket ds.Basket
	result := r.db.Find(&basket, "id_row = ?", id)
	if result.Error != nil {
		return basket, result.Error
	}
	return basket, nil
}

func (r *Repository) DeleteBasketRow(basket_row *ds.Basket) error {
	err := r.db.Model(&ds.Basket{}).Where("id_good = ?", basket_row.Id_good, "id_user = ?", basket_row.Id_user).Take(&basket_row).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&ds.Basket{}, "id_row = ?", basket_row.Id_row).Error
	return err
}

func (r *Repository) ChangeQuantity(basket_row *ds.Basket, quantity int) error {
	err := r.db.Model(&ds.Basket{}).Where("id_good = ?", basket_row.Id_good, "id_user = ?", basket_row.Id_user).Take(&basket_row).Error
	if err != nil {
		return err
	}
	err = r.db.Model(&ds.Basket{}).Where("id_good = ?", basket_row.Id_row).Update("quantity", quantity).Error
	return err
}

func (r *Repository) CreateOrder(params *ds.Orders) error {
	err := r.db.Create(params).Error

	return err
}

func (r *Repository) CreateGoodOrder(orderGood *ds.GoodOrder) error {
	err := r.db.Create(orderGood).Error
	return err
}

type ordersStatus struct {
	Id_order    uint
	Date        string
	Name        string
	Total       int
	Description string
	Login       string
}

func (r *Repository) GetOrder(id_user uint) ([]ordersStatus, error) {
	var order []ordersStatus
	err := r.db.Table("orders").Select("*").Joins("JOIN statuses on statuses.id_status = orders.status").Find(&order, "id_user = ?", id_user).Error

	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *Repository) GetAllOrders() ([]ordersStatus, error) {
	var order []ordersStatus
	err := r.db.Table("orders").Select("*").Joins("JOIN statuses on statuses.id_status = orders.status").Joins("JOIN users on users.id_user = orders.id_user").Find(&order).Error

	if err != nil {
		return nil, err
	}
	return order, nil
}

type GoodQuantity struct {
	Id_good     uint   `json:"Id_good"`
	Type        string `json:"type"`
	Company     string `json:"company"`
	Color       string `json:"color"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Quantity    uint   `json:"quantity"`
}

func (r *Repository) GetGoodOrder(id_order uint) ([]GoodQuantity, error) {
	var goods []GoodQuantity
	err := r.db.Table("good_orders").Select("*").Joins("JOIN goods on goods.id_good = good_orders.id_good").Find(&goods, "id_order = ?", id_order).Error

	if err != nil {
		return goods, err
	}
	return goods, nil
}

func (r *Repository) GetAllStatuses() ([]ds.Statuses, error) {
	var products []ds.Statuses
	result := r.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *Repository) ChangeStatus(id_order uint, id_status uint) error {
	err := r.db.Model(&ds.Orders{}).Where("id_order = ?", id_order).Update("status", id_status).Error
	return err
}
