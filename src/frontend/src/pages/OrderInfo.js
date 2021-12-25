import React, {useEffect, useState} from 'react';
import {Container} from "react-bootstrap";
import {getOrder} from "../http/purchaseAPI";
import {useParams} from "react-router-dom";
import RecommendationsList from "../components/RecommendationsList";

const OrderInfo = () => {
    const orderID = useParams()
    const [order, setOrder] = useState({})

    useEffect(() => {
        getOrder(orderID.id).then(data => setOrder(data))
    },[])

    return (
        <Container className={"mt-5"}>
            {console.log(order)}
            <h2>Your order #{order.order_id} is submitted!</h2>
            <p>Customer name: <strong>{order.full_name}</strong></p>
            <p>Customer address: <strong>{order.address}</strong></p>
            <p>Total price: <strong>{order.total_price}</strong></p>
            <RecommendationsList />
        </Container>
    );
};

export default OrderInfo;