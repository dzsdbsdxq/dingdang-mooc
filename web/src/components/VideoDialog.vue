<template>
  <el-dialog
    :id="currentPath"
    v-model="dialogVisible"
    :destroy-on-close="true"
    :draggable="true"
    :close-on-click-modal="false"
    :show-close="false"
    :append-to-body="true"
    :lock-scroll="true"
    class="videoDialog"
    :style="{backgroundColor:'#000',borderRadius:'6px',position:'relative'}"
    width="65%"
  >
    <template #header="{close,titleId}">
      <div class="video-dialog-header">
        <div class="title" :id="titleId">{{state.curData.name}}</div>
        <div class="right">
          <el-button @click="_changeVideoPlayerAreaFunc" type="success" class="right-el playlists" size="small" link :icon="Expand">视频列表</el-button>
          <el-button @click="innerVisible = true" type="primary" class="right-el download" size="small" link :icon="VideoPlay">客户端播放</el-button>
          <el-icon @click="close" class="right-el closeDialogBtn" size="20px" color="#ffffff"><CircleClose /></el-icon>
        </div>
      </div>
    </template>
    <div class="video-area-content">
      <div class="videoArea">
<!--        <Video v-if="state.lists.length>0" :key="state.curData.file_id" :item="state.curData" @_changeVideoPlayerArea="_changeVideoPlayerAreaFunc"></Video>-->
<!--      -->
        <ArtPlayerNg v-if="state.lists.length>0" :key="state.curData.file_id" :item="state.curData"></ArtPlayerNg>
      </div>

      <div @mouseleave.stop="changePlayListsLeaveFunc" class="playLists">
        <el-card class="play-lists-box-card">
          <template #header>
            <div class="card-header" style="color: #ffffff">
              <span>播放列表</span>
            </div>
          </template>
          <ul class="drawer-list" v-if="state.lists.length>0">
            <li
                @click="changePlayItemFunc(index)"
                v-for="(item,index) in state.lists"
                :key="index"
                class="drawer-item"
                :data-is-current='state.curIndex == index ? true:false'>
              <el-image class="play-thumbnail" :src="item.thumbnail" fit="fill" />
              <div class="meta">
                <div class="title">
                  <el-icon class="icon" v-if="state.curIndex == index" color="#637dff"><VideoPlay /></el-icon>
                  <span class="filename">{{item.name}}</span>
                </div>
                <div class="duration">{{item.size}}</div>
                <div class="time">{{item.updated_at}}</div>
              </div>
            </li>
          </ul>

        </el-card>
      </div>
    </div>
    <el-dialog
        v-model="innerVisible"
        width="30%"
        title="客户端播放"
        append-to-body
        style="padding: 0"
    >
      <el-card
          style="padding: 0;"
          class="box-card"
          :body-style="{
            display:'flex',
            flexDirection:'row',
            flexWrap:'nowrap',
            justifyContent:'space-between',
            alignItems:'center',
            alignContent:'center'
            }">
        <el-link :underline="false" :href="`iina://weblink/?url=${encodeURIComponent(state.curData.url)}`" target="_blank">
          <img style="width: 40px;height: 40px;" :src="Iina" alt="iina" title="iina" />
        </el-link>
        <el-link :underline="false" :href="`intent://${encodeURIComponent(state.curData.url)}#Intent;package=com.mxtech.videoplayer.ad;S.title=video.mp4;end`" target="_blank">
          <img style="width: 40px;height: 40px;" :src="MxPlayer" alt="MxPlayer" title="MxPlayer" />
        </el-link>
        <el-link :underline="false" :href="`intent://${encodeURIComponent(state.curData.url)}#Intent;package=com.mxtech.videoplayer.ad;S.title=video.mp4;end`" target="_blank">
          <img style="width: 40px;height: 40px;" :src="MxPlayerPro" alt="MxPlayerPro" title="MxPlayerPro" />
        </el-link>
        <el-link :underline="false" :href="`nplayer://${encodeURIComponent(state.curData.url)}`" target="_blank">
          <img style="width: 40px;height: 40px;" :src="Nplayer" alt="Nplayer" title="Nplayer" />
        </el-link>
        <el-link :underline="false" :href="`potplayer://${encodeURIComponent(state.curData.url)}`" target="_blank">
          <img style="width: 40px;height: 40px;" :src="PotPlayer" alt="PotPlayer" title="PotPlayer" />
        </el-link>
        <el-link :underline="false" :href="`vlc://${encodeURIComponent(state.curData.url)}`" target="_blank">
          <img style="width: 40px;height: 40px;" :src="Vlc" alt="Vlc" title="Vlc" />
        </el-link>
      </el-card>
      <div style="height: 15px;"></div>
      <div style="width: 100%;text-align: center;margin: 10px auto;">
        <el-button @click="downloadFunc" type="primary">
          <el-icon><Promotion /></el-icon>不限速下载视频
        </el-button>
      </div>

    </el-dialog>
  </el-dialog>
</template>

<script setup>
import {computed, nextTick, onMounted, reactive, ref} from "vue";
import Video from "@/components/Video";
import {API} from "@/utils/api";
import { Expand,VideoPlay } from '@element-plus/icons-vue'
import ArtPlayerNg from "@/components/ArtPlayerNg";

import Iina from "../assets/images/iina.webp"
import MxPlayer from "../assets/images/mxplayer.webp"
import MxPlayerPro from "../assets/images/mxplayer-pro.webp"
import Nplayer from "../assets/images/nplayer.webp"
import PotPlayer from "../assets/images/potplayer.webp"
import Vlc from "../assets/images/vlc.webp"
const state = reactive({
  curIndex:0,
  curPlayUrl:"",
  curData:{},
  lists:{}
})
const em = defineEmits(["_closeVideoDialog"])
const showVideoLists = ref(false);
const innerVisible = ref(false)
const props = defineProps({
  showVideoDialog:{
    type:Boolean,
    default:false
  },
  filePath:{
    type:String,
    default:""
  }
})
const _changeVideoPlayerAreaFunc = () => {
  let playLists = document.getElementsByClassName("playLists")[0];
  if(showVideoLists.value){
    playLists.style.right = '-400px';
  } else {
    playLists.style.right = '0px';
  }
  showVideoLists.value = !showVideoLists.value;
}
const dialogVisible = computed({
  get() {
    console.log("showVideoDialog:",props.showVideoDialog);
    return props.showVideoDialog
  },
  set(val) {
    em("_closeVideoDialog",val);
  }
})
const currentPath = computed({
  get() {
    props.filePath !== 'root' && getListsData(props.filePath);
    return props.filePath
  },
  set(val) {
  }
});
/**
 * 鼠标移出视频列表，隐藏视频列表处理
 */
const changePlayListsLeaveFunc = () => {
  showVideoLists.value = false;
  document.getElementsByClassName("playLists")[0].style.right = '-400px'
}
/**
 * 点击视频列表切换播放视频
 * @param index
 */
const changePlayItemFunc = (index) => {
  state.curIndex = index;
  state.curData = state.lists[index];
}
/**
 * 获取视频列表数据
 * @param filePath
 */
const getListsData = (filePath) => {
  API.getListsData({path:filePath}).then( res => {
    if(res.code == 200){
      state.lists = res.data;
      state.curData = res.data[0];
      //state.curPlayUrl = encodeURIComponent(state.curData.url)
      //getVideoPlayerSteamFunc(res.data);
    }
  }).catch( err => {
    console.log(err);
  });
}
const downloadFunc = () => {
  console.log(state.curData);
  const a = document.createElement('a');
  a.style.display = 'none';
  a.setAttribute('target', '_blank');
  /*
   * download的属性是HTML5新增的属性
   * href属性的地址必须是非跨域的地址，如果引用的是第三方的网站或者说是前后端分离的项目(调用后台的接口)，这时download就会不起作用。
   * 此时，如果是下载浏览器无法解析的文件，例如.exe,.xlsx..那么浏览器会自动下载，但是如果使用浏览器可以解析的文件，比如.txt,.png,.pdf....浏览器就会采取预览模式
   * 所以，对于.txt,.png,.pdf等的预览功能我们就可以直接不设置download属性(前提是后端响应头的Content-Type: application/octet-stream，如果为application/pdf浏览器则会判断文件为 pdf ，自动执行预览的策略)
   */
  a.setAttribute('download', state.curData.name);
  a.href = state.curData.url;
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
}

const getVideoPlayerSteamFunc = (dataLists) => {
  API.getPlayUrl({path:dataLists[0].file_id}).then( res => {
    console.log("getVideoPlayerSteamFunc:",res);
    if(res.code == 200){
      state.lists = dataLists;
      state.curData = res.data;
    }
  }).catch( err => {

  });
}
onMounted( () => {
})


</script>

<style lang="less">
.videoDialog{
  .el-dialog__header{
    padding: 0;
    margin: 0;
    width: 100%;
    .video-dialog-header{
      display: flex;
      flex-direction: row;
      flex-wrap: nowrap;
      justify-content: space-between;
      align-items: center;
      align-content: center;
      padding: 15px 20px;
      font-size: 16px;
      color: #fff;
      .title{

      }
      .right{
        display: flex;
        flex-direction: row;
        flex-wrap: nowrap;
        justify-content: flex-end;
        align-items: center;
        align-content: center;
        .right-el{
          margin-right: 10px;
          &.playlists{
            margin-right: 0px;
          }
          &.download{
            cursor: pointer;
          }
          &.closeDialogBtn{
            margin-right: 0px;
            cursor: pointer;
          }
        }
      }

    }
  }
  .el-dialog__body{
    padding: 0;
    z-index: 0;
    position: relative;
    overflow: hidden;
    .video-area-content{
      height: 550px;
      position: relative;
      .videoArea{
        width: 100%;
        height: 100%;
        &.move{
          -webkit-animation: mymove .6s; /* Chrome, Safari, Opera */
          animation: mymove .6s;
          animation-iteration-count: 1;
          -webkit-animation-iteration-count: 1;
        }
      }
      .playLists{
        position: absolute;
        height: calc(100% - 50px);
        width: 400px;
        top:0px;
        right: -400px;
        background: transparent!important;
        z-index: 999;
        overflow: scroll;
        transition: 0.5s;
        .play-lists-box-card{
          height: 100%;
          background-color: #313136;
          border: none;
          padding: 0;
          border-radius: 0;
          .el-card__header{
            border-bottom: 1px solid #84858d33;
            padding: 10px;
          }
          .el-card__body{
            padding: 0;
            overflow: scroll;
            height: 90%;
            .drawer-list{
              position: relative;
              padding: 0px;
              overflow: scroll;
              height: 100%;
              .drawer-item {
                position: relative;
                display: flex;
                justify-content: space-between;
                align-items: center;
                padding: 8px;
                border-radius: 5px;
                user-select: none;
                cursor: pointer;
                margin-bottom: 2px;
                &[data-is-current='false']:hover {
                  background-color: #84858d33;
                }
                &[data-is-current='true'] {
                  background-color: #84858d33;
                }
                .play-thumbnail{
                  width: 106px;
                  height: 60px;
                  margin-right: 12px;
                  background-color: #000000;
                  -ms-flex-negative: 0;
                  flex-shrink: 0;
                  background-size: cover;
                }
                .meta{
                  position: relative;
                  flex-grow: 1;
                  width: calc(100% - 106px - 8px * 2 - 12px);
                  .title {
                    display: flex;
                    flex-direction: row;
                    flex-wrap: nowrap;
                    justify-content: flex-start;
                    align-items: center;
                    align-content: center;
                    .icon {
                      font-size: 20px;
                      margin-right: 8px;
                    }

                    .filename {
                      font-size: 14px;
                      line-height: 1.5;
                      color: #ffffff;
                      max-width: 100%;
                      overflow: hidden;
                      white-space: nowrap;
                      -o-text-overflow: ellipsis;
                      text-overflow: ellipsis;
                    }
                  }
                  .duration{
                    font-size: 10px;
                    line-height: 1.6;
                    color: #ffffff;
                    opacity: .64;
                    max-width: 100%;
                    overflow: hidden;
                    white-space: nowrap;
                    -o-text-overflow: ellipsis;
                    text-overflow: ellipsis;
                  }
                  .time{
                    font-size: 12px;
                    line-height: 1.6;
                    color: #ffffffb8;
                    opacity: .64;
                    max-width: 100%;
                    overflow: hidden;
                    white-space: nowrap;
                    -o-text-overflow: ellipsis;
                    text-overflow: ellipsis;
                  }
                }
              }
            }
          }
        }
      }
    }
  }

}
</style>
