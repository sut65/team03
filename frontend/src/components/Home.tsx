import { Box, Container } from "@mui/material";
import * as React from "react";


function Home() {

    return (
        <div style={{
            backgroundImage: `url(${"https://img.freepik.com/free-photo/skyscrapers-sunset_1112-1870.jpg?w=1060&t=st=1673093915~exp=1673094515~hmac=cd71110c40d895948fd202933daada6d43b9ab1c5ec7595ce2cedcfac4b415fc"})`,
            backgroundSize: "cover",
            padding: 0,
            }}>
        <Container
            sx={{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                height: "88vh",
                overflow: "hidden",
            }}
        > 
        </Container>
    </div>
   
    );
   
   }
   
   export default Home;