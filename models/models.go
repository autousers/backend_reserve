package models

import (
	"time"
)

// Rooms schema of the rooms table.
type Rooms struct {
	ID         int8      `json:"id"`
	Room_name  string    `json:"room_name"`
	Room_desc  string    `json:"room_desc"`
	Created_at time.Time `json:"created_at"`
}

type Reserves struct {
	ID            int8      `json:"id"`
	Reserves_name string    `json:"reserves_name"`
	Reserves_room int8      `json:"reserves_room"`
	Start_date    time.Time `json:"start_date"`
	End_date      time.Time `json:"end_date"`
	Created_at    time.Time `json:"created_at"`
}
