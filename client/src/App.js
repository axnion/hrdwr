import React, { Component } from 'react';

class App extends Component {
    constructor() {
        super()
        this.state = {
            data: null
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
        this.setState({data: e.data})
    }

    render() {
        return (
            <div>
                {this.state.data}
            </div>
        )
    }
}

const sendInterval = function(connection) {
    setInterval( _ =>{
        connection.send("{cpu: 300}")
    }, 2000 )
}

// TODO: Create stats mocking wihhc is sent to echo.websocket.org for testing

export default App;
