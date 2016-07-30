import React from 'react';
import { Panel, Button, SplitButton, MenuItem, DropdownButton, ButtonToolbar } from 'react-bootstrap';

import ReactDOM from 'react-dom';
import SidebarTree from './SidebarTree'
import CodeEditor from './CodeEditor'
import Request from "../common/Request"


class Dcmp extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      data: {},
      code: ""
    };
  }


  componentDidMount() {
    this.refreshTree()
  }

  refreshTree() {
    Request.getKeyList(function(treeData) {
      this.setState({
        data: treeData
      })
    }.bind(this));
  }

  handleNew(eventKey) {
    //必须选择一个文件夹
    if (!this.state.cursor) {
      alert("必须选择一个文件夹");
      return;
    }
    if (!this.state.cursor.dir) {
      alert("必须选择一个文件夹");
      return;
    }

    var inputValue = prompt("请输入需要新建的名字", "")
    if (inputValue == "") {
      alert("请输入文件名");
      return;
    }

    let isDir = false;
    if (eventKey == "dir") {
      isDir = true
    }
    Request.newFile(this.state.cursor.path + "/" + inputValue, isDir, function(data) {
      //刷新列表
      this.refreshTree()
    }.bind(this));
  }

  handleDelete() {
    if (this.state.cursor) {
      let cursorNode = this.state.cursor;
      //只允许单个删除
      if (cursorNode.dir && cursorNode.children.length != 0) {
        alert("不允许批量删除");
        return;
      }
      Request.deleteFile(cursorNode.path, cursorNode.dir, function(data) {
        this.refreshTree();
      }.bind(this))
    }
  }

  handleSave() {
    let cursorNode = this.state.cursor;
    if (cursorNode && !cursorNode.dir) {
      Request.saveCode(cursorNode.path, this.state.code, function(data) {
        cursorNode.value = data.node.value;
      }.bind(this));
    }
  }

  handleToggle(node, toggled) {
    if (this.state.cursor) {
      this.state.cursor.active = false;
    }
    node.active = true;
    if (node.children) {
      node.toggled = toggled;
    }
    this.setState({
      cursor: node
    });
    //渲染view
    this.setState({
      code: node.value
    })
  }

  handleUpdateCode(newCode) {
    this.setState({
      code: newCode
    })
  }

  render() {
    return (
      <div>
      <div className="sidebar">
        <SidebarTree data={this.state.data} onToggle={this.handleToggle.bind(this)}/>
        <div className="sidebar-button">
        <ButtonToolbar>
          <DropdownButton bsStyle="success" title="新建" dropup noCaret id="dropdown-no-caret" onSelect={this.handleNew.bind(this)}>
            <MenuItem eventKey="file">文件</MenuItem>
            <MenuItem eventKey="dir">文件夹</MenuItem>
          </DropdownButton>
        <Button bsStyle="danger" onClick={this.handleDelete.bind(this)}>删除</Button>
        </ButtonToolbar>
        </div>
      </div>
      <div className="editor">
          <CodeEditor code={this.state.code} updateCode={this.handleUpdateCode.bind(this)}/>
          <Button className="save-btn" bsStyle="primary" onClick={this.handleSave.bind(this)}>保存</Button>
      </div>
  	</div>
      );
  }
}

export default Dcmp;
