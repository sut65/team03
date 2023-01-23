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

	// Set Data Blood
	db.Model(&Blood{}).Create(&Blood{
		Name: "A",
	})
	db.Model(&Blood{}).Create(&Blood{
		Name: "AB",
	})
	db.Model(&Blood{}).Create(&Blood{
		Name: "B",
	})
	db.Model(&Blood{}).Create(&Blood{
		Name: "O",
	})

	var Blooda Blood
	var Bloodab Blood
	var Bloodb Blood
	var Bloodo Blood
	db.Raw("SELECT * FROM bloods WHERE name = ?", "A").Scan(&Blooda)
	db.Raw("SELECT * FROM bloods WHERE  name = ?", "AB").Scan(&Bloodab)
	db.Raw("SELECT * FROM bloods WHERE name = ?", "B").Scan(&Bloodb)
	db.Raw("SELECT * FROM bloods WHERE  name = ?", "O").Scan(&Bloodo)

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
		Blood:        Blooda,
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
		Blood:        Bloodb,
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
		Blood:        Bloodab,
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
		Province: "Bangkok",
	}
	db.Model(&Province{}).Create(&Province1)

	Province2 := Province{
		Province: "Nakhon Ratchasima",
	}
	db.Model(&Province{}).Create(&Province2)

	Province3 := Province{
		Province: "Chonburi",
	}
	db.Model(&Province{}).Create(&Province3)

	Province4 := Province{
		Province: "Chiangmai",
	}
	db.Model(&Province{}).Create(&Province4)

	//Memberlevel
	Memberlevel1 := Memberlevel{
		Memberlevel: "Classis",
	}
	db.Model(&Memberlevel{}).Create(&Memberlevel1)

	Memberlevel2 := Memberlevel{
		Memberlevel: "Sliver",
	}
	db.Model(&Memberlevel{}).Create(&Memberlevel2)

	Memberlevel3 := Memberlevel{
		Memberlevel: "Gold",
	}
	db.Model(&Memberlevel{}).Create(&Memberlevel3)

	Memberlevel4 := Memberlevel{
		Memberlevel: "Plantinum",
	}
	db.Model(&Memberlevel{}).Create(&Memberlevel4)

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
		Memberlevel: Memberlevel1,
		Signin:       loginCustomer1,
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
		Memberlevel: Memberlevel2,
		Signin:       loginCustomer2,
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
	// var NoOrder1 Food
	// var PadThai Food
	// var PadKaphao Food
	// var Noodles Food
	// db.Raw("SELECT * FROM Foods WHERE Name = ?", "No Order").Scan(&NoOrder1)
	// db.Raw("SELECT * FROM Foods WHERE Name = ?", "Pad Thai").Scan(&PadThai)
	// db.Raw("SELECT * FROM Foods WHERE Name = ?", "Pad Kaphao").Scan(&PadKaphao)
	// db.Raw("SELECT * FROM Foods WHERE Name = ?", "Noodles").Scan(&Noodles)

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
	// var NoOrder2 Drink
	// var Pepsi Drink
	// var Mansome Drink
	// var Water Drink
	// db.Raw("SELECT * FROM drinks WHERE name = ?", "No Order").Scan(&NoOrder2)
	// db.Raw("SELECT * FROM drinks WHERE name = ?", "Pepsi").Scan(&Pepsi)
	// db.Raw("SELECT * FROM drinks WHERE name = ?", "Mansome").Scan(&Mansome)
	// db.Raw("SELECT * FROM drinks WHERE name = ?", "Water").Scan(&Water)

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
	// var NoOrder3 Accessories
	// var Plug Accessories
	// var Chair Accessories
	// var Table Accessories
	// var TableChair Accessories
	// var Bed Accessories
	// db.Raw("SELECT * FROM Accessories WHERE name=?", "No Order").Scan(&NoOrder3)
	// db.Raw("SELECT * FROM Accessories WHERE name=?", "Plug").Scan(&Plug)
	// db.Raw("SELECT * FROM Accessories WHERE name=?", "Chair").Scan(&Chair)
	// db.Raw("SELECT * FROM Accessories WHERE name=?", "Table").Scan(&Table)
	// db.Raw("SELECT * FROM Accessories WHERE name=?", "Table & Chair (small)").Scan(&TableChair)
	// db.Raw("SELECT * FROM Accessories WHERE name=?", "Bed").Scan(&Bed)

	// ===============     PAYMENT     ===============

	// ===============     สถานที่ชำระเงิน     ===============
	db.Model(&Place{}).Create(&Place{
		Name: "Counter Service 1",
	})
	db.Model(&Place{}).Create(&Place{
		Name: "Counter Service 2",
	})
	// var CS1 Place
	// var CS2 Place
	// db.Raw("SELECT * FROM places WHERE name = ?", "Counter Service 1").Scan(&CS1)
	// db.Raw("SELECT * FROM places WHERE name = ?", "Counter Service 2").Scan(&CS2)

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
	// var Ethereum Crypto
	// var DogECoin Crypto
	// var Tether Crypto
	// var BitCoin Crypto
	// db.Raw("SELECT * FROM cryptos WHERE name = ?", "Ethereum").Scan(&Ethereum)
	// db.Raw("SELECT * FROM cryptos WHERE name = ?", "DogECoin").Scan(&DogECoin)
	// db.Raw("SELECT * FROM cryptos WHERE name = ?", "Tether").Scan(&Tether)
	// db.Raw("SELECT * FROM cryptos WHERE name = ?", "BitCoin").Scan(&BitCoin)

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
		Picture:         "data:image/jpeg;base64,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",
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
		Picture:         "data:image/jpeg;base64,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",
		PaymentMethodID: 2,
	})
	// var PromptPay Bank
	// var SCB Bank
	// var Bangkok Bank
	// var Krungthai Bank
	// var Krungsri Bank
	// var Kasikorn Bank
	// db.Raw("SELECT * FROM banks WHERE name = ?", "PromptPay").Scan(&PromptPay)
	// db.Raw("SELECT * FROM banks WHERE name = ?", "SCB").Scan(&SCB)
	// db.Raw("SELECT * FROM banks WHERE name = ?", "Bangkok").Scan(&Bangkok)
	// db.Raw("SELECT * FROM banks WHERE name = ?", "Krungthai").Scan(&Krungthai)
	// db.Raw("SELECT * FROM banks WHERE name = ?", "Krungsri").Scan(&Krungsri)
	// db.Raw("SELECT * FROM banks WHERE name = ?", "Kasikorn").Scan(&Kasikorn)

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
	// var Cash PaymentMethod
	// var Transfer PaymentMethod
	// var Cryptox PaymentMethod
	// db.Raw("SELECT * FROM payment_methods WHERE name = ?", "Cash").Scan(&Cash)
	// db.Raw("SELECT * FROM payment_methods WHERE name = ?", "Transfer").Scan(&Transfer)
	// db.Raw("SELECT * FROM payment_methods WHERE name = ?", "Crypto").Scan(&Cryptox)

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
	db.Raw("SELECT * FROM check_in_outs WHERE id = ?", "1").Scan(&booking1)
	db.Raw("SELECT * FROM check_in_outs WHERE id = ?", "2").Scan(&booking2)
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
		// Payment:   b4001,
		CHK_PaymentStatus: s1001,
		Date_time:         time.Now(),
		Amount:            1760,
		Description:       "-",
		Employee:          Sobsa,
	})

	db.Model(&CHK_Payment{}).Create(&CHK_Payment{
		// Payment:   b4001,
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
		Booking:          booking1, //dump
		CheckInTime:      time.Now(),
		CheckOutTime:     time.Now(),
		CheckInOutStatus: checkout,
		Employee:         Sobsa,
	})

	db.Model(&CheckInOut{}).Create(&CheckInOut{
		Booking:          booking2, //dump
		CheckInTime:      time.Now(),
		CheckOutTime:     time.Now(),
		CheckInOutStatus: checkin,
		Employee:         Banana,
	})
	//ข้อมูลเลข Booking ยัง dump อยู่ (09/01/2023)
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

}
