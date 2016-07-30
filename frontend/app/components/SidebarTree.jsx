import { Treebeard, decorators } from 'react-treebeard';
import React from 'react';

class SidebarTree extends React.Component {
  constructor(props) {
    super(props);
  }
  render() {
    return (
      <div className="sidebar-tree">
        <Treebeard
      data={this.props.data}
      onToggle={this.props.onToggle}
      decorators={decorators}
      />
      </div>
      );
  }
}

decorators.Header = (props) => {
  const style = props.style;
  const iconType = props.node.children ? 'folder' : 'file-text';
  const iconClass = `fa fa-${iconType}`;
  const iconStyle = {
    marginRight: '5px'
  };
  return (
    <div style={style.base}>
            <div style={style.title}>
                <i className={iconClass} style={iconStyle}/>
                {props.node.name}
            </div>
        </div>
    );
};

export default SidebarTree;
