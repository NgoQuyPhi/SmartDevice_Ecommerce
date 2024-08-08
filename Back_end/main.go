package main

import (
	database "PJ/SmartDevice_Ecomerce/Back_end/Database"
	routers "PJ/SmartDevice_Ecomerce/Back_end/Router"
)

func main() {
	var DBSTR = "root:NgocBich1609@@@tcp(localhost:3306)/smartdevice_ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	database.Connect(DBSTR)
	r := routers.InitRouter()

	r.Run()
}
