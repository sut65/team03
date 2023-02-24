import {
    GetRoom,
    GetFoods,
    GetDrinks,
    AddService,
    UpdateFood,
    UpdateDrink,
    GetFoodItem,
    GetPriceFood,
    GetDrinkItem,
    GetPriceDrink,
    GetAccessories,
    UpdateAccessories,
    GetAccessorieItem,
    GetPriceAccessorie,
} from "./service/ServiceHttpClientService";
import { Button, Container, FormControl, Grid, Paper, Select, SelectChangeEvent, Snackbar, TextField, Typography } from "@mui/material";
import { DrinksInterface, FoodsInterface, ServicesInterface } from "../../models/modelService/IService";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { forwardRef, useEffect, useRef, useState } from "react";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { StorageInterface } from "../../models/IStorage";
import { Link as RouterLink } from "react-router-dom";
import { grey } from "@mui/material/colors";
import { Box } from "@mui/system";

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

function ServiceAdd() {

    const [service, setService] = useState<Partial<ServicesInterface>>({});
    const [time] = useState<Date | null>(new Date());

    const [food, setFood] = useState<FoodsInterface[]>([]);
    const [fooditem, setFoodItem] = useState(0);
    const [fooditemwant, setFoodItemWant] = useState(0);
    const [fooditemsum, setFoodItemSum] = useState(0);
    const [pricefood, setPriceFood] = useState(0);

    const [drink, setDrink] = useState<DrinksInterface[]>([]);
    const [drinkitem, setDrinkItem] = useState(0);
    const [drinkitemwant, setDrinkItemWant] = useState(0);
    const [drinkitemsum, setDrinkItemSum] = useState(0);
    const [pricedrink, setPriceDrink] = useState(0);

    const [accessorie, setAccessorie] = useState<StorageInterface[]>([]);
    const [accessorieitem, setAccessorieItem] = useState(0);
    const [accessorieitemwant, setAccessorieItemWant] = useState(0);
    const [accessorieitemsum, setAccessorieItemSum] = useState(0);
    const [priceaccessorie, setPriceAccessorie] = useState(0);

    const [pricetotal, setPriceTotal] = useState(0);

    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const [message, setAlertMessage] = useState("");
    const [room, setRoom] = useState('');
    const id_cus = localStorage.getItem("id");
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
        const id = event.target.id as keyof typeof ServiceAdd;
        const { value } = event.target;
        const getitemwant = event.target.value;

        setService({ ...service, [id]: value });
        setFoodItemWant(getitemwant)
    };

    const handleInputDrinkChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof ServiceAdd;
        const { value } = event.target;
        const getitemwant = event.target.value;

        setService({ ...service, [id]: value });
        setDrinkItemWant(getitemwant)
    };

    const handleInputAccessorieChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof ServiceAdd;
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

    async function add() {
        let serviceadd = {
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

        if (fooditemsum > 0 && fooditemsum < fooditem) {
            if (drinkitemsum > 0 && drinkitemsum < drinkitem) {
                if (accessorieitemsum > 0 && accessorieitemsum < accessorieitem) {
                    let res = await AddService(serviceadd);
                    if (res.status) {
                        await UpdateFood(foodupdate);
                        await UpdateDrink(drinkupdate);
                        await UpdateAccessories(accessoriesupdate);
                        setAlertMessage("Save Order Successfully");
                        setSuccess(true);
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
        getroom();
        getfoods();
        getdrinks();
        getaccessories();
    }, []);

    // when foodid, drinkid, Accessoriesid update
    useEffect(() => {
        if (status.current) {
            status.current = false;
        } else {
            getfooditem();
            setFoodItemSum(fooditem - fooditemwant);
            getdrinkitem();
            setDrinkItemSum(drinkitem - drinkitemwant);
            getaccessorieitem();
            setAccessorieItemSum(accessorieitem - accessorieitemwant);
            getfoodprice();
            getdrinkprice();
            getaccessorieprice();
            setPriceTotal((pricefood * fooditemwant) + (pricedrink * drinkitemwant) + (priceaccessorie * accessorieitemwant));
        }
    });

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
                            <Box
                                display="flex"
                            >
                                <Typography
                                    variant="h4"
                                    sx={{
                                        flexGrow: 1,
                                        fontFamily: "Comic Sans MS",
                                    }}
                                >
                                    Room {room}
                                </Typography>
                            </Box>
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
                                    padding: 4,
                                }}
                            >
                                <Grid
                                    container
                                    justifyContent="center"
                                    gap={10}
                                >
                                    <Grid item xs={3.3}>
                                        <FormControl fullWidth variant="outlined">
                                            <Select
                                                native
                                                value={service.FoodID + ""}
                                                onChange={handleChange}
                                                inputProps={{
                                                    name: "FoodID",
                                                }}
                                                sx={{
                                                    flexGrow: 1,
                                                    fontFamily: "Comic Sans MS",
                                                }}
                                            >
                                                <option aria-label="None" value="">
                                                    Choose Food
                                                </option>
                                                {food.map((item: FoodsInterface) => (
                                                    <option value={item.ID} key={item.ID}>{item.Name}</option>
                                                ))}
                                            </Select>
                                        </FormControl>
                                    </Grid>
                                    <Grid item xs={3.3}>
                                        <FormControl fullWidth variant="outlined">
                                            <Select
                                                native
                                                value={service.DrinkID + ""}
                                                onChange={handleChange}
                                                inputProps={{
                                                    name: "DrinkID",
                                                }}
                                                sx={{
                                                    flexGrow: 1,
                                                    fontFamily: "Comic Sans MS",
                                                }}
                                            >
                                                <option aria-label="None" value="">
                                                    Choose Drink
                                                </option>
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
                                                value={service.StorageID + ""}
                                                onChange={handleChange}
                                                inputProps={{
                                                    name: "StorageID",
                                                }}
                                                sx={{
                                                    flexGrow: 1,
                                                    fontFamily: "Comic Sans MS",
                                                }}
                                            >
                                                <option aria-label="None" value="">
                                                    Choose Accessories
                                                </option>
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
                                        required
                                        label="How much ?"
                                        InputProps={{
                                            style: { fontFamily: 'Comic Sans MS' },
                                        }}
                                        variant="standard"
                                        id="FoodItem"
                                        value={service.FoodItem}
                                        onChange={handleInputFoodChange}
                                    />
                                </Grid>
                                <Grid item xs={2.2}>
                                    <TextField
                                        required
                                        label="How much ?"
                                        InputProps={{
                                            style: { fontFamily: 'Comic Sans MS' },
                                        }}
                                        variant="standard"
                                        id="DrinkItem"
                                        value={service.DrinkItem}
                                        onChange={handleInputDrinkChange}
                                    />
                                </Grid>
                                <Grid item xs={2.2}>
                                    <TextField
                                        sx={{
                                            fontFamily: "Comic Sans MS",
                                        }}
                                        required
                                        label="How much ?"
                                        InputProps={{
                                            style: { fontFamily: 'Comic Sans MS' },
                                        }}
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
                                        onClick={add}
                                        variant="contained"
                                        color="success"
                                    >
                                        ORDER
                                    </Button>
                                </Grid>
                            </Grid>
                        </Paper>
                    </Paper>
                </Container>
            </ThemeProvider>
        </div >
    );
}

export default ServiceAdd