import React, { Component } from 'react'

class Volt extends Component {
    render() {
        return (
            <div className="feature">
                <p>{this.props.label}: {this.props.value} rpm</p>
            </div>
        )
    }
}

export default Volt