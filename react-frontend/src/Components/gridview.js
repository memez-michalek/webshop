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
        const item = this.props.item
      
        const items = store.getState().cart.shoppingCart[item.id]
        console.log("items")
        console.log(items)

        if (items !== undefined){
          store.dispatch(addItem(item,item.id, items["amount"]+1))
        }else {
          store.dispatch(addItem(item,item.id, 1))
        }
        
        
       
      }




    render(){
        const link = "product/"+ this.props.item.id
        const AddItem = {id: this.props.item.id, name: this.props.item.name, image: this.props.item.imageUrl, price: this.props.item.price}
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
            
            <button value={this.props.item} onClick={this.addItems}>Add to the cart</button>
           </div>
        )
    }


}   
export default GridView;

