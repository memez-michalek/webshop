import { Component } from "react";
import GridView from "./Components/gridview"
import image from "./Components/lajcior.jpg"
import React from "react"
import ReactDOM from "react-dom"
import axios from "axios"
import { BrowserRouter as Router, Link, Route, Switch } from 'react-router-dom';
import productview from "./Components/productview"
import TopRow from "./Components/toprow"
import Cookies from 'universal-cookie'
import Cart from "./Components/cart"
class Main extends React.Component{
    constructor(props){
        super(props)
        this.state = {prods: [], cartItems: {}}
    }
    addItems = (e) =>{
      e.preventDefault()
      const itemId = e.target.value
      
      const cartIds = this.state.cartItems
      if (cartIds[itemId] === undefined){
          cartIds[itemId] = 1
      }else{
        let v = cartIds[itemId]
        console.log("cart ids "+v)
        cartIds[itemId] = v+1
      }


      console.log(cartIds)
      this.setState({cartItems: cartIds})
      console.log("cart items" +this.state.cartItems)
    }
    removeItems = (e) =>{
        e.preventDefault()
        let newcart = {}
        const itemid = e.target.value
        console.log(itemid)
        const prevCart = this.state.cartItems
        
        console.log("wisnia bakajoko")
       for(const item in prevCart){
          console.log("item value "+ item)
          if(item === itemid){
            if(prevCart[item] > 1){
              newcart[item] = prevCart[item] -1
              this.setState({cartItems: newcart})
            }else{
              this.setState({cartItems: newcart})

            }
          }else{
            newcart[item] = prevCart[item]
          }

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
          <Router>
          <div>
          <TopRow></TopRow>
          <Cart cartItems={this.state.cartItems} removeItem={this.removeItems}></Cart>
          </div>
          
          <div>

            

              {this.state.prods.map(item=><GridView item={item} addItems={this.addItems}></GridView>)}
              



          </div>




        </Router>


    
          
        )
    }

}


export default Main;