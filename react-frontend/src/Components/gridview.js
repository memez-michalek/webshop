import react from "react"
import ReactDOM from "react-dom"
import { Link } from "react-router-dom";


class GridView extends react.Component{

   
    render(){
        
        return(
            
            <a href={`http://localhost:3000/product/${this.props.Id}`}>
            <div className="gridField">   
                <img src={this.props.ImageUrl}/>
                <h1 className="Name">{this.props.Name}</h1>
                <h2 className="Price">{this.props.Price} PLN</h2>
                <p className="Category">{this.props.Category}</p>
                
            </div>
            </a>
        )
    }


}   
export default GridView;

