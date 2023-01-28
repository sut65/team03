import { Grid, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TextField } from "@mui/material";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { PaymentsInterface } from "../../models/modelPayment/IPayment";
import { grey } from '@mui/material/colors';
import Container from "@mui/material/Container";
import Button from "@mui/material/Button";
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


function PaymentShow() {

    const [payment, setPayment] = useState<PaymentsInterface[]>([]);

    const id_cus = localStorage.getItem("uid");


    const getServices = async () => {
        const apiUrl = `http://localhost:8080/payment/customer/${id_cus}`;
        const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        };

        await fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setPayment(res.data);
                }
            });
    };

    useEffect(() => {
        getServices();
    }, []);

    return (
        <ThemeProvider theme={theme}>
            <Container maxWidth="md">
                <Grid container spacing={1} sx={{ padding: 3 }}>
                    <Grid item xs={12}>
                        <TextField
                            style={{ float: "right" }}
                            label=" "
                            defaultValue="This is your bill."
                            variant="standard"
                            InputProps={{
                                readOnly: true,
                            }}
                        />
                    </Grid>
                </Grid>

                <div>
                    <Container maxWidth="md">
                        <div style={{ height: 500, width: "100%", marginTop: "20px" }}>
                            <TableContainer >
                                <Table aria-label="simple table">
                                    <TableHead>
                                        <TableRow>
                                            <TableCell align="center" width="20%"> Payment </TableCell>
                                            <TableCell align="center" width="20%"> Customer name </TableCell>
                                            <TableCell align="center" width="20%"> Bill Time </TableCell>
                                            <TableCell align="center" width="20%"> Method </TableCell>
                                            <TableCell align="center" width="20%"> Name </TableCell>
                                            <TableCell align="center" width="20%"> Place </TableCell>
                                            <TableCell align="center" width="20%"> Slip </TableCell>
                                        </TableRow>
                                    </TableHead>

                                    <TableBody>
                                        {payment.map((item: PaymentsInterface) => (
                                            <TableRow key={item.ID}>
                                                <TableCell align="center">{item.ID}</TableCell>
                                                <TableCell align="center">{item.Customer.FirstName}</TableCell>
                                                <TableCell align="center">{moment(item.Time).format("DD/MM/YYYY HH:mm:ss")}</TableCell>
                                                <TableCell align="center">{item.PaymentMethod.Name}</TableCell>
                                                <TableCell align="center">{item.Method.Name}</TableCell>
                                                <TableCell align="center">{item.Place.Name}</TableCell>
                                                <TableCell align="center"><img src={`${item.Picture}`} width="75" height="90" /></TableCell>
                                            </TableRow>
                                        ))}
                                    </TableBody>
                                </Table>
                            </TableContainer>
                        </div>
                    </Container>
                </div>
            </Container>
        </ThemeProvider>
    );
}

export default PaymentShow