import { Component } from "react";

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
        this.setState({prods: data})
      }).catch(err=>{

        console.log(err)
      })

    }
    render(){
        
        return(
           
            <div>
                <p>{this.state.prods}</p>

            </div>
        )
    }

}


export default Main;