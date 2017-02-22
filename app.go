package main

import (
	"strings"

	"github.com/coreos/etcd/client"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/dcmp/etcd"
	"github.com/silenceper/dcmp/util"
)

type TreeNode struct {
	Name     string      `json:"name"`
	Dir      bool        `json:"dir"`
	Toggled  bool        `json:"toggled"`
	Path     string      `json:"path"`
	Value    string      `json:"value"`
	Children []*TreeNode `json:"children"`
}

func doIndex(c *gin.Context) {
	c.Redirect(302, "/frontend")
}

func doKeyList(c *gin.Context) {
	resp, err := etcd.GetKeyList(cfg.BasePath)
	if err != nil {
		util.RenderError(c, -1, err.Error())
		return
	}
	treeNodes := formatEtcdNodes(resp.Node)
	util.RenderSuccess(c, treeNodes)
}

func doKeyNew(c *gin.Context) {
	key, keyExists := c.GetPostForm("key")
	if !keyExists || key == "" {
		util.RenderError(c, -1, "参数错误")
		return
	}

	isDir, exists := c.GetPostForm("isDir")
	if !exists || isDir == "" {
		util.RenderError(c, -1, "参数错误")
		return
	}

	opts := &client.SetOptions{}
	if isDir == "yes" {
		opts.Dir = true
	}
	resp, err := etcd.SetKey(key, "", opts)
	if err != nil {
		util.RenderError(c, -1, err.Error())
		return
	}
	util.RenderSuccess(c, resp)
}

func doKeySave(c *gin.Context) {
	key, keyExists := c.GetPostForm("key")
	if !keyExists || key == "" {
		util.RenderError(c, -1, "参数错误")
		return
	}
	value, valueExists := c.GetPostForm("value")
	if !valueExists {
		util.RenderError(c, -1, "参数错误")
		return
	}

	resp, err := etcd.UpdateKey(key, value)
	if err != nil {
		util.RenderError(c, -1, err.Error())
		return
	}
	util.RenderSuccess(c, resp)
}

func doKeyDelete(c *gin.Context) {
	key, keyExists := c.GetPostForm("key")
	if !keyExists || key == "" {
		util.RenderError(c, -1, "参数错误")
		return
	}

	isDir, exists := c.GetPostForm("isDir")
	if !exists || isDir == "" {
		util.RenderError(c, -1, "参数错误")
		return
	}

	opts := &client.DeleteOptions{}
	if isDir == "yes" {
		opts.Dir = true
	}
	resp, err := etcd.DeleteKey(key, opts)
	if err != nil {
		util.RenderError(c, -1, err.Error())
		return
	}
	util.RenderSuccess(c, resp)
}

func formatEtcdNodes(node *client.Node) *TreeNode {
	treeNode := &TreeNode{}
	arr := strings.Split(node.Key, "/")
	count := len(arr)
	if count < 1 {
		count = 1
	}
	treeNode.Name = arr[count-1]
	treeNode.Path = node.Key
	treeNode.Dir = node.Dir
	treeNode.Value = node.Value
	treeNode.Toggled = true

	for _, v := range node.Nodes {
		treeNode.Children = append(treeNode.Children, formatEtcdNodes(v))
	}
	//必须返回一个空数组供前端渲染
	if treeNode.Dir && len(treeNode.Children) == 0 {
		treeNode.Children = []*TreeNode{}
	}
	return treeNode
}
