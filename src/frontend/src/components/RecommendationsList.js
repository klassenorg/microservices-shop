import {Container, Row} from "react-bootstrap";
import React, {useEffect, useState} from "react";
import RecommendationItem from "./RecommendationItem";
import {fetchRecommendations} from "../http/recommendationAPI";

const RecommendationsList = () => {
    const [recommendations, setRecommendations] = useState([])

    useEffect(() => {
        fetchRecommendations('4').then(data =>
            setRecommendations(data)
        )
    }, [])

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