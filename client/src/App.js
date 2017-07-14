import React, { Component } from 'react'
import createDataMock from "./DataMock"
import Cpu from './components/Cpu'
import './App.css'

class App extends Component {
    constructor() {
        super()
        this.state = {
            data: undefined
        }
    }

    componentDidMount() {
        this.connection = new WebSocket("wss://echo.websocket.org")

        this.connection.onmessage = e => {
            this.update(e)
        }

        sendInterval(this.connection)
    }

    update(e) {
        this.setState({data: JSON.parse(e.data)})
    }

    render() {
        if(!this.state.data) {
            return (
                <h1>Loading</h1>
            )
        } else {
            return (
                <div>
                    <div id="stat_container">
                        <div id="cpu" className="stat_group">
                            {this.state.data.Cpus.map(cpu => <Cpu key={cpu.Name} name={cpu.Name} usage={cpu.Usage} />)}
                        </div>
                        <div id="disk" className="stat_group">
                        </div>
                        <div id="memory" className="stat_group">
                        </div>
                        <div id="temp" className="stat_group">
                        </div>
                        <div id="fans" className="stat_group">
                        </div>
                        <div id="volt" className="stat_group">
                        </div>
                    </div>
                </div>
            )
        }
    }
}

const sendInterval = function(connection) {
    setInterval( _ =>{
        connection.send(JSON.stringify(createDataMock()))
    }, 2000 )
}


export default App
