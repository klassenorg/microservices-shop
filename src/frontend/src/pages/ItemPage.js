import React, {useContext, useEffect, useState} from 'react';
import {Button, Card, Col, Container, Image, Row} from "react-bootstrap";
import Recommendations from "../components/RecommendationsList";
import {useParams} from "react-router-dom";
import {Context} from "../index";
import {fetchOneProduct} from "../http/catalogAPI";
import {fetchCart} from "../http/cartAPI";

const ItemPage = () => {
    const productId = useParams()
    const {cart} = useContext(Context)
    const [product, setProduct] = useState({})

    useEffect(() => {
        fetchCart().then(data =>
            cart.setCart(data.cart)
        )
        fetchOneProduct(productId.id).then(data => setProduct(data))
        window.scrollTo(0, 0)
    }, [productId])



        return (
        <Container className={"py-5"}>
            <Row>
                <Col md={5}>
                    <Image fluid rounded src={product.image_path}/>
                </Col>
                <Col>
                    <Card>
                        <Card.Body>
                            <Card.Title>{product.name}</Card.Title>
                            <Card.Subtitle>{product.price} $</Card.Subtitle>
                            <Card.Text className={"mt-3"}>
                                {product.description}
                            </Card.Text>

                            <Button variant="outline-success" onClick={() => cart.addToCart(product.id, 1)}>Add to cart</Button>
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