import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { CHK_PaymentsInterface } from "../../models/ICHK_Payment";
import { GetCHK_Payments } from "./service/CHK_PaymentHttpClientService";

function CHK_Payments() {
    const [chk_payments, setCHK_Payments] = useState<CHK_PaymentsInterface[]>([]);

    useEffect(() => {
        getCHK_Payments();
    }, []);

    const getCHK_Payments = async () => {
        let res = await GetCHK_Payments();
        if (res) {
            setCHK_Payments(res);
        }
    };

    const columns: GridColDef[] = [
        { field: "ID", headerName: "ลำดับ", width: 50 },
        { field: "Payment", headerName: "รายการชำระเงิน", width: 250, valueFormatter: (params) => params.value.Name },
        { field: "Status", headerName: "สถานะการชำระเงิน", width: 250, valueFormatter: (params) => params.value.Type },
        { field: "Date_time", headerName: "วัน-เวลาที่ชำระเงิน", width: 150 },
        { field: "Amount", headerName: "จำนวนเงินที่ชำระ", width: 150 },
        { field: "Description", headerName: "คำอธิบายเพิ่มเติม", width: 150 },
        { field: "User", headerName: "ตรวจสอบโดย", width: 250,  valueFormatter: (params) => params.value.Name},
    ]

    return (
        <div>
            <Container maxWidth="md">
                <Box display="flex" sx={{ marginTop: 2, }}>
                    <Box flexGrow={1}>
                        <Typography component="h2" variant="h6" color="primary" gutterBottom>
                            ข้อมูลการตรวจสอบการชำระเงิน
                        </Typography>
                    </Box>
                    <Box>
                        <Button component={RouterLink} to="/chk_payment/create" variant="contained" color="primary">
                            ตรวจสอบการชำระเงิน
                        </Button>
                    </Box>
                    <Box>
                        <Button component={RouterLink} to="/chk_payment/create" variant="contained" color="warning">
                            แก้ไขการตรวจสอบการชำระเงิน
                        </Button>
                    </Box>
                </Box>
                <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
                    <DataGrid rows={chk_payments} getRowId={(row) => row.ID} columns={columns} pageSize={5} rowsPerPageOptions={[5]}/>
                </div>
            </Container>
        </div>
    );
}
export default CHK_Payments;