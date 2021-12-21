import React from 'react';
import {Button, Card, Col, Container, Image, Row} from "react-bootstrap";
import Recommendations from "../components/RecommendationsList";

const ItemPage = () => {
    const item = {"id":100000,"name":"Lemon","description":"The lemon is a round, slightly elongated fruit, it has a strong and resistant skin, with an intense bright yellow colour when it is totaly ripe, giving off a special aroma when it is cut.","price":30,"image_path":"/img/products/100000.jpeg"}
        return (
        <Container className={"py-5"}>
            <Row>
                <Col md={5}>
                    <Image fluid rounded src={item.image_path}/>
                </Col>
                <Col>
                    <Card>
                        <Card.Body>
                            <Card.Title>{item.name}</Card.Title>
                            <Card.Subtitle>{item.price} $</Card.Subtitle>
                            <Card.Text className={"mt-3"}>
                                {item.description}
                            </Card.Text>

                            <Button variant="outline-success">Add to cart</Button>
                        </Card.Body>
                    </Card>
                </Col>
            </Row>
            <Row className={"mt-5"}>
                <Recommendations />
            </Row>
        </Container>
    );
};

export default ItemPage;