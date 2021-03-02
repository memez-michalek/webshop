import { Component } from "react";
import GridView, {Grid} from "./Components/gridview"

import React from "react"
import ReactDOM from "react-dom"
import axios from "axios"






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
           
            <div className="jd">
              
            
            {this.state.prods.map(item=><GridView ImageUrl={item.imageUrl} Name={item.name} Price={item.price} Category={item.category} ></GridView>)}


            </div>
        )
    }

}


export default Main;