
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
export {
GetStorages,
};