package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	db *sql.DB
}

type Hotel struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Country     string  `json:"country"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Photo       string  `json:"photo"`
	Description string  `json:"description"`
	Stars       int     `json:"stars"`
	CityId      int     `json:"cityId"`
}

type City struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Photo  string `json:"photo"`
	Offers int    `json:"offers"`
	Rooms  int    `json:"rooms"`
}

type Booking struct {
	Id            int     `json:"id"`
	StartDatetime string  `json:"startDatetime"`
	EndDatetime   string  `json:"endDatetime"`
	Cost          float64 `json:"cost"`
	HotelId       int     `json:"hotelId"`
	RoomId        int     `json:"roomId"`
	CustomerId    int     `json:"customerId"`
}

func main() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", dbHost, username, dbName)
	fmt.Println(dbUri)

	db, err := sql.Open("postgres", dbUri)
	//"host=localhost user=postgres dbname=postgres port="+port+" sslmode=disable")
	db.SetMaxOpenConns(10)
	if err != nil {
		log.Println("failed to connect database", err)
		return
	}

	var app = App{db}

	r := gin.Default()
	r.Use(gin.Logger())

	app.initializeRoutes(r)

	r.Run()
}

func (app *App) initializeRoutes(r *gin.Engine) {
	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/city", app.getCities)
		apiRoutes.GET("/hotel", app.getHotels)
		apiRoutes.GET("/booking", app.getBookings)
		apiRoutes.POST("/booking", app.addBooking)

	}

	r.NoRoute(func(c *gin.Context) {
		render(c, gin.H{"payload": "not found"})
	})

}

func (app *App) getBookings(c *gin.Context) {
	var items []*Booking

	booking, isbooking := c.GetQuery("id")
	if isbooking {
		data := &Booking{}
		app.db.QueryRow("select * from bookings where id = $1", booking).Scan(
			&data.Id,
			&data.StartDatetime,
			&data.EndDatetime,
			&data.Cost,
			&data.HotelId,
			&data.RoomId,
			&data.CustomerId,
		)
		items = append(items, data)
	} else {
		rows, err := app.db.Query("SELECT * FROM bookings ORDER BY id asc")
		if isError(err, "Не удалось получить данные о брони") {
			render(c, gin.H{"payload": "not found"})
		}
		defer rows.Close()
		for rows.Next() {
			data := &Booking{}
			err := rows.Scan(&data.Id, &data.StartDatetime, &data.EndDatetime, &data.Cost, &data.HotelId, &data.RoomId,
				&data.CustomerId)
			if err == nil {
				items = append(items, data)
			} else {
				log.Println(err)
			}
		}
	}

	if len(items) == 0 {
		render(c, gin.H{"payload": "not found"})
	} else {
		render(c, gin.H{"payload": items})
	}
}

func (app *App) addBooking(c *gin.Context) {
	layout := "2006-01-02T15:04:05.000Z"
	startDatetime, startDatetimeErr := time.Parse(layout, c.PostForm("start_datetime"))
	endDatetime, endDatetimeErr := time.Parse(layout, c.PostForm("end_datetime"))
	if startDatetimeErr != nil || endDatetimeErr != nil {
		c.JSON(http.StatusBadRequest, "invalid datetime")
		return
	}
	cost := c.PostForm("cost")
	hotelId := c.PostForm("hotel_id")
	roomId := c.PostForm("room_id")
	customerId := c.PostForm("customer_id")
	sqlStatement := `INSERT INTO bookings (start_datetime, end_datetime, cost, hotel_id, room_id, customer_id)
 	  				 VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	id := 0
	var err = app.db.QueryRow(sqlStatement, startDatetime, endDatetime, cost, hotelId, roomId, customerId).Scan(&id)
	if err != nil {
		panic(err)
	}
}

func (app *App) getHotels(c *gin.Context) {

	var items []*Hotel

	id, isId := c.GetQuery("id")
	city, isCity := c.GetQuery("city")
	if isId {
		data := &Hotel{}
		app.db.QueryRow("select * from hotels where id = $1", id).Scan(
			&data.Id,
			&data.Name,
			&data.Country,
			&data.Address,
			&data.Latitude,
			&data.Longitude,
			&data.Photo,
			&data.Description,
			&data.Stars,
			&data.CityId,
		)
		if data.Name != "" {
			items = append(items, data)
		}
	} else {
		var rows *sql.Rows
		var err error
		if isCity {
			rows, err = app.db.Query("SELECT * FROM hotels where city_id = $1 ORDER BY id asc", city)
		} else {
			rows, err = app.db.Query("SELECT * FROM hotels ORDER BY id asc")
		}
		if isError(err, "Не удалось получить данные о городах") {
			render(c, gin.H{"payload": "not found"})
		}
		defer rows.Close()
		for rows.Next() {
			data := &Hotel{}
			err := rows.Scan(
				&data.Id,
				&data.Name,
				&data.Country,
				&data.Address,
				&data.Latitude,
				&data.Longitude,
				&data.Photo,
				&data.Description,
				&data.Stars,
				&data.CityId,
			)
			if err == nil {
				items = append(items, data)
			} else {
				log.Println(err)
			}
		}
	}
	render(c, gin.H{"payload": items})
}

func (app *App) getCities(c *gin.Context) {

	var items []*City

	city, isCity := c.GetQuery("id")
	if isCity {
		data := &City{}
		app.db.QueryRow("select * from cities where id = $1", city).Scan(
			&data.Id,
			&data.Name,
			&data.Photo,
			&data.Offers,
			&data.Rooms,
		)
		if data.Name != "" {
			items = append(items, data)
		}
	} else {
		rows, err := app.db.Query("SELECT * FROM cities ORDER BY id asc")
		if isError(err, "Не удалось получить данные о городах") {
			render(c, gin.H{"payload": "not found"})
		}
		defer rows.Close()
		for rows.Next() {
			data := &City{}
			err := rows.Scan(&data.Id, &data.Name, &data.Photo, &data.Offers, &data.Rooms)
			if err == nil {
				items = append(items, data)
			} else {
				log.Println(err)
			}
		}
	}

	if len(items) == 0 {
		render(c, gin.H{"payload": "not found"})
	} else {
		render(c, gin.H{"payload": items})
	}
}

func isError(err error, str string) bool {
	if err != nil {
		log.Println("Какая-то хуйня! \n"+str+"\n", err.Error())
		return true
	}
	return false
}

func render(c *gin.Context, data gin.H) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	c.JSON(http.StatusOK, data["payload"])
}
