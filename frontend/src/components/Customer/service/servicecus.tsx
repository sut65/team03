import { CustomerInterface } from "../../../models/modelCustomer/ICustomer" 



			// //==================================================Customer Routes
			// protected.GET("/customers", customer.ListCustomers)
			// protected.GET("/customer/:id", customer.GetCustomerByID)
			// protected.POST("/customers", customer.CreateCustomer)
			// protected.PATCH("/customers", customer.UpdateCustomer)
			// protected.DELETE("/customers/:id", customer.DeleteCustomer)
			// //Gender
			// protected.GET("/customers/genders", customer.ListGender)
			// protected.GET("/customer/genders/:id", customer.GetGender)
			// protected.POST("/customers/genders", customer.CreateGender)
			// protected.PATCH("/customers/genders", customer.UpdateGender)
			// protected.DELETE("/customers/genders/:id", customer.DeleteGender)
			// //Nametitle
			// protected.GET("/nametitles", customer.ListNametitle)
			// protected.GET("/customer/nametitles/:id", customer.GetNametitle)
			// protected.POST("/customers/nametitles", customer.CreateNametitle)
			// protected.PATCH("/customers/nametitles", customer.UpdateNametitle)
			// protected.DELETE("/customers/nametitles/:id", customer.DeleteNametitle)
			// //Province
			// protected.GET("/provinces", customer.ListProvince)
			// protected.GET("/customer/provinces/:id", customer.GetProvince)
			// protected.POST("/customers/provinces", customer.CreateProvince)
			// protected.PATCH("/customers/provinces", customer.UpdateProvince)
			// protected.DELETE("/customers/provinces/:id", customer.DeleteProvince)
			// //==================================================Customer Routes


            
            
            const apiUrl = "http://localhost:8080";

            async function GetCustomerByUID() {
                let uid = localStorage.getItem('id');
                const requestOptions = {
                    method: "GET",
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
                        "Content-Type": "application/json",
                    },
                }
            
                let res = await fetch(`${apiUrl}/customer/${uid}`, requestOptions)
                    .then((response) => response.json())
                    .then((res) => {
                        if (res.data) {
                            return res.data;
                        } else {
                            return false;
                        }
                    });
            
                return res;
            }


            async function GetNametitleByUID(data: CustomerInterface) {
                let nid = data.Nametitle_ID;
                const requestOptions = {
                    method: "GET",
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
                        "Content-Type": "application/json",
                    },
                }
            
                let res = await fetch(`${apiUrl}/customer/nametitles/${nid}`, requestOptions)
                    .then((response) => response.json())
                    .then((res) => {
                        if (res.data) {
                            return res.data;
                        } else {
                            return false;
                        }
                    });
            
                return res;
            }
            
            

                // Delete Booking
                async function DeleteCustomer(data: CustomerInterface) {
                    let c_id = data.ID;
                    const requestOptions = {
                        method: "DELETE",
                        headers: {
                            Authorization: `Bearer ${localStorage.getItem("token")}`,
                            "Content-Type": "application/json",
                        },
                        body: JSON.stringify(data),
                    }

                    let res = await fetch(`${apiUrl}/customers/${c_id}`, requestOptions)
                        .then((response) => response.json())
                        .then((res) => {
                            if (res.data) {
                                return res.data;
                            } else {
                                return false;
                            }
                        });

                    return res;
                }

                // Update customer
                async function UppdateCustomer(data: CustomerInterface) {
                    let cus_id = data.ID;
                    const requestOptions = {
                        method: "PATCH",
                        headers: {
                            Authorization: `Bearer ${localStorage.getItem("token")}`,
                            "Content-Type": "application/json",
                        },
                        body: JSON.stringify(data),
                    }

                    let res = await fetch(`${apiUrl}/customers/${cus_id}`, requestOptions)
                        .then((response) => response.json())
                        .then((res) => {
                            if (res.data) {
                                return res.data;
                            } else {
                                return false;
                            }
                        });

                    return res;
                }

                async function GetNametitle() {
                    const requestOptions = {
                        method: "GET",
                        headers: {
                            Authorization: `Bearer ${localStorage.getItem("token")}`,
                            "Content-Type": "application/json",
                        },
                    };
                
                    let res = await fetch(`${apiUrl}/nametitles`, requestOptions)
                        .then((response) => response.json())
                        .then((res) => {
                            if (res.data) {
                                return res.data;
                            } else {
                                return false;
                            }
                        });
                
                    return res;
                }

                export{
                    GetNametitle,
                    GetCustomerByUID,
                    UppdateCustomer,
                    DeleteCustomer,
                    GetNametitleByUID,
                }