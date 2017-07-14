import React, { Component } from 'react'

class Memory extends Component {
    render() {
        return (
            <div className="feature">
                <p>Total: {this.props.total} kB</p>
                <p>Used: {this.props.used} kB</p>
            </div>
        )
    }
}

export default Memory