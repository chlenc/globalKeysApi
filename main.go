package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type App struct {
	db *sql.DB
}

type THotel struct {
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

type TCity struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Photo  string `json:"photo"`
	Offers int    `json:"offers"`
	Rooms  int    `json:"rooms"`
}

type TBooking struct {
	Id            int     `json:"id"`
	StartDatetime string  `json:"startDatetime"`
	EndDatetime   string  `json:"endDatetime"`
	Cost          float64 `json:"cost"`
	HotelId       int     `json:"hotelId"`
	RoomId        int     `json:"roomId"`
	CustomerId    int     `json:"customerId"`
}

type TRooms struct {
	Id          int    `json:"id"`
	Room        int    `json:"room"`
	Persons     int    `json:"persons"`
	Floor       int    `json:"floor"`
	Housing     int    `json:"housing"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	CityId      int    `json:"city_id"`
	HotelId     int    `json:"hotel_id"`
}

func main() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	db, err := sql.Open("postgres", dbUri)
	//db, err := sql.Open("postgres", "host=localhost user=postgres dbname=postgres port="+"5432"+" sslmode=disable")
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
		apiRoutes.GET("/room", app.getRooms)
		apiRoutes.GET("/booking", app.getBookings)
		apiRoutes.POST("/booking", app.addBooking)

	}

	r.NoRoute(func(c *gin.Context) {
		render(c, gin.H{"payload": "not found"})
	})

}

func (app *App) getBookings(c *gin.Context) {
	var items []*TBooking

	booking, isbooking := c.GetQuery("id")
	hotel, ishotel := c.GetQuery("hotel")

	if isbooking {
		data := &TBooking{}
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
		var rows *sql.Rows
		var err error
		if ishotel {
			rows, err = app.db.Query("SELECT * FROM bookings where hotel_id = $1 ORDER BY id asc", hotel)
		} else {
			rows, err = app.db.Query("SELECT * FROM bookings ORDER BY id asc")
		}
		if isError(err, "Не удалось получить данные о брони") {
			render(c, gin.H{"payload": "not found"})
		}
		defer rows.Close()
		for rows.Next() {
			data := &TBooking{}
			err := rows.Scan(&data.Id, &data.StartDatetime, &data.EndDatetime, &data.Cost, &data.HotelId, &data.RoomId,
				&data.CustomerId)
			if err == nil {
				items = append(items, data)
			} else {
				log.Println(err)
			}
		}
	}

	render(c, gin.H{"payload": items})
}

func (app *App) addBooking(c *gin.Context) {

	decoder := json.NewDecoder(c.Request.Body)
	var res TBooking
	err := decoder.Decode(&res)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid data, err: "+err.Error())
		return
	}

	log.Println(res.CustomerId)
	sqlStatement := `INSERT INTO bookings (start_datetime, end_datetime, cost, hotel_id, room_id, customer_id)
 	 				 VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	id := 0
	err = app.db.QueryRow(
		sqlStatement,
		res.StartDatetime,
		res.EndDatetime,
		res.Cost,
		res.HotelId,
		res.RoomId,
		res.CustomerId,
	).Scan(&id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "database error, err: "+err.Error())
		return
	}
}

func (app *App) getHotels(c *gin.Context) {

	var items []*THotel

	id, isId := c.GetQuery("id")
	city, isCity := c.GetQuery("city")
	if isId {
		data := &THotel{}
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
			data := &THotel{}
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

func (app *App) getRooms(c *gin.Context) {

	var items []*TRooms

	id, isId := c.GetQuery("id")
	hotel, isHotel := c.GetQuery("hotel")
	if isId {
		data := &TRooms{}
		app.db.QueryRow("select * from rooms where id = $1", id).Scan(
			&data.Id,
			&data.Room,
			&data.Persons,
			&data.Floor,
			&data.Housing,
			&data.Description,
			&data.Price,
			&data.CityId,
			&data.HotelId,
		)
		items = append(items, data)
	} else {
		var rows *sql.Rows
		var err error
		if isHotel {
			rows, err = app.db.Query("SELECT * FROM rooms where hotel_id = $1 ORDER BY id asc", hotel)
		} else {
			rows, err = app.db.Query("SELECT * FROM rooms ORDER BY id asc")
		}
		if isError(err, "Не удалось получить данные о комнатах") {
			render(c, gin.H{"payload": "not found"})
		}
		defer rows.Close()
		for rows.Next() {
			data := &TRooms{}
			err := rows.Scan(
				&data.Id,
				&data.Room,
				&data.Persons,
				&data.Floor,
				&data.Housing,
				&data.Description,
				&data.Price,
				&data.CityId,
				&data.HotelId,
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

	var items []*TCity

	city, isCity := c.GetQuery("id")
	if isCity {
		data := &TCity{}
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
			data := &TCity{}
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
