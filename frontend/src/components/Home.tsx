import { Box, Button, Container, Grid, Paper } from "@mui/material";
import * as React from "react";
import Carousel from "react-material-ui-carousel";
import im10 from "../Image/im10.jpg"
import im6 from "../Image/im6.jpeg"
import im7 from "../Image/im7.jpg"
import im9 from "../Image/im9.png"
import { mt } from "date-fns/locale";


function Home() {
  
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
  <div>
     <Grid>
      {ImageC()}
     </Grid >
     <div className="grid-con2">
        <div className="grid-item-21">
          <img src="https://cf.bstatic.com/xdata/images/hotel/max1280x900/19061811.jpg?k=d8542777976a7e63a981e6a0fef321472713bc27d6454055df45265eee200536&o=&hp=1" />
        </div>
     </div>
     <div className="grid-con">
        <div className="grid-item-1">
          <img src="https://content.r9cdn.net/himg/80/b3/46/expediav2-13266-f26277-844954.jpg" />
        </div>
        <div className="grid-item-2">
          <div className="grid-item-info">
            <div>
              <h3>Handmade piece</h3>
              <p>Welcome to Hotel G03</p>
              <a href="#" className="black-btn">
                Check out
              </a>
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
              <h3>wow</h3>
              <p>Welcome to Hotel G03</p>
              <a href="#" className="black-btn">
                Check out
              </a>
            </div>
          </div>
        </div>
        <div className="grid-item-4">
          <div>
            <img src="https://assets.langhamhotels.com/is/image/langhamhotelsstage/cdakl-rooms-chairman-suite-bedroom-1680-945:Medium?wid=1680&hei=944"></img>
          </div>
          <div className="grid-item-info">
            <div>
              <h3>wowhhghgh</h3>
              <p>Welcome to Hotel G03</p>
              <a href="#" className="black-btn">
                Check out
              </a>
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
  );
}

export default Home;
