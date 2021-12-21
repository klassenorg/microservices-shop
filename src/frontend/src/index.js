import React, {createContext} from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import ItemStore from "./store/item";
import CartStore from "./store/cart";

export const Context = createContext(null)

ReactDOM.render(
    <Context.Provider value={
        {
            item: new ItemStore(),
            cart: new CartStore(),
        }
    }>
        <React.StrictMode>
            <App />
        </React.StrictMode>
    </Context.Provider>,
        document.getElementById('root')

);
