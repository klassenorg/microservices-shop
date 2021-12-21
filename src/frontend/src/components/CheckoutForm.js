import React, {useState} from 'react';
import {Button, Form} from "react-bootstrap";
import Cookies from "universal-cookie";
import {useNavigate} from "react-router-dom";

const CheckoutForm = () => {
    const cookies = new Cookies()
    const navigate = useNavigate()

    const [email, setEmail] = useState('klassenorg@gmail.com')
    const [phone, setPhone] = useState('+79774164899')
    const [name, setName] = useState('Arthur Klassen')
    const [city, setCity] = useState('Moscow')
    const [address, setAddress] = useState('ul. Pushkina, d. Kolotushkina, kv. 1')
    const [cardNumber, setCardNumber] = useState('1234123412341234')
    const [cardExpiration, setCardExpiration] = useState('12/34')
    const [cardCVV, setCardCVV] = useState('123')

    const createOrder = async () => {
        let orderId
        try {
            console.log(cookies.get("USER_ID"), email, phone, name, city, address, cardNumber, cardExpiration, cardCVV)
            orderId = "123456"
        } catch (e) {

        }

        navigate("/order/" + orderId)
    }

    return (
        <Form className={"d-flex flex-column"} autoComplete={"off"}>
            <Form.Text>Customer information</Form.Text>
            <Form.Control
                className={"mt-3"}
                placeholder={"Enter email..."}
                value={email}
                onChange={e => setEmail(e.target.value)}
                type={"email"}
                required
            />
            <Form.Control
                className={"mt-3"}
                placeholder={"Enter phone..."}
                value={phone}
                onChange={e => setPhone(e.target.value)}
                type={"tel"}
                required
            />
            <Form.Control
                className={"mt-3"}
                placeholder={"Enter name..."}
                value={name}
                onChange={e => setName(e.target.value)}
                required
            />
            <Form.Control
                className={"mt-3"}
                placeholder={"Enter city..."}
                value={city}
                onChange={e => setCity(e.target.value)}
                required
            />
            <Form.Control
                className={"mt-3"}
                placeholder={"Enter address..."}
                value={address}
                onChange={e => setAddress(e.target.value)}
                required
            />
            <Form.Control
                className={"mt-3"}
                placeholder={"Enter card number..."}
                value={cardNumber}
                onChange={e => setCardNumber(e.target.value)}
                required
            />
            <Form.Control
                className={"mt-3"}
                placeholder={"Enter card expiration date..."}
                value={cardExpiration}
                onChange={e => setCardExpiration(e.target.value)}
                required
            />
            <Form.Control
                className={"mt-3"}
                placeholder={"Enter card CVV..."}
                value={cardCVV}
                onChange={e => setCardCVV(e.target.value)}
                type={"password"}
                required
            />
            <Button className={"mt-3"} variant={"outline-success"} onClick={createOrder}>Create order</Button>
        </Form>
    );
};

export default CheckoutForm;