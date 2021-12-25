import React, {useContext, useEffect} from 'react';
import {Col, Container, Row} from "react-bootstrap";
import {Context} from "../index";
import CartItem from "../components/CartItem";
import CheckoutForm from "../components/CheckoutForm";
import {observer} from "mobx-react-lite";
import RecommendationsList from "../components/RecommendationsList";
import {fetchAllProducts} from "../http/catalogAPI";
import {fetchCart} from "../http/cartAPI";

const Cart = observer(() => {
    const {item} = useContext(Context)
    const {cart} = useContext(Context)


    useEffect(() => {
        fetchAllProducts().then(data =>
            item.setProducts(data.products)
        )
        fetchCart().then(data =>
            cart.setCart(data.cart)
        )
    }, [])

    const calculateTotalPrice = () => {
        let total = 0
        {Object.keys(cart.cart).map(itemId =>
            item.products.map(product =>
                product.id == itemId ? total+=(product.price * cart.cart[product.id]) : null
            )
        )}
        return total
    }

    return (
        <Container>
            <h2 className={"mt-3"}>Cart</h2>
            {cart.countAll() > 0
                ?
                <Row>
                    <Col md={8}>
                        {Object.keys(cart.cart).map(itemId =>
                            item.products.map(product =>
                                product.id == itemId ? <CartItem key={product.id} item={product}/> : null
                            )
                        )}
                    </Col>
                    <Col md={3}>
                        <h2>Total price: {calculateTotalPrice()}$</h2>
                        <CheckoutForm/>
                    </Col>
                </Row>
                :
                <Container>
                    <Row><h2>No item's in cart</h2></Row>
                    <Row><RecommendationsList /></Row>
                </Container>

            }
        </Container>
    );
});

export default Cart;