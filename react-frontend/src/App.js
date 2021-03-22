import logo from './logo.svg';
import './App.css';
import Main from './File'
import {Route, Link, BrowserRouter as Router, Switch} from "react-router-dom"
import Productview from "./Components/productview"
import Login from "./Components/Loginview"
import Register from "./Components/Registerview"
import Logout from "./Components/logoutview"
import { render } from 'react-dom';
import { resetWarningCache } from 'prop-types';
import react from "react"
import About from './Components/AboutView';
import Contact from './Components/ContactView';
import Checkout from './Components/Checkoutview'
class App extends react.Component{

  

  render(){
  return (
    <Router>
      
    <Switch>  
      <Route exact path="/">
        <Main></Main>
      </Route>
      <Route path="/product/:id" component={Productview}/>
      <Route path="/login" component={Login}/>
      <Route path="/register" component={Register}/>
      <Route path="/logout" component={Logout}></Route>
      <Route path="/about" component={About}></Route>
      <Route path="/contact" component={Contact}></Route>
      <Route path="/checkout" component={Checkout}></Route>
    </Switch>
    </Router>

   



  );
  }
}

export default App;
