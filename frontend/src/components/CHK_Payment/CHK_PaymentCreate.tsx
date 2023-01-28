import React, { useEffect, useState } from "react";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { Link as RouterLink } from "react-router-dom";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Divider from "@mui/material/Divider";
import FormControl from "@mui/material/FormControl";
import Grid from "@mui/material/Grid";
import Paper from "@mui/material/Paper";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import Snackbar from "@mui/material/Snackbar";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";

import { CHK_Payments, GetEmployeeByUID, GetStatuses, GetPayments} from "./services/CHK_PaymentHttpClientService";
import { EmployeeInterface } from "../../models/IEmployee";
import { StatusesInterface } from "../../models/modelCHK_Payment/IStatus";
import { PaymentsInterface } from "../../models/modelPayment/IPayment";
import { CHK_PaymentsInterface } from "../../models/modelCHK_Payment/ICHK_Payment";



const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function CHK_PaymentCreate() {
    const [chk_payment, setCHK_Payment] = useState<CHK_PaymentsInterface>({});
    const [payments, setPayments] = useState<PaymentsInterface[]>([]);
    const [statuses, setStatuses] = useState<StatusesInterface[]>([]);
    const [employees, setEmployees ] = useState<EmployeeInterface>();
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
        const name = event.target.name as keyof typeof chk_payment;
        setCHK_Payment({
            ...chk_payment,
            [name]: event.target.value,
        });
    };

    const getStatuses = async () => {
        let res = await GetStatuses();
        if (res) {
            setStatuses(res);
        }
    };

    const getPayments = async () => {
        let res = await GetPayments();
        if (res) {
            setPayments(res);
        }
    };

    const getEmployeeByUID = async () => {
        let res = await GetEmployeeByUID();
        if (res) {
            setEmployees(res);
        }
    }

    useEffect(() => {
        getStatuses();
        getPayments();
        getEmployeeByUID();
    }, []);

    //for convert data type string to int
    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function submit() {
        let data = {
            PaymentID: convertType(chk_payment.PaymentID),
            StatusID: convertType(chk_payment.StatusID),
            Date_time: chk_payment.Date_time,
            Amount: convertType(chk_payment.Amount),
            Description: chk_payment.Description,
            EmployeeID: convertType(chk_payment.EmployeeID),
        };

        let res = await CHK_Payments(data);
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
                    บันทึกการตตรวจสอบเสร็จสิ้น
                </Alert>
            </Snackbar>
            <Snackbar open={error} autoHideDuration={6000} onClose={handleClose} anchorOrigin={{ vertical: "top", horizontal: "center" }} >
                <Alert onClose={handleClose} severity="error">
                    ไม่สามารถบันทึกการตตรวจสอบได้
                </Alert>
            </Snackbar>
            <Paper>
                <Box display="flex" sx={{ marginTop: 2, }} >
                    <Box sx={{ paddingX: 2, paddingY: 1 }}>
                        <Typography component="h2" variant="h6" color="primary" gutterBottom >
                            ตรวจสอบการชำระเงิน
                        </Typography>
                    </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} sx={{ padding: 2 }}>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>เลือกรายการ การชำระเงิน</p>
                            <Select
                                native
                                value={chk_payment.PaymentID + ""}
                                onChange={handleChange}
                                inputProps={{
                                    name: "PaymentID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    กรุณาเลือกราย การชำระเงิน
                                </option>
                                {payments.map((item: PaymentsInterface) => (
                                    <option value={item.ID} key={item.ID}>
                                        {item.ID}
                                    </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>สถานะการชำระเงิน</p>
                            <Select
                                native
                                value={chk_payment.StatusID + ""}
                                onChange={handleChange}
                                inputProps={{
                                    name: "StatusID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    กรุณาเลือกสถานะการชำระเงิน
                                </option>
                                {statuses.map((item: StatusesInterface) => (
                                    <option value={item.ID} key={item.ID}>
                                        {item.Type}
                                    </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>วันที่ชำระเงิน</p>
                            {/* input from roomid andthen search booking where roomid and get start\stop day in recorded   */}
                            <LocalizationProvider dateAdapter={AdapterDateFns}>
                                <DatePicker
                                    value={chk_payment.Date_time}
                                    onChange={(newValue) => {
                                        setCHK_Payment({
                                            ...chk_payment,
                                            Date_time: newValue,
                                        });
                                    }}
                                    renderInput={(params) => <TextField {...params} />}
                                />
                            </LocalizationProvider>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>จำนวนเงิน</p>
                            <TextField value={chk_payment.Amount}
                                id="outlined-basic"
                                label="กรุณาใส่จำนวนเงิน"
                                variant="outlined"
                            />
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>หมายเหตุ</p>
                            <TextField value={chk_payment.Description}
                                id="outlined-basic"
                                label="ป้อนหมายเหตุที่นี่"
                                variant="outlined"
                            />
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>ตรวจสอบโดย</p>
                            <Select
                                native
                                value={chk_payment.EmployeeID + ""}
                                onChange={handleChange}
                                disabled
                                inputProps={{
                                    name: "EmployeeID",
                                }}
                            >
                                <option value={employees?.ID} key={employees?.ID}>
                                    {employees?.Employeename}
                                </option>
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={12}>
                        <Button
                            component={RouterLink}
                            to="/CPM"
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
                            บันทึก
                        </Button>
                    </Grid>
                </Grid>
            </Paper>
        </Container>
    );
}
export default CHK_PaymentCreate;