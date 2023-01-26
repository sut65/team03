import { Alert, Button, FormControl, Grid, Snackbar, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TextField } from "@mui/material";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { ServicesInterface } from "../../models/modelService/IService";
import { grey } from '@mui/material/colors';
import Container from "@mui/material/Container";
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


function ServiceDelete() {

    const [service, setService] = useState<ServicesInterface[]>([]);
    const [services, setServices] = useState<Partial<ServicesInterface>>({});
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const n_cus = localStorage.getItem("name");
    const id_cus = localStorage.getItem("id");

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

    function submit() {
        let data = {
            ID: services.ID ?? "",
        };

        const requestOptions = {
            method: "DELETE",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        };
        fetch(`http://localhost:8080/services/${services.ID}`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setSuccess(true);
                } else {
                    setError(true);
                }
            });
    }

    const getServices = async () => {
        const apiUrl = `http://localhost:8080/services/customer/${id_cus}`;
        const requestOptions = {
            method: "GET",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json",
            },
        };

        fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setService(res.data);
                }
            });
    };

    useEffect(() => {
        getServices();
    }, [success]);

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
                            DELETE SUCCESS
                        </Alert>
                    </Snackbar>

                    <Snackbar
                        open={error}
                        autoHideDuration={2000}
                        onClose={handleClose}
                        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
                    >
                        <Alert onClose={handleClose} severity="error">
                            DELETE FAIL
                        </Alert>
                    </Snackbar>
                    <Grid container spacing={2} rowGap={3}>
                        <Grid container columnGap={58.3}>
                            <Grid>
                                <FormControl>
                                    <TextField
                                        label="Customer Name"
                                        value={n_cus}
                                        InputProps={{
                                            readOnly: true,
                                        }}
                                        variant="standard"
                                    />
                                </FormControl>
                            </Grid>
                            <Grid>
                                <TextField
                                    label="Room Number"
                                    value={404}
                                    InputProps={{
                                        readOnly: true,
                                    }}
                                    variant="standard"
                                />
                            </Grid>
                        </Grid>

                        <Grid container columnGap={42}>
                            <Grid xs={4.5}>
                                <TextField
                                    fullWidth
                                    label=" "
                                    defaultValue="Choose bill number if you want to cancle."
                                    variant="standard"
                                    InputProps={{
                                        readOnly: true,
                                    }}
                                />
                            </Grid>
                            <Grid>
                                <TextField
                                    required
                                    label="Type Bill Number"
                                    variant="standard"
                                    id="ID"
                                    value={services.ID || ""}
                                    onChange={handleInputChange}
                                />
                            </Grid>
                        </Grid>
                    </Grid>

                    <div>
                        <Container maxWidth="md">
                            <div style={{ height: 500, width: "100%", marginTop: "20px" }}>
                                <TableContainer >
                                    <Table aria-label="simple table">
                                        <TableHead>
                                            <TableRow>
                                                <TableCell align="center" width="20%"> Bill Number </TableCell>
                                                <TableCell align="center" width="20%"> Customer name </TableCell>
                                                <TableCell align="center" width="20%"> Bill Time </TableCell>
                                                <TableCell align="center" width="20%"> Food </TableCell>
                                                <TableCell align="center" width="20%"> Drink </TableCell>
                                                <TableCell align="center" width="20%"> Accessorie </TableCell>
                                            </TableRow>
                                        </TableHead>

                                        <TableBody>
                                            {service.map((item: ServicesInterface) => (
                                                <TableRow key={item.ID}>
                                                    <TableCell align="center">{item.ID}</TableCell>
                                                    <TableCell align="center">{item.Customer.FirstName}</TableCell>
                                                    <TableCell align="center">{moment(item.Time).format("DD/MM/YYYY")}</TableCell>
                                                    <TableCell align="center">{item.Food.Name}</TableCell>
                                                    <TableCell align="center">{item.Drink.Name}</TableCell>
                                                    <TableCell align="center">{item.Accessories.Name}</TableCell>
                                                </TableRow>
                                            ))}
                                        </TableBody>
                                    </Table>
                                </TableContainer>
                            </div>
                        </Container>
                    </div>
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

export default ServiceDelete