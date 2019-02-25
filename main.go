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

//type rooms struct {
//	HotelId     string         `json:"hotel_id"`
//	Busy        pq.StringArray `json:"busy" gorm:"type:varchar(64)[]"`
//	Persons     int            `json:"persons"`
//	Floor       string         `json:"floor"`
//	Housing     string         `json:"housing"`
//	Room        string         `json:"room"`
//	Description string         `json:"description"`
//	Price       int            `json:"price"`
//}
//type users struct {
//	Phone    string `json:"phone"`
//	Mail     string `json:"mail"`
//	Name     string `json:"name"`
//	Passport string `json:"passport"`
//	Password string `json:"password"`
//}

//+
type City struct {
	Id     int
	Name   string
	Photo  string
	Offers int
	Rooms  int
}

func main() {
	//db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=postgres user=alesenka sslmode=disable")

	//url := os.Getenv("DATABASE_URL")
	//connection, _ := pq.ParseURL(url)
	//connection += " sslmode=require"
	//db, err := gorm.Open("postgres", connection)

	db, err := sql.Open("postgres", "host=localhost port=5432 dbname=postgres user=alesenka sslmode=disable")
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

//routes.go

func (app *App) initializeRoutes(r *gin.Engine) {
	r.GET("/city", func(c *gin.Context) {
		data := app.getCities(c)
		if len(data) == 0 {
			render(c, gin.H{"payload": "not found"})
		} else {
			render(c, gin.H{"payload": data})
		}
	})
	r.GET("/hotel", func(c *gin.Context) {
		data := app.getHotels(c)
		//if len(data) == 0 {
		//	render(c, gin.H{"payload": "not found"})
		//} else {
			render(c, gin.H{"payload": data})
		//}
	})

	//r.Use(setUserStatus())
	//r.GET("/", app.showTopCities)
	//r.GET("/search", app.search)
	//r.GET("/login", func(c *gin.Context) {
	//	if !loggedInCheck(c) {
	//		render(c, gin.H{"IsLogin": false}, "login.html")
	//	}
	//})
	//r.POST("/login", app.performLogin)
	//
	//r.GET("/registration", func(c *gin.Context) {
	//	if !loggedInCheck(c) {
	//		render(c, gin.H{"IsLogin": false}, "registration.html")
	//	}
	//})
	//r.POST("/registration", app.register)
	//
	//r.GET("/likes", func(c *gin.Context) {
	//	loggedInInterface, _ := c.Get("is_logged_in")
	//	loggedIn := loggedInInterface.(bool)
	//	render(c, gin.H{"IsLogin": loggedIn}, "likes.html")
	//})
	//r.POST("/likes", app.doLike)
	//
	//r.GET("/logout", logout)
	//
	//r.POST("/checkUsername", app.checkUsername)

	//====
	r.NoRoute(func(c *gin.Context) {
		render(c, gin.H{"payload": "not found"})
	})

}

//func (app *App) search(c *gin.Context) {
//
//	var out = map[string]interface{}{}
//	var reqSelector = ""
//	city, isCity := c.GetQuery("city")
//	date, isDate := c.GetQuery("date")
//	persons, isPersons := c.GetQuery("persons")
//
//	if isCity || isDate || isPersons {
//		reqSelector += " WHERE "
//	}
//
//	if isCity {
//		var i, _ = strconv.Atoi(city)
//		out["CitiesList"] = app.getCities(i)
//		reqSelector += " city_id = $1"
//	} else {
//		out["CitiesList"] = app.getCities(0)
//	}
//	if isDate {
//		out["Date"] = date
//	}
//	if isPersons {
//		out["Persons"] = persons
//	}
//
//	var items []*Hotel
//
//	var req = "SELECT id,name,address,photo,stars FROM Hotel" + reqSelector //and min price of room
//	rows, err := app.db.Query(req, city)
//	if isError(err, "Не удалось получить данные о отелях") {
//		return
//	}
//	defer rows.Close()
//	for rows.Next() {
//		data := &Hotel{}
//		err := rows.Scan(&data.Id, &data.Name, &data.Address, &data.Photo, &data.Stars)
//		minPrice := 0
//		app.db.QueryRow("SELECT MIN(price) from rooms where hotel_id = $1", data.Id).Scan(&minPrice)
//		log.Println(minPrice)
//		data.MinPrice = minPrice
//		if err == nil {
//			items = append(items, data)
//		}
//		//app.getOffersAndRooms(data)
//	}
//
//	out["IsLogin"] = false
//	out["IsHotel"] = true
//	out["Hotel"] = items
//
//	render(c, out, "home.html")
//
//}
//
//func (app *App) showTopCities(c *gin.Context) {
//	var items []*City
//
//	rows, err := app.db.Query("SELECT id,name,photo FROM cities ORDER BY id asc LIMIT 4")
//	if isError(err, "Не удалось получить данные о городах") {
//		return
//	}
//	defer rows.Close()
//	for rows.Next() {
//		data := &City{}
//		err := rows.Scan(&data.Id, &data.Name, &data.Photo)
//		if err == nil {
//			items = append(items, data)
//		}
//		app.getOffersAndRooms(data)
//	}
//	render(c, gin.H{
//		"IsLogin":    false,
//		"CitiesList": app.getCities(0), //
//		"IsHotel":   false,
//		"Cities":     items},
//		"home.html")
//}

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
			rows, err = app.db.Query("SELECT * FROM hotels where city_id = $1 ORDER BY id asc",city)
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

//func (app *App) getOffersAndRooms(data *City) {
//	app.db.QueryRow("select count(Hotel.id) from Hotel where Hotel.city_id = $1", data.Id).Scan(&data.Offers)
//	app.db.QueryRow("select count(rooms.id) from rooms where rooms.city_id = $1", data.Id).Scan(&data.Rooms)
//}
//
func isError(err error, str string) bool {
	if err != nil {
		log.Println("Какая-то хуйня! \n"+str+"\n", err.Error())
		return true
	}
	return false
}

//func loggedInCheck(c *gin.Context) bool {
//	loggedInInterface, err := c.Get("is_logged_in")
//	if !err {
//		return false
//	}
//	loggedIn := loggedInInterface.(bool)
//	if loggedIn {
//		render(c, gin.H{
//			"IsError": false,
//			"IsLogin": loggedIn,
//		}, "profile.html")
//		return true
//	}
//
//	return false
//}

func render(c *gin.Context, data gin.H) {
	//loggedInInterface, _ := c.Get("is_logged_in")
	//data["is_logged_in"] = loggedInInterface.(bool)

	c.JSON(http.StatusOK, data["payload"])

	//switch c.Request.Header.Get("Accept") {
	//case "application/json":
	//	//Respond with JSON
	//	c.JSON(http.StatusOK, data["payload"])
	//case "application/xml":
	//	//Respond with XML
	//	c.XML(http.StatusOK, data["payload"])
	//default:
	//	//Respond with HTML
	//	c.HTML(http.StatusOK, templateName, data)
	//
	//}
}

//
//func logout(c *gin.Context) {
//	c.SetCookie("token", "", -1, "", "", false, true)
//	c.Redirect(http.StatusTemporaryRedirect, "/")
//}
//
//func (app *App) performLogin(c *gin.Context) {
//	username := c.PostForm("Username")
//	password := c.PostForm("Password")
//
//	var temp user
//	app.db.First(&temp, "username = ?", username)
//	if (temp != user{}) && (temp.Password == password) {
//
//		token := generateSessionToken()
//		c.SetCookie("token", token, 3600, "", "", false, true)
//		c.Set("is_logged_in", true)
//
//		render(c, gin.H{"IsLogin": true}, "profile.html")
//
//	} else {
//		c.HTML(http.StatusBadRequest, "login.html", gin.H{"IsError": true})
//	}
//}
//func (app *App) register(c *gin.Context) {
//	var newUser = user{
//		Username:  c.PostForm("Username"),
//		Password:  c.PostForm("Password1"),
//		FirstName: c.PostForm("FirstName"),
//		LastName:  c.PostForm("LastName"),
//		Email:     c.PostForm("Email"),
//	}
//	if _, err := app.registerNewUser(newUser); err == nil {
//		token := generateSessionToken()
//		c.SetCookie("token", token, 3600, "", "", false, true)
//		c.Set("is_logged_in", true)
//
//		render(c, gin.H{"IsLogin": true}, "profile.html")
//
//	} else {
//		c.HTML(http.StatusBadRequest, "registration.html", gin.H{
//			"IsError":      true,
//			"ErrorMessage": err.Error()})
//
//	}
//}
//
//func (app *App) registerNewUser(newUser user) (*user, error) {
//
//	var temp user
//	app.db.First(&temp, "username = ?", newUser.Username)
//
//	if strings.TrimSpace(newUser.Password) == "" {
//		return nil, errors.New("The password can't be empty")
//	} else if (temp != user{}) {
//		return nil, errors.New("The username isn't available")
//	}
//	app.db.Create(&newUser)
//
//	return &temp, nil
//}
//
//func (app *App) checkUsername(c *gin.Context) {
//	username := c.PostForm("Username")
//	var temp user
//	var answer = false
//	app.db.First(&temp, "username = ?", username)
//	if (temp == user{}) {
//		answer = true
//	}
//
//	c.JSON(200, gin.H{
//		"success": answer,
//		"IsError": false,
//	})
//}
//
//func setUserStatus() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		if token, err := c.Cookie("token"); err == nil || token != "" {
//			c.Set("is_logged_in", true)
//		} else {
//			c.Set("is_logged_in", false)
//		}
//	}
//}
//func (app *App) doLike(c *gin.Context) {
//	var event = likeEvent{
//		Username: c.PostForm("login"),
//		Password: c.PostForm("password"),
//		Target:   c.PostForm("target"),
//		Error:    "",
//	}
//	IsError, err := likeTarget(c)
//	c.JSON(200, gin.H{
//		"error":   err,
//		"IsError": IsError,
//	})
//	event.Error = err
//	app.db.Create(&event)
//}

//func generateSessionToken() string {
//	return strconv.FormatInt(rand.Int63(), 16)
//}
