<template>
  <div class="player-view"></div>
</template>

<script setup>
import Artplayer from 'artplayer';
import {onMounted} from "vue";
const props = defineProps({
  item:{
    type:Object,
    default:{}
  }
})
const initArtPlayerFunc = () => {
  new Artplayer({
    container: '.player-view',
    url: props.item.url,
    muted: false,
    autoplay: false,
    pip: true,
    autoSize: true,
    autoMini: true,
    fullscreen: true,
    fullscreenWeb: true,
    whitelist: ['*']
  });
}

onMounted( () => {
  initArtPlayerFunc()
})

</script>

<style scoped>
.player-view{
  height: 100%;
  width: 100%;
  position: relative;
  background-color: #000000;
  z-index: 1;
}
</style>
