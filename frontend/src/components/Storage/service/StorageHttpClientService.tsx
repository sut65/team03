import { StorageInterface } from "../../../models/IStorage";

const apiUrl = "http://localhost:8080";
// protected.GET("/storages", storage.ListStorages)
// protected.GET("/storage/:id", storage.GetStorage)
// protected.POST("/storages", storage.CreateStorage)
// protected.PATCH("/storages", storage.UpdateStorage)
// protected.DELETE("/storages/:id", storage.DeleteStorage)

// protected.GET("/products", storage.ListProducts)
// protected.GET("/product/:id", storage.GetProduct)
// protected.POST("/products", storage.CreateProduct)
// protected.PATCH("/products", storage.UpdateProduct)
// protected.DELETE("/products/:id", storage.DeleteProduct)

// protected.GET("/product_types", storage.ListProductTypes)
// protected.GET("/product_type/:id", storage.GetProductType)
// protected.POST("/product_types", storage.CreateProductType)
// protected.PATCH("/pProduct_types", storage.UpdateProductType)
// protected.DELETE("/pProduct_types/:id", storage.DeleteProductType)


// List Storage
async function GetStorages() {
const requestOptions = {
    method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
};

let res = await fetch(`${apiUrl}/storages`, requestOptions)
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

// List Storage
async function GetStorage(id : any) {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/storage/${id}`, requestOptions)
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

// List ProtypeType
async function GetProductTypes() {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/product_types`, requestOptions)
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

// List Product
async function GetProducts() {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/products`, requestOptions)
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

// List Employee
async function GetEmployees() {
    const requestOptions = {
        method: "GET",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
    };

    let res = await fetch(`${apiUrl}/Employees`, requestOptions)
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

//Craete room
async function CreateStorage(data: StorageInterface) {
    const requestOptions = {
        method: "POST",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };

    let res = await fetch(`${apiUrl}/storages`, requestOptions)
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

//protected.PATCH("/storages/id", storage.Storage)
async function Storage(data: number) {
    let StorageID = data;
    const requestOptions = {
        method: "PATCH",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/storages/${StorageID}`, requestOptions)
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

// Delete Storage
async function DeleteStorage(data: number) {
    let StorageID = data;
    const requestOptions = {
        method: "DELETE",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }
    
    let res = await fetch(`${apiUrl}/storages/${StorageID}`, requestOptions)
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

//update room
async function UpdateStorage(data: StorageInterface) {
    const requestOptions = {
        method: "PUT",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
        body: JSON.stringify(data),
    }

    let res = await fetch(`${apiUrl}/storages`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                return {status: true,data: res.data};
            } else {
                return {status: false,data: res.error};
            }
        });

    return res;
}

export {
    GetStorages,
    GetProducts,
    GetProductTypes,
    CreateStorage,
    GetEmployees,
    Storage,
    DeleteStorage,
    GetStorage,
    UpdateStorage,
};