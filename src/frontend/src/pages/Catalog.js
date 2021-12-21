import React from 'react';
import {Container, Row} from "react-bootstrap";
import ItemList from "../components/ItemList";

const Catalog = () => {
    return (
        <Container>
            <Row className={"mt-2"}>
                <ItemList />
            </Row>
        </Container>
    );
};

export default Catalog;