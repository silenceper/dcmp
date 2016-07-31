import reqwest from "reqwest";

function processResponse(resp, cb) {
  if (resp.retCode != 0) {
    processError(resp.retMsg)
    return
  }
  cb(resp.data)
}

function processError(msg) {
  alert(msg)
}


export default {
  getKeyList: (cb) => {
    reqwest({
      url: "/dcmp/v1/key/list",
      method: "get",
      error: (err) => {
        console.log(err)
        processError("保存失败")
      },
      success: (resp) => {
        processResponse(resp, cb)
      }
    })
  },

  newFile: (key, isDir, cb) => {
    let newDir = "yes";
    if (!isDir) {
      newDir = "no";
    }
    reqwest({
      url: '/dcmp/v1/key/new',
      method: 'post',
      data: {
        key: key,
        isDir: newDir
      },
      error: (err) => {
        console.log(err)
        processError("新建失败")
      },
      success: (resp) => {
        processResponse(resp, cb)
      }
    })
  },

  deleteFile: (key, isDir, cb) => {
    console.log(isDir)
    let deleteDir = "yes";
    if (!isDir) {
      deleteDir = "no";
    }

    reqwest({
      url: '/dcmp/v1/key/delete',
      method: 'post',
      data: {
        key: key,
        isDir: deleteDir
      },
      error: (err) => {
        console.log(err)
        processError("保存失败")
      },
      success: (resp) => {
        processResponse(resp, cb)
      }
    })
  },

  saveCode: (key, code, cb) => {
    reqwest({
      url: '/dcmp/v1/key/save',
      method: 'post',
      data: {
        key: key,
        value: code
      },
      error: (err) => {
        console.log(err)
        processError("保存失败")
      },
      success: (resp) => {
        processResponse(resp, cb)
      }
    })
  }
}
