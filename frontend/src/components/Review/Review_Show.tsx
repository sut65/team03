import React, { useEffect, useState } from "react";

import { Link as RouterLink } from "react-router-dom";

import Typography from "@mui/material/Typography";

import Button from "@mui/material/Button";

import Container from "@mui/material/Container";

import Box from "@mui/material/Box";
import Carousel from "react-material-ui-carousel";
import PersonAddIcon from '@mui/icons-material/PersonAdd';
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { grey } from "@mui/material/colors";
import {
  Paper,
  Grid,
  Dialog,
  DialogContent,
  TextField,
  Rating,
  AppBar,
  IconButton,
} from "@mui/material";
import Review_Save from "./Review_Save";
import { ReviewInterface } from "../../models/IReview";
import Moment from "moment";
import Toolbar from "@mui/material/Toolbar";
import Logo1 from "../../Image/LOGO.png"
import im10 from "../../Image/im10.jpg"
import im7 from "../../Image/im7.jpg"
import im9 from "../../Image/im9.png"


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
  },
});

function Review_Show() {
  Moment.locale("th");
  const [openForCreate, setOpenForCreate] = React.useState(false);
  const [review, setReview] = React.useState<ReviewInterface[]>([]);
  // const c = { width: '200x', height: '200px' };
  // const d = { width: '0', height: '0' };
  const imgSize = { width: "200px", height: "200px" };
  //const [myImageStyle, setMyImage] = React.useState({width: '200px', height: '200px' })
  const [openImage, setOpenImage] = React.useState(false);
  const [img, setimg] = React.useState({});

  //ตรวจสอบภาพ
  // const CheckImage = (item: ReviewInterface) => {
  //   if (item.Reviewimage == "")
  //   {setMyImage({width: '200px', height: '200px' })}
  // };
  const CheckImage = (item: ReviewInterface) => {
    if (item.Reviewimage != "") {
      return (
        <img
          src={`${item.Reviewimage}`}
          style={imgSize}
          onDoubleClick={() => OpenImageonCilck(item)}
        />
      );
    }
  };


  const OpenImageonCilck = (item: ReviewInterface) => {
    setOpenImage(true);
    setimg(item.Reviewimage);
  };

  const closeImage = () => {
    setOpenImage(false);
  };

  const getReview = async () => {
    const apiUrl = `http://localhost:8080/Reviews`;

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
        //   console.log(res.data); //show ข้อมูล
        if (res.data) {
          setReview(res.data);
          console.log(res.data);
        } else {
          // console.log("else");
        }
      });
  };

  useEffect(() => {
    getReview();
  }, []);

  function Item(props: any) {
    return (
         <img src={props.item.Image} width= "100%" height="600px"/>
      );
  }

  var Slider = [
    {
      Image: im9,
    },
    {
      Image: im7,
    },
    {
      Image: im10,
    },
  ];

  function ImageC() {
    return (
      <Carousel>
        {Slider.map((item, i) => (<Item key={i} item={item} />))}
      </Carousel>
    );
  }

  return (
    <ThemeProvider theme={themeshow}>
    <AppBar position="fixed">
      <Toolbar>
        <div>
          <img src={Logo1} width= "75px" height="75px"/>
        </div>
        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', width: '60%'}}>
          <Typography variant="h6" color="secondary" noWrap component="div" marginLeft={2}>
            <div >
              G03 Hotel
            </div>
          </Typography>
        </Box>
        <Box sx={{ display: 'flex', width: '50%'}}>
          <Button component={RouterLink} to="/Homeshow"  color="secondary" sx={{ display: 'flex', width: '10%'}}>Home</Button>
          <Button component={RouterLink} to="/Roomhome"  color="secondary" sx={{ display: 'flex', width: '10%'}}>Room</Button>
          <Button component={RouterLink} to="/RW"  color="secondary" sx={{ display: 'flex', width: '10%'}}>Review</Button>
          <Button component={RouterLink} to="/RW"  color="secondary" sx={{ display: 'flex', width: '10%'}}>ABOUT</Button>
          <Button component={RouterLink} to="/RW"  color="secondary" sx={{ display: 'flex', width: '13%'}}>Content</Button>
          <IconButton component={RouterLink} to="/customer/create" sx={{ display: 'flex', width: '11%'}} color="secondary" >
            <PersonAddIcon />
          </IconButton>
        </Box>
        <Box sx={{ display: 'flex', width: '11%'}}>
          <Button component={RouterLink} to="/home" variant="contained" color="secondary" >BOOK NOW</Button>
        </Box>
        
      </Toolbar>

    </AppBar>
    <div>
      <Container maxWidth="xl">
        {ImageC()}
        <div >
          <Grid
            container
            direction="row"
            columns={6}
            spacing={3}
            sx={{ mt: 2 }}
            
          >
            {review.map((item: ReviewInterface) => (
              <Grid item xs="auto" key={item.ID}>
                <Paper>
                  
                  <Grid>
                    <TextField
                      sx={{ "& fieldset": { border: "none" } }}
                      value={item.Customer.Email}
                    />
                  </Grid>

                  <Grid container columns={12}>
                    <Grid
                      display={"flex"}
                      justifyContent={"center"}
                      item
                      xs={5}
                    >
                      <Rating
                        value={item.Star}
                        size="medium"
                        sx={{ mt: 1.7 }}
                      />
                    </Grid>
                    <Grid item xs={7}>
                      <TextField
                        sx={{ "& fieldset": { border: "none" } }}
                        value={`${Moment(item.Reviewdate).format(
                          "DD/MM/YYYY HH:mm A"
                        )}`}
                      />
                    </Grid>
                  </Grid>

                  <Grid>
                    <TextField
                      fullWidth
                      multiline
                      sx={{ "& fieldset": { border: "none" } }}
                      value={item.Comment}
                    />
                  </Grid>

                  <Grid container display={"flex"} justifyContent={"center"}>
                    <Grid>
                      {CheckImage(item)}
                      {/* <img  src={`${item.Reviewimage}`} onError={() => CheckImage(item)} style={myImageStyle} onDoubleClick={() => OpenImageonCilck(item)}/> */}
                    </Grid>
                  </Grid>

                </Paper>
              </Grid>
            ))}
          </Grid>
        </div>
        <Dialog fullWidth maxWidth="md" open={openImage} onClose={closeImage}>
          <DialogContent>
            <img src={`${img}`} width="100%" height="100%" />
          </DialogContent>
        </Dialog>
      </Container>
    </div>
    </ThemeProvider>
  );
}

export default Review_Show;
