package db

//GetRooms gets all rooms
func GetRooms() []Room {
	var rooms []Room
	db.Preload("Connections").Find(&rooms)
	return rooms
}
