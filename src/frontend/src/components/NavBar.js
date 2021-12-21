import React, {useContext} from 'react';
import {Button, Container, Nav, Navbar} from "react-bootstrap";
import {NavLink, useNavigate} from "react-router-dom";
import {Context} from "../index"
import {observer} from "mobx-react-lite";

const NavBar = observer(() => {
    const navigate = useNavigate()
    const {cart} = useContext(Context)
    return (
        <Navbar bg="dark" variant="dark">
            <Container>
                <NavLink style={{color: "white"}} to={"/"}>Microservices-shop</NavLink>
                <Nav className="ml-auto" style={{color:"white"}}>
                    <Button variant={"outline-light"} onClick={() => navigate("/checkout")}>Cart: {sumValues(cart.cart)} items</Button>
                </Nav>
            </Container>
        </Navbar>
    );
});

export default NavBar;

function sumValues(obj){
    let val
    Object.keys(obj).length > 0 ? val = Object.values(obj).reduce((a, b) => parseInt(a) + parseInt(b)) : val = 0
    return val
}
