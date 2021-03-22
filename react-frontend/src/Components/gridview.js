import react from "react"
import ReactDOM from "react-dom"
import { Link, NavLink } from "react-router-dom";


class GridView extends react.Component{
    constructor(props){
        super(props)
    }
    render(){
        const link = "product/"+ this.props.item.id
        return(
            <div className="productContainer">
                <NavLink to={link}>
           
            <div className="gridField">   
                <img src={this.props.item.imageUrl}/>
                <h1 className="Name">{this.props.item.name}</h1>
                <h2 className="Price">{this.props.item.price} PLN</h2>
                <p className="Category">{this.props.item.category}</p>
            </div>

            </NavLink>
            <button value={this.props.item.id}onClick={this.props.addItems}>Add to the cart</button>
           </div>
        )
    }


}   
export default GridView;

