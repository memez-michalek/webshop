import logo from './logo.svg';
import './App.css';
import Main from './File'
import {Route, Link, BrowserRouter as Router, Switch} from "react-router-dom"
import Productview from "./Components/productview"
function App() {
  return (
    <Router>
    <Switch>  
      <Route exact path="/">
        <Main></Main>
      </Route>
      <Route path="/product/:id" component={Productview}/>


    </Switch>
    </Router>

   



  );
}

export default App;
