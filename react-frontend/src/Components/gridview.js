import react from "react"
import ReactDOM from "react-dom"


class Grid extends react.Component{
    constructor(props){
        super(props)
        this.state = {}
    }


    render(){

        return(
            <div>
                <img src={this.props.imageUrl} alt="elo jd"></img>

            </div>
        )
    }


}


