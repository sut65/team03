import { Button, Container, createTheme, FormControl, FormLabel, Grid, Select, SelectChangeEvent, Snackbar, TextField } from "@mui/material";
import { AddPayment, GetDestination, GetMethodP, GetPaymentMethods, GetPicture, GetPlaces, GetPriceServiceCID } from "./service/PaymentHttpClientService";
import { MethodsInterface, PaymentMethodsInterface, PaymentsInterface, PlacesInterface } from "../../models/modelPayment/IPayment";
import { DateTimePicker, LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { forwardRef, useEffect, useRef, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { ThemeProvider } from "@emotion/react";
import { grey } from "@mui/material/colors";
import { ServicesInterface } from "../../models/modelService/IService";
import { GetService } from "../Services/service/ServiceHttpClientService";


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

function PaymentAddOut() {

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
    const [message, setAlertMessage] = useState("");
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const id_cus = localStorage.getItem("id");

    const [price, setPrice] = useState(0);

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

    const handleChange = (event: SelectChangeEvent) => {
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

    const convertType = (data: string | number | undefined | null) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function submit() {
        let data = {
            CustomerID: convertType(id_cus),
            PaymentMethodID: convertType(payment.PaymentMethodID),
            MethodID: convertType(payment.MethodID),
            PlaceID: convertType(payment.PlaceID),
            Price: price,
            Time: time,
            Picture: image,
        };

        console.log(data);
        
        let res = await AddPayment(data);
        if (res.status) {
            setAlertMessage("Save Payment Successfully");
            setSuccess(true);
        } else {
            setAlertMessage(res.message);
            setError(true);
        }
    }

    const getpaymentMethods = async () => {
        let res = await GetPaymentMethods();
        if (res) {
            setPaymet(res);
        }
    };
    const getplaces = async () => {
        let res = await GetPlaces();
        if (res) {
            setPlace(res);
        }
    };
    const getmethodp = async () => {
        let res = await GetMethodP(paymetid);
        if (res) {
            setMethod(res);
        }
    };
    const getdestination = async () => {
        let res = await GetDestination(metid);
        if (res) {
            setDestination(res);
        }
    };
    const getpicture = async () => {
        let res = await GetPicture(metid);
        if (res) {
            setPicture(res);
        }
    };
    const getpriceofroom = async () => {
        let res = await GetPriceServiceCID(id_cus);
        if (res) {
            setPrice(res);
        }
    };

    useEffect(() => {
        getpaymentMethods();
        getmethodp();
        getdestination();
        getpicture();
        getplaces();
        getpriceofroom();
    }, [paymetid, metid]);

    return (

        <div>
            <ThemeProvider theme={theme}>
                <Container maxWidth="md">
                    <Snackbar
                        id="success"
                        open={success}
                        autoHideDuration={2000}
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
                        autoHideDuration={2000}
                        onClose={handleClose}
                        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
                    >
                        <Alert onClose={handleClose} severity="error">
                            {message}
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
                                            value={payment.PaymentMethodID + ""}
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
                                        value={payment.PlaceID + ""}
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
                                        value={price}
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
                                component={RouterLink}
                                to="/ps"
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
                                COMMIT
                            </Button>
                        </Grid>
                    </Grid>
                </Container>
            </ThemeProvider>
        </div >
    );
}


export default PaymentAddOut