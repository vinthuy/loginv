const webpack = require('webpack')
module.exports = {

    entry: __dirname + "/app/index.js",//已多次提及的唯一入口文件
    output: {
        path: __dirname + "/public",//打包后的文件存放的地方
        filename: "bundle.js"//打包后输出文件的文件名
    },

    devServer: {
        contentBase: "./public",//本地服务器所加载的页面所在的目录
        historyApiFallback: true,//不跳转
        inline: true,//实时刷新
        proxy: {
            '/api': {
                target: 'http://localhost:8080',
                pathRewrite: { '^/api': '/' },
                changeOrigin: true
            }
        }
    },
    module: {

        rules: [
            {
                test: /(\.jsx|\.js)$/,
                use: {
                    loader: "babel-loader",
                    options: {
                        presets: [
                            "env", "react", "es2015", "stage-0"
                        ]
                    }
                },
                exclude: /node_modules/
            },
            {
                test: /\.(less|css)$/,
                use: [
                    {
                        loader: "style-loader"
                    }, {
                        loader: "css-loader",
                        // options: {
                        //     modules: true
                        // }
                    }, {
                        loader: "postcss-loader"
                    }, {
                        loader: "less-loader"
                    }
                ]
            }
        ]
    },
    plugins: [
        new webpack.BannerPlugin('版权所有，翻版必究'),
        new webpack.HotModuleReplacementPlugin()//热加载插件
    ],
}
