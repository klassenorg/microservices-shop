import React from 'react';
import {Button, Card, Col} from "react-bootstrap";
import {useNavigate} from "react-router-dom";

const RecommendationItem = ({item}) => {
    const navigate = useNavigate()
    return (
        <Col>
            <Card>
                <Card.Img alt={item.name + " image"} style={{cursor: "pointer"}} src={item.image_path} onClick={() => navigate("/products/" + item.id)}/>
                <Card.Body>
                    <Card.Title>{item.name}</Card.Title>
                    <Card.Text><strong>{item.price} $</strong></Card.Text>
                    <Button className={"w-50"} onClick={() => navigate("/products/" + item.id)}>See more...</Button>
                </Card.Body>
            </Card>
        </Col>
    );
};

export default RecommendationItem;