import {
    GetRoom,
    GetFoods,
    GetDrinks,
    UpdateFood,
    UpdateDrink,
    GetFoodItem,
    GetDrinkItem,
    GetPriceFood,
    GetFoodItemS,
    GetPriceDrink,
    GetDrinkItemS,
    UpdateService,
    GetAccessories,
    GetServiceByID,
    UpdateAccessories,
    GetAccessorieItem,
    GetPriceAccessorie,
    GetAccessoriesItemS,
} from "./service/ServiceHttpClientService";
import { Box, Button, Container, FormControl, Grid, Paper, Select, SelectChangeEvent, Snackbar, TextField, Typography } from "@mui/material";
import { DrinksInterface, FoodsInterface, ServicesInterface } from "../../models/modelService/IService";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { forwardRef, useEffect, useRef, useState } from "react";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { StorageInterface } from "../../models/IStorage";
import { Link as RouterLink } from "react-router-dom";
import { useParams } from "react-router-dom";
import { grey } from "@mui/material/colors";

const theme = createTheme({
    palette: {
        primary: {
            main: grey[800],
        },
        secondary: {
            main: grey[50],
        },
    },
});

const Alert = forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function ServiceUpdate() {

    const [service, setService] = useState<Partial<ServicesInterface>>({});
    const [time] = useState<Date | null>(new Date());

    const [food, setFood] = useState<FoodsInterface[]>([]);
    const [fooditem, setFoodItem] = useState(0);
    const [fooditems, setFoodItemS] = useState(0);
    const [fooditemwant, setFoodItemWant] = useState(0);
    const [fooditemsum, setFoodItemSum] = useState(0);
    const [pricefood, setPriceFood] = useState(0);

    const [drink, setDrink] = useState<DrinksInterface[]>([]);
    const [drinkitem, setDrinkItem] = useState(0);
    const [drinkitems, setDrinkItemS] = useState(0);
    const [drinkitemwant, setDrinkItemWant] = useState(0);
    const [drinkitemsum, setDrinkItemSum] = useState(0);
    const [pricedrink, setPriceDrink] = useState(0);

    const [accessorie, setAccessorie] = useState<StorageInterface[]>([]);
    const [accessorieitem, setAccessorieItem] = useState(0);
    const [accessorieitems, setAccessorieItemS] = useState(0);
    const [accessorieitemwant, setAccessorieItemWant] = useState(0);
    const [accessorieitemsum, setAccessorieItemSum] = useState(0);
    const [priceaccessorie, setPriceAccessorie] = useState(0);

    const [pricetotal, setPriceTotal] = useState(0);

    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const [message, setAlertMessage] = useState("");
    const [room, setRoom] = useState('');
    const id_cus = localStorage.getItem("id");
    const { id } = useParams();
    const status = useRef(true);

    const handleClose = (
        event?: React.SyntheticEvent | Event,
        reason?: string
    ) => {
        if (reason === "clickaway") {
            return;
        }
        setSuccess(false);
        setError(false);
    };

    // for combobox information
    const handleChange = (event: SelectChangeEvent) => {
        const name = event.target.name as keyof typeof service;
        setService({
            ...service,
            [name]: event.target.value,
        });
    };

    const handleInputFoodChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof ServiceUpdate;
        const { value } = event.target;
        const getitemwant = event.target.value;

        setService({ ...service, [id]: value });
        setFoodItemWant(getitemwant)
    };

    const handleInputDrinkChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof ServiceUpdate;
        const { value } = event.target;
        const getitemwant = event.target.value;

        setService({ ...service, [id]: value });
        setDrinkItemWant(getitemwant)
    };

    const handleInputAccessorieChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof ServiceUpdate;
        const { value } = event.target;
        const getitemwant = event.target.value;

        setService({ ...service, [id]: value });
        setAccessorieItemWant(getitemwant)
    };

    const convertType = (data: string | number | undefined | null) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };
    const convertTypeNotNull = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function update() {
        let serviceupdate = {
            ID: convertTypeNotNull(id),
            CustomerID: convertType(id_cus),
            Time: time,
            FoodID: convertTypeNotNull(service.FoodID),
            FoodItem: convertTypeNotNull(service.FoodItem),
            DrinkID: convertTypeNotNull(service.DrinkID),
            DrinkItem: convertTypeNotNull(service.DrinkItem),
            StorageID: convertTypeNotNull(service.StorageID),
            StorageItem: convertTypeNotNull(service.StorageItem),
            Total: pricetotal,
        };
        let foodupdate = {
            ID: convertTypeNotNull(service.FoodID),
            Item: fooditemsum,
        };
        let drinkupdate = {
            ID: convertTypeNotNull(service.DrinkID),
            Item: drinkitemsum,
        };
        let accessoriesupdate = {
            ID: convertTypeNotNull(service.StorageID),
            Quantity: accessorieitemsum,
        };

        if (fooditemsum > 0 && fooditemsum <= 100) {
            if (drinkitemsum > 0 && drinkitemsum <= 100) {
                if (accessorieitemsum > 0 && accessorieitemsum <= 100) {
                    let res = await UpdateService(serviceupdate);
                    if (res.status) {
                        await UpdateFood(foodupdate);
                        await UpdateDrink(drinkupdate);
                        await UpdateAccessories(accessoriesupdate);
                        setAlertMessage("Update Order Successfully");
                        setSuccess(true);
                        setInterval(() => {
                            window.location.assign("/ss");
                        }, 2000);
                    } else {
                        setAlertMessage(res.message);
                        setError(true);
                    }
                } else {
                    setAlertMessage("Accessorie not enough")
                    setError(true);
                }
            } else {
                setAlertMessage("Drink not enough")
                setError(true);
            }
        } else {
            setAlertMessage("Food not enough")
            setError(true);
        }
    }

    const getservicebyid = async () => {
        let res = await GetServiceByID(id);
        if (res) {
            setService(res);
        }
    };

    const getfoods = async () => {
        let res = await GetFoods();
        if (res) {
            setFood(res);
        }
    };
    const getfooditem = async () => {
        let res = await GetFoodItem(service.FoodID);
        if (res) {
            setFoodItem(res);
        }
    };
    const getdrinks = async () => {
        let res = await GetDrinks();
        if (res) {
            setDrink(res);
        }
    };
    const getdrinkitem = async () => {
        let res = await GetDrinkItem(service.DrinkID);
        if (res) {
            setDrinkItem(res);
        }
    };
    const getaccessories = async () => {
        let res = await GetAccessories();
        if (res) {
            setAccessorie(res);
        }
    };
    const getaccessorieitem = async () => {
        let res = await GetAccessorieItem(service.StorageID);
        if (res) {
            setAccessorieItem(res);
        }
    };
    const getroom = async () => {
        let res = await GetRoom(id_cus);
        if (res) {
            setRoom(res);
        }
    };
    const getfooditems = async () => {
        let res = await GetFoodItemS(id);
        if (res) {
            setFoodItemS(res);
            setFoodItemWant(res);
        }
    }
    const getdrinkitems = async () => {
        let res = await GetDrinkItemS(id);
        if (res) {
            setDrinkItemS(res);
            setDrinkItemWant(res);
        }
    }
    const getaccessorieitems = async () => {
        let res = await GetAccessoriesItemS(id);
        if (res) {
            setAccessorieItemS(res);
            setAccessorieItemWant(res);
        }
    }
    const getfoodprice = async () => {
        let res = await GetPriceFood(service.FoodID);
        if (res) {
            setPriceFood(res);
        }
    };
    const getdrinkprice = async () => {
        let res = await GetPriceDrink(service.DrinkID);
        if (res) {
            setPriceDrink(res);
        }
    };
    const getaccessorieprice = async () => {
        let res = await GetPriceAccessorie(service.StorageID);
        if (res) {
            setPriceAccessorie(res);
        }
    };

    useEffect(() => {
        getservicebyid();
        getroom();
        getfoods();
        getdrinks();
        getaccessories();
        getfooditems();
        getdrinkitems();
        getaccessorieitems();
    }, []);

    // when foodid, drinkid, Accessoriesid update
    useEffect(() => {
        if (status.current) {
            status.current = false;
        } else {
            getfooditem();
            setFoodItemSum((fooditem + fooditems) - fooditemwant);
            getdrinkitem();
            setDrinkItemSum((drinkitem + drinkitems) - drinkitemwant);
            getaccessorieitem();
            setAccessorieItemSum((accessorieitem + accessorieitems) - accessorieitemwant);
            getfoodprice();
            getdrinkprice();
            getaccessorieprice();
            setPriceTotal((pricefood * fooditemwant) + (pricedrink * drinkitemwant) + (priceaccessorie * accessorieitemwant));
        }
    });

    console.log(room);
    
    return (
        <div>
            <ThemeProvider theme={theme}>
                <Container
                    maxWidth="lg"
                    sx={{
                        width: "auto",
                        height: "auto",

                    }}>

                    <Snackbar
                        id="success"
                        open={success}
                        autoHideDuration={4000}
                        onClose={handleClose}
                        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
                    >
                        <Alert onClose={handleClose} severity="success">
                            {message}
                        </Alert>
                    </Snackbar>

                    <Snackbar
                        id="error"
                        open={error}
                        autoHideDuration={4000}
                        onClose={handleClose}
                        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
                    >
                        <Alert onClose={handleClose} severity="error">
                            {message}
                        </Alert>
                    </Snackbar>

                    <Paper
                        elevation={3}
                        sx={{
                            bgcolor: "#CDCDCDCD",
                            padding: 2,
                            marginBottom: 2,
                            boxShadow: 1,
                            marginTop: 4,
                        }}
                    >
                        <Paper
                            sx={{
                                bgcolor: "#white",
                                padding: 2,
                                marginBottom: 3,
                                boxShadow: 1,
                                marginTop: 0.5,
                            }}
                        >
                            <Grid
                                container
                                spacing={1}
                                item xs={12}
                            >
                                <Grid item xs={11}>
                                    <Typography
                                        variant="h4"
                                        sx={{
                                            flexGrow: 1,
                                            fontFamily: "Comic Sans MS",
                                        }}
                                    >
                                        Room {room}
                                    </Typography>
                                </Grid>
                                <Grid item xs={1}>
                                    <Typography
                                        variant="h4"
                                        sx={{
                                            flexGrow: 1,
                                            fontFamily: "Comic Sans MS",
                                        }}
                                    >
                                        Bill {id}
                                    </Typography>
                                </Grid>
                            </Grid>
                        </Paper>

                        <Paper
                            sx={{
                                bgcolor: "#white",
                                padding: 2,
                                marginBottom: 1,
                                boxShadow: 1,
                                marginTop: 1,
                            }}
                        >
                            <Typography
                                gutterBottom
                            >
                            </Typography>

                            <Grid
                                container
                                sx={{
                                    flexGrow: 1,
                                    fontFamily: "Comic Sans MS",
                                }}
                            >
                                <Grid
                                    container
                                    justifyContent="center"
                                    gap={14.5}
                                >
                                    <Box>
                                        <img src="https://images.unsplash.com/photo-1627308595186-e6bb36712645?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=735&q=80" height={350} width={250} />
                                    </Box>
                                    <Box>
                                        <img src="https://images.unsplash.com/photo-1502389743708-d00f658638bd?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=687&q=80" height={350} width={250} />
                                    </Box>
                                    <Box>
                                        <img src="https://images.unsplash.com/photo-1558682766-1ff2c50e7056?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=753&q=80" height={350} width={250} />
                                    </Box>
                                </Grid>
                            </Grid>

                            <Typography
                                gutterBottom
                            >
                            </Typography>

                            <Grid
                                container
                                spacing={1}
                                sx={{
                                    padding: 4
                                }}
                            >
                                <Grid
                                    container
                                    justifyContent="center"
                                    gap={10.5}
                                >
                                    <Grid item xs={3.3}>
                                        <FormControl fullWidth variant="outlined">
                                            <Select
                                                native
                                                disabled
                                                sx={{
                                                    fontFamily: "Comic Sans MS",
                                                }}
                                                value={service.FoodID + ""}
                                                onChange={handleChange}
                                                inputProps={{
                                                    name: "FoodID",
                                                }}
                                            >
                                                {food.map((item: FoodsInterface) => (
                                                    <option value={item.ID}>{item.Name}</option>
                                                ))}
                                            </Select>
                                        </FormControl>
                                    </Grid>
                                    <Grid item xs={3.3}>
                                        <FormControl fullWidth variant="outlined">
                                            <Select
                                                native
                                                disabled
                                                sx={{
                                                    fontFamily: "Comic Sans MS",
                                                }}
                                                value={service.DrinkID + ""}
                                                onChange={handleChange}
                                                inputProps={{
                                                    name: "DrinkID",
                                                }}
                                            >
                                                {drink.map((item: DrinksInterface) => (
                                                    <option value={item.ID}>{item.Name}</option>
                                                ))}
                                            </Select>
                                        </FormControl>
                                    </Grid>
                                    <Grid item xs={3.3}>
                                        <FormControl fullWidth variant="outlined">
                                            <Select
                                                native
                                                disabled
                                                sx={{
                                                    fontFamily: "Comic Sans MS",
                                                }}
                                                value={service.StorageID + ""}
                                                onChange={handleChange}
                                                inputProps={{
                                                    name: "StorageID",
                                                }}
                                            >
                                                {accessorie.map((item: StorageInterface) => (
                                                    <option value={item.ID}>{item.Product?.Name}</option>
                                                ))}
                                            </Select>
                                        </FormControl>
                                    </Grid>
                                </Grid>
                            </Grid>

                            <Typography
                                gutterBottom
                            >
                            </Typography>

                            <Grid
                                container
                                justifyContent="center"
                                gap={20}
                            >
                                <Grid item xs={2.2}>
                                    <TextField
                                        InputProps={{
                                            style: { fontFamily: 'Comic Sans MS' },
                                        }}
                                        required
                                        variant="standard"
                                        id="FoodItem"
                                        value={service.FoodItem}
                                        onChange={handleInputFoodChange}
                                    />
                                </Grid>
                                <Grid item xs={2.2}>
                                    <TextField
                                        InputProps={{
                                            style: { fontFamily: 'Comic Sans MS' },
                                        }}
                                        required
                                        variant="standard"
                                        id="DrinkItem"
                                        value={service.DrinkItem}
                                        onChange={handleInputDrinkChange}
                                    />
                                </Grid>
                                <Grid item xs={2.2}>
                                    <TextField
                                        InputProps={{
                                            style: { fontFamily: 'Comic Sans MS' },
                                        }}
                                        required
                                        variant="standard"
                                        id="StorageItem"
                                        value={service.StorageItem}
                                        onChange={handleInputAccessorieChange}
                                    />
                                </Grid>
                            </Grid>

                            <Typography
                                gutterBottom
                            >
                            </Typography>

                            <Grid container spacing={1} sx={{ padding: 3 }}>
                                <Grid item xs={12}>
                                    <Button
                                        sx={{
                                            fontFamily: "Comic Sans MS",
                                        }}
                                        component={RouterLink}
                                        to="/ss"
                                        variant="contained"
                                        color="error"

                                    >
                                        BACK
                                    </Button>

                                    <Button
                                        sx={{
                                            float: "right",
                                            fontFamily: "Comic Sans MS",
                                        }}
                                        onClick={update}
                                        variant="contained"
                                        color="success"
                                    >
                                        UPDATE
                                    </Button>
                                </Grid>
                            </Grid>
                        </Paper>
                    </Paper>
                </Container>
            </ThemeProvider >
        </div >
    );
}

export default ServiceUpdate