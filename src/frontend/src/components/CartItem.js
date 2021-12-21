import React, {useContext} from 'react';
import {Button, Col, Image, Row} from "react-bootstrap";
import {Context} from "../index";
import {observer} from "mobx-react-lite";

const CartItem = observer(({item}) => {
    const cart = useContext(Context)
    return (
        <Row className={"shadow rounded-3 p-2 mb-2"}>
            <Col md={4}>
                <a href={"/products/" + item.id}><Image fluid rounded src={item.image_path} alt={item.name+" image"}/></a>
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
                        Count: {cart.cart.count(item.id)}
                    </Col>
                </Row>
            </Col>
        </Row>
    );
});

export default CartItem;