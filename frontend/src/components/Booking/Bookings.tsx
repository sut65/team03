import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { BookingsInterface } from "../../models/modelBooking/IBooking";
import { GetBookingsBYUID, GetCustomerByUID } from "./services/BookingHttpClientService";
import { CustomerInterface } from "../../models/modelCustomer/ICustomer";

function Bookings() {
    const [bookings, setBookings] = useState<BookingsInterface[]>([]);
    const [customers, setCustomers] = useState<CustomerInterface>();

    useEffect(() => {
        getBookings();
        getCustomer();
    }, []);

    const getBookings = async () => {
        let res = await GetBookingsBYUID();
        if (res) {
            setBookings(res);
            console.log(res)
        }
    };
    
    const getCustomer = async() => {
        let res = await GetCustomerByUID();
        if (res) {
            setCustomers(res);
        }
    }

    const columns: GridColDef[] = [
        { field: "Booking_Number", headerName: "เลขที่การจอง", width: 150 },
        { field: "Branch", headerName: "สาขา", width: 150, valueFormatter: (params) => params.value.B_name, },
        { field: "Room", headerName: "ห้องพักหมายเลข", width: 150, valueFormatter: (params) => params.value.Room_No, },
        { field: "Start", headerName: "วันที่เริ่มเข้าพัก", width: 150 },
        { field: "Stop", headerName: "วันที่สิ้นสุดการเข้าพัก", width: 150 },
        { field: "Customer", headerName: "จองโดย", width: 150, valueFormatter: (params) => params.value.FirstName, },
        { field: "TotalAmount", headerName: "คิดเป็นเงิน/รายการ", width: 150 },
    ]

    return (
            <Container maxWidth="md">
                <Box display="flex" sx={{ marginTop: 2, }}>
                    <Box flexGrow={1}>
                        <Typography component="h2" variant="h6" color="primary" gutterBottom>
                            ข้อมูลการจองห้องพักของคุณ {customers?.FirstName}
                        </Typography>
                    </Box>
                    <Box>
                        <Button component={RouterLink} to="/Book/Create" variant="contained" color="primary">
                            จองห้องพัก
                        </Button>
                    </Box>
                    <Box>
                        <Button component={RouterLink} to="/Book/Edit" variant="contained" color="warning">
                            แก้ไขการจองห้องพัก
                        </Button>
                    </Box>
                    <Box>
                        <Button component={RouterLink} to="/Book/Delete" variant="contained" color="error">
                            ยกเลิกการจองห้องพัก
                        </Button>
                    </Box>
                </Box>
                <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
                    <DataGrid rows={bookings} getRowId={(row) => row.ID} columns={columns} pageSize={5} rowsPerPageOptions={[5]}/>
                </div>
            </Container>
    );
}

export default Bookings;