import React, {useContext} from 'react';
import {Button, Container, Nav, Navbar} from "react-bootstrap";
import {NavLink, useNavigate} from "react-router-dom";
import {Context} from "../index"
import {observer} from "mobx-react-lite";

const NavBar = observer(() => {
    const navigate = useNavigate()
    const {cart} = useContext(Context)

    const count = cart.countAll()

    return (
        <Navbar bg="dark" variant="dark">
            <Container>
                <NavLink style={{color: "white"}} to={"/"}>Microservices-shop</NavLink>
                <Nav className="ml-auto" style={{color:"white"}}>
                    <Button variant={"outline-light"} onClick={() => navigate("/checkout")} disabled={count < 1}>Cart{count > 0 ? `: ${count} items` : " is empty"}</Button>
                </Nav>
            </Container>
        </Navbar>
    );
});

export default NavBar;