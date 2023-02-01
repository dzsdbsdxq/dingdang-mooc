<template>
  <el-container>
    <el-header>
      <Header></Header>
      <div class="main-title">
        <h1 class="h-title">叮当慕课™</h1>
        <p class="p-title">
          撸码的真谛便是让你去感受代码的故事，让你感同身受，也别忘了不要“人没了，™代码还在”哦！
         </p>
      </div>
    </el-header>
    <el-main>
      <div class="notice">
        <span class="tips">视频无法播放或者黑屏等异常情况，请使用客户端播放或者下载本地播放</span>
        <span class="https">拒绝流量劫持，全面使用 HTTPS！</span>
      </div>
      <router-view/>
    </el-main>
    <el-footer>
      <Footer />
    </el-footer>
  </el-container>

</template>

<style lang="less">
.el-container{
  background: url("~@/assets/images/bg.img.png");
  .el-header{
    padding: 0;
    height: 300px;
    background: url("https://www.haoba.cc/upload/2023/01/wwu-1900x300.jpg") no-repeat;
    position: relative;
    .main-title{
      position: absolute;
      left: 0;
      right: 0;
      top: 80px;
      margin: auto;
      display: inline-block;
      text-align: center;
      .h-title{
        font-size: 43px;
        color: #ffffff;
        text-shadow: 5px 5px 5px rgb(0 0 0 / 15%);
      }
      .p-title{
        font-size: 18px;
        color: #ffffff;
        margin-top: 15px;
        text-shadow: 5px 5px 5px rgb(0 0 0 / 15%);
      }
    }
  }
  .el-main{
    width: 80%;
    margin: 0 auto;
    .notice{
      display: flex;
      flex-direction: row;
      flex-wrap: nowrap;
      justify-content: space-between;
      align-items: center;
      align-content: center;
      margin: 10px auto;
      .tips{
        color:indianred;
      }
      .https{
        color:green;
        cursor: crosshair;
      }
    }
  }
  .el-footer{
    padding: 0;
    height: auto;
  }

}
</style>
<script>
import Header from "@/components/Header";
import Footer from "@/components/Footer";
export default {
  components: {Header,Footer}
}
</script>

