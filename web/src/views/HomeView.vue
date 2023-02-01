<template>
  <div class="page-container">
    <div class="main-content" v-if="state.items.length > 0">
      <WebItem @_alertVideoDialog="alertVideoDialogFunc" v-for="(item, idx) in state.items" :key="idx" :item="item" />
    </div>
    <VideoDialog
        v-if="showVideoDialog"
        @_closeVideoDialog="_closeVideoDialogFunc"
        :show-video-dialog="showVideoDialog"
        :file-path="currentPath"
    >
    </VideoDialog>
  </div>
</template>

<script setup>
import WebItem from "../components/WebItem.vue";
import {API} from "@/utils/api";
import {onMounted, reactive, ref} from "vue";
import VideoDialog from "@/components/VideoDialog";
const showVideoDialog = ref(false)
const currentPath = ref("root")
const state = reactive({
  items:[]
})
const _closeVideoDialogFunc = (value) => {
  showVideoDialog.value = value;
}
const alertVideoDialogFunc = (data) => {
  currentPath.value = data.path;
  showVideoDialog.value = true;
}
const getIndexData = () => {
  API.getIndexData({}).then( res => {
    if(res.code == 200){
      state.items = res.data;
    }
  }).catch( err => {

  })
}
onMounted( () => {
  getIndexData();
})
</script>
