import {
    Box,
    Button,
    Container,
    Divider,
    FormControl,
    FormControlLabel,
    FormLabel,
    Grid,
    MenuItem,
    Paper,
    Radio,
    RadioGroup,
    Rating,
    Select,
    SelectChangeEvent,
    Snackbar,
    TextField,
    ThemeProvider,
    Typography,
  } from "@mui/material";
  import { grey } from "@mui/material/colors";
  import { createTheme } from "@mui/material/styles";
  import * as React from "react";
  import { DepartmentInterface } from "../../models/IEmployee";
  import {
    ReviewInterface,
    SystemworkInterface,
  } from "../../models/IReview";
  import { Link as RouterLink } from "react-router-dom";
  import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { CustomerInterface } from "../../models/modelCustomer/ICustomer";
  
  const bgbutton = createTheme({
    palette: {
      primary: {
        // Purple and grey play nicely together.
        main: grey[800],
      },
      secondary: {
        // Purple and grey play nicely together.
        main: grey[50],
      },
    },
  });
  
  const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
  
    ref
  ) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });
  
  function Review_Edit() {
    const [start, setStart] = React.useState<number | null>();
    const [review, setReview] = React.useState<Partial<ReviewInterface>>({});
    const [user, setUser] = React.useState<CustomerInterface>();
    const [department, setDepartment] = React.useState<DepartmentInterface[]>([]);
    const [systemwork, setSystemwork] = React.useState<SystemworkInterface[]>([]);
    const [success, setSuccess] = React.useState(false);
    const [error, setError] = React.useState(false);
    const [imageString, setImageString] = React.useState<string | ArrayBuffer | null>(null);
    const [reviewdate, setReviewdate] = React.useState<Date | null>(new Date());
  
    const getDepartment = async () => {
      const apiUrl = `http://localhost:8080/Departments`;
  
      const requestOptions = {
        method: "GET",
  
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
          "Content-Type": "application/json",
        },
      };
      //การกระทำ //json
      fetch(apiUrl, requestOptions)
        .then((response) => response.json()) //เรียกได้จะให้แสดงเป็น json ซึ่ง json คือ API
  
        .then((res) => {
          console.log(res.data); //show ข้อมูล
  
          if (res.data) {
            setDepartment(res.data);
          } else {
            console.log("else");
          }
        });
    };
    //---------------------Systemwork-------------------------------------
    const getSystemwork = async () => {
      const apiUrl = `http://localhost:8080/Systemworks`;
  
      const requestOptions = {
        method: "GET",
  
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
          "Content-Type": "application/json",
        },
      };
      //การกระทำ //json
      fetch(apiUrl, requestOptions)
        .then((response) => response.json()) //เรียกได้จะให้แสดงเป็น json ซึ่ง json คือ API
  
        .then((res) => {
          console.log(res.data); //show ข้อมูล
  
          if (res.data) {
            setSystemwork(res.data);
          } else {
            console.log("else");
          }
        });
    };
  
    const handleImageChange = (event: any) => {
      const image = event.target.files[0];
      const reader = new FileReader();
      reader.readAsDataURL(image);
      reader.onload = () => {
          const base64Data = reader.result;
          setImageString(base64Data)
      }
  }
  
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
  
    const handleInputChange = (
      event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
      const id = event.target.id as keyof typeof Review_Edit;
      const { value } = event.target;
      setReview({ ...review, [id]: value });
    };
  
    const handleChange = (event: SelectChangeEvent<number>) => {
      const name = event.target.name as keyof typeof review;
      setReview({
        ...review,
        [name]: event.target.value,
      });
    };
  
    function editReview() {
        let dataemployee = {
          Comment: review.Comment ?? "",
          Star: start,
          Reviewdate: reviewdate,
          Reviewimage: imageString,
          CustomerID: user?.ID ?? "",
          DepartmentID: review.DepartmentID,
          SystemworkID: review.SystemworkID,
        }
    
        console.log(dataemployee)
        const apiUrlUpdate = "http://localhost:8080/Reviews";
    
        const requestUpdateOptions = {
            method: "PATCH",
            headers: {
              Authorization: `Bearer ${localStorage.getItem("token")}`,
              "Content-Type": "application/json",
            },
            body: JSON.stringify(dataemployee),
        }
    
        fetch(apiUrlUpdate, requestUpdateOptions)
          .then((response) => response.json())
          .then((res) => {
            if (res.data) {
              setSuccess(true);
            //   setErrorMassage("");
            //   setBtnDisabled(!btnDisabled)
            } else {
              
              setError(true);
            //   setErrorMassage(res.error);
            }
          });
      }
  
    React.useEffect(() => {
      const getToken = localStorage.getItem("token");
      if (getToken) {
        setUser(JSON.parse(localStorage.getItem("user") || ""));
      }
      getDepartment();
      getSystemwork();
    }, []);
    return (
      <ThemeProvider theme={bgbutton}>
        <Container maxWidth="md">
          <Snackbar
            open={success}
            autoHideDuration={6000}
            onClose={handleClose}
            anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
          >
            <Alert onClose={handleClose} severity="success">
              บันทึกข้อมูลสำเร็จ
            </Alert>
          </Snackbar>
  
          <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="error">
              บันทึกข้อมูลไม่สำเร็จ
            </Alert>
          </Snackbar>
  
          <Paper>
            <Box
              display="flex"
              sx={{
                marginTop: 2,
              }}
            >
              <Box sx={{ paddingX: 2, paddingY: 1 }}>
                <Typography
                  component="h2"
                  variant="h6"
                  color="primary"
                  gutterBottom
                >
                  Review
                </Typography>
              </Box>
            </Box>
  
            <Divider />
  
            <Grid
              container
              spacing={3}
              sx={{ padding: 2 }}
              style={{ marginLeft: "10.5%" }}
            >
              {/* ComboboxDepartment */}
              <Grid item xs={3}>
                <FormLabel>Department</FormLabel>
                <FormControl fullWidth variant="outlined">
                  <Select
                    value={review.DepartmentID}
                    onChange={handleChange}
                    inputProps={{
                      name: "DepartmentID",
                    }}
                  >
                    <MenuItem value={0} key={0}>
                      เลือกแผนก
                    </MenuItem>
                    {department.map((item: DepartmentInterface) => (
                      <MenuItem value={item.ID}>{item.Name}</MenuItem>
                    ))}
                  </Select>
                </FormControl>
              </Grid>
  
              {/* ComboboxPosition */}
              <Grid item xs={3}>
                <FormLabel>System Work</FormLabel>
                <FormControl fullWidth variant="outlined">
                  <Select
                    value={review.SystemworkID}
                    onChange={handleChange}
                    inputProps={{
                      name: "SystemworkID",
                    }}
                  >
                    <MenuItem value={0} key={0}>
                      เลือกตำแหน่ง
                    </MenuItem>
                    {systemwork.map((item: SystemworkInterface) => (
                      <MenuItem value={item.ID}>{item.Name}</MenuItem>
                    ))}
                  </Select>
                </FormControl>
              </Grid>
            </Grid>
  
            <Grid
              container
              spacing={3}
              sx={{ padding: 2 }}
              style={{ marginLeft: "14.5%" }}
            >
              <Grid item xs={6}>
                <FormControl>
                  <FormLabel>Start</FormLabel>
                  <Rating
                    name="simple-controlled"
                  //   value={5}
                  onChange={(event, newValue) => {
                      setStart(newValue);
                    }}
                  />
                </FormControl>
              </Grid>
            </Grid>
  
            <Grid container spacing={3} sx={{ padding: 2 }}>
              <Grid item xs={10}>
                <FormControl fullWidth variant="outlined">
                  <FormLabel>Comment</FormLabel>
  
                  <TextField
                    id="Comment"
                    variant="outlined"
                    type="string"
                    size="medium"
                    value={review.Comment || ""}
                    onChange={handleInputChange}
                  />
                </FormControl>
              </Grid>
            </Grid>
  
              <Grid container spacing={3} sx={{ padding: 2 }}>
              <Grid item xs={10}>
                <FormControl fullWidth variant="outlined">
                  <FormLabel>Image</FormLabel>
                  <img src={`${imageString}`} width="500" height="500"/>
                  <input type="file" onChange={handleImageChange} />
                </FormControl>
              </Grid>
  
              {/* วันที่ทำงาน */}
              <Grid item xs={12}>
                <Button component={RouterLink} to="/RW" variant="contained">
                  Back
                </Button>
  
                <Button
                  style={{ float: "right" }}
                  onClick={editReview}
                  variant="contained"
                  color="primary"
                >
                  Submit
                </Button>
              </Grid>
  
            </Grid>
          </Paper>
        </Container>
      </ThemeProvider>
    );
  }
  
  export default Review_Edit;
  