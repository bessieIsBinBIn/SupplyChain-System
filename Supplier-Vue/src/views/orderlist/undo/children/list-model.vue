<template>
  <div>
    <el-table
      ref="multipleTable"
      class="my-table"
      :data="tableData"
      style="width: 100%"
      @sort-change="doSortChange"
      @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="50"></el-table-column>
      <el-table-column type="index" label="ID"></el-table-column>

      <el-table-column property="buyer" align="center" label="发起企业"  >
      </el-table-column>

      <el-table-column property="goodsname" align="center" label="对应货品"  >
      </el-table-column>

      <el-table-column property="createtime" align="center" label="发起时间" sortable="custom"  >
      </el-table-column>
      
      <el-table-column property="deadline" align="center" label="截止时间" sortable="custom" >
      </el-table-column>

      

      <!-- 可操作 -->
      <el-table-column align="center" width="220" label="操作">
        <template slot-scope="scope">
          <el-button plain size="mini" @click="clickFileName(scope.row)">查看详情</el-button>
          <el-button plain size="mini" @click="delOne(scope.row)">拒绝</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 素材详情时的弹窗 -->
    <el-dialog title="素材详情"
      class="detail-dialog"
      :visible.sync="showDetailDialog"
      :close-on-click-modal="false"
      :close-on-press-escape="false">
      <el-form class="detail-form" :model="detailForm">
        <el-form-item label="发起企业：" prop="buyer" label-width="100px">
          <span>{{detailForm.buyer}}</span>
        </el-form-item>
        <el-form-item label="对应货品：" prop="goodsname" label-width="100px">
          <span>{{detailForm.goodsname}}</span>
        </el-form-item>
        <el-form-item label="货品数量：" prop="goodsnum" label-width="100px">
          <span>{{detailForm.goodsnum}}</span>
        </el-form-item>
        <el-form-item label="订单总金额：" prop="orderamount" label-width="100px">
          <span>{{detailForm.orderamount}}</span>
        </el-form-item>
        <el-form-item label="发起时间：" prop="createtime" label-width="100px">
          <span>{{detailForm.createtime}}</span>
        </el-form-item>
        <el-form-item label="截止时间：" prop="deadline" label-width="100px">
          <span>{{detailForm.deadline }}</span>
        </el-form-item>
        <el-form-item label="订单备注：" prop="remark" label-width="100px">
          <span>{{detailForm.remark}}</span>
        </el-form-item>
      </el-form>
      <div slot="footer" class="">
        <el-button @click="accept()">接受</el-button>
        <el-button @click="delOne()">删 除</el-button>
      </div>
    </el-dialog>

    <!-- 图片查看，大图弹窗 -->
    <el-dialog
      class="img-dialog"
      title="图片详情"
      top="10vh"
      :visible.sync="showImgDialog"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :show-close="true">
      <el-image :src="imgDialogSrc" fit="contain">
        <div slot="error" class="image-slot">
          <i class="el-icon-picture-outline" style="font-size: 100px"></i>
        </div>
        <div slot="placeholder" class="image-slot">
          加载中<span class="dot">...</span>
        </div>
      </el-image>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary"
          @click="downloadImg"
          size="mini">下&nbsp;&nbsp;&nbsp;载</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
import { getUndoTable } from '@/api/excel'
export default {
  props: {
    nowPid: {},
    nowLevel: {},
    // tableDatas: {
    //   type: Array,
    //   default: () => {
    //     return []
    //   }
    // },
    // 勾选情况
    selectionArr: {
      type: Array,
      default: () => {
        return []
      }
    }
  },
  data() {
    return {
      // 素材详情
      showDetailDialog: false,
      // 素材详情里面的内容
      detailForm: {},
      tableData:[],
      multipleSelection: this.selectionArr,
      /* 显示图片弹窗 */
      showImgDialog: false,
      imgIdDialog: 0,
      imgDialogSrc: ''
    }
  },
  watch: {
    // tableDatas(val) {
    //   this.tableData = val
    // }
  },
  methods: {
    accept() {

    },
    // 排序
    doSortChange({ column, prop, order }) {
      let _order
      if(order === 'ascending') {
        _order = 0
      } else if(order === 'descending') {
        _order = 1
      } else {
        prop = ''
        _order = ''
      }
      // 如果需要全库数据的一个整体排序
      // 需要调用后台接口
      this.$emit('doSort', {
        prop,
        order: _order
      })
    },
    delOne(row) {
      let _arr = []
      _arr.push(row)
      this.$refs.multipleTable.clearSelection()
      this.toggleSelection(_arr)
      this.$emit('delOne', row.id)
    },
    /* 勾选项同步 */
    toggleSelection(rows) {
      if(rows && rows.length > 0) {
        rows.forEach(row => {
          this.$refs.multipleTable.toggleRowSelection(row, true)
        })
      } else {
        this.$refs.multipleTable.clearSelection()
      }
    },
    // 表格勾选项
    handleSelectionChange(val) {
      // this.tableData.forEach(item => {
      //   item.hasChecked = false
      // })
      // if(val.length > 0) {
      //   val.forEach(item => {
      //     this.tableData.forEach((v, i, _this) => {
      //       if(v.id === row.id) {
      //         v.hasChecked = true
      //       }
      //     })
      //   })
      // }
      this.multipleSelection = val
      this.$parent.$data.checkList = val
      console.log(this.$parent.$data.checkList)
      // this.$parent.$data.fileList = this.tableData
    },
    // 点击文件名
    clickFileName(row) {
      // 如果是目录
      if(row.materialType === 0) {
        this.$emit('goFolder', row)
      } else if(row.materialType === 3) { // 如果是图片
        this.imgDialogSrc = row.thumbnailUrl
        this.imgIdDialog = row.id
        this.showImgDialog = true
      } else {
        // 显示详情
        this.detailForm = row
        this.showDetailDialog = true
      }
    },
    getData() {
      getUndoTable().then(res => {
        console.log(res)
        this.tableData = res.data
      })
    },
  },
  created() {
    this.getData()
    this.multipleSelection = this.selectionArr
  },
  mounted() {
    this.toggleSelection(this.multipleSelection)
  }
}
</script>

<style scoped>
.my-table /deep/ th {
  background: transparent !important;
  font-weight: normal;
}
/*新建文件夹*/
.edit-input /deep/ .el-input__inner{
  padding: 0 8px;
}
.file-img {
  display: inline-block;
  vertical-align: middle;
  width: 20px;
  height: 20px;
  background-repeat: no-repeat;
  background-position: center;
  background-size: 100% 100%;
}
.file-img.folder {
  background-image: url('~@/assets/images/test/icon_folder_24_35.png');
}
.file-img.video {
  background-image: url('~@/assets/images/test/icon_video_24_94.png');
}
.file-img.word {
  background-image: url('~@/assets/images/test/icon_text.png');
}
.file-img.word.pdf {
  background-position: 0 0;
  background-size: auto 100%;
  background-image: url('~@/assets/images/test/icon_pdf_24_92.png');
}
.file-img.word.code {
  background-position: 0 0;
  background-size: auto 100%;
  background-image: url('~@/assets/images/test/icon_code.png');
}

.file-img.img {
  background-image: url('~@/assets/images/test/icon_picture_24_94.png')
}
.file-img.news {
  background-position: 0 0;
  background-size: auto 100%;
  background-image: url('~@/assets/images/test/icon_misc.png');
}
.btns {
  width: 22px;
  height: 25px;
  border-color: #409EFF;
  color: #409EFF;
  padding: 0 !important;
}
.btns.cancel {
  margin-left: 4px;
}


/* 详情时的弹窗*/
.detail-dialog /deep/ .el-dialog {
  max-width: 550px;
  border-radius: 4px;
}
.detail-dialog /deep/ .el-dialog__header {
  padding: 10px;
  background: -moz-linear-gradient(#F8F9FD 0%, #DEE0E1 100%);
  background: -webkit-linear-gradient(#F8F9FD 0%, #DEE0E1 100%);
  background: -o-linear-gradient(#F8F9FD 0%, #DEE0E1 100%);
  background: -ms-linear-gradient(#F8F9FD 0%, #DEE0E1 100%);
  background: linear-gradient(#F8F9FD 0%, #DEE0E1 100%);
  border-radius: 4px;
}
.detail-dialog /deep/ .el-dialog__footer {
  padding: 6px 0;
  background-color: #fafafa;
  border-top: 1px solid #eee;
  text-align: center;
  border-bottom-left-radius: 4px;
  border-bottom-right-radius: 4px;
}
.detail-dialog /deep/ .el-dialog__title {
  font-size: 15px;
  font-weight: bold;
}
.detail-dialog /deep/ .el-dialog__close {
  position: relative;
  top: -10px;
  background-color: #409EFF;
  color: #fff;
  border-radius: 2px;
  padding: 1px;
}

/* 编辑 */
.edit-dialog /deep/ .el-dialog {
  min-width: 750px;
  max-width: 900px;
  border-radius: 4px;
}
.edit-dialog.edit-folder-dialog /deep/ .el-dialog {
  width: 550px;
  max-width: 550px;
  min-width: 550px;
}
.edit-dialog /deep/ .el-dialog__header {
  padding: 10px;
  background: -moz-linear-gradient(#F8F9FD 0%, #DEE0E1 100%);
  background: -webkit-linear-gradient(#F8F9FD 0%, #DEE0E1 100%);
  background: -o-linear-gradient(#F8F9FD 0%, #DEE0E1 100%);
  background: -ms-linear-gradient(#F8F9FD 0%, #DEE0E1 100%);
  background: linear-gradient(#F8F9FD 0%, #DEE0E1 100%);
  border-radius: 4px;
}
.edit-dialog /deep/ .el-dialog__footer {
  padding: 6px 0;
  background-color: #fafafa;
  border-top: 1px solid #eee;
  text-align: center;
  border-bottom-left-radius: 4px;
  border-bottom-right-radius: 4px;
}
.edit-dialog /deep/ .el-dialog__title {
  font-size: 15px;
  font-weight: bold;
}
.edit-dialog /deep/ .el-dialog__close {
  position: relative;
  top: -10px;
  background-color: #409EFF;
  color: #fff;
  border-radius: 2px;
  padding: 1px;
}
.edit-dialog .del-tips {
  text-align: center;
  font-size: 15px;
  color: #333;
  margin-top: 20px;
  margin-bottom: 25px;
}

/*批量编辑*/
.share-theme {
  max-height: 150px;
  overflow-x: hidden;
  overflow-y: auto;
}
.share-org {
  display: flex;
  margin-top: 8px;
}
.share-org .left {
  display: inline-block;
  width: 280px;
  padding: 8px 0px;
  border: 1px solid #eee;
  border-radius: 4px;
  height: 250px;
  overflow-y: auto;
  overflow-x: hidden;
}
.share-org .right {
  display: inline-block;
  width: 180px;
  padding: 8px 0px;
  border: 1px solid #eee;
  border-radius: 4px;
  margin-left: 15px;
  height: 250px;
  overflow-y: auto;
  overflow-x: hidden;
}
.select-tag {
  position: relative;
  box-sizing: border-box;
  display: inline-block;
  width: 180px;
  border-bottom: 1px solid #eee;
  padding: 0 30px 0 15px;
  color: #333;
  /*margin-bottom: 5px;*/
}
.el-icon-close {
  position: absolute;
  left: 150px;
  top: 7px;
  cursor: pointer;
  font-size: 18px;
  color: red;
}
.el-icon-close:hover {
  border: 1px solid red;
  border-radius: 50%;
}
.topnav_box::-webkit-scrollbar {
  width: 5px;
  height:10px;
  background-color: #fff;
}  
.topnav_box::-webkit-scrollbar-track {
  -webkit-box-shadow: inset 0 0 6px rgba(0,0,0,0.3);
  border-radius: 10px;
  background-color: #eee;
}
.topnav_box::-webkit-scrollbar-thumb {
  border-radius: 10px;
  -webkit-box-shadow: inset 0 0 6px rgba(0,0,0,.3);
  background-color: #a5d6a7;
}

/*大图显示弹窗*/
.img-dialog /deep/ .el-dialog {
  max-width: 800px;
  border-radius: 4px;
}
.img-dialog /deep/ .el-dialog__header {
  padding: 10px;
  background: -moz-linear-gradient(#F8F9FD 0%, #DEE0E1 100%);
  background: -webkit-linear-gradient(#F8F9FD 0%, #DEE0E1 100%);
  background: -o-linear-gradient(#F8F9FD 0%, #DEE0E1 100%);
  background: -ms-linear-gradient(#F8F9FD 0%, #DEE0E1 100%);
  background: linear-gradient(#F8F9FD 0%, #DEE0E1 100%);
  border-radius: 4px;
}
.img-dialog /deep/ .el-dialog__body {
  text-align: center;
  padding: 10px;
}
.img-dialog /deep/ .el-dialog__body .el-image__inner {
  max-height: 60vh;
}
.img-dialog /deep/ .el-dialog__footer {
  padding: 6px 0;
  background-color: #fafafa;
  border-top: 1px solid #eee;
  text-align: center;
  border-bottom-left-radius: 4px;
  border-bottom-right-radius: 4px;
}
.img-dialog /deep/ .el-dialog__title {
  font-size: 15px;
  font-weight: bold;
}
.img-dialog /deep/ .el-dialog__close {
  position: relative;
  top: -10px;
  background-color: #409EFF;
  color: #fff;
  border-radius: 2px;
  padding: 1px;
}

.file-name /deep/ .el-link--inner {
  display: inline-block;
  width: 145px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

/deep/ .el-checkbox__inner {
  transform: scale(1.1)
}
</style>