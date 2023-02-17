import { ButtonGroup, Grid, Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";
import { ServicesInterface } from "../../models/modelService/IService";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { GetService } from "./service/ServiceHttpClientService";
import { Link as RouterLink } from "react-router-dom";
import Container from "@mui/material/Container";
import { useEffect, useState } from "react";
import { grey } from '@mui/material/colors';
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

function ServiceShow() {
    const [service, setService] = useState<ServicesInterface[]>([]);
    const id_cus = localStorage.getItem("id");

    const getservice = async () => {
        let res = await GetService(id_cus + "");
        if (res) {
            setService(res);
        }
    };

    useEffect(() => {
        getservice();
    }, []);

    return (
        <ThemeProvider theme={theme}>
            <Container maxWidth="md">
                <Grid container spacing={1} sx={{ padding: 3 }}>
                    <Grid item xs={12}>
                        <Button
                            component={RouterLink}
                            to="/sa"
                            variant="contained"
                            color="success"
                        >
                            ORDER
                        </Button>

                        <Button
                            style={{ float: "right" }}
                            component={RouterLink}
                            to="/sd"
                            variant="contained"
                            color="error"
                        >
                            CANCLE
                        </Button>
                    </Grid>
                </Grid>

                <div>
                    <Container maxWidth="xl">
                        <div style={{ height: 500, width: "100%", marginTop: "20px" }}>
                            <TableContainer >
                                <Table aria-label="simple table">
                                    <TableHead>
                                        <TableRow>
                                            <TableCell align="center" width="20%">Bill</TableCell>
                                            <TableCell align="center" width="20%">Customer</TableCell>
                                            <TableCell align="center" width="20%">Bill Time</TableCell>
                                            <TableCell align="center" width="20%">Food</TableCell>
                                            <TableCell align="center" width="20%">Item</TableCell>
                                            <TableCell align="center" width="20%">Drink</TableCell>
                                            <TableCell align="center" width="20%">Item</TableCell>
                                            <TableCell align="center" width="20%">Accessorie</TableCell>
                                            <TableCell align="center" width="20%">Item</TableCell>
                                            <TableCell align="center" width="20%">- Edit -</TableCell>
                                        </TableRow>
                                    </TableHead>

                                    <TableBody>
                                        {service.map((item: ServicesInterface) => (
                                            <TableRow key={item.ID}>
                                                <TableCell align="center">{item.ID}</TableCell>
                                                <TableCell align="center">{item.Customer?.FirstName}</TableCell>
                                                <TableCell align="center">{moment(item.Time).format("DD/MM/YYYY HH:mm:ss")}</TableCell>
                                                <TableCell align="center">{item.Food?.Name}</TableCell>
                                                <TableCell align="center">{item.FoodItem}</TableCell>
                                                <TableCell align="center">{item.Drink?.Name}</TableCell>
                                                <TableCell align="center">{item.DrinkItem}</TableCell>
                                                <TableCell align="center">{item.Storage?.Product?.Name}</TableCell>
                                                <TableCell align="center">{item.StorageItem}</TableCell>
                                                <TableCell align="center">
                                                    <ButtonGroup color="primary" aria-label="outlined primary button group">
                                                        <Button
                                                            color="warning"
                                                            component={RouterLink}
                                                            to={`/su/${item.ID}`}
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
            </Container>
        </ThemeProvider>
    );
}

export default ServiceShow