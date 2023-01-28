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
import { DeleteBooking, GetBooking, GetBookings, GetBranchs, GetCustomerByUID } from "./services/BookingHttpClientService";
import { BookingsInterface } from "../../models/modelBooking/IBooking";
import { BranchsInterface } from "../../models/modelBooking/IBranch";
import MenuItem from "@mui/material/MenuItem";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function BookingDelete() {
    const [u_bookings, setU_Bookings] = useState<BookingsInterface[]>([]); // for select
    const [s_booking, setS_Booking] = useState<BookingsInterface>(); // for Show 
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

    const onChangeU_Booking = async (event: SelectChangeEvent) => {
        let res = await GetBooking(event.target.value);
        console.log(res);
        if (res) {
            setS_Booking(res);
        }
    };

    const getBookings = async () => {
        let res = await GetBookings();
        if (res) {
            setU_Bookings(res);;
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

    const convertType_C = (data: string | number | null) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function submit() {
        let data = {
            ID: s_booking?.ID,
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
                                onChange={onChangeU_Booking}
                                inputProps={{
                                    name: "ID",
                                }}
                            >
                                {u_bookings.map((item: BookingsInterface) => item.CustomerID === convertType_C(localStorage.getItem('id')) && ( 
                                    <MenuItem key={item.ID} value={item.ID}>
                                        {item.ID}
                                    </MenuItem>
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
                            >
                                {branchs.map((item: BranchsInterface) => item.ID === s_booking?.BranchID && (
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
                            >
                                {rooms.map((item: RoomInterface) => item.ID === s_booking?.RoomID && (
                                    <option aria-label="None" value={item.ID} key={item.ID} selected>
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
                                    value={s_booking?.Start}
                                    disabled
                                    onChange={(newValue) => {
                                        setS_Booking({
                                            ...s_booking
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
                                    value={s_booking?.Stop}
                                    disabled
                                    onChange={(newValue) => {
                                        setS_Booking({
                                            ...s_booking,
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