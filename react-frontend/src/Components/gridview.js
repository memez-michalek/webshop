import react from "react"
import ReactDOM from "react-dom"


class GridView extends react.Component{
    constructor(props){
        super(props)
        this.state = {}
    }


    render(){
        console.log(this.props);
        return(
            <div className="gridField">
                
                <img className="Image" src={this.props.imageUrl} alt="elo jd"></img>
                <h1 className="Name">{this.props.Name}</h1>
                <h2 className="Price">{this.props.Price}</h2>
                <p className="Category">{this.props.Category}</p>
                
            </div>
        )
    }


}
export default GridView;

