import React from 'react';
import {Routes, Route} from 'react-router-dom'
import Catalog from "../pages/Catalog";
import ItemPage from "../pages/ItemPage";
import Cart from "../pages/Cart";
import NotFound from "../pages/NotFound";
import OrderInfo from "../pages/OrderInfo";

const AppRouter = () => {
    return (
        <Routes>
            <Route path={"/"} element={<Catalog />} exact />
            <Route path={"/products/:id"} element={<ItemPage />} exact />
            <Route path={"/checkout"} element={<Cart />} exact />
            <Route path={"/order/:id"} element={<OrderInfo />} exact />
            <Route path={"*"} element={<NotFound />} exact />
        </Routes>
    );
};

export default AppRouter;