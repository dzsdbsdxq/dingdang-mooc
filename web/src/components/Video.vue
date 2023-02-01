<template>
  <div class="item-box">
    <video ref="videoPlayer" class="video-js"></video>
  </div>
</template>


<script setup>
import {ref, onMounted, onUnmounted, onUpdated, getCurrentInstance, watchEffect, computed} from "vue";
import videojs from 'video.js';
import "video.js/dist/video-js.css"

const videoPlayer = ref(null)

const myPlayer = ref(null)
const showVideoLists = ref(false)

const em = defineEmits(["_changeVideoPlayerArea"])

const props = defineProps({
  item:{
    type:Object,
    default:{}
  }
})

const initVideoJsPlayerFunc = () => {
  console.log("initVideoJsPlayerFunc:",props.item);
  myPlayer.value = videojs(videoPlayer.value, {
    poster: props.item.thumbnail,
    autoplay:true,
    controls: true,
    sources: [
      {
        src: props.item.url,
        type: "video/mp4",
      }
    ]
  }, () => {
    myPlayer.value.play();
  })
  let myButton = myPlayer.value.controlBar.addChild("button");
  let myButtonDom = myButton.el();
  myButtonDom.innerHTML = "列表";
  myButtonDom.onclick = function(){
    showVideoLists.value = !showVideoLists.value;
    em("_changeVideoPlayerArea",{type:1,va:showVideoLists.value})
  }
}
const clickTest = () => {
  console.log("clickTest");
}

onMounted(()=>{
  initVideoJsPlayerFunc();
})

</script>
<style lang="less">
.item-box {
  height: 100%;
  width: 100%;
  position: relative;
  background-color: #000000;
  z-index: 1;
  .video-js{
    width: 100%;
    height: 100%;
    border-radius: 6px;
    font-size: 14px;
    button{
      &.vjs-big-play-button{
        top: 0;
        left: 0;
        bottom: 0;
        right: 0;
        margin: auto;
        width: 2em;
        height: 2em;
        border-radius: 50%;
        font-size: 3em;
        line-height: 2em;
        .vjs-icon-placeholder{
          &:before{
            font-size: 1.5em;
          }
        }
      }
    }

  }
}
</style>
