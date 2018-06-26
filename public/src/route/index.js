import React, { Component } from "react";
import {
    BrowserRouter as Router,
    Route,
    Link,
    Redirect,
    Switch,
    withRouter
} from "react-router-dom";

// import Home from "./home";
import Login from "./login";

const Register = () => <h1>Register</h1>;

function RouterConfig() {
    return (
        <Router>
            <Switch>
                <Route exact path="/" component={Register} />
                <Route exact path="/login" component={Login} />
            </Switch>
        </Router>
    );
}

export default RouterConfig;