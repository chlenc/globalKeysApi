// main.go

package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

type App struct {
	db *sql.DB
}

type Hotel struct {
	Id          int
	Name        string
	Country     string
	Address     string
	Latitude    float64
	Longitude   float64
	Photo       string
	Description string
	Stars       int
	CityId      int
}

type City struct {
	Id     int
	Name   string
	Photo  string
	Offers int
	Rooms  int
}

func main() {

	db, err := sql.Open("postgres", "host=localhost port=5432 dbname=postgres user=postgres sslmode=disable")
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
		apiRoutes.GET("/city", func(c *gin.Context) {
			data := app.getCities(c)
			if len(data) == 0 {
				render(c, gin.H{"payload": "not found"})
			} else {
				render(c, gin.H{"payload": data})
			}
		})
		apiRoutes.GET("/hotel", func(c *gin.Context) {
			data := app.getHotels(c)
			render(c, gin.H{"payload": data})
		})
	}

	r.NoRoute(func(c *gin.Context) {
		render(c, gin.H{"payload": "not found"})
	})

}

func (app *App) getHotels(c *gin.Context) []*Hotel {

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
			return items
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
	return items
}

func (app *App) getCities(c *gin.Context) []*City {

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
			return items
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
	return items
}

func isError(err error, str string) bool {
	if err != nil {
		log.Println("Какая-то хуйня! \n"+str+"\n", err.Error())
		return true
	}
	return false
}

func render(c *gin.Context, data gin.H) {
	//res.header("Access-Control-Allow-Origin", "*");
	//res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	log.Println(c.GetHeader("Access-Control-Allow-Origin"))
	c.JSON(http.StatusOK, data["payload"])
}
