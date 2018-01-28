import React from 'react'
import { render } from 'react-dom'
import GetRDate from "./GetRData"
import { Form, Icon, Input, Button, Checkbox } from 'antd';
const FormItem = Form.Item;
require('./login.css');

export class Login extends React.Component {
  handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        console.log('Received values of form: ', values);
      }
      var data = "name=" + values['userName'] +
        "&pwd=" + values['password']

      console.log(data)
      /*网络请求的配置*/
      var opts = {
        method: "post",
        headers: {
          "Content-type": "application/x-www-form-urlencoded"
        },
        body: data
      }
      fetch("/api/login", opts)
        .then((response) => {
          return response.text();
        }).then((responseText) => {
          console.log(responseText)
          this.setState({ remoteResult: responseText })
          return responseText;
        }).catch((error) => {
          alert(error)
        });


    });

  }
  render = () => {
    const { getFieldDecorator } = this.props.form;
    return (
      <Form onSubmit={this.handleSubmit} className="login-form" >
        <FormItem>
          {getFieldDecorator('userName', {
            rules: [{ required: true, message: 'Please input your username!' }],
          })(
            <Input prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />} placeholder="Username" />
            )}
        </FormItem>
        <FormItem>
          {getFieldDecorator('password', {
            rules: [{ required: true, message: 'Please input your Password!' }],
          })(
            <Input prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />} type="password" placeholder="Password" />
            )}
        </FormItem>
        <FormItem>
          {getFieldDecorator('remember', {
            valuePropName: 'checked',
            initialValue: true,
          })(
            <Checkbox>Remember me</Checkbox>
            )}
          <a className="login-form-forgot" href="">Forgot password</a>
          <Button type="primary" htmlType="submit" className="login-form-button">
            Log in
          </Button>
          Or <a href="">register now!</a>
        </FormItem>
      </Form>
    );
  }
}

export const WrappedNormalLoginForm = Form.create()(Login);

// render(<WrappedNormalLoginForm />, document.getElementById("root"));