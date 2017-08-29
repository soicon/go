package dbclient

import (
	"github.com/dmr/microservice/productservice/model"
	"github.com/boltdb/bolt"
	"log"
	"fmt"
	"encoding/json"
	"math/rand"
	"time"
	"strconv"
)

type IBoltClient interface {
	OpenBoltDb()
	QueryBrand() ([]Brand,error)
	QueryProduct(productId string) (model.Product, error)
	QueryAllProduct(brand string) ([]model.Product,error)
	NewProduct(product model.Product) (string,error)
	InitializeBucket()
}
type Brand struct {
	Name string `json:"name"`
}

// Real implementation
type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("products.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

}



// Creates an "AccountBucket" in our BoltDB. It will overwrite any existing bucket of the same name.
func (bc *BoltClient) InitializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("ProductBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

func GenerateId() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10000000-1)+1
}


// Seed (n) make-believe account objects into the AcountBucket bucket.
func (bc *BoltClient) NewProduct(product model.Product) (string,error) {

		key := product.Product_Brand+strconv.Itoa(GenerateId())

		// Create an instance of our Account struct
		acc := model.Product{
			key,
			product.Product_Name,
			product.Product_Brand,
			product.Product_Img,
			product.Product_Price,
		}

		// Serialize the struct to JSON
		jsonBytes, _ := json.Marshal(acc)

		// Write the data to the AccountBucket
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("ProductBucket"))
			err := b.Put([]byte(key), jsonBytes)
			if err != nil {
				return err
			}
			return nil
		})
		return key,nil
}

func RemoveDuplicate(list []Brand) []Brand{

	var temp []Brand
	temp = append(temp,list[0])
	for _,brand := range list{
		var isDup = false
		for _, temp := range temp{
			if temp.Name == brand.Name {
				isDup = true
			}
		}
		if !isDup{
			temp= append(temp, brand)
		}
	}
	return temp
}
func (bc *BoltClient) QueryBrand() ([]Brand,error) {
	var list []Brand
	var list_product []model.Product
	err:= bc.boltDB.View(func(tx *bolt.Tx) error {
		 b:= tx.Bucket([]byte("ProductBucket"))
		b.ForEach(func(k, v []byte) error {
			temp := model.Product{}
			json.Unmarshal(v,&temp)
			list_product = append(list_product, temp)
			return nil
		})
		for _, product := range list_product{
			list = append(list,Brand{product.Product_Brand})

		}
		list = RemoveDuplicate(list)
		if list ==  nil{
			return fmt.Errorf("No Brand Found")
		}
		return nil
	})
	if err != nil {
		return list,err
	}
	return list,nil
}

func (bc *BoltClient)QueryAllProduct(brand string) ([]model.Product,error){
	var list_product []model.Product
	err:= bc.boltDB.View(func(tx *bolt.Tx) error {
		b:= tx.Bucket([]byte("ProductBucket"))
		b.ForEach(func(k, v []byte) error {
			temp := model.Product{}
			json.Unmarshal(v,&temp)
			if temp.Product_Brand == brand {
				list_product = append(list_product, temp)
			}
			return nil
		})


		return nil
	})
	if err != nil {
		return list_product,err
	}
	return list_product,nil
}

func (bc *BoltClient) QueryProduct(productId string) (model.Product, error) {
	// Allocate an empty Account instance we'll let json.Unmarhal populate for us in a bit.
	account := model.Product{}

	// Read an object from the bucket using boltDB.View
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Read the bucket from the DB
		b := tx.Bucket([]byte("ProductBucket"))

		// Read the value identified by our accountId supplied as []byte
		accountBytes := b.Get([]byte(productId))
		if accountBytes == nil {
			return fmt.Errorf("No product found for " + productId)
		}
		// Unmarshal the returned bytes into the account struct we created at
		// the top of the function
		json.Unmarshal(accountBytes, &account)

		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})
	// If there were an error, return the error
	if err != nil {
		return model.Product{}, err
	}
	// Return the Account struct and nil as error.
	return account, nil
}