import { Alert, Button, Container, createTheme, FormControl, FormLabel, Grid, Select, SelectChangeEvent, Snackbar, TextField } from "@mui/material";
import { MethodsInterface, PaymentMethodsInterface, PaymentsInterface, PlacesInterface } from "../../models/modelPayment/IPayment";
import { DateTimePicker, LocalizationProvider } from "@mui/x-date-pickers";
import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { ThemeProvider } from "@emotion/react";
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

function PaymentAdd() {

    const [payment, setPayment] = useState<Partial<PaymentsInterface>>({});
    const [time, setTime] = useState<Date | null>(new Date());
    const [paymet, setPaymet] = useState<PaymentMethodsInterface[]>([]);
    const [paymetid, setPaymetid] = useState('');
    const [method, setMethod] = useState<MethodsInterface[]>([]);
    const [destination, setDestination] = useState<string | null>(null);
    const [metid, setMetid] = useState('');
    const [place, setPlace] = useState<PlacesInterface[]>([]);
    const [image, setImage] = useState<string | ArrayBuffer | null>(null);
    const [picture, setPicture] = useState<string | ArrayBuffer | null>(null);
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const id_cus = localStorage.getItem("id");

    const handleImageChange = (event: any) => {
        const image = event.target.files[0];

        const reader = new FileReader();
        reader.readAsDataURL(image);
        reader.onload = () => {
            const base64Data = reader.result;
            setImage(base64Data)
        }
    }

    const handleMet = (event: { target: { name: string; value: any; }; }) => {
        const name = event.target.name as keyof typeof payment;
        const getmetid = event.target.value;

        setMetid(getmetid);
        setPayment({
            ...payment,
            [name]: event.target.value
        });
    };

    const handlePayMet = (event: { target: { name: string; value: any; }; }) => {
        const name = event.target.name as keyof typeof payment;
        const getpaymetid = event.target.value;

        setPaymetid(getpaymetid);
        setPayment({
            ...payment,
            [name]: event.target.value
        });
    };

    const handleChange = (event: SelectChangeEvent<number>) => {
        const name = event.target.name as keyof typeof payment;
        setPayment({
            ...payment,
            [name]: event.target.value
        });
    };

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

    function submit() {
        let data = {
            CustomerID: typeof id_cus === "string" ? parseInt(id_cus) : 0,
            PaymentMethodID: typeof payment.PaymentMethodID === "string" ? parseInt(payment.PaymentMethodID) : 0,
            MethodID: typeof payment.MethodID === "string" ? parseInt(payment.MethodID) : 0,
            PlaceID: typeof payment.PlaceID === "string" ? parseInt(payment.PlaceID) : 0,
            Time: time,
            Picture: image,
        };

        const requestOptions = {
            method: "POST",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        };
        fetch(`http://localhost:8080/payment`, requestOptions)
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
    const fetchPaymentMethods = async () => {
        fetch(`${apiUrl}/paymentmethods`, requestOptions)
            .then(response => response.json())
            .then(res => {
                setPaymet(res.data);
            })
    }
    const fetchPlaces = async () => {
        fetch(`${apiUrl}/places`, requestOptions)
            .then(response => response.json())
            .then(res => {
                setPlace(res.data);
            })
    }
    const fetchMethodP = async () => {
        fetch(`${apiUrl}/methods/paymet/${paymetid}`, requestOptions)
            .then(response => response.json())
            .then(res => {
                setMethod(res.data);
            })
    }

    const getMethod = async () => {
        const apiUrl = `http://localhost:8080/method/${metid}`;
        const requestOptions = {
            method: "GET",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json",
            },
        };

        await fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setDestination(res.data.Destination);
                    setPicture(res.data.Picture);
                }
            });
    };

    useEffect(() => {
        fetchPaymentMethods();
        fetchPlaces();
        fetchMethodP();
        getMethod();
    }, [paymetid, metid]);

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
                            PAYMENT SUCCESS
                        </Alert>
                    </Snackbar>

                    <Snackbar
                        open={error}
                        autoHideDuration={2000}
                        onClose={handleClose}
                        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
                    >
                        <Alert onClose={handleClose} severity="error">
                            PAYMENT FAIL
                        </Alert>
                    </Snackbar>
                    <Grid container
                        direction="row"
                        justifyContent="center"
                        alignItems="stretch">
                        <Grid container spacing={1} sx={{ padding: 3 }} >
                            <Grid item xs={3}>
                                <img src={`${picture}`} width="150" height="150" />
                            </Grid>
                            <Grid
                                container
                                item xs={9}
                                columnGap={3}
                            >
                                <Grid item xs={5.75}>
                                    <FormControl fullWidth variant="outlined">
                                        <Select
                                            native
                                            value={payment.PaymentMethodID}
                                            onChange={handlePayMet}
                                            inputProps={{
                                                name: "PaymentMethodID",
                                            }}
                                        >
                                            <option aria-label="None" value="">
                                                Choose method
                                            </option>
                                            {paymet.map((item: PaymentMethodsInterface) => (
                                                <option value={item.ID} >{item.Name}</option>
                                            ))}
                                        </Select>
                                    </FormControl>
                                </Grid>
                                <Grid item xs={5.75}>
                                    <FormControl fullWidth variant="outlined">
                                        <Select
                                            native
                                            value={payment.MethodID}
                                            onChange={handleMet}
                                            inputProps={{
                                                name: "MethodID",
                                            }}
                                        >
                                            <option aria-label="None" value="">
                                                Choose way
                                            </option>
                                            {method.map((item: MethodsInterface) => (
                                                <option value={item.ID}>{item.Name}</option>
                                            ))}
                                        </Select>
                                    </FormControl>
                                </Grid>
                                <Grid item xs={12}>
                                    <FormControl fullWidth variant="outlined">
                                        <FormLabel> Destination </FormLabel>
                                        <TextField
                                            variant="outlined"
                                            value={destination}
                                            InputProps={{
                                                readOnly: true,
                                            }}
                                        />
                                    </FormControl>
                                </Grid>
                            </Grid>
                            <Grid item xs={6}>
                                <FormControl fullWidth variant="outlined">
                                    <FormLabel> Where </FormLabel>
                                    <Select
                                        native
                                        value={payment.PlaceID}
                                        onChange={handleChange}
                                        inputProps={{
                                            name: "PlaceID",
                                        }}
                                    >
                                        <option aria-label="None" value="">
                                            Choose Place
                                        </option>
                                        {place.map((item: PlacesInterface) => (
                                            <option value={item.ID}>{item.Name}</option>
                                        ))}
                                    </Select>
                                </FormControl>
                            </Grid>
                            <Grid item xs={6}>
                                <FormControl fullWidth variant="outlined">
                                    <FormLabel> When </FormLabel>
                                    <LocalizationProvider dateAdapter={AdapterDateFns}>
                                        <DateTimePicker
                                            renderInput={(props) => <TextField {...props} />}
                                            value={time}
                                            onChange={setTime}
                                        />
                                    </LocalizationProvider>
                                </FormControl>
                            </Grid>
                            <Grid item xs={12}>
                                <FormControl fullWidth variant="outlined">
                                    <FormLabel> Price </FormLabel>
                                    <TextField
                                        variant="outlined"
                                        value={200}
                                        InputProps={{
                                            readOnly: true,
                                        }}
                                    />
                                </FormControl>
                            </Grid>
                        </Grid>
                    </Grid>
                    <Grid container spacing={1} sx={{ padding: 3 }}>
                        <Grid item xs={12}>
                            <img src={`${image}`} width="270" height="270" />
                            <FormControl fullWidth variant="outlined">
                                <FormLabel> Slip </FormLabel>
                                <input type="file" onChange={handleImageChange} />
                            </FormControl>
                        </Grid>
                    </Grid>
                    <Grid container spacing={1} sx={{ padding: 3 }}>
                        <Grid item xs={12}>
                            <Button
                                style={{ float: "right" }}
                                onClick={submit}
                                variant="contained"
                                color="success"
                                component={RouterLink}
                                to="/ps"
                            >
                                COMMIT
                            </Button>
                        </Grid>
                    </Grid>
                </Container>
            </ThemeProvider>
        </div >
    );
}


export default PaymentAdd