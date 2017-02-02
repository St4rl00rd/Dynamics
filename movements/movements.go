package movements

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"fmt"
	// import postgress Driver (pq)
	_ "github.com/lib/pq"
	"github.com/st4rl00rd/dynamics/context"
)

/* MAIN FUNCTIONS */

// ShowCtl is the controller for MovementShow route
func ShowCtl(ctx *context.Context, w http.ResponseWriter, id int) {
	movement := GetMovement(ctx, id)
	if movement.ID == 0 {
		NotValidMovementID(w)
	} else {
		jData, err := json.Marshal(movement)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
		fmt.Println("OK Movement")
	}
}

// IndexCtl returns the data to MovementIndex Handler
func IndexCtl(ctx *context.Context, w http.ResponseWriter) {
	movements := GetMovements(ctx)
	jData, err := json.Marshal(movements)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
	fmt.Println("OK Movements")
}

/* DATABASE */

// GetMovement returns all the properties from a item on your db
func GetMovement(ctx *context.Context, idPrm int) *Movement {
	var m Movement
	idCnv := strconv.Itoa(idPrm)
	fmt.Println("SELECT * FROM movements WHERE id=" + idCnv)
	rows, err := ctx.Db.Query("SELECT * FROM movements WHERE id=" + idCnv)

	if err != nil {
		log.Fatal(err)
	}
	err = nil
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&m.ID, &m.Details, &m.SchemeID, &m.ProductID, &m.Quantity, &m.StoreFromID, &m.StoreToID, &m.UserFromID, &m.UserToID, &m.Ended, &m.CreatedAT, &m.UpdatedAT); err != nil {
			log.Fatal(err)
		}
	}
	if err != nil {
		return &Movement{}
	}
	return &m
}

// GetMovements returns all items from a table on your db
func GetMovements(ctx *context.Context) JSONResponse {
	response := JSONResponse{}
	fmt.Println("SELECT * FROM movements")
	movementsDB, err := ctx.Db.Query("SELECT * FROM movements")
	if err != nil {
		log.Fatal(err)
	}
	for movementsDB.Next() {
		err = nil
		var m Movement
		if err := movementsDB.Scan(&m.ID, &m.Details, &m.SchemeID, &m.ProductID, &m.Quantity, &m.StoreFromID, &m.StoreToID, &m.UserFromID, &m.UserToID, &m.Ended, &m.CreatedAT, &m.UpdatedAT); err != nil {
			log.Fatal(err)
		} else {
			response.InsertMovement(InsertProduct(ctx, m))
		}
	}

	return response
}

// InsertMovement Function
func (m *JSONResponse) InsertMovement(item Movement) []Movement {
	m.Movements = append(m.Movements, item)
	return m.Movements
}

// InsertProduct function
func InsertProduct(ctx *context.Context, m Movement) Movement {
	productID := strconv.FormatInt(m.ProductID, 10)
	product, err := ctx.Db.Query("SELECT * FROM products WHERE id =" + productID)
	if err != nil {
		fmt.Println(err)
	}
	for product.Next() {
		if errp := product.Scan(&m.Product.ID, &m.Product.SchemeID, &m.Product.Name, &m.Product.Details, &m.Product.CreatedAT, &m.Product.UpdatedAT); errp != nil {
			fmt.Println(errp)
		}
	}
	movement := InsertLastMovement(ctx, m)
	return movement
}

// InsertLastMovement function
func InsertLastMovement(ctx *context.Context, m Movement) Movement {
	productID := strconv.FormatInt(m.ProductID, 10)
	movement, err := ctx.Db.Query("SELECT * FROM movements WHERE product_id =" + productID)
	if err != nil {
		log.Fatal(err)
	}
	for movement.Next() {
		if errp := movement.Scan(&m.Product.LastMovement.ID, &m.Product.LastMovement.Details, &m.Product.LastMovement.SchemeID, &m.Product.LastMovement.ProductID, &m.Product.LastMovement.Quantity, &m.Product.LastMovement.StoreFromID, &m.Product.LastMovement.StoreToID, &m.Product.LastMovement.UserFromID, &m.Product.LastMovement.UserToID, &m.Product.LastMovement.Ended, &m.Product.LastMovement.CreatedAT, &m.Product.LastMovement.UpdatedAT); errp != nil {
			log.Fatal(errp)
		}
	}
	return m
}

/* SPECIAL FUNCTIONS */

// NotValidMovementID function resolves the http.ResponseWriter when Movement is not found in the DB
func NotValidMovementID(w http.ResponseWriter) {
	jData, err := json.Marshal("Not Valid Movement ID")
	println(string(jData))
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
