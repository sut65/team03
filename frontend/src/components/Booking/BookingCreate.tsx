import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import Container from "@mui/material/Container";
import Snackbar from "@mui/material/Snackbar";
import Paper from "@mui/material/Paper";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Grid from "@mui/material/Grid";
import FormControl from "@mui/material/FormControl";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";

import { BookingsInterface } from "../../models/modelBooking/IBooking";
import { BranchsInterface } from "../../models/modelBooking/IBranch";
import { RoomInterface } from "../../models/IRoom";
import { CustomerInterface } from "../../models/modelCustomer/ICustomer";
import { Bookings, GetBranchs, GetCustomerByUID } from "./services/BookingHttpClientService";
import { GetRooms } from "../Room/service/RoomHttpClientService";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function BookingCreate() {
    const [booking, setBooking] = useState<BookingsInterface>({
        Start: new Date(),
        Stop: new Date(),
    });
    const [branchs, setBranchs] = useState<BranchsInterface[]>([]);
    const [rooms, setRooms] = useState<RoomInterface[]>([]);
    const [customers, setCustomers] = useState<CustomerInterface>();
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const [message, setAlertMessage] = useState("");

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

    const handleChange = (event: SelectChangeEvent) => {
        const name = event.target.name as keyof typeof booking;
        setBooking({
            ...booking,
            [name]: event.target.value,
        });
    };

    const getBranchs = async () => {
        let res = await GetBranchs();
        if (res) {
            setBranchs(res);
        }
    };

    const getRooms = async () => {
        let res = await GetRooms();
        if (res) {
            setRooms(res);
        }
    };

    const getCustomer = async() => {
        let res = await GetCustomerByUID();
        if (res) {
            setBooking({
                ...booking,
                CustomerID: res.ID,
            });
            setCustomers(res);
        }
    }

    useEffect(() => {
        getBranchs();
        getRooms();
        getCustomer();
    }, []);

    //for convert data type string to int
    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function submit() {
        let data = {
            BranchID: convertType(booking.BranchID),
            RoomID: convertType(booking.RoomID),
            Start: booking.Start,
            Stop: booking.Stop,
            CustomerID: convertType(booking.CustomerID),
        };

        console.log(data);

        let res = await Bookings(data);
        if (res.status) {
            setAlertMessage("จองห้องพักสำเร็จ");
            setSuccess(true);
        } else {
            setAlertMessage(res.message);
            setError(true);
        }
    }

    return (
        <Container maxWidth="md">
            <Snackbar open={success} autoHideDuration={6000} onClose={handleClose} anchorOrigin={{ vertical: "top", horizontal: "center" }} >
                <Alert onClose={handleClose} severity="success">
                    {message}
                </Alert>
            </Snackbar>
            <Snackbar open={error} autoHideDuration={6000} onClose={handleClose} anchorOrigin={{ vertical: "top", horizontal: "center" }} >
                <Alert onClose={handleClose} severity="error">
                    {message}
                </Alert>
            </Snackbar>
            <Paper>
                <Box display="flex" sx={{ marginTop: 2, }} >
                    <Box sx={{ paddingX: 2, paddingY: 1 }}>
                        <Typography component="h2" variant="h6" color="primary" gutterBottom >
                            จองห้องพัก
                        </Typography>
                    </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} sx={{ padding: 2 }}>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>เลือกสาขาของโรงแรม</p>
                            <Select
                                native
                                value={booking.BranchID + ""}
                                onChange={handleChange}
                                inputProps={{
                                    name: "BranchID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    กรุณาเลือกสาขาที่จะเข้าพัก
                                </option>
                                {branchs.map((item: BranchsInterface) => (
                                    <option value={item.ID} key={item.ID}>
                                        {item.B_name}
                                    </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>ห้องพัก</p>
                            <Select
                                native
                                value={booking.RoomID + ""}
                                onChange={handleChange}
                                inputProps={{
                                    name: "RoomID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    กรุณาเลือกห้องพัก
                                </option>
                                {rooms.map((item: RoomInterface) => (
                                    <option value={item.ID} key={item.ID}>
                                        ห้องพักหมายเลข: {item.Room_No} ราคา: {item.Amount}
                                    </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>วันที่เข้าพัก</p>
                            <LocalizationProvider dateAdapter={AdapterDateFns}>
                                <DatePicker
                                    // disablePast
                                    value={booking.Start}
                                    onChange={(newValue) => {
                                        setBooking({
                                            ...booking,
                                            Start: newValue,
                                        });
                                    }}
                                    renderInput={(params) => <TextField {...params} />}
                                />
                            </LocalizationProvider>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>วันที่สิ้นสุดการพัก</p>
                            <LocalizationProvider dateAdapter={AdapterDateFns}>
                                <DatePicker
                                    // minDate={booking.Start && new Date(booking.Start.getTime() + (24 * 60 * 60 * 1000))}
                                    value={booking.Stop}
                                    onChange={(newValue) => {
                                        setBooking({
                                            ...booking,
                                            Stop: newValue,
                                        });
                                    }}
                                    renderInput={(params) => <TextField {...params} />}
                                />
                            </LocalizationProvider>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>จองโดย</p>
                            <Select
                                native
                                disabled
                                value={booking.CustomerID + ""}
                                onChange={handleChange}
                                inputProps={{
                                    name: "CustomerID",
                                }}
                            >
                                <option value={customers?.ID} key={customers?.ID}>
                                    {customers?.FirstName} {customers?.LastName}
                                </option>
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={12}>
                        <Button
                            component={RouterLink}
                            to="/Book"
                            variant="contained"
                            color="inherit"
                        >
                            กลับ
                        </Button>
                        <Button
                            style={{ float: "right" }}
                            onClick={submit}
                            variant="contained"
                            color="primary"
                        >
                            บันทึกการจอง
                        </Button>
                    </Grid>
                </Grid>
            </Paper>
        </Container>
    );
}
export default BookingCreate;