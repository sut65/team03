import * as React from "react";
import Carousel from "react-material-ui-carousel";
import im10 from "../Image/im10.jpg"
import im7 from "../Image/im7.jpg"
import im9 from "../Image/im9.png"
import Logo1 from "../Image/LOGO.png"

import Toolbar from "@mui/material/Toolbar";

import Typography from "@mui/material/Typography";

import { Link as RouterLink } from "react-router-dom";

import { createTheme, styled, useTheme } from "@mui/material/styles";
import { grey } from "@mui/material/colors";

import { AppBar, Box, Button, Grid, ThemeProvider } from "@mui/material";

const bgnavbar = createTheme({
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


function About() {
  const themep = useTheme();
  const [open, setOpen] = React.useState(false);

  //เรียกใช้งานรูปภาพเป็นสไลด์ๆ
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
//เรียกใช้ Carousel
  function ImageC() {
    return (
      <Carousel>
        {Slider.map((item, i) => (<Item key={i} item={item} />))}
      </Carousel>
    );
  }

  return (
    <ThemeProvider theme={bgnavbar}>
    <AppBar position="fixed">
      <Toolbar>
        <div>
          <img src={Logo1} width= "75px" height="75px"/>
        </div>
        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', width: '50%'}}>
          <Typography variant="h6" color="secondary" noWrap component="div" marginLeft={2}>
            <div >
              G03 Hotel
            </div>
          </Typography>
        </Box>
        <Box sx={{ display: 'flex', width: '35%'}}>
          <Button component={RouterLink} to="/Homeshow"  color="secondary" sx={{ display: 'flex', width: '10%'}} >Home</Button>
          <Button component={RouterLink} to="/Roomhome"  color="secondary" sx={{ display: 'flex', width: '10%'}}>Room</Button>
          <Button component={RouterLink} to="/RW"  color="secondary" sx={{ display: 'flex', width: '10%'}}>Review</Button>
          <Button component={RouterLink} to="/About"  color="secondary" sx={{ display: 'flex', width: '10%'}}>ABOUT</Button>

        </Box>
        <Box sx={{ display: 'flex', width: '6%'}}>
        <Button component={RouterLink} to="/home" variant="contained" color="secondary" >LOGIN</Button>
       
        </Box>
        <Box sx={{ display: 'flex', width: '10%'}}>
        <Button component={RouterLink} to="/home" variant="contained" color="secondary" >BOOK NOW</Button>
        </Box>
        
      </Toolbar>

    </AppBar>
    <div>
   <Grid>
    {ImageC()}
   </Grid >
   
   <div className="grid-con2">
      <div className="grid-item-21">
        <img src="https://i.pinimg.com/originals/ee/53/37/ee5337111b704bd8117495af62c8a1a7.jpg" />
      </div>
      <div className="grid-item-22">
        <div className="grid-item-info3">
          <h1>WELCOME TO THE G03 HOTEL</h1>
        </div>
      </div>
      <div className="grid-item-23">
        <div className="grid-item-info4">
          <h2>Our hotel has good service such as friendly staff and service with a smile. Parking service, Roomservice , Spa&Massage, Elegant restaurant with well decorated. Variety of food for you to taste. Complete hotel facilities Additional services available to you With 24-hour service and excellent security system. With a key card scanning door system There are security guards who are always ready to help you. You will be impressed. Satisfaction in various fields in our hotel</h2>
        </div>
      </div>
      <div className="grid-item-24">
        <img src="https://media-cdn.tripadvisor.com/media/photo-s/1c/ff/66/8a/modern-thai-cuisine-perfectly.jpg" />
      </div>
      <div className="grid-item-25">
        <img src="https://media-cdn.tripadvisor.com/media/photo-s/0c/17/73/d1/legrande-lounge.jpg" />
      </div>
   </div>

   <div className="grid-con">
      <div className="grid-item-1">
        <img src="https://content.r9cdn.net/himg/80/b3/46/expediav2-13266-f26277-844954.jpg" />
      </div>
      <div className="grid-item-2">
        <div className="grid-item-info">
          <div>
            <h2>RELAX AND ENJOY</h2>
            <p>with the G03 Hotel.</p>
          </div>
        </div>
        <div>
          <img src="https://files.guidedanmark.org/files/382/242733_Hotel_Ottilia__Brchner_Hotels_PR.jpg"></img>
        </div>
      </div>
      <div className="grid-item-3">
        <div>
          <img src="https://www.nzbusinesstraveller.co.nz/wp-content/uploads/2022/08/Copy-of-Chairman-Suite-Lounge-1024x683.jpg"></img>
        </div>
        <div className="grid-item-info2">
          <div>
            <h2>EXCELLENT SERVICE</h2>
            <p>Service is available 24 hours a day.</p>
            <p>Always serve customers with a smile.</p>
          </div>
        </div>
      </div>
      <div className="grid-item-4">
        <div>
          <img src="https://assets.langhamhotels.com/is/image/langhamhotelsstage/cdakl-rooms-chairman-suite-bedroom-1680-945:Medium?wid=1680&hei=944"></img>
        </div>
        <div className="grid-item-info">
          <div>
            <h2>LUXURY BY THE SEA</h2>
            <p>Meet the beautiful sea view.</p>
          </div>
        </div>
      </div>
      <div className="grid-item-5">
        <div>
          <img src="https://www.thedenizen.co.nz/wp-content/uploads/2021/01/Opera-Suite-Lounge-1.jpg"></img>
        </div>
      </div>
  </div>
</div>
</ThemeProvider>
  );
}

export default About;