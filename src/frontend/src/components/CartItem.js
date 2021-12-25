import React, {useContext} from 'react';
import {Button, Col, Image, Row} from "react-bootstrap";
import {Context} from "../index";
import {observer} from "mobx-react-lite";
import {useNavigate} from "react-router-dom";

const CartItem = observer(({item}) => {
    const cart = useContext(Context)
    const navigate = useNavigate()

    return (
        <Row className={"shadow rounded-3 p-2 mb-2"}>
            <Col md={4}>
                <Image onClick={() => navigate("/products/"+item.id)} fluid rounded src={item.image_path} alt={item.name+" image"}/>
            </Col>
            <Col md={8}>
                <Row>
                    <Col>
                        <h4>{item.name}</h4>
                    </Col>
                    <Col>
                        <Button className={"w-100"} variant={"outline-danger"} onClick={() => cart.cart.removeFromCart(item.id)}>Remove from cart</Button>
                    </Col>
                </Row>
                <Row>
                    <Col>
                        <strong>{item.price} $</strong>
                    </Col>
                    <Col>
                        <Row>
                            <Col>Count: {cart.cart.count(item.id)}</Col>
                            <Col>Total: <strong>{cart.cart.count(item.id) * item.price}$</strong></Col>
                        </Row>
                    </Col>
                </Row>
            </Col>
        </Row>
    );
});

export default CartItem;