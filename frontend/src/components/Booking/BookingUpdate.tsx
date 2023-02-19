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
import { GetBooking, GetBookings, GetBranchs, GetCustomerByUID, UppdateBooking } from "./services/BookingHttpClientService";
import { GetRooms } from "../Room/service/RoomHttpClientService";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function BookingUpdate() {
    const [u_bookings, setU_Bookings] = useState<BookingsInterface[]>([]);
    const [branchs, setBranchs] = useState<BranchsInterface[]>([]);
    const [rooms, setRooms] = useState<RoomInterface[]>([]);
    const [booking, setBooking] = useState<BookingsInterface>({});
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

    const handleSuccess = () => {
        const shouldConfirm = window.confirm('คุณแน่ใจแล้วหรือไม่ว่าจะแก้ไขการจอง');
        if (shouldConfirm) {
            submit();
        }
    };

    const handleChange = (event: SelectChangeEvent) => {
        const name = event.target.name as keyof typeof booking;
        setBooking({
            ...booking,
            [name]: event.target.value,
        });
    };

    const onChangeU_Booking = async (event: SelectChangeEvent) => {
        let res = await GetBooking(event.target.value);
        if (res) {
            setBooking(res);
            console.log(res);
        }
    };

    const getBookings = async () => {
        let res = await GetBookings();
        if (res) {
            setU_Bookings(res);
        }
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

    const getCustomer = async () => {
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
        getBookings();
        getBranchs();
        getRooms();
        getCustomer();
    }, []);

    //for convert data type string to int
    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    const convertType_C = (data: string | number | null) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function submit() {
        let data = {
            ID: convertType(booking.ID),
            Booking_Number: booking.Booking_Number,
            BranchID: convertType(booking.BranchID),
            RoomID: convertType(booking.RoomID),
            Start: booking.Start,
            Stop: booking.Stop,
            CustomerID: convertType(booking.CustomerID),
        };

        console.log(data);

        let res = await UppdateBooking(data);
        if (res.status) {
            setAlertMessage("แก้ไขการจองห้องพักสำเร็จ");
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
                            แก้ไขการจองห้องพัก
                        </Typography>
                    </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} sx={{ padding: 2 }}>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                        <p>เลือกรายการที่จะแก้ไข</p>
                        <Select
                                native
                                value={booking.ID + ""}
                                onChange={onChangeU_Booking}
                                inputProps={{
                                    name: "ID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    กรุณาเลือกรายการที่จะแก้ไข
                                </option>
                                {u_bookings.map((item: BookingsInterface) => item.CustomerID === convertType_C(localStorage.getItem('id')) && ( 
                                    <option value={item.ID} key={item.ID}>
                                        {item.Booking_Number}
                                    </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
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
                                        ห้องพักหมายเลข: {item.Room_No} ราคา: {item.Amount} บาท
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
                                    // minDate={booking.Start}
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
                                value={booking.CustomerID + ""}
                                onChange={handleChange}
                                disabled
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
                            onClick={handleSuccess}
                            variant="contained"
                            color="warning"
                        >
                            แก้ไขการจอง
                        </Button>
                    </Grid>
                </Grid>
            </Paper>
        </Container>
    );
}
export default BookingUpdate;