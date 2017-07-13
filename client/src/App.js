import React, { Component } from 'react';
import createDataMock from "./DataMock"

class App extends Component {
    constructor() {
        super()
        this.state = {
            data: {
                cpus: [],
                disks: [],
                memory: {},
                sensors: {}
            }
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
        return (
            <div>
                {JSON.stringify(this.state.data)}
            </div>
        )
    }
}

const sendInterval = function(connection) {
    setInterval( _ =>{
        connection.send(JSON.stringify(createDataMock()))
    }, 2000 )
}


export default App;
