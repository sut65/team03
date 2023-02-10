import {
    GetRoom,
    GetFoods,
    GetDrinks,
    AddService,
    UpdateFood,
    UpdateDrink,
    GetFoodItem,
    GetDrinkItem,
    GetAccessories,
    UpdateAccessories,
    GetAccessorieItem,
} from "./service/ServiceHttpClientService";
import { AccessoriesInterface, DrinksInterface, FoodsInterface, ServicesInterface } from "../../models/modelService/IService";
import { Button, Container, FormControl, Grid, Select, SelectChangeEvent, Snackbar, TextField } from "@mui/material";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { forwardRef, useEffect, useRef, useState } from "react";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { Link as RouterLink } from "react-router-dom";
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

function ServiceAdd() {

    const [service, setService] = useState<Partial<ServicesInterface>>({});
    const [time] = useState<Date | null>(new Date());

    const [food, setFood] = useState<FoodsInterface[]>([]);
    const [fooditem, setFoodItem] = useState(0);
    const [fooditemwant, setFoodItemWant] = useState(0);
    const [fooditemsum, setFoodItemSum] = useState(0);

    const [drink, setDrink] = useState<DrinksInterface[]>([]);
    const [drinkitem, setDrinkItem] = useState(0);
    const [drinkitemwant, setDrinkItemWant] = useState(0);
    const [drinkitemsum, setDrinkItemSum] = useState(0);

    const [accessorie, setAccessorie] = useState<AccessoriesInterface[]>([]);
    const [accessorieitem, setAccessorieItem] = useState(0);
    const [accessorieitemwant, setAccessorieItemWant] = useState(0);
    const [accessorieitemsum, setAccessorieItemSum] = useState(0);

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
            AccessoriesID: convertTypeNotNull(service.AccessoriesID),
            AccessoriesItem: convertTypeNotNull(service.AccessoriesItem),
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
            ID: convertTypeNotNull(service.AccessoriesID),
            Item: accessorieitemsum,
        };

        if (fooditemsum > 0) {
            if (drinkitemsum > 0) {
                if (accessorieitemsum > 0) {
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
                    setAlertMessage(`Accessories not enough now item have ${accessorieitem - 1}`);
                    setError(true);
                }
            } else {
                setAlertMessage(`Drink not enough now item have ${drinkitem - 1}`);
                setError(true);
            }
        } else {
            setAlertMessage(`Food not enough now item have ${fooditem - 1}`);
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
        let res = await GetAccessorieItem(service.AccessoriesID);
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
            setFoodItemSum(fooditem - fooditemwant)
            getdrinkitem();
            setDrinkItemSum(drinkitem - drinkitemwant)
            getaccessorieitem();
            setAccessorieItemSum(accessorieitem - accessorieitemwant)
        }
    });

    return (
        <div>
            <ThemeProvider theme={theme}>
                <Container maxWidth="md">
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

                    <Grid container spacing={1} sx={{ padding: 3 }}>
                        <Grid item xs={12}>
                            <TextField
                                fullWidth
                                label="Room Number"
                                value={room}
                                InputProps={{
                                    readOnly: true,
                                }}
                                variant="standard"
                            />
                        </Grid>
                    </Grid>

                    <Grid container margin={2} columnGap={4}>
                        <Grid>
                            <img src="https://images.unsplash.com/photo-1627308595186-e6bb36712645?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=735&q=80" alt="" width="250" height="350" />
                        </Grid>
                        <Grid>
                            <img src="https://images.unsplash.com/photo-1502389743708-d00f658638bd?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=687&q=80" alt="" width="250" height="350" />
                        </Grid>
                        <Grid>
                            <img src="https://images.unsplash.com/photo-1558682766-1ff2c50e7056?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=753&q=80" alt="" width="250" height="350" />
                        </Grid>
                    </Grid>

                    <Grid container spacing={1} sx={{ padding: 3 }}>
                        <Grid item xs={4}>
                            <TextField
                                label=" "
                                defaultValue="If you want to eat."
                                variant="standard"
                                InputProps={{
                                    readOnly: true,
                                }}
                            />
                        </Grid>
                        <Grid item xs={4}>
                            <TextField
                                label=" "
                                defaultValue="If you want to drink."
                                variant="standard"
                                InputProps={{
                                    readOnly: true,
                                }}
                            />
                        </Grid>
                        <Grid item xs={4}>
                            <TextField
                                label=" "
                                defaultValue="If you want to borrow."
                                variant="standard"
                                InputProps={{
                                    readOnly: true,
                                }}
                            />
                        </Grid>
                    </Grid>

                    <Grid container spacing={1} sx={{ padding: 3 }}>
                        <Grid item xs={4}>
                            <FormControl fullWidth variant="outlined">
                                <Select
                                    native
                                    value={service.FoodID + ""}
                                    onChange={handleChange}
                                    inputProps={{
                                        name: "FoodID",
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
                        <Grid item xs={4}>
                            <FormControl fullWidth variant="outlined">
                                <Select
                                    native
                                    value={service.DrinkID + ""}
                                    onChange={handleChange}
                                    inputProps={{
                                        name: "DrinkID",
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
                        <Grid item xs={4}>
                            <FormControl fullWidth variant="outlined">
                                <Select
                                    native
                                    value={service.AccessoriesID + ""}
                                    onChange={handleChange}
                                    inputProps={{
                                        name: "AccessoriesID",
                                    }}
                                >
                                    <option aria-label="None" value="">
                                        Choose Accessories
                                    </option>
                                    {accessorie.map((item: AccessoriesInterface) => (
                                        <option value={item.ID}>{item.Name}</option>
                                    ))}
                                </Select>
                            </FormControl>
                        </Grid>
                    </Grid>

                    <Grid container spacing={1} sx={{ padding: 3 }}>
                        <Grid item xs={4}>
                            <TextField
                                required
                                type='number'
                                label="How much ?"
                                variant="standard"
                                id="FoodItem"
                                value={service.FoodItem}
                                onChange={handleInputFoodChange}
                            />
                        </Grid>
                        <Grid item xs={4}>
                            <TextField
                                required
                                type='number'
                                label="How much ?"
                                variant="standard"
                                id="DrinkItem"
                                value={service.DrinkItem}
                                onChange={handleInputDrinkChange}
                            />
                        </Grid>
                        <Grid item xs={4}>
                            <TextField
                                required
                                type='number'
                                label="How much ?"
                                variant="standard"
                                id="AccessoriesItem"
                                value={service.AccessoriesItem}
                                onChange={handleInputAccessorieChange}
                            />
                        </Grid>
                    </Grid>

                    <Grid container spacing={1} sx={{ padding: 3 }}>
                        <Grid item xs={12}>
                            <Button
                                component={RouterLink}
                                to="/ss"
                                variant="contained"
                                color="error"
                            >
                                BACK
                            </Button>

                            <Button
                                style={{ float: "right" }}
                                onClick={add}
                                variant="contained"
                                color="success"
                            >
                                CONFIRM
                            </Button>
                        </Grid>
                    </Grid>
                </Container>
            </ThemeProvider>
        </div>
    );
}

export default ServiceAdd