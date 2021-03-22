import react from "react"
import ReactDOM from "react-dom"
import { Link, NavLink } from "react-router-dom";
import {store} from './redux-module/redux-store'
import {addItem} from './redux-module/addItemAction'


class GridView extends react.Component{
    constructor(props){
        super(props)
    }

    addItems = (e) =>{
        e.preventDefault()
        const item = e.target.value
        const items = store.getState().cart.shoppingCart[item.id]
        
        console.log(`items retrieved from the store sentino`)
        console.log(items)
        console.log(item.id)

        if (items !== undefined){
          store.dispatch(addItem(item, items+1))
        }else {
          store.dispatch(addItem(item, 1))
        }
        
        
       
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
            <button value={this.props.item}onClick={this.addItems}>Add to the cart</button>
           </div>
        )
    }


}   
export default GridView;

