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

	Customerrole := UserRole{
		RoleName: "Customer",
	}
	db.Model(&UserRole{}).Create(&Customerrole)

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

	//Doctor login
	loginCustomer1 := Signin{
		Username: "Sandee12@gmail.com",
		Password: SetupPasswordHash("SD123456"),
		UserRole: Customerrole,
	}
	db.Model(&Signin{}).Create(&loginCustomer1)
	loginCustomer2 := Signin{
		Username: "Nicha@gmail.com",
		Password: SetupPasswordHash("NC332548"),
		UserRole: Customerrole,
	}
	db.Model(&Signin{}).Create(&loginCustomer2)

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
	var Room1 Room
	var Room2 Room
	var Room3 Room
	db.Raw("SELECT * FROM rooms WHERE id = ?", "1").Scan(&Room1)
	db.Raw("SELECT * FROM rooms WHERE id = ?", "2").Scan(&Room2)
	db.Raw("SELECT * FROM rooms WHERE id = ?", "3").Scan(&Room3)

	//ระบบสมัครสมาชิก
	var Customer1 Customer
	var Customer2 Customer

	//Gender
	Gender1 := Gender{
		Gender: "Male",
	}
	db.Model(&Gender{}).Create(&Gender1)

	Gender2 := Gender{
		Gender: "Female",
	}
	db.Model(&Gender{}).Create(&Gender2)

	Gender3 := Gender{
		Gender: "Commom",
	}
	db.Model(&Gender{}).Create(&Gender3)

	Gender4 := Gender{
		Gender: "Neuter",
	}
	db.Model(&Gender{}).Create(&Gender4)

	Gender5 := Gender{
		Gender: "Other",
	}
	db.Model(&Gender{}).Create(&Gender5)

	//Province
	Province1 := Province{
		Province: "Amnat Charoen",
	}
	db.Model(&Province{}).Create(&Province1)

	Province2 := Province{
		Province: "Ang Thong",
	}
	db.Model(&Province{}).Create(&Province2)

	Province3 := Province{
		Province: "Ayutthaya",
	}
	db.Model(&Province{}).Create(&Province3)

	Province4 := Province{
		Province: "Bangkok",
	}
	db.Model(&Province{}).Create(&Province4)

	////////////////////////////////////เพิ่มม//////////////////////////////////////////////
	Province5 := Province{
		Province: "Bueng Kan",
	}
	db.Model(&Province{}).Create(&Province5)

	Province6 := Province{
		Province: "Buriram",
	}
	db.Model(&Province{}).Create(&Province6)

	Province7 := Province{
		Province: "Chachoengsao",
	}
	db.Model(&Province{}).Create(&Province7)

	Province8 := Province{
		Province: "Chai Nat",
	}
	db.Model(&Province{}).Create(&Province8)

	Province9 := Province{
		Province: "Chaiyaphum",
	}
	db.Model(&Province{}).Create(&Province9)

	Province10 := Province{
		Province: "Chanthaburi",
	}
	db.Model(&Province{}).Create(&Province10)

	Province11 := Province{
		Province: "Chiang Mai",
	}
	db.Model(&Province{}).Create(&Province11)

	Province12 := Province{
		Province: "Chiang Rai",
	}
	db.Model(&Province{}).Create(&Province12)

	Province13 := Province{
		Province: "Chonburi",
	}
	db.Model(&Province{}).Create(&Province13)

	Province14 := Province{
		Province: "Chumphon",
	}
	db.Model(&Province{}).Create(&Province14)

	Province15 := Province{
		Province: "Kalasin",
	}
	db.Model(&Province{}).Create(&Province15)

	Province16 := Province{
		Province: "Kamphaeng Phet",
	}
	db.Model(&Province{}).Create(&Province16)

	Province17 := Province{
		Province: "Kanchanaburi",
	}
	db.Model(&Province{}).Create(&Province17)

	Province18 := Province{
		Province: "Khon Kaen",
	}
	db.Model(&Province{}).Create(&Province18)

	Province19 := Province{
		Province: "Krabi",
	}
	db.Model(&Province{}).Create(&Province19)

	Province20 := Province{
		Province: "Lampang",
	}
	db.Model(&Province{}).Create(&Province20)

	Province21 := Province{
		Province: "Lamphun",
	}
	db.Model(&Province{}).Create(&Province21)

	Province22 := Province{
		Province: "Loei",
	}
	db.Model(&Province{}).Create(&Province22)

	Province23 := Province{
		Province: "Lopburi",
	}
	db.Model(&Province{}).Create(&Province23)

	Province24 := Province{
		Province: "Mae Hong Son",
	}
	db.Model(&Province{}).Create(&Province24)

	Province25 := Province{
		Province: "Maha Sarakham",
	}
	db.Model(&Province{}).Create(&Province25)

	Province26 := Province{
		Province: "Mukdahan",
	}
	db.Model(&Province{}).Create(&Province26)

	Province27 := Province{
		Province: "Nakhon Nayok",
	}
	db.Model(&Province{}).Create(&Province27)

	Province28 := Province{
		Province: "Nakhon Pathom",
	}
	db.Model(&Province{}).Create(&Province28)

	Province29 := Province{
		Province: "Nakhon Phanom",
	}
	db.Model(&Province{}).Create(&Province29)

	Province30 := Province{
		Province: "Nakhon Ratchasima",
	}
	db.Model(&Province{}).Create(&Province30)

	Province31 := Province{
		Province: "Nakhon Sawan",
	}
	db.Model(&Province{}).Create(&Province31)

	Province32 := Province{
		Province: "Nakhon Si Thammarat",
	}
	db.Model(&Province{}).Create(&Province32)

	Province33 := Province{
		Province: "Nan",
	}
	db.Model(&Province{}).Create(&Province33)

	Province34 := Province{
		Province: "Narathiwat",
	}
	db.Model(&Province{}).Create(&Province34)

	Province35 := Province{
		Province: "Nong Bua Lamphu",
	}
	db.Model(&Province{}).Create(&Province35)

	Province36 := Province{
		Province: "Nong Khai",
	}
	db.Model(&Province{}).Create(&Province36)

	Province37 := Province{
		Province: "Nonthaburi",
	}
	db.Model(&Province{}).Create(&Province37)

	Province38 := Province{
		Province: "Pathum Thani",
	}
	db.Model(&Province{}).Create(&Province38)

	Province39 := Province{
		Province: "Pattani",
	}
	db.Model(&Province{}).Create(&Province39)

	Province40 := Province{
		Province: "Phang Nga",
	}
	db.Model(&Province{}).Create(&Province40)

	Province41 := Province{
		Province: "Phatthalung",
	}
	db.Model(&Province{}).Create(&Province41)

	Province42 := Province{
		Province: "Phayao",
	}
	db.Model(&Province{}).Create(&Province42)

	Province43 := Province{
		Province: "Phetchabun",
	}
	db.Model(&Province{}).Create(&Province43)

	Province44 := Province{
		Province: "Phetchaburi",
	}
	db.Model(&Province{}).Create(&Province44)

	Province45 := Province{
		Province: "Phichit",
	}
	db.Model(&Province{}).Create(&Province45)

	Province46 := Province{
		Province: "Phitsanulok",
	}
	db.Model(&Province{}).Create(&Province46)

	Province47 := Province{
		Province: "Phrae",
	}
	db.Model(&Province{}).Create(&Province47)

	Province48 := Province{
		Province: "Phuket",
	}
	db.Model(&Province{}).Create(&Province48)

	Province49 := Province{
		Province: "Prachin Buri",
	}
	db.Model(&Province{}).Create(&Province49)

	Province50 := Province{
		Province: "Prachuap Khiri Khan",
	}
	db.Model(&Province{}).Create(&Province50)

	Province51 := Province{
		Province: "Ranong",
	}
	db.Model(&Province{}).Create(&Province51)

	Province52 := Province{
		Province: "Ratchaburi",
	}
	db.Model(&Province{}).Create(&Province52)

	Province53 := Province{
		Province: "Rayong",
	}
	db.Model(&Province{}).Create(&Province53)

	Province54 := Province{
		Province: "Roi Et",
	}
	db.Model(&Province{}).Create(&Province54)

	Province55 := Province{
		Province: "Sa Kaeo",
	}
	db.Model(&Province{}).Create(&Province55)

	Province56 := Province{
		Province: "Sakon Nakhon",
	}
	db.Model(&Province{}).Create(&Province56)

	Province57 := Province{
		Province: "Samut Prakan",
	}
	db.Model(&Province{}).Create(&Province57)

	Province58 := Province{
		Province: "Samut Sakhon",
	}
	db.Model(&Province{}).Create(&Province58)

	Province59 := Province{
		Province: "Samut Songkhram",
	}
	db.Model(&Province{}).Create(&Province59)

	Province60 := Province{
		Province: "Saraburi",
	}
	db.Model(&Province{}).Create(&Province60)

	Province61 := Province{
		Province: "Satun",
	}
	db.Model(&Province{}).Create(&Province61)

	Province62 := Province{
		Province: "Sing Buri",
	}
	db.Model(&Province{}).Create(&Province62)

	Province63 := Province{
		Province: "Sisaket",
	}
	db.Model(&Province{}).Create(&Province63)

	Province64 := Province{
		Province: "Songkhla",
	}
	db.Model(&Province{}).Create(&Province64)

	Province65 := Province{
		Province: "Sukhothai",
	}
	db.Model(&Province{}).Create(&Province65)

	Province66 := Province{
		Province: "Suphan Buri",
	}
	db.Model(&Province{}).Create(&Province66)

	Province67 := Province{
		Province: "Surat Thani",
	}
	db.Model(&Province{}).Create(&Province67)

	Province68 := Province{
		Province: "Surin",
	}
	db.Model(&Province{}).Create(&Province68)

	Province69 := Province{
		Province: "Tak",
	}
	db.Model(&Province{}).Create(&Province69)

	Province70 := Province{
		Province: "Trang",
	}
	db.Model(&Province{}).Create(&Province70)

	Province71 := Province{
		Province: "Trat",
	}
	db.Model(&Province{}).Create(&Province71)

	Province72 := Province{
		Province: "Ubon Ratchathani",
	}
	db.Model(&Province{}).Create(&Province72)

	Province73 := Province{
		Province: "Udon Thani",
	}
	db.Model(&Province{}).Create(&Province73)

	Province74 := Province{
		Province: "Uthai Thani",
	}
	db.Model(&Province{}).Create(&Province74)

	Province75 := Province{
		Province: "Uttaradit",
	}
	db.Model(&Province{}).Create(&Province75)

	Province76 := Province{
		Province: "Yala",
	}
	db.Model(&Province{}).Create(&Province76)

	Province77 := Province{
		Province: "Yasothon",
	}
	db.Model(&Province{}).Create(&Province77)

	/////////////////////////////////////////////////

	//Nametitle  =============================แก้ 1/23/2023===================================================================
	Nametitle1 := Nametitle{
		Nametitle: "Mr.",
	}
	db.Model(&Nametitle{}).Create(&Nametitle1)

	Nametitle2 := Nametitle{
		Nametitle: "Mrs.",
	}
	db.Model(&Nametitle{}).Create(&Nametitle2)

	Nametitle3 := Nametitle{
		Nametitle: "Ms.",
	}
	db.Model(&Nametitle{}).Create(&Nametitle3)

	//Csutomer
	password1, _ := bcrypt.GenerateFromPassword([]byte("SD123456"), 14)
	Customer1 = Customer{
		Gender:      Gender1,
		Province:    Province1,
		FirstName:   "Sandee",
		LastName:    "Masuk",
		Password:    string(password1),
		Age:         30,
		Phone:       "095-7456921",
		Email:       "Sandee12@gmail.com",
		Nametitle:  Nametitle1,
		Signin:      loginCustomer1,
	}
	db.Model(&Customer{}).Create(&Customer1)

	password2, _ := bcrypt.GenerateFromPassword([]byte("NC332548"), 14)
	Customer2 = Customer{
		Gender:      Gender2,
		Province:    Province2,
		FirstName:   "Nicha",
		LastName:    "Memak",
		Password:    string(password2),
		Age:         25,
		Phone:       "084-5215667",
		Email:       "Nicha@gmail.com",
		Nametitle:  Nametitle2,
		Signin:      loginCustomer2,
	}
	db.Model(&Customer{}).Create(&Customer2)

	//ระบบแจ้งซ่อม

	//set type data
	air := RepairType{
		Name: "Air Conditioner",
	}
	db.Model(&RepairType{}).Create(&air)

	toilet := RepairType{
		Name: "Toilet",
	}
	db.Model(&RepairType{}).Create(&toilet)

	light := RepairType{
		Name: "Light Bulb",
	}
	db.Model(&RepairType{}).Create(&light)

	fur := RepairType{
		Name: "Furniture",
	}
	db.Model(&RepairType{}).Create(&fur)

	elec := RepairType{
		Name: "Electrical appliance",
	}
	db.Model(&RepairType{}).Create(&elec)

	//set status data
	inprogress := RepairStatus{
		Name: "In Progress",
	}
	db.Model(&RepairStatus{}).Create(&inprogress)

	success := RepairStatus{
		Name: "Success",
	}
	db.Model(&RepairStatus{}).Create(&success)

	db.Model(&RepairReq{}).Create(&RepairReq{
		Room:         Room1,
		RepairType:   air,
		Note:         "air not cool",
		Time:         time.Now(),
		RepairStatus: success,
		Customer:     Customer1,
	})

	db.Model(&RepairReq{}).Create(&RepairReq{
		Room:         Room2,
		RepairType:   fur,
		Note:         "bed is broken",
		Time:         time.Now(),
		RepairStatus: inprogress,
		Customer:     Customer2,
	})

	db.Model(&RepairReq{}).Create(&RepairReq{
		Room:         Room3,
		RepairType:   air,
		Note:         "air not cool",
		Time:         time.Now(),
		RepairStatus: inprogress,
		Customer:     Customer1,
	})
	var RepairReq1 RepairReq
	var RepairReq2 RepairReq
	var RepairReq3 RepairReq
	db.Raw("SELECT * FROM repair_reqs WHERE id = ?", "1").Scan(&RepairReq1)
	db.Raw("SELECT * FROM repair_reqs WHERE id = ?", "2").Scan(&RepairReq2)
	db.Raw("SELECT * FROM repair_reqs WHERE id = ?", "3").Scan(&RepairReq3)

	// ===============     SERVICE     ===============

	// ===============     อาหาร     ===============
	db.Model(&Food{}).Create(&Food{
		Name:  "No Order",
		Price: 0,
		Item:  1,
	})
	db.Model(&Food{}).Create(&Food{
		Name:  "Pad Thai",
		Price: 65,
		Item:  50,
	})
	db.Model(&Food{}).Create(&Food{
		Name:  "Pad Kaphao",
		Price: 55,
		Item:  50,
	})
	db.Model(&Food{}).Create(&Food{
		Name:  "Noodles",
		Price: 55,
		Item:  50,
	})
	var Noodles Food
	db.Raw("SELECT * FROM Foods WHERE name = ?", "Noodles").Scan(&Noodles)

	// ===============     เครื่องดื่ม     ===============
	db.Model(&Drink{}).Create(&Drink{
		Name:  "No Order",
		Price: 0,
		Item:  1,
	})
	db.Model(&Drink{}).Create(&Drink{
		Name:  "Pepsi",
		Price: 15,
		Item:  50,
	})
	db.Model(&Drink{}).Create(&Drink{
		Name:  "Mansome",
		Price: 20,
		Item:  50,
	})
	db.Model(&Drink{}).Create(&Drink{
		Name:  "Water",
		Price: 10,
		Item:  50,
	})
	var Pepsi Drink
	db.Raw("SELECT * FROM drinks WHERE name = ?", "Pepsi").Scan(&Pepsi)

	// ===============     อุปกรณ์เสริม     ===============
	db.Model(&Accessories{}).Create(&Accessories{
		Name:  "No Order",
		Price: 0,
		Item:  1,
	})
	db.Model(&Accessories{}).Create(&Accessories{
		Name:  "Plug",
		Price: 40,
		Item:  15,
	})
	db.Model(&Accessories{}).Create(&Accessories{
		Name:  "Chair",
		Price: 70,
		Item:  7,
	})
	db.Model(&Accessories{}).Create(&Accessories{
		Name:  "Table",
		Price: 120,
		Item:  7,
	})
	db.Model(&Accessories{}).Create(&Accessories{
		Name:  "Table & Chair (small)",
		Price: 135,
		Item:  10,
	})
	db.Model(&Accessories{}).Create(&Accessories{
		Name:  "Bed",
		Price: 350,
		Item:  5,
	})
	var Bed Accessories
	db.Raw("SELECT * FROM Accessories WHERE name=?", "Bed").Scan(&Bed)

	// ===============     ตารางหลัก     ===============
	db.Model(&Service{}).Create(&Service{
		Customer:    Customer1,
		Time:        time.Now(),
		Food:        Noodles,
		Drink:       Pepsi,
		Accessories: Bed,
	})

	// ===============     PAYMENT     ===============

	// ===============     สถานที่ชำระเงิน     ===============
	db.Model(&Place{}).Create(&Place{
		Name: "Counter Service 1",
	})
	db.Model(&Place{}).Create(&Place{
		Name: "Counter Service 2",
	})
	var CS2 Place
	db.Raw("SELECT * FROM places WHERE name = ?", "Counter Service 2").Scan(&CS2)

	// ===============     สกุลเงินดิจิทอล     ===============
	db.Model(&Crypto{}).Create(&Crypto{
		Name:            "Ethereum",
		PublicKey:       "Asd54f98sadf84sa9d8f49sad4f",
		Picture:         "data:image/jpeg;base64,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",
		PaymentMethodID: 3,
	})
	db.Model(&Crypto{}).Create(&Crypto{
		Name:            "Tether",
		PublicKey:       "A4s59df4a9s8df489as4f89dsa4",
		Picture:         "data:image/jpeg;base64,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",
		PaymentMethodID: 3,
	})
	db.Model(&Crypto{}).Create(&Crypto{
		Name:            "DogECoin",
		PublicKey:       "Gfds98g79ds8fgsdf5gs4d65564",
		Picture:         "data:image/jpeg;base64,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",
		PaymentMethodID: 3,
	})
	db.Model(&Crypto{}).Create(&Crypto{
		Name:            "BitCoin",
		PublicKey:       "S6a5d4sdfsf9asd48f9asd6f24a",
		Picture:         "data:image/jpeg;base64,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",
		PaymentMethodID: 3,
	})
	db.Model(&Crypto{}).Create(&Crypto{
		Name:            "NULL",
		PublicKey:       "NULL",
		Picture:         "NULL",
		PaymentMethodID: 3,
	})
	var DogECoin Crypto
	db.Raw("SELECT * FROM cryptos WHERE name = ?", "DogECoin").Scan(&DogECoin)

	// ===============     ธนาคาร     ===============
	db.Model(&Bank{}).Create(&Bank{
		Name:            "PromptPay",
		Number:          "093-348-6360",
		Picture:         "data:image/jpeg;base64,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",
		PaymentMethodID: 2,
	})
	db.Model(&Bank{}).Create(&Bank{
		Name:            "SCB",
		Number:          "152-563-6599",
		Picture:         "data:image/jpeg;base64,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",
		PaymentMethodID: 2,
	})
	db.Model(&Bank{}).Create(&Bank{
		Name:            "Bangkok",
		Number:          "420-542-5416",
		Picture:         "data:image/jpeg;base64,iVBORw0KGgoAAAANSUhEUgAAAOEAAADhCAMAAAAJbSJIAAAAmVBMVEUTSZ////8AP5sAQZwARJ0AN5kAOpkANZgOR54APJoAQJwKRp4APZoAOZoANJgAMpeeq8/6+/3q7fXy9PmqttYaTaLU2uqUpMxJaa5uhbvX3Ou/x9/f4+/L0ubGzuPk6PI1W6izvtpie7WJm8dQb7GBlMR8kMKYp83v8PcrVaU+Yqtrg7suV6apttU7X6pHaK4AF5AAK5ZadrQ6wLDDAAALB0lEQVR4nO1dCXeqOhA2bAYDKqBYFaW4W7u86///cQ+0rWQm7m0TPPnOrVdb9MyQzD4TazUNDQ0NDQ0NDQ0Njb8HNXJYBquZZv7k8GDlP6x4UTyh+VXMZLsX5u5vVkM24ZeCPrdzbNozb97Pnyzm+cPu2Tb/mS3b7WnxgnY77Vn7oz2d5Je0l/njZGnJJv1CWBsyiOMpmf6XDEkcxdFgmK5GUUzeSDyYkDQl4yAe1FhGemQzJFkUJm9kGJKMeFQ27ZeBzVZpSOxoOSar+F82iibhMHLZC5knCVmQxYIsk5Q8vzg5h5MFyUajZEKyYVAlDgfrF2JGWRAOiDGMo340j5rWjKxbtZzD6Zz8a83IMqLFGuYcxqNkk3NIOsStCIfWlgQBcdMgGQeBm42TIFmsbPZE1mZE+sFiEdQ/gmhKIjKNggWJxuM4f0e2ClaebNIvBfNyUMMzzPw/s2Hs1STL1WfDswyLGex9mNVcz7M8I//XME3L8/JrjYosYQ6aI38snuxefv72+ynt9sgz21/zeQn9vuwxQOm8K5uGXwY1H2rFNDQ0NDQ0NDQ0NB4S7NHddvPjwVk0Z2RUl03Eb4KtA0I6D8wi9UKSY+HKJuTXYI/JDjNTNiW/BH+5Z5BENSabll+BMSFfGNiyifkNWIWW+cIjahtaT0kJ/crkyC9Ga0g4PFelHHcp3DnPIAkfbBGtDwIxfihtQ1mEOCRTRzZZPwfqjzCDD2X4naWIQRI9TJhhboUMEhI/iFWkLDjCIVk+hii2BscYJGT7CKLot48zSIJu9X3wxuQEg4UPXnVtw9ZACIfgdceXTeJ9oD4QwsG/DVjFSbVF0e7w7OSxL5TLaotiAy5Y7sZQOwbLWq+uKLI1YLBX2D9aA05qdUWR1hOelc9wwpw9iigiIbQ+t6PTewxRNKElfPoO6+sg1hhU0kFlXWD5SvEgbQBRbFcxHK4DSzhqlf5oQVGsoIMKzV7Ep/KRKFYuSXxCCPewgSjGfrWsIusCQevBSBBZxYrFik3gt4ywJkGRf6XSNlDKIlETMZLUykyl5JryCSyPWFOe1LZqw4DWTux4shq4EdOqlE7rY57wwbG1QQr3pRrFDA+UKILjnfzQcw0rIYrWC1iYzQmyHRB9ZFXw3tyUJ7pzimj2DG5H3/gzQm+FnfEkp6e1h7EALK5V995Q3uL9jPKwQeV00PwbQm8FMgDzs4XQxop/xxHTogpawFsbn18R5B4o7b0hb+0SYvGb1I0y8HJcZMHRwqtrMhrhTSJFayDfsVC1i6Ge8YRenF8yoQJ+VdN7Q4bi8hwhNKKJkhUplFu7xj1xgCPUVjHgh+WIq1xM5MzCtI4C8KY8ieF1DqYL3r5SbhHZK1iEjysXoQlyb5lqafAmCINQbu0cKAVirFi5BmaV4uv9Z5h7C5TqJ0LODL2BOhjwjxXap9QFzkz/lh1GffApC3WiYeeNJ214m2MJd0KgTDQMC0nfldBr4QBpFmTKpYBCh+SmPbr7pB/Z7T8PeOdT52YlaMBE5K274UeBGpyXd8Q+sI1xrEKivwV2Fnm6Qz9A51aFferB/t+ge8fOcjP4adI5pF0CaWrcwaGHWjU7sl1wH970n+aQrOUqG5SVLwoxd8BDN4xkchcRZqzvvemO4POeZS4iCs0LbO8Iz2FZuMBQZhJcdMtJ+3Z7SGFRQPYiCqQwR3L7PTdhJWoHiZLoZCKC7timtnhw4R4LexeYJaSHxLeGBNaRvv479v19MKdigm5Osfip+PMiWdG+nYgJImHzpm0FE5IH3KOe7wCFXdwH3DTIjBKSBwzl6BqDq/stOBmaX9/7Q71yn9HgX/lVIGebcoovsHn35uqOWDCksPa53I+UbUr9Mgkjn5+EDa6d1eZLcwuP16ttGWk3xuWfegbo4QqvS0D43JYf28C9GckQRN5WFGUKfmNd1fRrcHdn1yfmlzV1JKMHxSvzs+8u4FNlV9RWrPfyG8mmEGKXu18yvBqnXCwa7HxRkJS6eByd8cPs+1tj9cu/m0lIDttlB+Rt71eBoe3JZfqBchuSRPvP4t16GRmpepmAxScBLd7NOdfytUeTj8E+W1SoWV7YqYS+TFam6iuDSHm/5KKJe5tf+O8WFc4p7Py9883z8p25AL7l4Hz+2+XkjaTf04hceC3Bb2PlBEZ00ANgUmR4TqGaIGX+/v1RnLKWwWHZ4IeloS3QdLI8HfHDHpVSMzvn947+PlnDeVVhKeaFjUMn22qox4eE5XqaWS7TxJI5TMpRPWxwejqu6CkYRORGLTkOZa9hWuaQgk7vExOidVC655IDXFpKBofl1omQy8xYIAWXHvvuCweUC/leRs7xlcAhKxfdV3zAa/AGgIzEfjOceAaHuXA1jLEEXVp2lmFFDTaOCLMaFshaBKA3mKu0SbAWfJbmBYganM6bYwKpBVrY4UiQX/6M7O9HoqhTpg46xoj8LfQrqQOyv/BETD5vs5QQ5DfLBCC3EdZsUFYDNhSj0Ro+tphLiC24RBSuVXjA8K94KfNBg/4IdQVz5pBMJKSi+LoTLhpCHgblNLEB+E/xBAJ/AtqzhAjY4IKeHpYTuA+zQ+cI2sO4vwsU2m7v0bkdfJYhFFRjoELtfW1lBo/g+8BSZnBhWCKj/wt0YQgEBTUbbvb6iMKpw40gvLU5ZSzBWORocvQngu4lC7bv79aKwinansAjA7XShZS2GoN3XES5IhP0U0aFEw77iUVN/dTm97GcAiKoFUUi/xrajMStuaCSPRJljkE3oKTaU83nd9ubSBs4oCY4/gfc7UT0jW2MzxCTd0llbkiHsPSLurd50RSetgA7VuX1m4BOhaArcjxaQK/wEJaooCWV1qlQq4GjglJX4Hmg87DKmIkSvT7Y2VOJw3pwRmIkZFF0xO4eohP2qQ3i55tbO34EUPOnr4IoB1SWDhCdDsVseOSJ5IPA4FFBZCMYHUQHtO0hMoRmDXYJyx58hgnPItDD9xzawB1EB3v6W7ij5X9VBLMgTUkX33VRzyGebqNNdKTyVIHpfHQklGinmn14jeCYZLxDyVKJ2SdrjVjMfLg+JuJwClM79gSdqKwGgwWLcCCBpC9AfHDXIahbM6cDryBTRaaCiogWi9mU36m415Z3VLx37BcsFJDBLzAXNwuPaHkbenCFgjL5tA6/OiG/YCtdi5ZB6/hg+Whb2mRozCAoZfotE/uu4av0WRIAX9D92mkdCrqIw0NOwJ9ht27sqDeubmJ9U3LiMIdfC8zqgiPNeypMdCGIhJEsPhUOHmj6jHyNNVYxwUQhHcOh3kPEkqG5227HdKndx18rkK7VPdTM32J6w49CJzpIl7zSwgjCpS3uiSgEUwbmWtCM3mvRGv6WkndWc9+x6JKp4sd7M9EYTUxNH7mcT1ZL0LmumBUUgTZFdE/+Q17Puyn48hlRXKIehHtvhAQ0FuQ22k2VRfAA68i00DlEM1WNBAKtC2zAWYzNKuzQLxhdpFnOYa7k2VDHwWAJ+AySZ1VPMDsK6nyszjP2hTZKClQBFmzhO4oKqRgetHmZwhk3qqRieJjrCxTOXHE37TTOK5wKqhge1P0QeDgHtJ1qeDGncMrDiWYy+mR+HKJU7x5Dt4o2QgSTib7DMugrfijyNRCFVPFancPKfgLuMzx2sNI2QgS+tpu+yD5Y5xdAnadvu9GppBt6HtbnWGj4GDZCBOrPwsddwD0sf/7ysAv4CbP6XpqGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGxjH8D5HsoIGaPa9kAAAAAElFTkSuQmCC",
		PaymentMethodID: 2,
	})
	db.Model(&Bank{}).Create(&Bank{
		Name:            "Krungsri",
		Number:          "430-072-7783",
		Picture:         "data:image/jpeg;base64,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",
		PaymentMethodID: 2,
	})
	db.Model(&Bank{}).Create(&Bank{
		Name:            "Krungthai",
		Number:          "842-652-3265",
		Picture:         "data:image/jpeg;base64,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",
		PaymentMethodID: 2,
	})
	db.Model(&Bank{}).Create(&Bank{
		Name:            "Kasikorn",
		Number:          "475-869-9859",
		Picture:         "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBxIQEhUTEhATFhUVFhYSFRYWFhYYFxgWFRcYGBgWFhsYHSggGB0lHhYYITEhJSkrLy4uFx8zODMtOCgtLisBCgoKDg0OGxAQGy0mHyUrLS8wLS03Ly0tKy0tLS0tLSstLS4tLSstLy0tLS0tKy0tLS0tLS0tLS0tLS0tLS0tLf/AABEIAOEA4QMBIgACEQEDEQH/xAAbAAEAAgMBAQAAAAAAAAAAAAAABQYCAwQHAf/EAEAQAAIBAgMECAMFBgQHAAAAAAECAAMRBBIhBTFRYQYTIjJBcYGRobHRI0JScsEHFIKi8PEVJDOSYnSTsrPC0v/EABsBAQACAwEBAAAAAAAAAAAAAAABAgMFBgQH/8QAOxEAAgECAwUFBgQDCQAAAAAAAAECAxEEITESE0FRYQUiMnGhBhQzkbHBUoHh8EKC0RUjJHKSorLC8f/aAAwDAQACEQMRAD8A5YiJrjihERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAEREAREQBERAERMS9tfp/eG0tTPhsLWxNRUqEXKT4L68kursubMp8zeftNDVuA15X/Q3m6hha9XuU6j/AJQ7fITHvFwR1lD2Jxco7VepCHzl89F6s+5/P2n2Z1Nm4lRdqNUDiabgfKcnXMN+vqf1No3nNWL1PYivst0a8Jfk16pzOiJrWqD4+ltPib/OZq0umnoctjuzsTgZ7GIg4t6cU/Jq6fz8z7ERJPEIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiYVDa55fAg/rIbsrnqweEqYuvGhS8UnboubfRLN9EY1KgA03+HxBAIP95L9Hei9bGm/dS9jUblvCD7x+E29DejrY2rme4pL2mI0JPggPPx4DzEvXS/FHB0aD0gAKdQAKNFy5GBXTwtJp0tvvz0/fofR61bD9hYd4fCq8krylx83zfJaRWSzOnZHRHC4YC1IVGH36gDG/EDcvoJPqANBI/Y22KWLp56Z5Mp7yngwklNhGMUu6aWdeVd7yUtq/HU+SN2jsXD4kfa0UY/itZh5MNRJOcmOxtOgheq4VR4n5DieQhpPUhTcHtJ2tx09TznpH0CemC+GJqDUmme+B/w27/lv85TqbncTr4b+U9O2Dt843HGwK00pMEB3m7LctzNt3hOP9oPRcOrYmitmGtUD7w8XA4jx4jy18NWgrbdM3OB7Uo9o0nhsYlKEnZSfNZLPz8Ml5soitf5xNFF/687Cb5SMtpXOF7Z7Ln2binRlmtYvnH+q0fXPiIiJJqhERAEREAREQBERAEREAREQBERAEREANOeoCxAAvraw8+yPab33W5E+wM7+iOHFXHUVtcZyx/gGf5rMc82o8zufYuioe8YtrwRsvWUv+KXlfmer9HNljC0EpDeBdzxc94++nkBIb9pKE4VSBotRS3IEEA+5A9Zb5z4vDJVRqbi6sLEcps3FbOyjBilKvGak85Xz6v8AU8V2dtCph3FSk5Vh7EcGHiJ6j0Z6TU8YMpGSqBdk48Sp8R8RKP0l6KVMIS63ej+LxXkw/wDbd5SO6OVCuLoEHXrFHoxykexM88JShLZZz2Hq1sLVVOSybV15u11+/NcvQ+k/SlMJ9mgD1SL2v2V5t9PlPNtpbSq4l89Vyx8OA5KNwmG0a5qVXdtSzFj6k/29JvwmzCy9Y7CnS3Z2F7keCKNWPlpxIlZzlN24FcRiKuJm4rwrhwtzf9XkT/7MkJxFQ+Ap2P8AEVt8j7T0tluLGef9B8arYjqqSFKa02YXILOSyjMxGl7bgNBPQp6KXhNx2ckqCSd83+v/ALx1PEelmy/3TFOgHYJFRPyMdB6WI9JwUz4eXx3/ADl4/avhR9jV8e1TPwYfMyi0tw/rco+k8DWxVcTee0cPeuyKWKl4oSs30d4v5tRfmbIiJY+eCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAPj5N/2mTHQBgMfT59aB55G+khn3eenvpNuxcX+74qnUO5alz+W9m+BMpLuzjI772PlvMJiqC8Vk/O8ZL6r1Pd4nwGfZtDzGJFxY+MoO2ejyYfGYerSFqdSqqlfBWzA9nkRfTwtPQJB9KB2aH/MUvmZWSujBiKcZxV+DTXTM84pU6dBOuqKKjOzrSpnu2Q2Z38WFzYDxsbyNxmLes2eoxY7uQHgFG5RyE7tr3NLDAXOlWwHE1amgll6OdBiQKmK0G8Ugdf4z4eQ9/CeXZcnso55UalaW6prJWb5XaTu/t6LU1/szwTGpUrWIS2QHixIJA8svxno01UqSooVVCqBYACwA5ATbPVGOyrHQ4agqFNU07lH/aq/+Xpjx6wn2RvqJ5vR3Dz/AElw/anjs1anSB7iFj+Zzp8FH+6VGlut6+4H6gzX1c6z6G37YkqXs+oy1nNW/wBTl9ImcREk+dCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAmiuNb+O/3JP9ec3zFl/rz0ErON0brsHtT+zsZGrLwPuy8nx84uz56panqPQHbYxFAUy32lIBObJuVv0PMc5YsVjadK2dwCdw1LH8qjU+gnhmDxb4dw6MQb20LC4vqpykEg+cuWzdpJXHZ0Y6Mvjfyt2xwv1h5CenD19pbL1R2naeA3ct/SzpyzutFfP5Pg/y4FzHSCle2Wp7KT/tDZvhOXbWJWstE0jntiKRYC91sfvDevrIggkZbeHd1P8ALr/4/Sa1Pj+HS+63LNcZfLOn5ZnuaeULqxh0bwv2uHaqhC06dd7sCFVutexJOm5iZaa236K7szDwYWVT+VnKhvS8rTVM+ubNbXvZ7c/9Srl87r5iKRt2hpf717X/AIgy5v8AqPIWRjo0FSjsp3/RJfYtuF2rSqEANlY7lYWJ/L4N6EzLamOTD0mqubKguefADmTYespWLqLTBZyFG83Gh3a6r2vO1TzEqO29stiDkUsKY7qlm1P4iCxAPlaUq1lTXU2mA7Oni6ll4Vq/surOTH4xsRXeq+92LHfproByAsPSB9B7TXSpj4fO3tNk8UI2zerNH7VdqwxeIjQofDpZK2jlo7dIpbK/mtdO4iIlzlRERAEREAREQBERAEREAREQBERAEREAREQDFk/r9L8rzUudDmUn0v8ACb4K+XrrKygpHR9je0eI7OW6ktul+F8Oey/qmtl9Hm7JsnpuFUU8ThqTpuuqqrW5rbKfS0lf8dwDkMtZ6J3WdGYDkGQ5lHINblKE1McCfK39fGY/u/p5k/8AzLKrVj1Omh2l2Diu85SpPk019pR+TXkegPtjBb6mMzW1ASnVv6GsXA9Lec58Z03w9MEYbDZmP3qtiTzOpZvUiUf93HFfc/os+pR5H5n5Rvqz4W/fUmWL7Aod51XPotp/SKX+43Y7H1cQxZtATfKq2APJRp675pp0xbd/XqJsCjwt8j9Z9lFDO8s2aLtX2qqYin7vhI7uno/xNcssop8dnN/iEREuckIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiSmz+juKrpnpUCVO4kqoPlmIvJSb0LwhKbtFNvpmRcTbicO9NijqVZdCp3iapBR5ZMREnML0XrVDRyEEVg7ZtbIENjm4nUe8sot6GSnSnU8Kv+rsQcTox2GNKo9MkEozISNxym1x7TnkMo007MRESCBERAEREAREQBERAERBgG/BYKpXbJSps7b7AeHE+A9ZvxmyMRR/1KDrcgC+oJO4AjQnlLD+zzG2qPQKKQ6lma5DdkaDTw1PvG2MeThsPiqRdA1Rr0i7OmemxysM3NeW+ZVCLhe574Yak6O25O+eXk0nrra64q/Arm0NmVsPbrqTJm3E21tv3X15TGvs+qjrTamwdwGVd5Ia9iLcbH2lt2kj4/B0TRo2epWdmAZmUHtBmLN3Rf6CbOmGBr0nTE0agBo0kRsp7Q1YXsRqDmtLOks2tMv1LTwcUpTjdxtFp9HryzSt89Cp4/Y+IoANVosgJsCbEX4aE29ZlgNh4muualQdl46AemYi/pJbZ2Ieps7GZ3Z7NRYZiTYlxci+6WnHYWkVSpUdBRy01pK1VqYVMt2Zctsz8ByiNNPNFqWChU7ybtZO2V9WtbWtly4nnBwNUVOq6pusvbJY5r+X6zbtHZNfD2NaiyA6AmxBO+1wSLyxdJsc6Jg69Kqwdkqp1mmZlDKFzab7G/rNWAxTHZ1VqhNTq69JgGJOmdSRc30Nz7mQ4RTaMTw1PalC7uldPK1rJ+fHhZfMiKXR/FsgdcM5U6g6ajiBv+E0YbZVeq7U0oszqbMB90jwYnQS9bE2k2PqZzhCiXKNVGIYWOW4AUWud27jI7optBgMYr9tKYNQgk5mbdq177l+Mtuo5GZYKjJw2ZO0r56aK/GKa9SqbS2TXw9uupMl9ATYi/mLi/KW+ntHFYhEZdlI6KAFJa1gPwXt7gTg2biWq7PxYdiQhpMgYk5bvuBbW2km8QLUaVauazXRWatTfKKQNrHqlIBA8TlP6RBLVaF8PRSvKEmouKekebWd4vS2WS62sQ3TSlVq1aANNeuZCGpoSzAXuM2nM+xkJjNgYqipeph3VRvN1IHnlJtLpiazU9q0ghX7amFckbwqs1xbcTlGs4tg7QartKogzCkwqKULFlOUWvruJtfTjJlFSlnrexNbD051XtN3ctnhyVm+fXT0KngdkV8QCaVF3ANiRa1+FzpeWno5tLEYU0cHUpMherdS2/ITcqBzIOvMzn2LttnpphaeCD9WCbis1MaalmIAtqfE+M6tpUsm1MMoLEdlgGdnylgbhSxOmgkRSik0ymHhGnGNWlJ3vFPlm81mvo2QG1MDVr42ulKmzt1jmw4ZjqSdBOOlsmu9RqK0XNRdWTS4Atvvp4iT+0nq4PFV2qUQ9LEl179gy5r6MmqsL7ucsW1NpJSxlKjSFqlZ6ZrN4hVsFT1t7ecbuLvfmR7pTnKUpyae1mv8AM3s2y45O93xyyz8+wmy69Z2p06TM694XAy2Ntb2A1mP+G1rVD1bfY5etvYZM17XB8jul32W6HEbQSooFPMHZyxVQVY2DEEEcdOBnDSql12o5dXJFA5k7p0O6/Dd6GRu1bXn9yvuVNKPeebkuH8Knw/lWej6FLiImA1oiIgCIiAIiIAiIgEnsLbbYNmdKSMWAW7AmwF91iLXvr5CfdsdIHxKJTNJKaUyWVUVgLnx38z7yLLAaXI8NLW0/vAqDi3uPpDqTS2eH5G0pe7uioTxLjfWO7btxtfaV+ZK0ekdanhv3ZOypJJYAhiGJJXfoJtxXSmpUpGkaagsqo1QL2nVO6pMhA+/vab9Rp56aT7n5v7r9JO8ny9UZV7rayxj0t8OWnLxZfXqduF2q1OhVoBbrVyliQ1xk1FtZljdsvVoUqDAZaV8pGa5zcdfDdODrBxPuPpPhf8/uv0ld5O1reqKbvB7Oz727Wt8KWl728XPM7sdtZqtKjSZbCkHVSA1yHIvm9vCKW1mGGfDhew7KxNjmuuXd4W0nEH5v/L9I6wcW91+kbyd729UTu8Hdv3t5q3wpaaW8XIl9k9JHwyZBSR8rF0LqSUcrlLLblOXA7XeitZQAevTq3JGo4kW8dTOLP+f4fSDUHE/7h9JKqT5adUTs4RW/xb7un91LK6t+Ll6ZHbg9rtSo1aIW61cmYkG4yG4trOt+k9Q0TR6tblBRNXL2zTG5L7pD5+b/AA+kdYOLe4+kbyeiXqhGOEirRxb0t8KWnLxExU6TVGxFPEZBmpDKBrlIyka+P3jNGzNtvQrmuqXYljYg5e0NdxvI4VBxb3H0jrBxb3H0jeT1t6om2Fvte9u97/ClrpfxckdmydqNhnZggcOrK6MpKsrWuD7TbituVKmIXEWysuTIFXsqF7q28R9ZG5/z/wAv0n3rBxPuPpI3k7Wt6oqoYNRUVi3ZO/wpa318V/XXMlNr7fbEBV6taaqWYKq2BZjdmN+J+ZnzF7fepiVxRQB1KEAA5ezoPG8i8/5/5fpPuf8AP7r9JO8ny9UTJYVu7xb1T+FLhp/FwJrCdJ6lOpWqdWjCvY1EYEppfdrzMyxfSlqlJ6SYalSV7ByiFSQDe2+Qefm3uv0nzrOb+4+knez0t6om+GUXFYx2zy3T46/xXzf1PsQIkGmQiIgCIiAIiIAiIgFh6GVwta3WVFJDEhFBUqtNjdiWBFju0321mzpFjBUw9D7as+ZWYB0FmK1XXM1mNja4Fr6AelcRyL2JFwVNja4O8HiOUPUJABJIUWUE6AEk2HDUk+sybfdsetYpqi6VufPi19k/3e972dtUrR7IqOBTpXPWLe5qKhCabwSO9bxHORmx8VTO0HLZmzgqp0cghDfUMANFIuL8OcrhxdS2XrGsbAjMbWU3UWv4HdPgxD5s+ds9757nNfjffL77ToZ549tw17rT+t/r80n53NsSHwhy3Oak76jKAvXscxGa1+zlsATdlF5D4bEijQwzhWJzVrBGKnTKGJNmue0vhoBaRD7QrEEGs5BzXBZrHP3r6+PjxmujiaiFStRlyklbEjKW0JXhfxkOrd3Kzxqk01e6ja+XNO9rvlpck+lh+3Bu5ulN+02Zu0ga17Dde3pLHSS2FpvRAOJWhdQRqELHOyj7zDd6ykV6zVGzOzMx3liST6mZJi6gKkVGBQZUIYgqOC66DXcJCqJNsrTxUY1Jzs+95XWevmuHWxa+iGERaJark/zLGl2nCnq1DBmS57RLEDThHRyk9HrcOVqU3NQUxXVQ4zKO6ykd079PxSoVazPbM7Nl0W5Jtc304azoTaddSxFeoC/eIZrtbTU31kqolboWpYuENhbL7uV1xus/LvcnfqWHDUcmGRTYsuOyEjxsgBtyvJfalFOuRVCmhUrlcU33s4PZRvwrutxvz1oK13Ayh2ABzAAmwO64HHTfvmRxdTt/aP8Aad/tHtfm17XrCqJK1iyx0VHZ2XovRafno+jZdKNTEVqmIpYtAMOiubFVVaeXulCBv9TNjaYFTbOBhx9iFW4ZnIFbMdQFt4SlV9o1nUI9Z2UblZmI03aEzFMbVFrVGFhlFmYWB3ga6DlJ3q6+hZY+KvlJ3TzbTeb+Vlw4a+RbcRQp16eDpGy1upVqTnc1m1pt5gXB4/HvoMEr49u0AtTD91A7WO9Qp8CNDyM8/euzZbsxygBbk9kDcF4ek3ptKsCzCs4ZrFiGYFrbrkHWFVXL92sVWPipX2eP/Vx5rPPXirXLVs/Z/U4zFViKYFBvswzBEzVRcLdrAWVjpzmeEwi4XEYs5FekaHWqv3WRmBsD7gHlKW1ZyCpdipbOQSSC34jxPObP32rbL1rWtktmNsv4bX7vLdI3qWi6kRxkI2tF5Ny4PN316WaXXZWhb1waU8PRNMh6b4ykyE97KVIIbgQQQRymvpopRG6gDqnqt17DvdaGOVT+EAbuPzqK4moFCh2CghwLmwYaAjnzmRxdTtDrH7Zu/aPaPFte16xvE1awnjYunsKLWSWXC3Lo9H0v+eiIiYTXCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIAiIgCIiAIiIB//Z",
		PaymentMethodID: 2,
	})
	db.Model(&Bank{}).Create(&Bank{
		Name:            "NULL",
		Number:          "NULL",
		Picture:         "NULL",
		PaymentMethodID: 2,
	})
	var bnull Bank
	db.Raw("SELECT * FROM banks WHERE name = ?", "NULL").Scan(&bnull)

	// ===============     ช่องทางการชำระเงิน     ===============
	db.Model(&PaymentMethod{}).Create(&PaymentMethod{
		Name: "Cash",
	})
	db.Model(&PaymentMethod{}).Create(&PaymentMethod{
		Name: "Transfer",
	})
	db.Model(&PaymentMethod{}).Create(&PaymentMethod{
		Name: "Crypto",
	})
	var Cryptox PaymentMethod
	db.Raw("SELECT * FROM payment_methods WHERE name = ?", "Crypto").Scan(&Cryptox)

	// ===============     ตารางหลัก     ===============
	db.Model(&Payment{}).Create(&Payment{
		Customer:      Customer1,
		PaymentMethod: Cryptox,
		Bank:          bnull,
		Crypto:        DogECoin,
		Place:         CS2,
		Time:          time.Now(),
		Picture:       "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAASABIAAD/4gHYSUNDX1BST0ZJTEUAAQEAAAHIbGNtcwIQAABtbnRyUkdCIFhZWiAH4gADABQACQAOAB1hY3NwTVNGVAAAAABzYXdzY3RybAAAAAAAAAAAAAAAAAAA9tYAAQAAAADTLWhhbmSdkQA9QICwPUB0LIGepSKOAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAlkZXNjAAAA8AAAAF9jcHJ0AAABDAAAAAx3dHB0AAABGAAAABRyWFlaAAABLAAAABRnWFlaAAABQAAAABRiWFlaAAABVAAAABRyVFJDAAABaAAAAGBnVFJDAAABaAAAAGBiVFJDAAABaAAAAGBkZXNjAAAAAAAAAAV1UkdCAAAAAAAAAAAAAAAAdGV4dAAAAABDQzAAWFlaIAAAAAAAAPNUAAEAAAABFslYWVogAAAAAAAAb6AAADjyAAADj1hZWiAAAAAAAABilgAAt4kAABjaWFlaIAAAAAAAACSgAAAPhQAAtsRjdXJ2AAAAAAAAACoAAAB8APgBnAJ1A4MEyQZOCBIKGAxiDvQRzxT2GGocLiBDJKwpai5+M+s5sz/WRldNNlR2XBdkHWyGdVZ+jYgskjacq6eMstu+mcrH12Xkd/H5////2wBDAAcHBwcHBwwHBwwRDAwMERcRERERFx4XFxcXFx4kHh4eHh4eJCQkJCQkJCQrKysrKysyMjIyMjg4ODg4ODg4ODj/2wBDAQkJCQ4NDhkNDRk7KCEoOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozv/wAARCAgABZgDASIAAhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSExBhJBUQdhcRMiMoEIFEKRobHBCSMzUvAVYnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwD0ilpKK988kloooqQCiiigAooopgFFFFBoFFFFABRRRQA+iiipAKKKKACiiigBaKKKQBRRUlABRRRSAKKKKACiiigAooooAKKKKAH0UUUgCiiigAooooAKKKKCxR1qSox1qSpAKKKKACiiigAooooAKKKKACgUUCgB9FFFSAUUUUAFFFFABSjqKSlHUUASUUUVIBRRRQAUUUUAFFFFABRRRQAo60+mDrT6RoFFFFABRRRQAUUUUAFFFFAAOtPpg60+kwFooopAFFFFABRRRQAUUUUAKOtOpo606kwCiiikAUUUUAFFFFABSjrSUo60APoooqRoKKKKDQWiiigQUUUUAFFFFACjrTqaOtOpMAooooAKKKKACiiigApy9abTl60gHUUUUhhRRRQAUopKUUAFFFFIQUUUUDQop1NFOoZYUUUUgCiiigAooooAcvWnU1etOpMAooopAFFFFABSikpRQAUUUUgCiiigBy06mrTqACiiigAooooAKKKKBrcKKKKChVp9MWn1LAKKKKACiiikAUUUUgCnCm04UALRRRQAUUUUAFFFFABRRRQB55RRRXunkhRRRQAUUUUAFFFFMAooooNAooooAKKKKAH0UUVIBRRRQAUUUUALRRRSAKkqOpKQBRRRQAUUUUAFFFFABRRRQAUUUUAPooopAFFFFABRRRQWFA60UDrQBLRRRUgFFFFABRRRQAUUUUAFFFFABQKKBQA+iiipAKKKKACiiigApR1FJSjqKAJKKKKkAooooAKKKKACiiigAooooAUdafTB1p9I0CiiigAooooAKKKKACiiigAHWn0wdafSYC0UUUgCiiigAooooAKKKUdaQAOtOoopAFFFFABRRRQAUUUUAFKOtJSjrQA+iiipGgooooNBaKKKBBRRRQAUUUUAKOtOpo606kwCiiigAooooAKKKKACnL1ptOXrSAdRRRSGFFFFABSikpRQAUUUUhBRRRQAop1NFOoY0FFFFIsKKKKACiiigApy9aF606kwCiiikAUUUUAFKKSlFABRRRSAKKKKAHLTqatOoAKKKKACiiigAooooGtwooooKFWn0xafUsAooooAKKKKQBRRRSAKcKbThQAtFFFABRRRQAUUUUAFFFFAHnlFFFe6eSFFFFABRRRTAKKKKDQKKKKACiiigAoop9ABRRRUgFFFFABRRRQAtFFFIAqSo6kpAFFFFABRRRQAUUUUAFFFFABRRRQA+iiikAUUUUAFFFFABQOtFA60FktFFFSAUUUUAFFFFABRRRQAUUUUAFAooFAD6KKKkAooooAKKKKAClHUUg61JQAtFFFSAUUUUAFFFFABRRRQAUo60lKOtAD6KKKRoFFFFABRRRQAUUUUAFFFFACjrT6YOtPqQCiiigAooooAKKKKAClHWkpR1pAOooopAFFFFABRRRQAUUUUAFKOtJSjrQC3H0UUVJqFFFFAC0UUUCCiiigAooooAUdadTR1p1JgFFFFABRRRQAUUUUAFOXrTacvWkA6iiikMKKKKAClFJSigAooopCCiiigBRTqaOtOoY0FFFFIsKKKKACiiigBy9adUdOXrSYDqKKKQBRRRQAUopKUUAFFFFIAooooActOpq06gAooooAKKKKACiiiga3CiilWkUC0+iikAUUUUAFFFFIAooopAFOFNpwoAWiiigAooooAKKKKACiiigDzyiiivdPJCiiigAooopmgUUUUAFFFFABRRRQAU+iikAUUUUgCiiigBaKKKQBRRRQAVJRRSAKKKKACiiigAooooAKKKKACn0yn0gCiiigAooooAKKKKACgdaKUUFklFFFSAUUUUAFFFFABRRRQAUUUUAFAoFPoAKKKKkAooooAKKKKAAdakqMdakpMYtFFFIQUUUUAFFFFABRRRQAUo60lKOtAD6KKKRoFFFFABRRRQAUUUUAFFFFACjrT6YOtPqQCiiigAooooAKKKKAClHWkpR1pAOooopAFFFFABRRRQAUUUUAFKOtJSjrQC3H0UUVJqFFFFAC0UUUCCiiigAooooAUdadTR1p1JgFFFFABRRRQAUUUUAFOXrTacvWkA6iiikMKUUlKKACiiikIKKKKACiiigaFFOpop1DLCiiikAUUUUAFFFFADl606mr1p1JgFFFFIAooooAKUUlKKACiiikAUUUUAOWnU1adQAUUUUAFFFFABRRRQNbhRRRQUKtPpi0+pYBRRRQAUUUUgCiiikAU4U2nCgBaKKKACiiigAooooAKKKKAPPKKKK9480KKKKACiiigAooooAKKKKACiin0AFFFFSAUUUUAFFFFAC0UUUgCiipKACiiikAUUUUAFFFFABRRRQAUUUUAFPplPpAFFFFABRRRQAUUUUFi0DrQOtSVIBRRRQAUUUUAFFFFABRRRQAUUUUAAp9MFPpMAooopAFFFFABRRRQAo61JUY61JUgFFFFABRRRQAUUUUAFFFFABSjrSUo60Gg+iiikAUUUUAFFFFABRRRQAUo60lKOtAD6KKKkAooooAKKKKACiiigApR1pKUdaQDqKKKQBRRRQAUUUUAFFFFABSjrSUo60Atx9FFFSahRRRQAtFFFAgooooAKKKKAFHWnU0dadSYBRRRQAUUUUAFFFFABTl602nL1pAOooopDClFJSigAooopCCiiigAooooGhRTqaKdQywooopAFFFFABRRRQA5etOpq9adSYBRRRSAKKKKAClFJSigAooopAFFFFADlp1NWnUAFFFFABRRRQAUUUUDW4Uq0lKtIofRRRSAKKKKACiiikAUUUUgCnCm04UALRRRQAUUUUAFFFFABRRRQB55RRRXvHmhRRRQAUUUUAFFFFABRRT6ACiiipAKKKKACiiigBaKKKQBRRRQAVJRRSAKKKKACiiigAooooAKKKKACiiigB9FFFIAooooAKKKKACiiigsUdakqMdakqQCiiigAooooAKKKKACiiigAooooABT6YKfSYBRRRSAKKKKACiiigBR1qSox1qSpAKKKKACiiigAooooAKKKKAClHWkpR1oNB9FFFIAooooAKKKKACiiigApR1pKUdaAH0UUVIBRRRQAUUUUAFFFFABSjrSUo60gHUUUUgCiiigAooooAKKKKAClHWkpR1oBbj6KKKk1CiiigBaKKKBBRRRQAUUUUAKOtOpo606kwCiiigAooooAKKKKACnL1ptOXrSAdRRRSGFKKSlFABRRRSEFFFFABRRSr1oGgFOoopFhRRRQAUUUUAFFFFADl606mr1p1JgFFFFIAooooAKUUlKKACiiikAUUUUAOWnU1adQAUUUUAFFFFABRRRQNbhSrSUq0ih9FFFIAooooAKKKKQBRRRSAKcKbThQAtFFFABRRRQAUUUUAFFFFAHnlFFFe8eaFFFFABRRT6AGU+iikAUUUUgCiiigAooooAKKKKAFooopAFFFSUAFFFFIAooooAKKKKACiiigAooooAKKKfQAUUUUgCiiigAooooAKKKUUFgOtSUUVIBRRRQAUUUUAFFFFABRRRQAUCigUAPoooqQCiiigAooooAKUdRSUo6igCSiiipAKKKKACiiigAooooAKKKKAClHWgdafSNAooooAKKKKACiiigAooooAKUdaSlHWgB9FFFSAUUUUAFFFFABRRRQAUo60lKOtIB1FFFIAooooAKKKKACiiigBR1p9MHWn0mNBRRRSNAooooAWiiigQUUUUAFFFKOtAAOtOoopAFFFFABRRRQAUUUUAFOXrTacvWkA6iiikMKUUlKKACiiikIKKKKAClXrSUq9aBodRRRSLCiiigAooooAKKKKAHL1p1NXrTqTAKKKKQBSikpRQAUUUUgCiiigAoopy0AC06iigAooooAKKKKACiiiga3ClWkpVpFD6KKKQBRRRQAUUUUgCiiikA4UtIKWgAooooAKKKKACiiigAooooA8/ooor3DzQooooAKKKKACiiigAooooAKKKKACiiigBaKKKQBRRRQBJRRRSAKKKKACiiigAooooAKKKKACiiigAp9FFIAooooAKKKKACiiigsWgdaKB1qQJKKKKACiiigAooooAKKKKACiiigAoFFAoAfRRRUgFFFFABRRRQAUo6ikpR1oAkoooqQCiiigAooooAKKKKACiiigBR1p9MHWn0jQKKKKACiiigAooooAKKKKAClHWkHWn0mAtFFFIAooooAKKKKACiiigBR1p1NHWnUmAUUUUgCiiigAooooAKKKKAFHWn0wdafSY0FFFFI0CiiigBaKKKBBRRRQAUo60lKOtADqKKKQBRRRQAUUUUAFFFFABTh1oXrTqTGFFFFIApRSUooAKKKKQgooooAKVetJSr1oGh1FFFIsKKKKACiiigAooooAcvWnU1etOpMAooopAFKKSlFABRRRSAKKKKACnLTactADqKKKACiiigAooooAKKKKBrcKVaSlWkUPooopAFFFFABRRRSAKKKKQDhS0gpaACiiigAooooAKKKKACiiigDz+iiivcPNCiiigAooooAKKKKACiiigBaKKKQBRRRQAUUUUAFSUUUgCiiigAooooAKKKKACiiigAooooAKKKfQAUUUUgCiiigAooooAKKKKCxR1qSox1qSpAKKKKACiiigAooooAKKKKACiiigAFPpgp9JgFFFFIAooooAKKKKAAdakqMdakpMYtFFFIQUUUUAFFFFABRRRQAUUUUAKOtPpg60+kaBRRRQAUUUUAFFFFABRRRQADrT6YOtPpMBaKKKQBRRRQAUUUUAFFFFACjrTqaOtOpMAooopAFFFFABRRRQAUUUUAKOtPpg60+kxoKKKKRoFFFFAC0UUUCCiiigApR1pKUdaAHUUUUgCiiigAooooAKKKKAHL1p1NXrTqTGFFFFIApRSUooAKKKKQgooooAKVetJSr1oGh1FFFIsKKKKACiiigAoopy9aABetOoopAFFFFIApRSUooAKKKKQBRRRQAU5abTloAdRRRQAUUUUAFFFFABRRRQNbhSrSUq0ih9FFFIAooooAKKKKQBRRRSAcKWkFLQAUUUUAFFFFABRRRQAUUUUAef0UUV7h5oUUUUAFFFFAC0UUUgCiiigAooooAKKKkoAKKKKQBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABT6KKQBRRRQAUUUUAFFFFABSjrSUooLJKKKKkAooooAKKKKACiiigAooooAKKKB1oABT6KKQBRRRSAKKKKACiiigBR1qSmDrT6kAooooAKKKKACiiigAooooAKUdaSlHWg0H0UUUgCiiigAooooAKKKKACiiigBR1p9MHWn1IBRRRQAUUUUAFFFFABRRRQAo606mjrTqTAKKKKQBRRRQAUUUUAFFFKOtAAOtPoopDQUUUUjQWiiigQUUUUAFFFFABSjrSUo60AOooopAFFFFABRRRQAUUUUAOXrTqavWnUmMKKKKQBSikpRQAUUUUhBRRRQAU4daQdadQNBRRRSLCiiigAooooAKcvWm05etADqKKKkAooooAKUUlKKACiiikAUUUUAFOWm05aAHUUUUAFFFFABRRRQAUUUUDW4Uq0lKtIofRRRSAKKKKACiiikAUUUUgHClpBS0AFFFFABRRRQAUUUUAFFFFAHAUUUV7Z5oUUUUAFFFFABRRRQAUUVXuLiCyglu7uTZHHQBY/p1qKW8tYTiWRAfdh/jXhuv+NtT1J3hs2+zW38O377/AFNcQT5j75X3Ser80GftD6oiu7SX/Vzr/wB9D/GrXbPb1r5OiLxyb4+K63RfGmr6TLtkkN1B/cl5xSD2h9CUVQsL6LUrSK8thiOQZwav0GgUUUUAFFFFABRRRQAUUUUAFFFFABT6KKQBRRRQAUUUvNACUUvNJzQAUUc0tBYDrUlMANPqQCiiigAooooAKKKKACiiigAoooxu+XpmgApcGuH1zxzpWjym1X/S5RxtjOMH3Nc/b/FCISEXNidvs5JH4UiPaHrNLWZperabrFsLnTZRNH/Fjhkb0PtWnSLCiiigAooooAKUdRSUo60ASUUUVIBRSHpVO8vLfT7WS9uuIolLN9BQBcoyPWvBtU+JesXDuNOMdvEDtU4yx/OsX/hPfFX/AD+n/vlf8KrlI9ofSWR60oIr5s/4T3xV/wA/p/75X/Cl/wCE78VEhftp5I/hXP8AKiwe3vofSdFcT4E1XUdW0Z7jUpGkmE7AkqBxXbVJoFKOtA60+kWFFFFABRRRQAUUUUAFFFFABSjrSUo60APoooqQCiiigAooooAKKKKAClHWkpR1pAOooopAFFFFABRRRQAUUUUAFKOtJSjrQC3H0UUVJqFFFFAC0UUUCCiiigAooooAUdadTR1p1JgFFFFABRRRQAUUUUAFFFFADl606mjrTqTGFFFFIApRSUooAKKKKQgooooAUdadTR1p1DGgooopFhRRRQAUUUUAFOXrTacvWgB1FFFSAUUUUAFKKSlFABRRRSAKKKKACnLQtOoAKKKKACiiigAooooAKKKKBrcVafTFp9SygooooAKKKKACiiikAUUUUgHClpBS0AFFFFABRRRQAUUUUAFFFFAHAUUVJXtnmkdFSUUgI6KkooAjqSiigArzD4k38kdrBp0fSVy8v+4nAr0+vL/Hmg6rqc8N3Zp5sccTBwnHOaDOZ5PZWcmoXcNjDw0smwn2r3Gx8A+H7SPZdxPcS/8ATQ7f5V4XIlxaS7ZFkilT1ODXWaV4713Sx5Mzfaof7snJ/wC++tBED0y+8CeHbqP93Abd/WNyo/8AHs1zGlfDwpqZbU332kf3CnHmfWu80HxBa+ILb7VafIY+JI35reoL9mNiiEapHGgjjjHAFOoooNAoo5oznpQAUUUUAFFFFABRRRQAU+mU+kwCiiigAooooABzwKxrzxFomnzm3u7mO3kHUMSa2R1r5+8exyN4nnZVYgrFyASPuUBUPYf+Eu8M/wDQQh/M0f8ACXeGf+ghD+Zr5s8mf/nm/wD3yf8ACjyZ/wDnm/8A3yf8Kkz9ofSf/CXeGf8AoIQ/maD4t8MEYOoQ4+pr5s8mf/nm/wD3yaDBPj/Vv/3yaA5z6nsNVsNVUyafMk6pxlWIrRry34ZI6W15vUrmUdQR2969SoNIBRRRQWFFFFABRRRQAUUUUAKBk4rn/E9/Jpvh27vkOJFjAU+hkOK3vlIwenesrXNPGsabc2PBaVf3f+8BxSB7HzZpVg+salBYq2HmbLv6ZPJr0DxJ4DtdI0t9RtLkyPCyo+e4PWvPIHvdF1BJh8k9vJhQe5Brqde8ealr1gbKaNIYmILkDliKZyjfAF/LZ+IIYF+5cAxSf7y/dr6IrwP4c6Hc3WsJqcgIhtlOc92P3cete+UmdEAooopFBRRg0UAKOtSVGOtSVIBRRRQAHkYrzH4n3csei2sA+7NJg/8AAK9OzjmuT8aaNLrmjSxQDdLE29B6kUdRS2Z4T4T0y31jXLa0uxmJ/mYHuB2r17XIPBfhtYm1CwXM/TCtzj15rwrzLzTZ95zDNE/B6HI+tdDYWuseM9VSC8kklC4Dynoi+uauxzwPWdEsfBniSCaaw09B5ThCCrDg/jXEaxrfh/TNTubBdCglEDBAxZh+ua5m6bXPB2pz2sUpiLHG/wDvL9awo1ubycBB9ondsgdck0WL5z6K8E31pqOlSXdlapaRGZgEQ7ua6+ua8I6K+haDDZTcz582X2L10tZo3gKOtPpg60+gsKKMGigAooooAKKKKACiiigApR1pB1p9JgLRRRSAKKKKACiiigAooooAKUdaSlHWkA6iiikAUUUUAFFFFABRRRQAUo60lKOtALcfRRRUmoUUUUALRRRQIKKKKACiiigBR1p1NHWnUmAUUUUAFFFFABRRRQAU5etNpy9aQDqKKKQwooooAKUUlLQAUUUUhBRRRQAo606mjrTqGNBRRRSLCiiigAooooAKcvWm05etADqKKKkAooooAKUUlKKACiiikAUUUUAOWnU1adQAUUUUAFFFFABRRRQAUUUUDW4q0+mLT6llBRRRQAUUUUgCiiikAU4U2nCgBaKKKACiiigAooooAKKKKACiiigDhKKKK9o80KKKKACiiigAooooAKKKKAMPX9OsNQsp3uoI5Hjj4fH7zp618zDtX1Rqf/IOn/65H+VfK46ChGdQ9W+GX/HxfJ/sivX68d+GX+vvf90V7FQFMKa7pGu+Rgo9WOB+tcB4p8apo7mw0xBNd/xZ6J9K8fvdU1HUG33s8jj+6TlPyoDnPoyTWtFj/wBZdwr9HzU8WoWN1/x73EUn+6QK+V/lHpUkbOjGSL8x2oD2h9ZdelFeCeHvGupaTLHb30puLcDp/EPxr3Kzu7e/t47yzffHIKDT2hZqPz4RwZE/76H+NSdq+VbtpPtEvL/61u/+0aBVKh9TrJG52oyk+gIJp9eD/Dx3fxMiuX4hl7/SvezQOmISFBZiAB1J4qPz4P8Anon/AH0P8ayPFP8AyL+of9cW/pXzOWfafmfp60kFSofWSyI+djBsehB/lTq8s+GRJgvcvv8AnT+VepnoaAGNLEh2u6qfQkA0nnwf89E/76H+NeI/EhiviBMOV/dDpXC2rP8AaYuW++vf3FSZ+0PqymtHE33lR/fFOHRfov8AKiqNCPyYP7qflTWS0TG/y1z0zx/OpunNeY/E13j0+yZe87fyoA9JWO1c/J5bfTmpTDBg/Kh/CvlO3vb62YyW87x/7rmvUvC3jueWeOx1sgKcIkw6EngBvf3qTP2h60kQXouypGZVG5iFHqTgU0DnFcj46+Twrdj/AK5/+hUGh1nnwf8APRP++h/jT1dHGUYMPYg/yr5KDPn7zj3zXuHwxJOiXGW3/vz/AOgrQHtD0emPJHHjzGC56ZIH86fXlfxSeQQWAXpvk/lQWz0/z4P+eif99D/GpAQwBUgg9COa+SQ7gggtx719L+Fix0HT2fvAn8jQRTqHQYNGeMgj8elZ2o6na6TYSX985CrnCjqcdhXhOveNNV1iVkif7PbgELHHwSPVjSFUqHvM2q6ZbEpc3USH6iqyeINEdgi6hBkHj5wf6V8tFwSXZgWPUv1/OkBGRjFFiPaH0/qOh6F4hXfcLHIccSREZ/TkVw+peHPAnhuRZNTMzMTwm9mz7Fcc15Ja3l3ZP5tjI8bA5JU4HFSX2oXerXT3d2xkkkwPfinYPaHtMHxD8M2kS2tsjxRp9392Rj9a7DRdYsdetDfWOTEDjkfxV8vC3n3AeW//AHyf8K9g8J6zB4f8GSXt3uMnnvsTpk4XA/GkOEz1ZyqoWcgADkk4A/GsuTWtFtuJ72BPYtn+VfPet+JNT12dmuZXWE52xRnA+jeornDtAPQUWF7Q+rIdV0q5I+z3MT56bXx+hrQBBwQeK+QMDg/jkdfwrufDfjTUtFukiupPPtjgMH5Kr7UWBTPogdakqtbzx3MMd1BzHKu4YqdxlGHPQ9OvSoNyPz4B/wAtE/76H+NOWWJztR1Y+gINfJt/JL9rnVC+A7d/eu3+GjNJ4jbeXyID1NVyke0Pf6OQNwz9QM/lS5I5XrWRrer2miac2oXTHbHlUQdS3tUlkt1pumXSmW8t4pP9qVF/mapLqXh3TB5UU1tb46qCF/lxXz5rfiXVtelb7XI4hz8kKH5QPVh3rniFA4xV2M3UPqk3uh6uBE88F0DwFJRgPz5q5bafpllzZwxRE/3UA/I18kLjcGGAR39K67Q/GOtaK21JTPCD80cnzAj0BPSk0Sqmp9LUA81Q0vUY9XsItShUos4yykYKn6VyHizxtBoG7TrMB70jnPIjH96oRs3bU72WWKJC0zBQOu47f1rMk8QaFAdst7Ah9C2a+ZdR1jUtWkaS9uGlDfwj7orHwo9BWnszP6wfXMGqaddMPs9zFJnptf8ApmtDPc9K+OYzscPGcMOQR1Fd94c8fappLrDqL/abcEAh+So9qHTGq93Y+ie+KKr2t3De28d5AdySjKt3J9DViszcKKKKACiiigAHWn0wdafSYC0UUUgCiiigAooooAKKKKAFHWnU0dadSYBRRRSAKKKKACiiigAooooAUdafTB1p9JjQUUUUjQKKKKAFooooEFFFFABRRTWZVUs5wB1JoByUdWPAwadn0rj9Q8XWNmTFZ5nYHnPauQvPFOq3GfKkWNSeiDDfnXTSwdSoefiM0oU91dnr+R0pM14U+pahJ8zzSZ92xTVv70n/AF0n/AXOa6v7Kqdzk/t2HWB7x70fSvG7fxHrNqRsmLAfwvzmunsvG0TlRqMRhPTK9DXNVwNSC5mb0c1oTaO9oqG3uIbqITQNuQ96mrkPVhJzV47BTl60lKvWkIdRRRU3GFKKKKACiiigQUUUUAFFFKOtA0A606iikWFFFH0oAKKDx1oyPWgAoooPHXigBy9adTV9adSYBRRRSAKKKKAClFJRkUALRRRSAKKKKAHLTqatOoAKKKSgBaKKTr0oAWiiikAUUUlA0OWn0xSKfSZQUUmaPpQAtFJ0paQwooopCCnCm04UALRRRQAUUUUAFFFFABRRRQAUUUUAcJRRRXtHmhRRRQAUUUUAFFFFABRRRQBS1P8A5B0//XI/yr5XHQV9Ual/yDZ/+uMlfK46ChGdQ9T+GX+vvf8AdFexV478Mv8AX3v+6K9ioCB5X4u8EmQyatpI3SHmWPu3unpWX4b8BG+hF7rLmKM8xRD77D/aNelax4i0fRB/psxD/wB1eT+VYH/Cx9DPRJyf90UAb8HhbQbePEdkh7fPziuX8Q+BdOntZLjSh9nuIs/IP9W4HtXTaF4l0/xDHIbGN1WI4fnvXQYzwe9I0Pk7Pl/u37V6n8NdXkE0ujyH5D88Ps3euB16KODWrtI+0pq74Qmkg8S2ZQ8Gbb/30Kk50fR9fKl1/wAfk3/XVv8A0I19Xdq+Ubr/AI+5v+urf+hGrRpUOy+HP/IyR/8AXGX/ANBFe+Hoa8D+HP8AyMkf/XGX/wBBFe+HoaT3CmYfiv8A5F3UP+uLf0r5jPSvpzxV/wAi7qH/AFxb+lfMZ6VKCoez/DH/AI977/fX+VeoHpXl/wAMf+Pe+/31/lXqB6VTNKZ4X8Sv+RgT/rkK4S0/4+Yv+ui/zFd38Sv+RgT/AK5CuEtP+PmL/rov8xUmEtz6wGMLn0X+VczN4y8OQu8U10FkjbH3TXTLn5CvX5cflXhHxA0cafrC30I/c3AyP98dao3qHrml+JNG1S5Nvp90JJD8xQj09Ks6romna1DFDqK71iYlQPU183aJqD6TqNvfQ8eW/wA3+6x5/SvqGCRZUFxD/q2RXX/gVQFOoeKeL/BNro1odU0suIs4eN+cGvNMAnBr6f8AE0Ybw7exv18gt+XNfMAGQB60zOofRngrV5NW0WNrlyZoP3Unuw+7+lHjr/kVbv8A4B/6FXI/C+4kMl/bE9Qkn9K67x1/yKt3/wAA/wDQqDSP8M+de9e5fDD/AJAlx/13P/oK14b3r3L4X/8AIEuP+u5/9BWgxpnpJ715T8Uv9RYf78n8q9WPevKfil/qLD/ff+VBrUPGR1r6e8Lf8i9p2P8Angn8jXzCOtfT3hf/AJF3Tv8Argn8jQyaY3xJ4dtPENibWVnSRcsjA/KG7E+1eGJ4T1Ztc/sZo/3nHzAYTb/eP0r6T478Vkalr2kaSok1OdYz0UY3OfwpFzgYOleAtE06EG6jFxMfvGXlR9BWnc+FfDlzEUe0jAx95BjHvWIfiV4cVyq+a+O/l4/rTD8TfD+Puzj32ijULwPKvFfh9/Dmp/Y1+eOUb0P+zWBaTtbXMVwvWN1YZ9jXZeNvENj4gnt5NPWbECFSWAA5rhD90/SmYS3PrWARTQxTAIXeNWJ9eKbd6daX9obW5ijMJ7Y5zUGif8gWx/694v8A0Fa1Kk6meKj4a3Q1pLcSE2OC5k/iA9BXo1r4Q8P2MARbVHYDrIMlvrXSUUyOQ8c8eeEtP06xXWNNXySG2Op5XDenpXkXPavov4hf8ipcf76f+h187L94U0Z1Nz6H+HN0914bijbpbyPF+A5ruz0P0P8AKvNPhf8A8gF/+u7/APstelnofof5VL3NYHyPdf8AH3N/11f+ddz8M/8AkYx/17tXDXX/AB9zf9dX/nXc/DP/AJGMf9e7VTMIfGe/npXjvxXu5N9lZD7p3yn8Plr2QdeuPevK/idpEl1Y2+q243C3BV/XFZ9ToqniKKryKjZAJAOOvJr6Y07wroGn2yQi3R2IX5nAY8jt618y4554967jSviDr+mwrDviuI1+ULIOg+tWznp+Z6/qXgrQdRQ7rQQtj5XQ7Tn1wK43Qvhu1vqTT6owlgt2xEq/8tB/eb6VHa/FUhgL2xU+rRNyPoD1rvNF8XaFrY22s+2U8eXKdrZ9B2pam94dDpjsVVWMeUqjaMdq8I8ceDLmznm1rTBJNBI2ZFY5Zfc+or3jp1rB1/XrPw/apPeoXWU7cL1qCntqeVeFfh+2qxjUdWYwwt9xF6k16jF4U8PwxCCKxjOOCXGS1c3D8TvDpcK0c0Yz3UYrvrS7tr+3S8tG3xydDQTThA868R/DzTLi2kudGT7PKilvLPIYjsK8KKsrPHISrZJI9xX2ICAQW6f418oa7GINavYouNs0i/nVxZNWFj0j4Xau4mn0Od8gr5sP0H3x+Jr2ivmPwTM9v4osWj4DHy2+mDX00F+VW9ame5dHYdRRRSNQoopR1oAB1p9FFSAUUUUAFFFFABRRRQAUUUo60gAdadRRSAKKKKACiiigAooooAKKKUdaAAdafRRSGgooopGgtFFFAgooooAKKKjlljhjaWX7qjJot2FKSiuZ9CC9vbewgM9ycKOnue1eWav4ivNUkZV+SIdF9feoNb1eXVrtm/5YqcJWMemK9vCYNQXPPc+UzHMXVlyR2GjPIPU0A+lGSORxWlY6Vf6iR9kiwP72MV2zmoq7Z5kISm+WK1M3NKOTiu2h8D3hH7+5Rc9gOalfwPcBf3Nypb0IrnWOpdZHasuxFvhOEIpVYqeK277wzqdi25kJUfxLyPyrDYMGw3UV1UqlOorx2OOrQlT0maNhqd5pcgktH3KOsZ6GvWNH1u21iEPCcP8AxJ/dPsK8VqzaXU9ncLc2zbHTv7dxXJisHGprDc7cBmE6Ekuh70CB/FgetZl7qlranbHmR/QGucPiNdUthLYjEZ4J77u9Z3POepr81zfO5U5vD0NGtLn2tO9SKnHY1ZdavpG2oQi/rVRry9Y5MrH6GqtHtXy1bHV6j/eTNFC2pZF7fIdwkcfU1qW+u3MeBMoZe5rB4FKKqlmOIpfBMpq+h6DbXdveL5kTYPdatZrzZXkiYSRHBH6+1dlpmqJeqI3GyQDp6ivrsrzlV7U6u5EoaXNiiiivoDAKUdaSlHWgB1HvR2NefeLvHth4cBtLRRc3bD7ucqPr/hQo30ByS1Z3F3d2tlCZ7uVIUUZ3OcD/AOvXnmqfFLQLNgliHvnxyYxtQfia8F1jXdU125+0alP5vcIPuofQCsrNdMMOkrs5amIf2D125+L2osxFrZQxD1Zix+mKpL8W/Ee//U2xHoRivMDkdaTNX7Km9DL29Q9otvjBKCPttgrDuYm5/DNd5o/j/wAN6sUiWf7PI/Hlzdc/WvlvNKMEEHAz37j6e9J0IW0HHEzTuz7cBDENkEYwMdDTz6V8peHvG+seHisQkNzaj70T/fx/snsa+jfD/iHTPEVmtzpzluPmRvvofQ/41yzp2O2nV59TfooorM1CjOBRSEgc98gD8aYXtqcF4w8dWnhlhZwILi5PJVjgKMdTXnlp8XtWjk331rCYQedhO4D1riPGc8tx4q1F5eSJdg+ijFcwV3CuqFBNWOCriHzH2bpOq2es2Ud/YsDG46dwfetIV418Hp3ezv4W6LKhH/fNezVzVI8rsddGXNG4UUUVBoOWlNNHHWnFgvuRQCGOSoyOo9en415H4o+KMOmXDWOixpcyJw0rHCBvTisf4g+Pi/maJocmB0mnB6ewrxVgABtOR3PqfWumlRuuZnJVxLT5YntuifFyZrhItct40QnDTRE4XPqK9tt5Yp0WeFgySKGUjoQe9fEgbtX1d8P3nk8KWLXP3/LOP93PFTVhbUWGqzk2mdrRRRXM1dWO1K7Ck4HJH59Ky9Y1aw0aye81CQRoo4/vE+gr508T/EbVtclaDTy1rajgAfeb3NaQpXMqldQ0Z7jrHjbw5oxeK7ulaRR/q4xk/ga87uvjEoOLCwOwdDK3P6V4fjuxYsepJoJ7V1ww6tqcU8VJu0T1WT4veImbMdvbBfTBY1Yg+MGsIf8ASLSGQHqAWQj+ea8hPHWgEdqv2FMn6xUWp9Fab8WdCuW26hHJZnj5jh0J/DmvTNO1LT9TjFxp8yTIQOUOf07V8Tsc4ArS03VtT0m6W4sZzAy9x0PsfUVjPDRb902hi59T7XyD0pRXkvg34kQayy6frKrb3TcBgcI/+Br1lTx6jtXLOLi7M7KcubUWnCm04VmaC0UUUAFFFFABRRRQAUUUUAFFFFAHCUUUV7R5oUUUUAFFFFABRRRQAUUUUAUtS/5Bs/8A1xevlcdBX1RqQJ0+cD/nkf5V8rDoB7UIzqHqnwy/197/ALor2Mda8Y+G11b2s979oZVyoxuIGfzr2O3uIbkb4GDDPVSCP0oYQPnHxVPJceIbt5PvBtiH0ArIs7C81KQxWELyuOTtOK2vF1rJaeIbxJP4z5g/GpfCOrx6LrSTXH+ql+ST/GoMz0L4e6XqGmxXa6lbyW/mOp/eADt7V6QM1Wt57e6QS2sokT0yDWdr2vWGiWMk9xJgtkKoOTntx9au50Hz/r2Dr93j/nsan8KRGTxHYJjpMp/LNYkskk05nk/1jOXP416F8N9NefUpNVb7lrwPqeKg50e3V8o3X/H3N/11b/0I19XgcgV8tatB9m1a4gb+GaRf++jmrRpUOo+HP/IyR/8AXGX/ANBr3wg4NfMvhzVBo+sQajL/AKpSQ4+vevoyHVLC5hSe1nRkYZwWHH61AUyj4r/5FzUD/wBMW/mK+Yz0r2zxt4osV06bSLB/Omn2o4ByAOvavEyDyKEFQ9n+GP8Ax733++v8q9QPSvJvhvd21vDeLO6qWdMbmA7e9esnp+FUzSB4X8Sv+RgT/rkK4S0/4+Yv+ui/zFd38S+PECE/88hXCWpAuIieAHX+YqTCe59YrkbCOvy/yrlfGekjVNDmjQfvIMyR/hy1b8V5ZSMiRTq7HbgBgT09AavkAjB6Hig3Z8jFgBuPYZr3nwBq5vNGawlP76z+Q/7jcivK/FukDSNamiUfu5T5sf8A1zPUfnT/AAhq/wDZOtxSyMFim/dyk9B2B/Cg5z3TxIR/Yd6fSI5/75NfLw6D6V7R468Wae+nSaVp04mecjeyDjivGaEaVD1n4WRE3N/OeyIv/jxrsvHX/Iq3f/AP/QqpfD/TpLHQ/tMud923mj6Dj+VbHi+1N34avo0/hjV8f7pzSNF/DPmr+Kvcvhf/AMgS4/67n/0Fa8MJ9Dg/yr034eeIbTT559NvH8tJSpibplu+frTMaZ7ee9eVfFHmCwxz88n8q9LbUdPjUyvPGq9zuU4H514V478Rwa5eRW9jloLTI3rxkmg1qHBDrX074YOPDmnn0gT+Rr5iHXB4r6U8JXdo+iafAsylxCvy7hngHt1oZFM1Nd1P+xNKub8gO8YxGPViOK+Yry+udQuHurtvMlkbJY/wj0r3j4iCT/hGJNvTzVzXz3z2pIJmzpWhanrchTTo8gdWPA/OumX4a+I3U5WP/v7/APWrv/hqsQ8PAp97zn3flXoQIJwKLjhA+Xda8N3/AIeaBNQ2Dz1ZgFYnkVgn7p+letfFT/X2P/XOT+a15Mfun6UzOW59V6J/yBbH/r3i/wDQVrUrL0T/AJAtj/17xf8AoK1qUjrhsFFFJSEcX8Qv+RUuP99P/Q6+d1619EfEL/kVLj/fT/0OvndetNGFbc92+F//ACAH/wCu7/8Astelnofof5V5p8L/APkAP/13f/2WvSm+6fof5UnuawPki6/4/Jv+ur/zrufhn/yMf/bu1cNe/wDH/OO+9v512vw7mhg8QPJO6oogPLED+dUzCPxn0JTZI0ljaOVN6sMMp7j0qK3uLW5Ba3lVwOuCCPxxWTr3iCy8O2Ju7ti0hzsQHkmszpucFrfwyhmLXGizeUzEsY5eVJ9BXGXHgDxTDkJbCYf7BFdv4d+I8E0zW2vbYmZsxSAfJg9AR6+9enQ3NtdRCWCZJg/TYRVXZKhCep8p32l6ppp230MkPb5l4/Oqcbujq6MUZTkMOoI719M+ML/TIdCu4b942VwRGnVt1fMiZ+X6H8qad0Y1IH0x4O1qTXdCiublj5kZ8p27s3Y1jfEiyu77SIorWNpj5o4XrUXwvhkj0OV36NL8tel8444NZmyXuJHyBNBcW7mKdGidf4W65r1n4Wau+ZtElPBHnRfQff8AzNUPifpfk39tqsIIW5UpJ/10TpXF+GtR/sjWra+X7sZCOPY8Vo9UZp8jsfVXp9RXyl4jK/2/fE8j7Q9fRur+ItK0jTpL15ldlX5UVgckjjoa+Xp2e4mlnk5YuWYj+8/NTTHUOi8GxtceKLADoJf6Gvp7GGx6CvDPhfpbzanNq0n+rgj8tf8Aro/Ne6Up7mlHYKKKKRqFKOtIOtPpMBaKKKQBRRRQAUUUUAFFFFABSjrSUo60gHUUUUgCiiigAooooAKKKKAClHWkpR1oBbj6KKKk1CiiigBaKKKBBRRRQDdtRDjHPfgfWuJ8Y6o0UaaXCfnfBf8ACu4HWvE9avDfapcTH+E7PyrswFP2lS/Y8rNq/JSa7mX6gdFoxnqcUDHetrQdMOp30cZ+6p3N9BXt1J8kXPsfKUqcqk1CG7Nvw74c+2KL3URtT/lmnrXpESLEgiVBGB0A705I44/lj+6BhafXzmIxDqSuz7bC4SOHppR+J7i0UlFYtXOrcD0rite8MwXatd2n7uUD7o6Ed67ajAJx69PrV0qzpSTRhWoU68XCW58+mJ438qT7y8GqF7cFP3Yr0LxnpwtpP7Vi4EnEn1HSvJpWLuZD/FXZmePvRT7nyk8Hy1XF9DoPDeoC1uzA/wDqpuv+92r0DbtLK3UdPpXjqsU+deq817LpUUupLFg/eQE/QV+bZthZVKkeXdn0mU1rRcB0NvPcNst0Ln244+ta0Xh+9f52kWP2Iya6iC3gs0EUI7c1YHNenhuHaXInW3PVdTQ4uXRb5AeElA9BzWSVdMrINpHavSSxWsvUtNS9jMkf3xzWGP4fjGHNR6ExnqcTT4p5LeRZo/4Tk/SowpBYN1BxR0I+tfKJckmnujU9EtLhLuBZ079frVjPNcr4fnJke39s11S9K/R8txSxFCNZ77GE9xaUAEHPpSVmaxqcGjadNqVx92FSQP7x7CvRtci9tTjPHnjNPD0B0+wbddyjk/3FPVvrXzc8sk8jTytueQ5Y/wB4+pq3f6hPql3Lf3Tl5Z2LHP8ACOwqpx3rspQ5dTgr1LsSkOTwP/r/AIVsaLol/r98tjp6ZY8s5+7GP7x/wr6J8N+AtG0WMPKiXVyB80kg4z/sjtSqVbahSoc+p89WHhrX9TTzLCymdD/EeP54rUbwB4sVP+PCVvYFf8a+rcYUA4GOw6UmBWP1hnT9UPjS90zU9NcxajbSwOOzL8v/AH0OM1QJNfalxDDcReVcxiVD1DDK/UivIfFPwytpUe/8PgxS/e8jqj/7pPT6VUa93ZmFXDcqueDnpj1rX0TXNQ0C7S9s3wV6r/fH901myxy28r28yeVIhw6N1Df4VHjcQvc9DW9rnOtGfYPh3XbTxDp6ajanGRiRP7rVv18tfD/xK+iazHBLITa3R2yKegbsa+o1bIOTk5z/AIVwVadpHpUJ3VhaO4+ooo7j6ioZufIPiv8A5GO//wCu7Vz3cfUV0Piv/kY7/wD67tXPdx9RXpx+FHkS+Jnufwa/49tQ/wB6P+Rr22vEvg1/x7ah/vR/yNe21w1viZ6OH+AKUUlNbJ4HXt6fjWJshz/d6ZHpXh/xB8fbPM0HQ5P3mNs86/wnuBU3xA+IAjDaHoj7mORNOvRfYe9eFHIGT3OSfU+tdVGjf3mcWIr7wG8ZyDkEZB9fejPzAdj1oIyME4B6+9dd4R8J3fii8CY2WiY81/b0HvXS3bU44q7HeDvCd14ovV3r5VojAySDqQD0H16V9V2dtDZwLawKEjjACIP4VHaoNM0600q0SysYxHDGABjv71o159Sd2enSo8iuJVHU9RtdKspL68fZHGMk/wBKuk4Iz34/OvnX4o+Jf7R1D+wYHxb2v+sI53P6fhRThzSUSqsuWLZxvijxNfeJ743VydsKMRFF6D+9+NcyTSZzg9cjOfajhvlPIJxgdSewFelGPKuU8uT5ncUGrlpZX184isoJJmbgBFJz+PQV6x4R+GH2uNNR8Qq0anBS3U4Yj1Y17tZWNrYRiCzhSKMDACDHT1rlqV+XRHTTwvMrnyovgDxe67lsJVz/AAsynH61RvPCHifT033NhIEHcc/yzX2NgUhGCKyWIZr9TTPhohlO1sAjquMEfWkr6z8SeCtE8QRMs8QimOT50Yw2f9r1FfN/iTwzqXhu8+z3qEo3Eco+64Hp711U6/MuUwqUuQ5snHQ7W6qfQ19BfDjxs2oKmg6s+LgD92+fvAdq+fqfDPPa3EdzA2xo2yG7g9qKkOePKTSnyyTPuZQATjvz+NOrk/Buvr4i0WK+biQfJIP9odT+NdWK8yS5XY9SMuZXFooopFBRRRQAUUUUAFFFFABRRRQBwlFFFe0eaFFFFABRRRQAUUUUAFFFFACMAylT0IIP4183+JNAu9DvpIHjzC5ykgH8HYV9I02WOOdPJuVRo/RxmgJwPk1h/eH517r8Nsf8I6+3H+vl/wDZa6F/C3htuZNPi/Jv8a07Cxs9NgW3sohFGXMmB70EU4HL+LvCy+ILfz7f/j8h/wBUT/y09VPsK8JurS7sJzBdR+SycZcf1r6sqpeWNpqCeXfQxy/UUrhOB8tRXE0P/HrI6/8AXNiKZLLPKfMuGcn/AGzmvoG48B+Gpz/qMf7jkVPbeC/DNp0tFP8AvsWqRezPEdD8P6nr8/2eyXA7uRgAdzmvofSNLtdGsFsbRchON3d27savwxw26eTCEWMdAgwKfVXNKdMOvGcV478QPDkwm/ty2i/dv8k6qMkY6PgetexUYEgIPQ8GgD5HHPTrT1yp27tpPbHWvo678GeHr1zJNAE/658U2w8GeG7F98dorn1kYk1Jn7M888CeF5Lu5TV7+LFvEWCIRjzCfUe3auN1zTptI1Se0nzlJDhj0KPyOa+nRlV+RdiLwAKq3lhZ3yeXeQRTD/poM0D9mfK0akyJlf4h296+soseRETyPLT+VYI8J+GlIK6fCCDwQOn61vqFTYij5EGBQXTpnnPxC0C61C3TU7NPMkhH7xf9jtj6V4gwKHa42n0PH86+u+e1ZNzomkXbF7m0ilcjkkUETgeG+AV/4qm2wP4G/lX0Rzj5evasOz8OaDYzrdWtpHFMBhSA3A/Otw88CgunTPOviJo63+kf2lCPntG6d9h6/rXhHNfWzwR3EbQ3ADKylSD0IIwR+NYH/CIeFh006H8j/jQROmfNHIO0dT0HevQPCXgu71Wdbu+BitUIJBGCxB6YNezW2i6PaDFtaxRfRf8AHJrUQKPlXt60B7MSOOKFViT5EQYAFEiLLG0TjKuCpHqDxT6D0/woNmfM/ifQLnQNQktpE/clt8UmPvA9ifaubxnr0r6zvLK2vIfst7EskbDo/wA1crP8P/C07b/I8sjkKjlAfYjmi5h7M+elVpisQJbJwoUZOa9atfBdzB4PvDOgW8uUSTnkoqHIA9zXoml+HtH0f5rC2VHPU43H8zW0OvBx70DhA+RZVaN2ST5ZE4IPB/Wu0+Hg/wCKohOP4Jf/AEGvcrvQ9GvmL3llDI3diOWptn4f0Swn+1WtlFFNjAZQeB+dK4KmO1zSl1nS7jTpDhpl3IfRgOP1r5mv7G70y4e3uosSIduOgYeoPrX1h9eKoX2m2OpoY9Qt47hMcZHP4UF1KZ836L4k1XQHYaY4Ebcsj8gGvd/COvS+INLN/cBfMExU7BgdKrP4D8MynItdnsrHNb2laTp+jW5ttOjaOMnkMcnNDHCmeV/FQj7RYj1jk/mteTn7pz6V9WX+jaXqZibU4Fn8pSq7uetZp8H+FyCBp8I+g/8Ar0XM5U9TR0Mg6LY45/0eP/0Fa1KZHHHDGkcCCONBtAFPpGwUUoBzTmBYFQcZ70AcP8QuPCk4P/PRP/Q6+dlzmvra9srTU7d7e9jE0DsG2t3IrEHg7wvnA06D8v8A69NMipT1Oc+FwP8AYDnt57/+y16WM5+Xk1nWOm6fpkXkadbxwxZzhQev51oEEjAqXqaw2Pnbxv4budK1KW+t032k53LJjlfUH0FcIMH0Ir7AMSTKVmAYdweh+uax7nw14eujun0+Fz64x/6DincyqUzhPhQP9C1EqB/rYe3H3T1qz468HXmrzHWNLYyy7cPA54x6r6Gu9sNH0zSDIml28cIlIdtoPb8a0uvFQu44U9D5AmimtpXhuF8uT7pVxz+GadFcXcI/ctKvujYFfVl9pGl6kuy9t4pOxZx834Guan+H3haZt3kFf9xytacxn7M+c5pZJAXncsR3ZiTXQ6B4Y1HxDdeVbqY4FwXkxjA9q92tfBHhi1IdbVHK8jzMv/8AWrqYYY7eMRRqqIOioMCk5AqZW02wg0qwi060UKkICk+pq9kDk9BRRUG5ynjLSv7Y0CeDq8eZo/qK+ZMHA7Ng5z6g19isodShAYHjB6Gudbwf4XYl206FgxyThs59etCdjOdPU+Xy74B6Be9b2h+GtV12ZYbWF1jzzKeFA7n3r6Ii8MeG4CHi0+EbeeVJH6mttI0jQIiqijoEGBVupclUbGZo2lWeh2Uem2K5VFyWP8R7t9a1qSioNhaKKKCwHWn00dadSYC0UUUgCiiigAooooAKKKKAFHWnU0dadSYBRRRSAKKKKACiiigAooooAKUdaB1p9DGgoooqTQKKKKAFooooEFFFFAEM7bIXf0Un9K8I3bi0n97/ABr2/Uf+PC4/65t/KvEB91vwr18q3Z85nv2QjBLACvRfBEIMM92fvOwX8q84HevUfBf/ACDX/wCux/lWuO/gfM5cph/tJ2NFFFeEfWhRRRQAUGig0Cexy/jMRnw5deb0wMfXPFfP4bJIbru4/KvfPG//ACLVz/wH+deB9G59f6V52LWp4eZ/xfkIn3xXuPw/+fRRNJy+4x59lrw1eor3T4d/8i6P+u71GHhzyT7Dy/8AjfI7qiiivSPdCjPIPvRSHpQCOI1e3FvfMR/y05rMNbOv/wDH6n+7WLX5lmUbYmZ0Q2NPSZPL1GI+pxXdHgAfWuA03/kIwf74r0Dt+Jr6jhr+A/Uie4mQOvpXinxZ1SaOOz0SM8n95L/Ifzr2ojP6frxXy98Q7x7zxbeZPywlYh/wEV9TR+JHLX+FnDr7dBxU0MTzTJDEu93YBV9Se1R13nw402LUvE0fnDKW6GY/UEAfzrsm7RbOClHmmke3+DfDcXhzS0hZR9qlG+dh2J6CuwoAOBnqM7v6UV50pXep6sFyLlCiiikUFLxzuOB3pKcvWk9hXPFPin4ViMH/AAkNkmJIsCdQPvKf4vwrwz5SzcnjGB7EV9rX1rHe2c1nL92ZCn/fQxXxjd2sljdTWUv34JGjb/gJ4/SuzD1Lqxw16fK79yEHGOcen1r6u8DaqdX8OWty7bpY18qT/eTj+VfKB+7gd8Cvevg9dl7O/sT/AMs5VkH/AG0H/wBaiutCML8R7LR3H1FFHcfUVxM9I+QfFf8AyMd//wBd2rnu4+orofFf/Ix3/wD13aue7j6ivTj8KPIl8TPc/g1/x7ah/vR/yNe29q8S+DX/AB7ah/vR/wAjXtZzwV6joPWuGt8TPSwyvEDjoe/b1rx3x94+S136Hoz5nIIllBzsHoPxqT4h+Of7ODaHpL/6RJkSyD/lmPQfjXgB3t88nLE8n1PrWlGlf3mc+Iq290Qu7ktLyxPLf3j60hGRSNjjPNdT4X8K33iW/WO3+WFTmWQjhV9Pqa6b21OZJvYf4T8K33iW9WGP5YF5lkI4VfQe5/SvqfSdLsdHs00/T02RxDHuT6n1PvUej6VZaNZpp9hHsjjH4k+p9T71r1w1anMz0KNBU1d7hRRRWJte2pja9qKaRpNzqT/8sYyy/wC92r45nmkuJWuJfvuSz/7zHNfSHxYvPs3hkQd55lX8ua+aiME/U12YdaHFjJ3khrnjA6kjFet/DDwtFqt22u36AwW7bIl7M/qfpXk2M5x1wcV6BpHxF1jRNPh06whg8uNe6nOfwNa1IuUbIxpSUZJs+pEI5BPI6/Wn8V81/wDC2/En/PGD/vg/40f8Lb8Sf88YP++D/jXM8LM6/rMD6U4o4r5r/wCFt+JP+eMH/fB/xo/4W34k/wCeMH/fB/xqfq0xrEwPpTIrB8Q6Ha+INPexuFBJVtj90bHBFeE/8Lb8Sf8APGD/AL4P+NJ/wtzxJ/zxh/BSP6044eondBPEU5KzPOdQs7jTb2XTroYlhYq3v7/jVPJx8oycg59MVt69rlx4gvjqF3GqSEYJXvisUV3xtbXc4KjjvE9Y+Eer/ZtYm0pj+6uwXUHsy8n8xX0emOor438K3Js/EmnzjtcIv/fZ2f8As1fZCcDFefiFqd+FneNh9FFFcx1BRRRQAUUUUAFFFFABRRRQBwlFFFe0eaFFFFABRRRQAUU+ikAyin0UAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFBYUUUtABUlFFSAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFOpaKQBRRRSAKKKKACiiigBKUdRRSgHNAElFFFSAUUUUAFFFFABRRRQAUUUUAFFFKOtAAOtPoopGgUUUUAFFFFABRRRQAUUUUAFKOtJSjrQA+iiipAKKKKACiiigAooooAKKKUdaQAOtOoopAFFFFABRRRQAUUUUAFFFFACjrT6YOtPpMaCiiikaBRRRQAtFFFAgooooGijqX/IPn/65tXiHc/QV7lf5NjPj/nm38q8MAO4jHNezlP2j5nPN0IO9eo+C/wDkGv8A9dT/ACrzBRzXp/gr/kFn/fNXj/4HzOfJv94+R2NFFFeEfWhRRRQAUGiigNOpzPjC3e58O3SR9VUP+Rr5/JLRlm6lq+pWAZSrLuDDGPXNeb6p8PIbm4efTJhGXOWRxlVPtXJXo87ujzMfhpSalTPIVIUjmvdvAFrJaeHIfN+9IzMPpWFpnw5EE6yancrMFO7Yq4yR0B9q9NjVIgIUXaAOFHRQOwqcNRlC7Y8DhJQbnMlpRSUorrPTCkbpS0UAtzjdf/4/E/3axK3fEP8Ax9p/u1hZFfmuZf71M6IbF/Tf+Qjb/wC+K74/1NcFpn/IRg/3q709fxNfScMfwZepnMUfeH1FfI/iz/kZL/8A67tX1uSQMivljx9afZPFl8nZysn/AH2K+qw+7OPF7I4017B8HxH/AGnfuPvGJRXj+QCGPY16P8K75LPxQsMpwtxG6DPqcEfyrep8LOWh8aPpWigDHXqQP0oriPVCiiigApwptOFAmKe31H86+Q/FqhPEuoAf892r63uZltoHuZPuxKXP4DNfGOoXb3+oXF8//LeRpB9Ca2w+7OTGbIqZwcjsD/KvYfg6zjVL+Mngwq345FePd/wP8q9g+Dv/ACFL/wD64r/MVtV+FnPh/wCIj3+juPqKKO4+orhZ6h8g+K/+Rjv/APru1c93H1FdD4r/AORjv/8Aru1c93H1FenH4UeRL4me5/Br/j21D/ej/ka9sOccdf8ACvE/g1/x7ah/vR/yNe2e+Mkdq4a3xHpYZXjY+MtZlll1e8ln+80zk/nWcODXsvxF8ESJLJ4h0xMo/wA08eOn+0K8YB4B67hkN7eldtBpo4KlOzZLAsTzxrOxWMsAxHUA9T+FfXvhix0iw0mGPRsGAqDuH8ZPVj718fcAjNd74K8aXPhm5W2nYyWchwy/3eeorHEK6Lw9S0lE+p6KrWl3b3sKXFq4kjkG4EHP4VYriPRFprfdNOpKBrc8b+Mn/INsf+urf+gV4Ce30r6N+Ldq03h2K5X/AJd5g3/fQxXzguNoHcZH613UPgPMxPxC0dOtIRu+X8fyrqbLwX4l1C2S9s7Nnil+YMGBJ/WtW7amEY3djl6M12X/AAr/AMXf8+Mn5rR/wr/xd/z4yfmtL2qL9kzjc0Zrsv8AhX/i7/nxk/NaP+Ff+Lv+fGT81o9sg9kzjc0Zrsv+Ff8Ai7/nxk/NaP8AhX/i7/nxk/NaPbIfsmcbmkzmuz/4V74ubrZSfmtKvw98Xq2Usn3dssuPxoVZdwVJmHoORrmnk9PtUP8A6GK+0V6mvEfA/wANbjTrxNW13aZY/mSIHI3difp2r2yPOSSOvX61w158zO/CQ5YslooornOkKKKKACiiigAooooAKKKKAOHooor2TzQooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooLFooqSpAjqSiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKfTKfSYBRRRSAKKKKACiiigAoooHUUAOHWn0UVIBRRRQAUUUUAFFFFABRRRQAUUUUAFKOtA60+g0CiiikAUUUUAFFFFABRRRQAUUUUAA60+mjrTqTAWiiikAUUUUAFFFFABRRRQAUo60lKOtIB1FFFIAooooAKKKKACiiigApR1pKUdaAH0UUVI0FFFFBoLRRRQIKKKKACg0UUAxGAZSrdCMH6GvLtZ8M31pO81mDNE5yMfeFepUY71pQrOjLmicuKwlPERtPc8XtNB1O/k8sRPEp4JbivW9NsY9NtI7NOqjmtGkz2rSvjJ1tJGWCwMMO3bqFFFFcx3hRRRQAUUUUAOHWnU0dadSYwooopAFKKSlFABSEZBHrS0hGaQJ2dzM1bTxfRq8IG+MY59K5P7HeLlRFuPTpXf0uTXjY3JaWInzstTuzntI0l4D9oufvEYxXSck57UylHWu7CYOGHjyQJnuKRnvgjpXhXxc0t0ms9XjHyMpic/7S9K92rB8S6NFr2jzadJ1YZQ+jDpXbB2kjKr8DPkLpU9rcy2l1FdQHEkbhl+optxbz2dxJa3ORJESrZ46dKi9PXtXba55/kfX/AId1238QaVFqEHWQYYdwR1zW5XyV4X8T33hm9Elr+9ibiSI8bvofWvpTQfEuleIoPPsJQX/ijJwyH0x1P1rkqU7M76NXmVux0NFBODg8GjIJwDzWR0CggdTT88gdj3qPOD2989hXnni34h6docTWtg63F30AHKr7nH8qEruxMp8quZvxP8TLYaYNGtH/ANJuhk+yd6+eiAuR64I+nf8AWrF7e3OoXT3l25kklO4lucew9B7VWrthTsjy6s+Z3Gn/AB/lXsnwd/5CV9/1xX+Yrxs/4/yr2T4O/wDISvv+uK/zFKr8LKofHE99o7j6iijuPqK4WeofIPiv/kY7/wD67tXPdx9RXQ+K/wDkY7//AK7tXPdx9RXpx+FHkS+Jnufwa/49tQ/3o/5Gvbe1eJfBr/j21D/ej/ka9trhrfEenhVeNhkgDKQwyDwc8j8R6GvAPHngJtPaTWtFT/R2JaaHrsPqvtX0GKbJGkqGOQblPBFTCfLJMdWldWPh/gANnO4ZDf0pc4616t4+8Btpcsur6LHutHJM0eM+W3+yPSvKBkKM8g8gn+VehCfMuY82pHl0PQvA3jefw9MtndkvZOfm/wBjPcV9NW1zBdxrc2zh45Blcc5/+vXxJyDXongXxvP4cmWyvmMlm5wf+meT1H1rmrUrpzOjD1be6fT9FRQXEF1Ek9uwZHG4Ecgg1LXG9Dtv1MPxLpo1fQ7qw7yRnH1HNfHJV0/dy/eTKn8D0r7iJP8An1r5i+JPh3+x9ZN9ACLa8+cegf8Ai5966MJPdHNjKd1zHnIPPPTufQV7f8KPE8aBvDl2+N2Xt2PGR3WvEMVLDM8EizROY5FOUYdQRXVOHOuU4aU+WSkfb6sMD0xmnj1FeQeDPiTZ6jHHp2uyLb3X3Qzfdf6noDXrcciyIHBUg9CpyCK8+cHF2PWpy5tSWikyBQTj8azavoagcDrTS6Rgs5AA6k1k61remaHam71OdYUHY8lvYDqa+bPF3jq/8R3Hk2rvb2a52xrwxz/Ex7g9h2rWnSuc9WvyaH09FqFhcyeVBNG7DsrqT+QNXlAFfDVrJLazLJas8bqwKuhww+tfUXw38UT+I9JcXvFxbvsb3XsfxqqlFxjcVLEc75T0aiiiuc6AooooAKKKKACiiigAooooAKKKKAOHooor2TzQooooAKKKKACiiigAooooAKKKKACiiigAooooLCiiigAoopaACipKKkAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiin0AFFFFSAUUUUAFFFFABRRRQAUo60AHNSUAFFFFSAUUUUAFFFFABRRRQAUUUUAFFFKOtAwHWn0UUiwooooAKKKKACiiigAooooAKUdaSlHWgB9FFFSAUUUUAFFFFABRRRQAUUUUAKOtOpo606kwCiiikAUUUUAFFFFABRRRQAUo60DrT6GNBRRRUmgUUUUALRRRQIKKKKACiiigBR1p1NHWnUmAUUUUAFFFFABRRRQAUUUUAOHWnU0dadSYwooopAFKKSlFABRRRSEFFFFABSjrSU4daBoWg47jI9PWiikWePfEXwW9+P7f05P8ASFXM6Dnco7/UV4LvduJBgjt6V9tEcf0ryTxl8OY9RdtT0ACKc8tEeA5rop1OjOKvQveSPAPrUkU0sEiywsyOvRlJUj8RUl3bXFjO1ndo0Myn542Hf2PcVX6HHf0ro30OPVM7uz+JHimxRYvtCzKvAMi5P59a02+K/iXGFS3Bx94givMjlfvcfWkznpS9ii1UlfU6nU/GviXV4zHd3jCMnOyIbf16kVyik89TnnJ65pw5pe2R0FPksS5XYnTmj+lbWieH9U8QXIt9OiL8ElyPkA6ZJ6fhU/ibQJfDepnT5m83KIyy9A3HP5UXWw2tDnD/AI/yr2T4O/8AISvv+uK/zFeONn7v5ivbvg/aSefqF/giNgkY9z1rOt8DLofGj3OkPVfqKWkPYn1FcB6Z8g+K/wDkY7//AK7tXPDqPqK6PxZGY/El+mP+W7Vzobaw+or1I/CjyJ/Ez3L4Nf8AHtqH+9H/ACNe2ivFfg4hFpfvjgvGM/ga9rWuCt8TPSw/wCilpaQ1kbxV2RyxiVCjqGUjDKe4NfOvj3wJLpEj6xpC77RyTInUp9Pavo2mSRJMhjkUOjDDKehBq6c+VqRnVpcysfD5BXDHkEZB/pQeMMDyOg7GvUfHvgSbRZn1bSlMlmxJePGdp/wrzHlsHHUZ46AV6VN86uebUXI7Ho3gLxzJoMy6bqLF7CQ43HrGx9PavpeKRZkWaNgyuuQR396+H2BGCelfWPw+eeTwlYPc8vsIB9s8Vx4iFnc7MNO+h2lYXiHQ7TxBpkmnXgyrA7W/ut2NbtIRkEHpXMnZ3OuVranxprWkX+hX7afqK4kQYBHRh2NZGOMjpX194k8MWHiSzNvertkAPlzDqpr5r8SeENZ8OT7bqPfAchJYxkHHrjpXbSrc2h5deg4O62OTOcYrpNK8W+INGXbY3ciR4xtf5l49B2rm8jGc0uCBk9D3rZpNWZhGTi7o9Qi+LXiRVCSLA+P4tpJ/HpVO9+KHiy6GyKWOIdysfI/M152T60VChDojV1ptalm7vbq/nM95K8z/AN5mJ/IHpVYUd8VJFHJLIII1Ls/G1Rlj6YrSxm5NqxEeoHTnr7d6+jfhDo89jpM+oTgr9qcbARghVzj+dYHgz4XyGRNT8SDGMFLbqPq1e8QKEQIq7AowFHQCuOvV0cEd2GoWtMmooorjO0KKKKACiiigAooooAKKKKACiiigDh6KKK9k80KKKKCwooooICiiigAooooLCiiigAooooAWiiipAKKkooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKfQAUUUVIBRRRQAUUUUAFFFFABSjrSDrUtABRRRUgFFFFABRRRQAUUUUAFFFFABRRRQMKUdaB1p9BYUUUUgCiiigAooooAKKKKACiiigAHWn00dadSYC0UUUgCiiigAooooAKKKKAClHWkpR1pAOooopAFFFFABRRRQAUUUUAFFFKOtAAOtPoopDQUUUUjQWiiigQUUUUAFFFFABSjrSUo60AOooopAFFFFABRRRQAUUUUAFOHWm04daQDqKKKQwooooAKUUlLQAUUUUhBRRRQAo606mjrTqGNBRRRSLCkI3YXpnvS0tFwTtqYmraBo2tx+VqcEcwHAcjDg+xry3VPhCm4tpF3tUniOUZx9GHNe20lONRp6GU6PPqfMF18NPFlsdot1mGeGjfn8jVI+AfFzEJ9hk/NT/UV9WYpVHNavEOxmsKk7s+arP4WeKJ2AuPJt19WYkg/7orvdH+E+kWjrNrErXjgZwPljz9Bya9corN1m9C1SgtUU7W2tbGIQ2UYjhUfdVcLWJ4l8Lab4mtxFfLskj/1ci9RkV09JUJtO6LsnozxCD4OxpLtub8yQ5y2F+bHoDXr2laXZ6NaJYWCeXEg4HUn3JrRzS5qpTk1ZjVKC1QUdeCMj0oorMZ5N48+H8uvXA1XSyq3GP3sbHAfHTB9a8ztPh34ruZlhktvIAPO8grjuQa+pKXAreNdwVkYyw3M7mB4b0C18N6Wmm2vIByz92J6k10K02lWsZPmd2axXKrD6KKKRQUlLRUt2Vxp2ZFNEk0bRSKHRhhlPQg14B4p+GF9DdtdeHU82CQljGWwVPtX0JSVpSquPvIzqUOc+ZtG+F+vX86jU4/stvkFjuySAeRivo+yt4bO2jtbddkUahUHsO9WsUtOdVz1YU6fJoFKtJSrWZsPqGeCK4jMMyK8bcMrcgipqSlr0B+Z5brXwr0LUHMtgz2chyflGU/EGvO7/AOE3iS2P+iSQ3I9SxQ4+lfS1IK1jWkjCWGhLVHyQ/wAPfFyPt+wufdSMfnmrdv8ADPxfcNtNssfvKwwPyr6toqniJdCPqqWrPBNN+D1zuB1a8VV/iWFck+2T0r1XRfCmg6AobTrdVcDBkf5n/M109IOKzlVnJWbN4UoR1QA9uvenigUtZFhRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAHD0UUV7JwBRRRQAUUUUAFFHNFAC0UVJUgR0VJRQBHUlFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUU+ilcAooopAFFFFABRRRQAUUUUAFFFKAc0AKOtPooqQCiiigAooooAKKKKACiiigAooooAKUdaSlHWg0H0UUUgCiiigAooooAKKKKACiiigApR1pKUdaAH0UUVIBRRRQAUUUUAFFFFABRRRQAo606mjrTqTAKKKKQBRRRQAUUUUAFFFFABSjrQOtPoY0FFFFSaBRRRQAtFFFAgooooAKKKKAClHWgdadSAKKKKACiiigAooooAKKKKACnDrQvWnUmMKKKKQBSikpRQAUUUUhBRRRQAUUUo60DQDrTqKKRYUUUUAFFFFABRRRQAU5etC9adSYBRRRSAKKKKAClFJSigAooopAFFHbNFABThQtOoAKKKKACiiigAooo69KACijB9KMGkNbhTlpBx1pwIPSgoWiiikAUUUUDCiigUgCiiikIcKWkFLQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQBxFFSUV65wEdFSUUAR0VJRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFPoopAFFFFIAooooAKKKKACiiigAoopQDmgAHWpKKKkAooooAKKKKACiiigAooooAKKKKACiilHWgYDrT6KKRYUUUUAFFFFABRRRQAUUUUAFFFA60AKOtPooqQCiiigAooooAKKKKACiiigApR1pKUdaQDqKKKQBRRRQAUUUUAFFFFABRRRQAo60+mgc06kxoKKKKRoFFFLQAUUUUCCiiigAooooAUdadTR1p1JgFFFFABRRRQAUUUUAFFFOXrQADrTqKKkYUUUUAFKKSloAKKKKQgooooAKUdaSnDrQNC0UUUiwooooAKKKKACiiigBy9adTR1p1JgFFFFIAooooAKUUlKKAClX7wpKO4pAZ9lqlpqLzR2xINs5iYHgZrQritAvLWyGsXN3KscS3shLE8dB0q0PGOiqR886RnpI0L7T/wACIxj3oA6wdafVaC6t7qFLm3cSRt0YHIzWWmrxLrDaOYmVtnmBichucED6UAbtFNJIzg8YrIXV45tXk0eNSWhjDuw6DJxj60gNisjW9astCtBeX2/YTt+RdxzWwVCnivO/iZ/yAE/67LQNblq2+IXh25uI7aIygy8DcmBmug1nWdN0S2W5vydjfdAGWPHpXzRpAzqtrz/y3XP03c10vjrXW1jWmitmBitSY0x3Y9T+VBR6avxH8L5UxGQ7sAAIckn863dW8S6ZocUM14ZAtxym1cnpnkHGK8n+Hfh06jeLrF2gFvbNtjH9589fwra+Kj/uLAv98O/Hb8DQxnSx/Ebw0zhWklXP99MA/lmuxsb+z1KL7VZSrJGePlxwfevGLbwtpFx4I/tV0aK4WNiHz1OelX/hNcTH7dbf8shtcf7x4NSI9morL1PWdO0iJZ7+YRg52jPLY9B3rItvGfh2/ieRLkIqDJ38H8KYHV0dOtY1jrWlalA0+nzLKkfLHniqY8X+HdkjC8jbyhlgPf69aGM6XuK8w+JWo32m2tpJYTmFmkbOO4ArptP8XeH9VuhbWtwfMbgBhtz9M9a4v4rjFhYsT0kYZ/CpEd14SuJrvw9aXFwxeRk5Y9+a6OuT8HskfhazeVtirGSSTgda2oNW0y5l8m2uI5X9FdWP5A0AagpaRfY0tABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAHGUUUV65wBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFPoopAFFFFIAooooAKKKKACiiigAoopQDmgBR1p9FFSAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQMUdafTB1p9JlhRRRQAUUUUAFFFFABRRRQAUUUUAFKOtIOtPpMBaKKKQBRRRQAUUUUAFFFFABRRRQAo606mjrTqTAKKKKQBRRRQAUUUUAFFFFABSjrSUo60Atx9FFFSahRRRQAtFFFAgooooAKKKKAClHWkpR1oAdRRRSAKKKKACiiigAooooAKcvWm04daQDqKKKQwpRSUooAKKKKQgooooAKKKKAFHWnU0dadQxoKKKKRYUUUUAFFFFABTl602nL1oAdRRRUgFFFFABSikpRQAUUUUgCkOeMdcj+dLSHJwB13D+fNAHm3hzSo77WdTu7z54oLpgkfVN397HevRmiQxmIKNrdA3K/8AfPpXnHh3WFsdb1O1vcR28t23lP0Ab0J7etekyPHHE0sjYVed3t9aQHC2ca6F4qGl2mRaajGZUjJ+VZEPzFR2z6VpeKbeVIIdatBtm09g5I7xn74H4VTspBrviY6nbrm106Mxxv8A3nfrg967ORQ8ZjK7lIAYHuCOR+VAFT+0LU6edTz+5KebnttAyKx/CdvMLKTU7v8A19+5mOeoQ/cH4CuUeK8jb/hCeSjz+ZG3b7N94gn68V6hGqxgRIMKgAH07UwJa87+Jn/IAT/rsteiV538TP8AkX0P/TZaTGtzwWIyiUeR/rMnbj1zTSrbmHRySD/vVoaNzrFoPSYZ+ma6jx7oA0rVftNv8sF3lx2CsOSPrUrco9g8IXWn3Og2raeAsaAKy+jj72fqa4n4qYEFhuGcGT+VcZ4J8QNoepJHIzNa3JAkBH3Wzw1dj8VGUw6aR86lmKnOM8d/aqYHIWeh+LtS0NJLNjLabciMNgHHXNdP8Oddt7S5bQJ4vKlmckMepYdVP4dKs+F/Gmi6V4bht7lyZow2UUep71yHhmG51nxfHeWy42ymY+gUe/vUgWNfe48U+MG07zvLRX8qMYyAF6kV0niH4d2Wn6PJe6e7edEuXJ7jviuc8Qi68L+MX1COPcrSebHxwQ3UfWuh8S/EOwv9JkstLV/NlX594I2+vWgCx8N939j6kSOCzFT7bOlcR4M0K217Vza3mTAibnUcbvSu2+Gu0aLqJDZyzE+mdlZHwvBGt3Gf+eP9TQBieKtJh8MeIIDphZcbJEVjnbzzzXafE+QtYWDkbjub5T0JZQc/hWJ8TedftwP+eK/+hVsfE0E6ZpoHXJbP0QUgOPu/ENxe6Hp3hfTSWDIBLj7zOW4UH09a9c8I+E7fQLfzplSS8flnAxt/2R9K8evdCubDSbHxHYHCOoLFeqMDkGvZfBXiZfEFgVn4uocCQevow+tAHZrjnHFOoBB6UUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAcZRRRXrnAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABT6KKQBRRRSAKKKKACiiigAooooAKKKB1oAUdakooqQCiiigAooooAKKKKACiiigAooooAKKKKAClHWgdafQaBRRRSAKKKKACiiigAooooAKKKKACiilHWgAHWn0UVIBRRRQAUUUUAFFFFABRRRQAUo60DrTqQBRRRSAKKKKACiiigAooooAKKKKAFHWn00DmnUmNBRRRSNAoopaACiiigQUUUUAFFFFACjrTqaOtOpMAooooAKKKKACiiigAooooAcvWnU0dadSYwooopAFKKSlFABRRRSEFFFFABRRSjrQNAOtOoopFhRRRQAUUUUAFFFFABTh1oXrTqTAKKKKQBRRRQAUopKUUAFFFFIApDuxleo5paKAM1dJ04rPH5QaO6bfIpHBOMVhr4H0BZBI0TkA5EZmcoPwz+ldgOtOoAqQwQ20QghQRIg4EYwBWVqeo6jZTBLXTpbtSuQ0TKuT6HPQV0FJSA5bRdPvHv5Nd1kbbmVdkcQ6Rx5zg44znvXUAbcg9Sc06igArH1rRbLXbT7DfBmj3bvlbacj3rYopgcPB8PvDkM63CRyb1YOD5hPIrpNW0mx1i2NpfpvRjnrgjjHWtSiga3PP3+G/hhl2tDKy4wD5pJ/Ktq78JaRqNlbWF6sjx2oxGd5DY6ckV01OWkUcGfhz4XBH7qU/WQn+ddTpukafpCGLT4RHkYJ9a1aKVwMrUtJ07Vovs+pQLIDnqOR9D2rGtPBHhvT43itrVT5qlG8zLnB7ZPSuuoouBjWGi6XpsUsFhAsMcow4XPPGKp6P4U0fQ7hrqyRlkddrEnORnNdLRQ2BzOr+E9H126S8vlZnjAUEHHQ5qfV/DWma5FDBfKzJACEAOPvDBz69K36KkDLi0awi0tdG8sNbKu3afSsPTvBGiaTfLf2ayrIpP8R2/iPSuxFOoAaue/6U6iigAooooAKKKKACiiigAooooAKKKKACiiigDjKKKK9c4AooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACin0UrgFFFFIAooooAKKKKACiiigAooooAKKKKAFoHWgdakqQCiiigAooooAKKKKACiiigAooooAKKKKAClHWkpR1oAfRRRSNAooooAKKKKACiiigAooooAKKKB1oAUdafRRUgFFFFABRRRQAUUUUAFFFFABRRSjrSAB1p1FFIAooooAKKKKACiiigAooooAKUdaSlHWgFuPoooqTUKKKKAFooooEFFFFABRRRQAUo60lKOtADqKKKQBRRRQAUUUUAFFFFABTl602nL1pAOooopDCiiigApRSUtABRRRSEFFFFABTh1pB1p1A0FFFFIsKKKKACiiigAooooAcvWnU0dadSYBRRRSAKKKKAClFJSigAooopAFFFFADlp1NFOoAKKKKACiiigAooooAKKKKBrcVafTFp9SygooooAKKKKACiiikMKKKKQhRTqQUtABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQBxlFFFeucAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUU+ilcAooopAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFKOtABzUlABRRRUgFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAKOtOpo606kwFooooNAooooAKKKKACiiigAooooAKUdaQdafSYC0UUUgCiiigAooooAKKKKACiiigApR1oHWnUgCiiikAUUUUAFFFFABRRRQAUUUo60AA60+iikNBRRRSNAoopaACiiigQUUUUAFFFFACjrTqaOtOpMAooooAKKKKACiiigAooooAKcOtA606kxhRRRSAKUUlKKACiiikIKKKKACiiigaFHWnU0dadQywooopAFFFFABRRRQAU5etNpw60AOoooqQCiiigApRSUooAKKKKQBRRRQAU5abTloAdRRRQAUUUUAFFFFABRRRQNbhRRSrSKFWnUUUgCiiigAooopAFFFFIApRSUooAdRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAcZRRRXrnAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFPopXAZRT6KLgFFFFIAooooAKKKKACiiigAooooAKKKKACiiigAooooAKUZzQAc1JQAUUUVIBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABSjrQOtOoAKKKKQC0UUUGgUUUUAFFFFABRRRQAUUUo60AAzmn0UVABRRRTAKKKKACiiigAooooAKUdaSlHWkA6iiikAUUUUAFFFFABRRRQAUUUUAFOHWkHWn0MaCiiipNAoopaACiiigQUUUUAFFFFABSjrSUo60AOooopAFFFFABRRRQAUUUUAFFFOHWgAHWnUUVIwooooAKUUlLQAUUUUhBRRRQAUo60lOHWgaFooopFhRRRQAUUUUAFFFFADl606mr1p1JgFFFFIAooooAKUUlKKACiiikAUUUUAFOFC06gAooooAKKKKACiiigAooooGtwpVpKctIodRRRSAKKKKACiiikAUUUUgCnCkFOoAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigDjKKKK9c4AooooAKKKfQAyin0UrgMop9FFwGU+iigAooopAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFA60UoBzQBJRRRUgFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFKOtAAOtOoopALRRRQaBRRRQAUUUUAFFFFABRRRQADrT6aOtOpMBaKKKQBRRRQAUUUUAFFFFABRRRQAo606mjrTqTAKKKKQBRRRQAUUUUAFFFFABRRSjrQADrT6KKQ0FFFFI0FooooEFFFFABRRRQAUUUUAKOtOpo606kwCiiigAooooAKKKKACiiigBaUdaB1p1JjCiiikAUopKUUAFFFFIQUUUUAFFFFACjrTqQdaWhjQUUUUiwooooAKKKKACiinL1oAB1p1FFIAooopAFFFFAC0UUUgCiiigAooooActOpop1ABRRRQAUUUUAFFFFABRRRQNbirT6YtPqWUFFFFABRRRQAUUUUhhRRRSEKKdSCloAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigDj6KKK9U4AooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAFoooqQCiiigAooooAKMGlANPoAYAc0+iigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKUdaB1p1ABRRRSAWkoooAWiiig0CiiigAooooAKUdaQdafSAWiiipAKKKKYBRRRQAUUUUAFFFFABSjrSUo60gHUUUUgCiiigAooooAKKKKACiiigBR1p9NA5p1JjQUUUUjQKKKWgAooooEFFFFABRRRQAUo60lKOtADqKKKQBRRRQAUUUUAFFFFABS0lKOtACjrTqKKkYUUUUAFKKSloAKKKKQgooooAKUdaSnDrQNC0UUUiwooooAKKKKACiiigApw60L1p1JgFFFFIAooooAKKKKAFooopAFFFFABTlptOFADqKKKACiiigAooooAKKKKBrcKKKVaRQq06iikAUUUUAFFFFIAooopAFKKSlFADqKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooA4+iiivVOAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAWiiipAKKKKACiiigAoowaMGgAwaMGpKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKUdaSlA5oAdRRRSAKKKKAFooooNAooooAKKKKACiilGc0AABzT6KKgAooopgFFFFABRRRQAUUUUAFFFFABSjrSUo60gHUUUUgCiiigAooooAKKKKAClHWkpR1oBbj6KKKk1CiiigBaKKKBBRRRQAUUUUAFFFKOtAAOtOoopAFFFFABRRRQAUUUUAFFFLQAUo60DrTqTGFFFFIApRSUooAKKKKQgooooAKKKKAFHWnUg60tDGgooopFhRRRQAUUUUAFFFFADl606mjrTqTAKKKKQBRRRQAUUUUALRRRSAKKKKAHLTqatOoAKKKKACiiigAooooAKKKKBrcKVaSnLSKHUUUUgCiiigAooopDCiiikIKcKQU6gAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigDj6KKK9U4AooooAKKKKACiiigAooooAKKKKACiiigAooooAWiiipAKKKKACiiigAoowaMGgAwaMGpKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiilHWgAA5p1FFIAooooAKKKKBi0UUUFhRRRQAUUUoBzUAABzT6KKACiiimAUUUUAFFFFABRRRQAUUUUAFFFFACjrTqaOtOpMAooopAFFFFABRRRQAUUUUAFOHWkHWn0MaCiiipNAoopaACiiigQUUUUAFFFFABSjrQOtOpAFFFFABRRRQAUUUUAFFFFABSjrRSjrSAdRRRSGFFFFABSikpaACiiikIKKKKAClHWkpw60DQtFFFIsKKKKACiiigAooooAKcvWm04daAHUUUVIBRRRQAUUUUALRRRSAKKKKACiinLQACnUUUAFFFFABRRRQAUUUUAFFFFA1uKtPpq06pZQUUUUAFFFFIAooopAFFFKKAFFLRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQBx9FFFeqcAUUUUALRRRUgFFFFABRRRQAUUUUAFFFFABRRRQAUUUYNABRg0YNSZFAEeDRg1JRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABSgUCnUAFFFFIAooooAWiiig0CiiigAoooqAAdakpgBzT6YBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABSjrSUo60gHUUUUgCiiigAooooAKKKKACiiigBR1p9MHWn0mNBRRRSNBaKKKBBRRRQAUUUUAFKOtJSjrQA6iiikAUUUUAFFFFABRRRQAUUUtAAOtPpo606pYwooooAKUUlKKACiiikIKKKKACiiigBR1p1IOtLQxoKKKKRYUUUUAFFFFABRRRQA5etOpo606kwCiiikAUUUUAFFFFAC0UUUgCiiigApwoWnUAFFFFABRRRQAUUUUAFFFFA1uFKtJTlpFDqKKKQBRRRQAUUUUgCiiikAUopKcKAFooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooA5CiiivTOAKKKKACiiigAoowaMGgAoowaCOOlACEYDE8bOW9vrWbNrOj2zbLi8gjb0aRQf514b4y8Q67qGqf2V5TwQqfLihXIaQeue9V7X4c+KJoPtIjSJf7sj4f8hmsvbGnIj6Gguba5i8+2lSSP+8jAj8xU/XHvXym6a/4U1XeS9tdJ7n5/p/CRX0R4V8QReJdKW7+5cr+6nX+6w6Eexop1OfcdSmdBJNFDH5szqif3mOB+dUF1rR3fy0vYC3oJFz/OvCfFdn4q1TxG9hepI+85gEeQmPfHFUtQ+H/iLTLU30sSMmM/I+WHuaj2zFyI+mAylBICCrdD2NL6+1fOfgPxLeadq8Fg5L21yPLYO27a/Yj617xrmpWujWE+o3WcQIDtHdj2rWnNTHUpmlI6RJ5kjBV9ScCsga7oe/y/t1vu9PMXP86+drvUPEHjLUCieZK38EEZIRAf7xHStOb4a+JI7czeTHIw6xq/Ix71kq76F+zPoyOSKZPMhYOo7qcipARzjt19q+UdH13WvDV+XgLqsZAkhkJJ/I96+nNN1CLV7GLUbfGydATj2q6c+cipTNKiiitTIKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKXFGKdQAUUUUgCiiigAooooAWkoooAWiiig0ClAOaQdakqACiiimAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUo60DrTqQBRRRSAKKKKACiiigAooooAKKKKAFHWn00A5p1JjQUUUtI0CiiigQUUUUAFFFFABSjrQOtOpAFFFFABRRRQAUUUUAFFFFAC0o60g60+kxhRRRSAKKKKAClFJS0AFFFFIQUUUUAFKOtJTh1oGhaKKKRYUUUUAFFFFABRRRQAU5etNpw60AOoooqQCiiigAooooAWiiikAUUUUAFFFKvWgBRTqKKACiiigAooooAKKKKACiiiga3FWn0xafUsoKKKKACiiigAooopDCiiikIUU6kFLQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAchg0YNSUV6ZwEeDRg1JRQBHg0YNSUUAFFFFABRRRQB574q8U6P4e1SIX9o006rvVxg4rf8AD3iXT/EkDT2QKFMBkPUZ6ZHvWJ408J22uW8+oJuN3HCBCB046g+9eXfDrVjpniEW8jARXaeScnAD9QT9K5/aWnqa+z9w6v4s28QisboD94zlCfYVT+Esrm5vbUn5Sgl/EcCm/FLVLe4ls9PhdXMe9iVOcEnvV/4T2DrFe6oc7T+6A+lNaz0LfwHsJQbwwIOB8rEfMBVLVgjaZdIwziFv5Gr5qhqX/IPuf+uJ/ka2Zznyho7karZSJx++jz+lfR3jHw/e+ItMWxspBGVlLvv43DsK+c9GVhf2mR/y1X+Yr6+b7x+tcuH2OiofKNrqmteENUnS3kCXETGOQBcg4r6Z0q8GpaXbXrjBlhWTj1Pavmfxb/yM2pD/AKeG/wDQRX0X4WBHhzTc8fuY/wCZq6L+wEzwn4h2sVv4quEg+US+XIf+BCvVvhqzy+FLfcfuySJ/3ya8v+JZC+LH3cfuof5GvTfhiQfCqAcn7ROfwyKKfxhU+A9DoooroOcKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiilwaAEpcUYNOoAKKKKQBRRRQAUUUUAFFFFAxaKKKCwoooHWoAUA5p9FFMAooooAKKKKACiiigAooooAKKKKACiiigAoopR1pAA606iikAUUUUAFFFFABRRRQAUUUUAFKOopKUdRQBJRRRUggooooNAooooAKKKKACiilHWgAHWnUUUgCiiigAooooAKKKKACiiloAB1p9NA5p1SxhRRRQAUopKWgAooopCCiiigAooooAUdadSDrS0MaCiiikWFFFFABRRRQAUUU5etAAOtOoopAFFFFIAooooAKKKKAFooopAFFFFABThSL1p9ABRRRQAUUUUAFFFFABRRRQNbhRRSr1pFCrTqKKQBRRRQAUUUUgCiiikAUopKUUAOooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAOUooor0zgCiiigAooooAKKKKACiiigCPbkYPQ8Gvn7Wfh74jOq3Eul2wa38zfGS6KRn8a+haSs501Pc0p1D5/074X61cy7tVcW6f7Lb2r3HTtOttLtIrG1XbDEu1cdSf7zVoUUQgobCqVAPSqd/G8tlPHGMs8ZUD1JGKuUVoQfOGnfD/xVDe20kliqqkgZj5icAEe9fSB5bI9abS1nCCgtC51Ofc+f/EngbxJqOtXV3b2gkhlldgS6A8gY717L4fs57HRrW1uYxHNFEisAcjgnIrbpKIQsP2h4p418H+INY197vTbQSQGOIAh0Q5APqe1d34K03UdI0FLLUozFKskpPzB+CRjp612NFCp+/zh7QKKKK0MwooooAKKKKACiiigAooooAKKKKACiiigAopcGjBoAMGnUUUgCiiigAooooAKKKKACiiigBaKKKDQKKKKgApQDmgA5p9MAooooAKKKKACiiigAooooAKKKKACiiigAoopR1pAA606iikAUUUUAFFFFABRRRQAUUUUAFFFFABTgOaQdafQwQtFFFSaBRRRQAUUUUAFFFFABSjrQOtOpAFFFFABRRRQAUUUUAFFFFABSjrSUo60gH0UUUhhSikpRQAUUUUhBRRRQAUUUUAFOHWkHWnUDQUUUUiwooooAKKKKACiiigBaUdaF606kwCiiikAUUUUAFFFFAC0UUUgCiiigAoopR1oAUU6iigAooooAKKKKACiiigAooooGtwpy0i0+kygooopAFFFFABRRRSGFFFFIQU4Ugp1ABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAcpRRRXpnAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAXCil2k9BSUAFFGRRkUAtQopCQOpFO5oG0JRS800EHoaBC0UUUAFFFFABRRRQAUUUUAFFFLg0AJS4NGDTqACilopAJRS0UAJRRRQAUUUUAFFFFABS0UUDCiiioLCgdaKUA5pgPooooAKKKKACiiigAooooAKKKKACiiigAooooAKUZzQOtOpAFFFFIAooooAKKKKACiiigAooooAKKKUDmgAA5p9FFIaFooopFhRRRQAUUUUAFFFKOtAAOtOoopAFFFFABRRRQAUUUUAFFFFAC0oHNIOtPpMYUUUUgClFJSigAooopCCiiigAoopR1oGgHWnUUUiwooooAKKKKACiiigApaSnL1oAB1p1FFIAooopAFFFFABRRRQAtFFFIAooooAKcOtIOtPoAKKKKACiiigAooooAKKKKAW4UUUq9aRYq06iikAUUUUAFFFFIAooopAFFFKKAFFLRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAHKUUUV6ZwBRR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KNuWTfnrnBqWigCIQxR4MaJn/dFS/N/Hg+mO1FFACHp60h6Yc7/bbkU6igCFbeIZZVRWPfaOKmOTjJzj0GKKKACiiigAxnjpSqBjJ6joaSigEOI3DaxBB+tAQMMHgDtSDrTqRoJ8+3O5hnrg/wBKPmIDOck+hP6ilopAFFFFMAooooAKKKKACiilpAKvWn0wA5p9SAUUUUAFFFFABRRRQAUUUUAFKOtJSjrQA6iiikAUUUUAFFFFA0FFFFBYUUUtACr1p1NAOadSYBRRS0gCiiikAUUUUAFKOtJSjrQA+iiigAooooAKKKKACiiigEFKvWkpV60ix9FFFIAooooAKKKKBhRRRUiClFJThQAtFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQBzFFFFegcAUUUUAFFFFABRRRQAUUUUAFFJRQAtFFFABRRRSAKUYY42rj/apKKZehVmsrO4/18KH6CqLeH9Ff71ola9FR7KHY2hi68PgqMwh4X8Pj/lyi/Nv8a0YNM0624ghWP6DNXKKn2a/lKnjcRU3qMcAF6ED/dFJRRg1qc0nKXxsKKMGlwaCRMGlwafRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRSAKKKKCwooopjCiiigAooooAKXBoHWnUgGgHNOoopAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAC0oBzQAc0+pAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAQo606mjrTqTNAooooAKKKKACiiigAooopALSr1pKUA5qQH0UUUAFFFFABRRRQAUUUUAFFFFABSjrSUo60AOoozRSAKKKKACiiigaCiiigsKcvWkpQDmgB1FFFSAtFFFIAooooAKKKKAClHWkpR1oAfRRRQAUUUUAFFFFABRRRQAUoBzQvWnUmNbi0UUUigooooAKKKKkAoopRQAopaKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooA5iiiivQOAKKKKACiiigAooooAKMEkqOo60deKzNa1SDRNNl1G6QvFAcFR1NIDSory/8A4WromcfZbj8qQ/FbQxwbWfOM9ulL2i2NvZdT1Lvt7iiszSdTi1zT7fUIEKRSqWCnrgHB/I1p0zEKKMGjByFwcmgAooHPSlwaYCUUuDRg0AJg0YNSUUgGYNPoopgFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUhhRRRTLCiiigAoopR1pAABzTqKKQBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAtKAc0AHNPqQCiiigAooooAKKKKACiiigAooooAKKKKACiiigAoopR1oBCgHNLRRSNAooooAKKKKACiiigApaKUA5pAABzT6KKkAooooAKKKKACiiigAooooAKKKKACiiigBR1p1IAc0tDAKKKKQBRRRQNBRRS0FCr1p1NAOadSYwoopaQBRRRSAKKKKACiiigApR1pKUdaAH0UUUAFFFFABRRRQAUUUUAKvWnUgBBpaljQtFFFBQUUUUAFFFFSAU4Ugp1ABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAHMUUuDRg16BwCUUuDRg0AJRS4NGDQAlFLg0YNACVm61pFtrmnvpt5nyJDkgda08GnmkB4nr3gfwhoGmNeXjXDKW2om/wCYkc4rhfCvhmPxRqr20e63tYv3jfx8HjGa9I+Lhb+zbKNTgNOxP/fNY/wh/wCPu+/64r/6FmsFBc51uo/ZnsWj6Vb6PpcGk2+fKhXaM9eWya0cGlaRF+8wGDjnjn0oV0b7rKcjPBB49a2ujlsxScAmvPZPiX4XWZreRpiyNtBEbdc49fWvQj90/Svja6/5CMv/AF3b/wBDNZ1KnJsaU6Z9ixy+bCsmMbgCPcHoalqraf8AHpb/APXGP+VWq1MgooopgFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUhhRRRTLCiiigAooooAKKKKAClFJRQA/IoyKZSgHNIB1FFFIAooooAKKKKACiiigAooooAKKKKAFoHWigdakCSiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKUA5oHWnUgCiiigEFFFFBoFFFFABRRRSAcvWn0xetPqQCiiigAooooAKKKKACiiigAooooAKKKKAClHWkpwBzQAtFFFIAooooAKKKKAClHWkpaBofkUU0A5p9SywooopAFFFFABRRRQAUUUUAFKOtJSjrQA+iiigAooooAKKKKAClXrSU4A5pAOpKWikNBRRRQUFFFFAwoooqRCinU0U6gAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigDnKKKK9A4AooooAKKKKACiiigAoPSig9KAPIPi3/wAg+x/67P8A+gVkfCL/AI/L/wD64itf4t/8g+x/67P/AOgVkfCL/j7v/wDriK5/+Xh0f8uzsfic8kfhsPE7IPOT7vBrz34YTySeIyspZz5LfeYletd98USD4YyO0yf4V558LQR4kYHqIGJ+maJfxAh/DPos/dP0r42uv+QjL/13b/0M19kn7p+hr41uTnUJT/03Yf8Aj9KuGHPsG0/49Lf/AK4x/wAqtVVtD/okB/6ZRj8cVaroRlcKKKKZAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFIYUUUUywooooAKKKKACiiigAooooAKXBowadSAaAc06iikAUUUUAFFFFABRRRQAtFFFSAUUUUAFFFFABSgHNIOtSUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRSgHNAAAc06iikAUUUUAgooooNLhRRRQAUUUUgHL1p9MXrT6kAooooAKKKKACiiigAooooAKKKKACiiigBR1p1IAc0tDAKKKKQBRRRQAUUUUAFOXrTacvWhjQ+iiioLCiiigAooooAKKKKACiiigApR1pKcBzQA6iiigAooooAKKKKAFXrT6YvWn1LAKKKKACiiipBbhRRRQWFFFKKAFFLRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQBzlFFFegcAUUUUAFFFFABRRRQAUHpRRQB5B8XARp9hn/AJ7P/wCg4rI+ERAu9QYngRAV6V4s8Of8JLpP2ND5dxG2+GRui+xrw4+BvGlpcDbZySDcAzRvwQD14IrnmnznRBrksetfEqBpPCkjICwjljdsdlz1ryn4c6hFYeJoWuSEWWNo8se/avoqSzju7L7LeLuVoQrr6n/61fP+tfDTX9OnJ0iMXlvvJVlP70D8xTmvtiptH0Je3UFlZvdXDhECnknjOK+QAxub4TBTmaQkKOu5mHFbreH/ABddKsBtrqVT1XDf1avTfBvw6m0+7XVdaOJIxmKJedp9DSkufcuMlDY9chj8mKOI87VUH8BUlFFbnKFFFFMAooooAKKKKACiiigAooooAKKKKACiiigAooopDCiiimWFFFFAgooooGFFFFABRS4NGDSAMGjBp1FFwCiiikAUUUUAFFFFABRRRQAtFFFSAUUUUAFFFFABRRQOtACgHNPoooAKKKKACiiigAooooAKKKKACiiigAooooAKKKUA5oAADmnUUUgCiiigAooooAKKKKBoKKKKRYUtJTl60AABzT6KKkAooooAKKKKACiiigAooooAKKKKACnAHNIOtOoAKKKKQBRRRQAUUUUAFFFFABTl602nL1oY0PoooqCgooooGFFFFABRRRQAUUUUAKOtPpoHNOoAKKKKACiiigAoopV60gFAINOoopAFFFFABRRRUgtwooooLClFJSigB1FFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAHOUUUV6BwBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFADQZA24Ej8adRRSAKKKKY7BRRRQFmFFFFIQUUUUXBtLcKKKKYBRRRQAUUUUAFFFFIAooopl3CiiigAooopEhRRS4NBYYNGDTqKLgFFFFIAooooAKKKKACiiigAooooAKKKKAFoooqQCiiigAooooAKUA5pB1qSgAooooAKKKKACiiigAooooAKKKKACiiigAooooAUdadTQDmnUmAUUUUAFFFFABRRRQCCiiikaXFoooqQClXrSUq9aAH0UUUAFFFFABRRRQAUUUUAFFFFABRRSjrQAoHNLRRSAKKKKACiiigAooooAKKKKACnL1pOacoINJgOoooqRoKKKKCwooooAKKKKAClHWkpR1oAfRRRQAUUUUAFFFFABSr1pKVetIB9FFFIAooooAKKKKkFuFFFFBYUopKUUAOooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooA5yiiivQOAKKK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GnUAFFFFABRRRQAUUUUAFFFFABRRRQAUo60lKOtA0PooooLCiiigAooopMAoooqAClXrSdacAQeaAHUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFKKSigEPzRTRTs0FhRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAYFFFFdhwBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFA60UuDQA+iiigAooooAKKKKACiiigAooooNAooooFcKKKKBhRRRQAUUUUAFFFFABRRRQAUUUUXAKUA5oAOadSbAWiiioAKKKKACiiigAooooAKKKKACiiigBR1p9MHWn0AFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAIKUdaSnAHNBoLRRRSAKKKKGAtFFFQAUUUUAFKvWk604Ag0AOooooAKKKKACiiigAooooAKKKKACiiigApR1pKUdaBofRRRQVcKKKKBhRRRSYBRRR1qAFXrT6aAQeadQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABSikooBbj80U0U6gsKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAMCiiiuw4AooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAAdakpmDT6ACiiigAooooAKKKKACiiigAooooLuFFFFBAUUUUGgUUUUAFFFFABRRRQAUUUoBzSuAYNABzT6KgAooooAKKKKACiiigAooooAKKKKACiiigAooooAUdafTB1p9ABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQNIUdadSAHNLQywooopAFFFFIBaKKKkAooooAVetPpi9afQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAU4DmkHWn0AFFFFA0FFFFBYUUUUmAUq9aSlBwagB9FICDS0AFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAhRTs0ylFBY6ijNFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAGBRRRXYcAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRS4NACUUuDQAaADBoANPooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKB2CiiigsKKKKACiiigAoopQDmlcAAOafRRUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRSjrQAAHNPoooAKKKKACiiigAooooAKKKKACiiigAooooBBTgDmkHWnUGgUUUUrgFFFFIBaKKKkAooooAKOtFKvWgBQCDTqKKACiiigAooooAKKKKACiiigAooooAKKKUdaAFAOadRRQAUUUUDQUUUUFhRRRSYBRRRUAKDg04EGmUoODQA+ik3CgEGgBaKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooBbiinZplKKCx1FGaKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAwKKKK7DgCiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKXBoASijBpcGgAwaMGn0UAMwafRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQXdBRRRQAUUUUAFFFFACjrT6YOtPqLAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUoBzQAc0+kAUUUUAFFFFABRRRQAUUUUXAKKKKLgFFFFFwCiiii40xR1p1NHWn0my7oKKKKgLoKKKKAugooooC6CiiigLoKVetJSg4NAXQ+ikyDS0BdBRRRQF0FFFFAXQUUUUDCiiigAooooHYKcAc0AHNOoCwUUUUAgooopFhRRRQAUUUUMAoooqQCiiikAUq9aSlBwaAH0Um4UAg0ALRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFA1uKKdTRTs0FBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAYFFFFdhwBRRRQAUUUUAFFFFABRRRQAYNGDVj5KPkrPmZpyruV8GjBqx8lHyUczDlXcr4NGDVj5KPko5mHKu5XwaMGrHyUfJRzMOVdyvg0YNWPko+SjmYcq7lfBowasfJR8lHMw5V3K+DRg1Y+Sj5KOZhyruQYNPzUnyUfJRzMOVdyPNGak+Sj5KOZ9g5V3I80ZqT5KPko5n2DlXcjzRmpPko+SjmfYOVdyPNGak+Sj5KOZ9g5V3I80ZqT5KPkpc7DlXcjzRmpPko+SjnYcq7keaM1J8lHyUc7HyIjzRmpPko+SjnYciI80ZqT5KPko52HIiPNGak+Sj5KOdhyIjzRmpPko+SjnYciI80ZqT5KPko52HIiPNGak+Sj5KOdl8q7keaM1LhaMLRzsOVdyLNGalwtGFo52HKu5FmnZFPwtGFo5mHKu4zIoyKfhaMLRzMOVdxmRRkU/C0YWjmYcq7jMilp2F7UlHMw5V3EopaKOZhyruJS4ooHWjmYcq7htNG00/cKNwo5mHKu4zaaNpp+RS0czDlXcj2mlAINPoo5mHKu4UUUUw5V3CiiigOVdwooooDlXcKKKKGHKu4UUUVnYOVdwoooosHKgoooosHIgoooosHIhR1p2RTKXBoHyIdnNLimYPak+f0oDkRJijFR/P6UfP6UByIkxRio/n9KPn9KA5ESYowaj+f0oG8HpQHIiXaaNppmWoy1AciJACDTqhy1GWoDkRNRUWXNGJKA5ES0VFiSjElAKCJaKixJRiSgvlRLRUWJKMSUByolpR1qHElGJKBXLGRRkVXxJRiSgLljIoyKgxJRtk9KAuT5FGRUG2T0o2yelJhcnyKMioNsnpRtk9KlDJ8ijIqDbJ6UbZPSqAnyKMioCJByRTcvQwLORRkVWy9GXqbAWcigc9KgAkNBEvYUgLGCacAQaqjzgen607MvpQBaoqoXlHak8ySgC5RVPzJKPMkoAuUVT8ySjzJKALlFU/Mko8ySgC5RVPzJKPMkoAuUVT8ySjzJKALlFU/MkoDyHtQBcoqpul9KN0vpQBboqpul9KN0vpQBboqpul9KN0vpQBboqpul9KN0vpQNFyiqe6X0o3S+lBVy5SiqW6X0o3S+lAF6iqO6X0o3S+lAF6iqO6X0o3S+lAF6iqO6X0o3S+lAF6iqO6X0o3S+lAF6iqO6X0o3S+lAF6iqaGUnpVsHigBaKKKACiiigAooooAKKKKACiiigDAooorsOAKKKKACiiigAooooAKKKKAH7hRuFMooAfuFG4UyigB+4UbhTKKAH7hRuFMooAfkUtM2mn0AFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRUWAKKKKLAFFFFFi7hRRRRYLhRRRRYLhRRRQFwooopBcKKKKAuFKKSigY/IoyKZRQA/IoyKZS4NADsilpgBzT6ACiiigAooooAKKKKACiiigAooooAKKKKAAdafuFMooAfkUtMHWn0AFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAgooooNLoKKKUdaGF0KAc06iiswuFFFFABRRRQAUUUUAFFFHWgA60uDSgEGnUAMwaMGn0UANAINOoooAKKKKACiiigAooooAKXBoHWn0AMwaMGn0UANAOadRRSYBRRRUjQUUUUirhRRRTQwGO9OyPSm0VQDsj0pPlPakooaACF7UUUVIBRnFFFIBQ2eopQQe1NpQcGgB34UfhQCDS0AJ+FH4UtFACfhR+FLRQAn4UfhS0UAJ+FH4UtFACfhSjHpRRTAXijikooAXijikooAXijikooAXijikooAXijikooGh2RRkU2ikVcdkUcU2lFAXF/Cj8KXNFACfhR+FLRQAn4UfhS0UAJ+FH4UtFACEelApaKACiiigAooooAKKKKACiiigAooooAwKKKK7DgCiiigAooooAKKKKACiiigAooooAKKKKACiiigAoopdpoAfRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUrgFFFFQAUUUUAFFFFBoFFFKAc0AABzT6KKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAoooHWgBwBzTqKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigApwBzQAc06k2AUUUVAIKKKKDQKKKKACiiigA604Ag0i9afQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAo60+mDrT6ACiiigAooooAKKKKTAKKKKkaCiiikVcKKKKaGFFFFUAUUUUMAoooqQCiiikAoODTgQaZSg4NAD6KQEGloAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigEKKdmmUooLHUUZooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigDAooorsOAKKKKACiiigAooooAKKKKACiiigAooooAKKXaaMGgBKkpmDT6ACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooopXAKKKKLgFFFFFwCiiioAKKKKACiiigAoopQDQXcADmn0UUBcKKKKBhRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUDrRSgHNAD6KKKACiiigAooooAKKKKACiiigAoooouAUUUUXAKcAc0AHNLSbAWiiioAKKKKAQUUUUGlwooooAKKKKAFXrT6YvWn0AFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAKOtOyKZSjrQA+ikzmloAKKKKACiiigAooopMAoooqRoKKKKRVwooopoYUUUVQBRRRQwCiiipAKKKKQCr1p9MBxTtwoAWikBBpaACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooGhRTqaKdmgq4UUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAGBRS4NGDXYcAlFLg0YNABg0YNPooAZg0YNPooAZg0YNPooAZg0YNPooAZg0YNPooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooopXAKKKKgAooooAKKKKACiiigAooooAKKKKAFHWn0wdafQAUUUUAFFFFBdwooooC4UUUUDCiiigAooooAKKKKACiiigAoopQDmgAAOafRRQAUUUUAFFFFABRRRQAUUUUAFFFFDAKKKKzAKUdaSnAHNADqKKKACiiigAooooAKKKKAQUUUUGlwooooAVetPpi9afQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAKOtOzmmUooAfRSZzS0AFFFFABRRRQAUUUUmAUUUVI0FFFFIq4UUUU0MKKKKsAooopMAoooqQCiiikAq9afTAcU4EGgBaKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAQUUUUFNhSikooEtx+aKaKdQUFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAGJRRRXYcAUUUUAFFFFABRRRQAUUUUAFFFFABRSZGMg5+lGR68Hv2oAWiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKVwCiiii4BRRRRcAoooqACiiigAooooAKKKKACiiigAooooAKKKKAFHWn0wdafQAUUUUAFFFFABRRRQAUUUUGgUUUUAFFFFABRRRQAUUUoBzQAAHNPoooAKKKKACiiigAooooAKKKKACiiigAooooYBRRSjrWYCgHNOoooAKKKKACiiigAooooAKKKKACiiigEFFFHWg0FXrT6aAQadQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFACinZzTKUUAPopM5paACiiigAooooAKKKKTAKKKKkaCiiihFXCiiiqGFFFFMAooopMAoooqQClBxSUUgHgg0tMBxTgQaAFooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAQoNOzTKUHFBY6ijNFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAGJRRRXYcAUUUUAFFFFABRRRQAUUUUAFJuFKCCMqQfpUEi7omXIG5SPzGKAPJNb+Kn2a9kstGtxcGFirPJyeOuAOtbXhLx7D4iu30u8g+z3ZBZT1UgdRjsfavJdB1GTwF4kuP7UtnndVaMquBkZ+8N3WtTwjDP4h8cjWbSHyYfPMzEdAAMAD1zXNzM6XBI+jqKBk9BR+B/Kui5y3QUUmfUED1I4oDAgnsP88UxsWig5BAwcn0opJ3BhRRkc89O3ekJAxz179qYMWigZboCaUgjqDRcBKKQkBS2RgdaNwLBRznoRyPzoAWij+Et6ce9FK6BhRRkUuDRdAxKKKKgAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiinAGgAANOoooAKKKKACiiigAooooAKKKKC7hRRRQMKKKKACiilAOaAAA5p9FFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRSbAKUdaSlHWoAfRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAIOtOAINIvWn0GlwooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAUHFOzmmUoOKAH0Umc0tABRRRQAUUUUAFFFFJgFFFFTYaCiiihFXCiiiqGFFFFMAooopM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tboWCAJj1QjmvVL/AOJk2sRf2To9nIt1c/u8ykEYf5c4HTrSgVPc8atCDqca9/PUfjvFfaWDhG9gf0r4wMD6VqYgu08xrOcGQZ5JVu3rnHWvq3wxrcPiHSYtShhaDc7Jhju+71p01bcVfXY6GiiirOcKKKKACiiigAooooAKKKKACiiigApcGjBp9ADMGn0UUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQaBRRRQADrUlRjrUlABRRRQAUUUUAFFFFABRRRQwCiiiswCiiigAooooAKKKKACnAHNIOtPoAKKKKACiiigAooooAKKKKACjrRSr1oAUAg06iigAooooBBRRRQaXCiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigApQcUlFAD85paYDinZzQAtFFFABRRRQAUUUUAFFFFJgFFFFQNbhRRRTRYUUUVQBRRRQwCiiipsAUUUUgFXrT6YDinAg0ALRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAhQcU7OaZSg4oLHUUZzRQAUUUUAFFFFABRRRQAUUUUAYlFFFdBwBRRRQAUUUUAFFFFABRRRQAe4rgtT+HHh3V759RvRKZJDuYJKVG78q72igUJtbnHaT4D8MaM/nWlqWl/vSnfXX7TnGeT1b0/CnUUD52FRyRJNFJDJ91wRx6Hg1JRQB53bfC/wALW00c6fat0T7lzNnp07V6ISevfrRRQXObZw2u/D/Qdena8mjaGZhkmM4DN6sO9N0b4deHdGuVvERp5+7zHcPwAru6KfKg52czrvhDQvEbLJqkTFwMb0O1jg8ZP0p3h3wrpfhmOaHTAwSdg7bznkDArpKKQe0Zx+u+BvDviCVrq8gZJ243o2OKwIfhP4cjkzNLcTJ/cLAV6fRT5UHOynY2FnptstnYQrDGn3QvQ/73qferRA2sFAAfqp5FOopEHneq/DPw5ql0b3a8Erfe2H5SPYdq19A8F6H4cl+0WUZebGFkkOSo9K62inyov2jON1nwH4a1y6N5eQMszfeZGxmtzRtGsfD9imnaapWJGdvmOT81a1FIOdhRRRQQFFFFABRRRQAUUUUAFFFFABSikpRQA+iiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiig0CiiigAHWpKYAc0+gAooooAKKKKACiiigAooooYBRRRWYBRRRQAUUUUAFFFOAOaAAA5p1FFABRRRQAUUUUAFFFFABRRRQAdacAQaaODTwQaAFooooAKKKKACiiigEFFFFBpcKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAClBxSUUAPzmlpgOKdnNAC0UUUAFFFFABRRRSYBRRRU2GgooopFXCiiimhhRRRVgFFFFJgFFFFTYApQcUlFIB4INLTF60+gAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAW44UtNBxTs0FhRRRQAUUUUAFFFFABRRRQBiUUUV0HAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABSikpRQA+iiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigu4UoBzSgHNOoC4UUUUDCiiigAooooAKKKKACiiik2AUUUVABRRRQAUUUUAKOtPpg60+gAooooAKKKKACiiigAooooAKKKKACgcGiigB4INLUY4NP3CgBaKTINLQAUUUUAFFFFAIKKKKDS4UUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFKDikooAfnNLTB1p9ABRRRQAUUUUAFFFFJgFFFFTYaCiiihFXCiiiquMKKKKYBRRRSYBRRRUAKDinAg0ylBxQA+ikBBpaACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigApwptKDQNbjqKM5ooKCiiigAooooAKKKKAMSiiiug4AooooAKKKKACiiigAooooAKKKKHoD03Ciiii4BRRRRcAooooAKKKKYBRRRSuAUUUUAFFFFABRRRQAUUUULUHpuFFFFFwCiiigAooooWoBRRRRcAooooY1qFKKMGjBoEPooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAoHWigdaAJKKKKACiiii6NAoooougCiiii6AKKKKVwCiiioAKKKKACiiigAooooAUdafTQDmnUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAADg08EGmUA4NAElFICDS0AFFFFABRRRQCCiiig0uFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFACg4p2c0ylBxQA+ikBzS0AFFFFABRRRQAUUUUmAUUUVNhoKKKKEVcKKKKq4wooooYBRRRU2AKKKKQCr1p9MBxTgQaAFooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKBocKWmg4p2c0FXCiiigAooooAKKKKAMSiiiug4AooooAKKKKACiiigAooooAPrWE3ifw6mQ2oQAqSCN/p1rakIWKRv7ikn8q+NZ5DPM84/jZiBx36VrRgp7mVSbPq/wD4Snw5/wBBC3/77o/4Snw5/wBBC3/77r5KwScDP5Uc/wCcVr7FdzPnPrX/AISnw5/0ELf/AL7qSHxJoFxKsMF9A7uQFCtkknpXyL05OcfhXT+DSR4q0vAxidVxx0IJ/lQ6KSuaKbPq2ijtWfqmp2ukWEmoXhxGgzgckn0ArmSNG7FuaeG3iM9w6xxjqzHAH1Ncxc+OfCto+yS+jYj+5838q+fvEHirVPEFw7XEhSDP7uJfuqPf1rmGPJfOPwHP0ro9lbcz9ofUcHj/AMJT4AvkTP8AfBWuqgura6jEttIsqHuhzXxnkMQFwc+oBrY0jW9T0K4FxpszKSRlScqR6Yp+wXQjnPrvjGetFc/4a8QweItOS8iUI6/LIvof8K6CuZprc3TT2CiiikMKKKjaVEUyOdqqCST0AHWna4N2JD8vXvzXO3/i3w7pj+VeXkauP4Qdx/TNeLeMvHF7rE0thp8phsYzgqv3pD659PavN12qSwIQjq3TP4CtoUrbmUp32Pp6P4h+EpG2i7wfdSK6Wx1XTtSTfp86T8Zwhyfyr48YAncp/IirVnc3NjcLcWbtG687gc8j6VfsV0I5z7JHp0+tFcB4H8ZjxFbmzvQFu4ByP749frXf9s1zNNbm90FQ3Fzb2sXn3MixRjq7HCj8aoa1q9roenSajdnKoOFHVj2FfMWu+JtU1+4eW8lIj3ZjjHCqPcdz71pTpicl3PoS58d+FLVyj3yMR/cy38qSDx74TuCAt8iEnAD5Wvlotj584/Ac/SlyrHClTkZ5ANa+yj3MeZn2bb3EF1GJrZ1kQ91ORU+O/WvkfQ/EOp+HrgTWEpCkjehOVI7jHbNfT/h/XbfX9Mi1GAbc/K69wayqU2ty4VDcooorI1CiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACgdaKB1oAkooooYBRRRWZdwooooGFFFFABRRRQAUUUUAFFFFABRRRQAUo60lKOtAD6KKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAUHBpwINMpV60APooooAKKKKACiiigFuFFFFBoFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFACg4pwOaZSjrQA+iiigAooooAKKKKACiiikwCiiioGtwoooposKKKKq4BRRRQwCiiipsAUoOKSikA8EGlpgOKcCDQAtFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAgooooKbClBpKKBLcfnNFIKWgoKKKKACiiigDEoooroOAKKKKACiiigAooooAKKKKAKOoP5VhO//TN2/AA18dA/KMelfW/iKXyNBvpz0SF/8a+RhwPpXRQ0MKh7n4T8CeHtZ0G11K/hdpZQekjDocfzrov+Fa+FB1ikH1lb/CtTwxLb6X4Qs5rl1ijSEEsxwMsc/wAq5q/+KuiQS+VZ2s9wfUMEH61PNILR7miPhr4TYfLFIfpK3+FW7LwD4d0+7iv7aJxNEQy/vWwCPwrMsvijoN3J5VwstsfV/nH5ivQYLm3uolntpFkR/uspyD9KhxqLc0Si9mTk45rwn4r6u73kOjRMRHEBJIB/eboPevdxyQPWvk/xfc/bPE99cEknzQi/RTg/rVU0KoznFVmIA6kjFfRfhfwFpWm2SXOoRie6lUOxblUz2ArwjQoftWs2MMnR5wD+J4r7AyM7gMHG3HtWlWbJpxRxOteBtC1a2ZBAIZgDsdOMntn8a+aLu1ezu3tZOGQlSPccGvsvOOR25r5e8fQLb+LLoIMZ2tgf7QyaKU2OpA0Phzq0mm+II7Jm/c3mUb2YdK+kRyCfTrXxrZTtb3cMwJHlSKykemRmvsVHEsayHow3fnzU4jcVMlooorA2FHJxXnHxL1iSx0H7JCdr3j7D67R1/OvRxjIzXz98V7rz9ZtbbJHkW+cdsscj9K0pmdQ8tLDgKu3H617H8P8AwTZ3lp/beqr5m5sRxnkYHrXjLYIOelfX3h+H7PollCoAVIl/EkVrVkzKmineeEfDd5EYpLGFeMBkXBHvXzj4o0Obw3rD2KtlCPMjPX5ff3r6x6c14d8XYgl1YXOBueN0P/ATUUpu9i6kEeZ6Hqcmj6vBqELFfLcbiP4lJ+b9K+ukn82FZB/EAw9CCMiviwjIINfWPg66+2+GrJ2JJMe3J9UO2qxC6ipnknxU1d59WTSInPlWwBYDuzc15fFG8jrGgyzEAD3J4rY8RXIvfEF7dkkl5+M/3Rx/SpfCkIuvEWnxSDIMoJH05H8q0grLUiWrPc/DngLR9Js1e8iFxcyKGkZuQuewFT694E0PVrR/ItlguFB8socAnHGfxrvCQSW7njHtTQdp3enNcqmzocEfGM9u1vPJEeqEgj3Bx/OvR/hdq0ljr39ms2Yb1CcHsy9MVznjOBLbxNfwoMfvM49iM1j6PdNaavZ3KkgQzIcj+7uGa7PjTuc603PsQcgt6HFLTOoG76mn1wnUFFFFABRRRQAUUUUwWuwUUUUroGFFFFMAooopXQMKKKKACiiigAooooAKB1opwBzQA6iiihgFFFFZgFFFFBdwooooGFFFFABRRRQAUUUUAFFFFABSjrSUo60APopMiloAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAoBwaKKAHgg0tMBwacCDQAtFFFABRRRQCCiiig0uFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABSg4pKKAHg5paYDinA5oAWiiigAooooAKKKKTAKKKKmw0FFFFIq4UUUU0MKKKKsAooopMAoooqAClXrSUoOKAH0UgINLQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUDQoOKdnNMpwoKuLRRRQAUUUUAYlFFFdBwBRRRQAUUUUAFFFFABQelFFAHK+M5APC2olSP8AUn/x4ACvlQg9Dx9eP519j6jYW2q2bWF+nmRyDDbflOK4G4+FHh2QfumuIif9sOR/31W1OaMpxZ5d4k8RDVbWw0exbzLa0hTeiZ+eUDB/D3rlH07UI1y1vMi+uxs/yr6i0Dwppfh638u3XzJu9wwHmH+ldIy712SfOPcmr+sLsHsz4uBVhyoYZxnPeuz8H+LLjw9fLG5ZrGTHmRHnBPdfStT4jaBaaNqcd3YoEiugTsHQMOteccEY9a1j761Mn7r0PtJGV1WRDkMAwPqDyDXx9qb+Zqt1I3Xz2/8AQzX0x4HuZLzwvYS3Bywj8vPqFOAf0r5t1uFodcvIWGCs5B/FiR+lY0k72LqdyXwx/wAjDp//AF3T+Zr63r5H8MkDxBp5Padf0PNfXGc7v9k80q+uw6aYh6Gvmf4j/wDI3XX+5H/Kvpg9DXzN8RmB8X3Y9Ej/AJUUGi6hwhJAJHavsjTCX0+1kbvCufrtFfHUaGSQRL1YgD8eK+y7BNllbp0xGF/EKB/SjEGdMnooorA3FXlgPevmj4lyFvFEq/3YYgPyr6XHHNfN/wAT4TH4oDkcPbof++Rg1pTM6h5w3Q/SvsjSP+QTa/8AXKP/ANBr43bhT9K+x9JIGk2v/XKP/wBBrau02RTZfPSvFPjB97TvrL/Sva+vFeJfF9wzabjv5p/A4rGl8dzSoeL9eK+mPhvIz+EInPVGlA/77NfM9fUHw9iEPhK0BH+tEkn4M9b4hrlSM6ex81XDb7qV26mQ/wAzW94K/wCRp07/AK6j/wBBNY19C0OpXEDDBSYqfqCTWx4LI/4SfTj283+hFW2mtCD6tpD0pwyQT6HFIelcKOpny549/wCRtv8A/eX+VcgjFWDDqCCK67x4wPi3UMdnUfpXLWsZmuY4V6s6qPxNd1P4TllufZdu26KNz3j/AF4qUnAyaZEuI0HYIP1xTzwM1wnWUJNV0yGVoJbmNXU4ILd6uQyx3ESzQsGRxkEelfO3i7/kZbwscjf9O1e0+Ex/xTdmvTMYOfxoujVwSV2bVze2lmoe6kWMMSoLHHIpltqNhdsUtZ0kKgMdpzwa4D4mqv2C14/5bN/KsT4ZZF9dc5Hkrxj3pXFyK1z2OSWOFDJK6qoGSScAD1NYT+K/DkcnktfRbumASc/kK4j4gWWvz3EZQmW2YYRY+ME/3h3FYcfw712S2E2+GN2XdtO7J9uOKjmkFNRW7Pa7W8tL5d1nKsw/2DnH1qx2BPG7PX29a+YrC9vdEv8AMDFHhlIYA8EoeR9K+lrO6F7Zw3f/AD0jVsH0YZoCpBblM63o6/fu4VxwctjpVm5vrKyh+0XcyRRnozHAIPce1fMl8U+3TBwCTK3XpjNdqvh3xD4ojTU2xDHtCRK78EDvjtS5mNwS1Z6va+I9CvZPKtryJ2zgAHqfatrpnPbH459K+X9R06/0a/NjdnZKBuDJjH1r3PwXq8mr6MrXXzSwfIx/lTFUgtzraKKK0MgooooAKKXaaUA5ougAA5p1FFF0AUUUUMAooorMAooooAKKKKC7hRRRQMKKKKACiiigAooooAKKKKAFHWlyKbRQA/IpaYOtOyKAFooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAClBwaSigB4INLTAcGnAg0ALRRRQAUUUUAgooooNLhRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUoOKSigB4OaWmDrT6ACiiigAooooAKKKKTAKKKKmw0FFFFCKuFFFFVcYUUUUMAoooqbAFFFFIBV60+mA4pwOaAFooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigEFFFFBTYUoNJRQJbj85opBS0FBRRRQBiUUUV0HAFFFFABRRRQAUUUUAFFFFADWdUUySHaqjcWPAA9Sa8w1P4m6XY6rFZWyCWFZP382egP931rY8e6Zqmo6FINNmZfLbLxr/y0T+79BXzPtZcqQRn7ylen1rWnBGU53PsyOeCWJJo5FZJfuMDw30qTI6E4Pp3r5c0Hxrrfh9fJtXDwj/llIMg/Q9q6Kb4sa5JHsht7eFz/AB/Mx/nT9k+we0L/AMW72Friy04HLxKzMB2B6V4+Ac9CcDd+FWby7uL6drq7lM0zty3U4ra8L6Fc+I9US0gyIw2Zj2AHat4NRWpnLXY+hPA1vLbeFLATDBKbvwJJ/lXivxDsTY+KLh8YW4CyofXjn8a+kYIUt4RbQjCIoVBXAfEXw6dY0pb61QtcWbEgDuO9c9Opad2aTTcLHztZTtaXcVyP+WbhvyIJr7DtbyK9tIrqNgUlXeD65/r7V8b7T1IPHUen1rrdB8ba54ei+x2rK8AYsEkGevbNb1Kd9iITtufURljQ4dgCBk57Y9a+UPFWoLqviG8vUy0byhVI7ha2tX+Imv6tbmzBW2hbqIuG+ma4ZcqoIONvI9z71NOlbcJzubnhmw/tLXrGyHIMuT9F5yfavrXjAHQDpXj/AML/AA2bWJ9eu0KySr5UQb+6eSwr2Cs60k3oXBNFO91Gx02MTX8yQoehc4B+lZ+n+JdB1SXyNPvIpZB/CDz+teAfEK9mvPE91FLvZYCscYPC5IznHb61xQmkgfz4iyuhBBJw2RyMkUeydg9ofaHTrx9a8W+LVg2LLVFA2jMLnuD1H4V6xo9015pNtcT/ADGSFST/ALR61S8R6Kmu6PPpzgF5Fyv++v3fzqIPuVNXPkraWHAzX1R4J1RNQ8M2z5y6L5bjuGHTNfMN7a3Fjcy2k42SwnDj39q09E8R6l4dleXTJQqNjMbcqx7mumdNPYxhpufWufmwOSOwr57+KWox3WtxWkTBhaRYfHQMx6fWobv4o+ILmLyoo4YD/wA9FBzXnbyvO7NKd7sxZmPVie5qadK245zuMWNnkWFfvOQB9TxX1/o9qumaRa2gH+phUMB79a8D+Hfhp9Z1Qalcp/odtjDnoXHIA9ea+kQcnd3PGO1Z1pJvQqEWtz5c8d2Lad4ovEwAJWEqe4I/nXN6Ref2fqVtek4EMqscegIJ/SvdviV4dbVNOj1a0QtPaZBA7r3+tfPoQ4yRkA4Pt7GtaWq0M57n2bFdR3EEcysCsihwR3z0/Gh54oss7AbQWOewHNfMOheOtd0C3WyhZJYUJIWQdj2zU2s/EDXtZtzZOy28DdViGD9Ce4rNUWae0MHX75dT1u6vhkpNMcEdcLV/wZp/9o+JLC3IyEfzG+ic81zA3AAg7duQvvn1r374Y+Gm0+0fWbpCstyAiBuoQc5/GtqklT0M15HqoBO3dxjOf6U49DS0nHeuI6j5y8YMB4jux/t/0r2vwqceHbL/AK5f1rx7xvbvb+JbjzRt8w70z3GK9G8Ea7p82gxWbTossK7CGOO/asro6KivCyKfxNP/ABLrU/8ATZv5VhfDFh/aF1/1yX+dWPiFrVleS2+nWcqyMhZmx7j+dHwygbz7u7AzHtWPI/vZpjWkLM9fOW6n6DFIS4IbO4jilpD0rQ5kfNPiYAa7fiMYAmkx9Sa998OsTodnnvAmP++TXgXiP/kO3/8A13f+de9+HP8AkB2H/Xun/oJrM6KmsUkfOOo8384/6aN/Ovp3TkIsLcA/KIlAX6ivmLUP+QhP/wBdG/nX1Bp3/Hjb/wDXNP5UBV1jZHkvxMwNQspAoG6Jtx+la/wxLf2ddo3aRf1GayPib/x92P8A1zetf4Z/8ed7/wBdI/8A0CgJfBY9NooorS6OcKcAc0AHNOpXQBRRRUAFFFFABRRRV3QBRRRUAFFFFABRRRQAUUUUF3CiiigLhRRRQMKKKKACiiigAooooAKUdaSigB+RS0wdadkUALRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFAODRRQA8EGlqMHBp4INAC0UUUAFFFFAIKKKKDS4UUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAKDinA5plKDigB9FIDmloAKKKKACiiigAooopMAoooqBoKKKKaLCiiirAKKKKTAKKKKmwBSr1pKUHFIB9FIDmloAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAW4oOKdnNMpwoLuLRRRQBiUUUV0HAFFFFABRRRQAUUUUAFFFFACkcfK2MdBjg1554n+H2na45vLRvst1/Ey/df6jtXoVFVCbRPIj5pvPht4ptH2QxrcJ6q2T/KqMXgHxdIcfYHX/AHio/rX1JnP3gfwOKOcfMSx+gq/bSI9meCaT8KtTuW36rcJAnpEd3+FezaRothodmLLT0CoOvHzN9WrXDc88Cm1E5tlwgkFKeu5WxjgDHB+tJRUlHmnib4cWOryPfac32a7Y/Nj7j/Udq8zu/ht4qtnKrbCdOxRwTX0vSYz97B+mRWntpGXsz5it/h34vnwfsflf9dGUf1NeieHvhdaWj/adef7S/wDcU/LXrQVB6/iAaKPbSD2Y1URAqIoVFG0KOigdMU49OaKKzNTyDx14Ev8AVNRbWNJIkaZQJIy23BUYGPU1xum/DXxLd3Kx3cQtYQRvZmBYjvx9K+ky3GFG0UKcc8A+4JrT2rtYy5Cta26WcEVtD9yFAi/hU9FFZmpx/ibwbp3iNfOfEN2B8s6joPQjv9a8hv8A4ZeJYJCIFS8iHRlOD+VfR1AxyDnHscVUKskTyI+XIvAHiyV9n2IjHUswA/PNdjovwru5iH1udRFkfuo+Tj/er3IBehGAPTFODZJzwKv20iPZlO0s7ewt47SzURxxjCADhR/Un1q2MZ56UYNGDWRqLtOSQ2O2MZGK8v8AEvw0s9Tke+0lvs1w5y6/wP8A4V6lSVUZuGxM4LofMV38OfFdsxUWnnKOhRgTTbX4c+Lrgg/ZPJ56yMox/Ovp7Yp+9z9MikCAfd/UA1ftpEezPKvDnwvsbFhc6232qXOdg+6DXqwCjAAAGMAD+EDoBTqKiUnPVmkIrqFIenTPtS0VIzmfEXhm28QW6LI2yeL/AFcuOmex9hXlN14D8QwO2xI5GBwrI2Mj3GK97609Tt+6xFTyIuE2tzwzTvh5rk8o+07LaLqzZ3Oa9g0rS7TRrNLGxXbHGOvdyepNaLYY5IJ/GilZinNsKQ0tFWSeM6x4F1++1S6u4BDsmlZ1zJg4JyMjFep6PaTWWmWtpPjfDEqNg5GQMcGtOiosy+dnhl38PvEc13JMgg2u5YZk7E/SvbLNGgtYopPvIig49RU9FFmHOzzvxp4Z1TXp7aTTxGREjK299vJ9OK0PBeg6hoVvcxagEBldWXY27gLg5rtKKLMXOwpwBzTR1qSizJCiiikAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQXcKKKKBhRRRQAUUUUAFFFFABSjrSUUAPyKWmDrTsigBaKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAoBwaKKAHgg0tRg4NPBBoAWiiigAooooBBRRRQaXCiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigBR1p9MBxTgc0ALRRRQAUUUUAFFFFJgFFFFTYaCiiihFXCiiiquMKKKKYBRRRSYBRRRUAKDinA5plKvWgB9FFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABSg0lFA1uPzmikFFBRi0UUV0HAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABS4NJUlADMGjBp9FABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAA61JUY61JQAUUUVmAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUF3CiiigYUUUUAFFFFABRRRQAUo60lFAD8ilpg607IoAWiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKVetJQDg0ASUUgINLQAUUUUAFFFFAIKKKKDS4UUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUoOKSigB4OaWmL1p9ABRRRQAUUUUAFFFFJgFFFFQNBRRRTRYUUUVVwCiiihgFFFFTYApQcUlFIB4OaWmL1p9ABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQNLUcKKBRQUYtFFFdBwBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAU/IplFAElFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUXQBRRRRdAFFFFF0AUUUUAA60/cKZQOtAElFJuFG4VFmAtFFFIAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooLuFFFFAwooooAKKKKACiiigApR1pKKAH5FLTB1p2RQAtFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAoODTgQaZSg4NAD6KQEGloAKKKKACiiigFuFFFFBoFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAKDinA5plKvWgB9FFFABRRRQAUUUUmAUUUVNhoKKKKRVwooopoYUUUVYBRRRSYBRRRUAKDinA5plKvWgB9FFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAtwooooLuOFFAooAxaKKK6DgCiiigAooooAKKKKACiiigAooooC4UUUUBcKKKKAuFFFFAXCiiigAooooAKKKKACiiigAooooAKKKKAH5FLUdSUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRWYBRRRQAUUUUAFFFFaXQBRRRQAUDrRRQA/cKWox1p+4VFmAtFJuFLSAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooGFFFFBVwooooGFFFFABRRRQAUUUUAFKOtJRQA/IpaYOtPoAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigBV60+mA4NOBBoAWiiigAooooBBRRRQXcKKKKBhRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFKDikooAeDmlpgOKcDmgBaKKKACiiigAooopMAoooqbDQUUUUIq4UUUVVxhRRRQwCiiipsAUoOKSikA8HNLTAcU4HNAC0UUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQNbjhRQKKCjFoooroOC4UUUUBcKKKKAuFFFFAXCiiigLhRRRQQFFFFABRRRQAUUUUAFFFFBdwooooAKKKKACiiigLhRRS4NACUuDS4NOoAbg06iigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKizAKKKKLMAooopAFFFFABRRRWl0AUUUUXQBRRRQADrT9wplA61FmBJRSbhS0WYBRRRSAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigYUUUUFXCiiigYUUUUAFFFFABRRRQAo607IplKOtAD6KTIpaACiiigAooooAKKKKACiiigAooooAKKKKACiiigApQcGkooAeCDS0wHBpwINAC0UUUAFFFFAIKKKKC7hRRRQMKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigApy02lBxQA+ikBzS0AFFFFABRRRQAUUUUmAUUUVA0FFFFNFhRRRVgFFFFJgFFFFQAU5abSg4oAfRSA5paACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooGtxwooFFBRi0UUV0HmhRRRQAUUUUAFFFFABRRRQAUUUUAFFFIGznHOP1ouhPQWig5Aye3WigdgooooAKKKKACiiigu4UUUUCuFFFFBIVJUdOyKC7jqKTIpaAuFFFFAXCiiigAooooAKKKSi6EmmLRSE7Rnr9KU8Ej070wugooopDCiiigAooooAKKKKACiiigAoooqLMAooopAFFFFABRRRQAUUUVpdAFFFFAAOtP3CmUUAP3ClqMdafuFRZgLRRRSAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigYUUUUFXCiiigYUUUUAFFFFABRRRQAo607IplKOtAD6KTIpaACiiigAooooAKKKKACiiigAooooAKKKKACiiigApV60lKDg0APopAQaWgAooooAKKKKBrcKKKKCwooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigBy06mA4pwOaAFooooAKKKKACiiikwCiiipsNBRRRQirhRRRVXGFFFFDAKKKKmwBRRRSActOpq06gAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAW44UUgooLuY1FFFdB5oUUUUAFFFFABRRRQAUUUUAFFFHoByW7d/1ougWuwfWvC/E3xG1/SNdvLC2W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pQ8e+NtW8M6jBZackRSaDzG3rnncV6/hW14C8QXnijTZ7y+SNHimMQ2jttB9fevNPi//AMh6z/69P/ajV13wf/5Ad5/19n/0WtASgvZnaeMdau9D0GfUrFVMqFAu7BGCcc15/wCDfH2s6/rkOmXaRLEUdv3a7TkfjXU/Ej/kUrj6x/8AoVePfC848XW5HaOSgIQVrn0F4i1228O6ZLqNyC2w7VUfxORkD6eprwW7+KviiabzLVo7aLrt27/wNenfE/S7q+8NpJYgv9lmEjBeSRtwSB3xXmnw51bwvpVxd/20kYkkICSSAlc+hHY0PQIQVj13wdrep654dOrap5ZlO7yyEwPlHevLIfip4jluI4SsOzzChwuDtY4H417nbXenXulPLpjpJbhW27BhQcHp9K+QLP8A4+4/+uy/zNAQgrH2VezzW2nXF7HgypCXUHpuUZ5rwjTfij4hub+G0kjgCTTIpIXnazYOPevc9V/5At1/1wf+VfImi/8AIYtP+u8f/odD8wpwTWh9kkbgEY4P584zXjXjD4h674f8Qz6dZpC8aBCNy4PIzXs5++Pr/Svl74m/8jjc/wC7H/6DQTT7HvXg7WLrXfD9vqV0oWR2cNj2JH9K6muB+Gf/ACJ1t/vyf+htXfUGc97BRRRWgBRRRQAUUUUAFFFFRZgFFFFIAooooAKKKK0ugCiiii6AKB1oooAfuBpajHWn7gaizAWiiikAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAwooooKuFFFFAwooooAKKKKAClFJRQA/OaWmCnA5oAWiiigAooooAKKKKACiiigAooooAKKKKACiiigApQcGkooAeCDS0wHBpwINAC0UUUAFFFFAIKKKKC7hRRRQMKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigBy06mA4pwOaAFooooAKKKKACiiikwCiiipsNBRRRQirhRRRVXGFFFFDAKKKKmwBSg4pKKQDwc0tNWnUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFACiigUUwMaiiitzgCiiigAooooAKKKKACiiigDwD4v8A/IXsv+uD1yGna9e2fh2TQtKZhPc3O92QEsE245x711/xeOdYsR/0wPXjqa2vhDaQnT729Cr5zSeUW4JwPSoszr+xYj+E1jqFtdXst7BO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TVp1ABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFACiigUUwMaiiitzgCiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACnbhTaKAH5FLUdPyKCroWiiigdwooooAKKKKACiiigAooooAKKKKACiiigAoooqLMAooooswCiiikAUUUUAFFFFaXQBQODRRQwHhgaWoxwaeGBqLMBaKKKQBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUDQUUUUFhRRRQAUUUUAFFFFACjrT6YKcDmgBaKKKACiiigAooooAKKKKACiiigAooooAKKKKAClBwaSigB4INLTF60+gAooooAKKKKBrcKKKKCwooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACnLTaUHFAD6KQHNLQAUUUUAFFFFABRRRSY0FFFFQVcKKKKaGFFFFWAUUUUmAUUUVACg4pwOaZTloAdRRRQAUUUUAFFFFABRRRQAUUUUAFFFFACiigUUwMaiiitzgCiiigAooooAKKKKACiiigAooJ4yBn2qCa5htkMs7CNF+8zEAD601rsD03J6K5w+L/DAODqEP5mta01LT75Fksp0mVuhQ5/lQ01uTzIu0UAgnANFJ6blBRRRQAUUDtngt2NV57y1tYzLdSrEo/vnFNa7A9NyxRXODxf4YL7P7QhyPc1s2t9ZXqCSzmSZW5BQ5/lQ01uTzItUU0MCcCnUm7blBRTHljjUvIwVV6seAKwZvFnhu3fy5r+JT9aaTexPMjoaKzbPWdJ1DiyuopW/uqwzWix24B6ntQ1bcdxaKKKQwooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigB+RS1HTtwoAdRTc06gu4UUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFRZgFFFFIAooooAKKKK0ugCgcGiihgPDA0tRjg08MDUWYC0UUUgCiiigAooooAKKKKACiiigAooooAKKKKACiiigYUUUUFXCiiigYUUUUAFFFFABSjikooAeDmlpg60+gAooooAKKKKACiiigAooooAKKKKACiiigAooooAUHBpwINMpV60APooooAKKKKAW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cwUUUUAFFFFABRRRQAUUUUAFFFFABS0lFADs06o6fmgq6FooooC4UUUUDCiiigAooooAKKKKACiiigAoooqLMAooopAFFFFABRRRWl0AUUUUMAHBp4INMpR1qLMB9FFFIAooooAKKKKACiiigAooooAKKKKACiiigAooooGFFFFBVwooooGFFFFABRRRQAUo60lKDigB9FIDmloAKKKKACiiigAooooAKKKKACiiigAooooAKUHFJRQA8EGlpi9afQAUUUUAFFFFA1uFFFFBYUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABTlptOWgB1FFFABRRRQAUUUUmAUUUVNhoKKKKEVcKKKKq4wooooYBRRRU2AKUHFJRSAeDmlpq06gAooooAKKKKACiiigAooooAUUUCimBhUUUVucAUUUUAFFFFABRRRQA/cKTcKbRQA7cKYfmHzAZ9ev6GlooAp/2dpxbc1vH+CircarEhSPC+m1QKWiq5mRZi7TT+QQy9e+aTcKTcKksrtZ20knmyRoX/vbBmrWOeTnHQUm4UtLmZFmFRlAy7HVWX/a5/Q1JRTLKA0zTw277PF/3yKsrCsaFY1Cem3ipqKrmZFmN2mndaKKksQg427jt9DVWSws5m3SwRN/vIGq3RS5mBDHBHCNsKIF/ugAY/EVKNwPXGO/U0tFHMwCiiimAUuRtwOD9aSigAxnhicDpg0UUUAFFFFACEbuDtx7rmjaMbTjHsMUtFADdo6FQwHTtQFAO5flP1zTqKAEIU53Ln3zjNGMABSQPQ4I/lS0UAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQA/NLUdPzQAtFFFBVwooooGFFFFABRRRQAUUUUAFFFFDAKKKKizAKKKKQBRRRQgCiiitLoAoHBoooYDwwNLUYODTwwNRZgLRRRSAKKKKACiiigAooooAKKKKACiiigAooooAKKKKBoKKKKCrhRRRQMKKKKACiiigBR1p9MBxTgc0ALRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAoOKcCDTKVetAD6KKKACiiigFuFFFFBd0FFFFAwooooAKKKKACiiigAooooAKKKKACiiigAooooAKUHFJRQA8HNLTVp1ABRRRQAUUUUAFFFFJgFFFFQNbhRRRTRYUUUVYBRRRSYBRRRUAOWnU1adQAUUUUAFFFFABRRRQAUUUUAKKKBRTAwqKKK3OAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACn7hTKKAHbhS7hTKKAJKKbuFG4UAOopNwpaACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAH5pajp+aAFooooLCiiigAooooAKKKKACiiigAooooYBRRRUWYBRRRSAKKKKEAUUUVpdAFAODRRQwHhgaWmL1p9RZgFFFFIAooooAKKKKACiiigAooooAKKKKACiiigAooooGgooooLCiiigAooooAKUdaSlBxQA+ikBzS0AFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFKDikooAeCDS0xetPoAKKKKACiiiga3CiiigsKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigBQcU4HNMpy0AOooooAKKKKACiiikwCiiioGtwoooposKKKKsAooopMAoooqAFBxTgc0ynLQA6iiigAooooAKKKKACiiigBRRQKKYGFRRRW5590FFFFAXQUUUUBdBRRRQF0FFFFAXQUUUUDCiiigAooooAKKKKACiiigAooooAKfuFMooAfuFLUdO3CgB1FJuFLQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFLSUUAPzS0wdafQXcKKKKACiiigAooooAKKKKACiiigAooooYBRRRUWYBRRRSAKKKKEAUUUVpdAAODTwwNMoBwaGBJRTdwo3CswHUUgbNLQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQNBRRRQWFFFFABRRRQAUUUUAKOtPpgOKcDmgBaKKKACiiigAooooAKKKKACiiigAooooAKKKKAFBxTgQaZSr1oAfRRRQAUUUUAtwooooLugooooGFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABTlptKDigB9FIDmloAKKKKACiiikwCiiipsNBRRRQirhRRRVXGFFFFDAKKKKmwBTlptKDikA+ikBzS0AFFFFABRRRQAUUUUAKKKBRTAwqKKK3PMCiiigAooooAKKKKACiiigAooooLugooooC6CiiigLoKKKKAugooooGFFFFABRRRQAUUUUAFO3Cm0UAP3ClqOnbhQA6ik3CloAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAUU7NMpRQA+ikzS0F3CiiigAooooAKKKKACiiigAooooAKKKKGAUUUAZOKzAKKlELnpQYXAzQXyS7EVFO2OO1Nww7GhNE8rCikyfQ0gOT0q+ZBysdRQeKbuFTYLMcDg07cKj3ClyKOViuP3CjcKbRRysLjtwo3Cm0UWYXHbhSg5plAODSAkopu4UbhQA6im7hRuFADqKQHNLQOwUUUUFhRRRQAUUUUAFFFFABSjrSUoOKAH0UgOaWgAooooAKKKKACiiigAooooAKKKKACiiigApV60lKDigB9FIDmloAKKKKACiiiga3CiiigsKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigBy06mrTqACiiigAooooAKKKKTAKKKKga3CiiimiwoooqwCiiikwCiiioActOpgOKcDmgBaKKKACiiigAooooAUUUCimBhUUUVueYFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUF3QUUUUBdBRRRQF0FFFFAwooooAKKKKACnbhTaKAH7hS1HT9woAWiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAFFOzTKUUAPopM0tBdwooooAKKKKACiiigAooooAKKKKGAUUUlRZgTxzspAbpV5WDjIrIJBGKnhnZCB2qWjaFTozRxQVBpFcMMin1JuknsRmNT2pDEh4xUtFK4cqIDbxmm/ZkqzRRzMORFU2ykYFR/ZPTrV6iq52T7KJQNs/qKT7PIK0KOafOw9lEzTC4ppicdq0+aWjnYeyRklXAzg03n0Na+KQqD2pcwvZmTk+lAOa1DGhGCKb5EfpTuT7IzjxSAg1om3jIpptk7UXQvZFJetPJxVj7Njp1pDbuR1FF0Lkl2K24UbhUxtnAzkUwwuKLoOSXYaDmlo8txyRS4I6g0XQcj7CUUE4pu4d6Lk2HUUm4HpzS0xXCiiigYo60+mA4pwOaAFooooAKKKKACiiigAooooAKKKKACiiigAooooAVetPpgOKcDmgLC0UUUCCiiigaYUUUUF8yCiiigOZBRRRQF0FFFFAwooooAKKKKACiiigAooooAKKKKAFBxTgc0ynLQA6iiigAooooAKKKKTAKKKKga3CiiimiwoooqwCiiikwCiiioAKctNpy0AOooooAKKKKACiiigBRRQKKYGFRRRW55gUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQXdBRRRQMKKKKACiiigAooooAfuFLUdO3CgB1FJuFLQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFACjrT6YKdmgBaKKKC7hRRRQAUUUUAFFFFABRRRRdCugpD0peacF5pNoqzIaKs7BSbBUD5WRLJIpGK0omDDJ61RKigF15FJm8Hbc1KKqLcoq/N1oN5EPWoNk77FuiqRvoh2NNbUIsfKDmgC/mk3D1rM/tEf3TSHUOOFoA1Nw7UZPpWT/aD/3QKQ38mOKANfPtQDntWIb6bFMN7MRjNAG/nFJuFc+buYjrTftMx6mgDosj1oLADJNc208jDGSPpUe9/wC8TQB03mx/3hSebGP4hXNbmPU0h5HU0AdL58Q/iFNNxEO9c2B6k0uM9DigDoTdRAU03sQGea50sVPJzTh8w70Abv2+H0NNa/ixwPzrIWB2OAD9auJp7H5mIxQFiY34/uikN47LxHmp0s4FHzDNWlSNB8opoTiluUkaZ+iBacysv3utTvOF4XrVQyFuW5NWRKcbaC0UgOaWgwClHWkpQcUAPopAc0tABRRRQAUUUUAFFFFABRRRQAUUUlAC0AZpQCTinqhzSuilFjCMUL1qcJ604KBRcuxDQOanwKUACi4+REIUml2GpuO1JSuHIiLYaNhqWii4/ZkRUimkEVOKDg0XD2ZXpCcVYwKQqDTuLksQA5pakKelNZStF0Kw2ikJxQDmi4haKKKYBRRRQAUUUUAFFFFABTlptKDigB9FIDmloAKKKKACiiikwCiiipsNBRRRQirhRRRVXGFFFFDAKKKKmwBTlptKDikA+ikBzS0AFFFFABRRRQAoopM44opgYdFFFbnmBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQXdBRRRQF0FFFFAwooooAKKKKACn7hTKKAJKKTdS0AFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABSikooAfmlpg60+gAooooLCiiigApCcDNBOBmnKN1K6AFBanhCDk1IoGMU7bUF+zGgDNLkdqCMUlBSTCkPSlooLGYpOe1PPSmUMChMDu5NQAepqeY/NUVZm0FoIQCKQKAc06igYUUUUAFFFFACHmkxTqKAG4pMYp3SkJGKAG0U7a2MqM1PHbytztoArfXijj1rUWwZhlyMVaWzgUYxQBiBHbopqzHZSydcAe9bKoqDCin5bvjFAGamnqD8xBq0trCvarFFAxAoAwBRzjFKDTWcIN1MTdgYgDniqEs/wDCDUc1wznZ2qEKCeOtNIwnO477x4609VbPJpqqQc1KvWqMxQMGnUUUAFFFFACjrT6YOtPoAKKKKACiiigAoopCcUALQBmnKhbpUqRlTk0ropRdyIIScVIsZB5qXApaTZoopDVUA5p9JRUlBRRRQAUUUUAFFFFALcKKKKCwooooAKKKKAYUYzRRQTYaVBprLjpUlFNA4EBBFJVjApjDjiquRysiopSMUlFxWCiiimIKKKKACiiigBy06mrTqACiiigAooooAKKKKTAKKKKga3CiiimiwoooqwCiiikwCiiioActOpq06gAooooAKKKKAGt1FFDdRRQBi0UUV0HmBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFBd0FFFFAwooooAKKKKACn7qZRQBJRTdwpdwoAWiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigBRTs0ylFAD6KTNLQVdBSUoGTipAhzSui0mxFQtU4UAYpwAAxRUGnIhoXBp1FFBQHkUwjFPpD0oAZRRRQAh6UynnpTKGBnScvmmUrH5iKSszojsFFFN3YODQIdRSZHY08Ru3RTQA2kJAGauJZTOM8D61bjsUH3+TQBkDJ6CpBFIeimtxbaJe1ThQowKAMVLGRxyQPrVtLKMDD81fxSYoAiWKNBhRUgLD0pcUlAC54pKKKACiiigApD0pagluFiO3qadgbtuLNIsS+prNklZ+h4pjMztuJoAycU0jCc7jVUk1KqkHNCqQc1IBmqMwAzTgMGgDBp1ABRRRQAUUUUAKDinA5plKOtAD6KKKBBRSE4qSNcmgqOrGAE09UJOKnCgUoAFJs15UAAApaKKgoKKKKACiiigAooooAKKKKACiiigEFFFFBdwooooAKKKKACiiigAooooAKKKKAYjDcOKiZStTUYzTRDTK9FSsOOKjIIqrkWEooopiCiiigBQcU4HNMpy0AOooooAKKKKACiiikwCiiioGtwoooposKKKKsAooopMAoooqAFBxTgc0ynLQA6iiigAooooAQjJFFLRQBg7qN1NoroPMHbqN1NooAduo3U2igB26jdTaKAJKKbuo3UAOopu6jdQA6iiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooLugooooC4UUUUDCiiigAooooAfupajp+6gBaKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAFHWnUzpUyDcRSuhpNiopJBqcLg04AAYoqDeGgUUUUFhRRRQAUh6UtIelADKKKKAEPSoycDNSHpUMh2oTQxrczjyxNA5o7n3pQQDk1mbCEAAknpVu1t1uo94Ix6ViXM2Ts55q1pd2Yrjy2+63T2oA6FLKFR0qyqKgwopwIOCO9LQAnNAFLRQAUUUUAFFFFACHpTacelNoAKKKKACkJAGSaCyqMk1mzyknahppEuaRPLchQQvWqGS5y3WmgMTyakVeas5pVLsQKTTlUg5pwGKcBmgAAzTgMGgDBp1ABRRRQAUUUUAFFFFABSg4pKKBDwc0E4poOKlRd1A0rgi7qsqABQAAMUtS2bxikgoooqRhRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQC3CiiigsKKKKACiiigAooooAKKKKACiiigGFNcZHFOopohohKkU2pmXNRlStUmRYbRRRTEFOWm05aAHUUUUAFFFFABRRRSYBRRRU2GtwooooRYUUUVYBRRRSYBRRRUAFOWm05aAHUUUUAFFFFABRRRQBz9FFFdB5gUUUUAFFFFABRRRQAUUUUAFFFFADt1G6m0UAOzTqjp26gB1FN3UbqAHUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUF3QUUUUDCiiigAooooAfupajHWpKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKMZ4oYWY9VzVpcDio40OOtSheazNYK246iiigoKKKKC7oKKKKBhSHpS0h6UAMooooATrVec4jIqwTgZqrcHCkdc0Ma3KVVriURjYOSe9LPKqLjuazclzyc5qDcR2ydxoR8MCDg1KLeaThR171fh09RzL1osyeZHQWFwJ4RwcrV6si12wsFXpWtnp70WBST2FooopDCiiigAooo6UAIelNpSaSgAprMqjJpkkiqvvVB2djweKdhcyHSyFztHFQCLJ96mUZ4708IQaswkrshWI56in+WRzmpgvNPC+tAuRFYgihetWioNMZcDNArMjooooCzCiiigQUUUUAFFFFABRRT413Nigai2ORCTVpQAKAABS1LZtGKSCiiipGFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQC3CiiigsKKKKACiiigAooooAKKKKACiiigAooooBhTWGRTqKaIaISpFNqZlyKh74qkyLBSg4pKKYh4OaWmrTqACiiigAooooAKKKKTGgoooqUVcKKKKq4wooooYBRRRU2AKUHFJRSAeDmlpq06gAooooAKKKKAOfoooroPMCiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAduo3UzOD8wI/CloG01uSUU3dRuoEOooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigu6CiiigYUUUUAFP3UyigCSik3UtABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQOwU9FJOaZjPFWY0OKV0WtSRTjin00LzTqgsKKKKACiiigEFFFFBoFIelLSHpQAyiiigBrdKp3hKAYGT7VepCFb7wzQBz6WbzvmTpV9LGGMZxkitDAA4pp5FRYrmZGuEGFFNIJOak2mjaaslkeD2rSgfcvPaqOMc0sLmNsnkUmODtualFIGBGaWoNk7hRRRQAUh6UuaQkYoFcbVeWbaNq9abNNlSi9fWqgDE8mqsJyVheWbJNOCk0qrzTwMVRkNVSDmpAM0AZpwGDQAAYNOoooAKKKKAGlQRxTCpAzUtIV3DFAMgopxXFNoM2FFFFABSjmkpR1oAUIWOBVpAAMd6YinOalAwaTNoKy1HUUUVBQUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQC3CiiigsKKKKACiiigAooooAKKKKACiiigAooooAKKKKAYVEUOc5qWimiGiv3xRUmw5zmjYfWquiLMatOoCEd6dtNFw5WNop200hBFFwsxKKKKLhZhRRRQwsFFFFQC3CiiimiwoooqwCiiikwCiiioActOpq06gAooooAKKKKAOfoooroPMCiiigAooooAKKKKACiiigApxRgfmGB60xiQpI6gH+VedaL4tZbm4j1eRio6KqkmldDUW9j0aiuaTxbosjrGkjFn6DYa6WmOUWtwooooJCmT5SCRwcEIxB+gNY174i0rT5zb3MjKw/wBk1QufFukPAyxSkPtOMrxnHFK6NPZy7FPwprd/qt3Ml3IXjES4xzz6129cJ4S1W9v7q4iuGU7Y1P7tVT8OK7umFVNOzCiiigzHbqdUTtHGvmSOET1Y4qOO7tpjiGVHPoGFK6Gotq5ZopDkdRx60vPYZpiCioTc2wLBpUG372WAxUuQQrKcq3INK6DpcWiiimDdtwooooBa7BRRRQAUUUUxXCiiikVZhRRSL82fagQtFQyTwwjMzomf7zAU6KWGf/UOsmOu05pXRpbS5JRQBkE+gyarC7tiuTKg5xgsKLoEr7FmioPtNsOTLHj/AHhU+QSNpyCMg+tMGrbhT91RgqVL5G0dSeAKYJY8gFlBJ4G4En8jSugLFFIGzxRuUqXyNo6k8AUXQlrsLRUQnj6FgCeg3Ak/kak3UxtW3FooJwMmopZ4oeJmCHGeT2pXQLXYloqvHdW83+okV/oRU7Ha21uOM5ouhyTjuLRUUk8MJxM4Tvye1JHc283+okR8dcEUXQKLauiaig5D7SO2c0UXQvIKKKY8iRqHlIVTwCe5piWuw+imLLC7bIpFcgZOD0qQAk4FK6G9NxKKaXRRuZgF6ZPAoikhmYLDIrnGTg9KLoai2OpePWpRFzyaeI0zxRdFtMhVTuFWlOOKNlKF5qBwVtx1FFFBQUUUUAFFFFABRRRQXdBSHpS0HkUDI6KUjFJQAUUUUAIaTFOooAbikxT6Q80AMIyKZjHNSYoHvQBPbucbTVrPNZytscN2q+pzzUWNIMfSHpQTimPIqKSfypFsUsAMmqMsxLYFNeQufl4qPYc5JqrGbYvJNOC80BeaeBiqIADFOAzQBmnAYNAABg06iigAooooAKKKKAClBxSUUANZc1EeDirAGagYYOaDOSdxtFFFACE4qWJctTAMnFW0UAZoZUU73JOAMUUUVBsFFFFIAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigEFFFFBdwooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAYUUUUEpBRRRQUwooooJsxCM9KbtNPophZkZGKKeRmmkEU76CadhKKKKklIKKKKaKCiiirAKKKKTAKKKKgBy06mA4pwOaAFooooAKKKKAOfoooroPMCiiigAooooAKKKKACiiigDI1LW9P0rat0/wA7dFHWuVS98Katdlp4Cjt3OVH6VpTeEYZ9Ua+knd45fvo/J/4CegrE8X6Fp9lbxXdnHsAO0jccmoaa3Oqm4Xtc6j+wNDtAL0oRsUvndnIXuKm03xFpurTmCyLswXdypFZOh3T3PhuZZm3NDG0QJ9MZFc94K51Z3I3MYu5PbiquhKDd+boeg6nq1lpMSS3rFQ5wABmpNN1K01W3+02bFl3beRiuX8bbjp1vtYqTIRx9D/hVzwbn+xgCxY+cV59iKLon2f7vm6lGafwxqmp+RLEzTFipJDYHJFZc0vguCcxGB2VfvOuSBWDFbi98QfZWJUPcEAgnONx612WueGtLt9Ie5tIvKaAZ4JwwzjmoWux0N2tfqbejWmixxNqGkopEn8XtVy51bTrOXybmYI+AcH3rjfAcsga5hONq8gdqs6/4Xu9T1E3luVwVC4OePerujCdFe099ncB0YBkbcpAII96VmKKWUZKgkD3FV7O2FrbRWxO7y0Ck+pHerNM5zySS31nxBqksEwdcnIV+APyqrq+iXegyxyiQEy/ddcgZHbmvZDtQmWTacclm4P515X4l1hNYvY7KxBkhgJ2kd2b0HcVmdtGpzOyWh3Ph3UZNS0qOaY5dDtb3PrUmu6r/AGXpxnY7WPCjufejRLD+ydLSKTAkX55PYd/yrhrl5vFWtmAbvs8J4OCRt9ePWrujKKTn5HKXInB827yzTfNyTk+4r3ey5s4FPQxIPxxnNeW+M4kj1VIkUqiQhVGMYr1GybFnb8HHlqentioNK04yhtYsk8ZrnpvFGl2959hlLiTeI8bT1NdF0rx3VefE5OMn7SmSSfatG0tzGjBS0keuySrFC87/AHUUsfoBk1j6d4h03VLj7PaF2bYX+6egOK0L7jTrg/8ATKT/ANBNeZ+ByRq5YcEwP0J9RSugpU7q6PWaZu4Bx1zn2x60/ntWH4juns9GuJYuG4XP160zJIgvvFWjWD+S8pklHVIxu/XpUNp4w0a6fy2domPHzDP6iuO8L6bZagZr3UH/AHcRAJPGfrj9au+JtG0lNOXUNKQRjftJUnB5I/pWXMzrdFJ2PS0ZXUMhyDyPenVyPgy9lutMdZjlon8sH2rrq1ujnl7rtIKoarcTW2nyXFqhaQAhQOobsav5xyKXP93vQQkeP2Oi6n4ikkupsDbwzS5AB9gOao3Ed94ZvjHG+1kIY7c4YfjXslzcW1pC11dFVUdW6D8vWvJb6S48U60fIU4k+Vc87VHc4rM66c76W0PTZbn7Toz3jZDNbuSB/u15ZpeiXWsecYZFURnJ3bsfpXqGpRC30WaCIE/uCox7D+tedaLPqsNpNaaVDIbiQgk4xgfjQKmo2LY8EaiMN5sS56EbjXpttE8dvFE/WNNp+vrXlK674j0i5Md3I57srAYx3xXdavqwl8PSahbkgzqoUA8gtxV3RE4vRnIeItcl1e5bS9PyYVOw7f42PGPpXS6B4Ys9M23FwFa4Izk5wv69a4HQ3vY5ZG02PzJ1zjA+7nqTnv6VqN4g8S6Xdf8AEwc89mUcj0qDXkTbitz1bgqpzjfnFeWeItbm1e6OlWGTArbPl/jY8flXX6rqu7w4dQteGuVCgd13ccV51oL3sckh02PzLgZAIH3c9Sc96CadLQ73QfDNnpe24nVXuSM/7v09664KTyvPrXlH/CQeJdLvB/aDE57EDkelem2N3He2kN4nSQZOOn0q7ozqLS7JbqdLa2kuJPuxqWP4V47D/aHifUziQgY3A/whc9DivSvEbsmgXO3rgj864vwGv+l3Ck/8sgRj2ODUPTcqEfd5kZeq6Pe+G5EuI5A248MuQM/jXo+m6wt1oh1G448lTu/AVk+Nxv0mFOMbxj1zjNc3pk0i+ENRXr86gfRv8KDSK54JvcoQDUPFGpnEpAI3AnO0KexAo1TR77w3LFPHKGEnIZc7cjtzWv4EQfarkZwfKUj8Ditvx182kxoOF8wFfUcZoLc7S5Ejf0TUTqemw3L8NyDn1WpodZ0u4uDa29wruOMD19K5fwTuk0ydCcASHHsGFU9N8H31lqsN5IUEcT78jqxoMmoqTueihjn5Rk15n4q1KXUrv+x9PbKxHcSOnHXpXWeItWXSLB5EP71vkTHYnvXPeG9HNrpdxqk6sZpVbBI6KRzVtoKcFa5k+BiX1aRiDzFuwTXrgjZh6V5L4BYvrDtg/wCqK4x3BrvvE2tLo2nyS7wsjfIo7jPfFZ3KnT1ON8XanNf3S6RpxysJLHb32j5s/SoPAOTq0mcn9xnk9M1reG9HNro91q1yP30yMELddmP5nvWP8PnB1OTrlogCce9FzSMUkev0DrRg/lQOtMysySikDUtAgooooAKKKKACiiigAooooBBRRRQaAeRTCMU+g8igCOinbTRtNADaKdtNG00ANop200hGKAGnmm4p9IeaAGGrEDkHaahxQMryKALkkm1TjrVBgzNuJpSWJyTSjmpsVzMaF5p+KAMU4c1RI0DFOAzTtpoAwaAADBp1FFABRRRQAUUUUAFFFFABRRRQAoOKjccZp9IwyMUCexBRSng4oAzxQQSRoSc1aAwKjjXHNS1LZtFWWoUUUVJQUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFALcKKKKCwooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAClFJSigGIwzTCMVKRTWXNBFiOiiimgCiiirAKKKKTAKKKKgApy02nLQA6iiigAooooA5+iiiug8wKKKKACiiigAooooAKKKKAK91d21mAbiRU3fd3nbn864jxjqNncWkVnbypI7Et8jAgfjXQa1oFvrWwzOUMf3cc/zrJtPBVrFLvu5TMg6JgKB+VE9TeDilzX1HeG7Rx4em3KVM4ZlB6/dxXKeFL63sdTDXRKK6Fdx7EnvXrsaJEipGMBOFHtXH6h4Nsry4a4ildC/UdqizClVup8/UzvF+oWdxbw20EgkaNyzFPmAyCMcfWtPwXNE2mCJWXzBOWKk4IXI5qzD4V06KxksmBcyHJc8H9Kn0Xw9Bo08lwk8krOAvzBegosynOPJa557puB4oickbTdsM59GNeka8wbQrnHQxk5/Gs638K29vqS6mJ3LiTzMYGOpNb99aJf2UllIdqupUFe2TRBNbhOpFuFnscH4G/wCPm6UY5UHJOK6fUvE2n6Xc/ZbouXwD8nTml0bw/Fo8kjpO8vmDHzgVV1bwxDqV8bl5GiJXaQOaLMXPGc7tnSxSpNGssZyHUMvuDT2ZVUs3QDJ+gqO1txBEkEPSJAqn2FPdd6MnTcCPzqzn23PKdc8SyaxL9jtnEduG2+/447Vu6GfDWlL5z3ULznjdk4Uf7PH86cfA1sR/x9yg7s8BaD4FtS2TeS4/3FqLM7Zzpv4Wbl3cQ6xpNymmyCdm+UFDzj3rg7TSvFlirC1DQ7+G2EfdHSvQNH0WHRYnSGV5S5zlwB/KtstnjAoszCE1T21PDdXTU47tk1V2aUqu3cQT9eK9E8LrqoSUag7+XtHl+4qxrHhuHV7r7XLcSI4UKMAYx710MSCGBLdeVjUKM+1FmaVK0XCyRLmvHNZkWHxG8jEHbOrEA56c17DXLXXhW3utROomeRHMgkwoGOOxomm9jOjJKTbLN9rulNpcjxzq/mxsoAPO4gjFcb4HjdtSaVRlVjZSenJwf6V0N74Ltbi4MsM7RqTu24GAT1xW9pei2ekQ+XbAsT94v60WZanGMLJlLUfFOm6Zd/ZJ9xIAJK9Bmo/FMqSaDLLEdytscHHFRan4Ug1G9a7EpQOBuUdyK2tQ0xL/AE7+zfMaNMKAR221Yk4dzy/RdJu9Zt3tIJliiQ7pTzyT9K6XWdMk0jwr9jLCQRvu3DgHJJ7/AFro9H0SPR1ljimeRZcZ3gZyPpVjVtMj1ezazmdkDEEEdsVFmEqvv3OX8Cf8eU4yOZs9fatQ+LNKF99gJfduC7scZPFXdG0WLR43jSZ5Q77vmA6YxWQfCEB1BrwTuqM4coO5BzRZjbhKTbZ2PUgDv0NVL68jsbOS8l+4i549TxirQwNoHRelUtRsU1GxksJGKrIByOoIOc1Zimjya81KfxDfYu5lggU4UHO3HqcV3el3XhfSLcw2l3EGbG5yfzxx0qg3gS1bAN5Ngey0n/CB2Z4N5KQevyLUWZ0ynCXWx1cmqWMdkdR3CSJV35T5uKw9I8VaffB2uZBaSb/lzgZX+lbdjp0dlpo0wsZowoXDAAHnviuau/A9jLIz2szwK/JXAYZ9Oe1FmYwkluYPjDUrG9nghtnEphB3snO7PvWq2nTnwSkLA70G8464HPArSsvB9jbSrLPK8+3orAAAj6V1e3KBHwQOAB0x6UWZpOacVY828F6haWb3MF3IkMkhDIWOOB6nsTUPjPULK8lt4bWRZXiB8wqcg59DW7qHgqxupmuLRvszP1QDchPrzzmpLDwZp9rMJ7ljcMOgIAAPtiizK543crmfJYTt4ISPa25MNx1CrznFVvBV/aWjXFvdSJFLIwZCxxwOxPTJr0nAKhSBtA27e2K4u/8ABVjczNcWjfZi/wB5ANyE+vPOaLMUKtlqYfjS+trqaBLRhK8KneVIOc+hrs/DdtLbaJbwSjafvFT1GelZ2n+DNMtphNcu9ww6A4AB9sV2OMkN3xg/h0osyakk42RQ1e1e8025t4h80iEKPfFeWeFNQt9N1JvtbeWrxlCx4CkNnmvZK5LVfCFjqU/2qOR4JDxhcFMH2omm9h0pJXjLY5/xfq1ndQw2tpIJdrBiyHIwBiruh6ZPJ4TmtyuJZ97KD+lWbPwPZwzebdTNMB0AAUfjXaqixjbGMKMBR6CizCU0opRPHvCeo22m6i/2tvLV4yhY8BSrZ5rU8Y6rZ3cMNraSCXDBiy8jGMV0Oq+D7HUZzdRyPBIeCFwUwfaq1n4Is4JvNupmmA6AAKPxosy1ONuZvUs+EYDbaMZJ12+ad+COgHrU1n4r0m4vBZRbizHaCRxn0FdIY0VDGnyrjAx1ArlNP8H29rqCXomZgjbwp9feizM4tSbbM7xZo2rajqEbWSFkVc8f3vXms2S18ciAq5kCKhBG5cbR+FeuKMDkmmTRCaF4ScB1K569fakaQdopHguj/wBsyzkaIzmUruPQcZ5xXZ+J9G1jUbi3eCIyskY3NnBD+o7Vv6F4Tt9CuftMFw7/ACbMFVGB3wRXWn+6Cdo/zzUWZpOSbPIZbbx0sDrI03lhCCC64249MVz+inWXmP8AYjMZNis3T7oPP1r3ieITwvCTgOpXPXqMdK5jQ/CcGhXH2iC4d/k2YZVGB3wRRYOZHURCURxicAybBuJ9alpMHOQcnvmlqyG0A61JUY608NmggWiiigAooooAKKKKACiiigEFFFFBd0FFFFAwooooAKKKKACgjIoooAYRikqQjIpu2gBtIeacRikoAbigDFOoHNAAOacBg0AYNOoAKKKKACiiigAooooAKKKKACiiigAooooHYKKKKV0IZ5ZJ605YjnrT15NSAYNDZSghQMDFLRRUFWCiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAW4UUUUFhRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFKKSlFADqDRRQDIWGKSpGGaYRimiLCUUUVVwCiiihgFFFFQAU5abTloAdRRRQAUUUUAc/RRRXQeYFFFFABRRRQAUUUUAFFFFABRRRQAUUUU7MdgooopCTvsFFFHTmgArE1rVv7HtEudpYMcVrC4tyR+9j/77X/GuZ8TX2npZwpdAToZSNsbDildDgmtza0vURqdil5gqGOMVo1i+H3tn0iJ4VMcZZiEPJ447VtUyajtuFFFFOzHcKKKKLMTkluFFFFIdwooooDpcdmnVHTs07MdmOooooswsFFFFISd9gooooBu24UUUnPbk0AtdhaKD8rBT1Iow3QAmgAoo+Y5AHI7Hig8MFPUildDasFFGG6AE0nPIxyOxp3RTaW4tFIGXds74zSjJycHigFrsFFAw33T9c8UpDDnGaY3o7MQU7dSBWb7oNINxJAHI7HipuhX0uSUUN8rbepxnikyeyk0XQXFopqurHAOCOx4p1F0DdtwooX5vwoO4AkKePwouhtO1wIyMVLGhHOajUFm2jnAyatqjbcgZ/ShtFU4sKKQZPbGOuaOeuOKg0FopdrEZUZ/SmjJ7Yx1zQOwtFJTscZHI9qBpN7CUUAMV34x9aSgGmhaB1pccZHI9qRdxXfg/Q9aBWJKKO2R0pu72b8v/AK9A7MdRTd3IGDz7dKcTgZoCzCimswGOpHsM0oYH/wCvxQKwtFBOBSMwGOpHqBQNJvYWim7h2z+PFOJAJHXHpQCTCilAJGRzTQwJwO1BpYWignAzTWYLjqR6gUCWo6imluOAT+lKWXJUHJHpQAtFHzYyAaTOfu8n06UDsxaKQnDFRzj0pfmxnBoEB5GKbto3ZHyDJHbpTvpzSugG7aAMGnYbspNITg/zouOwtFBIyQOcelHOMgGmIKKRSG6GlY7RlulABRSgMRnBFGDjNK47CUUUUxBRRRQAUUoBJxUgTBpXQ0mRDJpwUmpgAKOKTZpYYEweadgUtFSHKgwKKKKASCiiigphRRRQRYKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigFuFFFFBYUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABSg4pKKAHA5paaKdQAhGaay5p9IaBMhopSMUlNEhRRRVMAoooqbAFKDikopAPBzS01adQAUUUUAc/RRRXTZnmBRRRSAKKKKACiiigAooooAikuLeJissqoQu75jipEZZAGQ5DdD6ivL/HP/IRiQAEeVznP9K7izjnfRYUifbI0KhGPYkdaV0bOCSuzRuLu1tT/AKRKsY/2jg/l1pIL21uBuhkVx/snJ/KvOrLwhf3rvdarI0QDbVbOWb35rO1vQ5fD86SQSMQ/IcNgj644pc7NFhot2TPXJJEiUuxwq/ePpUcd1bTELFIjMecAisWHdr/hzY7bJJo8MfQqeTXnNtbvZ66lnaNvIkEZcZx65xTujOnT38j2B7q3iYpLIqEAH5jilneH7OzM4COpG7tyMV5p44I/tNAFBBjHX/61dHqmf+ER4PPkqc+9F0N07Wv1MKPwXZupMd7G2z72FGB+OaS38LWMkggg1CJyDnChW/Wsvw7ps+rLNZK/kRLgttzk0zWtIk0OdBG+4MCwIG0jB74qOlzpvK/LzK56ppNgNLsEtAwfbu5x6nNSz3trbAG4kWPdyoc4JH07Vn2N9JLoMd8f9YImYn3ArzbTLUa9fumoXBiZl37hyT9GPAq7o5YU+a7fQ9bgu7S6IFrNHLkZwjAkfhUxYKhkYgKO5OBXn8fhC5tr6KW2uSYjg714b6HFb3iDS7zU7VY7STaAcFRmhSb2JaitzWXVNMaTyluYyw/2uPzq6CCnmA5T1FeaX/g8WemNdeezNGu4jHHXHFXfBF5PL51hIxdUwy7jS52XKnFpyXQ7ee7t7UKbmRYw3qf50R3drLEZ45UeMdWVsgGvJruCXUvETQXkpi3uVG45UAdxXbHS4NL8OXdvbyGdRkknAzx7U7onkVrnSx3NtKzJFKrsoBIU5qZsIvmSEKvqTgV5n4GVft8zYC/IM46VT1fUZdY1U2In8q3Dbc54FF1a5Spe9yHpn9o6dv8ALFzHuHX5uPz71cXnBHQ8gjnI9a85uPBUX2I3FvdtIygsFx8pA54pvgzU7iO9OlzOzowJGTyCOwpc7KVCDu1LY9GmuLe3iM9zIsSdixxmq0GqafcH9xPG/wBGrltX8N3uqaqj+duhIy2f4f6ZrmNd0A6GsdzDOZN5wNwwQR34p3YlSg3ZSPXC2OgzkZFQzXdrbkrNKqYGTk1iaFqEl1oYu5PvRhs57lRXndlAfEGqlNQnMOSxBzwR7ds0XQqcLXXY9et720uzi2mjk+jCrX8Ww53emP615x/whVzDOjWd0XTIYk8EAfTvXX63HdvpE0Fkzl1UAHPNF0TOKubAwQWUhgODg9KSuA8H2eqw3Uhu1dYcZIc5yfWvQKYqyUXaJzfiHXBotr8mfMfgN6ZrhRc+Kr0NdxtKY8ZBUgV6pd2VpersuYw4qKWa00uyJJ2QIDj6jtSbSNaVSz5bXOK8N+Jbqa8XTdR+fcDtc9Q3YH3rf8Ra4NFtgEz5r8bvTNee6DA+oa+ksQO1ZPNPsAc167d2VpfLtuYw4qBVElqeVi48VXitdxtKY8ZBUgVt+G/EtzNeLp2o/vNwO1z1DDsfeuzmntNLsSSdkCA4+o7V5T4et21DX0liB2LJ5p+gOaDWElUV7HeeKNautJiWO0XmTrKBnGe2PWuQ2eMJ4ftpaRU6jnGR64r1Z4kuFEEiqyZyMjkGuK1DxglnLLZiBnkTKhiwwPwq7kUZ8zskV/DXiW7vbpNOvwH3KSr9CCOx966HxDrR0u1Qxn95IWVcdQQOtcf4T0ae4vF1efb5KFipVgSWPt2qLxy7HUYY+yQg/wDAietZ8zLlBe0Kkdx4q1FPtsDymI5OVOMYrd8NeJbme8XTdS+YtwrnqG7A+9ddoaLDpFssPypsDEepavKtW22viCVo+PLmBGPUc0wUlKTgkei+JdeGiWyJGSJZs/N6Vwn2jxXdK14hlMeMgqQK9UudPtL8bLuMS4AGfwzTLqa10qwZmYLCinaPUjtQZ0alvdtc47wz4murm8GnajtckfK5659PrXoJb0rxvwxaSX+uJPEPljcyH8DXpet6mml6e9yWAcgiMH+f0FATgr2Ob8U6/PbXP2PTnKvEMzMp+77Ung3U7++nnF7K03lgEZ/2qytO0qVtF1DWrwZkkicICcnB7n19qseAv+Pi5bIwY0J55GKDRxTjZHW+ItcOi2qeR/rpiwB9MdyK4FLnxdqSm+jeUx4LAocDipPHMnmawkKNny4FI+pzz+lepaCiR6RaJEMKYwSPdqDWMEldnF+FvFF1dXQ0vU8MxGFc9c+h966DxRrp0e2jMJzLKPkx2968v1Eiz8Su0XHl3AIx7HmtTx3IX1kRg8RRrs9gRk/ypXQuREK3Pi/UE+3RSSmMgtlDgDFdJ4V8U3V1dDTNUwWI+Vz1z6fWuz0NEi0m0SIYRogce7V41fbbTxK5j4CXIxjtzj+tFynFWPdz0rzXxd4kuYr02OlyGMwjMrL2PcD3FdZ4h1ePR9PkuCwEj8Rj6cZrz+x0eQeHdQ1q9G6aWNtmc5wxGSfX2pkQ03N3wTqd/qMtyL2VpRGBjP8AtV6CeleYfD7iW95HRM89K7HxHrUOi2BkZts0mREOvPr9KV0E1qcl4t8SXMV8bHSpCnkj96y9j3A9xVzwRqV/qElz9tlMvlgYz71g2ejSJ4c1DWbwbp5Y225znDEZJ9Se1aHw9B8y9CkH7nf2ouW0jY8a6le6baW7WcpiLyMCR14GeK5CK88byojxtckSDKkAEH8cVs/ENyws06nczbR1+YAU2Xxn/ZkMNjpkQmEUfzu+7GfQAUyLMseH5/FDasBqf2gxdCGA2/icV1+vazFotkbh8FmO2MZ6k1i6D4ug1iRbW4QwysMqQ3DH0xR4g8MHWbuG6WXYBgOjZI47jFAWtuefP4g8SiIXAuHjjkdthHQgd69V8MXdxfaLFc3LGSQsRk9eK47xzbQ2drYW1su2NGdAPYYNdR4NGfD0I/2n/nSuU+Wxe13WotGsmnfG9jtjGepP+FeSv4g8SrEtx9oeNJHbYR0IHevQfEPhdtYv4blZNqjAdWyRx6Y6Vg+OreGzt7C2t12xoXQD2GDTIhsdl4bup77RIbi6Yu5JyT7cVieLvEOoaYVs7FCu7/loBnJ9BWp4NAPh+AH1f+ddFLDHKoEqK2zlfY/jQF7M8cceNIbc6jO0qqAWwDzge1dV4S8TXeqTnT74Kzqu5X6E4xx+tUtX8cR+XcWUUB8z5otzMCOR1AFReC9BnWf+2JyoUg+XtbOSe5A6UrlvY7TX9ai0axadiN7nbGM9Se/4V5K/iDxKsSTm5eNJWbZjoQDjNega/wCFm1fUoboS7VGA6sCRx1xjpmsHx1BFaJYW1uu1E3oB7AimTDY7jw1dT32jQXNyxeRieT7cVi+K/E76QEs7b5ZZFJZhyVwfT3rQ8I5/4R23wcH5ufxrzvxqr23iJrqThZFUKTwDt60BaxPHH42uof7QjeTYeRzjj6V0nhbxVcalMdM1IATBSUYdSR1H1rXtfE2jf2ZHO86qUX7h+96YHrXlmiyyf8JHDLCMBpm2E+jdR+VK5TPRPFXiZtI22FruE7qCSOcVxLT+NDEdQbz1jwWz0AFet3mlafqH/H3EH2sMHvxzyaq69f2mmaTKZcBDGUjA6ndx0oJWpz3hLxTPqkhsL/HmBMrIOM//AKqm8W+I5dHijs7TKzOpfPcAH+tcn4E0+WbVFuQP3cCnJ9d3FdD420We4CatbEF4lMbA91zkVBpZGDFD42uLf+0A7hGG4YYZ/Ktzwn4nuNRuP7N1IAzFSVf3XqPrVe18d20FgkD27LNGmwHIAzUXgrRZpL1tbnACAN5a5+bc/UkdhimhNM0vFniLU7CZbCwQpu4DgZ3H2rlpx42tLY6hcvIqgbsA9B9K9jljjdQ0qAmPlfUfnXmuseOI5raexhtyHfdEWZgQPfAqyUjX8JeJbnVneyvQpkjXIccbucHipPFviaXSQtvacTSJuB7AZ5yPWsvwToM0EjatcFRvRhHtbP3u5xXLeL3M/iOcHohRQD6YBP55pXRdkPEvjCSFtRWSYpgMGzxXZeEfE82qO1hqGPOUZVv73qK7O3iVLNLZBiPyguPqK8P0hhB4nhK5H78qPoxxUolnvVKOaZggFv8AaP8AhUignmrJswC5pwTB5pwXFOqWxqLDAoooqS0gooooKCiiigAooooAKKKKACiiigGFFFFBFgooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooBBRRRQXcKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAUU6minUAFIaWg0CZE1NqRhmogecU1uSLRRRVgFFFFJgFFFFQA5adTVp1ABRRRQBz9FFFdh5gUUUVFmAUUUUWYBRRRRZgFB6UUUWYHlvjgj+04/Xyc4/pXdW9x9k0OO5kHEMCOQD6DpWZ4i0FtYMdzbMA6DByOopuh6NqFrZ3Vnqh3JKNqgHOBWdmbTknGyOPsZ9b8QXxSO4KGNiWJztUHvgdaNf0m+06KP7Zc+erHCjBHT61qw+H/EGk33m6fsdR8uQcAp7+9bfiLR9R1e3tEtgrMh3Nk4/nRZmzqxU7xZX0iSSDwm88X31Rj+Zrn/AAXGJdWlMxBkRSQD/ET6H2ru9I0yay0cWF8oz8wYKc5DVxc/hDWLScy6UwIDZXnawHoc0WZnCcUp3e5F41IbVFCnO2MA49fSum1bK+FCp6iBD/8AWrGsvCWo3F0tzqxAAbc6k5Jrr9asJ7/S3tLVRucAAHsBRZjnUi3Cz2OV8DEg3R77RSeOWObb1IZf61reGdFvtJEv2sIPMAHyuDjFN8S6NfauYvsYU+WSfmOM5os+SwKcfbc19C94bjMmgQq4BUxsCD6Zrmbjwgl1K8uj3KhVY+YhydvoFxXVWGmSwaKlhdkpIEKHbz1rkW8L+INOlaXTJN57fNx+PrRZhTnFKV3uZMz6z4ZvVSWUljyATkEDrXrsM5lWO46B1Vv++hnFed23hHU7y4E+rvkdWBOT+FejIqIqxqPkUjA9gMVcNNyazi9mZmtsf7Huj28o/wA64jwKM39wv95RXfalbPdadLaxEbpEKgnpmuX8M6JqGk3ck12EAZcDawPIqLMUJJU5RZzviC7i1TVlsoo0i8t/LZh1yeM11VzpiaD4fvLeSVnJUYJ757Vn614Tubq6a/0sgGTBZTwd3rmsz/hG/E96fIupj5ffdIDRZmnNDltcn8DgC9uGXDBUXPpk9qj8aWVrBLD9niWMybixHrXaaJo0Oj25iUhnc5dgMA/hWX4m0W/1WSJrNVYR56nGc0NPksSqvv3NPw7mTRrWQnJKEc+nSvOvDpP/AAksRH/PWT9c4/lXpukW02naXBa3K4eNSDjnqa5DSfDWp2OsJfTeX5SOWyGHOc/40WYQnFRkm9z0HC9Rkdzj1rh/HP8Ax6W4/wCmhH9a7jtXNeKNIvdWt4ls1UlH3nLAdsVZnSaU7sj8HDdoaqQDlyMHoaxbrwfb3k7y6RdAKGJkQ87WP93HSum0TSbjT9HNheHaX3fMhyRu9PcVyUvhjXbO4ebSpN2ejA7SR/tZ61FmaKceaTvuZU6az4bvVSeVumQM5BAr1e1vEuLKO+mby1ljDknoPrXn9v4V1m+uFfVpB5akEhjluOwPpXbajpxuNJk060wpKhUz04Ocfj0osxzlB7MuW95b3mfskiybTyFPNWuvSuH8K+HtQ0u6lubuNUUgjANdwCByegrQxk7bamdqWpW2l25muzg/wqOrH2rzdzrHiy+yqskSnC8fKB6kd66rxPod/qy2/wBmCssWc7mxnPT8q5ZPCHiGP/UvGmeu2Y/4VnNN7HRFwSvfU9E0zR7XSLcwWi4LYLsRznvj2p+o6nbaZbme7O0/wr3b6V5/B4W8QRTJJJIm1WBOJSTge3euh8T6JqGsC3FoqssajO5sc+tFmZyhG/xHLN/bHiy+LBSkSnC5HygepHevRdM0e10i3MFohG4gsx6k98e1edp4Q8Qx/wCpeNM9dsx/wqza+FvEEdzFJLKmxXBbEpJwD6d6LM3qcv2Wd9earZacU+1yBDIwVR15PTOOg96qy6Bo95G1w0SkyZJdT096o694aXVdtzAwjuAuw/3WHofT61yg8N+KFBtVbEY4wkgAIoszGEI6S5rMy9DnlsNehitT+78wxsAeoJxnFavjmJl1CKduFaLb+I/xrotB8J/2dMuoagQ9woKgL0Ge5962tX0iLV7VoXwJMgox6D60WZpKcedu4mgXEU2jQeWwJVVQjPTHevLtRxqPiCUW3zLJOAreueP0rRbwp4it98FsflP8QbArpPD/AIT/ALNm+3X5BkAwirzj1z70WY4Sgm5X1Oj1TVbbSIGlnPzHGEHU8Y4rzfZrPiy+MjApEpwMj5QPUjvXV+KdBv8AVpIpbQKVQYO9sZPrXKr4R8Qx8QvGgPULMef0osyYOCV76npGm6Xa6Pam3slKZ+Yseue/4V534zaeTVBG6GSGJVKrzjDdQSKltfC/iGK5illkTarAtiUnge3evTjHE7eYyLuPqMn2osyYyUZM8oufFV3PYNp4tUjjaMx5DNkA9+eKzfD2s3GkXB+zwLJ5gVScEkYPWvW9VsTd6bc2sAHmSxsq5CgZPv2rC8KaBfaXPJJfLFgrhRncc0WZpCatuc348jZdShuivMsC9OOmf8a9G8P3CS6NbMhH7tArDI4x3puuaHb61Yi1fh0O6Nj2btn2rzc+EPEtvugtGYqT1D4BpGrmuVIzLlRqfiZvs/zJJcDafXmtfx/bvDq63GMJJGoH1AIx+tdP4b8IjSpvt98wafGEVei/X3rc17QrfW7YRMdsicox6A1FmLmRJoFxHLo9s6Efu0VWGRxt715BcAaj4kZrf5kkuRtPrzWofCPiW3DW9oWKk9Q+Aa6zw34RXSZvt18wefGEVei/X3osHMjlfGktzLrbRTJvht9uxedpB5OcUl54vvbvTpNN+zRxxvH5eVLZA9eeK9gMcTne6KXP8RGayta09r7SbmztwvmSxlVyAvP1qyVJXPIvDmt3OlXMi29ukvnlAcg5ABwa0/GVxPJr5WRDJDCi7UIO3B5PTmup8LeH7zTJJTqKR4cYUA5NdqY1dt7ohY9TtyaizKlJNnkN94wvL3TZNNNtHGjx+XlS2QPx4rM8Pa9PpE8ghgWQTMgYgEsADgmvXNb097/SbmztgnmSptXICjOR3rB8K6De6VLK9+sfzDCgHJosHMjW15tCWLzNcSM5I8v1IP8AKr9pY6almq2scfkEZ4wcj3J5rL8Q+HIdbRZIyI54/uE/dI9DXD/8Ir4uiRoIHGzoT5nGPYVY7owoR5fiEf2fjYbgGL6d8V76Q25sDhsfyrhfD3hBdKn/ALRv2DTjhAvIUf41p+KdN1LUoI00tthDbmJbaCPSkxSkmznPiJu8u0KqWIdiQB6j/wCtXR+DQU8PQFgR8zdeOprhJPB3iSfidonA6ZmP+FOg8IeJ4QiI6BVYHCzH19MVNhtRsewAnkZzzivNPiJu2WZVS2HckAeo/wDrV0HiTTtU1Ozgh05trxkFyW29vWuHk8HeJJ+J2icDpmY/4VZNPTc73weGXw9BkEHLcHjrWld6zp1nPHa3MqiSU4Azkfj6V5nB4Q8TQhEWRVRGBwkpJ4PpXTa/4Q/tSUahZsqXOwK6H7px60i58t9DR1Dw1octrLcNEqlgz7wfavPvBN7NDrkcEX+pmDIVHTABIbB+lWv+EX8VurW3mnyyMEeYMYrrvDfhRNC/0y6YSXJXZlfuqv09amxLeh2YD4IOOeleYfETcDZlVZsM5IA9ef6V0ninS9U1KKIaY23YfmJfaDXDSeDvEk/+vMT46ZmP+FWKGm56B4RUjw9bhht+919+auaxoljrEAjv15T7rDkivOLbwl4mgMarIgRGBwspPH0r1K5N9Daj7CiyShcDcwAB96RUndnm2qaDo/h2I3jSvcTOCII5MYDHj9Kg8CaY9zfSavMuY4xgegl6HHtVl/COu6vfi51qSIByfM2PnC/7Ixwa9KtLK2sLZbW3GyNF2j/E+9TYG0Zmt67aaHAHuTmVuUjHUmvObbTtY8W332y7JSEtyccBfQD1Nb3inwxqmqX6XNoEZFTb+8bHPrWAng3xLGu2OREB7CY4/lVlR5Uj1Sx0+30y1+y2aYReR/eP1NeXeOb24k1r7GWIhRVDAHAOf8K0NH8Ma/aalBc3cimJGywWUsSPTFdH4l8MR65tuIWWOdf73Q/X8KixKYyy8M6C+loRbI5ZMlznr6/hXAeHbyax177PasXieUx9eT9fYVqjw540WNrWOUiEcKBIMflXR+HPCLaVci+vGVplHyAdieCTRYpyR0N1rWnW10lhcSr5rEjHbpnk9qx9U8M6GbGacxLG21n3g+1UfEHgwX9w1/pjCOV+XRs7WPrntiub/wCEX8VzIbZ5SYyMEeYMVRN0R+Bb2eLVxbqcQzIw2jkfKMg/jVbxjE8PiKVyPlkCsD+AH9K9D8NeGI9CzcTESXTqFZh90AdhUviTw5HrsAKEJcLnDHpg9qmwXNiC6hOnpd7hsEYZuegArxXQ42u/EkMkYyomLj3AOc1pf8Ij4nUfZoifLwFPz4BFd14Y8KLoe66mcPcvxkdFX0FFhnXKGLlewyT+NTjAGKQADp360tNsdgoooqSgooooAKKKKACiiigAooooAKKKKACiiigAooooBhRRRQQFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFALcKKKKCwooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAFFOpop1ABQaKKAY081HsIOc1NikIzTJsQkEUVNikK5p3CxFRTimKQjFDYrCUUUVIDlp1NWnUAFFFFAHP0UUV13R5d0FFFFF0F0FFFFF0F0FFFFF0F0FFFFF0MDlhg8D2oXjGABjrjvRRU2YCBV6Ece1Ln+IqrMO5oooswDjqoCt6ik2ofvDd9aWikAhVD1GfrQVB+bJDe3SlooAF+XocfQChgp6qD9aKKADC9Vyp9ulIAO4/KlooATanYfnS0UUACjaT6HtQoC8A4HbAFFFAC575Jb1NNwepY5+gpaKACnEqeqgn1NNooAX3UkH9KTn1/DAxRRQAdOlDBTztBJ65oooAVcAYAAHtS5HTaD7nrTaKAHkLgYAz6mg8jFJmnUAIcbhwNoH45paKKAGlQSMqDj160/j0P5mkooLuhcgcgH86aOB0BYnqfSlooC4vHofzNHHofzpKKBczEHU0mCOhz7EDmnUUDugGFHygKT1xSUtFAXQ3GRhwCPSgDGMAADsKdRQMDgnJUHHfvS5HofzNJRQA4Mcj/E0m3LbiaQdakoExCARg1PEWxgnj0AFQ1NHQyqZKcEYNJj1A46elOorM2EAUcqMHvig4IweaWigBuPUDjp6UoCjlRg98UtFABSYB4PIpaKADJxjP4ACiiigBMA8HkUuTjGfwwKKKAAdeaOWPzHI9MDFFA4NBd0PK5GB8p7YpoQKcjk9804NmloID/PWkPIx1+tLRQXdDcHOeD/AJ7U7/PWiigLoQjIwcj6GmhOck/lT6KAuhhU4wDn2I4pduFwny564p1FAXQwIFORk+uaf/nrRRQMQjIx/Wkx3wAR6U6igBCMkH096XAPBoooATAyDgcd+9L/AJ60UUAAP+c0wJzljmn0UANwxOWbPtjinDpzj8KKKADr3xTCpxgHI9CP8KfRQAgAUYUYHeggEYPNLRQNCBd3UAn9KnVVUYUYpFTHOafUtloKKKKksKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigGFFFFBAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAtwooooLCiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAUHFKDmm0ooAdRRRQAUUUUAFFFFACEZppXNPooBkW2kIxUpFNZaCLDVp1IBiloAKKKKAOd3UbqbRXQeQO3UbqbRQBJRTd1G6gB1FN3UbqAHUUUUF3QUUUVd0F0FFFFMYUUUVFmAUUUUWYBRRRSAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigBafUdOzQA6iiigAooooAKKKKACiiigAooooAKKKKACiiig0CiiigAFO3U2gdaBElTR1DUkTZOKBw0LFFFFZm4UUUUAFFFFABRRRQAUUUUAFFFFABRRRQAo60+mDrT6ACiiigAooooAKKKKACiiigEFFFFBoFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFPC55pnWp1GBSZSTYo4GKKKKguwUUUUFBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAwooooICiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigFuFFFFBYUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABSikpRQA6iiigAooooAKKKKACiiigApDS0UCY2kpxFIRigmwlFFFAHN0UUV02Z5AUUUUWYBRRRRZgFFFFIB26jdTaKAJKKbuo3UAOoooq7ou6Ciiii6C6CiiimMKKKKizAKKKKQBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUtJRQBJRTc06gAooooAKKKKACiiigAooooAKKKKACiiig0CiiigAoFFFADt2eKki+VvrUI61KDgg0Ai3RSKdwzS1mbrXYKKKKACiiigAooooAKKKKACiiigAooooABwaeGzTKUdaAH0UUUAFFFFABRRRQAUUUUAgooooLugooooGFFFFABRRRQAUUUUAFFFFABSE4GaWkPIoAeozzUw6YqNBgVJUtmsAoooqSgooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigGFFFFBAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAtwooooLCiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKUUlKKAHUUUUAFFFFABRRRQAUUUUAFFFFABSGlooExuKKdiimKxzFFFFdZ44UUUUAFFFFABRRRUWYBRRRRZgFFFFIB26jdTaKAJKKbup1ABRRRV3Rd0FFFFMYUUUVFmAUUUUgCiiigAooooAKKKKACiiigAooooAKKKKACiiigBafUdOzQA6iiigAooooAKKKKACiiigAooooAKKKKC7hRRRQMKKKKAAdakzjmoxTs54oAsRvnjFTVVQ7TVocjNRZm1OSsFFFFIYUUUUAFFFFABRRRQAUUUUAFFFFABSjrSUo60APooooAKKKKACiiigAooooAKKKKAQUUUUGgUUUUAFFFFABRRRQAUUUUAFHWilUZOKGNEyjApaBwMUVmaWCiiigoKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooBhRRRQQFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFALcKKKKCwooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAClFJSigB1FFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAHMUUUV2HihRRRQAUUUUAFFFFABRRRQAUUUVFmAUUUUgCnbqbRQBJRTd1OoAKKKKu6LugooopjCiiiswCiiigAooooAKKKKACiiigAooooAKKKKACiiigApaSigCSim5p1ABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFBoFFFFABQOtFA60CHk4GasRHI61AeaFyrA54oHDTcuUU1SCKdUWN0r7BRRRSAKKKKACiiigAooooAKKKKACgcGiigB4bNLTB1p9ABRRRQAUUUUAFFFFABRRRQCCiiig0CiiigAooooAKKKKACiiigBCcDNSJyc1GRkVKgwKTGtySiiioNgooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooE2FFFFBIUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAgooooLuFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUoOKSigBwOaWkFLQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQBzFFFFdh4oUUUUAFFFFABRRRQAUUUUAFFFFABRRRUWYBRRRSAKduptFAElFN3U6mAUUUVV0XdBRRRU2YwooopAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAC0+mU+gAooooAKKKKACiiigAooooAKKKKACiiig0CiiigAoFFFADt1OPIqMdakoEORivWrAINVTyKcrFevNDLhN9S1RTFcNT+KizNbBRRRSAKKKKACiiigAooooAKKKKAFHWn0wdafQAUUUUAFFFFABRRRQAUUUUAgooooNAooooAKKKKACiiigAooooADUy9KhNTL0pMa3HUUUVBsFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQF0FFFFAmwooooJCiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAW4UUUUFhRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAOFLSCloAKKKKACiiigAooooAKKKKACiiigAooooA5iiiiuw8UKKKKACiiigAooooAKKKKACiiigAooooAKKKKizAKKKKQBUlR07dQA6iiimAUUUVZoFFFFZgFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABTs02lFAD6KKKACiiigAooooAKKKKACiiigAooooLuFFFFAwooooAB1qSox1qSgApCMjFLRQAi5XvUqucjNR0UFc7LQYU7iqgJFPD8jNRZl8yLFFR+Yp4Bp24UWYcyHUU0MKdx60WZVgoo49aOPWizCwUUcetHHrRZhYUdafUYIBp24UWYWHUU3cKNwoswsOopu4UbhRZhYdRTdwo3CkFh1FN3CjcKBodRTdwo3Cgu46im7hRuFAXHUU3cKNwoC46im7hRuFAXHUU0MDTsigLgamXpUBIqWNgwpMa3JKKKKg2ugooooC6CiiigLoKKKKAugooooC6CiiigLoKKKKAugooooE2FFFFBIUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFALcKKKKCwooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAHClpBS0AFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAcxRRRXYeKFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRUWYBRRRSAdup1R1JTAKKKKq6LugoooqbMYUUUUgCiiigAooooAKKKKACiiigAooooAKKKKAClFJRQBJRTc06gAooooAKKKKACiiigAooooAKKKKACiiig0CiiigAHWpKjFODUAOooooAKKKKACg8jFFFADQoHIp3PrRRQAoJFO3mmUUD52P3mjeaZRQHOx+80bzTKKA52P3mjeaZRQHOx+80bzTKKA52P3mjeaZRQw52P3mgMTTKUVFmHOx+T60ZPrTQc06izDnYZPrRk+tFFFmPmYZPrRk+tFFFmHMwyfWjJ9aKKLMOZhk+tGT60UUWYczFBNLuNNoosw5mO3GpYHw231qCnxnDg0jSM3c0A2aWmgY5p1QdIUUUUgCiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAW4UUUUFhRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAOFLSCloAKKKKACiiigAooooAKKKKACiiigAooooA5iiiiuw8UKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooorMAp26m0UASUUUUwCiiirNAooorMAooooAKKKKACiiigAooooAKKKKACiiigAooooAUU+mCn0AFFFFABRRRQAUUUUAFFFFABRRRQAUUUUGgUUUUAFA60UDrQBJRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAKOtPpg60+gEFFFFBYUUUUAFFFFABRRRQAUUUUPYAoB2nNFJjNZjRoo28A9KfVeB8rjFTg5NSzqhJWHUUUVIwooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigFuFFFFBYUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFADhS0gpaACiiigAooooAKKKKACiiigAooooAKKKKAOYooorsPFCiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKizAKKKKQElFN3U6mAUUUVZoFFFFZgFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFACin0wU+gAooooAKKKKACiiigAooooAKKKKACiiig0CiiigAoHWigdaAJKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigBR1p9MHWn0AgooooLCiiigAooooAKKKKACiiih7AFFFFZgSwttO31q6Bjms5etXlbOKTNabWxLRSA5NLUG4UUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQC3CiiigsKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigBwpaQUtABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAHLbqN1NorsPFHbqdUdO3UAOooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiiswCpKjqSmAUUUVZoFFFFRZgFFFFIAooooAKKKKACiiigAooooAKKKKACiiigBRT6YKfQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQaBRRRQAUDrRQOtAElFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAo60+mDrT6AQUUUUFhRRRQAUUUUAFFFFABRRRQ9gCiiiswCpon+YA1DSc9qBx3NJetPqCI/L71MDk1FjqWuwtFFFIYUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFALcKKKKCwooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAHClpBS0AFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAcpRRRXYeKFFFFADt1OqOpKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooorMAqSo6kpgFFFFWaBRRRQAUUUVmAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAU7NNpRQA+iiigAooooAKKKKACiiigAooooAKKKKDQKKKKACgdaKB1oAkooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAFHWn0wdafQCCiiigsKKKKACiiigAooooAKKKKHsAUUUVmAUUUUAPjYhge1XUIPNZ9TRybSM0NMuE7bl6imK4an1FjoWuwUUUUhhRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQC3CiiigsKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigBQcUoOabThQAtFFFABRRRQAUUUUAFFFFABRRSE4oACcUUhOaKBXRy1FS+TJ6UeTJ6V13R5HIyKipfJk9KPJk9KLoORkVO3U/yZPSjyZPSi6DkYlFP8uT+7R5cn92i6DkYyin+XJ/do8uT+7RdByMZRT/Lk/u0eXJ/doug5GMop/lyf3aPLk/u0XQcjGUU/wAuT+7R5cn92i6DkYyin+XJ/do8uT+7RdByMZRT/Lk/u0eXJ/doug5GMop/lyf3aPLk/u1AcjGVJSeXJ/dqTy5PSmHIxlFO8uT0o8uT0qrouzG0U7y5PSl8uT0ougsxlFSeVJ6UeVJ6VA+VkdFSeVJ6UeVJ6UBysjop3lyelHlyelAcrG0U/wAuT0pfKk9KA5WR0VJ5UnpR5UnpQHKyOipPKk9KPKk9KA5WR0VJ5UnpR5UnpQHKyOlFP8qT0o8qT0oDlYUU7y5PSjy5PSgOVjaKd5cnpR5cnpQHKxtFO8uT0o8uT0oDlY2ineXJ6UeXJ6UroOVjaKd5cnpR5cnpRdBysbRTvLk9KPLk9KLouzG0U7y5PSjy5PSi6FZjaB1p3lyelKI5M9KLoLMWineXJ6UFJB2ougsxtFGJP7tKFk/u0XQWYlFO8uT0o8uT0ougsxtFO8uT0o8uT0ougsxtFO8uT0o8uT0ouh2Y2ineXJ6UeXJ6UXQcrG0U7y5PSjy5PSi6DlY2ineXJ6UeXJ6UXQcrG0U7y5PSjy5PSi6DlY2ineXJ6UeXJ6UXQcrG0U7y5PSjy5PSi6DlYg60+kCSelLtk/u0XQ+VhRRtk/u0bZP7tF0VZhRRtk/u0bZP7tF0FmFFG2T+7Rtk/u0XQWYUUbZP7tG2T+7RdBZhRRtk/u0oWQ9qG0FmJRTvLk9KPLk9KgLMbRTvLk9KPLk9KEFmNop3lyelHlyelW2gaZJHJtPNW1YNVDy5PSpY2dDkis2XTk1uXaKQNxmk3VNje46im7qN1IY6im7qN1ADqKbuo3UAOopu6jdQA6im7qN1ADqKQHNLQAUUUUAFFFITigBaKbupQc0ALRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQCCiiigu6CiiigLoKKKKAugooooC6CiiigLhRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAoGaUDFIDilzQAtFJmjNAC0UmaAc0ALRRRQAUUUE4oAQnFNZqUnNNIzQJsAc0UAYooJP/9k=",
	})
	var payment1 Payment
	db.Raw("SELECT * FROM payments WHERE id = ?", "1").Scan(&payment1) // dump test by earth
	// ============================================================================ Booking
	//---------------------------------Branch data-----------------------
	b4001 := Branch{
		B_name: "Bangkok",
	}
	db.Model(&Branch{}).Create(&b4001)

	b4002 := Branch{
		B_name: "Pattaya",
	}
	db.Model(&Branch{}).Create(&b4002)

	b4003 := Branch{
		B_name: "Chiang Mai",
	}
	db.Model(&Branch{}).Create(&b4003)

	b4004 := Branch{
		B_name: "Phuket Town",
	}
	db.Model(&Branch{}).Create(&b4004)
	// ============================================================================ Booking
	//ใส่ไว้ก่อนนะเราต้องใช้เชื่อมกับตาราง checkin-out by joon => patch 21/1/2566 by earth
	db.Model(&Booking{}).Create(&Booking{
		Branch:   b4001,
		Room:     Room1,
		Start:    time.Now(),
		Stop:     time.Now(),
		Customer: Customer1,
	})

	db.Model(&Booking{}).Create(&Booking{
		Branch:   b4002,
		Room:     Room2,
		Start:    time.Now(),
		Stop:     time.Now(),
		Customer: Customer2,
	})

	var booking1 Booking
	var booking2 Booking
	db.Raw("SELECT * FROM bookings WHERE id = ?", "1").Scan(&booking1)
	db.Raw("SELECT * FROM bookings WHERE id = ?", "2").Scan(&booking2)
	// ============================================================================ Check Payment
	// ------------------------- Status ------------------
	s1001 := CHK_PaymentStatus{
		Type: "ยังไม่ได้รับการชำระเงิน",
	}
	db.Model(&CHK_PaymentStatus{}).Create(&s1001)

	s1002 := CHK_PaymentStatus{
		Type: "ได้รับการชำระเงินเรียบร้อยแล้ว",
	}
	db.Model(&CHK_PaymentStatus{}).Create(&s1002)
	// ------------------------ CHK_Payment --------------------
	db.Model(&CHK_Payment{}).Create(&CHK_Payment{
		Payment:           payment1,
		CHK_PaymentStatus: s1001,
		Date_time:         time.Now(),
		Amount:            1760,
		Description:       "-",
		Employee:          Sobsa,
	})

	db.Model(&CHK_Payment{}).Create(&CHK_Payment{
		Payment:           payment1,
		CHK_PaymentStatus: s1002,
		Date_time:         time.Now(),
		Amount:            3200,
		Description:       "-",
		Employee:          Banana,
	})

	var chk_payment1 CHK_Payment
	var chk_payment2 CHK_Payment
	db.Raw("SELECT * FROM check_in_outs WHERE id = ?", "1").Scan(&chk_payment1)
	db.Raw("SELECT * FROM check_in_outs WHERE id = ?", "2").Scan(&chk_payment2)
	// ============================================================================ Check Payment
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
		Booking:          booking1,
		CheckInTime:      time.Now(),
		CheckOutTime:     time.Now(),
		CheckInOutStatus: checkout,
		Employee:         Sobsa,
	})

	db.Model(&CheckInOut{}).Create(&CheckInOut{
		Booking:          booking2,
		CheckInTime:      time.Now(),
		CheckOutTime:     time.Now(),
		CheckInOutStatus: checkin,
		Employee:         Banana,
	})

	var CheckInOut1 CheckInOut
	var CheckInOut2 CheckInOut
	db.Raw("SELECT * FROM check_in_outs WHERE id = ?", "1").Scan(&CheckInOut1)
	db.Raw("SELECT * FROM check_in_outs WHERE id = ?", "2").Scan(&CheckInOut2)

	//******************ระบบ review********************
	// Set Data Systemwork
	db.Model(&Systemwork{}).Create(&Systemwork{
		Name: "Booking",
	})
	db.Model(&Systemwork{}).Create(&Systemwork{
		Name: "Sevice",
	})
	db.Model(&Systemwork{}).Create(&Systemwork{
		Name: "Check Payment",
	})
	db.Model(&Systemwork{}).Create(&Systemwork{
		Name: "Payment",
	})
	db.Model(&Systemwork{}).Create(&Systemwork{
		Name: "Review Hotel",
	})
	db.Model(&Systemwork{}).Create(&Systemwork{
		Name: "Subscribe",
	})

	var Bookingsys Systemwork
	var Sevicesys Systemwork
	var Checkpaymentsys Systemwork
	var Paymentsys Systemwork
	var Reviewsys Systemwork
	var Subscribesys Systemwork
	db.Raw("SELECT * FROM systemworks WHERE name = ?", "Booking").Scan(&Bookingsys)
	db.Raw("SELECT * FROM systemworks WHERE name = ?", "Sevice").Scan(&Sevicesys)
	db.Raw("SELECT * FROM systemworks WHERE name = ?", "Check Payment").Scan(&Checkpaymentsys)
	db.Raw("SELECT * FROM systemworks WHERE name = ?", "Payment").Scan(&Paymentsys)
	db.Raw("SELECT * FROM systemworks WHERE name = ?", "Review Hotel").Scan(&Reviewsys)
	db.Raw("SELECT * FROM systemworks WHERE name = ?", "Subscribe").Scan(&Subscribesys)

	timedaterv1 := time.Date(2023, 10, 15, 0, 0, 0, 0, time.Local)
	timeyearrv2 := time.Date(2023, 5, 15, 0, 0, 0, 0, time.Local)

	db.Model(&Review{}).Create(&Review{
		Customer:   Customer1,
		Comment:    "ประทับใจทุกอย่างของโรงแรม",
		Start:      "5",
		Reviewdate: timedaterv1,
		Systemwork: Subscribesys,
		Department: Housekeeping,
	})
	db.Model(&Review{}).Create(&Review{
		Customer:   Customer2,
		Comment:    "โรงแรมสวย ดี บริการดี",
		Start:      "4",
		Reviewdate: timeyearrv2,
		Systemwork: Checkpaymentsys,
		Department: Salemarketing,
	})

	//=====================================================Storage
	// Data Storage
	//Product Data
	TV := Product{
		Name:  "TV",
		Price: 5000,
	}
	db.Model(&Product{}).Create(&TV)

	fridge := Product{
		Name:  "fridge",
		Price: 4500,
	}
	db.Model(&Product{}).Create(&fridge)

	microwave := Product{
		Name:  "microwave",
		Price: 2000,
	}
	db.Model(&Product{}).Create(&microwave)

	aircon := Product{
		Name:  "airconditioner",
		Price: 7000,
	}
	db.Model(&Product{}).Create(&aircon)

	lamp := Product{
		Name:  "lamp",
		Price: 500,
	}
	db.Model(&Product{}).Create(&lamp)

	bedspread := Product{
		Name:  "bedspread",
		Price: 100,
	}
	db.Model(&Product{}).Create(&bedspread)

	blanket := Product{
		Name:  "blanket",
		Price: 200,
	}
	db.Model(&Product{}).Create(&blanket)

	bolstercushion := Product{
		Name:  "bolster cushion",
		Price: 300,
	}
	db.Model(&Product{}).Create(&bolstercushion)

	duvet := Product{
		Name:  "duvet",
		Price: 150,
	}
	db.Model(&Product{}).Create(&duvet)

	bed := Product{
		Name:  "bed",
		Price: 1000,
	}
	db.Model(&Product{}).Create(&bed)

	sofa := Product{
		Name:  "sofa",
		Price: 800,
	}
	db.Model(&Product{}).Create(&sofa)

	chair := Product{
		Name:  "chair",
		Price: 700,
	}
	db.Model(&Product{}).Create(&chair)

	wardrobe := Product{
		Name:  "wardrobe",
		Price: 2500,
	}
	db.Model(&Product{}).Create(&wardrobe)

	drawer := Product{
		Name:  "drawer",
		Price: 900,
	}
	db.Model(&Product{}).Create(&drawer)

	door := Product{
		Name:  "door",
		Price: 600,
	}
	db.Model(&Product{}).Create(&door)

	window := Product{
		Name:  "window",
		Price: 500,
	}
	db.Model(&Product{}).Create(&window)

	mirror := Product{
		Name:  "mirror",
		Price: 300,
	}
	db.Model(&Product{}).Create(&mirror)

	drape := Product{
		Name:  "drape",
		Price: 150,
	}
	db.Model(&Product{}).Create(&drape)

	picture := Product{
		Name:  "picture",
		Price: 400,
	}
	db.Model(&Product{}).Create(&picture)

	vase := Product{
		Name:  "vase",
		Price: 200,
	}
	db.Model(&Product{}).Create(&vase)

	//ProductType Data
	Electrical := ProductType{
		Name: "Electrical",
	}
	db.Model(&ProductType{}).Create(&Electrical)

	Bedding := ProductType{
		Name: "Bedding",
	}
	db.Model(&ProductType{}).Create(&Bedding)

	Fur := ProductType{
		Name: "Furniture",
	}
	db.Model(&ProductType{}).Create(&Fur)

	Interior := ProductType{
		Name: "InteriorItems",
	}
	db.Model(&ProductType{}).Create(&Interior)

	//Storage Data
	//Storage1
	db.Model(&Storage{}).Create(&Storage{
		Employee:    Sobsa,
		Product:     TV,
		ProductType: Electrical,
		Quantity:    50,
		Time:        time.Now(),
	})
	//Storage2
	db.Model(&Storage{}).Create(&Storage{
		Employee:    Hanoi,
		Product:     blanket,
		ProductType: Bedding,
		Quantity:    55,
		Time:        time.Now(),
	})
	//Storage3
	db.Model(&Storage{}).Create(&Storage{
		Employee:    Banana,
		Product:     chair,
		ProductType: Fur,
		Quantity:    60,
		Time:        time.Now(),
	})
	var Storage1 Storage
	var Storage2 Storage
	var Storage3 Storage
	db.Raw("SELECT * FROM storages WHERE id = ?", "1").Scan(&Storage1)
	db.Raw("SELECT * FROM storages WHERE id = ?", "2").Scan(&Storage2)
	db.Raw("SELECT * FROM storages WHERE id = ?", "3").Scan(&Storage3)

}
