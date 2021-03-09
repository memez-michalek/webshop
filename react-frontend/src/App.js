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
      <Route></Route>
    </Switch>
    </Router>

   



  );
  }
}

export default App;
