import React from 'react';
import {Button, Col, Container, Image, Row} from "react-bootstrap";
import {useNavigate} from "react-router-dom";

const NotFound = () => {
    const navigate = useNavigate()
    return (
        <Container>
            <Row className={"mt-3"}>
                <Col className={"d-flex justify-content-center align-content-center"}>404 Page not found</Col>
                <Col className={"d-flex justify-content-center align-content-center"}>
                    <Button className={"d-flex justify-content-center align-content-center"} onClick={() => navigate("/")}>Go to main page</Button>
                </Col>
            </Row>
            <Row className={"mt-3"}>
                <Col><Image src={"/img/products/100000.jpeg"}/></Col>
                <Col>{/*TODO:add something*/}</Col>
            </Row>
        </Container>
    );
};

export default NotFound;