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
            const requestOptionsGet = {
                method: "GET",
                headers: {
                  Authorization: `Bearer ${localStorage.getItem("token")}`,
                  "Content-Type": "application/json",
                },
              };

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
            };
            const GetCustomerByID = async (uid: string) => {
                let res = await fetch(`${apiUrl}/customer/${uid}`, requestOptionsGet)
                  .then((response) => response.json())
                  .then((result) => {
                    return result.data ? result.data : false;
                  });
              
                return res;
              };


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
            
            async function Customers(data: CustomerInterface) {
                const requestOptions = {
                    method: "POST",
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify(data),
                };
            
                let res = await fetch(`${apiUrl}/customers`, requestOptions)
                    .then((response) => response.json())
                    .then((res) => {
                        if (res.data) {
                            return { status: true, message: res.data };
                        } else {
                            console.log(res.error);
                            return { status: false, message: res.error };
                        }
                    });
            
                return res;
            }
            

            async function GetCustomerlist() {
                const requestOptions = {
                    method: "GET",
                      headers: {
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
                        "Content-Type": "application/json",
                      },
                };
            
                let res = await fetch(`${apiUrl}/customers`, requestOptions)
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
            
            
            async function DeleteCustomer(data: number) {
                let customer= data;
                const requestOptions = {
                    method: "DELETE",
                      headers: {
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
                        "Content-Type": "application/json",
                      },
                    body: JSON.stringify(data),
                }
                
                let res = await fetch(`${apiUrl}/customers/${customer}`, requestOptions)
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
                async function UpdateCustomer(data: CustomerInterface) {
                    let cu_id = data.ID
                    const requestOptions = {
                        method: "PATCH",
                          headers: {
                            Authorization: `Bearer ${localStorage.getItem("token")}`,
                            "Content-Type": "application/json",
                          },
                        body: JSON.stringify(data),
                    }
                
                    let res = await fetch(`${apiUrl}/customersupdate/${cu_id}`, requestOptions)
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
                    UpdateCustomer,
                    DeleteCustomer,
                    GetNametitleByUID,
                    GetCustomerByID,
                    GetCustomerlist,
                    Customers,
                }