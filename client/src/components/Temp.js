import React, { Component } from 'react'

class Temp extends Component {
    render() {
        return (
            <div className="feature">
                <p>{this.props.label}: {this.props.value} &deg;C</p>
            </div>
        )
    }
}

export default Temp