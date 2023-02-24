import {
    Box,
    Grid,
    Table,
    Paper,
    styled,
    TableRow,
    TableHead,
    TableCell,
    TableBody,
    Typography,
    ButtonGroup,
    TableContainer,
} from "@mui/material";
import { ServicesInterface } from "../../models/modelService/IService";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { GetService } from "./service/ServiceHttpClientService";
import { tableCellClasses } from "@mui/material/TableCell";
import { Link as RouterLink } from "react-router-dom";
import EditIcon from '@mui/icons-material/Edit';
import Container from "@mui/material/Container";
import { grey } from '@mui/material/colors';
import { useEffect, useState } from "react";
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
    const name = localStorage.getItem("name");

    const getservice = async () => {
        let res = await GetService(id_cus + "");
        if (res) {
            setService(res);
        }
    };

    const TableCellHead = styled(TableCell)(({ theme }) => ({
        [`&.${tableCellClasses.head}`]: {
            backgroundColor: "#808080",
            color: theme.palette.common.white,
            fontFamily: "Comic Sans MS",
            fontSize: 16,
        },
    }));
    const TableCellHeadY = styled(TableCell)(({ theme }) => ({
        [`&.${tableCellClasses.head}`]: {
            backgroundColor: "#FF7F00",
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

    useEffect(() => {
        getservice();
    }, []);

    return (
        <ThemeProvider theme={theme}>
            <Container
                maxWidth="lg"
                sx={{
                    width: "auto",
                    height: "auto",

                }}>
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
                        elevation={3}
                        sx={{
                            bgcolor: "#white",
                            padding: 2,
                            marginBottom: 3,
                            boxShadow: 1,
                            marginTop: 0.5,
                        }}
                    >
                        <Typography
                            component="h2"
                            variant="h4"
                            sx={{
                                flexGrow: 1,
                                fontFamily: "Comic Sans MS",
                            }}
                            gutterBottom
                        >
                            Requirement
                        </Typography>
                        <Box sx={{
                            fontFamily: "Comic Sans MS"
                        }}>
                            ระบบโรงแรมจะทำให้การใช้งาน หรือทำงานของพนักงาน และ ลูกค้า ให้ง่ายและสะดวกสบายมากขึ้น
                            ระบบบริการเสริม เป็นระบบที่ให้ลูกค้าสะดวกสบายและให้ห้องนั้นมีประสิทธิภาพ เช่น อาหาร เครื่องดื่ม หรือขาดอุปกรณ์เสริมต่างๆ ภายในห้อง
                            เช่น ปลั๊กสามตา กาน้ำร้อน ไมโครเวฟ หรือเตียง และสามารถใส่จำนวนที่ต้องการ เมื่อทำการบันทึก จะบันทึกเวลา ณ ขณะนั้น
                            หากต้องการสินค้าเพิ่ม หรือไม่ต้องการสินค้านั้นแล้วสามารถแก้ไขรายการสินค้าได้ และจะทำการอัพเดทเวลา และ ราคาที่ทำการแก้ไขนั้นๆ
                            หากรายการที่ทำการสั่งไปไม่ต้องการแล้ว สามารถยกเลิกสินค้าได้ นอกจากนี้ยังแสดงรายการที่ลูกค้าทำการใช้บริการเสริม และรายการจากการสั่งสินค้าจะนำไปชำระเงินในภายหลัง
                        </Box>
                    </Paper>

                    <Paper
                        sx={{
                            bgcolor: "#white",
                            padding: 2,
                            marginBottom: 1,
                            boxShadow: 1,
                            marginTop: 0.5,
                        }}
                    >
                        <Container maxWidth="xl">
                            <Box
                                sx={{
                                    padding: 3,
                                }}
                            >
                                <Typography
                                    variant="h5"
                                    sx={{
                                        float: "right",
                                        flexGrow: 1,
                                        fontFamily: "Comic Sans MS",
                                    }}
                                    gutterBottom
                                >
                                    {name}'s Bill Service
                                </Typography>
                            </Box>
                            <TableContainer component={Paper}>
                                <Table sx={{ minWidth: 500 }}>
                                    <TableHead>
                                        <TableRow>
                                            <TableCellHead align="center" width="20%" sx={{ border: 1 }}>Bill</TableCellHead>
                                            <TableCellHead align="center" width="20%" sx={{ border: 1 }}>Bill Time</TableCellHead>
                                            <TableCellHead align="center" width="20%" sx={{ border: 1 }}>Food</TableCellHead>
                                            <TableCellHead align="center" width="20%" sx={{ border: 1 }}>Item</TableCellHead>
                                            <TableCellHead align="center" width="20%" sx={{ border: 1 }}>Drink</TableCellHead>
                                            <TableCellHead align="center" width="20%" sx={{ border: 1 }}>Item</TableCellHead>
                                            <TableCellHead align="center" width="20%" sx={{ border: 1 }}>Accessorie</TableCellHead>
                                            <TableCellHead align="center" width="20%" sx={{ border: 1 }}>Item</TableCellHead>
                                            <TableCellHead align="center" width="20%" sx={{ border: 1 }}>Total</TableCellHead>
                                            <TableCellHeadY align="center" width="20%" sx={{ border: 1 }}>Edit</TableCellHeadY>
                                        </TableRow>
                                    </TableHead>

                                    <TableBody>
                                        {service.map((item: ServicesInterface) => (
                                            <TableRow key={item.ID}>
                                                <TableCellBody align="center">{item.ID}</TableCellBody>
                                                <TableCellBody align="center">{moment(item.Time).format("DD/MM/YYYY HH:mm:ss")}</TableCellBody>
                                                <TableCellBody align="center">{item.Food?.Name}</TableCellBody>
                                                <TableCellBody align="center">{item.FoodItem}</TableCellBody>
                                                <TableCellBody align="center">{item.Drink?.Name}</TableCellBody>
                                                <TableCellBody align="center">{item.DrinkItem}</TableCellBody>
                                                <TableCellBody align="center">{item.Storage?.Product?.Name}</TableCellBody>
                                                <TableCellBody align="center">{item.StorageItem}</TableCellBody>
                                                <TableCellBody align="center">{item.Total}</TableCellBody>
                                                <TableCellBody align="center">
                                                    <ButtonGroup color="primary" aria-label="outlined primary button group">
                                                        <Button
                                                            color="warning"
                                                            component={RouterLink}
                                                            to={`/su/${item.ID}`}
                                                        >
                                                            <EditIcon /></Button>
                                                    </ButtonGroup>
                                                </TableCellBody>
                                            </TableRow>
                                        ))}
                                    </TableBody>
                                </Table>
                            </TableContainer>

                            <Typography
                                gutterBottom
                            >
                            </Typography>

                            <Grid container spacing={1} sx={{ padding: 3 }}>
                                <Grid
                                    container
                                    justifyContent="center"
                                    gap={40}
                                >
                                    <Button
                                        sx={{
                                            fontFamily: "Comic Sans MS",
                                        }}
                                        component={RouterLink}
                                        to="/sa"
                                        variant="contained"
                                        color="success"
                                    >
                                        ORDER
                                    </Button>
                                    <Button
                                        sx={{
                                            fontFamily: "Comic Sans MS",
                                        }}
                                        component={RouterLink}
                                        to="/pao"
                                        variant="contained"
                                        color="primary"
                                    >
                                        PAYMENT
                                    </Button>
                                    <Button
                                        sx={{
                                            float: "right",
                                            fontFamily: "Comic Sans MS",
                                        }}
                                        component={RouterLink}
                                        to="/sd"
                                        variant="contained"
                                        color="error"
                                    >
                                        CANCLE
                                    </Button>
                                </Grid>
                            </Grid>
                        </Container>
                    </Paper>
                </Paper>
            </Container>
        </ThemeProvider>
    );
}

export default ServiceShow