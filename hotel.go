package main

import "fmt"

type Room struct {
	ID       int
	Type     string
	Status   bool // true fo reserve,false for available
	BedCount int
	Price    int
}

var Rooms []Room = GenerateRooms()

func main() {
	input := ""
	for input != "exit" {
		fmt.Println("Enter a command:")
		fmt.Println("1 : Room list")
		fmt.Println("2 : Add room")
		fmt.Println("3 : Reserve room")
		fmt.Scanln(&input)
		switch input {
		case "1":
			GetRoomList()
		case "2":
			AddRoom()
		case "3":
			ReserveRoom()
		case "exit":
			fmt.Println("Existing...")
			break
		default:
			fmt.Println("Invalid command")
		}
	}

}

func GetRoomList() {
	for _, room := range Rooms {
		fmt.Printf("%+v \n", room)
	}

}

func GetRoomFromInput() Room {
	var room Room = Room{Status: false}
	fmt.Println("Enter room information line by line (ID,Type,Status,BedCount,Price)")
	fmt.Scanln(&room.ID)
	fmt.Scanln(&room.Type)
	fmt.Scanln(&room.BedCount)
	fmt.Scanln(&room.Price)

	return room

}

func AddRoom() {
	room := GetRoomFromInput()
	Rooms = append(Rooms, room)

}

func ReserveRoom() {

}

func CalculateRoomPrice() {

}

func GenerateRooms() []Room {
	rooms := []Room{}

	rooms = append(rooms, Room{ID: 1, Type: "Single", Status: false, BedCount: 1, Price: 100})
	rooms = append(rooms, Room{ID: 2, Type: "Single", Status: false, BedCount: 1, Price: 120})
	rooms = append(rooms, Room{ID: 3, Type: "Single", Status: false, BedCount: 1, Price: 150})
	rooms = append(rooms, Room{ID: 4, Type: "Double", Status: false, BedCount: 2, Price: 220})
	rooms = append(rooms, Room{ID: 5, Type: "Double", Status: false, BedCount: 2, Price: 220})
	rooms = append(rooms, Room{ID: 6, Type: "Double", Status: false, BedCount: 2, Price: 250})
	rooms = append(rooms, Room{ID: 7, Type: "Double", Status: false, BedCount: 2, Price: 230})
	rooms = append(rooms, Room{ID: 8, Type: "Standard", Status: false, BedCount: 4, Price: 300})
	rooms = append(rooms, Room{ID: 9, Type: "Standard", Status: false, BedCount: 4, Price: 320})
	rooms = append(rooms, Room{ID: 10, Type: "Standard", Status: false, BedCount: 4, Price: 360})

	return rooms
}
