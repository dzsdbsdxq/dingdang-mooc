<template>
  <div class="footer">
    <div class="logos">
      <a title="推荐使用FireFox浏览器" target="_blank" href="https://www.mozilla.org/" class="mozilla-logo">Mozilla Firefox</a>
      <a title="推荐使用Chrome浏览器" target="_blank" href="https://chrome.google.com/" class="google-logo">Google Chrome</a>
    </div>
    <ul class="navUl">
      <li class="li">
        <a>© 2023 工作笔记</a>
        <a>Powered by 咸菜拌白饭</a>
        <a>
          <el-tooltip
              class="box-item"
              effect="dark"
              content="请将#号替换为@"
              placement="top"
          >
            联系邮箱：1335244575#qq.com
          </el-tooltip></a>
        <a href="https://github.com/dzsdbsdxq" target="_blank" class="github">Github</a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'Footer'
}
</script>

<style lang="less" scoped>
.footer{
  width: 100%;
  padding: 30px 40px;
  border-top: 1px solid #dddddd;
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  justify-content: space-between;
  align-items: center;
  align-content: center;
  .logos{
    width: 260px;
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    justify-content: flex-start;
    align-items: center;
    align-content: center;
    .mozilla-logo{
      background-image: url("~@/assets/images/firefox.svg");
      background-repeat: no-repeat;
      background-size: 90px,48px;
      overflow: hidden;
      text-indent: 120%;
      white-space: nowrap;
      display: inline-block;
      height: 38px;
      width: 100px;
      flex-shrink: 0;
    }
    .google-logo{
      margin-left: 20px;
      background-image: url("~@/assets/images/Google_Chrome_icon.svg");
      background-repeat: no-repeat;
      background-size: 90px,48px;
      overflow: hidden;
      text-indent: 120%;
      white-space: nowrap;
      display: inline-block;
      height: 32px;
      width: 100px;
      flex-shrink: 0;
    }
  }
  .navUl{
    color: #4a4a4f;
    font-size: 13px;
    font-weight: bold;
    .li{
      display: flex;
      flex-direction: row;
      flex-wrap: nowrap;
      justify-content: flex-end;
      align-items: center;
      align-content: center;
      a{
        margin: 0px 10px;
        &.github{
          background: #000;
          color: #fff;
          padding: 2px 4px;
        }
      }
    }
  }
}
</style>
