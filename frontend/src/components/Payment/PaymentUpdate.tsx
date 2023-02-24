import { Button, Container, createTheme, FormControl, FormLabel, Grid, Paper, Select, SelectChangeEvent, Snackbar, TextField, Typography } from "@mui/material";
import { GetPlaces, GetMethodP, GetPicture, UpdatePayment, GetDestination, GetPaymentByID, GetPaymentMethods } from "./service/PaymentHttpClientService";
import { MethodsInterface, PaymentMethodsInterface, PaymentsInterface, PlacesInterface } from "../../models/modelPayment/IPayment";
import { DateTimePicker, LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { Link as RouterLink, useParams } from "react-router-dom";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { forwardRef, useEffect, useState } from "react";
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

const Alert = forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function PaymentUpdate() {

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
    const { id } = useParams();

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

    const convertTypeNotNull = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function submit() {
        let data = {
            ID: convertTypeNotNull(id),
            CustomerID: convertType(id_cus),
            PaymentMethodID: convertType(payment.PaymentMethodID),
            MethodID: convertType(payment.MethodID),
            PlaceID: convertType(payment.PlaceID),
            Time: time,
            Picture: image,
        };

        let res = await UpdatePayment(data);
        if (res.status) {
            setAlertMessage("Update Payment Successfully");
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
            console.log();

        }
    };
    const getpaymentbyid = async () => {
        let res = await GetPaymentByID(id);
        if (res) {
            setPayment(res);
            setPaymetid(res.PaymentMethodID)
            setMetid(res.MethodID)
            setImage(res.Picture)
        }
    };

    useEffect(() => {
        getpaymentMethods();
        getmethodp();
        getdestination();
        getpicture();
        getplaces();
    }, [paymetid, metid]);

    useEffect(() => {
        getpaymentbyid();
    }, [])

    return (

        <div>
            <ThemeProvider theme={theme}>
                <Container
                    maxWidth="lg"
                    sx={{
                        width: "auto",
                        height: "auto",
                        fontFamily: "Comic Sans MS",
                    }}>
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
                                marginBottom: 1,
                                boxShadow: 1,
                                marginTop: 0.5,
                            }}
                        >
                            <Grid
                                container
                                direction="row"
                                justifyContent="center"
                                alignItems="stretch">
                                <Grid
                                    container
                                    sx={{
                                        flexGrow: 1,
                                    }}>
                                    <Grid
                                        item xs={3}
                                        sx={{
                                            paddingLeft: 7
                                        }}
                                    >
                                        <img src={`${picture}`} width="150" height="150" />
                                    </Grid>
                                    <Grid
                                        container
                                        item xs={9}
                                        columnGap={3}
                                    >
                                        <Grid item xs={5.82}>
                                            <FormControl fullWidth variant="outlined">
                                                <Select
                                                    native
                                                    sx={{
                                                        fontFamily: "Comic Sans MS",
                                                    }}
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
                                        <Grid item xs={5.82}>
                                            <FormControl fullWidth variant="outlined">
                                                <Select
                                                    native
                                                    sx={{
                                                        fontFamily: "Comic Sans MS",
                                                    }}
                                                    value={payment.MethodID + ""}
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
                                                <FormLabel sx={{ fontFamily: "Comic Sans MS" }}> Destination </FormLabel>
                                                <TextField
                                                    InputProps={{
                                                        style: { fontFamily: 'Comic Sans MS' },
                                                        readOnly: true,
                                                    }}
                                                    variant="outlined"
                                                    value={destination}
                                                />
                                            </FormControl>
                                        </Grid>
                                    </Grid>
                                    <Grid
                                        container
                                        spacing={1}
                                        sx={{
                                            padding: 1
                                        }}
                                    >
                                        <Grid
                                            container
                                            gap={3.5}
                                        >
                                            <Grid item xs={5.8}
                                                sx={{
                                                    paddingLeft: 1
                                                }}>
                                                <FormControl fullWidth variant="outlined">
                                                    <FormLabel sx={{ fontFamily: "Comic Sans MS" }}> Where </FormLabel>
                                                    <Select
                                                        native
                                                        sx={{
                                                            fontFamily: "Comic Sans MS",
                                                        }}
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
                                            <Grid
                                                item xs={5.8}>
                                                <FormControl fullWidth variant="outlined">
                                                    <FormLabel sx={{ fontFamily: "Comic Sans MS" }}> When </FormLabel>
                                                    <LocalizationProvider dateAdapter={AdapterDateFns}>
                                                        <DateTimePicker
                                                            value={payment.Time}
                                                            renderInput={(props) => <TextField {...props} />}
                                                            onChange={setTime}

                                                        />
                                                    </LocalizationProvider>
                                                </FormControl>
                                            </Grid>
                                        </Grid>
                                    </Grid>
                                    <Grid
                                        container
                                        spacing={1}
                                        sx={{
                                            padding: 2
                                        }}
                                    >
                                        <FormControl fullWidth variant="outlined">
                                            <FormLabel sx={{ fontFamily: "Comic Sans MS" }}> Price </FormLabel>
                                            <TextField
                                                variant="outlined"
                                                value={payment.Price}
                                                InputProps={{
                                                    style: { fontFamily: 'Comic Sans MS' },
                                                    readOnly: true,
                                                }}
                                            />
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
                                    <img src={`${image}`} width="140" height="140" />
                                    <FormControl fullWidth variant="outlined">
                                        <FormLabel sx={{ fontFamily: "Comic Sans MS" }} > Slip </FormLabel>
                                        <input type="file" onChange={handleImageChange} />
                                    </FormControl>
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
                                        to="/ps"
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
                                        onClick={submit}
                                        variant="contained"
                                        color="success"
                                    >
                                        COMMIT
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


export default PaymentUpdate