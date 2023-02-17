import { ButtonGroup, Grid, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TextField } from "@mui/material";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { PaymentsInterface } from "../../models/modelPayment/IPayment";
import { grey } from '@mui/material/colors';
import { GetPayment } from "./service/PaymentHttpClientService";
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
    const id_cus = localStorage.getItem("id");

    const getpayment = async () => {
        let res = await GetPayment(id_cus + "");
        if (res) {
            setPayment(res);
        }
    };

    useEffect(() => {
        getpayment();
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
                                            <TableCell align="center" width="20%"> Slip Time </TableCell>
                                            <TableCell align="center" width="20%"> Method </TableCell>
                                            <TableCell align="center" width="20%"> Name </TableCell>
                                            <TableCell align="center" width="20%"> Place </TableCell>
                                            <TableCell align="center" width="20%"> Price </TableCell>
                                            <TableCell align="center" width="20%"> Slip </TableCell>
                                            <TableCell align="center" width="20%"> - Edit - </TableCell>
                                        </TableRow>
                                    </TableHead>

                                    <TableBody>
                                        {payment.map((item: PaymentsInterface) => (
                                            <TableRow key={item.ID}>
                                                <TableCell align="center">{item.ID}</TableCell>
                                                <TableCell align="center">{item.Customer?.FirstName}</TableCell>
                                                <TableCell align="center">{moment(item.Time).format("DD/MM/YYYY HH:mm:ss")}</TableCell>
                                                <TableCell align="center">{item.PaymentMethod?.Name}</TableCell>
                                                <TableCell align="center">{item.Method?.Name}</TableCell>
                                                <TableCell align="center">{item.Place?.Name}</TableCell>
                                                <TableCell align="center">{item.Price}</TableCell>
                                                <TableCell align="center"><img src={`${item.Picture}`} width="75" height="90" /></TableCell>
                                                <TableCell align="center">
                                                    <ButtonGroup color="primary" aria-label="outlined primary button group">
                                                        <Button
                                                            color="warning"
                                                            component={RouterLink}
                                                            to={`/pu/${item.ID}`}
                                                        >
                                                            Edit</Button>
                                                    </ButtonGroup>
                                                </TableCell>
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
                            to="/pai"
                            variant="contained"
                            color="success"
                        >
                            PAYMENT CHECKIN
                        </Button>

                        <Button
                            style={{ float: "right" }}
                            component={RouterLink}
                            to="/pao"
                            variant="contained"
                            color="success"
                        >
                            PAYMENT CHECKOUT
                        </Button>
                    </Grid>
                </Grid>
            </Container>
        </ThemeProvider>
    );
}

export default PaymentShow