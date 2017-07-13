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
                <div id="stat_container">
                    <div id="cpu" className="stat_group">
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
                {this.state.data}
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
