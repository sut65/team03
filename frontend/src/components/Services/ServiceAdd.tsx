import { AccessoriesInterface, DrinksInterface, FoodsInterface, ServicesInterface } from "../../models/modelService/IService";
import { Alert, Button, Container, FormControl, Grid, Select, SelectChangeEvent, Snackbar, TextField } from "@mui/material";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { useEffect, useState } from "react";
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

function ServiceAdd() {

    const [service, setService] = useState<Partial<ServicesInterface>>({});
    const [time] = useState<Date | null>(new Date());
    const [food, setFood] = useState<FoodsInterface[]>([]);
    const [drink, setDrink] = useState<DrinksInterface[]>([]);
    const [Accessorie, setAccessorie] = useState<AccessoriesInterface[]>([]);
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const id_cus = localStorage.getItem("id");

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
    const handleChange = (
        event: SelectChangeEvent<number>
    ) => {
        const name = event.target.name as keyof typeof service;
        setService({
            ...service,
            [name]: event.target.value
        });
    };

    function submit() {
        let data = {
            CustomerID: typeof id_cus === "string" ? parseInt(id_cus) : 0,
            Time: time,
            FoodID: typeof service.FoodID === "string" ? parseInt(service.FoodID) : 0,
            DrinkID: typeof service.DrinkID === "string" ? parseInt(service.DrinkID) : 0,
            AccessoriesID: typeof service.AccessoriesID === "string" ? parseInt(service.AccessoriesID) : 0,
        };


        const requestOptions = {
            method: "POST",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        };
        fetch(`http://localhost:8080/service`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setSuccess(true);
                } else {
                    setError(true);
                }
            });
    }

    const apiUrl = "http://localhost:8080";
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };
    const fetchFoods = async () => {
        fetch(`${apiUrl}/foods`, requestOptions)
            .then(response => response.json())
            .then(res => {
                setFood(res.data);
            })
    }
    const fetchDrinks = async () => {
        fetch(`${apiUrl}/drinks`, requestOptions)
            .then(response => response.json())
            .then(res => {
                setDrink(res.data);
            })
    }
    const fetchAccessories = async () => {
        fetch(`${apiUrl}/accessories`, requestOptions)
            .then(response => response.json())
            .then(res => {
                setAccessorie(res.data);
            })
    }

    useEffect(() => {
        fetchFoods();
        fetchDrinks();
        fetchAccessories();
    }, []);

    return (
        <div>
            <ThemeProvider theme={theme}>
                <Container maxWidth="md">
                    <Snackbar
                        open={success}
                        autoHideDuration={2000}
                        onClose={handleClose}
                        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
                    >
                        <Alert onClose={handleClose} severity="success">
                            ORDER SUCCESS
                        </Alert>
                    </Snackbar>

                    <Snackbar
                        open={error}
                        autoHideDuration={2000}
                        onClose={handleClose}
                        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
                    >
                        <Alert onClose={handleClose} severity="error">
                            ORDER FAIL
                        </Alert>
                    </Snackbar>

                    <Grid container spacing={1} sx={{ padding: 3 }}>
                        <Grid item xs={12}>
                            <TextField
                                fullWidth
                                label="Room Number"
                                value={404}
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
                                    value={service.FoodID}
                                    onChange={handleChange}
                                    inputProps={{
                                        name: "FoodID",
                                    }}
                                >
                                    <option aria-label="None" value="">
                                        Choose Food
                                    </option>
                                    {food.map((item: FoodsInterface) => (
                                        <option value={item.ID}>{item.Name}</option>
                                    ))}
                                </Select>
                            </FormControl>
                        </Grid>
                        <Grid item xs={4}>
                            <FormControl fullWidth variant="outlined">
                                <Select
                                    native
                                    value={service.DrinkID}
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
                                    value={service.AccessoriesID}
                                    onChange={handleChange}
                                    inputProps={{
                                        name: "AccessoriesID",
                                    }}
                                >
                                    <option aria-label="None" value="">
                                        Choose Accessories
                                    </option>
                                    {Accessorie.map((item: AccessoriesInterface) => (
                                        <option value={item.ID}>{item.Name}</option>
                                    ))}
                                </Select>
                            </FormControl>
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
                                onClick={submit}
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