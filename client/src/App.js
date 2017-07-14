import React, { Component } from 'react'
import createDataMock from "./DataMock"
import Cpu from './components/Cpu'
import Disk from './components/Disk'
import Memory from './components/Memory'
import Temp from './components/Temp'
import Fan from './components/Fan'
import Volt from './components/Volt'
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
                        <div id="cpu">
                            <h5>CPU</h5>
                            {this.state.data.Cpus.map(
                                cpu => <Cpu 
                                    key={cpu.Name} 
                                    name={cpu.Name} 
                                    usage={cpu.Usage} 
                                />
                            )}
                        </div>
                        <div id="disk">
                            <h5>ROM</h5>
                            {this.state.data.Disks.map(
                                disk => <Disk 
                                    key={disk.Name} 
                                    name={disk.Name} 
                                    total={disk.Total} 
                                    used={disk.Used} 
                                />
                            )}
                        </div>
                        <div id="memory">
                            <h5>RAM</h5>
                            <Memory 
                                total={this.state.data.Memory.Total}
                                used={this.state.data.Memory.Used}
                            />
                        </div>
                        <div id="temp">
                            <h5>Temperature</h5>
                            {this.state.data.Sensors.Temps.map(
                                temp => <Temp 
                                    key={temp.Label}
                                    label={temp.Lable}
                                    value={temp.Value}
                                />
                            )}
                        </div>
                        <div id="fans">
                            <h5>Fans</h5>
                            {this.state.data.Sensors.Fans.map(
                                fan => <Fan 
                                    key={fan.Label}
                                    label={fan.Lable}
                                    value={fan.Value}
                                />
                            )}
                        </div>
                        <div id="volt">
                            <h5>Voltage</h5>
                            {this.state.data.Sensors.Volt.map(
                                volt => <Volt 
                                    key={volt.Label}
                                    label={volt.Lable}
                                    value={volt.Value}
                                />
                            )}
                        </div>
                    </div>
                    
                    <div id="graph_container">
                        <p>testing</p>
                    </div>
                </div>
            )
        }
    }
}

const sendInterval = function(connection) {
    setInterval( _ =>{
        connection.send(JSON.stringify(createDataMock()))
    }, 1000 )
}


export default App
