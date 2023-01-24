import React, { useEffect } from "react";

import { Link as RouterLink } from "react-router-dom";

import Typography from "@mui/material/Typography";

import Button from "@mui/material/Button";

import Container from "@mui/material/Container";

import Box from "@mui/material/Box";


import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { UsersInterface } from "../../models/IUser";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { grey } from "@mui/material/colors";
import { EmployeeInterface } from "../../models/IEmployee";
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";
import moment from "moment";

const themeshow = createTheme({
  palette: {
    primary: {
      // Purple and green play nicely together.
      main: grey[800],
    },
    secondary: {
      // This is green.A700 as hex.
      main: "#e8f5e9",
    },
  }
});


function Manage_Show() {

  const [employee, setEmployee] = React.useState<EmployeeInterface[]>([]);
  
 const getEmployee = async () => {
   const apiUrl = "http://localhost:8080/Employees";
   const requestOptions = {
     method: "GET",
     headers: { Authorization: `Bearer ${localStorage.getItem("token")}`, "Content-Type": "application/json" },
   };


   fetch(apiUrl, requestOptions)
     .then((response) => response.json())
     .then((res) => {
       console.log(res.data);
       if (res.data) {
        setEmployee(res.data);
       }
     });
 };


 useEffect(() => {
   getEmployee();
 }, []);

 return (
<ThemeProvider theme={themeshow}>
  <div>
  <Container maxWidth="md">
  <Box
        display="flex"
        sx={{
          marginTop: 2,
        }}
      >
        <Box flexGrow={1}>
          <Typography
            component="h2"
            variant="h6"
            color="primary"
            gutterBottom
          >
            Employee data table
          </Typography>
        </Box>

        <Box>
          <Button 
            component={RouterLink}
            to="/Man"
            variant="contained"
            color="primary"
          >
            <Typography
              color="secondary"
              component="div"
              sx={{ flexGrow: 1 }}
            >
              Record Employee Information

            </Typography>
          </Button>
        </Box>
      </Box>
        <div>
          <Container maxWidth="md">
            <div style={{ height: 500, width: "100%", marginTop: "50px" }}>
              <TableContainer >
                <Table aria-label="simple table">
                  <TableHead>
                    {/* หัวข้อตาราง */}
                    <TableRow>
                      <TableCell align="center" width="20%"> Personal ID </TableCell>
                      <TableCell align="center" width="20%"> Name </TableCell>
                      <TableCell align="center" width="20%"> Department </TableCell>
                      <TableCell align="center" width="20%"> Position </TableCell>
                      <TableCell align="center" width="20%"> Location </TableCell>
                      <TableCell align="center" width="20%"> Username </TableCell>
                      <TableCell align="center" width="20%"> Password </TableCell>
                      <TableCell align="center" width="20%"> Email </TableCell>
                      <TableCell align="center" width="20%"> Tel </TableCell>
                      <TableCell align="center" width="20%"> Salary </TableCell>
                      <TableCell align="center" width="20%"> Gender </TableCell>
                      <TableCell align="center" width="20%"> Address </TableCell>
                      <TableCell align="center" width="20%"> Date of Birth </TableCell>
                      <TableCell align="center" width="20%"> Year of Start </TableCell>
                      <TableCell align="center" width="20%"> Officer </TableCell>

                    </TableRow>
                  </TableHead>
                  
                  <TableBody>
                    {employee.map((item: EmployeeInterface) => (
                      <TableRow key={item.ID}>
                        <TableCell align="center">{item.PersonalID}</TableCell>
                        <TableCell align="center">{item.Employeename}</TableCell>
                        <TableCell align="center">{item.Department?.Name}</TableCell>
                        <TableCell align="center">{item.Position?.Name}</TableCell>
                        <TableCell align="center">{item.Location?.Name}</TableCell>
                        <TableCell align="center">{item.Eusername}</TableCell>
                        <TableCell align="center">{item.Password}</TableCell>
                        <TableCell align="center">{item.Email}</TableCell>
                        <TableCell align="center">{item.Phonenumber}</TableCell>
                        <TableCell align="center">{item.Salary}</TableCell>
                        <TableCell align="center">{item.Gender}</TableCell>
                        <TableCell align="center">{item.Address}</TableCell>
                        <TableCell align="center">{moment(item.DateOfBirth).format("DD/MM/YYYY")}</TableCell>
                        <TableCell align="center">{moment(item.YearOfStart).format("DD/MM/YYYY")}</TableCell>
                        <TableCell align="center">{item.Officer?.Officername}</TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>
            </div>
          </Container>
        </div>
  </Container>
 </div>
</ThemeProvider>
 );

}


export default Manage_Show;