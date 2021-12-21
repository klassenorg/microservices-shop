import React from 'react';
import {Col, Container, Row} from "react-bootstrap";
import Cookies from 'universal-cookie';


const Footer = () => {
    const cookies = new Cookies()
    return (
        <Container fluid className={"bg-dark mt-5 py-5"}>
            <Row>
                <Col className={"align-content-center justify-content-center text-center text-white"}>© <a href={"https://github.com/klassenorg/microservices-shop"} style={{color: "white"}}>artkls</a></Col>
            </Row>
            <Row>
                <hr className={"border border-top border-1 border-light mt-5"}/>
            </Row>
            <Row>
                <Col className={"text-light"}>
                    USER_ID: {cookies.get("USER_ID")}
                </Col>

            </Row>
        </Container>
    );
};

export default Footer;