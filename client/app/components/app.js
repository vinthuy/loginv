import Main from "./main"
import { render } from "react-dom";
import React, { Component } from "react"

export class App1 extends React.Component{
    render = () => {
        return (<div>
            <Main />
        </div>)
    }

}
// render(<App1 />, document.body)