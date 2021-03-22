
import GridView from "./Components/gridview"
import image from "./Components/lajcior.jpg"
import React from "react"
import ReactDOM from "react-dom"
import axios from "axios"
import { BrowserRouter as Router, Link, Route, Switch } from 'react-router-dom';

import TopRow from "./Components/toprow"
import Cookies from 'universal-cookie'
import Cart from "./Components/cart"
import Logout from "./Components/logoutview"
import Productview from "./Components/productview"
import Login from "./Components/Loginview"

import Register from "./Components/Registerview"
import Checkout from "./Components/Checkoutview"


import {store} from './Components/redux-module/redux-store'
import {addItem} from './Components/redux-module/addItemAction'



import About from "./Components/AboutView"

class Main extends React.Component{
    constructor(props){
        super(props)
        this.state = {prods: []}
    }
    
  
  


    componentDidMount() {
      axios.post("http://localhost:8080/init").then(response=>{
    
        const data = response.data
        console.log(data)
        const obj = JSON.parse(data)
        console.log(obj)
        this.setState({prods: obj})
      }).catch(err=>{

        console.log(err)
      })

    }
 


    render(){
        const cookie = new Cookies()
        //console.log(cookie.get("token"))
        console.log(this.state.prods)
        return(
          <div>
          <div>
          <TopRow></TopRow>
          <Cart></Cart>
          </div>
          
          <div>

            

              {this.state.prods.map(item=><GridView item={item} ></GridView>)}
              
            


          </div>
       
         

        </div>


    
          
        )
    }

}


export default Main;