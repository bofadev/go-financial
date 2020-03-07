import React from 'react'
import '../css/reviewTableItem.css'
import 'bulma/css/bulma.css'

class ReviewTableItem extends React.Component {

    constructor(props){
        super(props)
    }

    onDelete = () => {
        this.props.onDeleteItem(this.props.review.Key)
    }
    

    render(){

        const colorAsset = '#33ffa5'
        const colorLiability = '#ffe633'
        let divStyle = {
            backgroundColor: ""
          };

        if(this.props.review.Type === "Asset"){
            divStyle.backgroundColor = colorAsset
        }
        else {
            divStyle.backgroundColor = colorLiability
        }

        return(            
            <tr style={divStyle}>
                <td>{this.props.review.Type}</td>
                <td>{this.props.review.Name}</td>
                <td>${this.props.review.Balance}</td>
                <td><button className="button is-danger" onClick={() => this.onDelete()}>delete</button></td>
            </tr>
        )
    }
}

export default ReviewTableItem