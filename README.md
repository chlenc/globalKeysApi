# globalKeysApi

This api used in [globalKeys app](https://github.com/YablochnayaVega/GlobalKeys)

To start watch server run `realize s globalkeys`

To resrt local database go to `/sql` and run `reload_base.sh`
To reset heroku database run `heroku_reload.sh`


Run the following command to deploy your code to heroku

`godep save`

`git add .`

`git commit -m " <Your message> "`

`git push heroku master`

current api address https://globalkeys.herokuapp.com/

# globalKeysApi documentation

## Cities api

##### Data format

Name | Type | Required | Default value | Description
---- | ---- | -------- |-------------- |-------------
id | autoincrement int | false |  | Increments by 1 as data is added
name | string | true |  | Name of the city
photo | string | false | null | Link to city image
offers | int | false | 0 | Count of offers in this city
rooms | int | false | 0 | Count of avalible rooms in this city

##### Avalible methods
Type | Dethod |  Required parameters | Optional parameters | Description
---- | ------ | -------------------- | ------------------- | -----------
GET | /api/city |  | id | Get cities list

## Hotels api

##### Data format

Name | Type | Required | Default value | Description
---- | ---- | -------- |-------------- |-------------
id | autoincrement int | false |  | Increments by 1 as data is added
name | string | true |  | Name of the Hotel
country | string | false | 'Россия' | City where the hotel is located
address | string | false | 'Не указан' | Hotel address
latitude | float | false | 46.2062966 | Hotel latitude
longitude | float | false | 6.1466899 | Hotel longitude
photo | string | false | null | Link to hotel image
description | HTML | false | 'Нет описания' | Hotel description
stars | int | false | 0 | Stars count
cityId | int | false |  | Identifier of the city where the hotel is located

##### Avalible methods
Type | Dethod |  Required parameters | Optional parameters | Description
---- | ------ | -------------------- | ------------------- | -----------
GET | /api/hotel |  | id, city | Get hotels list


## Rooms api

##### Data format

Name | Type | Required | Default value | Description
---- | ---- | -------- |-------------- |-------------
id | autoincrement int | false |  | Increments by 1 as data is added
room | int | true |  | Room number
persons | int | true |  | How many people fit in a room
floor | int | false | 1 | Room floor
housing | int | false | 1 | Room housing
description | HTML | false | 'Нет описания' | Description in HTML
price | int | false | 1 | Price per day
city_id | int | false | 0 | City id
hotel_id | int | true |  | Hotel id


##### Avalible methods
Type | Dethod |  Required parameters | Optional parameters | Description
---- | ------ | -------------------- | ------------------- | -----------
GET | /api/room |  | id, hotel | Get rooms list


## Bookings api

##### Data format

Name | Type | Required | Default value | Description
---- | ---- | -------- |-------------- |-------------
id | autoincrement int | false |  | Increments by 1 as data is added
start_datetime | string | true |  | Start booking datetime
end_datetime | string | true |  | End booking datetime
cost | int | true |  | Booking cost
hotelId | int | true |  | Hotel id
roomId | int | true |  | Room id
customerId | int | true |  | Customer id


##### Avalible methods
Type | Dethod |  Required parameters | Optional parameters | Description
---- | ------ | -------------------- | ------------------- | -----------
GET | /api/booking |  | id, hotel | Get bookings list
POST | /api/booking | start_datetime, end_datetime, cost, hotel_id, room_id, customer_id | | Add new booking




