import React from 'react'
import '../css/mainContent.css'
import 'bulma/css/bulma.css'
import ReviewTableItem from './reviewTableItem'
import Form from './Form'

class MainContent extends React.Component {

    constructor(props){
        super(props)
        this.state = {
            reviews: [],
            netWorth: 0,
            totalAssets: 0,
            totalLiabilities: 0
        }
    }

    getData = () => {        
                
        fetch("/reviews/get")
        .then(r => r.json())
        .then(data => {
            this.setState(prev => {
                return {
                    ...prev,
                    reviews: data
                }
            })

            fetch("/getTotals")
            .then(r => r.json())
            .then(data => {
                this.setState(prev => {
                    return {
                        ...prev,
                        netWorth: data.NetWorth,
                        totalAssets: data.TotalAssets,
                        totalLiabilities: data.TotalLiabilities
                    }
                })
            })
        })
    }
    onFormDirty = () => {
        this.getData()
    }
    onDeleteItem = (key) => {

        if(window.confirm("Delete the entry?") === false)
        {
            return;
        }

        console.log(`Deleting review - [key:${key}]`)

        const url = `/review/delete?key=${key}`;
        fetch(url)
        .then(r => r.json())
        .then(r => {
            console.log(r)
            this.getData()
        })
    }

    renderTableRow = (item) => {
        return <ReviewTableItem review={item} onDeleteItem={key => this.onDeleteItem(key)}/>
    }
    renderTableRows = () => {        
        return this.state.reviews.map(this.renderTableRow)
    }

    componentDidMount(){
        this.getData()
    }
    

    render(){
        return(
            <div className="main-content">
                <div className="left-content">
                    <table className="table is-striped is-hoverable">
                        <thead>
                            <tr>
                                <th>Type</th>
                                <th>Name</th>
                                <th>Balance</th>
                            </tr>
                        </thead>
                        <tbody>                                
                            {this.renderTableRows()}
                        </tbody>
                    </table>
                </div>
                <div className="right-content">
                    <div className="follow">

                        <table className="table">
                            <thead>
                                <tr>
                                    <th>Net Worth</th>
                                    <th>Assets</th>
                                    <th>Liabilities</th>
                                </tr>
                            </thead>
                            <tbody>
                                    <tr>
                                        <td>${this.state.netWorth}</td>
                                        <td>${this.state.totalAssets}</td>
                                        <td>${this.state.totalLiabilities}</td>
                                    </tr>
                            </tbody>
                        </table>
                        <hr />
                        
                        <p className="title right-title">Add New Item</p>
                        <Form setFormDirty={() => this.onFormDirty()}/>
                    </div>
                </div>
            </div>
        )
    }
}

export default MainContent