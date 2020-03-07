import React from 'react'
import 'bulma/css/bulma.css'

class Form extends React.Component {


    constructor(props){
        super(props)
        this.state={
            type: "Asset",
            name: "",
            balance: -1
        }


        this.MIN_BALANCE = 0;
        this.MAX_BALANCE = 9999999;
    }

    isNameValid = (val) => {
        if(val === ""){
            return false
        }

        return true
    }
    isBalanceValid = (val) => {

        if(val === "" || val < this.MIN_BALANCE || val > this.MAX_BALANCE)
        {
            return false;
        }

        return true
    }

    GetNameHelpText = () => {

        if(this.isNameValid(this.state.name) == false)
        {
            return(
                <p class="help is-danger">
                    Please provide a name
                </p>
            )
        }
    }
    GetBalanceHelpText = () => {

        if(this.isBalanceValid(this.state.balance) == false)
        {
            return(
                <p class="help is-danger">
                    Value must be within range {this.MIN_BALANCE}...{this.MAX_BALANCE}
                </p>
            )
        }
    }

    clearNameBalance = () => {

        const inputName = document.getElementById("name-input")
        const inputBalance = document.getElementById("balance-input")
        const btn = document.getElementById("submit-btn")

        btn.setAttribute("disabled", "true")

        inputName.value = "";
        inputBalance.value = "";

        this.setState(prev =>{
            return({
                ...prev,
                name: "",
                balance: -1
            })
        })
    }

    onClickSubmit = () => {

        if(this.isNameValid(this.state.name) === false || this.isBalanceValid(this.state.balance) === false)
        {            
            return
        }
        
        console.log(`Submitting new review - [type:${this.state.type}] [name:${this.state.name}] [balance:${this.state.balance}]`)

        const url = `/review/add?type=${this.state.type}&name=${this.state.name}&balance=${this.state.balance}`;
        fetch(url)
        .then(r => r.json())
        .then(r => {
            console.log(r)
            this.props.setFormDirty()
        })

        this.clearNameBalance()
    }
    onSelectTypeChange = (event) => {
        const val = event.target.value;
        
        this.setState(prev =>{
            return{
                ...prev,
                type: val
            }
        })
    }
    onInputNameChange = (event) => {
        const val = event.target.value;

        const input = document.getElementById("name-input")
        const btn = document.getElementById("submit-btn")

        let everythingValid = true
        if(this.isNameValid(val) == false)
        {
            everythingValid = false
            input.classList.add("is-danger")
            btn.setAttribute("disabled", "true")
        }
        else {
            input.classList.remove("is-danger")
        }
        if(this.isBalanceValid(this.state.balance) == false)
        {
            everythingValid = false
            btn.setAttribute("disabled", "true")
        }

        if(everythingValid === true)
        {
            btn.removeAttribute("disabled")
            btn.onclick = () => {this.onClickSubmit();}
        }

        this.setState(prev =>{
            return{
                ...prev,
                name: val
            }
        })
    }
    onInputBalanceChange = (event) => {
        const val = event.target.value;

        const input = document.getElementById("balance-input")
        const btn = document.getElementById("submit-btn")

        let everythingValid = true
        if(this.isBalanceValid(val) == false)
        {
            everythingValid = false
            input.classList.add("is-danger")
            btn.setAttribute("disabled", "true")
        }
        else
        {
            input.classList.remove("is-danger")
        }        
        if(this.isNameValid(this.state.name) == false)
        {
            everythingValid = false
            btn.setAttribute("disabled", "true")
        }

        if(everythingValid === true)
        {
            btn.removeAttribute("disabled")
            btn.onclick = () => {this.onClickSubmit();}
        }

        this.setState(prev =>{
            return{
                ...prev,
                balance: val
            }
        })
    }

    render(){
        return(
            <div className="form box">

                <div class="field">
                    <label class="label">Type</label>
                    <div class="control">
                        <div class="select">
                            <select onChange={this.onSelectTypeChange}>
                                <option>Asset</option>
                                <option>Liability</option>
                            </select>
                        </div>
                    </div>
                </div>

                <label class="label">Name</label>
                <p class="control has-icons-left">
                    <input id="name-input" class="input is-danger" type="text" placeholder="Name" onChange={this.onInputNameChange}/>
                    <span class="icon is-small is-left">
                        <i class="fa fa-user-circle"></i>
                    </span>
                </p>
                {this.GetNameHelpText()}

                <label class="label">Balance</label>
                <p class="control has-icons-left">
                    <input id="balance-input" class="input is-danger" type="number" min="0" max="9999999" placeholder="Amount" onChange={this.onInputBalanceChange}/>
                    <span class="icon is-small is-left">
                        <i class="fa fa-dollar"></i>
                    </span>
                </p>
                {this.GetBalanceHelpText()}

                <div className="field">
                    <p className="control">
                        <button disabled id="submit-btn" className="button is-link" onclick={() => {console.log("yerp")}}>Submit</button>
                    </p>
                </div>

            </div>
        )
    }
}

export default Form