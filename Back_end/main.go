package main

import (
	repositories "PJ/SmartDevice_Ecomerce/Back_end/Repositories"
	routers "PJ/SmartDevice_Ecomerce/Back_end/Router"
)

func main() {
	var DB string = "root:NgocBich1609@@@tcp(localhost:3306)/smartdevice_ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	repositories.Connect(DB)
	r := routers.InitRouter()

	r.Run()
}
