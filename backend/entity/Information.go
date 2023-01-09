package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// เข้ารหัส
func SetupPasswordHash(pwd string) string {
	var password, _ = bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(password)
}

func SetupIntoDatabase(db *gorm.DB) {
	//ระบบจัดการข้อมูลพนักงาน

	//ตำแหน่ง
	Officerrole := UserRole{
		RoleName: "Officer",
	}
	db.Model(&UserRole{}).Create(&Officerrole)

	Employeerole := UserRole{
		RoleName: "Employee",
	}
	db.Model(&UserRole{}).Create(&Employeerole)

	//login
	loginOfficer1 := Signin{
		Username: "OFSongsawang",
		Password: SetupPasswordHash("Songsawang01"),
		UserRole: Officerrole,
	}
	db.Model(&Signin{}).Create(&loginOfficer1)

	loginOfficer2 := Signin{
		Username: "OFMoonnight",
		Password: SetupPasswordHash("Moonnight02"),
		UserRole: Officerrole,
	}
	db.Model(&Signin{}).Create(&loginOfficer2)

	//Employee login
	loginEmployee1 := Signin{
		Username: "ESobsa",
		Password: SetupPasswordHash("Sobsa01"),
		UserRole: Employeerole,
	}
	db.Model(&Signin{}).Create(&loginEmployee1)
	loginEmployee2 := Signin{
		Username: "EHanoi",
		Password: SetupPasswordHash("Hanoi02"),
		UserRole: Employeerole,
	}
	db.Model(&Signin{}).Create(&loginEmployee2)

	loginEmployee3 := Signin{
		Username: "EBanana",
		Password: SetupPasswordHash("Banana03"),
		UserRole: Employeerole,
	}
	db.Model(&Signin{}).Create(&loginEmployee3)

	// Set Data Officer
	db.Model(&Officer{}).Create(&Officer{
		Officername: "Phonsak songsang",
		Tel:         "0981522594",
		Signin:      loginOfficer1,
	})
	db.Model(&Officer{}).Create(&Officer{

		Officername: "Moonnight machine",
		Tel:         "0981521111",
		Signin:      loginOfficer2,
	})

	// Set Data Department
	db.Model(&Department{}).Create(&Department{
		Name: "Reception",
	})
	db.Model(&Department{}).Create(&Department{
		Name: "Housekeeping",
	})
	db.Model(&Department{}).Create(&Department{
		Name: "Chef",
	})
	db.Model(&Department{}).Create(&Department{
		Name: "Accounting",
	})
	db.Model(&Department{}).Create(&Department{
		Name: "Sale & Marketing",
	})

	// Set Data Position
	db.Model(&Position{}).Create(&Position{
		Name: "Department head",
	})
	db.Model(&Position{}).Create(&Position{
		Name: "Full time employee",
	})
	db.Model(&Position{}).Create(&Position{
		Name: "Part time employee",
	})

	// Set Data Location
	db.Model(&Location{}).Create(&Location{
		Name: "Kitchen Room",
	})
	db.Model(&Location{}).Create(&Location{
		Name: "Cashier",
	})
	db.Model(&Location{}).Create(&Location{
		Name: "Hotel Front",
	})
	db.Model(&Location{}).Create(&Location{
		Name: "Housekeeping staff Room",
	})
	db.Model(&Location{}).Create(&Location{
		Name: "Marketing staff room",
	})

	var OFSongsawang Officer
	var OFMoonnight Officer
	db.Raw("SELECT * FROM officers WHERE officername = ?", "Phonsak songsang").Scan(&OFSongsawang)
	db.Raw("SELECT * FROM officers WHERE  officername = ?", "Moonnight machine").Scan(&OFMoonnight)

	var Reception Department
	var Housekeeping Department
	var Chef Department
	var Accounting Department
	var Salemarketing Department
	db.Raw("SELECT * FROM departments WHERE name = ?", "Reception").Scan(&Reception)
	db.Raw("SELECT * FROM departments WHERE name = ?", "Housekeeping").Scan(&Housekeeping)
	db.Raw("SELECT * FROM departments WHERE name = ?", "Chef").Scan(&Chef)
	db.Raw("SELECT * FROM departments WHERE name = ?", "Accounting").Scan(&Accounting)
	db.Raw("SELECT * FROM departments WHERE name = ?", "Sale & Marketing").Scan(&Salemarketing)

	var Head Position
	var Full Position
	var Part Position
	db.Raw("SELECT * FROM positions WHERE name = ?", "Department head").Scan(&Head)
	db.Raw("SELECT * FROM positions WHERE name = ?", "Full time employee").Scan(&Full)
	db.Raw("SELECT * FROM positions WHERE name = ?", "Part time employee").Scan(&Part)

	var Kitchen Location
	var Cashier Location
	var Hotelfront Location
	var Housekeepingstaff Location
	var Marketing Location
	db.Raw("SELECT * FROM locations WHERE name = ?", "Kitchen Room").Scan(&Kitchen)
	db.Raw("SELECT * FROM locations WHERE name = ?", "Cashier").Scan(&Cashier)
	db.Raw("SELECT * FROM locations WHERE name = ?", "Hotel Front").Scan(&Hotelfront)
	db.Raw("SELECT * FROM locations WHERE name = ?", "Housekeeping staff Room").Scan(&Housekeepingstaff)
	db.Raw("SELECT * FROM locations WHERE name = ?", "Marketing staff room").Scan(&Marketing)

	timedate1 := time.Date(1950, 2, 16, 0, 0, 0, 0, time.Local)
	timeyear1 := time.Date(1987, 2, 16, 0, 0, 0, 0, time.Local)
	timedate2 := time.Date(1965, 9, 9, 0, 0, 0, 0, time.Local)
	timeyear2 := time.Date(1995, 9, 9, 0, 0, 0, 0, time.Local)
	timedate3 := time.Date(1970, 5, 13, 0, 0, 0, 0, time.Local)
	timeyear3 := time.Date(1990, 5, 13, 0, 0, 0, 0, time.Local)

	db.Model(&Employee{}).Create(&Employee{
		PersonalID:   1430099536148,
		Employeename: "Sobsa tugwan",
		Email:        "Sobsa@gmail.com",
		Eusername:    "ESobsa",
		Password:     SetupPasswordHash("Sobsa01"),
		Department:   Chef,
		Position:     Head,
		Location:     Kitchen,
		Salary:       50000,
		Phonenumber:  "0905452001",
		Gender:       "Male",
		DateOfBirth:  timedate1,
		YearOfStart:  timeyear1,
		Blood:        "A",
		Address:      "219 m.10, nongprajak s, nongsham d, Ayutthaya 13000",
		Officer:      OFSongsawang,
		Signin:       loginEmployee1,
	})

	db.Model(&Employee{}).Create(&Employee{
		PersonalID:   1400000005257,
		Employeename: "Hanoi slotmachine",
		Email:        "Hanoi@gmail.com",
		Eusername:    "EHanoi",
		Password:     SetupPasswordHash("Hanoi02"),
		Department:   Housekeeping,
		Position:     Full,
		Location:     Housekeepingstaff,
		Salary:       25000,
		Phonenumber:  "0982542528",
		Gender:       "Female",
		DateOfBirth:  timedate2,
		YearOfStart:  timeyear2,
		Blood:        "B",
		Address:      "157 m.1, seesad s, dokpeeb d,Nakhonratchasima 30000",
		Officer:      OFSongsawang,
		Signin:       loginEmployee2,
	})

	db.Model(&Employee{}).Create(&Employee{
		PersonalID:   1585289653250,
		Employeename: "Banana amoi",
		Email:        "Banana@gmail.com",
		Eusername:    "EBanana",
		Password:     SetupPasswordHash("Eanana03"),
		Department:   Salemarketing,
		Position:     Head,
		Location:     Marketing,
		Salary:       50000,
		Phonenumber:  "0824545304",
		Gender:       "Female",
		DateOfBirth:  timedate3,
		YearOfStart:  timeyear3,
		Blood:        "AB",
		Address:      "426 m.6, yabyol s, nangrong d,Buriram 31000",
		Officer:      OFMoonnight,
		Signin:       loginEmployee3,
	})

	var Sobsa Employee
	var Hanoi Employee
	var Banana Employee
	db.Raw("SELECT * FROM employees WHERE employeename  = ?", "Sobsa tugwan").Scan(&Sobsa)
	db.Raw("SELECT * FROM employees WHERE employeename  = ?", "Hanoi slotmachine").Scan(&Hanoi)
	db.Raw("SELECT * FROM employees WHERE employeename  = ?", "Banana amoi").Scan(&Banana)

	//Room Data
	//RoomType Data
	Standard := RoomType{
		Size:    "Standard",
		Bedsize: "Single Bedded",
		Bedtype: "King Sized Bed",
	}
	db.Model(&RoomType{}).Create(&Standard)

	Superior := RoomType{
		Size:    "Superior",
		Bedsize: "Twin Bedded",
		Bedtype: "Queen Sized Bed",
	}
	db.Model(&RoomType{}).Create(&Superior)

	Deluxe := RoomType{
		Size:    "Deluxe",
		Bedsize: "Double Bedded",
		Bedtype: "Double Bed",
	}
	db.Model(&RoomType{}).Create(&Deluxe)

	Suite := RoomType{
		Size:    "Suite",
		Bedsize: "Triple Bedded",
		Bedtype: "King Sized Bed",
	}
	db.Model(&RoomType{}).Create(&Suite)

	//RoomZone Data
	A := RoomZone{
		Name: "A",
	}
	db.Model(&RoomZone{}).Create(&A)

	B := RoomZone{
		Name: "B",
	}
	db.Model(&RoomZone{}).Create(&B)

	C := RoomZone{
		Name: "C",
	}
	db.Model(&RoomZone{}).Create(&C)

	D := RoomZone{
		Name: "D",
	}
	db.Model(&RoomZone{}).Create(&D)

	//State Data
	on := State{
		Name: "on",
	}
	db.Model(&State{}).Create(&on)

	off := State{
		Name: "off",
	}
	db.Model(&State{}).Create(&off)

	//Room Data
	//Room1
	db.Model(&Room{}).Create(&Room{
		Employee: Sobsa,
		RoomType: Standard,
		RoomZone: A,
		State:    on,
		Time:     time.Now(),
	})
	//Room2
	db.Model(&Room{}).Create(&Room{
		Employee: Hanoi,
		RoomType: Superior,
		RoomZone: B,
		State:    on,
		Time:     time.Now(),
	})
	//Room3
	db.Model(&Room{}).Create(&Room{
		Employee: Banana,
		RoomType: Deluxe,
		RoomZone: C,
		State:    off,
		Time:     time.Now(),
	})

	//ระบบ check In-Out

	//set status data

	checkin := CheckInOutStatus{
		Name: "Checked In",
	}
	db.Model(&CheckInOutStatus{}).Create(&checkin)

	checkout := CheckInOutStatus{
		Name: "Checked Out",
	}
	db.Model(&CheckInOutStatus{}).Create(&checkout)

	//set check In-Out data

	db.Model(&CheckInOut{}).Create(&CheckInOut{
		Booking:          3001, //dump
		CheckInTime:      time.Now(),
		CheckOutTime:     time.Now(),
		CheckInOutStatus: checkout,
		Employee:         Sobsa,
	})

	db.Model(&CheckInOut{}).Create(&CheckInOut{
		Booking:          3002, //dump
		CheckInTime:      time.Now(),
		CheckOutTime:     time.Now(),
		CheckInOutStatus: checkin,
		Employee:         Banana,
	})

	//ระบบแจ้งซ่อม

	//set type data
	// air := RepairType{
	// 	Name: "Air Conditioner",
	// }
	// db.Model(&RepairType{}).Create(&air)

	// toilet := RepairType{
	// 	Name: "Toilet",
	// }
	// db.Model(&RepairType{}).Create(&toilet)

	// light := RepairType{
	// 	Name: "Light Bulb",
	// }
	// db.Model(&RepairType{}).Create(&light)

	// fur := RepairType{
	// 	Name: "Furniture",
	// }
	// db.Model(&RepairType{}).Create(&fur)

	// elec := RepairType{
	// 	Name: "Electrical appliance",
	// }
	// db.Model(&RepairType{}).Create(&elec)

	// //set status data
	// inprogress := RepairStatus{
	// 	Name: "In Progress",
	// }
	// db.Model(&RepairStatus{}).Create(&inprogress)

	// success := RepairStatus{
	// 	Name: "Success",
	// }
	// db.Model(&RepairStatus{}).Create(&success)

	// db.Model(&RepairReq{}).Create(&RepairReq{
	// 	Room:         12,
	// 	RepairType:   air,
	// 	Note:         "air not cool",
	// 	Time:         time.Now(),
	// 	RepairStatus: success,
	// 	User:         2,
	// })

	// db.Model(&RepairReq{}).Create(&RepairReq{
	// 	Room:         11,
	// 	RepairType:   fur,
	// 	Note:         "bed is broken",
	// 	Time:         time.Now(),
	// 	RepairStatus: inprogress,
	// 	User:         1,
	// })

	// db.Model(&RepairReq{}).Create(&RepairReq{
	// 	Room:         12,
	// 	RepairType:   air,
	// 	Note:         "air not cool",
	// 	Time:         time.Now(),
	// 	RepairStatus: inprogress,
	// 	User:         3,
	// })
}