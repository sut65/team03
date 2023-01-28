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

import { RoomInterface } from "../../models/IRoom";
import { CustomerInterface } from "../../models/modelCustomer/ICustomer";
import { GetRooms } from "../Room/service/RoomHttpClientService";
import { DeleteBooking, GetBookings, GetBranchs, GetCustomerByUID } from "./services/BookingHttpClientService";
import { BookingsInterface } from "../../models/modelBooking/IBooking";
import { BranchsInterface } from "../../models/modelBooking/IBranch";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function BookingDelete() {
    const [booking, setBooking] = useState<BookingsInterface>({});
    const [u_bookings, setU_Bookings] = useState<BookingsInterface[]>([]);
    const [branchs, setBranchs] = useState<BranchsInterface[]>([]);
    const [rooms, setRooms] = useState<RoomInterface[]>([]);
    const [customers, setCustomers] = useState<CustomerInterface>();
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);

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

    const getCustomerByUID = async () => {
        let res = await GetCustomerByUID();
        if (res) {
            setCustomers(res);
        }
    }

    useEffect(() => {
        getBookings();
        getBranchs();
        getRooms();
        getCustomerByUID();
    }, []);

    async function submit() {
        let data = {
            ID: booking.ID,
        };

        let res = await DeleteBooking(data);
        if (res) {
            setSuccess(true);
        } else {
            setError(true);
        }
    }

    return (
        <Container maxWidth="md">
            <Snackbar open={success} autoHideDuration={3000} onClose={handleClose} anchorOrigin={{ vertical: "top", horizontal: "center" }} >
                <Alert onClose={handleClose} severity="success">
                    จองห้องพักสำเร็จ
                </Alert>
            </Snackbar>
            <Snackbar open={error} autoHideDuration={6000} onClose={handleClose} anchorOrigin={{ vertical: "top", horizontal: "center" }} >
                <Alert onClose={handleClose} severity="error">
                    ไม่ไม่สามารถจองห้องพักได้
                </Alert>
            </Snackbar>
            <Paper>
                <Box display="flex" sx={{ marginTop: 2, }} >
                    <Box sx={{ paddingX: 2, paddingY: 1 }}>
                        <Typography component="h2" variant="h6" color="primary" gutterBottom >
                            ยกเลิกการจองห้องพัก
                        </Typography>
                    </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} sx={{ padding: 2 }}>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>เลือกรายการจองที่จะยกเลิก</p>
                            <Select
                                native
                                value={booking.ID + ""}
                                onChange={handleChange}
                                inputProps={{
                                    name: "ID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    กรุณาเลือกรายการจองที่จะแก้ไข
                                </option>
                                {u_bookings.map((item: BookingsInterface) => (
                                    <option value={item.ID} key={item.ID}>
                                        {item.ID}
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
                                disabled
                                value={booking.BranchID + ""}
                                onChange={handleChange}
                                inputProps={{
                                    name: "BranchID",
                                }}
                            >
                                {branchs.map((item: BranchsInterface) => item.ID === booking.BranchID && (
                                    <option aria-label="None" value={item.ID} key={item.ID} selected>
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
                                disabled
                                value={booking.RoomID + ""}
                                onChange={handleChange}
                                inputProps={{
                                    name: "RoomID",
                                }}
                            >
                                {rooms.map((item: RoomInterface) => item.ID === booking.RoomID && (
                                    <option value={item.ID} key={item.ID} selected>
                                        {item.ID}
                                    </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>วันที่เข้าพัก</p>
                            {/* input from roomid andthen search booking where roomid and get start\stop day in recorded   */}
                            <LocalizationProvider dateAdapter={AdapterDateFns}>
                                <DatePicker
                                    value={booking.Start}
                                    disabled
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
                                    value={booking.Stop}
                                    disabled
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
                    <Grid item xs={12}>
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
                                    {customers?.FirstName}," ", {customers?.LastName}
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
                            color="error"
                        >
                            ยกเลิกการจอง
                        </Button>
                    </Grid>
                </Grid>
            </Paper>
        </Container>
    );
}
export default BookingDelete;