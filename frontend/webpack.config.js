var webpack = require('webpack');
module.exports = {
  // 表示入口文件
  entry: "./app/main.js",
  // 表示输出文件
  output: {
    path: __dirname,
    filename: "./dist/bundle.js"
  },
  // 表示这个依赖项是外部lib，遇到require它不需要编译，
  // 且在浏览器端对应window.React
  //externals: {
  //  'react': 'window.React'
  //},
  //
  resolve: {
    extensions: ['', '.js', '.jsx']
  },
  plugins: [
    new webpack.optimize.UglifyJsPlugin({
      output: {
        comments: false,
      },
      compress: {
        warnings: false
      }
    }),
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: JSON.stringify(process.env.NODE_ENV),
      },
    }),
  ],
  module: {
    loaders: [{
      test: /\.(js|jsx)$/,
      loader: 'babel-loader',
      exclude: /node_modules/,
      query: {
        presets: ['es2015', 'react']
      }
    },
      {
        test: /\.css$/,
        loader: "style!css"
      },
    ],
  }
};
