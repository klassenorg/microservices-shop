import React from 'react';
import {Button, Card, Col} from "react-bootstrap";
import {useNavigate} from "react-router-dom";

const addToCart = async (id) => {

    try {
        console.log(`add item ${id} to cart`)
    } catch (e) {
        console.log(e)
    }
}

const Item = ({item}) => {
    const navigate = useNavigate()
    return (
        <Col md={3} sm={6} className={"mt-3"}>
            <Card style={{ width: '16rem'}} >
                <Card.Img style={{cursor: "pointer"}} variant="top" src={item.image_path} onClick={() => navigate("/products/"+item.id)}/>
                <Card.Body>
                    <Card.Title>{item.name}</Card.Title>
                    <Card.Text>
                        <strong>{item.price} $</strong>
                    </Card.Text>
                    <Col>
                        <Button variant="outline-info" onClick={() => navigate("/products/"+item.id)}>See more...</Button>
                    </Col>
                    <Col>
                        <Button className={"mt-1"} variant="outline-info" onClick={() => addToCart(item.id)}>Add to cart</Button>
                    </Col>
                </Card.Body>
            </Card>
        </Col>
    );
};

export default Item;