import React, { Component } from 'react'

class Fan extends Component {
    render() {
        return (
            <div className="feature">
                <p>{this.props.label}: {this.props.value} rpm</p>
            </div>
        )
    }
}

export default Fan