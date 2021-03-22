
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
import Checkout from "./Components/views/Checkoutview"


import store from './Components/redux-module/redux-store'
import {addItem} from './Components/redux-module/addItemAction'
import {removeItems} from './Components/redux-module/removeItemAction'


import About from "./Components/views/AboutView"

class Main extends React.Component{
    constructor(props){
        super(props)
        this.state = {prods: []}
    }
    
    addItems = (e) =>{
      e.preventDefault()
      const itemId = e.target.value
      const items = store.getState().reducer.shoppingCart[itemId]
      console.log(`items retrieved from the store `)
      console.log(typeof(items))
      
      if (items !== undefined){
        store.dispatch(addItem(itemId, items+1))
      }else {
        store.dispatch(addItem(itemId, 1))
      }
      
      
     
    }
    removeItems = (e) =>{
        e.preventDefault()
        const itemid = e.target.value
        const currentItems = store.getState().reducer.shoppingCart
        const copy = currentItems
        if (currentItems[itemid]> 1){
          store.dispatch(removeItems(itemid, currentItems[itemid]-1))  
          
        }else{
          
          store.dispatch(removeItems(itemid, 0))  

        }
       
        
       
  
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
          <Cart removeItems={this.removeItems}></Cart>
          </div>
          
          <div>

            

              {this.state.prods.map(item=><GridView item={item} addItems={this.addItems}></GridView>)}
              
            


          </div>
          
          <Switch>
          
          <Route path="/about" component={About}></Route>
          <Route path="/product/:id" component={Productview}/>
          <Route path="/login" component={Login}/>
          <Route path="/register" component={Register}/>
          <Route path="/logout" component={Logout}/>
          
          </Switch>
         

        </div>


    
          
        )
    }

}


export default Main;