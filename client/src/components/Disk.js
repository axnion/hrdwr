import React, { Component } from 'react'

class Disk extends Component {
    render() {
        return (
            <div className="feature">
                <p>Total: {this.props.total} kB</p>
                <p>Used: {this.props.used} kB</p>
                <p>{this.props.name}</p>
            </div>
        )
    }
}

export default Disk