import { Component } from "react";

import React from "react"
import ReactDOM from "react-dom"
import axios from "axios"






class Main extends React.Component{
    constructor(props){
        super(props)
        this.state = {}
    }

    render(){
        const details =  ()=>{
            axios.post("http://localhost:8080/init").then(function(response){
        return response
    }).catch(function(error){
        return error
    })
        }
        
        
        return(
           
            <div>
                <p>output {details}</p>

            </div>
        )
    }

}


export default Main;