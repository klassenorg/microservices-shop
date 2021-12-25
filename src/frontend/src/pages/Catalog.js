import React, {useContext, useEffect} from 'react';
import {Container, Row} from "react-bootstrap";
import ItemList from "../components/ItemList";
import {observer} from "mobx-react-lite";
import {Context} from "../index";
import {fetchAllProducts} from "../http/catalogAPI";
import {fetchCart} from "../http/cartAPI";

const Catalog = observer(() => {
    const {item} = useContext(Context)
    const {cart} = useContext(Context)
    useEffect(() => {
        fetchCart().then(data =>
            cart.setCart(data.cart)
        )
        fetchAllProducts().then(data =>
            item.setProducts(data.products)
        )
    }, [])

    return (
        <Container>
            <Row className={"mt-2"}>
                <ItemList />
            </Row>
        </Container>
    );
});

export default Catalog;