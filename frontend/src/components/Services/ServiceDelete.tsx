import { Box, Button, Grid, Paper, Snackbar, styled, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TextField, Typography } from "@mui/material";
import { DeleteService, GetAccessorieItem, GetAccessoriesItemSn, GetRoom, GetService, GetServiceByIDn, UpdateAccessories } from "./service/ServiceHttpClientService";
import { ServicesInterface } from "../../models/modelService/IService";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { forwardRef, useEffect, useRef, useState } from "react";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { tableCellClasses } from "@mui/material/TableCell";
import { Link as RouterLink } from "react-router-dom";
import Container from "@mui/material/Container";
import { grey } from '@mui/material/colors';
import moment from "moment";

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

function ServiceDelete() {

    const [service, setService] = useState<ServicesInterface[]>([]);
    const [services, setServices] = useState<Partial<ServicesInterface>>({});
    const [servicess, setServicess] = useState<Partial<ServicesInterface>>({});

    const [accessorieitem, setAccessorieItem] = useState(0);
    const [accessorieitems, setAccessorieItemS] = useState(0);
    const [accessorieitemsum, setAccessorieItemSum] = useState(0);

    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const [message, setAlertMessage] = useState("");
    const [room, setRoom] = useState('');
    const n_cus = localStorage.getItem("name");
    const id_cus = localStorage.getItem("id");
    const status = useRef(true);

    const TableCellHead = styled(TableCell)(({ theme }) => ({
        [`&.${tableCellClasses.head}`]: {
            backgroundColor: "#808080",
            color: theme.palette.common.white,
            fontFamily: "Comic Sans MS",
            fontSize: 16,
        },
    }));

    const TableCellBody = styled(TableCell)(({ theme }) => ({
        [`&.${tableCellClasses.body}`]: {
            backgroundColor: "#white",
            color: theme.palette.common.black,
            fontFamily: "Comic Sans MS",
            fontSize: 12,
        },
    }));

    const handleSuccess = () => {
        const shouldConfirm = window.confirm('Are you sure you want to cancel the Service?');
        if (shouldConfirm) {
            confirm();
        }
    };

    const handleInputChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof ServiceDelete;
        const { value } = event.target;
        setServices({ ...services, [id]: value });
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

    const convertTypeNotNull = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function confirm() {
        let accessoriesupdate = {
            ID: convertTypeNotNull(servicess.StorageID),
            Quantity: accessorieitemsum,
        };
        let res = await DeleteService(services.ID);
        if (res.status) {
            await UpdateAccessories(accessoriesupdate);
            setAlertMessage("Cancle Order Successfully");
            setSuccess(true);
        } else {
            setAlertMessage(res.message);
            setError(true);
        }
    }

    const getservice = async () => {
        let res = await GetService(id_cus + "");
        if (res) {
            setService(res);
        }
    };
    const getservicebyid = async () => {
        let res = await GetServiceByIDn(services.ID);
        if (res) {
            setServicess(res);
        }
    };
    const getaccessorieitem = async () => {
        let res = await GetAccessorieItem(servicess.StorageID);
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
    const getaccessorieitems = async () => {
        let res = await GetAccessoriesItemSn(services.ID);
        if (res) {
            setAccessorieItemS(res);
        }
    }

    useEffect(() => {
        getservice();
        getroom();
    }, [success]);

    useEffect(() => {
        if (status.current) {
            getservicebyid();
            status.current = false;
        } else {
            getaccessorieitem();
            getaccessorieitems();
            setAccessorieItemSum(accessorieitem + accessorieitems);
            status.current = true;
        }
    });

    return (
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
                            <Grid item xs={10}>
                                <Typography
                                    variant="h4"
                                    sx={{
                                        flexGrow: 1,
                                        fontFamily: "Comic Sans MS",
                                    }}
                                >
                                    {n_cus}'s Bill
                                </Typography>
                            </Grid>
                            <Grid item xs={2}>
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
                        <Container maxWidth="xl">
                            <Box
                                sx={{
                                    padding: 3,
                                }}
                            >
                                <Grid
                                    container
                                    justifyContent="center"
                                    gap={20}
                                >
                                    <Typography
                                        variant="h5"
                                        sx={{
                                            flexGrow: 1,
                                            fontFamily: "Comic Sans MS",
                                        }}
                                        gutterBottom
                                    >
                                        Choose bill number if you want to cancle
                                    </Typography>

                                    <TextField
                                        required
                                        InputProps={{
                                            style: { fontFamily: 'Comic Sans MS' },
                                        }}
                                        style={{ float: "right" }}
                                        label="Type Bill Number"
                                        variant="standard"
                                        id="ID"
                                        value={services.ID || ""}
                                        onChange={handleInputChange}
                                    />
                                </Grid>
                                <Typography
                                    gutterBottom
                                >
                                </Typography>
                            </Box>

                            <Container maxWidth="xl">
                                <TableContainer component={Paper}>
                                    <Table sx={{ minWidth: 500 }}>
                                        <TableHead>
                                            <TableRow>
                                                <TableCellHead align="center" width="20%" sx={{ border: 1 }}> Bill Number </TableCellHead>
                                                <TableCellHead align="center" width="20%" sx={{ border: 1 }}> Customer name </TableCellHead>
                                                <TableCellHead align="center" width="20%" sx={{ border: 1 }}> Bill Time </TableCellHead>
                                                <TableCellHead align="center" width="20%" sx={{ border: 1 }}> Food </TableCellHead>
                                                <TableCellHead align="center" width="20%" sx={{ border: 1 }}>  Item </TableCellHead>
                                                <TableCellHead align="center" width="20%" sx={{ border: 1 }}> Drink </TableCellHead>
                                                <TableCellHead align="center" width="20%" sx={{ border: 1 }}>  Item </TableCellHead>
                                                <TableCellHead align="center" width="20%" sx={{ border: 1 }}> Accessorie </TableCellHead>
                                                <TableCellHead align="center" width="20%" sx={{ border: 1 }}>  Item </TableCellHead>
                                                <TableCellHead align="center" width="20%" sx={{ border: 1 }}> Total </TableCellHead>
                                            </TableRow>
                                        </TableHead>

                                        <TableBody>
                                            {service.map((item: ServicesInterface) => (
                                                <TableRow key={item.ID}>
                                                    <TableCellBody align="center">{item.ID}</TableCellBody>
                                                    <TableCellBody align="center">{item.Customer?.FirstName}</TableCellBody>
                                                    <TableCellBody align="center">{moment(item.Time).format("DD/MM/YYYY HH:mm:ss")}</TableCellBody>
                                                    <TableCellBody align="center">{item.Food?.Name}</TableCellBody>
                                                    <TableCellBody align="center">{item.FoodItem}</TableCellBody>
                                                    <TableCellBody align="center">{item.Drink?.Name}</TableCellBody>
                                                    <TableCellBody align="center">{item.DrinkItem}</TableCellBody>
                                                    <TableCellBody align="center">{item.Storage?.Product?.Name}</TableCellBody>
                                                    <TableCellBody align="center">{item.StorageItem}</TableCellBody>
                                                    <TableCellBody align="center">{item.Total}</TableCellBody>
                                                </TableRow>
                                            ))}
                                        </TableBody>
                                    </Table>
                                </TableContainer>
                            </Container>
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
                                        onClick={handleSuccess}
                                        variant="contained"
                                        color="success"
                                    >
                                        CANCLE
                                    </Button>
                                </Grid>
                            </Grid>
                        </Container>
                    </Paper>
                </Paper>
            </Container>
        </ThemeProvider >
    );
}

export default ServiceDelete