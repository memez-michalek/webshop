import { Component } from "react";
import GridView, {Grid} from "./Components/gridview"
import image from "./Components/lajcior.jpg"
import React from "react"
import ReactDOM from "react-dom"
import axios from "axios"
import { BrowserRouter as Router, Link, Route, Switch } from 'react-router-dom';
import productview from "./Components/productview"




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
        
        return(
          <Router>
            <div>
             
                {this.state.prods.map(item=><GridView Id={item.id}ImageUrl={item.imageUrl} Name={item.name} Price={item.price} Category={item.category} ></GridView>)}
                



            </div>




          </Router>
          
        )
    }

}


export default Main;