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
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";

import { GetCHK_Payments, GetEmployeeByUID, GetStatuses, GetPayments, GetCHK_Payment, UppdateCHK_Payment } from "./services/CHK_PaymentHttpClientService";
import { EmployeeInterface } from "../../models/IEmployee";
import { CHK_PaymentStatusesInterface } from "../../models/modelCHK_Payment/IStatus";
import { CHK_PaymentsInterface } from "../../models/modelCHK_Payment/ICHK_Payment";
import { PaymentsInterface } from "../../models/modelPayment/IPayment";
import { DateTimePicker } from "@mui/x-date-pickers/DateTimePicker";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function CHK_PaymentUpdate() {
    const [chk_payment, setCHK_Payment] = useState<CHK_PaymentsInterface>({});
    const [u_chkpayments, setU_CHKPayments] = useState<CHK_PaymentsInterface[]>([]);
    const [payments, setPayments] = useState<PaymentsInterface[]>([]);
    const [statuses, setStatuses] = useState<CHK_PaymentStatusesInterface[]>([]);
    const [employees, setEmployees] = useState<EmployeeInterface>();
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
        const name = event.target.name as keyof typeof chk_payment;
        setCHK_Payment({
            ...chk_payment,
            [name]: event.target.value,
        });
    };

    const handleSuccess = () => {
        const shouldConfirm = window.confirm('คุณแน่ใจแล้วหรือไม่ว่าจะแก้ไขการตรวจสอบการชำระเงิน');
        if (shouldConfirm) {
            submit();
        }
    };

    const onChangeU_CHKPayment = async (event: SelectChangeEvent) => {
        let res = await GetCHK_Payment(event.target.value);
        console.log(res);
        if (res) {
            setCHK_Payment(res);
        }
    };

    const handleInputChange_Text = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
        const id = event.target.id as keyof typeof chk_payment;
        const { value } = event.target;
        setCHK_Payment({ ...chk_payment, [id]: value, });
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

    const getCHK_Payments = async () => {
        let res = await GetCHK_Payments();
        if (res) {
            setU_CHKPayments(res);
        }
    }

    const getEmployeeByUID = async () => {
        let res = await GetEmployeeByUID();
        if (res) {
            setEmployees(res);
        }
    }

    useEffect(() => {
        getStatuses();
        getPayments();
        getCHK_Payments();
        getEmployeeByUID();
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
            ID: convertType(chk_payment.ID),
            PaymentID: convertType(chk_payment.PaymentID),
            CHK_PaymentStatusID: convertType(chk_payment.CHK_PaymentStatusID),
            Date_time: chk_payment.Date_time,
            Amount: convertType(chk_payment.Amount),
            Description: chk_payment.Description,
            EmployeeID: convertType_C(localStorage.getItem('id')),
        };

        console.log(data)

        let res = await UppdateCHK_Payment(data);
        if (res.status) {
            setAlertMessage("แก้ไขรายการตรวจสอบการชำระเงินสำเร็จ");
            setSuccess(true);
            setInterval(() => {
                window.location.assign("/CPM");
            }, 2000);
        } else {
            setAlertMessage(res.message);
            setError(true);
        }
    }

    return (
        <Container maxWidth="md">
            <Snackbar open={success} autoHideDuration={3000} onClose={handleClose} anchorOrigin={{ vertical: "top", horizontal: "center" }} >
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
                            แก้ไขการตรวจสอบการชำระเงิน
                        </Typography>
                    </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} sx={{ padding: 2 }}>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>เลือกรายการการชำระเงิน</p>
                            <Select
                                native
                                value={chk_payment.PaymentID + ""}
                                onChange={onChangeU_CHKPayment}
                                inputProps={{
                                    name: "PaymentID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    กรุณาเลือกรายการชำระเงิน
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
                                value={chk_payment.CHK_PaymentStatusID + ""}
                                onChange={handleChange}
                                inputProps={{
                                    name: "CHK_PaymentStatusID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    กรุณาเลือกสถานะการชำระเงิน
                                </option>
                                {statuses.map((item: CHK_PaymentStatusesInterface) => (
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
                            <LocalizationProvider dateAdapter={AdapterDateFns}>
                                <DateTimePicker
                                    // disableFuture
                                    inputFormat="dd/MM/yyyy HH:mm"
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
                            <TextField
                                type="number"
                                value={chk_payment.Amount || " "}
                                id="Amount"
                                label="กรุณาใส่จำนวนเงิน"
                                variant="outlined"
                                inputProps={{
                                    name: "Amount",
                                    // min: 0
                                }}
                                onChange={handleInputChange_Text}
                            />
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>หมายเหตุ</p>
                            <TextField
                                type="text"
                                value={chk_payment.Description || " "}
                                id="Description"
                                label="ป้อนหมายเหตุที่นี่"
                                variant="outlined"
                                inputProps = {{
                                    name: "Description"
                                }}
                                onChange = {handleInputChange_Text}
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
                            onClick={handleSuccess}
                            variant="contained"
                            color="warning"
                        >
                            บันทึกการแก้ไข
                        </Button>
                    </Grid>
                </Grid>
            </Paper>
        </Container>
    );
}
export default CHK_PaymentUpdate;