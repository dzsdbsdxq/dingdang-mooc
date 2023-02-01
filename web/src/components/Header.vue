<template>
  <div class="header">
    <div class="container">
      <div class="menu fhNav">
        <ul class="inner-nav clear">
          <li><a href="#" title="首页"><el-icon size="20"><HomeFilled /></el-icon>首页</a></li>
          <li><a target="_blank" href="https://www.haoba.cc/s/about.html" title="作者"><el-icon size="20"><UserFilled /></el-icon>作者</a></li>
          <li><a target="_blank" href="https://www.haoba.cc/" title="博客"><el-icon size="20"><Promotion /></el-icon>博客</a></li>
          <li><a @click="lincenceFunc" href="javascript:void(0);" title="免责"><el-icon size="20"><WarningFilled /></el-icon>免责</a></li>
        </ul>
      </div>
    </div>
    <el-dialog
        v-model="dialogVisible"
        title="免责声明"
        :append-to-body="true"
        width="30%"
    >
      <div>本站只提供WEB页面服务，本站不存储、不制作任何视频，不承担任何由于内容的合法性及健康性所引起的争议和法律责任。
        若本站收录内容侵犯了您的权益，请附说明联系邮箱，本站将第一时间处理。
        如有内容侵犯了你的权益，请联系删除相关内容。</div>
      <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="dialogVisible = false">
          知道了
        </el-button>
      </span>
      </template>
    </el-dialog>
  </div>

</template>

<script setup>
import {ref} from "vue";

const dialogVisible = ref(false);
const lincenceFunc = () => {
  dialogVisible.value = true;
}
</script>

<style scoped lang="less">
.header{
  width: 100%;
  position: absolute;
  top: 0;
  z-index: 1;
  .container{
    width:100%;
    margin: 0 auto;
    position: relative;
    .menu{
      float: left;
      position: relative;
      .bottomLine{
        position: absolute;
        left: 50px;
        bottom:0;
        background-color: #333;
        transition: all 0.3s ease;
        height: 2px;
        width: 145px;
      }
      ul{
        li{
          float: left;
          line-height: 60px;
          width: 80px;
          text-align: center;
          font-size: 14px;
          position: relative;
          list-style: none;
          a{
            font-size: 14px;
            color: #ffffff;
            display: flex;
            flex-direction: row;
            flex-wrap: wrap;
            justify-content: center;
            align-items: center;
            align-content: center;
            .el-icon{
              margin-right: 5px;
            }
            &:hover{
              color: #e83e8c;
            }
          }
          &.selectedNav{
            font-weight:bold;
          }
          &:hover{
            font-weight:bold;
            color: #e83e8c;
          }
        }
      }
    }
  }
}
</style>
