import {
  Box,
  Button,
  Container,
  Divider,
  FormControl,
  FormLabel,
  Grid,
  MenuItem,
  Paper,
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
  CustomerInterface,
  ReviewInterface,
  SystemworkInterface,
} from "../../models/IReview";
import { Link as RouterLink } from "react-router-dom";
import MuiAlert, { AlertProps } from "@mui/material/Alert";

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

function Review_Save() {
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
    const id = event.target.id as keyof typeof Review_Save;

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

  function submit() {
    let data = {
      Comment: review.Comment ?? "",
      Star: start,
      Reviewdate: reviewdate,
      Reviewimega: imageString,
      CustomerID: user?.ID ?? "",
      DepartmentID: review.DepartmentID,
      SystemworkID: review.SystemworkID,
    };

    const apiUrl = "http://localhost:8080/Reviews";

    const requestOptions = {
      method: "POST",

      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },

      body: JSON.stringify(data),
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        if (res.data) {
          window.location.reload();
          setSuccess(true);
        } else {
          setError(true);
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
                onClick={submit}
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

export default Review_Save;
