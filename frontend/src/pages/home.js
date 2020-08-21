import React from 'react'
import {connect} from 'react-redux'
import Main from '../layouts/main'
import Header from "../components/header";
import Todos from "../components/todos";

class Home extends React.Component {
    constructor(props){
        super(props);
        this.state = {
            title: ''
        }
    }

    addOne = () => {
        this.props.increasecount()
    }

    removeOne = () => {
        this.props.decreasecount()
    }

    componentDidMount(){
        this.setState({title: 'Home', count: this.props.count})
    }
    render(){
        return (
            <Main>
                <Header
                    title={this.state.title}
                />
                <h1>{this.state.title}</h1>
                <p>An example of head elements coming from state.</p>

                <div style={{display: 'flex', alignItems: 'center', alignContent: 'center'}}>
                    <button onClick={this.removeOne}>Decrease</button>
                    <p>{this.props.counter}</p>
                    <button onClick={this.addOne}>Increase</button>
                </div>

                <Todos />
            </Main>
        )
    }
}

const mapStateToProps = state => {
    return {
        counter: state.counter
    }
}

const mapDispatchToProps = dispatch => {
    return {
        increasecount: () => dispatch({
            type: 'INCREASE_COUNT'
        }),
        decreasecount: () => dispatch({
            type: 'DECREASE_COUNT'
        })
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(Home);
