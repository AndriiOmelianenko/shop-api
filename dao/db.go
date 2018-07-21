package dao

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/AndriiOmelianenko/shop-api/models"
	"github.com/icrowley/fake"
	"github.com/urfave/cli"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ShopDAO struct {
	Server   string
	Database string
}

var DB *mgo.Database

const (
	COLLECTION_CATEGORIES = "categories"
	COLLECTION_ITEMS      = "items"
	COLLECTION_ORDERS     = "orders"
	COLLECTION_ORDEREDS   = "oredereds"
)

func (m *ShopDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		fmt.Println("error connecting to mongodb:", err)
	}
	DB = session.DB(m.Database)
}

//func (m *ShopDAO) isDatabase(session *mgo.Session) bool {
//	databases, err := session.DatabaseNames()
//	if err != nil {
//		fmt.Println("error getting db names", err)
//	}
//	for _, database := range databases {
//		if database == m.Database {
//			return true
//		}
//	}
//	fmt.Printf("database %v not found\n", m.Database)
//	return false
//}

func SeedDatabase(c *cli.Context) error {
	fmt.Println("Seed database with random values...")
	if DB == nil {
		mongodb := ShopDAO{Server: c.GlobalString("mongo"), Database: c.GlobalString("dbname")}
		mongodb.Connect()
	}

	// just because
	err := DB.DropDatabase()
	if err != nil {
		fmt.Println("error dropping database")
	}

	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	firstLevelCategoriesNum := 5 + rand.Intn(11)
	firstLevelCategories := generateFirstLevelCategories(firstLevelCategoriesNum)

	secondLevelCategories := models.Categories{}
	for _, v := range firstLevelCategories {
		secondLevelCategoriesNum := 5 + rand.Intn(16) // in each first level category //
		secondLevelCategories = append(secondLevelCategories, generateSecondLevelCategories(secondLevelCategoriesNum, v)...)
	}
	categories := append(firstLevelCategories, secondLevelCategories...)

	b := DB.C(COLLECTION_CATEGORIES).Bulk()
	for _, v := range categories {
		b.Insert(v)
	}
	_, err = b.Run()
	if err != nil {
		fmt.Println("error seeding categories collection", err)
	}

	totalItems := models.Items{}
	for _, v := range categories {
		wg.Add(1)
		itemsNum := 50 + rand.Intn(101) // in each category
		items := generateItems(itemsNum, v)
		totalItems = append(totalItems, items...)
		go func() {
			defer wg.Done()
			b := DB.C(COLLECTION_ITEMS).Bulk()
			for _, v := range items {
				b.Insert(v)
			}
			_, err := b.Run()
			if err != nil {
				fmt.Println("error seeding items records:", err)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Generated first Level Categories ", len(firstLevelCategories))
	fmt.Println("Generated second Level Categories", len(secondLevelCategories))
	fmt.Println("Generated total Categories", len(categories))
	fmt.Println("Generated items ", len(totalItems))
	return nil
}

func generateFirstLevelCategories(number int) models.Categories {
	categories := models.Categories{}
	for i := 0; i < number; i++ {
		////newUUID, err := uuid.NewV4()
		//if err != nil {
		//	fmt.Println("error getting new uuid:", err)
		//}
		category := models.Category{
			ID:          bson.NewObjectId(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Alias:       fake.ProductName(),
			Title:       fake.Title(),
			Description: fake.ProductName(),
			Logo:        fake.ProductName(),
			//ParentID:  uuid.UUID{},
		}
		categories = append(categories, category)
	}
	return categories
}

func generateSecondLevelCategories(number int, category models.Category) models.Categories {
	categories := models.Categories{}
	for i := 0; i < number; i++ {
		category := models.Category{
			ID:          bson.NewObjectId(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Alias:       fake.ProductName(),
			Title:       fake.Title(),
			Description: fake.ProductName(),
			Logo:        fake.ProductName(),
			//ParentID:  category.ID,
		}
		categories = append(categories, category)
	}
	return categories
}

func generateItems(number int, category models.Category) models.Items {
	items := models.Items{}
	for i := 0; i < number; i++ {
		item := models.Item{
			ID:          bson.NewObjectId(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Alias:       fake.ProductName(),
			Title:       fake.Product(),
			Description: fake.ProductName(),
			Pictures:    fake.ProductName(),
			Price:       rand.Intn(1000),
			Count:       rand.Intn(100),
			//CategoryID:  category.ID,
		}
		items = append(items, item)
	}
	return items
}

//
//func generateOrders(number int) models.Orders {
//	orders := models.Orders{}
//	for i := 0; i < number; i++ {
//		order := models.Order{
//			ID:        i,
//			CreatedAt: time.Time{},
//			UpdatedAt: time.Time{},
//			Status: func() string {
//				decision := rand.Intn(3)
//				if decision == 0 {
//					return "Completed"
//				} else if decision == 1 {
//					return "Created"
//				}
//				return "Shipped"
//			}(),
//			Sum: rand.Intn(1000),
//		}
//		orders = append(orders, order)
//	}
//	return orders
//}
//
//func generateOrdereds(number int, order models.Order, item models.Item) models.Ordereds {
//	ordereds := models.Ordereds{}
//	for i := 0; i < number; i++ {
//		ordered := models.Ordered{
//			ID:        i,
//			CreatedAt: time.Time{},
//			UpdatedAt: time.Time{},
//			OrderID:   order.ID,
//			ItemID:    item.ID,
//			ItemCnt:   1,
//			ItemSum:   100,
//		}
//		ordereds = append(ordereds, ordered)
//	}
//	return ordereds
//}
//
