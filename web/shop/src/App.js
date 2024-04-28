import React, {Fragment} from 'react'
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
} from "react-router-dom"

import Home from "./pages/Home"
import Login from "./pages/Login";
import Register from "./pages/Register";
import BookDetail from "./pages/BookDetail";
import OrderList from "./pages/OrderList";
import Checkout from "./pages/Checkout";
import Nav from "./components/Nav";

function App() {
    return (
        <Fragment>
            <Router>
                <Nav/>
                <Switch>
                    <Route path="/Book/:id">
                        <BookDetail/>
                    </Route>
                    <Route path="/login">
                        <Login/>
                    </Route>
                    <Route path="/register">
                        <Register/>
                    </Route>
                    <Route path="/orders">
                        <OrderList/>
                    </Route>
                    <Route path="/checkout">
                        <Checkout/>
                    </Route>
                    <Route path="/">
                        <Home/>
                    </Route>
                </Switch>
            </Router>
        </Fragment>
    );
}

export default App;
