import React, { Component } from 'react'

class Cpu extends Component {
    render() {
        return (
            <div className="feature">
                <p>{Math.round(Number(this.props.usage) * 100)} %</p>
                <p>{this.props.name}</p>
            </div>
        )
    }
}

export default Cpu