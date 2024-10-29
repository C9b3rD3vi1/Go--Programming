package main


type car struct {
	brand      string
	model      string
	doors      int
	mileage    int
	frontWheel wheel
	backWheel  wheel
}


type wheel struct {
	frontWheel string
	backWheel string
}