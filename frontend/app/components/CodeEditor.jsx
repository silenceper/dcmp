import React from 'react';
import Codemirror from 'react-codemirror';

class CodeEditor extends React.Component {
  render() {
    var options = {
      lineNumbers: true
    };
    return (
      <Codemirror value={this.props.code} onChange={this.props.updateCode} options={options} />
    )
  }
}

export default CodeEditor;
