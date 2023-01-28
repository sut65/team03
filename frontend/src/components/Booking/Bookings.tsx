import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { BookingsInterface } from "../../models/modelBooking/IBooking";
import { GetBookings } from "./services/BookingHttpClientService";

function Bookings() {
    const [bookings, setBookings] = useState<BookingsInterface[]>([]);

    useEffect(() => {
        getBookings();
    }, []);

    const getBookings = async () => {
        let res = await GetBookings();
        if (res) {
            setBookings(res);
        }
    };

    const columns: GridColDef[] = [
        { field: "ID", headerName: "ลำดับ", width: 50 },
        { field: "Branch", headerName: "สาขา", width: 250, valueFormatter: (params) => params.value.Name, },
        { field: "Room", headerName: "ห้องพักหมายเลข", width: 250, valueFormatter: (params) => params.value.ID, },
        { field: "Start", headerName: "วันที่เริ่มเข้าพัก", width: 150 },
        { field: "Stop", headerName: "วันที่สิ้นสุดการเข้าพัก", width: 150 },
        { field: "Customer", headerName: "จองโดย", width: 250, valueFormatter: (params) => params.value.FirstName, },
    ]

    return (
            <Container maxWidth="md">
                <Box display="flex" sx={{ marginTop: 2, }}>
                    <Box flexGrow={1}>
                        <Typography component="h2" variant="h6" color="primary" gutterBottom>
                            ข้อมูลการจองห้องพัก
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