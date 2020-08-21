import React from 'react';
import axios from 'axios'
import {config} from '../config'

export default class Todos extends React.Component {
    constructor(props){
        super(props);
        this.state = {
            todos: [],
            newtodo: {
                title: '',
                description: ''
            }
        }
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.markCompleted = this.markCompleted.bind(this);
        this.deleteTodo = this.deleteTodo.bind(this);
    }

    handleChange = e => {
        let newtodo = this.state.newtodo;
        let name = e.target.name;
        newtodo[name] = e.target.value;
        this.setState({newtodo})
    }

    handleSubmit = e => {
        e.preventDefault()
        axios.post(`${config.apiurl}`, this.state.newtodo)
            .then(response => {
                this.setState({
                    todos: [...this.state.todos, response.data],
                    newtodo: {
                        title: '',
                        description: ''
                    }
                })
            })
            .catch(err => {
                console.log(err)
            })
    }

    markCompleted = (e, id) => {
        let todoIndex = this.state.todos.findIndex(x => x._id == id)
        let todo = this.state.todos[todoIndex]
        todo.completed = true
        axios.put(`${config.apiurl}/${id}`, todo)
            .then(response => {
                let alltodos = this.state.todos
                let todos = alltodos.length === 1 ? [] : alltodos.splice(todoIndex, 1)
                this.setState({todos})
            })
            .catch(err => {
                console.log(err)
            })
    }

    deleteTodo = (e, id) => {
        let todoIndex = this.state.todos.findIndex(x => x._id == id)
        axios.delete(`${config.apiurl}/${id}`)
            .then(response => {
                let alltodos = this.state.todos
                let todos = alltodos.length === 1 ? [] : alltodos.splice(todoIndex, 1)
                this.setState({todos})
            })
            .catch(err => {
                console.log(err)
            })
    }

    componentDidMount(){
        axios.get(config.apiurl)
            .then(response => {
                this.setState({todos: response.data})
            })
            .catch(err => {
                console.log(err)
            })
    }
    render(){
        return (
            <div style={{display: 'flex'}}>
                <form onSubmit={this.handleSubmit}>
                    <div>
                        <label>Title</label>
                        <input
                            type={'text'}
                            name={'title'}
                            value={this.state.newtodo.title}
                            onChange={this.handleChange}
                            required
                        />
                    </div>
                    <div>
                        <label>Description</label>
                        <textarea
                            name={'description'}
                            value={this.state.newtodo.description}
                            onChange={this.handleChange}
                            required
                        />
                    </div>
                    <button type={'submit'}>Add Todo</button>
                </form>
                <table>
                    <thead>
                        <tr>
                            <th>#</th>
                            <th>Title</th>
                            <th>Description</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                    {this.state.todos && this.state.todos.map((el, index) =>
                        <tr key={el._id}>
                            <td>{index + 1}</td>
                            <td>{el.title}</td>
                            <td>{el.description}</td>
                            <td>
                                <button
                                    type={'button'}
                                    onClick={(e) => this.markCompleted(e, el._id)}
                                >
                                    Mark Completed
                                </button>
                                <button
                                    type={'button'}
                                    onClick={(e) => this.deleteTodo(e, el._id)}
                                >
                                    Delete
                                </button>
                            </td>
                        </tr>
                    )}
                    </tbody>
                </table>
            </div>
        )
    }
}