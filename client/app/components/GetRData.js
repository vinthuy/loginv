import React, { Component } from 'react'

class GetRDate extends React.Component {
    static defaultProps = {
        remoteUrl: ''
    }
    propTypes = {
        remoteUrl: React.PropTypes.string.isRequired
    }

    state = {
        remoteResult: ''
    };
    getRequest = (url) => {
        /*网络请求的配置*/
        var opts = {
            method: "GET"
        }
        fetch(url, opts)
            .then((response) => {
                return response.text();
            }).then((responseText) => {
                console.log(responseText)
                this.setState({ remoteResult: responseText })
                return responseText;
            }).catch((error) => {
                alert(error)
            });

    }
    postRequest = (url, data) => {
        /*网络请求的配置*/
        var opts = {
            method: "POST",
            data: data
        }
        fetch(url, opts)
            .then((response) => {
                return response.text();
            }).then((responseText) => {
                console.log(responseText)
                this.setState({ remoteResult: responseText })
                return responseText;
            }).catch((error) => {
                alert(error)
            });

    }
    componentDidMount() {
        this.getRequest(this.props.remoteUrl)
    }
    render() {
        return (
            <div>remote:{this.state.remoteResult}</div>
        );
    }
}

export default GetRDate //default export only one .mutli can import {}