import React, {useContext} from 'react';
import {Col, Container, Row} from "react-bootstrap";
import {Context} from "../index";
import CartItem from "../components/CartItem";
import CheckoutForm from "../components/CheckoutForm";
import {observer} from "mobx-react-lite";

const Cart = observer(() => {
    const {item} = useContext(Context)
    const {cart} = useContext(Context)
    return (
        <Container>
            <h2 className={"mt-3"}>Cart</h2>
            <Row>
                <Col md={8}>
                    {Object.keys(cart.cart).map(itemId =>
                        item.products.map(product =>
                             product.id == itemId ? <CartItem key={product.id} item={product}/> : null
                        )
                    )}
                </Col>
                <Col md={3}>
                    <CheckoutForm/>
                </Col>
            </Row>
        </Container>
    );
});

export default Cart;