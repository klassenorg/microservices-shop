import {Container, Row} from "react-bootstrap";
import React from "react";
import RecommendationItem from "./RecommendationItem";

const RecommendationsList = () => {
    const recommendations = [{
        "id": 100003,
        "name": "Qiwi",
        "description": "Qiwi is green inside.",
        "price": 55,
        "image_path": "/img/products/100003.jpeg"
    }, {
        "id": 100001,
        "name": "Apple",
        "description": "Apples are the ideal fruit to eat at any time, having a positive role in the achievement of nourish balance. Their skin may be green, yellow or reddish, and the meat taste ranges from a bitter to sweet flavour.",
        "price": 15,
        "image_path": "/img/products/100001.jpeg"
    }, {
        "id": 100000,
        "name": "Lemon",
        "description": "The lemon is a round, slightly elongated fruit, it has a strong and resistant skin, with an intense bright yellow colour when it is totaly ripe, giving off a special aroma when it is cut.",
        "price": 30,
        "image_path": "/img/products/100000.jpeg"
    }, {
        "id": 100002,
        "name": "Orange",
        "description": "Orange is orange.",
        "price": 25,
        "image_path": "/img/products/100002.jpeg"
    }]

    return (
        <Container>
            <hr className="border border-top border-1 border-dark mt-5"/>
            <Row className={"align-content-center justify-content-center text-center mt-1"}>
                <h3>You also may be interested</h3>
            </Row>
            <Row md={4} s={2}>
                {recommendations.map(recommendation =>
                    <RecommendationItem key={recommendation.id} item={recommendation}/>
                )}
            </Row>

        </Container>
    );
};
export default RecommendationsList;